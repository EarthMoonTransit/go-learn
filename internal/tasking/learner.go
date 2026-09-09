package tasking

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const learnerWorkspaceMarker = ".go-learn-workspace.json"

// LearnerTaskInfo is the intentionally public subset of Task Contract v2.
// Verification internals, grader paths, reference paths, and mutant ids are
// deliberately absent.
type LearnerTaskInfo struct {
	ID                   string
	Title                string
	Activity             string
	SubmissionKind       string
	PrimarySkill         string
	SecondarySkills      []string
	Prerequisites        []string
	NewConcepts          []string
	Difficulty           string
	EstimatedTimeMinutes TimeRange
	VerificationProfile  string
	Prompt               string
	ContentDigest        string
}

type LearnerWorkspaceFile struct {
	Path     string
	Content  []byte
	Editable bool
}

type LearnerVerificationScope string

const (
	LearnerPublicChecks LearnerVerificationScope = "public"
	LearnerFullChecks   LearnerVerificationScope = "full"
)

type LearnerVerificationResult struct {
	Passed          bool
	Scope           LearnerVerificationScope
	SubmissionKind  string
	Summary         string
	Diagnostic      string
	FailureCategory string
}

type learnerPackage struct {
	data   packageData
	layout packageLayout
	info   LearnerTaskInfo
	files  []LearnerWorkspaceFile
}

// ListReadyLearnerTasks lists only ready Task Contract v2 packages. Legacy
// packages remain authoring-compatible but are not exposed through this CLI.
func ListReadyLearnerTasks(root string) ([]LearnerTaskInfo, error) {
	tasks, err := DiscoverTasks(root, []State{StateReady})
	if err != nil {
		return nil, err
	}
	result := make([]LearnerTaskInfo, 0, len(tasks))
	for _, located := range tasks {
		version, err := documentSchemaVersion(filepath.Join(located.Path, "task.yaml"))
		if err != nil {
			return nil, fmt.Errorf("inspect %s: %w", located.ID, err)
		}
		if version != 2 {
			continue
		}
		loaded, err := loadLearnerPackage(root, located.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, loaded.info)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result, nil
}

func LoadReadyLearnerTask(root, id string) (LearnerTaskInfo, error) {
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		return LearnerTaskInfo{}, err
	}
	return loaded.info, nil
}

// LearnerWorkspaceFiles returns only learner-visible material. Hints are read
// progressively by the CLI and are not copied. Hidden grader, reference, and
// mutant artifacts never enter this file set.
func LearnerWorkspaceFiles(root, id string) ([]LearnerWorkspaceFile, error) {
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		return nil, err
	}
	return cloneLearnerFiles(loaded.files), nil
}

func LearnerHints(root, id string) ([]string, error) {
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		return nil, err
	}
	content, err := readRegularFile(filepath.Join(loaded.data.Located.Path, "hints", "HINTS.md"), 1<<20)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, raw := range strings.Split(string(content), "\n") {
		line := strings.TrimSpace(raw)
		separator := strings.Index(line, ". ")
		if separator < 1 {
			continue
		}
		ordinal := line[:separator]
		validOrdinal := true
		for _, character := range ordinal {
			if character < '0' || character > '9' {
				validOrdinal = false
			}
		}
		if validOrdinal {
			result = append(result, strings.TrimSpace(line[separator+2:]))
		}
	}
	if len(result) == 0 {
		return nil, errors.New("task has no parseable progressive hints")
	}
	return result, nil
}

// LearnerSolutionFiles returns the reference submission only. Callers must
// enforce an explicit reveal action and must not persist the returned content
// in progress state or the learner workspace.
func LearnerSolutionFiles(root, id string) ([]LearnerWorkspaceFile, error) {
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		return nil, err
	}
	var sources []mappedFile
	switch loaded.data.TaskV2.Submission.Kind {
	case "implementation":
		for _, source := range loaded.layout.Reference {
			sources = append(sources, mappedFile{Source: source, Relative: filepath.ToSlash(filepath.Join("subject", filepath.Base(source)))})
		}
	case "tests":
		for _, source := range loaded.layout.Reference {
			sources = append(sources, mappedFile{Source: source, Relative: filepath.ToSlash(filepath.Join("subject", filepath.Base(source)))})
		}
	case "workspace-patch":
		sources = append(sources, loaded.layout.WorkspaceReference...)
	default:
		return nil, fmt.Errorf("unsupported submission kind %q", loaded.data.TaskV2.Submission.Kind)
	}
	result := make([]LearnerWorkspaceFile, 0, len(sources))
	for _, source := range sources {
		content, err := readRegularFile(source.Source, 8<<20)
		if err != nil {
			return nil, err
		}
		result = append(result, LearnerWorkspaceFile{Path: source.Relative, Content: content, Editable: true})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Path < result[j].Path })
	return result, nil
}

