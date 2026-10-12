package tasking

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"go-practice-bank/internal/sources"
	"go-practice-bank/internal/toolchainpolicy"
	"go-practice-bank/internal/triage"
)

var sentenceEndPattern = regexp.MustCompile(`[.!?](?:\s|$)`)

type checker struct {
	report QualityReport
}

func (s *CheckSession) checkLocated(located LocatedTask) (QualityReport, error) {
	root := s.root
	if _, err := os.Lstat(filepath.Join(located.Path, "lifecycle.yaml")); errors.Is(err, os.ErrNotExist) && located.State == StateReady {
		return s.checkLegacy(located)
	}
	if located.State == StateRejected || located.State == StateRetired {
		return checkInactive(root, located)
	}
	c := checker{report: QualityReport{
		SchemaVersion: 1, TaskID: located.ID, State: located.State,
		FailingGateIDs: []string{}, StarterStatus: "not-applicable", ReferenceStatus: "not-applicable", RaceStatus: "not-applicable",
		Mutants: MutantSummary{Survived: []string{}},
	}}
	data, loadProblems := loadPackage(root, located)
	c.gate("package-metadata", loadProblems)

	contractProblems, layout := checkPackageContract(root, data)
	c.gate("package-contract", contractProblems)

	metadataProblems := checkMetadata(root, data)
	c.gate("taxonomy-and-lifecycle", metadataProblems)

	promptProblems := checkPrompt(data.Prompt)
	c.gate("prompt-contract", promptProblems)

	provenanceProblems := checkProvenance(root, data)
	c.gate("provenance-license", provenanceProblems)

	reviewProblems := checkReview(data.Review)
	c.gate("review-evidence", reviewProblems)

	apiProblems := checkPublicAPI(layout)
	if data.ContractVersion == 2 {
		apiProblems = checkPublicAPIV2(data, layout)
	}
	c.gate("public-api-stability", apiProblems)

	determinismProblems := checkTestSources(layout)
	if data.ContractVersion == 2 {
		determinismProblems = checkTestSourcesV2(data, layout)
	}
	c.gate("deterministic-tests", determinismProblems)

	c.gate("task-type-diversity", nil)
	c.report.Gates[len(c.report.Gates)-1].Status = "not-applicable"
	c.report.Gates[len(c.report.Gates)-1].Details = []string{"repository-level distribution is reported by task-report"}

	staticPassed := len(loadProblems) == 0 && len(contractProblems) == 0 && len(metadataProblems) == 0 && len(promptProblems) == 0 && len(provenanceProblems) == 0 && len(apiProblems) == 0 && len(determinismProblems) == 0
	if !staticPassed {
		for _, id := range []string{"starter-public", "starter-full", "reference-full", "mutant-coverage"} {
			c.gate(id, []string{"execution skipped because static gates failed"})
		}
		c.notApplicable("race", "execution skipped because static gates failed")
		c.finish()
		return c.validatedReport(root)
	}
	if data.ContractVersion == 2 {
		s.executeV2(&c, data, layout)
		c.finish()
		return c.validatedReport(root)
	}

	starterPublic := s.runner.RunVariant(variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, Implementation: layout.Starter, Tests: layout.PublicTests})
	c.gate("starter-public", expectFailure("starter with public tests", starterPublic))
	starterFull := s.runner.RunVariant(variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, Implementation: layout.Starter, Tests: appendCopy(layout.PublicTests, layout.GraderTests)})
	c.gate("starter-full", expectFailure("starter with full tests", starterFull))

	referenceFirst := s.runner.RunVariant(variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, Implementation: layout.Reference, Tests: appendCopy(layout.PublicTests, layout.GraderTests)})
	referenceSecond := s.runner.RunVariant(variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, Implementation: layout.Reference, Tests: appendCopy(layout.PublicTests, layout.GraderTests)})
	var referenceProblems []string
	if !referenceFirst.Compiled {
		referenceProblems = append(referenceProblems, describeVariantFailure("reference with full tests", referenceFirst))
	} else if !referenceFirst.Passed {
		referenceProblems = append(referenceProblems, describeVariantFailure("reference with full tests", referenceFirst))
	}
	if !referenceSecond.Compiled || !referenceSecond.Passed {
		referenceProblems = append(referenceProblems, describeVariantFailure("repeated reference run", referenceSecond))
	}
	if referenceFirst.Canonical != referenceSecond.Canonical {
		referenceProblems = append(referenceProblems, "repeated reference runs have different observable test events")
	}
	c.gate("reference-full", referenceProblems)

	c.report.Mutants.Required = data.Review.MinimumMutants
	c.report.Mutants.Total = len(layout.Mutants)
	var mutantProblems []string
	for _, mutant := range layout.Mutants {
		result := s.runner.RunVariant(variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, Implementation: mutant.Files, Tests: layout.GraderTests})
		if !result.Compiled {
			mutantProblems = append(mutantProblems, describeVariantFailure("mutant "+mutant.ID, result))
			continue
		}
		if result.Passed {
			c.report.Mutants.Survived = append(c.report.Mutants.Survived, mutant.ID)
			mutantProblems = append(mutantProblems, fmt.Sprintf("mutant %s survives full tests", mutant.ID))
			continue
		}
		if result.FailureKind != variantFailureTest {
			mutantProblems = append(mutantProblems, describeVariantFailure("mutant "+mutant.ID, result))
			continue
		}
		c.report.Mutants.Detected++
	}
	if len(layout.Mutants) < data.Review.MinimumMutants {
		mutantProblems = append(mutantProblems, fmt.Sprintf("found %d mutants; review requires at least %d", len(layout.Mutants), data.Review.MinimumMutants))
	}
	sort.Strings(c.report.Mutants.Survived)
	c.gate("mutant-coverage", mutantProblems)

	if isConcurrencyTask(data.Task) {
		race := s.runner.RunVariant(variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, Race: true, Implementation: layout.Reference, Tests: appendCopy(layout.PublicTests, layout.GraderTests)})
		if race.FailureKind == variantFailureRaceUnsupported {
			c.infrastructureUnsupported("race", race.Diagnostic)
		} else {
			var raceProblems []string
			if !race.Compiled || !race.Passed {
				raceProblems = append(raceProblems, describeVariantFailure("reference go test -race with full tests", race))
			}
			c.gate("race", raceProblems)
		}
	} else {
		c.notApplicable("race", "task does not assess goroutines, channels, sync, or context")
	}

	c.finish()
	return c.validatedReport(root)
}

