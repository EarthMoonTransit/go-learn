//go:build taskintegration

package fuzznative

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"go-practice-bank/internal/fuzznative/protocol"
	"go-practice-bank/internal/isolation"
)

func integrationEngine(t *testing.T) *Engine {
	t.Helper()
	environment, err := isolation.NewEnvironment("native-fuzz-fixture")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := environment.Close(); err != nil {
			t.Error(err)
		}
	})
	return NewEngine(environment)
}

func TestTrustedActualRuntimeQueryDoesNotBuildOrRunAuthoredCode(t *testing.T) {
	engine := integrationEngine(t)
	first, err := engine.ActualRuntime(60 * time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if err := first.Validate(); err != nil {
		t.Fatal(err)
	}
	before := engine.Metrics()
	second, err := engine.ActualRuntime(60 * time.Second)
	if err != nil || second != first {
		t.Fatalf("operation runtime changed: %+v %+v %v", first, second, err)
	}
	if after := engine.Metrics(); after != before {
		t.Fatalf("cached runtime spawned another query: %+v %+v", before, after)
	}
	if before.Workspaces != 1 || before.GoRuns != 1 || before.HelperRuns != 0 || before.NativeRuns != 0 || engine.helper != "" {
		t.Fatalf("metadata query built or ran authored code: %+v helper=%q", before, engine.helper)
	}
}

func nativeRequest(variant string, sources map[string]string) Request {
	return Request{Identity: Identity{ModulePath: "task.local/integration-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}, GoVersion: "1.27", Implementation: map[string]string{"hex.go": subject(variant)}, Sources: sources, Timeout: 60 * time.Second}
}

func assertNativeResult(t *testing.T, result Result, valid, detected bool) {
	t.Helper()
	if result.Valid != valid || result.Detected != detected {
		t.Fatalf("native result valid=%v detected=%v, want %v/%v: phase=%s diagnostic=%s", result.Valid, result.Detected, valid, detected, result.Phase, result.Diagnostic)
	}
	if valid && len(result.Canonical) != 64 {
		t.Fatalf("eligible result lacks canonical observation: %+v", result)
	}
	if !valid && (result.Canonical != "" || result.Diagnostic == "") {
		t.Fatalf("invalid result lacks diagnostic or retains observation: %+v", result)
	}
}

func TestNativeKernelCorrectnessAndWeakeningControls(t *testing.T) {
	engine := integrationEngine(t)
	for _, variant := range append(append([]string{}, correct...), defects...) {
		t.Run("strong/"+variant, func(t *testing.T) {
			request := nativeRequest(variant, source(strong, -1))
			request.Vet = strings.HasPrefix(variant, "correct-")
			previous := ""
			for repeat := 0; repeat < 2; repeat++ {
				result := engine.Verify(request)
				assertNativeResult(t, result, true, !strings.HasPrefix(variant, "correct-"))
				if request.Vet && !result.Vetted {
					t.Fatal("declared vet did not run")
				}
				if repeat > 0 && result.Canonical != previous {
					t.Fatal("stable native fixture changed semantic observation")
				}
				previous = result.Canonical
			}
		})
	}
	roundtrip := `decoded,ok:=subject.Decode(subject.Encode(input));if !ok||!bytes.Equal(decoded,input){t.Error("roundtrip mismatch")}`
	for _, variant := range []string{"paired-uppercase", "paired-xor"} {
		t.Run("weak-roundtrip/"+variant, func(t *testing.T) {
			assertNativeResult(t, engine.Verify(nativeRequest(variant, source(roundtrip, -1))), true, false)
		})
	}
	noStatus := beforeDecode + `decoded,_:=subject.Decode(encoded);if !bytes.Equal(decoded,input){t.Error("decoded value mismatch")}`
	t.Run("weak-status-omission", func(t *testing.T) {
		assertNativeResult(t, engine.Verify(nativeRequest("decode-false-status", source(noStatus, -1))), true, false)
	})
	for i, kind := range []string{"drop-zero", "drop-high", "reverse", "truncate"} {
		for _, side := range []string{"encode-", "decode-"} {
			t.Run("removed-critical-seed/"+side+kind, func(t *testing.T) {
				assertNativeResult(t, engine.Verify(nativeRequest(side+kind, source(strong, i+2))), true, false)
			})
		}
	}
	t.Run("removed-letter-seed", func(t *testing.T) {
		assertNativeResult(t, engine.Verify(nativeRequest("paired-uppercase", source(strong, 6))), true, false)
	})
	for _, method := range []string{"Error", "Errorf", "Fatal", "Fatalf"} {
		t.Run("assertion/"+method, func(t *testing.T) {
			body := strings.ReplaceAll(strong, `t.Error("decoded payload/status mismatch")`, `t.`+method+`("line one\nline two")`)
			assertNativeResult(t, engine.Verify(nativeRequest("decode-drop-zero", source(body, -1))), true, true)
		})
	}
	for _, id := range []string{"a-ab-001", strings.Repeat("a", 76) + "-001"} {
		t.Run("module-boundary/"+id, func(t *testing.T) {
			request := nativeRequest("correct-copy", source(strong, -1))
			request.Identity.ModulePath = "task.local/" + id
			for name, s := range request.Sources {
				request.Sources[name] = strings.ReplaceAll(s, "task.local/integration-hex-001", request.Identity.ModulePath)
			}
			assertNativeResult(t, engine.Verify(request), true, false)
		})
	}
}

func TestNativeKernelPreservesReviewedAcceptedForms(t *testing.T) {
	engine := integrationEngine(t)
	data, err := os.ReadFile("frame/testdata/finite-forms.json")
	if err != nil {
		t.Fatal(err)
	}
	var forms []struct {
		Name     string
		Accepted bool
		Sources  map[string]string
	}
	if err := json.Unmarshal(data, &forms); err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, form := range forms {
		if !form.Accepted {
			continue
		}
		count++
		t.Run(form.Name, func(t *testing.T) {
			for name, s := range form.Sources {
				form.Sources[name] = strings.ReplaceAll(s, "task.local/typed-hex-001", "task.local/integration-hex-001")
			}
			request := nativeRequest("correct-loop", form.Sources)
			assertNativeResult(t, engine.Verify(request), true, false)
		})
	}
	if count != 28 {
		t.Fatalf("accepted-form matrix incomplete: %d", count)
	}
}

func TestNativeKernelNeverCreditsLateInfrastructureFailures(t *testing.T) {
	engine := integrationEngine(t)
	for _, variant := range []string{"raw-panic", "raw-goexit", "raw-exit1", "raw-output-over", "raw-io-after-error", "raw-timeout"} {
		t.Run(variant, func(t *testing.T) {
			request := nativeRequest(variant, source(strong, -1))
			if variant == "raw-timeout" {
				request.Timeout = 5 * time.Second
			}
			result := engine.Verify(request)
			assertNativeResult(t, result, false, false)
			if result.Phase != "native" {
				t.Fatalf("fault fixture did not reach native execution: %+v", result)
			}
		})
	}
	t.Run("late-call-quota", func(t *testing.T) {
		body := strong + `;if !ok{` + strings.Repeat(`subject.Decode(encoded);`, 7) + `}`
		assertNativeResult(t, engine.Verify(nativeRequest("decode-false-status", source(body, -1))), false, false)
	})
	t.Run("unused-decode-is-not-participation", func(t *testing.T) {
		tests := source(`encoded:=subject.Encode(input);_ = encoded`, -1)
		tests["fuzz_test.go"] += `func unusedDecode(wire []byte){subject.Decode(wire)}`
		result := engine.Verify(nativeRequest("correct-copy", tests))
		assertNativeResult(t, result, false, false)
		if !strings.Contains(result.Diagnostic, "participation") {
			t.Fatalf("missing useful runtime participation feedback: %+v", result)
		}
	})
	t.Run("static-rejection-before-native", func(t *testing.T) {
		before := engine.Metrics()
		result := engine.Verify(nativeRequest("correct-copy", source(`encoded:=subject.Encode(input);_ = encoded`, -1)))
		assertNativeResult(t, result, false, false)
		if result.Phase != "admission" || !strings.Contains(result.Diagnostic, "frame.participation") || engine.Metrics().NativeRuns != before.NativeRuns {
			t.Fatalf("source rejection bypassed or opaque: %+v", result)
		}
	})
}

func TestNativeKernelMaximumAndRace(t *testing.T) {
	engine := integrationEngine(t)
	for _, test := range []struct {
		name, variant   string
		assertions      int
		valid, detected bool
	}{
		{"maximum-correct", "correct-copy", 64, true, false},
		{"maximum-64-assertions", "decode-xor-boundary", 64, true, true},
		{"maximum-65-assertions", "decode-xor-boundary", 65, false, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			result := engine.Verify(nativeRequest(test.variant, maximum(test.assertions)))
			assertNativeResult(t, result, test.valid, test.detected)
			if test.valid && (result.Callbacks != 64 || result.Calls != 512) {
				t.Fatalf("maximum fixture incomplete: %+v", result)
			}
			if test.detected && result.Assertions != 4096 {
				t.Fatalf("maximum assertion trace incomplete: %+v", result)
			}
		})
	}
	for _, variant := range []string{"correct-copy", "decode-drop-zero", "raw-race"} {
		t.Run("race/"+variant, func(t *testing.T) {
			request := nativeRequest(variant, source(strong, -1))
			request.Race = true
			result := engine.Verify(request)
			if result.RaceUnsupported {
				t.Skip(result.Diagnostic)
			}
			assertNativeResult(t, result, variant != "raw-race", variant == "decode-drop-zero")
			if variant == "raw-race" && result.Phase != "native" {
				t.Fatalf("race fixture never executed: %+v", result)
			}
		})
	}
}

// Builds and executes only the trusted frontend. No learner callback is run.
func TestTrustedHelperStandaloneBindings(t *testing.T) {
	engine := integrationEngine(t)
	if err := engine.prepare("1.27", 60*time.Second); err != nil {
		t.Fatal(err)
	}
	identity := Identity{ModulePath: "task.local/helper-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}
	base := protocol.Request{Version: protocol.Version, Identity: identity, Runtime: engine.runtime, KernelDigest: SourceDigest(), Sources: map[string]string{"fuzz_test.go": "package subject_test"}, Exports: map[string]protocol.Export{}}
	run := func(request protocol.Request, args []string) (Process, []byte, []byte) {
		t.Helper()
		raw, err := json.Marshal(request)
		if err != nil {
			t.Fatal(err)
		}
		raw = append(raw, '\n')
		process, output, err := runProcess(engine.environment, filepath.Dir(engine.helper), engine.helper, args, raw, isolation.ExecutionModeNormal, 60*time.Second, protocol.MaxResponse, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		return process, output, raw
	}
	p, output, raw := run(base, nil)
	if failure := processFailure(p, output, nil); failure != "" {
		t.Fatal(failure)
	}
	response, err := protocol.ValidateResponse(output, raw, base)
	if err != nil || response.Accepted || !strings.Contains(response.Diagnostic, "frame.import") {
		t.Fatalf("expected useful source rejection bound to actual helper: %+v err=%v", response, err)
	}
	for _, test := range []struct {
		name   string
		mutate func(*protocol.Request)
		args   []string
	}{
		{"wrong-runtime", func(r *protocol.Request) { r.Runtime.Version = "go1.27.999" }, nil},
		{"wrong-target-platform", func(r *protocol.Request) { r.Runtime.GOOS = "unsupported" }, nil},
		{"wrong-kernel", func(r *protocol.Request) { r.KernelDigest = strings.Repeat("0", 64) }, nil},
		{"wrong-protocol", func(r *protocol.Request) { r.Version-- }, nil},
		{"invalid-module", func(r *protocol.Request) { r.Identity.ModulePath = "task.local/../escape" }, nil},
		{"helper-arguments", func(*protocol.Request) {}, []string{"--accept"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			request := base
			test.mutate(&request)
			p, output, _ := run(request, test.args)
			if p.Exit != 97 || p.TimedOut || p.Truncated || len(output) == 0 {
				t.Fatalf("helper mismatch did not fail closed: %+v %.300s", p, output)
			}
		})
	}
	if engine.metrics.NativeRuns != 0 {
		t.Fatal("binding controls executed authored code")
	}
}