// LearnerSubmissionFiles validates the persistent workspace boundary and
// returns only learner-owned Go files. It is used for safe review-context
// export, including submissions with syntax or API errors. Hidden checker
// material is never part of the result.
func LearnerSubmissionFiles(root, id, workspace string) ([]LearnerWorkspaceFile, error) {
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		return nil, err
	}
	submission, err := inspectLearnerSubmission(loaded, workspace)
	if err != nil {
		return nil, err
	}
	var sources []mappedFile
	switch loaded.data.TaskV2.Submission.Kind {
	case "implementation":
		for _, source := range submission.implementation {
			sources = append(sources, mappedFile{Source: source, Relative: filepath.ToSlash(filepath.Join("subject", filepath.Base(source)))})
		}
	case "tests":
		for _, source := range submission.tests {
			sources = append(sources, mappedFile{Source: source, Relative: filepath.ToSlash(filepath.Join("subject", filepath.Base(source)))})
		}
	case "workspace-patch":
		sources = append(sources, submission.workspace...)
	}
	result := make([]LearnerWorkspaceFile, 0, len(sources))
	for _, source := range sources {
		content, err := readRegularFile(source.Source, 8<<20)
		if err != nil {
			return nil, err
		}
		result = append(result, LearnerWorkspaceFile{Path: source.Relative, Content: content, Editable: true})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Path < result[j].Path })
	return result, nil
}

// VerifyLearnerWorkspace always assembles and executes a separate temporary Go
// module. Public mode returns sanitized public diagnostics. Full mode uses the
// hidden grader internally and returns only a coarse summary.
func VerifyLearnerWorkspace(root, id, workspace string, scope LearnerVerificationScope) (LearnerVerificationResult, error) {
	if scope != LearnerPublicChecks && scope != LearnerFullChecks {
		return LearnerVerificationResult{}, fmt.Errorf("invalid learner verification scope %q", scope)
	}
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		return LearnerVerificationResult{}, err
	}
	submission, err := inspectLearnerSubmission(loaded, workspace)
	if err != nil {
		return LearnerVerificationResult{}, err
	}
	if err := validateLearnerSubmission(loaded, submission); err != nil {
		result := LearnerVerificationResult{
			Scope: scope, SubmissionKind: loaded.info.SubmissionKind,
			Summary: "submission does not satisfy the public code contract", FailureCategory: "submission",
		}
		if scope == LearnerPublicChecks {
			result.Diagnostic = sanitizeVariantDiagnostic(nil, err, "", workspace, id)
		}
		return result, nil
	}
	runner, err := newRealExecutionRunner()
	if err != nil {
		return LearnerVerificationResult{}, err
	}
	defer runner.Close()

	plan := loaded.data.TaskV2.Verification
	base := variantRequest{TaskID: id, GoVersion: learnerGoVersion(loaded), TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired}
	switch loaded.data.TaskV2.Submission.Kind {
	case "implementation":
		tests := loaded.layout.PublicTests
		if scope == LearnerFullChecks {
			tests = appendCopy(loaded.layout.PublicTests, loaded.layout.GraderTests)
		}
		request := base
		request.Implementation, request.Tests = submission.implementation, tests
		return verifyLearnerVariant(runner, request, loaded, scope)
	case "tests":
		request := base
		request.Implementation, request.Tests = loaded.layout.VerificationReference, submission.tests
		if scope == LearnerPublicChecks {
			return learnerPublicResult(runner.RunVariant(request), "tests"), nil
		}
		return verifyLearnerMutation(runner, request, loaded, submission.tests), nil
	case "workspace-patch":
		request := workspaceRequest{
			TaskID: id, GoVersion: learnerGoVersion(loaded), TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired,
			Skeleton: loaded.layout.WorkspaceSkeleton, Overlay: submission.workspace,
		}
		if scope == LearnerPublicChecks {
			request.Tests, request.Checks = loaded.layout.WorkspacePublic, plan.PublicChecks
			return learnerPublicResult(runner.RunWorkspace(request), "workspace-patch"), nil
		}
		request.Tests = appendMapped(loaded.layout.WorkspacePublic, loaded.layout.WorkspaceGrader)
		request.Checks = plan.GraderChecks
		return verifyLearnerWorkspaceFull(runner, request, loaded), nil
	default:
		return LearnerVerificationResult{}, fmt.Errorf("unsupported submission kind %q", loaded.data.TaskV2.Submission.Kind)
	}
}

