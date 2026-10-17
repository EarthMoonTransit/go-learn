//go:build taskintegration

package learner

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func TestWorkerRevisionContinuesOldWorkspaceWithRealVerifier(t *testing.T) {
	s := workerRevisionService(t)
	root, err := filepath.EvalSymlinks(s.root)
	if err != nil {
		t.Fatal(err)
	}
	s.root = root
	historicalWorkerFixture(t, s.root)
	if err := saveProgress(s.home, Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}); err != nil {
		t.Fatal(err)
	}
	const id = "worker-completion-tests-001"
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	applyWorkerScaffoldFixture(t, s.root)
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
		t.Fatalf("gated example public check: %+v, %v", result, err)
	}
	if result, err := s.Complete(); err != nil || result.Passed || result.Diagnostic != "" {
		t.Fatalf("proposed starter must retain learner-owned cases: %+v, %v", result, err)
	}
	wrong := strings.ReplaceAll(string(starter), "work.calls != 1", "work.calls != 99")
	wrong = strings.ReplaceAll(wrong, "phase callback calls = %d, want 1", "phase callback calls = %d, want 99")
	if wrong == string(starter) {
		t.Fatal("gated example fixture changed")
	}
	install([]byte(wrong))
	if result, err := s.Check(); err != nil || result.Passed || !strings.Contains(result.Diagnostic, "TestWorkerBlockedPhase") || !strings.Contains(result.Diagnostic, "blocked phase callback calls = 1, want 99") || !strings.Contains(result.Diagnostic, "completed phase callback calls = 1, want 99") {
		t.Fatalf("missing named phase actual/want feedback: %+v, %v", result, err)
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
		t.Fatalf("reference must detect all eight existing defects: %+v, %v", result, err)
	}
	p, err := s.Progress()
	if err != nil || len(p.Completions) != 1 {
		t.Fatalf("expected one new completion: %+v, %v", p.Completions, err)
	}
	evidence := p.Completions[0]
	if evidence.TaskID != id || evidence.RouteRole != "introduction" || !slices.Equal(evidence.Skills, []string{"goroutines.launch-lifecycle", "channels.operations-blocking"}) {
		t.Fatalf("revision must not add assessed skills or change role: %+v", evidence)
	}
}

// The supplied example must report blocking mistakes as assertions, and still
// leave premature completion and overlapping independence to the learner.
// Only this temporary repository's checker-owned binding is changed.
func TestWorkerGatedExampleHasBoundedDefectFeedback(t *testing.T) {
	s := workerRevisionService(t)
	root, err := filepath.EvalSymlinks(s.root)
	if err != nil {
		t.Fatal(err)
	}
	s.root = root
	historicalWorkerFixture(t, s.root)
	applyWorkerScaffoldFixture(t, s.root)
	verify := integrationVerifier(t, s.root)
	const id = "worker-completion-tests-001"
	canonical := filepath.Join(s.root, "tasks/ready", id)
	for _, c := range []struct {
		name       string
		passed     bool
		diagnostic string
	}{
		{"inline-work", false, "blocked phase: Start has not returned, want return before release"},
		{"never-close", false, "completion channel did not close"},
		{"close-before-work", true, ""},
		{"skip-work", false, "callback calls0 want1"},
		{"double-work", false, "callback calls2 want1"},
		{"nil-channel", false, "nil completion"},
		{"shared-active-completion", true, ""},
		{"send-before-close", false, "completion sent a value instead of closing"},
	} {
		t.Run(c.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(canonical, "verification/mutants", c.name, "subject.go"))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(canonical, "verification/reference/subject.go"), data, 0600); err != nil {
				t.Fatal(err)
			}
			data, err = os.ReadFile(filepath.Join(canonical, "provenance.yaml"))
			if err != nil {
				t.Fatal(err)
			}
			var provenance tasking.ProvenanceV2
			if err := yaml.Unmarshal(data, &provenance); err != nil {
				t.Fatal(err)
			}
			digest, err := tasking.TaskContentDigest(canonical)
			if err != nil {
				t.Fatal(err)
			}
			provenance.Project.ContentDigest = digest
			writeYAMLFixture(t, filepath.Join(canonical, "provenance.yaml"), provenance)
			files, err := tasking.LearnerWorkspaceFiles(s.root, id)
			if err != nil {
				t.Fatal(err)
			}
			workspace := t.TempDir()
			for _, file := range files {
				target := filepath.Join(workspace, filepath.FromSlash(file.Path))
				if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(target, file.Content, 0600); err != nil {
					t.Fatal(err)
				}
			}
			result, err := verify(s.root, id, workspace, tasking.LearnerPublicChecks)
			if err != nil || result.Passed != c.passed {
				t.Fatalf("gated example %s: %+v, %v", c.name, result, err)
			}
			if !c.passed && (result.FailureCategory != "test" || !strings.Contains(result.Diagnostic, c.diagnostic)) {
				t.Fatalf("need useful bounded assertion, not compile error/timeout: %+v", result)
			}
		})
	}
}