func (c *checker) validatedReport(root string) (QualityReport, error) {
	if err := validateSchema(root, "task-quality-report.schema.json", c.report); err != nil {
		return QualityReport{}, fmt.Errorf("generated quality report is invalid: %w", err)
	}
	return c.report, nil
}

func (c *checker) gate(id string, problems []string) {
	status := "pass"
	if len(problems) != 0 {
		status = "fail"
	}
	details := append([]string(nil), problems...)
	if details == nil {
		details = []string{}
	}
	sort.Strings(details)
	c.report.Gates = append(c.report.Gates, Gate{ID: id, Status: status, Details: details})
}

func (c *checker) notApplicable(id, detail string) {
	c.report.Gates = append(c.report.Gates, Gate{ID: id, Status: "not-applicable", Details: []string{detail}})
}

func (c *checker) infrastructureUnsupported(id, detail string) {
	c.report.Gates = append(c.report.Gates, Gate{ID: id, Status: "infrastructure-unsupported", Details: []string{detail}})
}

func (c *checker) finish() {
	c.report.TechnicalStatus = "pass"
	c.report.ReviewStatus = "not-applicable"
	for _, gate := range c.report.Gates {
		if gate.ID == "review-evidence" {
			c.report.ReviewStatus = gate.Status
			continue
		}
		if gate.Status == "fail" || gate.Status == "infrastructure-unsupported" {
			c.report.TechnicalStatus = "fail"
			c.report.FailingGateIDs = append(c.report.FailingGateIDs, gate.ID)
		}
		switch gate.ID {
		case "starter-public", "starter-full":
			if gate.Status == "fail" {
				c.report.StarterStatus = "fail"
			} else if c.report.StarterStatus != "fail" {
				c.report.StarterStatus = gate.Status
			}
		case "reference-full":
			c.report.ReferenceStatus = gate.Status
		case "race":
			c.report.RaceStatus = gate.Status
		}
	}
	if c.report.ReviewStatus == "fail" {
		c.report.FailingGateIDs = append(c.report.FailingGateIDs, "review-evidence")
	}
	c.report.FailingGateIDs = uniqueSorted(c.report.FailingGateIDs)
	switch c.report.State {
	case StateDraft:
		c.report.PromotionEligible = c.report.TechnicalStatus == "pass"
		c.report.Passed = c.report.TechnicalStatus == "pass"
	case StateInReview:
		c.report.PromotionEligible = c.report.TechnicalStatus == "pass" && c.report.ReviewStatus == "pass"
		c.report.Passed = c.report.TechnicalStatus == "pass"
	case StateReady:
		c.report.PromotionEligible = c.report.TechnicalStatus == "pass" && c.report.ReviewStatus == "pass"
		c.report.Passed = c.report.PromotionEligible
	default:
		c.report.PromotionEligible = false
		c.report.Passed = c.report.TechnicalStatus == "pass"
	}
}

type packageLayout struct {
	TaskDir               string
	GoMod                 string
	Starter               []string
	PublicTests           []string
	GraderTests           []string
	Reference             []string
	Mutants               []mutantLayout
	WorkspaceSkeleton     []mappedFile
	WorkspaceStarter      []mappedFile
	WorkspaceReference    []mappedFile
	WorkspacePublic       []mappedFile
	WorkspaceGrader       []mappedFile
	WorkspaceMutants      []workspaceMutantLayout
	VerificationReference []string
}

type mutantLayout struct {
	ID    string
	Files []string
}

type workspaceMutantLayout struct {
	ID    string
	Files []mappedFile
}

func checkPackageContract(root string, data packageData) ([]string, packageLayout) {
	if data.ContractVersion == 2 {
		return checkPackageContractV2(root, data)
	}
	return checkPackageContractV1(root, data)
}