type learnerSubmission struct {
	implementation []string
	tests          []string
	workspace      []mappedFile
}

func loadLearnerPackage(root, id string) (learnerPackage, error) {
	located, err := FindTask(root, id)
	if err != nil {
		return learnerPackage{}, err
	}
	if located.State != StateReady {
		return learnerPackage{}, fmt.Errorf("task %q is not ready", id)
	}
	data, loadProblems := loadPackage(root, located)
	if len(loadProblems) != 0 {
		return learnerPackage{}, fmt.Errorf("ready task %q has invalid metadata: %s", id, loadProblems[0])
	}
	if data.ContractVersion != 2 {
		return learnerPackage{}, fmt.Errorf("task %q uses legacy contract v%d; learner CLI requires v2", id, data.ContractVersion)
	}
	contractProblems, layout := checkPackageContractV2(root, data)
	problems := append([]string{}, contractProblems...)
	problems = append(problems, checkMetadata(root, data)...)
	problems = append(problems, checkPrompt(data.Prompt)...)
	problems = append(problems, checkProvenance(root, data)...)
	problems = append(problems, checkReview(data.Review)...)
	problems = append(problems, checkPublicAPIV2(data, layout)...)
	problems = append(problems, checkTestSourcesV2(data, layout)...)
	problems = uniqueSorted(problems)
	if len(problems) != 0 {
		return learnerPackage{}, fmt.Errorf("ready task %q is not safe for learner export: %s", id, problems[0])
	}
	digest, err := TaskContentDigest(located.Path)
	if err != nil {
		return learnerPackage{}, err
	}
	info := LearnerTaskInfo{
		ID: id, Title: data.TaskV2.Title, Activity: data.TaskV2.Activity,
		SubmissionKind: data.TaskV2.Submission.Kind, PrimarySkill: data.TaskV2.Skills.Primary,
		SecondarySkills: append([]string(nil), data.TaskV2.Skills.Secondary...),
		Prerequisites:   append([]string(nil), data.TaskV2.Prerequisites...),
		NewConcepts:     append([]string(nil), data.TaskV2.NewConcepts...),
		Difficulty:      data.TaskV2.Difficulty, EstimatedTimeMinutes: data.TaskV2.EstimatedTimeMinutes,
		VerificationProfile: data.TaskV2.Verification.Profile, Prompt: string(data.Prompt), ContentDigest: digest,
	}
	loaded := learnerPackage{data: data, layout: layout, info: info}
	loaded.files, err = buildLearnerWorkspaceFiles(loaded)
	if err != nil {
		return learnerPackage{}, err
	}
	return loaded, nil
}

