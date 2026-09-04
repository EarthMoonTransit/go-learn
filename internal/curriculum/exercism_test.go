package curriculum

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
	"gopkg.in/yaml.v3"
)

func TestExercismV3EvidenceAndMappingSemantics(t *testing.T) {
	root := curriculumRepositoryRoot(t)
	mapping, evidence, source, contracts, inventory, candidates := loadExercismFixtures(t, root)
	if err := Validate(mapping, evidence, source, 3, contracts, inventory.Concepts, candidates); err != nil {
		t.Fatalf("repository fixture must validate: %v", err)
	}

	t.Run("deprecated mapping target", func(t *testing.T) {
		changed := cloneContracts(contracts)
		id := mapping.Concepts[0].SkillIDs[0]
		changed[id] = SkillContract{Status: "deprecated", Assignable: false}
		err := Validate(mapping, evidence, source, 3, changed, inventory.Concepts, candidates)
		assertCurriculumErrorContains(t, err, "non-assignable skill")
	})

	t.Run("invalid calendar date", func(t *testing.T) {
		changed := evidence
		changed.Reviews = append([]EvidenceReview(nil), evidence.Reviews...)
		changed.Reviews[0].ReviewedAt = "2026-02-30"
		err := Validate(mapping, changed, source, 3, contracts, inventory.Concepts, candidates)
		assertCurriculumErrorContains(t, err, "lacks review metadata")
	})

	t.Run("manual review requires artifact hash", func(t *testing.T) {
		changed := evidence
		changed.Reviews = append([]EvidenceReview(nil), evidence.Reviews...)
		changed.Reviews[0].Basis = "manual-artifact-review"
		changed.Reviews[0].ArtifactPath = ""
		changed.Reviews[0].ArtifactSHA256 = ""
		err := Validate(mapping, changed, source, 3, contracts, inventory.Concepts, candidates)
		assertCurriculumErrorContains(t, err, "requires artifact_path and artifact_sha256")
	})

	t.Run("duplicate candidate source path", func(t *testing.T) {
		changed := append([]catalog.Candidate(nil), candidates...)
		changed = append(changed, candidates[0])
		err := Validate(mapping, evidence, source, 3, contracts, inventory.Concepts, changed)
		assertCurriculumErrorContains(t, err, "duplicate candidate source path")
	})

	t.Run("inventory mismatch", func(t *testing.T) {
		err := Validate(mapping, evidence, source, 3, contracts, inventory.Concepts[:len(inventory.Concepts)-1], candidates)
		assertCurriculumErrorContains(t, err, "concept map inventory differs")
	})
}

func TestConceptInventoryRequiresFrozenSortedSet(t *testing.T) {
	root := curriculumRepositoryRoot(t)
	_, _, source, _, inventory, _ := loadExercismFixtures(t, root)
	if err := ValidateInventory(inventory, source); err != nil {
		t.Fatalf("repository inventory must validate: %v", err)
	}
	inventory.Concepts[1] = inventory.Concepts[0]
	assertCurriculumErrorContains(t, ValidateInventory(inventory, source), "empty, duplicated, or unsorted")
}

func loadExercismFixtures(t *testing.T, root string) (ConceptMap, EvidenceFile, sources.Source, map[string]SkillContract, ConceptInventory, []catalog.Candidate) {
	t.Helper()
	mapping, err := LoadConceptMap(root)
	if err != nil {
		t.Fatal(err)
	}
	evidence, err := LoadEvidence(root)
	if err != nil {
		t.Fatal(err)
	}
	inventory, err := LoadConceptInventory(root)
	if err != nil {
		t.Fatal(err)
	}
	registry, err := sources.LoadRegistry(root)
	if err != nil {
		t.Fatal(err)
	}
	source, err := registry.Find("exercism-go")
	if err != nil {
		t.Fatal(err)
	}
	candidates, err := catalog.ReadJSONL(root, "catalog/candidates/exercism-go.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	taxonomyBytes, err := os.ReadFile(filepath.Join(root, "curriculum", "skills.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var taxonomy struct {
		Domains []struct {
			Skills []struct {
				ID         string `yaml:"id"`
				Status     string `yaml:"status"`
				Assignable bool   `yaml:"assignable"`
			} `yaml:"skills"`
		} `yaml:"domains"`
	}
	if err := yaml.Unmarshal(taxonomyBytes, &taxonomy); err != nil {
		t.Fatal(err)
	}
	contracts := map[string]SkillContract{}
	for _, domain := range taxonomy.Domains {
		for _, skill := range domain.Skills {
			contracts[skill.ID] = SkillContract{Status: skill.Status, Assignable: skill.Assignable}
		}
	}
	return mapping, evidence, source, contracts, inventory, candidates
}

func curriculumRepositoryRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate test file")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func cloneContracts(values map[string]SkillContract) map[string]SkillContract {
	result := make(map[string]SkillContract, len(values))
	for id, contract := range values {
		result[id] = contract
	}
	return result
}

func assertCurriculumErrorContains(t *testing.T, err error, fragment string) {
	t.Helper()
	if err == nil || !strings.Contains(err.Error(), fragment) {
		t.Fatalf("error %v does not contain %q", err, fragment)
	}
}
