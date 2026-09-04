package tasking

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLearnerAPIContainsOnlyUsableDeclarations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "subject.go")
	content := `package subject
const secret = "SECRET_BODY_VALUE"
func helper() string { return secret }
// SECRET_COMMENT
func Clamp(panic, low, high int) (result int) { return low + high + panic }
func Combine(values []int, options map[string]bool, extra ...int) ([]byte, error) { return nil, nil }
func Apply(callback func(PRIVATE_ARG int) (PRIVATE_RESULT int)) int { return callback(5) }
`
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	stub, err := learnerTestAPI([]string{path})
	if err != nil || len(stub) == 0 {
		t.Fatalf("stub=%q err=%v", stub, err)
	}
	for _, hidden := range []string{"SECRET", "PRIVATE_ARG", "PRIVATE_RESULT", "helper", "low + high", "const secret"} {
		if strings.Contains(string(stub), hidden) {
			t.Fatalf("exported hidden content %q", hidden)
		}
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "learn_api.go", stub, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Go's type checker resolves these declarations without executing code.
	config := types.Config{}
	pkg, err := config.Check("subject", fset, []*ast.File{file}, nil)
	if err != nil || pkg.Scope().Lookup("Clamp") == nil || pkg.Scope().Lookup("Combine") == nil {
		t.Fatalf("API does not resolve: %v", err)
	}
	if !strings.Contains(string(stub), "use learn check or learn complete") {
		t.Fatal("missing local-run guidance")
	}
}

func TestLearnerAPIPlatformFilesFallBackWithoutDuplicateDeclarations(t *testing.T) {
	directory := t.TempDir()
	var paths []string
	for _, platform := range []string{"darwin", "linux"} {
		path := filepath.Join(directory, "subject_"+platform+".go")
		if err := os.WriteFile(path, []byte("package subject\nfunc Value() int {return 3}\n"), 0600); err != nil {
			t.Fatal(err)
		}
		paths = append(paths, path)
	}
	stub, err := learnerTestAPI(paths)
	if err != nil || len(stub) != 0 {
		t.Fatalf("platform API exported: %q err=%v", stub, err)
	}
	for _, tag := range []string{"//go:build !windows", "// +build linux"} {
		path := filepath.Join(directory, "subject.go")
		if err := os.WriteFile(path, []byte(tag+"\n\npackage subject\nfunc Value() int {return 3}\n"), 0600); err != nil {
			t.Fatal(err)
		}
		stub, err := learnerTestAPI([]string{path})
		if err != nil || len(stub) != 0 {
			t.Fatalf("build-constrained API exported: %q err=%v", stub, err)
		}
	}
	path := filepath.Join(directory, "public_api.go")
	if err := os.WriteFile(path, []byte("package subject\nfunc Value() int {return 3}\n"), 0600); err != nil {
		t.Fatal(err)
	}
	stub, err = learnerTestAPI([]string{path})
	if err != nil || len(stub) == 0 {
		t.Fatalf("ordinary underscore filename rejected: %q err=%v", stub, err)
	}
}

func TestLearnerAPIUnsupportedShapesDoNotLeakPartialStructure(t *testing.T) {
	cases := []string{
		`type secret struct { Value int }; func Read() secret { return secret{} }`,
		`type int string; func Read() int { return "hidden" }`,
		`const Public = "hidden"; func Read() int { return 1 }`,
		`type Record struct { private int }; func Read() int { return 1 }`,
		`func Read[T any](value T) T { return value }`,
		`type secret int; func (secret) Read() int { return 1 }`,
		`const count = 3; func Read() [count]int { return [count]int{} }`,
	}
	for _, content := range cases {
		path := filepath.Join(t.TempDir(), "subject.go")
		if err := os.WriteFile(path, []byte("package subject\n"+content), 0600); err != nil {
			t.Fatal(err)
		}
		stub, err := learnerTestAPI([]string{path})
		if err != nil || len(stub) != 0 {
			t.Fatalf("unsupported API exported %q: %v", stub, err)
		}
	}
}

func TestLearnerAPIAidIsReadonlyOptionalAndExcludedFromSubmission(t *testing.T) {
	workspace := t.TempDir()
	if err := os.Mkdir(filepath.Join(workspace, "subject"), 0700); err != nil {
		t.Fatal(err)
	}
	testPath := filepath.Join(workspace, "subject", "subject_test.go")
	if err := os.WriteFile(testPath, []byte("package subject_test\n"), 0600); err != nil {
		t.Fatal(err)
	}
	api := []byte("package subject\nfunc Clamp(int,int,int) int { panic(\"editor only\") }\n")
	loaded := learnerPackage{
		data: packageData{TaskV2: TaskV2{Submission: Submission{Kind: "tests"}}},
		files: []LearnerWorkspaceFile{
			{Path: "subject/subject_test.go", Editable: true},
			{Path: learnerAPIPath, Content: api},
		},
	}
	assertTestsOnly := func() {
		t.Helper()
		submission, err := inspectLearnerSubmission(loaded, workspace)
		if err != nil || len(submission.tests) != 1 || submission.tests[0] != testPath {
			t.Fatalf("incorrect test selection: %+v err=%v", submission, err)
		}
	}
	assertTestsOnly() // Pre-existing workspace, with no generated declaration.
	apiPath := filepath.Join(workspace, filepath.FromSlash(learnerAPIPath))
	if err := os.WriteFile(apiPath, api, 0600); err != nil {
		t.Fatal(err)
	}
	assertTestsOnly()
	if err := os.WriteFile(apiPath, []byte("package subject\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := inspectLearnerSubmission(loaded, workspace); err == nil || !strings.Contains(err.Error(), "read-only") {
		t.Fatalf("modified API accepted: %v", err)
	}
	if err := os.WriteFile(apiPath, api, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(workspace, "subject", "extra.go"), []byte("package subject\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := inspectLearnerSubmission(loaded, workspace); err == nil {
		t.Fatal("an undeclared implementation entered a test submission")
	}
}