func buildLearnerWorkspaceFiles(loaded learnerPackage) ([]LearnerWorkspaceFile, error) {
	var mapped []struct {
		file     mappedFile
		editable bool
	}
	mapped = append(mapped,
		struct {
			file     mappedFile
			editable bool
		}{mappedFile{Source: filepath.Join(loaded.data.Located.Path, "TASK.md"), Relative: "TASK.md"}, false},
		struct {
			file     mappedFile
			editable bool
		}{mappedFile{Source: filepath.Join(loaded.data.Located.Path, "go.mod"), Relative: "go.mod"}, false},
	)
	switch loaded.data.TaskV2.Submission.Kind {
	case "implementation":
		for _, source := range loaded.layout.Starter {
			mapped = append(mapped, struct {
				file     mappedFile
				editable bool
			}{mappedFile{Source: source, Relative: filepath.ToSlash(filepath.Join("subject", filepath.Base(source)))}, true})
		}
	case "tests":
		for _, source := range loaded.layout.Starter {
			mapped = append(mapped, struct {
				file     mappedFile
				editable bool
			}{mappedFile{Source: source, Relative: filepath.ToSlash(filepath.Join("subject", filepath.Base(source)))}, true})
		}
	case "workspace-patch":
		for _, source := range loaded.layout.WorkspaceSkeleton {
			mapped = append(mapped, struct {
				file     mappedFile
				editable bool
			}{source, false})
		}
		for _, source := range loaded.layout.WorkspaceStarter {
			mapped = append(mapped, struct {
				file     mappedFile
				editable bool
			}{source, true})
		}
	default:
		return nil, fmt.Errorf("unsupported submission kind %q", loaded.data.TaskV2.Submission.Kind)
	}
	seen := map[string]bool{}
	result := make([]LearnerWorkspaceFile, 0, len(mapped))
	if loaded.data.TaskV2.Submission.Kind == "tests" {
		declarations, err := learnerTestAPI(loaded.layout.VerificationReference)
		if err != nil {
			return nil, err
		}
		if len(declarations) > 0 {
			seen[learnerAPIPath] = true
			result = append(result, LearnerWorkspaceFile{Path: learnerAPIPath, Content: declarations})
		}
	}
	for _, item := range mapped {
		if err := cleanRelativePath(item.file.Relative); err != nil {
			return nil, err
		}
		if seen[item.file.Relative] {
			return nil, fmt.Errorf("learner workspace path collision: %s", item.file.Relative)
		}
		seen[item.file.Relative] = true
		content, err := readRegularFile(item.file.Source, 8<<20)
		if err != nil {
			return nil, err
		}
		result = append(result, LearnerWorkspaceFile{Path: item.file.Relative, Content: content, Editable: item.editable})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Path < result[j].Path })
	return result, nil
}

func isLearnerIDEMetadata(relative string) bool {
	return relative == ".idea" || strings.HasPrefix(relative, ".idea/")
}

