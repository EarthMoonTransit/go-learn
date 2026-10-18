//go:build taskintegration

package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// All authored source reaches native Go only through the disposable runner.
// This infrastructure fixture grants no catalog or coverage credit.
func TestNativeFuzzTaskCheckAndPublicFeedback(t *testing.T) {
	root, data := fuzzFixtureRepository(t)
	before, err := TaskContentDigest(data.Located.Path)
	if err != nil {
		t.Fatal(err)
	}
	session, err := NewCheckSession(root)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := session.Close(); err != nil {
			t.Error(err)
		}
	})
	checks, err := session.CheckAll([]State{StateDraft})
	if err != nil {
		t.Fatal(err)
	}
	for _, checked := range checks.Tasks {
		if !checked.Report.Passed {
			t.Fatalf("mixed-profile check %s failed: %+v", checked.TaskID, checked.Report.Gates)
		}
		if checked.TaskID == syntheticFuzzID && checked.Report.Mutants.Detected != 2 {
			t.Fatalf("missing eligible defect detection: %+v", checked.Report)
		}
	}
	metrics := checks.Execution
	if len(checks.Tasks) != 2 || metrics.FuzzRuns != 20 || metrics.FuzzNativeRuns != 20 || metrics.FuzzHelperRuns != 20 || metrics.VariantRuns != 5 || metrics.RuntimeQueries != 1 || metrics.ChildGoRuns != 6+metrics.FuzzGoRuns || metrics.Workspaces != 6+metrics.FuzzWorkspaces {
		t.Fatalf("incorrect v4 mixed-profile execution accounting: %+v", metrics)
	}
	if err := ValidateCheckCollection(root, checks, []State{StateDraft}); err != nil {
		t.Fatal(err)
	}
	problems, layout := checkPackageContractV2(root, data)
	if len(problems) != 0 {
		t.Fatal(problems)
	}
	loaded := learnerPackage{data: data, layout: layout}
	workspace := t.TempDir()
	path := filepath.Join(workspace, "fuzz_test.go")
	reference := mustRead(t, layout.Reference[0])
	// A nil-only requirement rejects a value-equivalent declared correct
	// implementation. This failure must be visible during public Check.
	bad := strings.Replace(string(reference), "wire:=subject.Encode(input);", "wire:=subject.Encode(input);if len(input)==0&&wire!=nil{t.Error(\"invalid nil-only assumption\")};", 1)
	mustWrite(t, path, []byte(bad))
	result := verifyLearnerFuzz(session.runner, loaded, []string{path}, LearnerPublicChecks)
	if result.Passed || result.FailureCategory != "public-behavior" || !strings.Contains(result.Diagnostic, "invalid nil-only assumption") {
		t.Fatalf("missing actionable public assertion feedback: %+v", result)
	}
	mustWrite(t, path, reference)
	result = verifyLearnerFuzz(session.runner, loaded, []string{path}, LearnerPublicChecks)
	if !result.Passed {
		t.Fatalf("corrected invariant still rejected: %+v", result)
	}
	// A late native fault after an assertion must invalidate detection through
	// the tasking boundary, rather than gaining mutant credit.
	faulty := filepath.Join(workspace, "hex.go")
	mustWrite(t, faulty, []byte("package subject\nfunc Encode(in []byte)[]byte{return []byte{0}}\nfunc Decode(in []byte)([]byte,bool){panic(\"late native fault\")}\n"))
	request := fuzzRequest{variantRequest: variantRequest{TaskID: syntheticFuzzID, GoVersion: session.policy.GoLanguageVersion, TimeoutSeconds: 60, Tests: []string{path}, Implementation: []string{faulty}}, ModulePath: "task.local/" + syntheticFuzzID, Plan: *data.TaskV2.Verification.Fuzz}
	fault := session.runner.RunFuzz(request)
	if fault.Valid || fault.Detected {
		t.Fatalf("late fault counted as detection: %+v", fault)
	}
	// Rejection of unsupported source precedes any authored native execution.
	for _, test := range []struct {
		name, source string
		diagnostics  []string
	}{
		{"wrong-target", strings.Replace(string(reference), "func FuzzHex(", "func FuzzOther(", 1), []string{"frame."}},
		{"no-registration", strings.Replace(string(reference), "f.Fuzz(", "f.Other(", 1), []string{"type-check original source:", "f.Other undefined"}},
		{"forbidden-import", strings.Replace(string(reference), "\"encoding/hex\"", "\"os\";\"encoding/hex\"", 1), []string{"frame."}},
		{"ignored-name", string(reference), []string{"frame."}},
	} {
		t.Run(test.name, func(t *testing.T) {
			name := "fuzz_test.go"
			if test.name == "ignored-name" {
				name = "_fuzz_test.go"
			}
			submitted := filepath.Join(workspace, name)
			mustWrite(t, submitted, []byte(test.source))
			request.Tests = []string{submitted}
			before := session.runner.Metrics()
			rejected := session.runner.RunFuzz(request)
			public := verifyLearnerFuzz(session.runner, loaded, []string{submitted}, LearnerPublicChecks)
			after := session.runner.Metrics()
			if rejected.Valid || rejected.Detected || rejected.Diagnostic == "" || public.Passed || public.FailureCategory != "native-eligibility" || after.FuzzNativeRuns != before.FuzzNativeRuns {
				t.Fatalf("unsupported source reached authored native execution: %+v public=%+v before=%+v after=%+v", rejected, public, before, after)
			}
			for _, diagnostic := range test.diagnostics {
				if !strings.Contains(public.Diagnostic, diagnostic) {
					t.Fatalf("public diagnostic %q omits %q", public.Diagnostic, diagnostic)
				}
			}
		})
	}
	after, err := TaskContentDigest(data.Located.Path)
	if err != nil || before != after {
		t.Fatalf("source fixture was mutated: %v", err)
	}
	if _, err := os.Stat(filepath.Join(data.Located.Path, "subject")); !os.IsNotExist(err) {
		t.Fatalf("native workspace leaked into source fixture: %v", err)
	}
}