func checkPackageContractV1(root string, data packageData) ([]string, packageLayout) {
	var problems []string
	layout := packageLayout{TaskDir: data.Located.Path, GoMod: filepath.Join(data.Located.Path, "go.mod")}
	for _, name := range []string{"TASK.md", "task.yaml", "lifecycle.yaml", "go.mod", "review/review.yaml", "review/notes.md", "hints/HINTS.md", "provenance.yaml"} {
		path := filepath.Join(data.Located.Path, filepath.FromSlash(name))
		info, err := os.Lstat(path)
		if err != nil || !info.Mode().IsRegular() {
			problems = append(problems, fmt.Sprintf("required regular file is missing: %s", name))
		} else if info.Size() == 0 {
			problems = append(problems, fmt.Sprintf("required file is empty: %s", name))
		}
	}
	if _, err := os.Lstat(filepath.Join(data.Located.Path, "go.work")); err == nil {
		problems = append(problems, "go.work is not allowed in a task package")
	}
	problems = append(problems, checkPureGoBoundary(data.Located.Path)...)
	goMod, err := os.ReadFile(layout.GoMod)
	if err == nil {
		policy, policyErr := toolchainpolicy.Load(root)
		if policyErr != nil {
			problems = append(problems, "load toolchain policy: "+policyErr.Error())
		} else {
			problems = append(problems, toolchainpolicy.ValidateGoMod(goMod, policy.ModulePrefix+data.Located.ID, policy.GoLanguageVersion, policy.ForbidToolchainDirective)...)
		}
	}
	layout.Starter, problems = collectGoDirectory(data.Located.Path, "starter", directoryImplementation, problems)
	layout.PublicTests, problems = collectGoDirectory(data.Located.Path, "tests/public", directoryTestHarness, problems)
	layout.GraderTests, problems = collectGoDirectory(data.Located.Path, "tests/grader", directoryTestHarness, problems)
	layout.Reference, problems = collectGoDirectory(data.Located.Path, "solution/reference", directoryImplementation, problems)

	mutantRoot := filepath.Join(data.Located.Path, "solution", "mutants")
	entries, err := os.ReadDir(mutantRoot)
	if err != nil {
		problems = append(problems, "required directory is missing: solution/mutants")
	} else {
		for _, entry := range entries {
			if !entry.IsDir() || !slugPattern.MatchString(entry.Name()) {
				problems = append(problems, fmt.Sprintf("invalid mutant entry: %s", entry.Name()))
				continue
			}
			files, next := collectGoDirectory(data.Located.Path, filepath.ToSlash(filepath.Join("solution/mutants", entry.Name())), directoryImplementation, problems)
			problems = next
			layout.Mutants = append(layout.Mutants, mutantLayout{ID: entry.Name(), Files: files})
		}
		sort.Slice(layout.Mutants, func(i, j int) bool { return layout.Mutants[i].ID < layout.Mutants[j].ID })
		if len(layout.Mutants) == 0 {
			problems = append(problems, "solution/mutants must contain at least one mutant directory")
		}
	}
	return uniqueSorted(problems), layout
}

func checkPureGoBoundary(taskDirectory string) []string {
	nativeExtensions := map[string]bool{
		".a": true, ".c": true, ".cc": true, ".cpp": true, ".cxx": true,
		".dll": true, ".dylib": true, ".h": true, ".hh": true, ".hpp": true,
		".m": true, ".mm": true, ".o": true, ".s": true, ".so": true,
		".sx": true, ".syso": true,
	}
	var problems []string
	err := filepath.WalkDir(taskDirectory, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(taskDirectory, path)
		if err != nil {
			return err
		}
		relative = filepath.ToSlash(relative)
		extension := strings.ToLower(filepath.Ext(entry.Name()))
		if nativeExtensions[extension] {
			problems = append(problems, fmt.Sprintf("native source or binary file is forbidden by the pure-Go boundary: %s", relative))
			return nil
		}
		if extension != ".go" {
			return nil
		}
		fileset := token.NewFileSet()
		parsed, err := parser.ParseFile(fileset, path, nil, parser.ParseComments)
		if err != nil {
			return nil // collectGoDirectory reports syntax errors with task-facing paths.
		}
		for _, imported := range parsed.Imports {
			importPath, err := strconv.Unquote(imported.Path.Value)
			if err == nil && importPath == "C" {
				problems = append(problems, fmt.Sprintf("cgo import is forbidden by the pure-Go boundary: %s", relative))
			}
		}
		for _, group := range parsed.Comments {
			for _, comment := range group.List {
				text := strings.ToLower(comment.Text)
				if strings.Contains(text, "#cgo") || strings.Contains(text, "//go:cgo_") {
					problems = append(problems, fmt.Sprintf("cgo compiler or linker directive is forbidden by the pure-Go boundary: %s", relative))
				}
			}
		}
		return nil
	})
	if err != nil {
		problems = append(problems, "inspect pure-Go boundary: "+err.Error())
	}
	return uniqueSorted(problems)
}

type directoryKind int

const (
	directoryImplementation directoryKind = iota
	directoryTestHarness
)