func inspectLearnerSubmission(loaded learnerPackage, workspace string) (learnerSubmission, error) {
	info, err := os.Lstat(workspace)
	if err != nil {
		return learnerSubmission{}, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return learnerSubmission{}, errors.New("learner workspace must be a real directory")
	}
	expected := map[string]LearnerWorkspaceFile{}
	allowedDirectories := map[string]bool{".": true}
	for _, file := range loaded.files {
		if isLearnerIDEMetadata(file.Path) {
			return learnerSubmission{}, errors.New("task content cannot use the reserved .idea metadata directory")
		}
		expected[file.Path] = file
		for directory := filepath.ToSlash(filepath.Dir(file.Path)); directory != "."; directory = filepath.ToSlash(filepath.Dir(directory)) {
			allowedDirectories[directory] = true
		}
	}
	if loaded.data.TaskV2.Submission.Kind != "workspace-patch" {
		allowedDirectories["subject"] = true
	}
	found := map[string]bool{}
	err = filepath.WalkDir(workspace, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(workspace, path)
		if err != nil {
			return err
		}
		relative = filepath.ToSlash(relative)
		if relative == "." {
			return nil
		}
		fileInfo, err := os.Lstat(path)
		if err != nil {
			return err
		}
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symlink is not allowed in learner workspace: %s", relative)
		}
		// IDE settings remain private workspace metadata. Inspect the entire
		// subtree (including links and file kinds), but never select its files.
		if isLearnerIDEMetadata(relative) {
			if relative == ".idea" && !fileInfo.IsDir() {
				return errors.New("learner .idea metadata must be a directory")
			}
			if !fileInfo.IsDir() && !fileInfo.Mode().IsRegular() {
				return fmt.Errorf("non-regular learner file is not allowed: %s", relative)
			}
			return nil
		}
		if entry.IsDir() {
			if !allowedDirectories[relative] {
				return fmt.Errorf("unexpected directory in learner workspace: %s", relative)
			}
			return nil
		}
		if !fileInfo.Mode().IsRegular() {
			return fmt.Errorf("non-regular learner file is not allowed: %s", relative)
		}
		if relative == learnerWorkspaceMarker {
			found[relative] = true
			return nil
		}
		if declared, ok := expected[relative]; ok {
			found[relative] = true
			if !declared.Editable {
				content, err := readRegularFile(path, 8<<20)
				if err != nil {
					return err
				}
				if !bytes.Equal(content, declared.Content) {
					return fmt.Errorf("learner read-only file was modified: %s", relative)
				}
			}
			return nil
		}
		if loaded.data.TaskV2.Submission.Kind != "workspace-patch" && filepath.Dir(relative) == "subject" && filepath.Ext(relative) == ".go" {
			found[relative] = true
			return nil
		}
		return fmt.Errorf("unexpected file in learner workspace: %s", relative)
	})
	if err != nil {
		return learnerSubmission{}, err
	}
	for _, file := range loaded.files {
		// This optional editor aid did not exist in older test workspaces.
		// Its absence must not invalidate their tests or progress.
		if loaded.data.TaskV2.Submission.Kind == "tests" && file.Path == learnerAPIPath && !file.Editable {
			continue
		}
		if loaded.data.TaskV2.Submission.Kind == "workspace-patch" || !file.Editable {
			if !found[file.Path] {
				return learnerSubmission{}, fmt.Errorf("required learner workspace file is missing: %s", file.Path)
			}
		}
	}

	var result learnerSubmission
	switch loaded.data.TaskV2.Submission.Kind {
	case "implementation", "tests":
		entries, err := os.ReadDir(filepath.Join(workspace, "subject"))
		if err != nil {
			return learnerSubmission{}, errors.New("learner subject directory is missing")
		}
		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
				return learnerSubmission{}, fmt.Errorf("learner subject may contain only Go files: %s", entry.Name())
			}
			path := filepath.Join(workspace, "subject", entry.Name())
			if loaded.data.TaskV2.Submission.Kind == "implementation" {
				if strings.HasSuffix(entry.Name(), "_test.go") {
					return learnerSubmission{}, fmt.Errorf("implementation submission may not add test file %s", entry.Name())
				}
				result.implementation = append(result.implementation, path)
			} else {
				if declared, ok := expected["subject/"+entry.Name()]; ok && declared.Path == learnerAPIPath && !declared.Editable {
					// The filesystem inspection above verified these exact bytes.
					// Editor declarations never enter the verification submission.
					continue
				}
				if !strings.HasSuffix(entry.Name(), "_test.go") {
					return learnerSubmission{}, fmt.Errorf("test submission may contain only *_test.go files: %s", entry.Name())
				}
				result.tests = append(result.tests, path)
			}
		}
		sort.Strings(result.implementation)
		sort.Strings(result.tests)
	case "workspace-patch":
		for _, editable := range loaded.data.TaskV2.Submission.EditablePaths {
			path := filepath.Join(workspace, filepath.FromSlash(editable))
			if !found[editable] {
				return learnerSubmission{}, fmt.Errorf("editable workspace file is missing: %s", editable)
			}
			result.workspace = append(result.workspace, mappedFile{Source: path, Relative: editable})
		}
	}
	return result, nil
}

// Code correctness is separate from the filesystem boundary: a learner must
// be able to export broken code, and an invalid submission is a failed attempt.
func validateLearnerSubmission(loaded learnerPackage, submission learnerSubmission) error {
	switch loaded.data.TaskV2.Submission.Kind {
	case "implementation":
		if len(submission.implementation) == 0 {
			return errors.New("implementation submission has no Go files")
		}
		starterAPI, err := apiFingerprint(loaded.layout.Starter)
		if err != nil {
			return err
		}
		learnerAPI, err := apiFingerprint(submission.implementation)
		if err != nil {
			return fmt.Errorf("cannot parse learner code: %w", err)
		}
		if learnerAPI != starterAPI {
			return errors.New("learner submission changed the public API")
		}
	case "tests":
		if len(submission.tests) == 0 {
			return errors.New("test submission has no Go test files")
		}
		var problems []string
		for _, path := range submission.tests {
			for _, problem := range inspectLearnerTestSource(path) {
				problems = append(problems, filepath.Base(path)+": "+problem)
			}
		}
		if problems = uniqueSorted(problems); len(problems) != 0 {
			return errors.New(problems[0])
		}
	case "workspace-patch":
		starterAPI, err := workspaceAPIFingerprint(loaded.layout.WorkspaceSkeleton, loaded.layout.WorkspaceStarter)
		if err != nil {
			return err
		}
		learnerAPI, err := workspaceAPIFingerprint(loaded.layout.WorkspaceSkeleton, submission.workspace)
		if err != nil {
			return fmt.Errorf("cannot parse learner code: %w", err)
		}
		if learnerAPI != starterAPI {
			return errors.New("learner workspace changed the public API")
		}
	}
	return nil
}

