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
		{"launch-provided-job-001", "duplicate-runner", []string{"held job 0 calls = 2, want 1"}},
		{"launch-provided-job-001", "inline-runner", []string{"StartJob blocked with held work, want return before Release"}},
		{"launch-provided-job-001", "premature-release", []string{"held job 0 completed before Release, want pending"}},
		{"launch-provided-job-001", "shared-job", []string{"job 1 returned supplied handle = false, want true"}},
		{"launch-provided-job-001", "skip-runner", []string{"held job 0 calls = 0, want 1"}},
		{"launch-provided-job-001", "wrong-handle", []string{"job 0 returned supplied handle = false, want true"}},
		{"replace-pending-value-001", "close-channel", []string{"channel after drain = closed, want open"}},
		{"replace-pending-value-001", "drop-replacement", []string{"pending value = absent, want next 7"}},
		{"replace-pending-value-001", "restore-old", []string{"pending value = 0, want next 7"}},
		{"replace-pending-value-001", "return-next", []string{"ReplacePending(0,7) returned 7, want old 0"}},
		{"replace-pending-value-001", "return-zero", []string{"ReplacePending(-3,8) returned 0, want old -3"}},
		{"replace-pending-value-001", "send-before-receive", []string{"ReplacePending(0,7) blocked on full buffer, want receive before send"}},
		{"merge-two-streams-001", "buffer-after-right-close", []string{"available input 23 output = absent, want progress from either open input"}},
		{"merge-two-streams-001", "close-on-first-input", []string{"one input closed output = (0,false), want pending until remaining input"}},
		{"merge-two-streams-001", "keep-closed-case", []string{"one input closed output = (0,false), want pending until remaining input"}},
		{"merge-two-streams-001", "reverse-left-order", []string{"available input 11 output = absent, want progress from either open input"}},
		{"merge-two-streams-001", "serial-inputs", []string{"available input 21 output = absent, want progress from either open input"}},
		{"merge-two-streams-001", "skip-second-input", []string{"available input 21 output = absent, want progress from either open input"}},
		{"merge-two-streams-001", "zero-on-close", []string{"one input closed output = (0,true), want pending until remaining input"}},
		{"cancelable-forward-001", "background-context", []string{"after cancellation/EOF output = pending, want closed"}},
		{"cancelable-forward-001", "close-caller-input", []string{"caller input = closed, want caller-owned open channel"}},
		{"cancelable-forward-001", "collect-until-eof", []string{"live input output = pending, want value 9 before EOF"}},
		{"cancelable-forward-001", "drop-zero", []string{"forwarded order/occurrences = [3 -2 3], want [0 3 -2 3]"}},
		{"cancelable-forward-001", "leave-output-open", []string{"after cancellation/EOF output = pending, want closed"}},
		{"cancelable-forward-001", "missing-receive-cancel", []string{"after cancellation/EOF output = pending, want closed"}},
		{"cancelable-forward-001", "missing-send-cancel", []string{"after cancellation/EOF output = 7, want closed after queued values"}},
		{"cancelable-forward-001", "retain-input-receiver", []string{"stage still receives caller input after cancellation, want no retained receiver"}},
		{"cancelable-forward-001", "reverse-input", []string{"forwarded order/occurrences = [3 -2 3 0], want [0 3 -2 3]"}},
		{"timeout-operation-001", "background-parent", []string{"child value = <nil>, want marker"}},
		{"timeout-operation-001", "call-twice", []string{"operation calls = 2, want 1"}},
		{"timeout-operation-001", "cancel-before-operation", []string{"child before operation returns = context canceled, want live context"}},
		{"timeout-operation-001", "doubled-budget", []string{"child deadline =", "want ("}},
		{"timeout-operation-001", "no-release", []string{"child after return = live, want canceled/released"}},
		{"timeout-operation-001", "no-timeout", []string{"child deadline =", "want ("}},
		{"timeout-operation-001", "release-only-on-error", []string{"child after return = live, want canceled/released"}},
		{"timeout-operation-001", "replace-operation-error", []string{"operation result = (41,<nil>), want unchanged (41,operation result)"}},
		{"timeout-operation-001", "return-before-cleanup", []string{"wrapper returned before operation cleanup, want pending until operation returns"}},
		{"timeout-operation-001", "reuse-earlier-parent", []string{"operation received parent, want distinct derived context"}},
		{"timeout-operation-001", "skip-canceled-parent", []string{"operation result = (0,context canceled), want unchanged (19,operation result)"}},
		{"timeout-operation-001", "successful-result-canceled", []string{"operation result = (41,context canceled), want unchanged (41,<nil>)"}},
		{"timeout-operation-001", "zero-result", []string{"operation result = (0,operation result), want unchanged (41,operation result)"}},
		{"gcd-reduction-001", "subtraction-only", []string{"GCD remainder reduction in loop = false, want true"}},
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
		{"formatter-contract-001", "omit-right", []string{"Pair.Format(\"L\",\"R\") = \"L:\", want \"L:R\""}},
		{"formatter-contract-001", "double-format", []string{"Format calls = 2, want 1"}},
		{"formatter-contract-001", "trim-result", []string{"Render caller formatter = \"value:\", want unchanged"}},
		{"resource-finalize-001", "close-before-write", []string{"calls/order = [close write], want [write close]"}},
		{"resource-finalize-001", "deduplicate-error-text", []string{"close wrapper identity = false, want true", "close cause identity = false, want true"}},
		{"resource-finalize-001", "empty-skips-calls", []string{"calls/order = [], want [write close]"}},
		{"resource-finalize-001", "ignore-close-error", []string{"close identity = false, want true"}},
		{"resource-finalize-001", "ignore-short-write", []string{"short-write identity = false, want true"}},
		{"resource-finalize-001", "lose-write-context", []string{"write context =", "want error text naming write"}},
		{"resource-finalize-001", "overwrite-write-error", []string{"write identity = false, want true", "short-write identity = false, want true"}},
		{"resource-finalize-001", "recover-write-panic", []string{"write panic = <nil>, want original"}},
		{"resource-finalize-001", "retry-write", []string{"calls/order = [write write close], want [write close]"}},
		{"resource-finalize-001", "skip-close-on-write-error", []string{"calls/order = [write], want [write close]"}},
		{"resource-finalize-001", "wrong-data", []string{"Write data = [], want [97 98 99]"}},
	} {
		t.Run(fixture.id+"/"+fixture.mutant, func(t *testing.T) {
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
	for _, fixture := range []struct{ id, file string }{
		{"independent-scaled-copy-001", "tolerated-output.txt"},
		{"value-average-001", "tolerated-output.txt"},
		{"category-total-001", "tolerated-output.txt"},
		{"gcd-reduction-001", "tolerated-output.txt"},
		{"merge-two-streams-001", "tolerated-output.txt"},
		{"merge-two-streams-001", "tolerated-output-function-declaration.txt"},
		{"cancelable-forward-001", "tolerated-output.txt"},
		{"timeout-operation-001", "tolerated-output.txt"},
	} {
		t.Run(fixture.id+"/"+fixture.file, func(t *testing.T) {
			id := fixture.id
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
			content, err := os.ReadFile(filepath.Join(canonical, "review", fixture.file))
			if err != nil {
				t.Fatal(err)
			}
			filename := "subject.go"
			if id == "gcd-reduction-001" {
				filename = "implementation.go"
			}
			if err := os.WriteFile(filepath.Join(workspace, "subject", filename), content, 0600); err != nil {
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
