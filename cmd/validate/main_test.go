package main

import (
	"bytes"
	"strings"
	"testing"

	"go-practice-bank/internal/validation"
)

func TestCoverageUsageErrorsBeforeRepositoryAccess(t *testing.T) {
	for _, args := range [][]string{
		{"-coverage", "-foundation-coverage"},
		{"-phase1", "-coverage"},
		{"-phase1", "-foundation-coverage"},
		{"-foundation-coverage"},
		{"-checks", "missing", "-run-commands"},
	} {
		var out, diagnostic bytes.Buffer
		if code := run(append([]string{"-root", t.TempDir()}, args...), &out, &diagnostic); code != 2 || out.Len() != 0 || diagnostic.Len() == 0 {
			t.Fatalf("%v: exit %d, stdout %q, stderr %q", args, code, out.String(), diagnostic.String())
		}
	}
}

func TestNumericSummaryCannotPublishFailedOrPartialEvidence(t *testing.T) {
	var out bytes.Buffer
	printSummary(&out, validation.Result{Curriculum: validation.Curriculum{Version: 4}, SelectedScope: "full", NumericUnavailable: "task-check: failed package reports: fixture-001"})
	if !strings.Contains(out.String(), "numeric evaluation unavailable: task-check: failed package reports: fixture-001") || strings.Contains(out.String(), "full-bank numeric coverage: pass") || !strings.Contains(out.String(), "not established by numeric counts") {
		t.Fatalf("failed evidence summary: %s", out.String())
	}
	out.Reset()
	printSummary(&out, validation.Result{Curriculum: validation.Curriculum{Version: 4}, SelectedScope: "foundation", NumericCoverage: &validation.NumericCoverage{Foundation: validation.NumericScopeCoverage{MinimaMet: true}, Specialist: validation.NumericScopeCoverage{PrimaryDeficit: 8, TotalDeficit: 16}, ReadyTarget: validation.NumericReadyTarget{Actual: 410, Min: 350, Max: 500, Met: true}}})
	for _, want := range []string{"foundation minima: pass", "specialist minima: fail; deficits 8 primary / 16 total", "full-bank numeric coverage: fail; ready target 410/350..500 (pass)"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("missing %q in %s", want, out.String())
		}
	}
}
