package tasking

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

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
	content := "package subject\n\nfunc Identity(value int) int {\n\treturn value\n}\n"
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

func TestPromotionRollsBackMetadataWhenSecondWriteFails(t *testing.T) {
	root := fixtureRepository(t)
	taskPath := taskFile(root, "task.yaml")
	lifecyclePath := taskFile(root, "lifecycle.yaml")
	originalTask, _ := os.ReadFile(taskPath)
	originalLifecycle, _ := os.ReadFile(lifecyclePath)
	promotionTestHooks.beforeWrite = func(name string) error {
		if name == "lifecycle.yaml" {
			return errors.New("synthetic second write failure")
		}
		return nil
	}
	t.Cleanup(func() {
		promotionTestHooks = struct {
			afterLock   func()
			beforeWrite func(string) error
		}{}
	})
	if _, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "test rollback", Date: "2026-09-04"}); err == nil {
		t.Fatal("injected failure was ignored")
	}
	afterTask, _ := os.ReadFile(taskPath)
	afterLifecycle, _ := os.ReadFile(lifecyclePath)
	if !bytes.Equal(originalTask, afterTask) || !bytes.Equal(originalLifecycle, afterLifecycle) {
		t.Fatal("ordinary promotion failure did not restore exact metadata bytes")
	}
}

func TestPromotionLockRejectsConcurrentWriter(t *testing.T) {
	root := fixtureRepository(t)
	locked := make(chan struct{})
	release := make(chan struct{})
	var once sync.Once
	promotionTestHooks.afterLock = func() {
		once.Do(func() {
			close(locked)
			<-release
		})
	}
	t.Cleanup(func() {
		promotionTestHooks = struct {
			afterLock   func()
			beforeWrite func(string) error
		}{}
	})
	done := make(chan error, 1)
	go func() {
		_, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "first", Reason: "hold lock", Date: "2026-09-04"})
		done <- err
	}()
	select {
	case <-locked:
	case <-time.After(5 * time.Second):
		t.Fatal("first promotion did not acquire lock")
	}
	_, secondErr := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "second", Reason: "compete", Date: "2026-09-04"})
	if secondErr == nil || !strings.Contains(secondErr.Error(), "locked") {
		t.Fatalf("concurrent promotion was not rejected: %v", secondErr)
	}
	close(release)
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

func TestPromotionRejectsPredatedTransition(t *testing.T) {
	root := fixtureRepository(t)
	_, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "old date", Date: "2026-09-03"})
	if err == nil || !strings.Contains(err.Error(), "predates") {
		t.Fatalf("predated transition was accepted: %v", err)
	}
}

func TestReadyPromotionValidatesFutureState(t *testing.T) {
	root := fixtureRepository(t)
	if _, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "review", Date: "2026-09-04"}); err != nil {
		t.Fatal(err)
	}
	if _, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateReady, Actor: "reviewer", Reason: "approve", Date: "2026-09-04"}); err != nil {
		t.Fatal(err)
	}
	located, err := FindTask(root, syntheticID)
	if err != nil || located.State != StateReady {
		t.Fatalf("ready promotion failed: %#v %v", located, err)
	}
}

func TestCheckerRejectsAuthoredReferenceAndMutantBindings(t *testing.T) {
	root := fixtureRepository(t)
	referenceFixture := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "adversarial", "noop-reference-binding", "solution_test.go")
	mutantFixture := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "adversarial", "unconditional-mutant-failure", "mutant_test.go")
	if err := isolation.CopyFile(referenceFixture, taskFile(root, "solution/reference/solution_test.go")); err != nil {
		t.Fatal(err)
	}
	if err := isolation.CopyFile(mutantFixture, taskFile(root, "solution/mutants/constant-zero/mutant_test.go")); err != nil {
		t.Fatal(err)
	}
	assertGateFails(t, checkWithoutError(t, root), "package-contract")
}

func TestNoOpPublicAndGraderCannotSatisfyDynamicGates(t *testing.T) {
	t.Run("public", func(t *testing.T) {
		root := fixtureRepository(t)
		replaceDirectoryWithFixture(t, root, "tests/public", "noop-public")
		assertGateFails(t, checkWithoutError(t, root), "starter-public")
	})
	t.Run("grader", func(t *testing.T) {
		root := fixtureRepository(t)
		replaceDirectoryWithFixture(t, root, "tests/grader", "noop-grader")
		realistic := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "adversarial", "surviving-realistic-mutant", "identity.go")
		if err := os.Remove(taskFile(root, "solution/mutants/constant-zero/identity.go")); err != nil {
			t.Fatal(err)
		}
		if err := isolation.CopyFile(realistic, taskFile(root, "solution/mutants/constant-zero/identity.go")); err != nil {
			t.Fatal(err)
		}
		report := checkWithoutError(t, root)
		assertGateFails(t, report, "mutant-coverage")
	})
}

