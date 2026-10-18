package validation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/testrepo"
)

func TestNumericCoverageRequiresSuccessfulCurrentEvidence(t *testing.T) {
	root := testrepo.Copy(t, repositoryRoot(t))
	checks := testrepo.Collection(t, root)
	testrepo.WriteArtifacts(t, root, checks)
	passed, err := ValidateRoot(root, Options{Checks: &checks})
	if err != nil || passed.NumericCoverage == nil || passed.NumericUnavailable != "" || len(passed.Coverage) != 145 || len(passed.SelectedSkillIDs) != 145 {
		t.Fatalf("passing evidence did not enable full arithmetic: %v / %s", err, passed.NumericUnavailable)
	}
	// This valid incomplete bank has numeric deficits, not an ordinary error.
	runs := 0
	foundation, err := validateRoot(root, Options{RunCommands: true, EnforceFoundationCoverage: true}, func(gotRoot string, states []tasking.State) (tasking.CheckCollection, error) {
		runs++
		if gotRoot != root || states != nil {
			t.Error("generated collection must cover the entire requested root")
		}
		// Synthetic metadata substitutes only the execution dependency. The
		// ordinary loader, freshness checks and eligibility path remain real.
		return checks, nil
	})
	if err == nil || !strings.Contains(err.Error(), "foundation coverage:") || foundation.NumericCoverage == nil || foundation.SelectedScope != "foundation" || len(foundation.SelectedSkillIDs) != 141 || len(foundation.Coverage) != 145 {
		t.Fatalf("foundation enforcement filtered the inventory or hid deficits: %v", err)
	}
	if runs != 1 {
		t.Fatalf("RunCommands generated %d collections, want exactly one", runs)
	}
	// An independent ordinary error outside collection inputs must prevent a
	// partial Result from carrying an authoritative numeric evaluation.
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module invalid\n\ngo 1.27\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := tasking.LoadCheckCollection(root, "checks-fixture.yaml"); err != nil {
		t.Fatalf("ordinary-error fixture unexpectedly made collection stale: %v", err)
	}
	partial, err := ValidateRoot(root, Options{Checks: &checks})
	if err == nil || !strings.Contains(err.Error(), "root module must be") || len(partial.Coverage) != 145 || partial.NumericCoverage != nil || partial.NumericUnavailable != "ordinary repository validation failed" {
		t.Fatalf("partial ordinary-error Result exposed numeric evaluation: %v / %s", err, partial.NumericUnavailable)
	}
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module go-practice-bank\n\ngo 1.27\n"), 0600); err != nil {
		t.Fatal(err)
	}
	failed := checks
	failed.Tasks = append([]tasking.CheckedTask(nil), checks.Tasks...)
	found := false
	for i := range failed.Tasks {
		if failed.Tasks[i].TaskID == "literal-bindings-001" {
			found = true
			r := &failed.Tasks[i].Report
			r.Passed, r.PromotionEligible = false, false
			r.TechnicalStatus, r.ReferenceStatus = "fail", "fail"
			r.FailingGateIDs = []string{"reference-full"}
			r.Gates = []tasking.Gate{{ID: "reference-full", Status: "fail", Details: []string{"synthetic evidence fixture"}}}
		}
	}
	if !found {
		t.Fatal("fixture lost literal-bindings-001")
	}
	path := testrepo.WriteArtifacts(t, root, failed)
	loaded, err := tasking.LoadCheckCollection(root, path)
	if err != nil {
		t.Fatal(err)
	}
	// Matching artifacts deliberately let ordinary validation succeed. This
	// assertion prevents the regression from merely exercising stale reports.
	ordinary, err := ValidateRoot(root, Options{Checks: &loaded})
	if err != nil {
		t.Fatalf("failed current evidence must reach the separate eligibility guard: %v", err)
	}
	if ordinary.NumericCoverage != nil || !strings.Contains(ordinary.NumericUnavailable, "literal-bindings-001") || len(ordinary.Coverage) != 145 {
		t.Fatalf("failed report received an authoritative result: %+v", ordinary.NumericCoverage)
	}
	blocked, err := ValidateRoot(root, Options{Checks: &loaded, EnforceFoundationCoverage: true})
	if err == nil || !strings.Contains(err.Error(), "failed package reports: literal-bindings-001") || blocked.NumericCoverage != nil {
		t.Fatalf("foundation enforcement accepted failed current evidence: %v", err)
	}
	// Without execution evidence, the old ordinary validator still succeeds;
	// raw encounters remain available, with no new authoritative pass.
	without, err := ValidateRoot(root, Options{})
	if err != nil || without.NumericCoverage != nil || len(without.Coverage) != 145 || without.NumericUnavailable == "" {
		t.Fatalf("ordinary compatibility/evidence boundary changed: %v", err)
	}

}

func TestNumericCoverageOptionsRejectNewInvalidCombinations(t *testing.T) {
	for _, tc := range []struct {
		options Options
		want    string
	}{
		{Options{EnforceCoverage: true, EnforceFoundationCoverage: true, RunCommands: true}, "mutually exclusive"},
		{Options{EnforceFoundationCoverage: true}, "requires a current -checks collection"},
	} {
		result, err := ValidateRoot(t.TempDir(), tc.options)
		if err == nil || !strings.Contains(err.Error(), tc.want) || result.NumericCoverage != nil {
			t.Fatalf("invalid selection was accepted: %v", err)
		}
	}
}