func collectGoDirectory(root, relative string, kind directoryKind, problems []string) ([]string, []string) {
	directory := filepath.Join(root, filepath.FromSlash(relative))
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, append(problems, fmt.Sprintf("required directory is missing: %s", relative))
	}
	var files []string
	nonTests := 0
	tests := 0
	for _, entry := range entries {
		if entry.IsDir() {
			problems = append(problems, fmt.Sprintf("nested directory is not allowed in %s: %s", relative, entry.Name()))
			continue
		}
		if filepath.Ext(entry.Name()) != ".go" {
			problems = append(problems, fmt.Sprintf("only Go files are allowed in %s: %s", relative, entry.Name()))
			continue
		}
		if strings.HasSuffix(entry.Name(), "_test.go") {
			tests++
		} else {
			nonTests++
		}
		files = append(files, filepath.Join(directory, entry.Name()))
		fileset := token.NewFileSet()
		parsed, parseErr := parser.ParseFile(fileset, filepath.Join(directory, entry.Name()), nil, parser.PackageClauseOnly|parser.ImportsOnly)
		if parseErr != nil {
			problems = append(problems, fmt.Sprintf("cannot parse %s/%s: %v", relative, entry.Name(), parseErr))
		} else if kind == directoryImplementation && parsed.Name.Name != "subject" {
			problems = append(problems, fmt.Sprintf("implementation package in %s/%s must be subject", relative, entry.Name()))
		} else if kind == directoryTestHarness && parsed.Name.Name != "subject_test" {
			problems = append(problems, fmt.Sprintf("test package in %s/%s must be subject_test", relative, entry.Name()))
		}
	}
	sort.Strings(files)
	if len(files) == 0 {
		problems = append(problems, fmt.Sprintf("%s must contain at least one Go file", relative))
	}
	if kind == directoryImplementation && tests != 0 {
		problems = append(problems, fmt.Sprintf("starter tests must live in tests/public, not %s", relative))
	}
	if kind == directoryTestHarness && tests == 0 {
		problems = append(problems, fmt.Sprintf("%s must contain a *_test.go file", relative))
	}
	if kind == directoryTestHarness && nonTests != 0 {
		problems = append(problems, fmt.Sprintf("%s may contain only *_test.go files", relative))
	}
	if kind == directoryImplementation && nonTests == 0 {
		problems = append(problems, fmt.Sprintf("%s must contain implementation files", relative))
	}
	return files, problems
}

func checkMetadata(root string, data packageData) []string {
	_, contracts, err := triage.LoadSkillContracts(root)
	return checkMetadataWithContracts(data, contracts, err)
}

func checkMetadataWithContracts(data packageData, contracts map[string]triage.SkillContract, err error) []string {
	var problems []string
	if data.Task.ID != data.Located.ID {
		problems = append(problems, "task id must match its directory")
	}
	if data.Lifecycle.State != data.Located.State {
		problems = append(problems, "lifecycle state must match its directory")
	}
	if want := qualityByState[data.Located.State]; data.Task.QualityStatus != want {
		problems = append(problems, fmt.Sprintf("quality_status must be %q in state %s", want, data.Located.State))
	}
	problems = append(problems, validateLifecycle(data.Lifecycle)...)
	if len(data.Task.NewConcepts) > 2 {
		problems = append(problems, "new_concepts must contain at most two skills for Phase 3A tasks")
	}
	if data.ContractVersion == 2 {
		problems = append(problems, checkMetadataV2(data)...)
	}
	if err != nil {
		return append(problems, "cannot load taxonomy: "+err.Error())
	}
	assessed := append([]string{data.Task.Skills.Primary}, data.Task.Skills.Secondary...)
	all := append(append([]string{}, assessed...), data.Task.Prerequisites...)
	all = append(all, data.Task.NewConcepts...)
	for _, id := range all {
		contract, ok := contracts[id]
		if !ok || !contract.Assignable || contract.Status != "active" {
			problems = append(problems, fmt.Sprintf("skill %q is not an active assignable taxonomy skill", id))
		}
	}
	for _, group := range [][]string{data.Task.Skills.Secondary, data.Task.Prerequisites, data.Task.NewConcepts} {
		if duplicate := firstDuplicate(group); duplicate != "" {
			problems = append(problems, fmt.Sprintf("duplicate skill %q", duplicate))
		}
	}
	if contains(data.Task.Skills.Secondary, data.Task.Skills.Primary) {
		problems = append(problems, "primary skill cannot also be secondary")
	}
	if overlap := firstOverlap(assessed, data.Task.Prerequisites); overlap != "" {
		problems = append(problems, fmt.Sprintf("assessed skill %q cannot also be a prerequisite", overlap))
	}
	if overlap := firstOverlap(data.Task.NewConcepts, data.Task.Prerequisites); overlap != "" {
		problems = append(problems, fmt.Sprintf("new concept %q cannot also be a prerequisite", overlap))
	}
	for _, concept := range data.Task.NewConcepts {
		if !contains(assessed, concept) {
			problems = append(problems, fmt.Sprintf("new concept %q must be primary or secondary", concept))
		}
	}
	available := append(append([]string{}, assessed...), data.Task.Prerequisites...)
	for _, assessedID := range assessed {
		if contract, ok := contracts[assessedID]; ok {
			for _, prerequisite := range contract.Prerequisites {
				if !contains(available, prerequisite) {
					problems = append(problems, fmt.Sprintf("direct prerequisite %q of %q is not declared", prerequisite, assessedID))
				}
			}
		}
	}
	return uniqueSorted(problems)
}

