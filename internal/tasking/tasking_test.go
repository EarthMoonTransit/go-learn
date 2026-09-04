package tasking

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

const syntheticID = "synthetic-identity-001"

func TestCheckValidDraft(t *testing.T) {
	root := fixtureRepository(t)
	report, err := Check(root, syntheticID)
	if err != nil {
		t.Fatal(err)
	}
	if !report.Passed {
		t.Fatalf("valid fixture failed: %#v", report.Gates)
	}
	if report.Mutants.Total != 1 || report.Mutants.Detected != 1 || len(report.Mutants.Survived) != 0 {
		t.Fatalf("unexpected mutant result: %#v", report.Mutants)
	}
}

func TestCheckRejectsInvalidMetadata(t *testing.T) {
	root := fixtureRepository(t)
	replaceInFile(t, taskFile(root, "task.yaml"), "functions.signatures-results", "functions.not-a-skill")
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "taxonomy-and-lifecycle")
}

func TestCheckRejectsMissingPublicTests(t *testing.T) {
	root := fixtureRepository(t)
	if err := os.RemoveAll(taskFile(root, "tests/public")); err != nil {
		t.Fatal(err)
	}
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "package-contract")
}

func TestCheckRejectsMissingGrader(t *testing.T) {
	root := fixtureRepository(t)
	if err := os.RemoveAll(taskFile(root, "tests/grader")); err != nil {
		t.Fatal(err)
	}
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "package-contract")
}

func TestCheckRejectsStarterThatPasses(t *testing.T) {
	root := fixtureRepository(t)
	replaceInFile(t, taskFile(root, "starter/identity.go"), "return value + 1", "return value")
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "starter-public")
	assertGateFails(t, report, "starter-full")
}

func TestCheckRejectsFailingReference(t *testing.T) {
	root := fixtureRepository(t)
	replaceInFile(t, taskFile(root, "solution/reference/identity.go"), "return value", "return value + 1")
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "reference-full")
}

func TestCheckRejectsSurvivingMutant(t *testing.T) {
	root := fixtureRepository(t)
	path := taskFile(root, "solution/mutants/constant-zero/identity.go")
	content := "package identity\n\nfunc Identity(value int) int {\n\treturn value\n}\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "mutant-coverage")
	if len(report.Mutants.Survived) != 1 || report.Mutants.Survived[0] != "constant-zero" {
		t.Fatalf("surviving mutant was not reported: %#v", report.Mutants)
	}
}

func TestCheckRejectsPathTraversal(t *testing.T) {
	root := fixtureRepository(t)
	if _, err := Check(root, "../synthetic-identity-001"); err == nil {
		t.Fatal("path traversal id was accepted")
	}
}

func TestCheckRejectsSymlinkEscape(t *testing.T) {
	root := fixtureRepository(t)
	outside := filepath.Join(t.TempDir(), "outside.md")
	if err := os.WriteFile(outside, []byte("outside"), 0o644); err != nil {
		t.Fatal(err)
	}
	path := taskFile(root, "ATTRIBUTION.md")
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, path); err != nil {
		t.Fatal(err)
	}
	report := checkWithoutError(t, root)
	assertGateFails(t, report, "package-metadata")
}

func TestPromotionRejectsFailingGates(t *testing.T) {
	root := fixtureRepository(t)
	if err := os.RemoveAll(taskFile(root, "tests/grader")); err != nil {
		t.Fatal(err)
	}
	_, err := Promote(root, PromoteOptions{
		ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "start review", Date: "2026-09-04",
	})
	if err == nil {
		t.Fatal("promotion with a missing grader was accepted")
	}
	if _, statErr := os.Stat(taskFile(root, "task.yaml")); statErr != nil {
		t.Fatalf("failed promotion moved or damaged source: %v", statErr)
	}
}

func TestPromotionRejectsDirectDraftToReady(t *testing.T) {
	root := fixtureRepository(t)
	_, err := Promote(root, PromoteOptions{
		ID: syntheticID, To: StateReady, Actor: "reviewer", Reason: "skip review", Date: "2026-09-04",
	})
	if err == nil || !strings.Contains(err.Error(), "not allowed") {
		t.Fatalf("direct draft to ready must be rejected, got %v", err)
	}
}

func TestDuplicateTaskIDRejected(t *testing.T) {
	root := fixtureRepository(t)
	duplicate := filepath.Join(root, "tasks", "rejected", syntheticID)
	if err := isolation.CopyTree(filepath.Join(root, "tasks", "drafts", syntheticID), duplicate); err != nil {
		t.Fatal(err)
	}
	if _, err := FindTask(root, syntheticID); err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("expected duplicate id error, got %v", err)
	}
}

