package tasking

import (
	"fmt"
	"strings"
	"testing"

	"go-practice-bank/internal/fuzznative"
)

type scriptedFuzzRunner struct {
	*fakeExecutionRunner
	Calls []fuzzRequest
	Run   func(fuzzRequest, int) fuzznative.Result
	seen  map[string]int
}

func (runner *scriptedFuzzRunner) RunFuzz(request fuzzRequest) fuzznative.Result {
	runner.Calls = append(runner.Calls, request)
	key := request.Tests[0] + "/" + request.Implementation[0]
	if runner.seen == nil {
		runner.seen = map[string]int{}
	}
	runner.seen[key]++
	return runner.Run(request, runner.seen[key])
}

func eligibleFuzz(request fuzzRequest, detected bool) fuzznative.Result {
	assertions := 0
	diagnostic := ""
	canonical := strings.Repeat("a", 64)
	if detected {
		assertions = 1
		diagnostic = "wire invariant mismatch"
		canonical = strings.Repeat("b", 64)
	}
	return fuzznative.Result{Valid: true, Detected: detected, Canonical: canonical, Diagnostic: diagnostic, Vetted: request.Vet, Callbacks: 1, Calls: 2, Assertions: assertions}
}

func ordinaryFuzzOutcome(request fuzzRequest, _ int) fuzznative.Result {
	mutant := request.Implementation[0] == "one" || request.Implementation[0] == "two"
	detected := mutant && (request.Tests[0] == "reference" || request.Implementation[0] == "one")
	return eligibleFuzz(request, detected)
}

func checkScriptedFuzz(t *testing.T, run func(fuzzRequest, int) fuzznative.Result) (QualityReport, *scriptedFuzzRunner) {
	t.Helper()
	runner := &scriptedFuzzRunner{fakeExecutionRunner: newFakeExecutionRunner(), Run: run}
	session, err := newCheckSessionWithRunner(repositoryRoot(t), runner)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	plan := VerificationPlan{Profile: "fuzz-seed-tests", RaceRequired: true, Vet: true, RequiredMutants: []string{"one", "two"}, Limits: VerificationLimits{CommandTimeoutSeconds: 30, DeterminismRuns: 2}, Fuzz: &FuzzPlan{Grammar: "hex-codec-v1", Target: "FuzzHex", CorrectVariants: []string{"loop", "empty"}}}
	data := packageData{Task: Task{ID: "scripted-hex-001"}, TaskV2: TaskV2{Verification: plan}}
	layout := packageLayout{Starter: []string{"starter"}, Reference: []string{"reference"}, VerificationReference: []string{"canonical"}, VerificationCorrect: []correctVariantLayout{{ID: "loop", Files: []string{"loop"}}, {ID: "empty", Files: []string{"empty"}}}, Mutants: []mutantLayout{{ID: "one", Files: []string{"one"}}, {ID: "two", Files: []string{"two"}}}}
	c := checker{report: QualityReport{State: StateDraft, Mutants: MutantSummary{Survived: []string{}}}}
	session.executeV2Fuzz(&c, data, layout)
	c.finish()
	return c.report, runner
}

func TestFuzzOrchestrationRepeatsAllCorrectVariantsAndMutants(t *testing.T) {
	report, runner := checkScriptedFuzz(t, ordinaryFuzzOutcome)
	if !report.Passed || report.Mutants.Detected != 2 || report.StarterStatus != "pass" || report.ReferenceStatus != "pass" || report.RaceStatus != "pass" {
		t.Fatalf("valid finite plan failed: %+v", report)
	}
	if len(runner.Calls) != 20 || len(runner.seen) != 10 {
		t.Fatalf("missing repeated implementation checks: %d calls %v", len(runner.Calls), runner.seen)
	}
	for key, count := range runner.seen {
		if count != 2 {
			t.Fatalf("%s repeated %d times", key, count)
		}
	}
	for _, request := range runner.Calls {
		if !request.Race || request.ModulePath != "task.local/scripted-hex-001" || request.Plan.Grammar != "hex-codec-v1" || request.Plan.Target != "FuzzHex" {
			t.Fatalf("lost declared execution identity: %+v", request)
		}
		mutant := request.Implementation[0] == "one" || request.Implementation[0] == "two"
		if request.Vet == mutant {
			t.Fatalf("vet must apply to every correct invocation: %+v", request)
		}
	}
	if runner.metrics.VariantRuns != 0 {
		t.Fatal("fuzz orchestration used the generic unobserved variant runner")
	}
}