func validateLifecycle(lifecycle Lifecycle) []string {
	var problems []string
	if len(lifecycle.Transitions) == 0 {
		return []string{"lifecycle must contain at least one transition"}
	}
	current := State("")
	previousDate := ""
	for index, transition := range lifecycle.Transitions {
		if transition.From != current {
			problems = append(problems, fmt.Sprintf("lifecycle transition %d has non-contiguous from state", index))
		}
		if current == "" {
			if transition.To != StateDraft {
				problems = append(problems, "initial lifecycle transition must enter draft")
			}
		} else if !allowedTransition(current, transition.To) {
			problems = append(problems, fmt.Sprintf("lifecycle transition %s -> %s is not allowed", current, transition.To))
		}
		if previousDate != "" && transition.At < previousDate {
			problems = append(problems, fmt.Sprintf("lifecycle transition %d predates the previous transition", index))
		}
		previousDate = transition.At
		current = transition.To
	}
	if current != lifecycle.State {
		problems = append(problems, "lifecycle state must equal the last transition target")
	}
	if lifecycle.UpdatedAt != lifecycle.Transitions[len(lifecycle.Transitions)-1].At {
		problems = append(problems, "lifecycle updated_at must equal the last transition date")
	}
	return problems
}

func checkPrompt(content []byte) []string {
	if len(content) == 0 {
		return []string{"TASK.md is empty"}
	}
	text := string(content)
	var problems []string
	if countWords(text) > 160 {
		problems = append(problems, fmt.Sprintf("TASK.md has %d words; maximum is 160", countWords(text)))
	}
	if strings.Contains(strings.ToLower(text), "todo") {
		problems = append(problems, "TASK.md contains a TODO placeholder")
	}
	sections := markdownSections(text)
	for _, name := range []string{"Context", "API", "Requirements", "Constraints"} {
		if strings.TrimSpace(sections[name]) == "" {
			problems = append(problems, "TASK.md is missing non-empty ## "+name)
		}
	}
	context := strings.TrimSpace(sections["Context"])
	if sentenceEndPattern.FindAllStringIndex(context, -1) != nil && len(sentenceEndPattern.FindAllStringIndex(context, -1)) > 2 {
		problems = append(problems, "TASK.md context must contain at most two sentences")
	}
	if !strings.Contains(sections["API"], "```go") {
		problems = append(problems, "TASK.md API must contain an explicit fenced Go API")
	}
	requirements := countBullets(sections["Requirements"])
	if requirements < 4 || requirements > 7 {
		problems = append(problems, fmt.Sprintf("TASK.md must contain 4..7 requirements; found %d", requirements))
	}
	if countBullets(sections["Constraints"]) == 0 {
		problems = append(problems, "TASK.md must contain at least one explicit constraint")
	}
	for name := range sections {
		lower := strings.ToLower(name)
		if lower == "solution" || lower == "approach" || lower == "implementation notes" {
			problems = append(problems, "TASK.md must not explain the solution")
		}
	}
	return uniqueSorted(problems)
}

func checkProvenance(root string, data packageData) []string {
	if data.ProvenanceVersion == 2 {
		return checkProvenanceV2(root, data)
	}
	return checkProvenanceV1(root, data)
}

func checkProvenanceV1(root string, data packageData) []string {
	var problems []string
	registry, err := sources.LoadRegistry(root)
	if err != nil {
		return []string{"cannot load source registry: " + err.Error()}
	}
	source, err := registry.Find(data.Task.Source.ID)
	if err != nil {
		return []string{err.Error()}
	}
	if err := cleanRelativePath(data.Task.Source.Path); err != nil {
		problems = append(problems, "task source path is not a safe repository-relative path")
	}
	if data.Task.Source.Repository != source.Repository || data.Task.Source.Commit != source.Commit || data.Task.Source.License != source.License.Expression {
		problems = append(problems, "task source metadata must match the audited registry entry")
	}
	p := data.Provenance.Source
	if p.ID != data.Task.Source.ID || p.Repository != data.Task.Source.Repository || p.Path != data.Task.Source.Path || p.Commit != data.Task.Source.Commit || p.License != data.Task.Source.License {
		problems = append(problems, "provenance source must exactly match task.yaml source")
	}
	if p.UsageMode != source.UsageMode {
		problems = append(problems, "provenance usage_mode must match the source registry")
	}
	if data.Provenance.Transformation.Type != data.Task.TransformationType {
		problems = append(problems, "provenance transformation type must match task.yaml")
	}
	for _, value := range append(append([]string{data.Provenance.Transformation.Summary}, data.Provenance.Transformation.IndependentlyAuthored...), data.Provenance.Transformation.Excluded...) {
		if strings.Contains(strings.ToLower(value), "todo") {
			problems = append(problems, "provenance contains an unresolved TODO placeholder")
			break
		}
	}
	if source.UsageMode == "blocked" {
		problems = append(problems, "blocked source cannot back a task")
	}
	if (source.UsageMode == "reference-only" || source.UsageMode == "concepts-only") && data.Task.TransformationType != "metadata-only" {
		problems = append(problems, "reference-only and concepts-only sources permit only metadata-only provenance")
	}
	if data.Task.TransformationType == "original" && data.Task.Source.ID != "go-practice-bank" {
		problems = append(problems, "original tasks require the registered go-practice-bank source")
	}
	if err := cleanRelativePath(data.Provenance.AttributionFile); err != nil {
		problems = append(problems, "attribution_file must be a safe relative path")
	} else {
		path := filepath.Join(data.Located.Path, filepath.FromSlash(data.Provenance.AttributionFile))
		if info, err := os.Lstat(path); err != nil || !info.Mode().IsRegular() {
			problems = append(problems, "declared attribution_file is missing")
		}
	}
	return uniqueSorted(problems)
}

