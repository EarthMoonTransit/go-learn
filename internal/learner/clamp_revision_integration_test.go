//go:build taskintegration

package learner

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestClampRevisionUsesExistingWorkspaceAndRealVerifier(t *testing.T) {
	s := goLandTestService(t)
	root, err := filepath.EvalSymlinks(s.root)
	if err != nil {
		t.Fatal(err)
	}
	s.root = root
	prepareHistoricalClampFixture(t, s.root)
	if err := saveProgress(s.home, Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}); err != nil {
		t.Fatal(err)
	}
	const id = "clamp-tests-001"
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	applyClampRevisionFixture(t, s.root)
	s.verify = integrationVerifier(t, s.root)
	if got, err := s.Resume(id); err != nil || got != workspace {
		t.Fatalf("old workspace must resume without reset: %s, %v", got, err)
	}
	if result, err := s.Check(); err != nil || !result.Passed {
		t.Fatalf("old starter must remain a valid public submission: %+v, %v", result, err)
	}
	if result, err := s.Complete(); err != nil || result.Passed || result.Diagnostic != "" {
		t.Fatalf("old incomplete starter must still fail full verification privately: %+v, %v", result, err)
	}
	starter, err := os.ReadFile(filepath.Join(s.root, "tasks/ready", id, "starter/subject_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	install := func(content string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(workspace, "subject/subject_test.go"), []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
	}
	install(string(starter))
	if result, err := s.Check(); err != nil || !result.Passed {
		t.Fatalf("both proposed examples must compile and pass: %+v, %v", result, err)
	}
	for _, tc := range []struct{ name, compare, expected, diagnostic string }{
		{"TestClampInsideRange", "got != 5", "want 5", "Clamp(5,0,10)=5, want 99"},
		{"TestClampBelowRange", "got != -1", "want -1", "Clamp(-3,-1,8)=-1, want 99"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// A deliberately wrong learner expectation proves discovery and
			// actual/expected public feedback without exposing hidden mutants.
			wrong := strings.Replace(string(starter), tc.compare, "got != 99", 1)
			wrong = strings.Replace(wrong, tc.expected, "want 99", 1)
			install(wrong)
			result, err := s.Check()
			if err != nil || result.Passed || !strings.Contains(result.Diagnostic, tc.name) || !strings.Contains(result.Diagnostic, tc.diagnostic) {
				t.Fatalf("missing discovered assertion diagnostic: %+v, %v", result, err)
			}
		})
	}
	// The proposed scaffold is guidance, not a new grading requirement: a
	// single test with plain Fail calls must still pass public and full checks.
	reference, err := os.ReadFile(filepath.Join(s.root, "tasks/ready", id, "solution/reference/subject_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	plain := regexp.MustCompile(`t\.Fatalf\([^\n]*\)`).ReplaceAllString(string(reference), "t.Fail()")
	plain = regexp.MustCompile(`\n}\nfunc TestClamp\w+\(t \*testing.T\) \{`).ReplaceAllString(plain, "")
	if strings.Count(plain, "func Test") != 1 || strings.Contains(plain, "Fatalf") {
		t.Fatal("fixture must use one test and no formatted diagnostics")
	}
	install(plain)
	if result, err := s.Check(); err != nil || !result.Passed {
		t.Fatalf("alternative organization failed public checks: %+v, %v", result, err)
	}
	if result, err := s.Complete(); err != nil || !result.Passed {
		t.Fatalf("alternative organization failed full checks: %+v, %v", result, err)
	}
	progress, err := s.Progress()
	if err != nil || len(progress.Completions) != 1 {
		t.Fatalf("expected one new successful completion: %+v, %v", progress.Completions, err)
	}
	evidence := progress.Completions[0]
	if evidence.TaskID != id || evidence.RouteRole != "practice" || !slices.Equal(evidence.Skills, []string{"expressions.comparison", "testing.basics-diagnostics"}) {
		t.Fatalf("new completion must record the revised skills and original role: %+v", evidence)
	}
}
