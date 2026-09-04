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
	"time"
	"unicode"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/sources"
	"go-practice-bank/internal/triage"
)

const taskCommandTimeout = 60 * time.Second

var sentenceEndPattern = regexp.MustCompile(`[.!?](?:\s|$)`)

type checker struct {
	report QualityReport
}

func Check(root, id string) (QualityReport, error) {
	located, err := FindTask(root, id)
	if err != nil {
		return QualityReport{}, err
	}
	return checkLocated(root, located)
}

func checkLocated(root string, located LocatedTask) (QualityReport, error) {
	c := checker{report: QualityReport{
		SchemaVersion: 1, TaskID: located.ID, State: located.State,
		Mutants: MutantSummary{Survived: []string{}},
	}}
	data, loadProblems := loadPackage(root, located)
	c.gate("package-metadata", loadProblems)

	contractProblems, layout := checkPackageContract(data)
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
	c.gate("public-api-stability", apiProblems)

	determinismProblems := checkTestSources(layout)
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

	environment, err := isolation.NewEnvironment("go-task-check")
	if err != nil {
		return QualityReport{}, err
	}
	defer environment.Close()

	starterPublic := runTaskPackages(environment, data.Located.Path, data.Task.ID, false, "./tests/public")
	c.gate("starter-public", expectFailure("starter with public tests", starterPublic))
	starterFull := runTaskPackages(environment, data.Located.Path, data.Task.ID, false, "./tests/public", "./tests/grader")
	c.gate("starter-full", expectFailure("starter with full tests", starterFull))

	referenceFirst := runTaskPackages(environment, data.Located.Path, data.Task.ID, false, "./solution/reference")
	referenceSecond := runTaskPackages(environment, data.Located.Path, data.Task.ID, false, "./solution/reference")
	var referenceProblems []string
	if !referenceFirst.Compiled {
		referenceProblems = append(referenceProblems, "reference does not compile with full tests")
	} else if !referenceFirst.Passed {
		referenceProblems = append(referenceProblems, "reference fails full tests")
	}
	if !referenceSecond.Compiled || !referenceSecond.Passed {
		referenceProblems = append(referenceProblems, "reference fails the repeated full test run")
	}
	if referenceFirst.Canonical != referenceSecond.Canonical {
		referenceProblems = append(referenceProblems, "repeated reference runs have different observable test events")
	}
	c.gate("reference-full", referenceProblems)

	c.report.Mutants.Required = data.Review.MinimumMutants
	c.report.Mutants.Total = len(layout.Mutants)
	var mutantProblems []string
	for _, mutant := range layout.Mutants {
		result := runTaskPackages(environment, data.Located.Path, data.Task.ID, false, "./solution/mutants/"+mutant.ID)
		if !result.Compiled {
			mutantProblems = append(mutantProblems, fmt.Sprintf("mutant %s does not compile", mutant.ID))
			continue
		}
		if result.Passed {
			c.report.Mutants.Survived = append(c.report.Mutants.Survived, mutant.ID)
			mutantProblems = append(mutantProblems, fmt.Sprintf("mutant %s survives full tests", mutant.ID))
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
		race := runTaskPackages(environment, data.Located.Path, data.Task.ID, true, "./solution/reference")
		var raceProblems []string
		if !race.Compiled || !race.Passed {
			raceProblems = append(raceProblems, "reference fails go test -race with full tests")
		}
		c.gate("race", raceProblems)
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

func (c *checker) finish() {
	c.report.Passed = true
	for _, gate := range c.report.Gates {
		if gate.Status == "fail" {
			c.report.Passed = false
			break
		}
	}
}

type packageLayout struct {
	TaskDir     string
	GoMod       string
	Starter     []string
	PublicTests []string
	GraderTests []string
	Reference   []string
	Mutants     []mutantLayout
}

type mutantLayout struct {
	ID    string
	Files []string
}

func checkPackageContract(data packageData) ([]string, packageLayout) {
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
	goMod, err := os.ReadFile(layout.GoMod)
	if err == nil {
		problems = append(problems, validateGoMod(data.Located.ID, string(goMod))...)
	}
	layout.Starter, problems = collectGoDirectory(data.Located.Path, "starter", directoryImplementation, problems)
	layout.PublicTests, problems = collectGoDirectory(data.Located.Path, "tests/public", directoryTestHarness, problems)
	layout.GraderTests, problems = collectGoDirectory(data.Located.Path, "tests/grader", directoryTestHarness, problems)
	layout.Reference, problems = collectGoDirectory(data.Located.Path, "solution/reference", directoryExecutableSolution, problems)

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
			files, next := collectGoDirectory(data.Located.Path, filepath.ToSlash(filepath.Join("solution/mutants", entry.Name())), directoryExecutableSolution, problems)
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

type directoryKind int

const (
	directoryImplementation directoryKind = iota
	directoryTestHarness
	directoryExecutableSolution
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
	}
	sort.Strings(files)
	if len(files) == 0 {
		problems = append(problems, fmt.Sprintf("%s must contain at least one Go file", relative))
	}
	if kind == directoryImplementation && tests != 0 {
		problems = append(problems, fmt.Sprintf("starter tests must live in tests/public, not %s", relative))
	}
	if kind == directoryTestHarness && tests == 0 {
		problems = append(problems, fmt.Sprintf("%s must contain a *_test.go binding", relative))
	}
	if kind == directoryExecutableSolution && (tests == 0 || nonTests == 0) {
		problems = append(problems, fmt.Sprintf("%s must contain implementation and *_test.go binding files", relative))
	}
	return files, problems
}

func validateGoMod(id, content string) []string {
	var problems []string
	moduleSeen := false
	goSeen := false
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "module":
			moduleSeen = len(fields) == 2 && fields[1] == "task.local/"+id
			if !moduleSeen {
				problems = append(problems, "go.mod module must be task.local/"+id)
			}
		case "go":
			goSeen = len(fields) == 2 && fields[1] == "1.24"
			if !goSeen {
				problems = append(problems, "go.mod must declare go 1.24")
			}
		case "require", "replace", "exclude", "retract", "tool", "godebug":
			problems = append(problems, fmt.Sprintf("go.mod directive %s is not allowed", fields[0]))
		}
	}
	if !moduleSeen {
		problems = append(problems, "go.mod is missing the required module directive")
	}
	if !goSeen {
		problems = append(problems, "go.mod is missing the required go directive")
	}
	return uniqueSorted(problems)
}

func checkMetadata(root string, data packageData) []string {
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
	_, contracts, err := triage.LoadSkillContracts(root)
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

type variantResult struct {
	Compiled  bool
	Passed    bool
	Canonical string
}

func runTaskPackages(environment *isolation.Environment, taskDirectory, taskID string, race bool, packages ...string) variantResult {
	workspace, err := environment.NewWorkspace("variant")
	if err != nil {
		return variantResult{}
	}
	if err := isolation.CopyTree(taskDirectory, workspace); err != nil {
		return variantResult{}
	}
	compileArgs := []string{"test", "-run", "^$", "-count=1"}
	compileArgs = append(compileArgs, packages...)
	compileOutput, compileErr := environment.Run(workspace, taskCommandTimeout, "go", compileArgs...)
	_ = compileOutput
	if compileErr != nil {
		return variantResult{}
	}
	args := []string{"test", "-json", "-count=1"}
	if race {
		args = append(args, "-race")
	}
	args = append(args, packages...)
	output, runErr := environment.Run(workspace, taskCommandTimeout, "go", args...)
	return variantResult{Compiled: true, Passed: runErr == nil, Canonical: canonicalTestEvents(output, taskID)}
}

func canonicalTestEvents(output []byte, taskID string) string {
	var events []string
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
		if event.Test != "" || event.Action == "pass" || event.Action == "fail" || event.Action == "skip" {
			events = append(events, strings.Join([]string{event.Action, strings.ReplaceAll(event.Package, "task.local/"+taskID, "task"), event.Test}, "|"))
		}
	}
	return strings.Join(events, "\n")
}

func expectFailure(label string, result variantResult) []string {
	if !result.Compiled {
		return []string{label + " does not compile"}
	}
	if result.Passed {
		return []string{label + " unexpectedly passes"}
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