func checkProvenanceV2(root string, data packageData) []string {
	registry, err := sources.LoadRegistry(root)
	if err != nil {
		return []string{"cannot load source registry: " + err.Error()}
	}
	identity, err := loadProjectIdentity(root)
	if err != nil {
		return []string{err.Error()}
	}
	licensePolicy, err := loadOutboundLicensePolicy(root)
	if err != nil {
		return []string{err.Error()}
	}
	return checkProvenanceV2WithInputs(data, registry, identity, licensePolicy)
}

func checkProvenanceV2WithInputs(data packageData, registry sources.Registry, identity projectIdentity, licensePolicy outboundLicensePolicy) []string {
	var problems []string
	p := data.ProvenanceV2
	if p.Project.Identity != identity.ID || p.Project.Repository != identity.Repository || p.Project.License != identity.License {
		problems = append(problems, "provenance project must match config/project-identity.yaml")
	}
	if identity.License != licensePolicy.OriginalContentLicense {
		problems = append(problems, "project identity license must match outbound original-content policy")
	}
	digest, digestErr := TaskContentDigest(data.Located.Path)
	if digestErr != nil {
		problems = append(problems, "compute task content digest: "+digestErr.Error())
	} else if p.Project.ContentDigest != digest {
		problems = append(problems, "provenance content_digest does not match authored task content; computed "+digest)
	}
	if strings.Contains(strings.ToLower(p.Creation.Summary), "todo") {
		problems = append(problems, "provenance creation summary contains an unresolved TODO placeholder")
	}
	expectedManifest := filepath.ToSlash(filepath.Join(identity.ReleaseManifestDirectory, data.Task.ID+".yaml"))
	if p.Release.Manifest != expectedManifest {
		problems = append(problems, "release manifest path must bind the task id under the configured release directory")
	}
	if p.Kind == "original" && len(p.Sources) != 0 {
		problems = append(problems, "original provenance must use internal project identity, not external source records")
	}
	compatible := make(map[string]bool, len(licensePolicy.CompatibleSourceLicense))
	for _, expression := range licensePolicy.CompatibleSourceLicense {
		compatible[expression] = true
	}
	seen := map[string]bool{}
	for _, reference := range p.Sources {
		if seen[reference.RegistryID] {
			problems = append(problems, fmt.Sprintf("duplicate provenance source %q", reference.RegistryID))
		}
		seen[reference.RegistryID] = true
		source, err := registry.Find(reference.RegistryID)
		if err != nil {
			problems = append(problems, err.Error())
			continue
		}
		if reference.Repository != source.Repository || reference.Commit != source.Commit || reference.License != source.License.Expression || reference.UsageMode != source.UsageMode {
			problems = append(problems, fmt.Sprintf("provenance source %q must exactly match the audited registry entry", reference.RegistryID))
		}
		if source.UsageMode == "blocked" {
			problems = append(problems, fmt.Sprintf("blocked source %q cannot back a task", reference.RegistryID))
		}
		if source.UsageMode == "adapt" && !compatible[source.License.Expression] {
			problems = append(problems, fmt.Sprintf("source %q license %q is not in the outbound compatibility allowlist", reference.RegistryID, source.License.Expression))
		}
		problems = append(problems, validateSourceRole(reference)...)
		artifactPaths := map[string]bool{}
		for _, artifact := range reference.Artifacts {
			if artifactPaths[artifact.Path] {
				problems = append(problems, fmt.Sprintf("source %q repeats artifact path %q", reference.RegistryID, artifact.Path))
			}
			artifactPaths[artifact.Path] = true
			if err := cleanRelativePath(artifact.Path); err != nil {
				problems = append(problems, fmt.Sprintf("source %q artifact path is unsafe", reference.RegistryID))
			}
			if artifact.SHA256 == strings.Repeat("0", 64) {
				problems = append(problems, fmt.Sprintf("source %q artifact digest is still an authoring stub", reference.RegistryID))
			}
		}
	}
	if err := cleanRelativePath(p.AttributionFile); err != nil {
		problems = append(problems, "attribution_file must be a safe relative path")
	} else if info, err := os.Lstat(filepath.Join(data.Located.Path, filepath.FromSlash(p.AttributionFile))); err != nil || !info.Mode().IsRegular() {
		problems = append(problems, "declared attribution_file is missing")
	}
	return uniqueSorted(problems)
}

