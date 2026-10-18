package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/fuzznative"
)

func scriptedLearnerFuzz(t *testing.T, scope LearnerVerificationScope, run func(fuzzRequest, int) fuzznative.Result) (LearnerVerificationResult, *scriptedFuzzRunner) {
	t.Helper()
	directory := t.TempDir()
	if err := os.WriteFile(filepath.Join(directory, "go.mod"), []byte("module task.local/learner-hex-001\n\ngo 1.27\n"), 0600); err != nil {
		t.Fatal(err)
	}
	loaded := learnerPackage{data: packageData{Located: LocatedTask{Path: directory}, Task: Task{ID: "learner-hex-001"}, TaskV2: TaskV2{Verification: VerificationPlan{Profile: "fuzz-seed-tests", RaceRequired: true, Vet: true, Limits: VerificationLimits{CommandTimeoutSeconds: 30, DeterminismRuns: 2}, Fuzz: &FuzzPlan{Grammar: "hex-codec-v1", Target: "FuzzHex", CorrectVariants: []string{"loop", "empty"}}}}}, layout: packageLayout{VerificationReference: []string{"canonical"}, VerificationCorrect: []correctVariantLayout{{ID: "loop", Files: []string{"loop"}}, {ID: "empty", Files: []string{"empty"}}}, Mutants: []mutantLayout{{ID: "one", Files: []string{"one"}}, {ID: "two", Files: []string{"two"}}}}}
	runner := &scriptedFuzzRunner{fakeExecutionRunner: newFakeExecutionRunner(), Run: run}
	return verifyLearnerFuzz(runner, loaded, []string{"learner"}, scope), runner
}

func TestFuzzLearnerPublicUsesEveryCorrectVariantAndRepeats(t *testing.T) {
	result, runner := scriptedLearnerFuzz(t, LearnerPublicChecks, func(q fuzzRequest, _ int) fuzznative.Result { return eligibleFuzz(q, false) })
	if !result.Passed || len(runner.Calls) != 6 {
		t.Fatalf("public correctness incomplete: %+v, %d calls", result, len(runner.Calls))
	}
	for _, q := range runner.Calls {
		if !q.Race || !q.Vet || q.ModulePath != "task.local/learner-hex-001" || q.Implementation[0] == "one" || q.Implementation[0] == "two" {
			t.Fatalf("wrong public execution boundary: %+v", q)
		}
	}
}

func TestFuzzLearnerPublicRejectsInvalidCorrectImplementationAssertion(t *testing.T) {
	for _, scope := range []LearnerVerificationScope{LearnerPublicChecks, LearnerFullChecks} {
		t.Run(string(scope), func(t *testing.T) {
			result, runner := scriptedLearnerFuzz(t, scope, func(q fuzzRequest, _ int) fuzznative.Result {
				r := eligibleFuzz(q, q.Implementation[0] == "empty")
				if r.Detected {
					r.Diagnostic = "empty values may have either nil representation"
				}
				return r
			})
			if result.Passed || result.FailureCategory != "public-behavior" || len(runner.Calls) != 6 {
				t.Fatalf("valid implementation incorrectly rejected without useful feedback: %+v %d calls", result, len(runner.Calls))
			}
			if scope == LearnerPublicChecks && !strings.Contains(result.Diagnostic, "nil representation") {
				t.Fatalf("public invariant feedback lost: %+v", result)
			}
			if scope == LearnerFullChecks && result.Diagnostic != "" {
				t.Fatal("full output exposed private implementation details")
			}
		})
	}
}

func TestFuzzLearnerFullDistinguishesInvalidFromSurvivingMutants(t *testing.T) {
	for _, kind := range []string{"strong", "survivor", "late-fault", "changing"} {
		t.Run(kind, func(t *testing.T) {
			result, runner := scriptedLearnerFuzz(t, LearnerFullChecks, func(q fuzzRequest, n int) fuzznative.Result {
				mutant := q.Implementation[0] == "one" || q.Implementation[0] == "two"
				r := eligibleFuzz(q, mutant)
				if q.Implementation[0] == "one" {
					switch kind {
					case "survivor":
						return eligibleFuzz(q, false)
					case "late-fault":
						r.Valid = false
						r.Canonical = ""
						r.Diagnostic = "internal-one raw race 98"
					case "changing":
						if n == 2 {
							r.Canonical = strings.Repeat("c", 64)
						}
					}
				}
				return r
			})
			if len(runner.Calls) != 10 || result.Diagnostic != "" {
				t.Fatalf("full verification skipped repeats or exposed hidden details: %+v %d calls", result, len(runner.Calls))
			}
			want := ""
			if kind == "survivor" {
				want = "hidden-verification"
			} else if kind != "strong" {
				want = "native-eligibility"
			}
			if result.Passed != (kind == "strong") || result.FailureCategory != want {
				t.Fatalf("wrong full outcome for %s: %+v", kind, result)
			}
			if strings.Contains(result.Summary, "internal-one") {
				t.Fatal("private mutant diagnostic leaked")
			}
		})
	}
}
