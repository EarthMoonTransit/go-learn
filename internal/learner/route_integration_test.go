//go:build taskintegration

package learner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

func TestRealOpeningLessonRunsThroughLearnerWorkspace(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	service, err := New(Config{Root: root, Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	first, err := service.Next()
	if err != nil || first.Task.ID != "literal-bindings-001" {
		t.Fatalf("first=%+v err=%v", first, err)
	}
	canonical := filepath.Join(root, "tasks", "ready", first.Task.ID)
	before, err := tasking.TaskContentDigest(canonical)
	if err != nil {
		t.Fatal(err)
	}
	workspace, err := service.Start(first.Task.ID)
	if err != nil {
		t.Fatal(err)
	}
	next, err := service.Next()
	if err != nil || !next.Resume {
		t.Fatalf("next=%+v err=%v", next, err)
	}
	result, err := service.Check()
	if err != nil || result.Passed {
		t.Fatalf("starter=%+v err=%v", result, err)
	}
	// Each fixture violates one explicit binding requirement. Full-check mutant
	// coverage alone cannot detect a public suite that silently omits a binding.
	for _, fixture := range []struct{ mutant, diagnostic string }{
		{"floating-count", "count = 12 (float64), want int 12"},
		{"integer-scale", "scale = 1 (int), want float64 1.25"},
		{"rune-label", "label = 71 (int32), want string G"},
		{"disabled", "enabled = false (bool), want bool true"},
	} {
		t.Run(fixture.mutant, func(t *testing.T) {
			content, err := os.ReadFile(filepath.Join(canonical, "solution", "mutants", fixture.mutant, "subject", "bindings.go"))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, "subject", "bindings.go"), content, 0600); err != nil {
				t.Fatal(err)
			}
			result, err := service.Check()
			if err != nil || result.Passed || !strings.Contains(result.Diagnostic, fixture.diagnostic) {
				t.Fatalf("public check must explain the violated requirement: result=%+v err=%v", result, err)
			}
		})
	}
	// Read canonical reference as data; execution occurs only in the learner
	// runner's disposable copy, never in the canonical source tree.
	content, err := os.ReadFile(filepath.Join(canonical, "solution", "reference", "subject", "bindings.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(workspace, "subject", "bindings.go"), content, 0600); err != nil {
		t.Fatal(err)
	}
	result, err = service.Check()
	if err != nil || !result.Passed {
		t.Fatalf("reference public=%+v err=%v", result, err)
	}
	result, err = service.Complete()
	if err != nil || !result.Passed {
		t.Fatalf("complete=%+v err=%v", result, err)
	}
	next, err = service.Next()
	if err != nil || next.Task.ID != "zero-results-001" {
		t.Fatalf("next=%+v err=%v", next, err)
	}
	progress, err := service.Progress()
	if err != nil || !progress.Tasks[first.Task.ID].Completed || progress.Tasks[first.Task.ID].Attempts != 7 {
		t.Fatalf("progress=%+v err=%v", progress, err)
	}
	after, err := tasking.TaskContentDigest(canonical)
	if err != nil || before != after {
		t.Fatal("canonical task changed")
	}
}
