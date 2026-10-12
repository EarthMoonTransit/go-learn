//go:build taskintegration

package learner

import (
	"io/fs"
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
	verify := integrationVerifier(t, root)
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
		{"package-inventory-001", "reset-sold", []string{"SetOnHand changed unrelated state", "want (9,7)"}},
		{"record-result-001", "trim-label", []string{"exact label:"}},
		{"error-context-wrap-001", "formatter-instead-of-error", []string{"Error text:", "want read: error text with original cause"}},
		{"bounded-copy-001", "reassign-sentinel", []string{"ErrLimit sentinel identity changed"}},
		{"bounded-copy-001", "close-on-success", []string{"successful copy", "closes got1,1 want0,0"}},
		{"buffered-flush-001", "empty-skips-flush", []string{"pending buffer", "want\"prefix|\",0"}},
		{"single-json-value-001", "partial-on-reject", []string{"partial object", "wantzero,ErrJSON"}},
		{"cleanup-order-001", "late-helper", []string{"Helper must precede acquire"}},
		{"cleanup-order-001", "rewritten-close-error", []string{"Errorf must preserve format and original close error"}},
		{"writer-double-001", "initial-snapshot", []string{"initial snapshot", "want empty/0"}},
		{"owned-values-001", "retained-worker", []string{"deadlock"}},
		{"select-ready-case-001", "closed-only-cancel", []string{"readable cancel", "want0,canceled"}},
		{"mutex-counter-001", "no-lock", []string{"DATA RACE"}},
		{"concurrent-counter-read-001", "unlocked-read", []string{"DATA RACE"}},
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
			service.verify = verify
			workspace, err := service.Start(fixture.id)
			if err != nil {
				t.Fatal(err)
			}
			install := func(parts ...string) {
				t.Helper()
				source := filepath.Join(append([]string{canonical}, parts...)...)
				err := filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
					if walkErr != nil {
						return walkErr
					}
					if entry.IsDir() {
						return nil
					}
					rel, err := filepath.Rel(source, path)
					if err != nil {
						return err
					}
					// A fixed-tests payload is rooted at subject; workspace patches
					// already carry package-relative paths for every editable file.
					if !strings.ContainsRune(rel, filepath.Separator) {
						rel = filepath.Join("subject", rel)
					}
					content, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					target := filepath.Join(workspace, rel)
					if err = os.MkdirAll(filepath.Dir(target), 0700); err != nil {
						return err
					}
					return os.WriteFile(target, content, 0600)
				})
				if err != nil {
					t.Fatal(err)
				}
			}
			// Read source as data. The verifier executes only its disposable copy.
			if fixture.id == "writer-double-001" && fixture.mutant == "initial-snapshot" {
				// A tooling-owned defect isolates the newly public baseline
				// without changing or republishing the accepted content package.
				install("solution", "reference")
				path := filepath.Join(workspace, "subject", "subject.go")
				content, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				old := "func (w *RecordingWriter) Snapshot() ([]byte, int) {"
				if strings.Count(string(content), old) != 1 {
					t.Fatal("snapshot regression fixture no longer matches reference API")
				}
				changed := strings.Replace(string(content), old, old+` if w.calls == 0 { return []byte("unexpected"), 1 };`, 1)
				if err := os.WriteFile(path, []byte(changed), 0600); err != nil {
					t.Fatal(err)
				}
			} else {
				install("solution", "mutants", fixture.mutant)
			}
			result, err := service.Check()
			if err != nil || result.Passed {
				t.Fatalf("public defect check: result=%+v err=%v", result, err)
			}
			for _, want := range fixture.diagnostics {
				if !strings.Contains(result.Diagnostic, want) {
					t.Errorf("public diagnostic %q does not contain %q", result.Diagnostic, want)
				}
			}
			install("solution", "reference")
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
	verify := integrationVerifier(t, root)
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
			service.verify = verify
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
