package tasking

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeclaredSubjectWorkspaceDirectoryIsAllowed(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "subject"), 0700); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "subject", "value.go")
	if err := os.WriteFile(path, []byte("package subject\n"), 0600); err != nil {
		t.Fatal(err)
	}
	loaded := learnerPackage{
		data:  packageData{TaskV2: TaskV2{Submission: Submission{Kind: "workspace-patch", EditablePaths: []string{"subject/value.go"}}}},
		files: []LearnerWorkspaceFile{{Path: "subject/value.go", Editable: true}},
	}
	if submission, err := inspectLearnerSubmission(loaded, root); err != nil || len(submission.workspace) != 1 {
		t.Fatalf("declared directory rejected: %+v %v", submission, err)
	}
	if err := os.WriteFile(filepath.Join(root, "subject", "extra.go"), []byte("package subject\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := inspectLearnerSubmission(loaded, root); err == nil {
		t.Fatal("undeclared file accepted")
	}
}
func TestSubjectDirectoryIsNotImplicitlyAllowedForWorkspacePatches(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"core", "subject"} {
		if err := os.Mkdir(filepath.Join(root, dir), 0700); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "core", "value.go"), []byte("package core\n"), 0600); err != nil {
		t.Fatal(err)
	}
	loaded := learnerPackage{
		data:  packageData{TaskV2: TaskV2{Submission: Submission{Kind: "workspace-patch", EditablePaths: []string{"core/value.go"}}}},
		files: []LearnerWorkspaceFile{{Path: "core/value.go", Editable: true}},
	}
	if _, err := inspectLearnerSubmission(loaded, root); err == nil {
		t.Fatal("undeclared subject directory accepted")
	}
}