func TestAdversarialCorrectReferenceRunsSharedSuites(t *testing.T) {
	root := fixtureRepository(t)
	source := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "adversarial", "correct-reference", "identity.go")
	if err := os.Remove(taskFile(root, "solution/reference/identity.go")); err != nil {
		t.Fatal(err)
	}
	if err := isolation.CopyFile(source, taskFile(root, "solution/reference/identity.go")); err != nil {
		t.Fatal(err)
	}
	report := checkWithoutError(t, root)
	if report.ReferenceStatus != "pass" || !report.Passed {
		t.Fatalf("correct reference did not run canonical suites: %#v", report)
	}
}

func TestOldAuthoredBindingBypassWasReproducible(t *testing.T) {
	root := repositoryRoot(t)
	environment, err := isolation.NewEnvironment("old-binding-proof")
	if err != nil {
		t.Fatal(err)
	}
	defer environment.Close()
	for _, testCase := range []struct {
		name     string
		fixture  string
		wantPass bool
	}{
		{"noop-reference", "noop-reference-binding", true},
		{"forced-mutant-failure", "unconditional-mutant-failure", false},
	} {
		workspace, _ := environment.NewWorkspace(testCase.name)
		if err := os.WriteFile(filepath.Join(workspace, "go.mod"), []byte("module task.local/old-binding\n\ngo 1.27\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(workspace, "identity.go"), []byte("package subject\nfunc Identity(v int) int { return v + 1 }\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		source := filepath.Join(root, "internal", "testdata", "tasking", "adversarial", testCase.fixture)
		entries, _ := os.ReadDir(source)
		if err := isolation.CopyFile(filepath.Join(source, entries[0].Name()), filepath.Join(workspace, entries[0].Name())); err != nil {
			t.Fatal(err)
		}
		_, runErr := environment.RunGo(workspace, 30*time.Second, "test", "./...")
		if (runErr == nil) != testCase.wantPass {
			t.Fatalf("%s pass=%t, want %t", testCase.name, runErr == nil, testCase.wantPass)
		}
	}
}

func TestDiscoverTasksRejectsUnexpectedAndDuplicateEntries(t *testing.T) {
	t.Run("unexpected", func(t *testing.T) {
		root := fixtureRepository(t)
		if err := os.Mkdir(filepath.Join(root, "tasks", "drafts", "malformed"), 0o755); err != nil {
			t.Fatal(err)
		}
		if _, err := DiscoverTasks(root, nil); err == nil || !strings.Contains(err.Error(), "unexpected") {
			t.Fatalf("unexpected directory was ignored: %v", err)
		}
	})
	t.Run("duplicate outside filter", func(t *testing.T) {
		root := fixtureRepository(t)
		if err := isolation.CopyTree(taskFile(root, "."), filepath.Join(root, "tasks", "rejected", syntheticID)); err != nil {
			t.Fatal(err)
		}
		if _, err := CheckAll(root, []State{StateDraft}); err == nil || !strings.Contains(err.Error(), "duplicate") {
			t.Fatalf("filtered duplicate was ignored: %v", err)
		}
	})
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
	if len(first.Tasks) != 1 {
		t.Fatalf("aggregate task count = %d", len(first.Tasks))
	}
	item := first.Tasks[0]
	if item.TechnicalStatus != "pass" || item.ReviewStatus != "pass" || !item.PromotionEligible || item.StarterStatus != "pass" || item.ReferenceStatus != "pass" || item.Mutants.Detected != 1 {
		t.Fatalf("aggregate omits check evidence: %#v", item)
	}
}

func TestAggregateArtifactCheckDetectsStaleReportWithoutWriting(t *testing.T) {
	root := fixtureRepository(t)
	report, err := BuildAggregateReport(root)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "reports"), 0o755); err != nil {
		t.Fatal(err)
	}
	yamlContent, _ := MarshalAggregateYAML(report)
	markdownPath := filepath.Join(root, "reports", "task-quality.md")
	yamlPath := filepath.Join(root, "reports", "task-quality.yaml")
	if err := os.WriteFile(markdownPath, MarshalAggregateMarkdown(report), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(yamlPath, yamlContent, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := CheckAggregateArtifacts(root, "reports/task-quality.md", "reports/task-quality.yaml", report); err != nil {
		t.Fatal(err)
	}
	beforeYAML, _ := os.ReadFile(yamlPath)
	if err := os.WriteFile(markdownPath, []byte("stale\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := CheckAggregateArtifacts(root, "reports/task-quality.md", "reports/task-quality.yaml", report); err == nil {
		t.Fatal("stale report was accepted")
	}
	afterYAML, _ := os.ReadFile(yamlPath)
	if !bytes.Equal(beforeYAML, afterYAML) {
		t.Fatal("check mode changed another report artifact")
	}
}

func TestCheckAllFindsBrokenDraftAndInReviewPackages(t *testing.T) {
	t.Run("draft", func(t *testing.T) {
		root := fixtureRepository(t)
		if err := os.RemoveAll(taskFile(root, "tests/grader")); err != nil {
			t.Fatal(err)
		}
		collection, err := CheckAll(root, []State{StateDraft})
		if err != nil {
			t.Fatal(err)
		}
		if len(collection.Reports) != 1 || collection.Reports[0].Passed {
			t.Fatalf("broken draft disappeared or passed: %#v", collection.Reports)
		}
	})
	t.Run("in-review", func(t *testing.T) {
		root := fixtureRepository(t)
		moved, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "review", Date: "2026-09-04"})
		if err != nil {
			t.Fatal(err)
		}
		if err := os.RemoveAll(filepath.Join(moved.Path, "tests", "grader")); err != nil {
			t.Fatal(err)
		}
		collection, err := CheckAll(root, []State{StateInReview})
		if err != nil {
			t.Fatal(err)
		}
		if len(collection.Reports) != 1 || collection.Reports[0].Passed {
			t.Fatalf("broken in-review package disappeared or passed: %#v", collection.Reports)
		}
	})
}

func TestFutureStateValidationRejectsLifecycleMismatch(t *testing.T) {
	root := fixtureRepository(t)
	taskYAML, _ := os.ReadFile(taskFile(root, "task.yaml"))
	lifecycleYAML, _ := os.ReadFile(taskFile(root, "lifecycle.yaml"))
	located, err := FindTask(root, syntheticID)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateFuturePromotion(root, located, StateInReview, taskYAML, lifecycleYAML); err == nil {
		t.Fatal("mismatched future lifecycle was accepted")
	}
}

func TestPromotionDestinationCollisionPreservesSource(t *testing.T) {
	root := fixtureRepository(t)
	original, _ := os.ReadFile(taskFile(root, "task.yaml"))
	if err := isolation.CopyTree(taskFile(root, "."), filepath.Join(root, "tasks", "review", syntheticID)); err != nil {
		t.Fatal(err)
	}
	if _, err := Promote(root, PromoteOptions{ID: syntheticID, To: StateInReview, Actor: "reviewer", Reason: "collision", Date: "2026-09-04"}); err == nil {
		t.Fatal("destination collision was accepted")
	}
	after, _ := os.ReadFile(taskFile(root, "task.yaml"))
	if !bytes.Equal(original, after) {
		t.Fatal("collision changed source metadata")
	}
}

func TestLegacyCompatibilityUsesExplicitPolicy(t *testing.T) {
	root := legacyFixtureRepository(t)
	report, err := Check(root, "defined-types-001")
	if err != nil {
		t.Fatal(err)
	}
	if !report.Passed || !report.Legacy || report.PromotionEligible {
		t.Fatalf("unexpected legacy result: %#v", report)
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
	for _, directory := range []string{"config", "schemas", "curriculum", "sources", "tasks/drafts", "tasks/review", "tasks/ready", "tasks/rejected", "tasks/retired"} {
		if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(directory)), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	for _, relative := range []string{
		"schemas/task.schema.json", "schemas/task-lifecycle.schema.json", "schemas/task-review.schema.json", "schemas/task-provenance.schema.json", "schemas/task-quality-report.schema.json", "schemas/task-aggregate-report.schema.json", "schemas/toolchain-policy.schema.json",
		"config/toolchain.yaml",
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

func legacyFixtureRepository(t *testing.T) string {
	t.Helper()
	repository := repositoryRoot(t)
	root := t.TempDir()
	for _, directory := range []string{"config", "schemas", "tasks/ready"} {
		if err := os.MkdirAll(filepath.Join(root, directory), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	for _, relative := range []string{"config/toolchain.yaml", "schemas/task.schema.json", "schemas/task-quality-report.schema.json"} {
		if err := isolation.CopyFile(filepath.Join(repository, relative), filepath.Join(root, relative)); err != nil {
			t.Fatal(err)
		}
	}
	directory := filepath.Join(root, "tasks", "ready", "defined-types-001")
	if err := os.MkdirAll(directory, 0o755); err != nil {
		t.Fatal(err)
	}
	manifest := `schema_version: 1
id: defined-types-001
title: Synthetic legacy compatibility fixture
task_type: implement
skills:
  primary: functions.signatures-results
  secondary: []
prerequisites: []
new_concepts:
  - functions.signatures-results
difficulty: introductory
estimated_time_minutes:
  min: 5
  max: 10
source:
  id: synthetic
  repository: https://invalid.example/synthetic
  path: synthetic
  commit: "0000000000000000000000000000000000000000"
  license: MIT
transformation_type: adapted
validation_commands:
  - go test ./...
quality_status: ready
`
	for name, content := range map[string]string{
		"task.yaml": manifest,
		"go.mod":    "module task.local/synthetic-legacy\n\ngo 1.24\n",
		"value.go":  "package legacy\nfunc Value() int { return 1 }\n",
	} {
		if err := os.WriteFile(filepath.Join(directory, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func replaceDirectoryWithFixture(t *testing.T, root, relative, fixture string) {
	t.Helper()
	destination := taskFile(root, relative)
	if err := os.RemoveAll(destination); err != nil {
		t.Fatal(err)
	}
	source := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "adversarial", fixture)
	if err := isolation.CopyTree(source, destination); err != nil {
		t.Fatal(err)
	}
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
