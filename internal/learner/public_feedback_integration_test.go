//go:build taskintegration

package learner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

// Full mutant coverage cannot protect a public assertion when a hidden test
// also detects the same defect. These regressions deliberately use Check only.
func TestReviewedDefectsHavePublicFeedback(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	for _, fixture := range []struct {
		id, mutant  string
		diagnostics []string
	}{
		{"boolean-policies-001", "exclusive-or", []string{"AnyActive(", "got false want true"}},
		{"token-values-001", "trim-space", []string{"TokenValue(\"ten \")", "got 10 want 0"}},
		{"column-occupancy-001", "case-normalized", []string{"CountColumn key \"a\"", "got 1 want 0"}},
	} {
		t.Run(fixture.id, func(t *testing.T) {
			canonical := filepath.Join(root, "tasks", "ready", fixture.id)
			before, err := tasking.TaskContentDigest(canonical)
			if err != nil {
				t.Fatal(err)
			}
			service, err := New(Config{Root: root, Home: t.TempDir()})
			if err != nil {
				t.Fatal(err)
			}
			workspace, err := service.Start(fixture.id)
			if err != nil {
				t.Fatal(err)
			}
			install := func(parts ...string) {
				t.Helper()
				content, err := os.ReadFile(filepath.Join(append([]string{canonical}, parts...)...))
				if err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(workspace, "subject", "subject.go"), content, 0600); err != nil {
					t.Fatal(err)
				}
			}
			// Read source as data. The verifier executes only its disposable copy.
			install("solution", "mutants", fixture.mutant, "subject.go")
			result, err := service.Check()
			if err != nil || result.Passed {
				t.Fatalf("public defect check: result=%+v err=%v", result, err)
			}
			for _, want := range fixture.diagnostics {
				if !strings.Contains(result.Diagnostic, want) {
					t.Errorf("public diagnostic %q does not contain %q", result.Diagnostic, want)
				}
			}
			install("solution", "reference", "subject.go")
			result, err = service.Check()
			if err != nil || !result.Passed {
				t.Fatalf("reference public check: result=%+v err=%v", result, err)
			}
			after, err := tasking.TaskContentDigest(canonical)
			if err != nil || before != after {
				t.Fatalf("canonical task changed: before=%s after=%s err=%v", before, after, err)
			}
		})
	}
}
