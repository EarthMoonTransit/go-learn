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
		{"dictionary-update-001", "clear-on-missing", []string{"absent changed map"}},
		{"dynamic-description-001", "partial-error-result", []string{"Describe subject.FancyNumber", "fancy 0.0"}},
		{"dynamic-description-001", "wrong-interface-priority", []string{"bothBox", "want \"box 9.0\""}},
		{"category-total-001", "all-categories-in-period", []string{"total got", "want1.5,nil"}},
		{"category-total-001", "success-reassigned", []string{"sentinel identity changed"}},
		{"checked-withdrawal-001", "success-reassigned", []string{"sentinel identity changed"}},
		{"reserved-copy-001", "zero-extra-alias", []string{"zero extra shares input"}},
		{"bounded-slice-view-001", "prefix-cleared", []string{"call changed input"}},
		{"writer-greeting-001", "close-on-error", []string{"closes got 1 want 0"}},
		{"writer-greeting-001", "close-on-short", []string{"closes got 1 want 0"}},
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

// Valid numerical alternatives must pass the same public/full path as the
// learner. These files are review fixtures, installed as data into a temp home.
func TestReviewedToleranceFixturesPassLearnerVerification(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"independent-scaled-copy-001", "value-average-001", "category-total-001"} {
		t.Run(id, func(t *testing.T) {
			canonical := filepath.Join(root, "tasks", "ready", id)
			before, err := tasking.TaskContentDigest(canonical)
			if err != nil {
				t.Fatal(err)
			}
			service, err := New(Config{Root: root, Home: t.TempDir()})
			if err != nil {
				t.Fatal(err)
			}
			workspace, err := service.Start(id)
			if err != nil {
				t.Fatal(err)
			}
			content, err := os.ReadFile(filepath.Join(canonical, "review", "tolerated-output.txt"))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, "subject", "subject.go"), content, 0600); err != nil {
				t.Fatal(err)
			}
			public, err := service.Check()
			if err != nil || !public.Passed {
				t.Fatalf("tolerated public result=%+v err=%v", public, err)
			}
			full, err := service.Complete()
			if err != nil || !full.Passed {
				t.Fatalf("tolerated full result=%+v err=%v", full, err)
			}
			after, err := tasking.TaskContentDigest(canonical)
			if err != nil || before != after {
				t.Fatalf("canonical changed before=%s after=%s err=%v", before, after, err)
			}
		})
	}
}