func TestFuzzOrchestrationDoesNotCreditInvalidOrChangingEvidence(t *testing.T) {
	for _, test := range []struct {
		name, gate   string
		mutate       func(fuzzRequest, int, fuzznative.Result) fuzznative.Result
		wantDetected int
	}{
		{"invalid-starter-even-with-another-survivor", "starter-full", func(q fuzzRequest, _ int, r fuzznative.Result) fuzznative.Result {
			if q.Tests[0] == "starter" && q.Implementation[0] == "one" {
				r.Valid = false
				r.Canonical = ""
				r.Diagnostic = "late panic after assertion"
			}
			return r
		}, 2},
		{"late-race-after-real-reference-assertion", "mutant-coverage", func(q fuzzRequest, _ int, r fuzznative.Result) fuzznative.Result {
			if q.Tests[0] == "reference" && q.Implementation[0] == "one" {
				r.Valid = false
				r.Canonical = ""
				r.Diagnostic = "raw race exit 98"
			}
			return r
		}, 1},
		{"changing-native-observation", "mutant-coverage", func(q fuzzRequest, n int, r fuzznative.Result) fuzznative.Result {
			if q.Tests[0] == "reference" && q.Implementation[0] == "one" && n == 2 {
				r.Canonical = strings.Repeat("c", 64)
			}
			return r
		}, 1},
		{"reference-kills-only-one-repetition", "mutant-coverage", func(q fuzzRequest, n int, r fuzznative.Result) fuzznative.Result {
			if q.Tests[0] == "reference" && q.Implementation[0] == "one" && n == 2 {
				return eligibleFuzz(q, false)
			}
			return r
		}, 1},
		{"missing-native-observation", "mutant-coverage", func(q fuzzRequest, _ int, r fuzznative.Result) fuzznative.Result {
			if q.Tests[0] == "reference" && q.Implementation[0] == "one" {
				r.Canonical = ""
			}
			return r
		}, 1},
		{"starter-already-detects-every-mutant", "starter-full", func(q fuzzRequest, _ int, r fuzznative.Result) fuzznative.Result {
			if q.Tests[0] == "starter" && q.Implementation[0] == "two" {
				return eligibleFuzz(q, true)
			}
			return r
		}, 2},
	} {
		t.Run(test.name, func(t *testing.T) {
			report, _ := checkScriptedFuzz(t, func(q fuzzRequest, n int) fuzznative.Result { return test.mutate(q, n, ordinaryFuzzOutcome(q, n)) })
			assertGateFails(t, report, test.gate)
			if report.Mutants.Detected != test.wantDetected {
				t.Fatalf("wrong Detection credit: got %d want %d", report.Mutants.Detected, test.wantDetected)
			}
		})
	}
}

func TestFuzzReferenceCorrectnessPrecedesPrivateDetection(t *testing.T) {
	report, runner := checkScriptedFuzz(t, func(q fuzzRequest, n int) fuzznative.Result {
		if q.Tests[0] == "reference" && q.Implementation[0] == "empty" {
			return eligibleFuzz(q, true)
		}
		return ordinaryFuzzOutcome(q, n)
	})
	assertGateFails(t, report, "reference-full")
	assertGateFails(t, report, "mutant-coverage")
	if report.Mutants.Detected != 0 {
		t.Fatal("incorrect reference tests received private mutant credit")
	}
	for _, q := range runner.Calls {
		if q.Tests[0] == "reference" && (q.Implementation[0] == "one" || q.Implementation[0] == "two") {
			t.Fatal("reference mutants ran before public correctness was established")
		}
	}
}

func TestFuzzRepetitionRejectsZeroOrSingleRunAndMissingVet(t *testing.T) {
	runner := &scriptedFuzzRunner{fakeExecutionRunner: newFakeExecutionRunner(), Run: ordinaryFuzzOutcome}
	request := fuzzRequest{variantRequest: variantRequest{Implementation: []string{"canonical"}, Tests: []string{"reference"}}}
	for _, runs := range []int{0, 1, 4} {
		set := repeatFuzz(runner, request, runs, fmt.Sprint(runs))
		if len(set.Problems) == 0 || set.AllDetected || set.AllSurvived {
			t.Fatalf("invalid repetition count certified: %+v", set)
		}
	}
	if len(runner.Calls) != 0 {
		t.Fatal("invalid repetition count executed checks")
	}
	runner.Run = func(q fuzzRequest, _ int) fuzznative.Result { r := eligibleFuzz(q, false); r.Vetted = false; return r }
	request.Vet = true
	if set := repeatFuzz(runner, request, 2, "correct"); len(set.Problems) == 0 || set.AllSurvived {
		t.Fatalf("missing declared vet passed: %+v", set)
	}
}
