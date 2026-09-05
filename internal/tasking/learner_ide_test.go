package tasking

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ideaWorkspaceFixture(t *testing.T, kind string) (learnerPackage, string) {
	t.Helper()
	workspace := t.TempDir()
	name := "subject/value.go"
	if kind == "tests" {
		name = "subject/value_test.go"
	}
	if kind == "workspace-patch" {
		name = "core/value.go"
	}
	files := []LearnerWorkspaceFile{{Path: name, Content: []byte("package subject\n"), Editable: true}, {Path: "go.mod", Content: []byte("module fixture\n")}}
	for _, file := range files {
		path := filepath.Join(workspace, filepath.FromSlash(file.Path))
		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, file.Content, 0600); err != nil {
			t.Fatal(err)
		}
	}
	loaded := learnerPackage{data: packageData{TaskV2: TaskV2{Submission: Submission{Kind: kind, EditablePaths: []string{name}}}}, files: files}
	return loaded, workspace
}

func ideaWrite(t *testing.T, workspace, relative string, content []byte) {
	t.Helper()
	path := filepath.Join(workspace, filepath.FromSlash(relative))
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatal(err)
	}
}

func TestLearnerIDEMetadataExcludedFromEverySubmissionKind(t *testing.T) {
	for _, kind := range []string{"implementation", "tests", "workspace-patch"} {
		t.Run(kind, func(t *testing.T) {
			loaded, workspace := ideaWorkspaceFixture(t, kind)
			before, err := inspectLearnerSubmission(loaded, workspace)
			if err != nil {
				t.Fatal(err)
			}
			metadata := map[string][]byte{
				".idea/workspace.xml":                  []byte("<project>private IDE state</project>"),
				".idea/modules.xml":                    []byte("<project/>"),
				".idea/inspectionProfiles/Project.xml": []byte("<component/>"),
				".idea/do_not_execute.go":              []byte("not even valid Go; IDE metadata must never enter compilation"),
			}
			for path, content := range metadata {
				ideaWrite(t, workspace, path, content)
			}
			for i := 0; i < 2; i++ {
				after, err := inspectLearnerSubmission(loaded, workspace)
				if err != nil {
					t.Fatal(err)
				}
				if strings.Join(before.implementation, "\n") != strings.Join(after.implementation, "\n") || strings.Join(before.tests, "\n") != strings.Join(after.tests, "\n") || len(before.workspace) != len(after.workspace) {
					t.Fatalf("IDE metadata changed selected submission: %+v", after)
				}
				for j, file := range after.workspace {
					if file != before.workspace[j] {
						t.Fatalf("IDE metadata changed workspace selection: %+v", after)
					}
				}
				for path, content := range metadata {
					actual, err := os.ReadFile(filepath.Join(workspace, path))
					if err != nil || !bytes.Equal(actual, content) {
						t.Fatalf("IDE metadata was changed: %s %v", path, err)
					}
				}
			}
		})
	}
}

func TestLearnerIDEMetadataDoesNotRelaxWorkspaceBoundary(t *testing.T) {
	for _, kind := range []string{"implementation", "tests", "workspace-patch"} {
		for _, scenario := range []string{"idea-file", "root-symlink", "nested-symlink", "unknown-directory", "similar-prefix", "nested-idea", "fixed-change", "extra-file"} {
			t.Run(kind+"/"+scenario, func(t *testing.T) {
				loaded, workspace := ideaWorkspaceFixture(t, kind)
				if scenario != "idea-file" && scenario != "root-symlink" {
					ideaWrite(t, workspace, ".idea/workspace.xml", []byte("<project/>"))
				}
				switch scenario {
				case "idea-file":
					ideaWrite(t, workspace, ".idea", []byte("invalid"))
				case "root-symlink":
					if err := os.Symlink(t.TempDir(), filepath.Join(workspace, ".idea")); err != nil {
						t.Fatal(err)
					}
				case "nested-symlink":
					if err := os.Symlink(t.TempDir(), filepath.Join(workspace, ".idea", "escape")); err != nil {
						t.Fatal(err)
					}
				case "unknown-directory":
					if err := os.Mkdir(filepath.Join(workspace, "unlisted"), 0700); err != nil {
						t.Fatal(err)
					}
				case "similar-prefix":
					ideaWrite(t, workspace, ".idea-other/workspace.xml", []byte("private"))
				case "nested-idea":
					ideaWrite(t, workspace, filepath.Dir(loaded.files[0].Path)+"/.idea/workspace.xml", []byte("private"))
				case "fixed-change":
					ideaWrite(t, workspace, "go.mod", []byte("changed"))
				case "extra-file":
					ideaWrite(t, workspace, "extra.go", []byte("package extra\n"))
				}
				if _, err := inspectLearnerSubmission(loaded, workspace); err == nil {
					t.Fatal("unsafe workspace accepted")
				}
			})
		}
	}
}

func TestLearnerIDEPathCannotBeDeclaredAsTaskContent(t *testing.T) {
	for _, kind := range []string{"implementation", "tests", "workspace-patch"} {
		for _, path := range []string{".idea", ".idea/submitted.go"} {
			for _, editable := range []bool{false, true} {
				t.Run(kind+"/"+path+"/"+map[bool]string{false: "fixed", true: "editable"}[editable], func(t *testing.T) {
					loaded, workspace := ideaWorkspaceFixture(t, kind)
					content := []byte("package metadata\n")
					ideaWrite(t, workspace, path, content)
					loaded.files = append(loaded.files, LearnerWorkspaceFile{Path: path, Content: content, Editable: editable})
					if editable {
						loaded.data.TaskV2.Submission.EditablePaths = append(loaded.data.TaskV2.Submission.EditablePaths, path)
					}
					if _, err := inspectLearnerSubmission(loaded, workspace); err == nil {
						t.Fatal("task-declared IDE path accepted")
					}
				})
			}
		}
	}
}
