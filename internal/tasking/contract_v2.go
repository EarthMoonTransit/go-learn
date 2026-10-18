package tasking

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go-practice-bank/internal/toolchainpolicy"
)

func checkPackageContractV2(root string, data packageData) ([]string, packageLayout) {
	profile := data.TaskV2.Verification.Profile
	switch profile {
	case "fixed-tests":
		problems, layout := checkPackageContractV1(root, data)
		problems = append(problems, compareRequiredMutants(data.TaskV2.Verification.RequiredMutants, mutantIDs(layout.Mutants))...)
		return uniqueSorted(problems), layout
	case "mutation-tests":
		return checkMutationLayout(root, data)
	case "workspace-tests":
		return checkWorkspaceLayout(root, data)
	case "fuzz-seed-tests":
		return checkFuzzLayout(root, data)
	case "benchmark-contract":
		return []string{"verification profile " + profile + " is reserved but not executable in " + CheckContractVersion}, packageLayout{TaskDir: data.Located.Path}
	default:
		return []string{"unknown verification profile"}, packageLayout{TaskDir: data.Located.Path}
	}
}

func checkV2CommonLayout(root string, data packageData) ([]string, packageLayout) {
	var problems []string
	layout := packageLayout{TaskDir: data.Located.Path, GoMod: filepath.Join(data.Located.Path, "go.mod")}
	for _, name := range []string{"TASK.md", "task.yaml", "lifecycle.yaml", "go.mod", "review/review.yaml", "review/notes.md", "hints/HINTS.md", "provenance.yaml", "ATTRIBUTION.md"} {
		path := filepath.Join(data.Located.Path, filepath.FromSlash(name))
		info, err := os.Lstat(path)
		if err != nil || !info.Mode().IsRegular() {
			problems = append(problems, "required regular file is missing: "+name)
		} else if info.Size() == 0 {
			problems = append(problems, "required file is empty: "+name)
		}
	}
	if _, err := os.Lstat(filepath.Join(data.Located.Path, "go.work")); err == nil {
		problems = append(problems, "go.work is not allowed in a task package")
	}
	problems = append(problems, checkPureGoBoundary(data.Located.Path)...)
	if goMod, err := os.ReadFile(layout.GoMod); err == nil {
		policy, policyErr := toolchainpolicy.Load(root)
		if policyErr != nil {
			problems = append(problems, "load toolchain policy: "+policyErr.Error())
		} else {
			problems = append(problems, toolchainpolicy.ValidateGoMod(goMod, policy.ModulePrefix+data.Located.ID, policy.GoLanguageVersion, policy.ForbidToolchainDirective)...)
		}
	}
	return problems, layout
}

func checkMutationLayout(root string, data packageData) ([]string, packageLayout) {
	problems, layout := checkV2CommonLayout(root, data)
	layout.Starter, problems = collectGoDirectory(data.Located.Path, "starter", directoryTestHarness, problems)
	layout.Reference, problems = collectGoDirectory(data.Located.Path, "solution/reference", directoryTestHarness, problems)
	layout.VerificationReference, problems = collectGoDirectory(data.Located.Path, "verification/reference", directoryImplementation, problems)
	layout.Mutants, problems = collectMutants(data.Located.Path, "verification/mutants", problems)
	problems = append(problems, compareRequiredMutants(data.TaskV2.Verification.RequiredMutants, mutantIDs(layout.Mutants))...)
	return uniqueSorted(problems), layout
}