func TestPromotionUsesAtomicDirectoryMove(t *testing.T) {
	root := fixtureRepository(t)
	moved, err := Promote(root, PromoteOptions{
		ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "all authoring checks passed", Date: "2026-09-04",
	})
	if err != nil {
		t.Fatal(err)
	}
	if moved.State != StateInReview {
		t.Fatalf("state = %s, want in-review", moved.State)
	}
	if _, err := os.Stat(filepath.Join(root, "tasks", "drafts", syntheticID)); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("draft source still exists after move: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "tasks", "review", syntheticID, "task.yaml")); err != nil {
		t.Fatalf("review destination is incomplete: %v", err)
	}
	report, err := Check(root, syntheticID)
	if err != nil || !report.Passed {
		t.Fatalf("moved package is invalid: report=%#v err=%v", report, err)
	}
	entries, err := os.ReadDir(filepath.Join(root, "tasks", "review"))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != syntheticID {
		t.Fatalf("unexpected atomic staging residue: %#v", entries)
	}
}

func TestAggregateReportIsDeterministic(t *testing.T) {
	root := fixtureRepository(t)
	first, err := BuildAggregateReport(root)
	if err != nil {
		t.Fatal(err)
	}
	second, err := BuildAggregateReport(root)
	if err != nil {
		t.Fatal(err)
	}
	if err := ValidateAggregateReport(root, first); err != nil {
		t.Fatalf("aggregate report schema: %v", err)
	}
	firstYAML, _ := MarshalAggregateYAML(first)
	secondYAML, _ := MarshalAggregateYAML(second)
	if !bytes.Equal(firstYAML, secondYAML) {
		t.Fatal("machine-readable reports differ")
	}
	if !bytes.Equal(MarshalAggregateMarkdown(first), MarshalAggregateMarkdown(second)) {
		t.Fatal("Markdown reports differ")
	}
}

func TestInitCreatesScaffoldWithoutEducationalContent(t *testing.T) {
	root := fixtureRepository(t)
	created, err := Init(root, InitOptions{
		Slug: "tooling-draft", PrimarySkill: "functions.signatures-results", SourceID: "exercism-go",
		SourcePath: "exercises/practice/hello-world", Actor: "author", Date: "2026-09-04",
	})
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != "tooling-draft-001" {
		t.Fatalf("id = %q, want tooling-draft-001", created.ID)
	}
	for _, directory := range []string{"starter", "tests/public", "tests/grader", "solution/reference", "solution/mutants"} {
		entries, err := os.ReadDir(filepath.Join(created.Path, filepath.FromSlash(directory)))
		if err != nil || len(entries) != 0 {
			t.Fatalf("%s must be an empty scaffold: entries=%v err=%v", directory, entries, err)
		}
	}
	content, err := os.ReadFile(filepath.Join(created.Path, "provenance.yaml"))
	if err != nil || !bytes.Contains(content, []byte("TODO: record")) {
		t.Fatalf("provenance stub missing: %v", err)
	}
}

func fixtureRepository(t *testing.T) string {
	t.Helper()
	repositoryRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	for _, directory := range []string{"schemas", "curriculum", "sources", "tasks/drafts", "tasks/review", "tasks/ready", "tasks/rejected", "tasks/retired"} {
		if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(directory)), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	for _, relative := range []string{
		"schemas/task.schema.json", "schemas/task-lifecycle.schema.json", "schemas/task-review.schema.json", "schemas/task-provenance.schema.json", "schemas/task-quality-report.schema.json", "schemas/task-aggregate-report.schema.json",
		"curriculum/skills.yaml", "sources/registry.yaml",
	} {
		if err := isolation.CopyFile(filepath.Join(repositoryRoot, filepath.FromSlash(relative)), filepath.Join(root, filepath.FromSlash(relative))); err != nil {
			t.Fatalf("copy repository contract %s: %v", relative, err)
		}
	}
	fixture := filepath.Join(repositoryRoot, "internal", "testdata", "tasking", syntheticID)
	destination := filepath.Join(root, "tasks", "drafts", syntheticID)
	if err := isolation.CopyTree(fixture, destination); err != nil {
		t.Fatal(err)
	}
	if err := isolation.CopyTree(filepath.Join(repositoryRoot, "tasks", "_template"), filepath.Join(root, "tasks", "_template")); err != nil {
		t.Fatal(err)
	}
	return root
}

func taskFile(root, relative string) string {
	return filepath.Join(root, "tasks", "drafts", syntheticID, filepath.FromSlash(relative))
}

func replaceInFile(t *testing.T, path, old, replacement string) {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	updated := strings.Replace(string(content), old, replacement, 1)
	if updated == string(content) {
		t.Fatalf("%q not found in %s", old, path)
	}
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		t.Fatal(err)
	}
}

func checkWithoutError(t *testing.T, root string) QualityReport {
	t.Helper()
	report, err := Check(root, syntheticID)
	if err != nil {
		t.Fatal(err)
	}
	return report
}

func assertGateFails(t *testing.T, report QualityReport, id string) {
	t.Helper()
	for _, gate := range report.Gates {
		if gate.ID == id {
			if gate.Status != "fail" {
				t.Fatalf("gate %s status = %s, want fail", id, gate.Status)
			}
			return
		}
	}
	t.Fatalf("gate %s not found", id)
}