func verifyLearnerVariant(runner *realExecutionRunner, request variantRequest, loaded learnerPackage, scope LearnerVerificationScope) (LearnerVerificationResult, error) {
	if scope == LearnerPublicChecks {
		return learnerPublicResult(runner.RunVariant(request), loaded.info.SubmissionKind), nil
	}
	first, deterministic := repeatLearnerVariant(loaded.data.TaskV2.Verification.Limits.DeterminismRuns, func() variantResult {
		return runner.RunVariant(request)
	})
	if !first.Passed || !deterministic {
		return learnerFullFailure(first, deterministic, loaded.info.SubmissionKind), nil
	}
	if auxiliary := learnerVariantAuxiliary(runner, request, loaded); auxiliary != nil {
		return *auxiliary, nil
	}
	return learnerFullPass(loaded.info.SubmissionKind), nil
}

func verifyLearnerWorkspaceFull(runner *realExecutionRunner, request workspaceRequest, loaded learnerPackage) LearnerVerificationResult {
	first, deterministic := repeatLearnerVariant(loaded.data.TaskV2.Verification.Limits.DeterminismRuns, func() variantResult {
		return runner.RunWorkspace(request)
	})
	if !first.Passed || !deterministic {
		return learnerFullFailure(first, deterministic, loaded.info.SubmissionKind)
	}
	if auxiliary := learnerWorkspaceAuxiliary(runner, request, loaded); auxiliary != nil {
		return *auxiliary
	}
	return learnerFullPass(loaded.info.SubmissionKind)
}

func verifyLearnerMutation(runner *realExecutionRunner, referenceRequest variantRequest, loaded learnerPackage, tests []string) LearnerVerificationResult {
	first, deterministic := repeatLearnerVariant(loaded.data.TaskV2.Verification.Limits.DeterminismRuns, func() variantResult {
		return runner.RunVariant(referenceRequest)
	})
	if !first.Passed || !deterministic {
		result := learnerFullFailure(first, deterministic, loaded.info.SubmissionKind)
		if first.Passed {
			result.Summary = "submitted tests are nondeterministic"
		} else {
			result.Summary = "submitted tests reject the canonical public behavior"
		}
		return result
	}
	for _, mutant := range loaded.layout.Mutants {
		request := referenceRequest
		request.Implementation, request.Tests = mutant.Files, tests
		result := runner.RunVariant(request)
		if !result.Compiled || result.Passed || result.FailureKind != variantFailureTest {
			return LearnerVerificationResult{
				Passed: false, Scope: LearnerFullChecks, SubmissionKind: loaded.info.SubmissionKind,
				Summary:         "one or more private defect models survived the submitted tests",
				FailureCategory: "hidden-verification",
			}
		}
	}
	if auxiliary := learnerVariantAuxiliary(runner, referenceRequest, loaded); auxiliary != nil {
		return *auxiliary
	}
	return learnerFullPass(loaded.info.SubmissionKind)
}

func learnerVariantAuxiliary(runner *realExecutionRunner, request variantRequest, loaded learnerPackage) *LearnerVerificationResult {
	plan := loaded.data.TaskV2.Verification
	if plan.Vet {
		vet := request
		vet.Vet, vet.Race = true, false
		if result := runner.RunVariant(vet); !result.Passed {
			value := learnerAuxiliaryFailure("vet", loaded.info.SubmissionKind)
			return &value
		}
	}
	if plan.RaceRequired {
		race := request
		race.Race = true
		if result := runner.RunVariant(race); !result.Passed {
			value := learnerAuxiliaryFailure("race", loaded.info.SubmissionKind)
			return &value
		}
	}
	return nil
}

