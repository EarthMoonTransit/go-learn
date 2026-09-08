package tasking

import (
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestTaskV2SyntheticProfiles(t *testing.T) {
	root := repositoryRoot(t)
	fixtures := []struct {
		profile string
		id      string
	}{
		{"fixed", "synthetic-fixed-v2-001"},
		{"mutation", "synthetic-mutation-v2-001"},
		{"workspace", "synthetic-workspace-v2-001"},
	}
	for _, fixture := range fixtures {
		t.Run(fixture.profile, func(t *testing.T) {
			runner := newFakeExecutionRunner()
			session, err := newCheckSessionWithRunner(root, runner)
			if err != nil {
				t.Fatal(err)
			}
			defer session.Close()
			path := filepath.Join(root, "internal", "testdata", "tasking", "v2", fixture.profile, fixture.id)
			report, err := session.checkLocated(LocatedTask{ID: fixture.id, State: StateDraft, Path: path})
			if err != nil {
				t.Fatal(err)
			}
			if !report.Passed {
				var details []string
				for _, gate := range report.Gates {
					if gate.Status == "fail" {
						details = append(details, gate.ID+": "+strings.Join(gate.Details, "; "))
					}
				}
				t.Fatalf("v2 fixture failed: %s", strings.Join(details, " | "))
			}
		})
	}
}

func TestTaskV2RejectsUnconditionalLearnerFailure(t *testing.T) {
	repository := repositoryRoot(t)
	path := filepath.Join(repository, "internal", "testdata", "tasking", "v2", "adversarial", "unconditional-failure", "submission_test.go")
	problems := inspectLearnerTestSource(path)
	if !containsText(problems, "unconditional") {
		t.Fatalf("expected unconditional failure rejection, got %v", problems)
	}
}

func TestTaskV2RejectsTestMainExit(t *testing.T) {
	root := repositoryRoot(t)
	path := filepath.Join(root, "internal", "testdata", "tasking", "v2", "adversarial", "testmain-exit", "submission_test.go")
	if problems := inspectLearnerTestSource(path); !containsText(problems, "TestMain") {
		t.Fatalf("expected TestMain/os.Exit rejection, got %v", problems)
	}
}

func TestWorkspaceOverlayContractRejectsGraderModification(t *testing.T) {
	root := repositoryRoot(t)
	var fixture struct {
		Editable  []string `yaml:"editable_paths"`
		Submitted []string `yaml:"submitted_paths"`
	}
	if err := decodeStrictFile(filepath.Join(root, "internal", "testdata", "tasking", "v2", "adversarial", "workspace-overlays", "grader-modification.yaml"), &fixture); err != nil {
		t.Fatal(err)
	}
	var actual []mappedFile
	for _, path := range fixture.Submitted {
		actual = append(actual, mappedFile{Relative: path})
	}
	problems := exactMappedPaths("learner overlay", fixture.Editable, actual)
	if len(problems) == 0 {
		t.Fatal("grader modification must not be accepted as an editable overlay")
	}
}

func TestTaskV2SchemaRejectsPathTraversal(t *testing.T) {
	root := repositoryRoot(t)
	path := filepath.Join(root, "internal", "testdata", "tasking", "v2", "adversarial", "invalid-editable-path", "task.yaml")
	var task TaskV2
	if err := decodeAndValidate(root, path, "task-v2.schema.json", &task); err == nil {
		t.Fatal("v2 schema accepted an editable path traversal")
	}
}

func TestStaleCollectionFixturesAreRejected(t *testing.T) {
	root := fixtureRepository(t)
	for _, test := range []struct{ name, fragment string }{{"stale-v1.yaml", "task-check/v3"}, {"stale-v2.yaml", "task-check/v3"}} {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "collections", test.name)
			var collection CheckCollection
			if err := decodeStrictFile(path, &collection); err != nil {
				t.Fatal(err)
			}
			err := ValidateCheckCollection(root, collection, []State{})
			if err == nil || !strings.Contains(err.Error(), test.fragment) {
				t.Fatalf("expected stale collection error containing %q, got %v", test.fragment, err)
			}
		})
	}
}

