// Package testrepo builds disposable metadata fixtures for repository tests.
// Its synthetic collections are not acceptance evidence and execute no task code.
package testrepo

import (
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

// Copy includes the complete bank and ordinary validation inputs, without the
// source checkout's Git data, ignored evidence, binaries or dependency tree.
func Copy(t testing.TB, source string) string {
	t.Helper()
	root := t.TempDir()
	for _, name := range []string{"adapters", "catalog", "config", "curriculum", "docs", "provenance", "reports", "schemas", "sources", "tasks", "theory", ".github"} {
		if err := isolation.CopyTree(filepath.Join(source, name), filepath.Join(root, name)); err != nil {
			t.Fatal(err)
		}
	}
	for _, name := range []string{"AGENTS.md", "CONTRIBUTING.md", "LICENSE", "README.md", "THIRD_PARTY_NOTICES", "go.mod", "go.sum"} {
		if err := isolation.CopyFile(filepath.Join(source, name), filepath.Join(root, name)); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

// Collection reconstructs synthetic reports from the accepted metadata and
// computes current input digests. It never calls CheckAll or a native runner.
func Collection(t testing.TB, root string) tasking.CheckCollection {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, "reports/task-quality.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var aggregate tasking.AggregateReport
	if err := yaml.Unmarshal(b, &aggregate); err != nil {
		t.Fatal(err)
	}
	quality := map[string]tasking.AggregateTask{}
	for _, item := range aggregate.Tasks {
		quality[item.ID] = item
	}
	located, err := tasking.DiscoverTasks(root, nil)
	if err != nil {
		t.Fatal(err)
	}
	policyDigest, err := tasking.ToolchainPolicyDigest(root)
	if err != nil {
		t.Fatal(err)
	}
	checks := tasking.CheckCollection{SchemaVersion: 1, ContractVersion: tasking.CheckContractVersion, ToolchainPolicyDigest: policyDigest, Execution: tasking.ExecutionMetrics{Sessions: 1, TaskChecks: len(located)}}
	for _, item := range located {
		inputDigest, err := tasking.TaskInputDigest(root, item)
		if err != nil {
			t.Fatal(err)
		}
		q, exists := quality[item.ID]
		if !exists || q.TechnicalStatus != "pass" {
			t.Fatalf("synthetic passing fixture requires accepted metadata for %s", item.ID)
		}
		checks.Tasks = append(checks.Tasks, tasking.CheckedTask{TaskID: item.ID, State: item.State, InputDigest: inputDigest, Report: tasking.QualityReport{
			SchemaVersion: 1, TaskID: item.ID, State: item.State, Passed: true,
			TechnicalStatus: q.TechnicalStatus, ReviewStatus: q.ReviewStatus, PromotionEligible: q.PromotionEligible,
			Legacy: q.Legacy, FailingGateIDs: append([]string{}, q.FailingGateIDs...),
			StarterStatus: q.StarterStatus, ReferenceStatus: q.ReferenceStatus, RaceStatus: q.RaceStatus,
			Gates: []tasking.Gate{}, Mutants: q.Mutants,
		}})
	}
	return checks
}

// WriteArtifacts makes the synthetic collection and both task-quality artifacts
// mutually current, so a failed report is not masked by an unrelated stale file.
func WriteArtifacts(t testing.TB, root string, checks tasking.CheckCollection) string {
	t.Helper()
	content, err := tasking.MarshalCheckCollection(checks)
	if err != nil {
		t.Fatal(err)
	}
	path := "checks-fixture.yaml"
	if err := os.WriteFile(filepath.Join(root, path), content, 0600); err != nil {
		t.Fatal(err)
	}
	report, err := tasking.BuildAggregateReportFromChecks(root, checks)
	if err != nil {
		t.Fatal(err)
	}
	if err := tasking.ValidateAggregateReport(root, report); err != nil {
		t.Fatal(err)
	}
	data, err := tasking.MarshalAggregateYAML(report)
	if err != nil {
		t.Fatal(err)
	}
	for relative, b := range map[string][]byte{"reports/task-quality.yaml": data, "reports/task-quality.md": tasking.MarshalAggregateMarkdown(report)} {
		if err := os.WriteFile(filepath.Join(root, filepath.FromSlash(relative)), b, 0600); err != nil {
			t.Fatal(err)
		}
	}
	return path
}