func learnerWorkspaceAuxiliary(runner *realExecutionRunner, request workspaceRequest, loaded learnerPackage) *LearnerVerificationResult {
	plan := loaded.data.TaskV2.Verification
	if plan.Vet {
		vet := request
		vet.Vet, vet.Race = true, false
		if result := runner.RunWorkspace(vet); !result.Passed {
			value := learnerAuxiliaryFailure("vet", loaded.info.SubmissionKind)
			return &value
		}
	}
	if plan.RaceRequired {
		race := request
		race.Race = true
		if result := runner.RunWorkspace(race); !result.Passed {
			value := learnerAuxiliaryFailure("race", loaded.info.SubmissionKind)
			return &value
		}
	}
	return nil
}

func repeatLearnerVariant(runs int, run func() variantResult) (variantResult, bool) {
	var first variantResult
	canonical := ""
	deterministic := true
	for index := 0; index < runs; index++ {
		result := run()
		if index == 0 {
			first = result
			canonical = result.Canonical
		}
		if !result.Passed {
			return result, true
		}
		if index > 0 && result.Canonical != canonical {
			deterministic = false
		}
	}
	return first, deterministic
}

func learnerPublicResult(result variantResult, submissionKind string) LearnerVerificationResult {
	value := LearnerVerificationResult{Passed: result.Passed, Scope: LearnerPublicChecks, SubmissionKind: submissionKind}
	if result.Passed {
		value.Summary = "public checks passed"
		return value
	}
	value.FailureCategory = string(result.FailureKind)
	value.Diagnostic = result.Diagnostic
	switch result.FailureKind {
	case variantFailureCompile, variantFailureAssembly:
		value.Summary = "submission did not compile with public checks"
	case variantFailureTimeout:
		value.Summary = "public checks timed out"
	default:
		value.Summary = "public checks found a problem"
	}
	return value
}

func learnerFullFailure(result variantResult, deterministic bool, submissionKind string) LearnerVerificationResult {
	value := LearnerVerificationResult{Passed: false, Scope: LearnerFullChecks, SubmissionKind: submissionKind, FailureCategory: "hidden-verification"}
	if !deterministic {
		value.Summary = "submission produced nondeterministic test events"
		return value
	}
	switch result.FailureKind {
	case variantFailureCompile, variantFailureAssembly:
		value.Summary = "submission did not compile in full verification"
	case variantFailureTimeout:
		value.Summary = "full verification timed out"
	default:
		value.Summary = "full checks found a behavior problem"
	}
	return value
}

func learnerFullPass(submissionKind string) LearnerVerificationResult {
	return LearnerVerificationResult{Passed: true, Scope: LearnerFullChecks, SubmissionKind: submissionKind, Summary: "full verification passed"}
}

func learnerAuxiliaryFailure(gate, submissionKind string) LearnerVerificationResult {
	return LearnerVerificationResult{Passed: false, Scope: LearnerFullChecks, SubmissionKind: submissionKind, Summary: gate + " verification failed", FailureCategory: gate}
}

func learnerGoVersion(loaded learnerPackage) string {
	content, err := os.ReadFile(filepath.Join(loaded.data.Located.Path, "go.mod"))
	if err == nil {
		for _, line := range strings.Split(string(content), "\n") {
			fields := strings.Fields(line)
			if len(fields) == 2 && fields[0] == "go" {
				return fields[1]
			}
		}
	}
	return "1.27"
}

func readRegularFile(path string, limit int64) ([]byte, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return nil, fmt.Errorf("path is not a regular file: %s", path)
	}
	if info.Size() > limit {
		return nil, fmt.Errorf("file exceeds %d bytes: %s", limit, path)
	}
	return os.ReadFile(path)
}

func cloneLearnerFiles(input []LearnerWorkspaceFile) []LearnerWorkspaceFile {
	result := make([]LearnerWorkspaceFile, len(input))
	for index, file := range input {
		result[index] = file
		result[index].Content = append([]byte(nil), file.Content...)
	}
	return result
}
