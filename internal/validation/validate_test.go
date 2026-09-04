package validation

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

func TestRepositoryPhase1(t *testing.T) {
	root := repositoryRoot(t)
	result, err := ValidateRoot(root, Options{AllowEmpty: true})
	if err != nil {
		t.Fatalf("phase 1 repository must validate: %v", err)
	}
	if got, want := len(result.Curriculum.Domains), 28; got != want {
		t.Fatalf("domains = %d, want %d", got, want)
	}
	if got, want := len(result.Coverage), 142; got != want {
		t.Fatalf("skills = %d, want %d", got, want)
	}
	primary, total := 0, 0
	for _, item := range result.Coverage {
		primary += item.RequiredPrimary
		total += item.RequiredTotal
	}
	if primary != 409 || total != 693 {
		t.Fatalf("minimums = %d primary / %d total, want 409 / 693", primary, total)
	}
	if len(result.Registry.Sources) < 1 {
		t.Fatal("expected audited sources")
	}
	deprecated := 0
	for _, domain := range result.Curriculum.Domains {
		for _, skill := range domain.Skills {
			if skill.Status == "deprecated" {
				deprecated++
			}
		}
	}
	if deprecated != 35 || len(result.Migration.Entries) != 99 || len(result.Migration.Additions) != 8 {
		t.Fatalf("migration inventory = %d deprecated / %d entries / %d additions", deprecated, len(result.Migration.Entries), len(result.Migration.Additions))
	}
}

func TestRepositoryValidationReusesCheckCollection(t *testing.T) {
	root := repositoryRoot(t)
	var aggregate tasking.AggregateReport
	if err := decodeStrict(filepath.Join(root, "reports", "task-quality.yaml"), &aggregate); err != nil {
		t.Fatal(err)
	}
	located, err := tasking.DiscoverTasks(root, nil)
	if err != nil {
		t.Fatal(err)
	}
	byID := map[string]tasking.AggregateTask{}
	for _, item := range aggregate.Tasks {
		byID[item.ID] = item
	}
	toolchainDigest, err := tasking.ToolchainPolicyDigest(root)
	if err != nil {
		t.Fatal(err)
	}
	collection := tasking.CheckCollection{
		SchemaVersion: 1, ContractVersion: tasking.CheckContractVersion,
		ToolchainPolicyDigest: toolchainDigest,
		Execution:             tasking.ExecutionMetrics{Sessions: 1, TaskChecks: len(located)},
	}
	for _, item := range located {
		digest, err := tasking.TaskInputDigest(root, item)
		if err != nil {
			t.Fatal(err)
		}
		quality := byID[item.ID]
		collection.Tasks = append(collection.Tasks, tasking.CheckedTask{
			TaskID: item.ID, State: item.State, InputDigest: digest,
			Report: tasking.QualityReport{
				SchemaVersion: 1, TaskID: item.ID, State: item.State,
				Passed: quality.TechnicalStatus == "pass", TechnicalStatus: quality.TechnicalStatus,
				ReviewStatus: quality.ReviewStatus, PromotionEligible: quality.PromotionEligible,
				Legacy: quality.Legacy, FailingGateIDs: quality.FailingGateIDs,
				StarterStatus: quality.StarterStatus, ReferenceStatus: quality.ReferenceStatus,
				RaceStatus: quality.RaceStatus, Gates: []tasking.Gate{}, Mutants: quality.Mutants,
			},
		})
	}
	if _, err := ValidateRoot(root, Options{AllowEmpty: true, Checks: &collection}); err != nil {
		t.Fatalf("validation rejected reusable collection: %v", err)
	}
}

func TestDecodeStrictRejectsUnknownFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "skills.yaml")
	content := "version: 1\ntrack: go-fundamentals\nunknown_field: true\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	var curriculum Curriculum
	err := decodeStrict(path, &curriculum)
	if err == nil || !strings.Contains(err.Error(), "field unknown_field not found") {
		t.Fatalf("expected unknown-field error, got %v", err)
	}
}

func TestDetectCycles(t *testing.T) {
	skills := map[string]Skill{
		"one.a": {ID: "one.a", Prerequisites: []string{"two.b"}},
		"two.b": {ID: "two.b", Prerequisites: []string{"one.a"}},
	}
	problems := detectCycles(skills)
	if len(problems) == 0 {
		t.Fatal("expected prerequisite cycle")
	}
}

func TestCurriculumV3SemanticContracts(t *testing.T) {
	t.Run("active contract", func(t *testing.T) {
		curriculum := loadCurriculum(t)
		skill := findSkill(t, &curriculum, "types.zero-values")
		skill.CommonMisconceptions = nil
		_, _, problems := validateCurriculum(curriculum)
		assertProblemContains(t, problems, "lacks the full teaching contract")
	})

	t.Run("deprecated prerequisite", func(t *testing.T) {
		curriculum := loadCurriculum(t)
		skill := findSkill(t, &curriculum, "types.zero-values")
		skill.Prerequisites = []string{"types.constants-iota"}
		_, _, problems := validateCurriculum(curriculum)
		assertProblemContains(t, problems, "non-assignable prerequisite")
	})

	t.Run("invalid minimum tier", func(t *testing.T) {
		curriculum := loadCurriculum(t)
		skill := findSkill(t, &curriculum, "types.zero-values")
		skill.MinimumOccurrences = Minimum{Primary: 1, Total: 2}
		_, _, problems := validateCurriculum(curriculum)
		assertProblemContains(t, problems, "minimum tier")
	})

	t.Run("primary plan exceeds task target", func(t *testing.T) {
		curriculum := loadCurriculum(t)
		for domainIndex := range curriculum.Domains {
			for skillIndex := range curriculum.Domains[domainIndex].Skills {
				skill := &curriculum.Domains[domainIndex].Skills[skillIndex]
				if skill.Status == "active" {
					skill.MinimumOccurrences = Minimum{Primary: 4, Total: 6}
				}
			}
		}
		_, _, problems := validateCurriculum(curriculum)
		assertProblemContains(t, problems, "outside target")
	})

	t.Run("cycle", func(t *testing.T) {
		curriculum := loadCurriculum(t)
		findSkill(t, &curriculum, "types.declarations-inference").Prerequisites = []string{"types.zero-values"}
		findSkill(t, &curriculum, "types.zero-values").Prerequisites = []string{"types.declarations-inference"}
		_, _, problems := validateCurriculum(curriculum)
		assertProblemContains(t, problems, "prerequisite cycle")
	})
}