func validateSourceRole(source ProvenanceSourceV2) []string {
	var problems []string
	roleAllowed := func(allowed ...string) bool { return contains(allowed, source.Role) }
	usageAllowed := func(value string, allowed ...string) bool { return contains(allowed, value) }
	switch source.UsageMode {
	case "reference-only":
		if !roleAllowed("factual-reference", "comparison-only") {
			problems = append(problems, fmt.Sprintf("reference-only source %q cannot use role %q", source.RegistryID, source.Role))
		}
		for _, artifact := range source.Artifacts {
			if !usageAllowed(artifact.Usage, "factual-reference", "comparison-only", "excluded") {
				problems = append(problems, fmt.Sprintf("reference-only source %q cannot classify artifact as %q", source.RegistryID, artifact.Usage))
			}
		}
	case "concepts-only":
		if !roleAllowed("concept-seed", "factual-reference", "comparison-only") {
			problems = append(problems, fmt.Sprintf("concepts-only source %q cannot use role %q", source.RegistryID, source.Role))
		}
		for _, artifact := range source.Artifacts {
			if !usageAllowed(artifact.Usage, "concept-only", "factual-reference", "comparison-only", "excluded") {
				problems = append(problems, fmt.Sprintf("concepts-only source %q cannot classify artifact as %q", source.RegistryID, artifact.Usage))
			}
		}
	}
	return problems
}

func checkReview(review Review) []string {
	var problems []string
	if review.Decision != "approve" {
		problems = append(problems, "review decision is not approve")
	}
	checks := []struct {
		value bool
		name  string
	}{
		{review.StoryFree, "story_free"},
		{review.PromptTestAlignment, "prompt_test_alignment"},
		{review.PublicAPIStable, "public_api_stable"},
		{review.ProvenanceVerified, "provenance_verified"},
		{review.LicenseVerified, "license_verified"},
		{review.ReferenceOnlyCleanRoomVerified, "reference_only_clean_room_verified"},
		{review.SolutionLeakageChecked, "solution_leakage_checked"},
	}
	for _, check := range checks {
		if !check.value {
			problems = append(problems, "review evidence is false: "+check.name)
		}
	}
	if review.MinimumMutants < 1 {
		problems = append(problems, "review must require at least one mutant")
	}
	return problems
}

func checkPublicAPI(layout packageLayout) []string {
	starter, err := apiFingerprint(layout.Starter)
	if err != nil {
		return []string{"cannot extract starter API: " + err.Error()}
	}
	var problems []string
	reference, err := apiFingerprint(layout.Reference)
	if err != nil {
		problems = append(problems, "cannot extract reference API: "+err.Error())
	} else if starter != reference {
		problems = append(problems, "starter and reference exported APIs differ")
	}
	for _, mutant := range layout.Mutants {
		fingerprint, err := apiFingerprint(mutant.Files)
		if err != nil {
			problems = append(problems, fmt.Sprintf("cannot extract mutant %s API: %v", mutant.ID, err))
		} else if starter != fingerprint {
			problems = append(problems, fmt.Sprintf("mutant %s exported API differs from starter", mutant.ID))
		}
	}
	return problems
}

func apiFingerprint(paths []string) (string, error) {
	fileset := token.NewFileSet()
	var declarations []string
	packageName := ""
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(fileset, path, nil, 0)
		if err != nil {
			return "", err
		}
		if packageName == "" {
			packageName = file.Name.Name
		} else if packageName != file.Name.Name {
			return "", errors.New("implementation files use multiple packages")
		}
		for _, declaration := range file.Decls {
			switch value := declaration.(type) {
			case *ast.FuncDecl:
				if !ast.IsExported(value.Name.Name) {
					continue
				}
				declarations = append(declarations, normalizedFunctionSignature(fileset, value))
			case *ast.GenDecl:
				for _, specification := range value.Specs {
					switch spec := specification.(type) {
					case *ast.TypeSpec:
						if ast.IsExported(spec.Name.Name) {
							declarations = append(declarations, value.Tok.String()+" "+nodeString(fileset, spec))
						}
					case *ast.ValueSpec:
						exported := false
						for _, name := range spec.Names {
							if ast.IsExported(name.Name) {
								exported = true
							}
						}
						if exported {
							declarations = append(declarations, value.Tok.String()+" "+nodeString(fileset, spec))
						}
					}
				}
			}
		}
	}
	sort.Strings(declarations)
	return packageName + "\n" + strings.Join(declarations, "\n"), nil
}

func normalizedFunctionSignature(fileset *token.FileSet, declaration *ast.FuncDecl) string {
	functionType := *declaration.Type
	functionType.Params = normalizedFieldList(declaration.Type.Params)
	functionType.Results = normalizedFieldList(declaration.Type.Results)
	receiver := ""
	if declaration.Recv != nil && len(declaration.Recv.List) != 0 {
		receiver = "(" + nodeString(fileset, declaration.Recv.List[0].Type) + ") "
	}
	return "func " + receiver + declaration.Name.Name + strings.TrimPrefix(nodeString(fileset, &functionType), "func")
}

