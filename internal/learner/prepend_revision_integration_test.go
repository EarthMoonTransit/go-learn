//go:build taskintegration

package learner

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestPrependRevisionContinuesOldWorkspaceWithRealVerifier(t *testing.T) {
	s := prependRevisionService(t)
	root, err := filepath.EvalSymlinks(s.root)
	if err != nil {
		t.Fatal(err)
	}
	s.root = root
	historicalPrependFixture(t, s.root)
	if err := saveProgress(s.home, Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}); err != nil {
		t.Fatal(err)
	}
	const id = "prepend-values-tests-001"
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	applyPrependBridgeFixture(t, s.root)
	s.verify = integrationVerifier(t, s.root)
	if got, err := s.Resume(id); err != nil || got != workspace {
		t.Fatalf("old workspace must resume: %s, %v", got, err)
	}
	if result, err := s.Check(); err != nil || !result.Passed {
		t.Fatalf("historical starter public check: %+v, %v", result, err)
	}
	if result, err := s.Complete(); err != nil || result.Passed || result.Diagnostic != "" {
		t.Fatalf("historical incomplete tests must fail privately: %+v, %v", result, err)
	}
	install := func(content []byte) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(workspace, "subject/subject_test.go"), content, 0600); err != nil {
			t.Fatal(err)
		}
	}
	starter, err := os.ReadFile(filepath.Join(s.root, "tasks/ready", id, "starter/subject_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	install(starter)
	if result, err := s.Check(); err != nil || !result.Passed {
		t.Fatalf("both variadic examples must compile and pass: %+v, %v", result, err)
	}
	if result, err := s.Complete(); err != nil || result.Passed || result.Diagnostic != "" {
		t.Fatalf("proposed starter must retain uncovered private defects: %+v, %v", result, err)
	}
	for _, tc := range []struct{ name, expected, diagnostic string }{
		{"TestStartingCase01", "[]int{5, 8, 3, 2, 6}", "PrependItems: got [5 8 3 2 6], want [99]"},
		{"TestPrependExpandedPrefix", "[]int{9, 4, 2}", "PrependItems expanded prefix: got [9 4 2], want [99]"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wrong := strings.ReplaceAll(string(starter), tc.expected, "[]int{99}")
			install([]byte(wrong))
			result, err := s.Check()
			if err != nil || result.Passed || !strings.Contains(result.Diagnostic, tc.name) || !strings.Contains(result.Diagnostic, tc.diagnostic) {
				t.Fatalf("missing public actual/expected feedback for %s: %+v, %v", tc.name, result, err)
			}
		})
	}
	reference, err := os.ReadFile(filepath.Join(s.root, "tasks/ready", id, "solution/reference/subject_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	install(reference)
	if result, err := s.Check(); err != nil || !result.Passed {
		t.Fatalf("reference public check: %+v, %v", result, err)
	}
	if result, err := s.Complete(); err != nil || !result.Passed {
		t.Fatalf("reference must detect all existing defects: %+v, %v", result, err)
	}
	p, err := s.Progress()
	if err != nil || len(p.Completions) != 1 {
		t.Fatalf("expected one new completion: %+v, %v", p.Completions, err)
	}
	evidence := p.Completions[0]
	if evidence.TaskID != id || evidence.RouteRole != "introduction" || !slices.Equal(evidence.Skills, []string{"functions.variadic"}) {
		t.Fatalf("prerequisite correction must not add assessed skills or change role: %+v", evidence)
	}
}