func TestProvenanceV2MultiSourceRoles(t *testing.T) {
	root := repositoryRoot(t)
	taskPath := filepath.Join(root, "internal", "testdata", "tasking", "v2", "fixed", "synthetic-fixed-v2-001")
	for _, test := range []struct {
		name      string
		fixture   string
		wantError string
	}{
		{"combined", "combined.yaml", ""},
		{"forbidden-reference-copy", "forbidden-reference-copy.yaml", "reference-only"},
	} {
		t.Run(test.name, func(t *testing.T) {
			var provenance ProvenanceV2
			path := filepath.Join(root, "internal", "testdata", "tasking", "provenance-v2", test.fixture)
			if err := decodeAndValidate(root, path, "task-provenance-v2.schema.json", &provenance); err != nil {
				t.Fatal(err)
			}
			data := packageData{
				Located: LocatedTask{ID: "synthetic-fixed-v2-001", State: StateDraft, Path: taskPath},
				Task:    Task{ID: "synthetic-fixed-v2-001"}, ProvenanceVersion: 2, ProvenanceV2: provenance,
			}
			problems := checkProvenanceV2(root, data)
			if test.wantError == "" && len(problems) != 0 {
				t.Fatalf("valid combined provenance rejected: %v", problems)
			}
			if test.wantError != "" && !containsText(problems, test.wantError) {
				t.Fatalf("expected %q rejection, got %v", test.wantError, problems)
			}
		})
	}
}

func TestInitDefaultsToV2WithValidOriginalProvenance(t *testing.T) {
	root := fixtureRepository(t)
	created, err := Init(root, InitOptions{Slug: "original-contract", PrimarySkill: "functions.signatures-results", Actor: "fixture-author", Date: "2026-09-04"})
	if err != nil {
		t.Fatal(err)
	}
	data, problems := loadPackage(root, created)
	if len(problems) != 0 {
		t.Fatalf("generated v2 metadata is invalid: %v", problems)
	}
	if data.ContractVersion != 2 || data.ProvenanceVersion != 2 || data.ProvenanceV2.Kind != "original" {
		t.Fatalf("task-init did not default to v2 original provenance: %+v", data)
	}
	if problems := checkProvenanceV2(root, data); len(problems) != 0 {
		t.Fatalf("original provenance scaffold is invalid: %v", problems)
	}
}

func TestCheckCollectionRejectsV1ContractAfterV2Bump(t *testing.T) {
	root := fixtureRepository(t)
	collection, err := fakeSession(t, root).CheckAll([]State{StateDraft})
	if err != nil {
		t.Fatal(err)
	}
	collection.ContractVersion = "task-check/v1"
	if err := ValidateCheckCollection(root, collection, []State{StateDraft}); err == nil || !strings.Contains(err.Error(), "task-check/v3") {
		t.Fatalf("stale v1 collection was not rejected by v2 contract: %v", err)
	}
}

func TestTaskV2PromotionAndAggregateCompatibility(t *testing.T) {
	root := fixtureRepository(t)
	id := "synthetic-fixed-v2-001"
	source := filepath.Join(repositoryRoot(t), "internal", "testdata", "tasking", "v2", "fixed", id)
	destination := filepath.Join(root, "tasks", "drafts", id)
	if err := isolation.CopyTree(source, destination); err != nil {
		t.Fatal(err)
	}
	if _, err := promoteForTest(t, root, PromoteOptions{ID: id, To: StateInReview, Actor: "reviewer", Reason: "review v2 contract", Date: "2026-09-04"}); err != nil {
		t.Fatal(err)
	}
	if _, err := promoteForTest(t, root, PromoteOptions{ID: id, To: StateReady, Actor: "reviewer", Reason: "approve v2 contract", Date: "2026-09-04"}); err != nil {
		t.Fatal(err)
	}
	manifest := filepath.Join(root, "tasks", "ready", id, "task.yaml")
	version, err := documentSchemaVersion(manifest)
	if err != nil || version != 2 {
		t.Fatalf("promotion did not preserve v2 manifest: version=%d err=%v", version, err)
	}
	report, err := buildAggregateForTest(t, root)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, task := range report.Tasks {
		if task.ID == id {
			found = task.TaskType == "implement"
		}
	}
	if !found {
		t.Fatal("aggregate report did not normalize the promoted v2 task")
	}
}

func containsText(values []string, fragment string) bool {
	for _, value := range values {
		if strings.Contains(value, fragment) {
			return true
		}
	}
	return false
}