func TestMigrationCompletenessAndCardinality(t *testing.T) {
	root := repositoryRoot(t)
	curriculum := loadCurriculum(t)
	skills, _, curriculumProblems := validateCurriculum(curriculum)
	if len(curriculumProblems) != 0 {
		t.Fatalf("fixture curriculum invalid: %v", curriculumProblems)
	}

	t.Run("missing old id", func(t *testing.T) {
		migration := loadMigration(t)
		migration.Entries = migration.Entries[1:]
		_, problems := validateMigration(root, migration, skills)
		assertProblemContains(t, problems, "missing v2 skill")
	})

	t.Run("split needs multiple targets", func(t *testing.T) {
		migration := loadMigration(t)
		for index := range migration.Entries {
			if migration.Entries[index].MigrationType == "split" {
				migration.Entries[index].NewIDs = migration.Entries[index].NewIDs[:1]
				break
			}
		}
		_, problems := validateMigration(root, migration, skills)
		assertProblemContains(t, problems, "requires at least two targets")
	})
}

func TestDeprecatedTaskSkillRequiresExactLegacyResolution(t *testing.T) {
	curriculum := loadCurriculum(t)
	skills, _, _ := validateCurriculum(curriculum)
	const (
		taskID       = "legacy-check-001"
		field        = "skills.primary"
		oldID        = "types.constants-iota"
		resolvedID   = "types.constants"
		manifestHash = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	)
	index := migrationIndex{legacyByKey: map[string]MigrationTaskRef{
		strings.Join([]string{taskID, field, oldID}, "\x00"): {
			TaskID: taskID, Field: field, OldID: oldID, ManifestSHA256: manifestHash,
			ResolvedSkillID: resolvedID, CoverageBehavior: "count-resolved",
		},
	}}
	if problems := validateTaskSkill("fixture", taskID, field, oldID, manifestHash, skills, index); len(problems) != 0 {
		t.Fatalf("exact legacy resolution rejected: %v", problems)
	}
	problems := validateTaskSkill("fixture", taskID, field, oldID, "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", skills, index)
	assertProblemContains(t, problems, "without an exact legacy resolution")
}

func TestCoverageCountsOnlyReadyPrimaryAndSecondary(t *testing.T) {
	curriculum := Curriculum{
		CoveragePolicy: Policy{CountedStatus: "ready"},
		Domains: []Domain{{ID: "x", Skills: []Skill{
			activeTestSkill("x.primary", Minimum{Primary: 2, Total: 3}),
			activeTestSkill("x.secondary", Minimum{Primary: 1, Total: 2}),
		}}},
	}
	tasks := []Task{
		{QualityStatus: "ready", Skills: TaskSkills{Primary: "x.primary", Secondary: []string{"x.secondary"}}},
		{QualityStatus: "draft", Skills: TaskSkills{Primary: "x.primary", Secondary: []string{"x.secondary"}}},
	}
	coverage := calculateCoverage(curriculum, tasks, map[string]string{"x.primary": "x", "x.secondary": "x"}, migrationIndex{})
	byID := map[string]Coverage{}
	for _, item := range coverage {
		byID[item.SkillID] = item
	}
	if got := byID["x.primary"]; got.ActualPrimary != 1 || got.ActualTotal != 1 {
		t.Fatalf("primary coverage = %+v", got)
	}
	if got := byID["x.secondary"]; got.ActualPrimary != 0 || got.ActualTotal != 1 {
		t.Fatalf("secondary coverage = %+v", got)
	}
}

func activeTestSkill(id string, minimum Minimum) Skill {
	assignable := true
	return Skill{ID: id, Status: "active", Assignable: &assignable, MinimumOccurrences: minimum}
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate test file")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func loadCurriculum(t *testing.T) Curriculum {
	t.Helper()
	var value Curriculum
	if err := decodeStrict(filepath.Join(repositoryRoot(t), "curriculum", "skills.yaml"), &value); err != nil {
		t.Fatal(err)
	}
	return value
}

func loadMigration(t *testing.T) Migration {
	t.Helper()
	var value Migration
	if err := decodeStrict(filepath.Join(repositoryRoot(t), "curriculum", "migrations", "v2-to-v3.yaml"), &value); err != nil {
		t.Fatal(err)
	}
	return value
}

func findSkill(t *testing.T, curriculum *Curriculum, id string) *Skill {
	t.Helper()
	for domainIndex := range curriculum.Domains {
		for skillIndex := range curriculum.Domains[domainIndex].Skills {
			skill := &curriculum.Domains[domainIndex].Skills[skillIndex]
			if skill.ID == id {
				return skill
			}
		}
	}
	t.Fatalf("skill %q not found", id)
	return nil
}

func assertProblemContains(t *testing.T, problems []error, fragment string) {
	t.Helper()
	for _, problem := range problems {
		if problem != nil && strings.Contains(problem.Error(), fragment) {
			return
		}
	}
	t.Fatalf("problems %v do not contain %q", problems, fragment)
}
