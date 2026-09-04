package tasking

import (
	"errors"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestVariantFailureClassification(t *testing.T) {
	tests := []struct {
		name     string
		race     bool
		compiled bool
		output   string
		err      error
		kind     variantFailureKind
		phase    variantFailurePhase
	}{
		{name: "compile", err: errors.New("exit status 1"), kind: variantFailureCompile, phase: variantPhaseCompile},
		{name: "test", compiled: true, err: errors.New("exit status 1"), kind: variantFailureTest, phase: variantPhaseTest},
		{name: "timeout", compiled: true, err: isolation.ErrCommandTimeout, kind: variantFailureTimeout, phase: variantPhaseInfrastructure},
		{name: "race cgo", race: true, output: "go: -race requires cgo", err: errors.New("exit status 2"), kind: variantFailureRaceCGO, phase: variantPhaseCompile},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			kind, phase := classifyVariantFailure(test.race, test.compiled, []byte(test.output), test.err)
			if kind != test.kind || phase != test.phase {
				t.Fatalf("got %s/%s, want %s/%s", kind, phase, test.kind, test.phase)
			}
		})
	}
}

func TestVariantDiagnosticIsSanitizedAndBounded(t *testing.T) {
	root := "/private/tmp/go-task-check-sensitive"
	workspace := root + "/workspaces/task-123"
	line := `{"Action":"output","Package":"task.local/synthetic-identity-001/subject","Output":"` + workspace + `/subject/value.go:2: failure\n"}` + "\n"
	line += strings.Repeat("diagnostic ", maxVariantDiagnosticRunes)
	diagnostic := sanitizeVariantDiagnostic([]byte(line), errors.New("exit status 1"), root, workspace, syntheticID)
	if strings.Contains(diagnostic, root) || strings.Contains(diagnostic, workspace) || strings.Contains(diagnostic, syntheticID) {
		t.Fatalf("diagnostic leaked task or temporary paths: %q", diagnostic)
	}
	if !strings.Contains(diagnostic, "<workspace>/subject/value.go") {
		t.Fatalf("workspace was not replaced: %q", diagnostic)
	}
	if len([]rune(diagnostic)) > maxVariantDiagnosticRunes {
		t.Fatalf("diagnostic is not bounded: %d runes", len([]rune(diagnostic)))
	}
}

func TestExpectedFailureRejectsInfrastructureAndCompileFailures(t *testing.T) {
	for _, result := range []variantResult{
		{FailureKind: variantFailureCompile, FailurePhase: variantPhaseCompile, Diagnostic: "compile failed"},
		{Compiled: true, FailureKind: variantFailureTimeout, FailurePhase: variantPhaseInfrastructure, Diagnostic: "timeout"},
		{Compiled: true, FailureKind: variantFailureRaceCGO, FailurePhase: variantPhaseCompile, Diagnostic: "cgo unavailable"},
	} {
		if problems := expectFailure("starter", result); len(problems) == 0 {
			t.Fatalf("%s was accepted as an expected test failure", result.FailureKind)
		}
	}
}