func checkWorkspaceLayout(root string, data packageData) ([]string, packageLayout) {
	problems, layout := checkV2CommonLayout(root, data)
	var next []string
	layout.WorkspaceSkeleton, next = collectMappedFiles(data.Located.Path, "workspace", nil)
	problems = append(problems, next...)
	problems = append(problems, exactMappedPaths("workspace skeleton", data.TaskV2.Submission.WorkspaceSkeleton, layout.WorkspaceSkeleton)...)
	layout.WorkspaceStarter, next = collectMappedFiles(data.Located.Path, "starter", nil)
	problems = append(problems, next...)
	problems = append(problems, exactMappedPaths("starter overlay", data.TaskV2.Submission.EditablePaths, layout.WorkspaceStarter)...)
	layout.WorkspaceReference, next = collectMappedFiles(data.Located.Path, "solution/reference", nil)
	problems = append(problems, next...)
	problems = append(problems, exactMappedPaths("reference overlay", data.TaskV2.Submission.EditablePaths, layout.WorkspaceReference)...)
	layout.WorkspacePublic, next = collectMappedFiles(data.Located.Path, "tests/public", nil)
	problems = append(problems, next...)
	layout.WorkspaceGrader, next = collectMappedFiles(data.Located.Path, "tests/grader", nil)
	problems = append(problems, next...)
	if len(layout.WorkspacePublic) == 0 {
		problems = append(problems, "tests/public must contain at least one Go test file")
	}
	if len(layout.WorkspaceGrader) == 0 {
		problems = append(problems, "tests/grader must contain at least one Go test file")
	}
	for _, group := range [][]mappedFile{layout.WorkspacePublic, layout.WorkspaceGrader} {
		for _, file := range group {
			if !strings.HasSuffix(file.Relative, "_test.go") {
				problems = append(problems, "workspace test harness may contain only *_test.go files: "+file.Relative)
			}
		}
	}
	mutantRoot := filepath.Join(data.Located.Path, "solution", "mutants")
	entries, err := os.ReadDir(mutantRoot)
	if err != nil {
		problems = append(problems, "required directory is missing: solution/mutants")
	} else {
		for _, entry := range entries {
			if !entry.IsDir() || !slugPattern.MatchString(entry.Name()) {
				problems = append(problems, "invalid mutant entry: "+entry.Name())
				continue
			}
			files, fileProblems := collectMappedFiles(data.Located.Path, filepath.ToSlash(filepath.Join("solution/mutants", entry.Name())), nil)
			problems = append(problems, fileProblems...)
			problems = append(problems, exactMappedPaths("mutant "+entry.Name()+" overlay", data.TaskV2.Submission.EditablePaths, files)...)
			layout.WorkspaceMutants = append(layout.WorkspaceMutants, workspaceMutantLayout{ID: entry.Name(), Files: files})
		}
		sort.Slice(layout.WorkspaceMutants, func(i, j int) bool { return layout.WorkspaceMutants[i].ID < layout.WorkspaceMutants[j].ID })
	}
	var ids []string
	for _, mutant := range layout.WorkspaceMutants {
		ids = append(ids, mutant.ID)
	}
	problems = append(problems, compareRequiredMutants(data.TaskV2.Verification.RequiredMutants, ids)...)
	return uniqueSorted(problems), layout
}

func collectMutants(taskRoot, relative string, problems []string) ([]mutantLayout, []string) {
	root := filepath.Join(taskRoot, filepath.FromSlash(relative))
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, append(problems, "required directory is missing: "+relative)
	}
	var mutants []mutantLayout
	for _, entry := range entries {
		if !entry.IsDir() || !slugPattern.MatchString(entry.Name()) {
			problems = append(problems, "invalid mutant entry: "+entry.Name())
			continue
		}
		files, next := collectGoDirectory(taskRoot, filepath.ToSlash(filepath.Join(relative, entry.Name())), directoryImplementation, problems)
		problems = next
		mutants = append(mutants, mutantLayout{ID: entry.Name(), Files: files})
	}
	sort.Slice(mutants, func(i, j int) bool { return mutants[i].ID < mutants[j].ID })
	if len(mutants) == 0 {
		problems = append(problems, relative+" must contain at least one mutant directory")
	}
	return mutants, problems
}

func collectMappedFiles(taskRoot, relative string, problems []string) ([]mappedFile, []string) {
	root := filepath.Join(taskRoot, filepath.FromSlash(relative))
	if info, err := os.Lstat(root); err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return nil, append(problems, "required directory is missing: "+relative)
	}
	var files []mappedFile
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == root || entry.IsDir() {
			return nil
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		local, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		local = filepath.ToSlash(local)
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			problems = append(problems, "workspace input must be a regular file: "+filepath.ToSlash(filepath.Join(relative, local)))
			return nil
		}
		if filepath.Ext(local) != ".go" {
			problems = append(problems, "workspace input may contain only Go files: "+filepath.ToSlash(filepath.Join(relative, local)))
			return nil
		}
		files = append(files, mappedFile{Source: path, Relative: local})
		return nil
	})
	if err != nil {
		problems = append(problems, "inspect "+relative+": "+err.Error())
	}
	sort.Slice(files, func(i, j int) bool { return files[i].Relative < files[j].Relative })
	return files, problems
}