func normalizedFieldList(input *ast.FieldList) *ast.FieldList {
	if input == nil {
		return nil
	}
	result := &ast.FieldList{Opening: input.Opening, Closing: input.Closing}
	for _, field := range input.List {
		count := len(field.Names)
		if count == 0 {
			count = 1
		}
		for range count {
			clone := *field
			clone.Doc = nil
			clone.Comment = nil
			clone.Names = nil
			result.List = append(result.List, &clone)
		}
	}
	return result
}

func nodeString(fileset *token.FileSet, node ast.Node) string {
	if node == nil {
		return ""
	}
	var output bytes.Buffer
	if err := format.Node(&output, fileset, node); err != nil {
		return "<invalid>"
	}
	return output.String()
}

func checkTestSources(layout packageLayout) []string {
	var problems []string
	for _, path := range append(append([]string{}, layout.PublicTests...), layout.GraderTests...) {
		fileProblems := inspectTestSource(path)
		for _, problem := range fileProblems {
			problems = append(problems, filepath.Base(path)+": "+problem)
		}
	}
	for _, path := range layout.PublicTests {
		content, err := os.ReadFile(path)
		if err == nil {
			lower := strings.ToLower(string(content))
			if strings.Contains(lower, "solution/reference") || strings.Contains(lower, "tests/grader") {
				problems = append(problems, filepath.Base(path)+": public test leaks hidden implementation or grader paths")
			}
		}
	}
	return uniqueSorted(problems)
}

func inspectTestSource(path string) []string {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, path, nil, 0)
	if err != nil {
		return []string{"cannot parse: " + err.Error()}
	}
	imports := map[string]string{}
	for _, item := range file.Imports {
		value, _ := strconv.Unquote(item.Path.Value)
		name := filepath.Base(value)
		if item.Name != nil {
			name = item.Name.Name
		}
		imports[name] = value
	}
	var problems []string
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		identifier, ok := selector.X.(*ast.Ident)
		if !ok {
			return true
		}
		importPath := imports[identifier.Name]
		key := importPath + "." + selector.Sel.Name
		switch {
		case key == "time.Sleep":
			problems = append(problems, "time.Sleep is not allowed")
		case importPath == "time" && (selector.Sel.Name == "Now" || selector.Sel.Name == "After" || selector.Sel.Name == "Tick"):
			problems = append(problems, key+" is an uncontrolled clock dependency")
		case importPath == "math/rand" || importPath == "crypto/rand":
			problems = append(problems, key+" is an uncontrolled randomness dependency")
		case importPath == "net" || importPath == "net/http" || importPath == "os/exec":
			problems = append(problems, key+" is an external-process or network dependency")
		case key == "os.Getenv" || key == "os.LookupEnv":
			problems = append(problems, key+" is an uncontrolled environment dependency")
		}
		return true
	})
	return uniqueSorted(problems)
}

func appendCopy(left, right []string) []string {
	result := append([]string(nil), left...)
	return append(result, right...)
}

func canonicalTestEvents(output []byte, taskID string) (string, bool) {
	var events []string
	compiled := false
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		var event struct {
			Action  string `json:"Action"`
			Package string `json:"Package"`
			Test    string `json:"Test"`
		}
		if json.Unmarshal(scanner.Bytes(), &event) != nil {
			continue
		}
		if event.Test != "" && event.Action == "run" {
			compiled = true
		}
		if event.Test != "" || event.Action == "pass" || event.Action == "fail" || event.Action == "skip" {
			events = append(events, strings.Join([]string{event.Action, strings.ReplaceAll(event.Package, "task.local/"+taskID, "task"), event.Test}, "|"))
		}
	}
	return strings.Join(events, "\n"), compiled
}

func expectFailure(label string, result variantResult) []string {
	if !result.Compiled {
		return []string{describeVariantFailure(label, result)}
	}
	if result.Passed {
		return []string{label + " unexpectedly passes"}
	}
	if result.FailureKind != variantFailureTest {
		return []string{describeVariantFailure(label, result)}
	}
	return nil
}

func isConcurrencyTask(task Task) bool {
	for _, id := range append([]string{task.Skills.Primary}, task.Skills.Secondary...) {
		domain, _, _ := strings.Cut(id, ".")
		switch domain {
		case "goroutines", "channels", "sync", "context":
			return true
		}
	}
	return false
}

func markdownSections(content string) map[string]string {
	result := map[string]string{}
	current := ""
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "## ") {
			current = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			if _, ok := result[current]; !ok {
				result[current] = ""
			}
			continue
		}
		if current != "" {
			result[current] += line + "\n"
		}
	}
	return result
}

func countBullets(content string) int {
	count := 0
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "- ") {
			count++
		}
	}
	return count
}

func countWords(content string) int {
	return len(strings.FieldsFunc(content, func(r rune) bool {
		return unicode.IsSpace(r) || strings.ContainsRune("#*`[](){}<>,:;", r)
	}))
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func firstDuplicate(values []string) string {
	seen := map[string]bool{}
	for _, value := range values {
		if seen[value] {
			return value
		}
		seen[value] = true
	}
	return ""
}

func firstOverlap(left, right []string) string {
	for _, value := range left {
		if contains(right, value) {
			return value
		}
	}
	return ""
}

func uniqueSorted(values []string) []string {
	set := map[string]bool{}
	for _, value := range values {
		if value != "" {
			set[value] = true
		}
	}
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
