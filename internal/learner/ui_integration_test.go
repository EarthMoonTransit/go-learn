//go:build taskintegration

package learner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

func TestUIRealVerificationAcrossSubmissionKinds(t *testing.T) {
	for _, fixture := range []struct{ id, editable, reference, weak string }{
		{"named-block-001", "subject/bindings.go", "subject/bindings.go", "solution/mutants/wrong-element-type/subject/bindings.go"},
		{"zero-results-001", "subject/subject.go", "subject.go", "solution/mutants/empty-map/subject.go"},
		{"clamp-tests-001", "subject/subject_test.go", "subject_test.go", "starter/subject_test.go"},
	} {
		t.Run(fixture.id, func(t *testing.T) {
			s := uiTestService(t, "")
			s.verify = tasking.VerifyLearnerWorkspace
			// Exercise the existing UI flow after moving a legacy workspace.
			if err := saveProgress(s.home, Progress{SchemaVersion: 1, Tasks: map[string]*TaskProgress{}}); err != nil {
				t.Fatal(err)
			}
			u := testLocalUI(t, s)
			serial := 0
			run := func(action string) uiJob {
				t.Helper()
				serial++
				id := fmt.Sprintf("integration-request-%03d", serial)
				w := uiHTTP(u, "POST", "/api/action", uiActionBody(t, s, action, fixture.id, id), nil)
				if w.Code != 202 {
					t.Fatalf("%s admission: %d %s", action, w.Code, w.Body)
				}
				job := waitUIJob(t, u, id)
				if job.State != "finished" {
					t.Fatalf("%s failed: %+v", action, job)
				}
				return job
			}
			course, err := s.CourseView()
			if err != nil || course.NextTaskID != "literal-bindings-001" {
				t.Fatalf("initial route: %+v %v", course, err)
			}
			start := run("start")
			workspace := start.Result.Workspace
			assertLearnerWorkspaceSafe(t, workspace)
			idePath := filepath.Join(workspace, ".idea", "inspectionProfiles", "metadata.go")
			if err := os.MkdirAll(filepath.Dir(idePath), 0700); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(idePath, []byte("INVALID GO IDE METADATA"), 0600); err != nil {
				t.Fatal(err)
			}
			oldWorkspace := workspace
			plan, err := s.Organize(false)
			if err != nil || len(plan.Moves) != 1 {
				t.Fatalf("organize: %+v %v", plan, err)
			}
			view, err := s.TaskView(fixture.id)
			if err != nil || view.WorkspaceUnavailable || view.Workspace == oldWorkspace {
				t.Fatalf("organized view: %+v %v", view, err)
			}
			workspace = view.Workspace
			idePath = filepath.Join(workspace, ".idea", "inspectionProfiles", "metadata.go")
			// Intentionally invalid editable code reaches only the disposable runner.
			if err := os.WriteFile(filepath.Join(workspace, fixture.editable), []byte("package subject\nthis is invalid Go\n"), 0600); err != nil {
				t.Fatal(err)
			}
			failed := run("check")
			if failed.Result.Verification.Passed || failed.Result.Verification.Diagnostic == "" {
				t.Fatalf("compiler failure missing: %+v", failed)
			}
			fullFailure := run("complete")
			if fullFailure.Result.Verification.Passed || fullFailure.Result.Verification.Diagnostic != "" {
				t.Fatal("full failure passed or leaked diagnostics")
			}
			// These valid submissions pass public/reference checks but fail hidden
			// graders (or miss required mutants). This reaches the private checker,
			// unlike the compiler-error case above.
			weak, err := os.ReadFile(filepath.Join(s.root, "tasks", "ready", fixture.id, fixture.weak))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, fixture.editable), weak, 0600); err != nil {
				t.Fatal(err)
			}
			if job := run("check"); !job.Result.Verification.Passed {
				t.Fatalf("weak fixture must pass public: %+v", job)
			}
			fullFailure = run("complete")
			if fullFailure.Result.Verification.Passed || fullFailure.Result.Verification.Diagnostic != "" {
				t.Fatal("hidden failure passed or leaked diagnostics")
			}
			payload, _ := json.Marshal(fullFailure)
			for _, private := range []string{"solution/reference", "tests/grader", "mutant", "go-learn-check"} {
				if strings.Contains(string(payload), private) {
					t.Fatalf("full result exposed %q", private)
				}
			}
			run("hint")
			reference, err := os.ReadFile(filepath.Join(s.root, "tasks", "ready", fixture.id, "solution", "reference", fixture.reference))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, fixture.editable), reference, 0600); err != nil {
				t.Fatal(err)
			}
			if job := run("check"); !job.Result.Verification.Passed {
				t.Fatalf("correct fixture public failed: %+v", job)
			}
			if job := run("complete"); !job.Result.Verification.Passed {
				t.Fatalf("correct fixture full failed: %+v", job)
			}
			view, err := s.TaskView(fixture.id)
			if err != nil || !view.Completed || view.Review == "" || len(view.UsedHints) != 1 || view.Attempts != 6 {
				t.Fatalf("completed view: %+v %v", view, err)
			}
			// A new UI and Service share the saved result; no in-memory job needed.
			restarted := testLocalUI(t, uiTestService(t, s.home))
			w := uiHTTP(restarted, "GET", "/api/task?id="+fixture.id, "", nil)
			var restored TaskView
			if err := json.Unmarshal(w.Body.Bytes(), &restored); err != nil {
				t.Fatal(err)
			}
			if w.Code != 200 || !restored.Completed || restored.Review == "" {
				t.Fatal("restart lost completion/review")
			}

			metadata, err := os.ReadFile(idePath)
			if err != nil || string(metadata) != "INVALID GO IDE METADATA" {
				t.Fatal("IDE metadata changed")
			}
		})
	}
}