func exactMappedPaths(label string, expected []string, actual []mappedFile) []string {
	var paths []string
	for _, file := range actual {
		paths = append(paths, file.Relative)
	}
	expected = append([]string(nil), expected...)
	sort.Strings(expected)
	if strings.Join(expected, "\n") != strings.Join(paths, "\n") {
		return []string{fmt.Sprintf("%s paths are %v; require exactly %v", label, paths, expected)}
	}
	return nil
}

func mutantIDs(mutants []mutantLayout) []string {
	result := make([]string, 0, len(mutants))
	for _, mutant := range mutants {
		result = append(result, mutant.ID)
	}
	return result
}

func compareRequiredMutants(required, actual []string) []string {
	required, actual = append([]string(nil), required...), append([]string(nil), actual...)
	sort.Strings(required)
	sort.Strings(actual)
	if strings.Join(required, "\n") != strings.Join(actual, "\n") {
		return []string{fmt.Sprintf("required_mutants %v must exactly match package mutants %v", required, actual)}
	}
	return nil
}

func checkMetadataV2(data packageData) []string {
	var problems []string
	plan := data.TaskV2.Verification
	if data.TaskV2.EstimatedTimeMinutes.Min > data.TaskV2.EstimatedTimeMinutes.Max {
		problems = append(problems, "estimated time min cannot exceed max")
	}
	for _, group := range [][]string{plan.PublicChecks, plan.GraderChecks, plan.ReferenceChecks} {
		for _, check := range group {
			if !contains(plan.Packages, check) && !contains(plan.Packages, "./...") {
				problems = append(problems, fmt.Sprintf("verification check %q is not declared in packages", check))
			}
		}
	}
	if len(plan.RequiredMutants) == 0 && contains([]string{"fixed-tests", "mutation-tests", "workspace-tests", "fuzz-seed-tests"}, plan.Profile) {
		problems = append(problems, "executable v2 profiles require at least one required mutant")
	}
	if len(plan.RequiredMutants) < data.Review.MinimumMutants {
		problems = append(problems, fmt.Sprintf("verification requires %d mutants; review requires at least %d", len(plan.RequiredMutants), data.Review.MinimumMutants))
	}
	if data.SubmissionOverlap() != "" {
		problems = append(problems, "editable path must not also be part of the immutable workspace skeleton: "+data.SubmissionOverlap())
	}
	for _, path := range data.TaskV2.Submission.EditablePaths {
		for _, reserved := range []string{"tests/", "solution/", "verification/", "review/", "workspace/", "starter/"} {
			if strings.HasPrefix(path, reserved) {
				problems = append(problems, fmt.Sprintf("editable path %q targets checker-owned or metadata content", path))
			}
		}
	}
	return problems
}

func (data packageData) SubmissionOverlap() string {
	for _, editable := range data.TaskV2.Submission.EditablePaths {
		if contains(data.TaskV2.Submission.WorkspaceSkeleton, editable) {
			return editable
		}
	}
	return ""
}

func checkPublicAPIV2(data packageData, layout packageLayout) []string {
	switch data.TaskV2.Verification.Profile {
	case "fixed-tests":
		return checkPublicAPI(layout)
	case "mutation-tests":
		canonical, err := apiFingerprint(layout.VerificationReference)
		if err != nil {
			return []string{"cannot extract verification reference API: " + err.Error()}
		}
		var problems []string
		for _, mutant := range layout.Mutants {
			fingerprint, err := apiFingerprint(mutant.Files)
			if err != nil || fingerprint != canonical {
				problems = append(problems, "verification mutant "+mutant.ID+" does not preserve the public API")
			}
		}
		return problems
	case "fuzz-seed-tests":
		return checkFuzzPublicAPI(layout)
	case "workspace-tests":
		starterAPI, err := workspaceAPIFingerprint(layout.WorkspaceSkeleton, layout.WorkspaceStarter)
		if err != nil {
			return []string{"cannot extract workspace starter API: " + err.Error()}
		}
		referenceAPI, err := workspaceAPIFingerprint(layout.WorkspaceSkeleton, layout.WorkspaceReference)
		if err != nil || referenceAPI != starterAPI {
			return []string{"workspace starter and reference exported APIs differ"}
		}
	}
	return nil
}

