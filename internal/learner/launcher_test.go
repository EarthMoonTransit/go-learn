package learner

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestGoLandSharedProjectPreservesLessonState(t *testing.T) {
	for _, version := range []int{1, 3} {
		t.Run(fmt.Sprint(version), func(t *testing.T) {
			s := goLandTestService(t)
			if err := saveProgress(s.home, Progress{SchemaVersion: version, Tasks: map[string]*TaskProgress{}}); err != nil {
				t.Fatal(err)
			}
			lessons := []struct{ id, file string }{
				{"literal-bindings-001", "subject/bindings.go"},
				{"zero-results-001", "subject/subject.go"},
				{"clamp-tests-001", "subject/subject_test.go"},
			}
			workspaces := map[string]string{}
			for _, lesson := range lessons {
				workspace, err := s.Start(lesson.id)
				if err != nil {
					t.Fatal(err)
				}
				workspaces[lesson.id] = workspace
			}
			p, err := s.Progress()
			if err != nil {
				t.Fatal(err)
			}
			p.Tasks[lessons[0].id].Attempts = 4
			p.Tasks[lessons[0].id].UsedHints = []int{1}
			p.Tasks[lessons[0].id].ActivitySeconds = 321
			p.Tasks[lessons[0].id].Completed = true
			if err := saveProgress(s.home, p); err != nil {
				t.Fatal(err)
			}
			before := persistedBytes(t, s)
			tree, err := workspaceTree(filepath.Join(s.home, "workspaces"), "")
			if err != nil {
				t.Fatal(err)
			}
			// No actual IDE or authored code runs in this regression.
			for _, lesson := range lessons {
				want := ideTarget{filepath.Join(s.home, "workspaces"), filepath.Join(workspaces[lesson.id], lesson.file)}
				var calls int
				s.launch = func(got ideTarget) error {
					calls++
					if got != want {
						t.Fatalf("target = %+v; want %+v", got, want)
					}
					return nil
				}
				view, err := s.TaskView(lesson.id)
				if err != nil || view.IDEProject != want.Project || view.Workspace != workspaces[lesson.id] {
					t.Fatalf("view = %+v, %v", view, err)
				}
				for range 2 {
					result, err := s.PerformAction("open", lesson.id, view.Revision)
					if err != nil || result.Workspace != workspaces[lesson.id] {
						t.Fatalf("open = %+v, %v", result, err)
					}
				}
				if calls != 2 {
					t.Fatalf("launcher called %d times; want 2", calls)
				}
			}
			afterTree, err := workspaceTree(filepath.Join(s.home, "workspaces"), "")
			if err != nil || !reflect.DeepEqual(tree, afterTree) || !bytes.Equal(before, persistedBytes(t, s)) {
				t.Fatalf("open changed solution files, markers, paths or progress: %v", err)
			}
		})
	}
}

func TestGoLandRejectsUnsafeLessonFiles(t *testing.T) {
	s := goLandTestService(t)
	id := "literal-bindings-001"
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	s.launch = func(ideTarget) error { t.Fatal("unsafe target reached launcher"); return nil }
	view, err := s.TaskView(id)
	if err != nil {
		t.Fatal(err)
	}
	before := persistedBytes(t, s)
	file := filepath.Join(workspace, "subject", "bindings.go")
	if err := os.Remove(file); err != nil {
		t.Fatal(err)
	}
	if _, err := s.PerformAction("open", id, view.Revision); err == nil {
		t.Fatal("missing lesson file accepted")
	}
	if err := os.Symlink(filepath.Join(workspace, "go.mod"), file); err != nil {
		t.Fatal(err)
	}
	if _, err := s.PerformAction("open", id, view.Revision); err == nil {
		t.Fatal("symlinked lesson file accepted")
	}
	if err := os.Rename(filepath.Join(workspace, "subject"), filepath.Join(workspace, "saved-subject")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("saved-subject", filepath.Join(workspace, "subject")); err != nil {
		t.Fatal(err)
	}
	if _, err := s.PerformAction("open", id, view.Revision); err == nil {
		t.Fatal("symlinked parent accepted")
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("rejected open changed progress")
	}
}

func TestGoLandLauncherPassesProjectAndFileAsSeparateArguments(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fake POSIX launcher")
	}
	bin := t.TempDir()
	output := filepath.Join(bin, "arguments")
	if err := os.WriteFile(filepath.Join(bin, "goland"), []byte("#!/bin/sh\nprintf '%s\\n' \"$@\" > \"$GOLAND_TEST_ARGS\"\n"), 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", bin)
	t.Setenv("GOLAND_TEST_ARGS", output)
	project := filepath.Join(t.TempDir(), "Course with spaces $(not-a-command)")
	target := ideTarget{project, filepath.Join(project, "01-first-steps", "01-lesson", "subject", "lesson.go")}
	if err := launchGoLand(target); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(output)
	if err != nil || string(data) != target.Project+"\n"+target.File+"\n" {
		t.Fatalf("launcher arguments = %q, %v", data, err)
	}
	for _, invalid := range []ideTarget{{"relative", target.File}, {project, "relative.go"}, {project, filepath.Join(filepath.Dir(project), "outside.go")}} {
		if err := launchGoLand(invalid); err != ErrIDEUnavailable {
			t.Fatalf("unsafe arguments accepted: %+v", invalid)
		}
	}
}

func goLandTestService(t *testing.T) *Service {
	t.Helper()
	s := uiTestService(t, "")
	root := t.TempDir()
	// Keep repeated compatibility checks bounded to the three submission kinds.
	for _, dir := range strings.Fields("config schemas curriculum sources tasks/ready/literal-bindings-001 tasks/ready/zero-results-001 tasks/ready/clamp-tests-001") {
		if err := isolation.CopyTree(filepath.Join(s.root, dir), filepath.Join(root, dir)); err != nil {
			t.Fatal(err)
		}
	}
	if err := isolation.CopyFile(filepath.Join(s.root, "docs/LEARNER_START.md"), filepath.Join(root, "docs/LEARNER_START.md")); err != nil {
		t.Fatal(err)
	}
	s.root = root
	return s
}