func mappedSources(files []mappedFile) []string {
	result := make([]string, 0, len(files))
	for _, file := range files {
		result = append(result, file.Source)
	}
	return result
}

func workspaceAPIFingerprint(groups ...[]mappedFile) (string, error) {
	byDirectory := map[string][]string{}
	for _, group := range groups {
		for _, file := range group {
			directory := filepath.ToSlash(filepath.Dir(file.Relative))
			byDirectory[directory] = append(byDirectory[directory], file.Source)
		}
	}
	var directories []string
	for directory := range byDirectory {
		directories = append(directories, directory)
	}
	sort.Strings(directories)
	var result []string
	for _, directory := range directories {
		fingerprint, err := apiFingerprint(byDirectory[directory])
		if err != nil {
			return "", err
		}
		result = append(result, directory+"\n"+fingerprint)
	}
	return strings.Join(result, "\n"), nil
}

func checkTestSourcesV2(data packageData, layout packageLayout) []string {
	switch data.TaskV2.Verification.Profile {
	case "fixed-tests":
		return checkTestSources(layout)
	case "mutation-tests":
		var problems []string
		for _, path := range append(append([]string{}, layout.Starter...), layout.Reference...) {
			for _, problem := range inspectLearnerTestSource(path) {
				problems = append(problems, filepath.Base(path)+": "+problem)
			}
		}
		return uniqueSorted(problems)
	case "fuzz-seed-tests":
		// Typed finite admission uses actual pinned compiler exports in the
		// native kernel, before any submitted test binary can run. Generic
		// Test-entrypoint rules do not describe the explicit FuzzHex grammar.
		return nil
	case "workspace-tests":
		var problems []string
		for _, file := range append(append([]mappedFile{}, layout.WorkspacePublic...), layout.WorkspaceGrader...) {
			for _, problem := range inspectTestSource(file.Source) {
				problems = append(problems, file.Relative+": "+problem)
			}
		}
		return uniqueSorted(problems)
	}
	return nil
}

func inspectLearnerTestSource(path string) []string {
	problems := inspectTestSource(path)
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, path, nil, 0)
	if err != nil {
		return append(problems, "cannot parse learner test: "+err.Error())
	}
	if file.Name.Name != "subject_test" {
		problems = append(problems, "learner test package must be subject_test")
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
	referencesAPI := false
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Body == nil {
			continue
		}
		if function.Name.Name == "TestMain" {
			ast.Inspect(function.Body, func(node ast.Node) bool {
				if isImportedSelectorCall(node, imports, "os", "Exit") {
					problems = append(problems, "TestMain may not call os.Exit")
				}
				return true
			})
		}
		if strings.HasPrefix(function.Name.Name, "Test") {
			for _, statement := range function.Body.List {
				if isUnconditionalFailure(statement) {
					problems = append(problems, "test contains an unconditional top-level failure")
				}
			}
		}
	}

	ast.Inspect(file, func(node ast.Node) bool {
		// A method call on a locally constructed task value has no package
		// qualifier at the call site. Its type reference, field access, or
		// method expression is valid API use as well as a free-function call.
		// Actual behavioral sufficiency is enforced by required mutants.
		selector, ok := node.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		identifier, ok := selector.X.(*ast.Ident)
		if ok && strings.HasPrefix(imports[identifier.Name], "task.local/") {
			referencesAPI = true
		}
		return true
	})
	if !referencesAPI {
		problems = append(problems, "learner tests must reference the task public API")
	}
	return uniqueSorted(problems)
}

func isImportedSelectorCall(node ast.Node, imports map[string]string, imported, method string) bool {
	call, ok := node.(*ast.CallExpr)
	if !ok {
		return false
	}
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || selector.Sel.Name != method {
		return false
	}
	identifier, ok := selector.X.(*ast.Ident)
	return ok && imports[identifier.Name] == imported
}

func isUnconditionalFailure(statement ast.Stmt) bool {
	expression, ok := statement.(*ast.ExprStmt)
	if !ok {
		return false
	}
	call, ok := expression.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	if identifier, ok := call.Fun.(*ast.Ident); ok && identifier.Name == "panic" {
		return true
	}
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	return contains([]string{"Fatal", "Fatalf", "Fail", "FailNow"}, selector.Sel.Name)
}
