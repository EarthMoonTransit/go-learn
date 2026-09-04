package validation

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	v2TaxonomyPath = "curriculum/migrations/snapshots/skills-v2.yaml"
	v2MappingPath  = "curriculum/migrations/snapshots/exercism-map-v2.yaml"
	v2EvidencePath = "curriculum/migrations/snapshots/exercism-evidence-v2.yaml"
)

type migrationIndex struct {
	byOld       map[string]MigrationEntry
	legacyByKey map[string]MigrationTaskRef
}

type baselineEvidence struct {
	Reviews []struct {
		ID      string `yaml:"id"`
		SkillID string `yaml:"skill_id"`
	} `yaml:"reviews"`
}

func validateMigration(root string, migration Migration, skills map[string]Skill) (migrationIndex, []error) {
	index := migrationIndex{byOld: map[string]MigrationEntry{}, legacyByKey: map[string]MigrationTaskRef{}}
	var problems []error
	if migration.SchemaVersion != 1 || migration.FromTaxonomyVersion != 2 || migration.ToTaxonomyVersion != 3 || migration.Decision != "docs/adr/0003-taxonomy-v3.md" {
		problems = append(problems, errors.New("migration: version range and ADR decision must identify taxonomy v2 to v3"))
	}
	if migration.Baseline.TaxonomyPath != v2TaxonomyPath || migration.Baseline.MappingPath != v2MappingPath || migration.Baseline.EvidencePath != v2EvidencePath {
		problems = append(problems, errors.New("migration: baseline paths must identify the immutable v2 snapshots"))
	}

	var old V2Curriculum
	oldBytes, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(v2TaxonomyPath)))
	if err != nil {
		problems = append(problems, fmt.Errorf("migration: read v2 taxonomy snapshot: %w", err))
	} else {
		if digestBytes(oldBytes) != migration.Baseline.TaxonomySHA256 {
			problems = append(problems, errors.New("migration: v2 taxonomy snapshot hash mismatch"))
		}
		if err := yaml.Unmarshal(oldBytes, &old); err != nil {
			problems = append(problems, fmt.Errorf("migration: decode v2 taxonomy snapshot: %w", err))
		}
	}
	if old.Version != 2 {
		problems = append(problems, fmt.Errorf("migration: snapshot taxonomy version must be 2, got %d", old.Version))
	}

	oldIDs := map[string]bool{}
	var sortedOldIDs []string
	for _, domain := range old.Domains {
		for _, skill := range domain.Skills {
			if oldIDs[skill.ID] {
				problems = append(problems, fmt.Errorf("migration: duplicate v2 snapshot skill %q", skill.ID))
			}
			oldIDs[skill.ID] = true
			sortedOldIDs = append(sortedOldIDs, skill.ID)
		}
	}
	sort.Strings(sortedOldIDs)
	if len(sortedOldIDs) != migration.Baseline.SkillCount || migration.Baseline.SkillCount != 99 {
		problems = append(problems, fmt.Errorf("migration: v2 skill count is %d, baseline declares %d", len(sortedOldIDs), migration.Baseline.SkillCount))
	}
	if digestBytes([]byte(strings.Join(sortedOldIDs, "\n")+"\n")) != migration.Baseline.SkillIDsSHA256 {
		problems = append(problems, errors.New("migration: v2 skill id inventory hash mismatch"))
	}

	mappingBytes, mappingErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(v2MappingPath)))
	if mappingErr != nil {
		problems = append(problems, fmt.Errorf("migration: read v2 mapping snapshot: %w", mappingErr))
	} else if digestBytes(mappingBytes) != migration.Baseline.MappingSHA256 {
		problems = append(problems, errors.New("migration: v2 mapping snapshot hash mismatch"))
	}
	evidenceBySkill := map[string][]string{}
	evidenceBytes, evidenceErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(v2EvidencePath)))
	if evidenceErr != nil {
		problems = append(problems, fmt.Errorf("migration: read v2 evidence snapshot: %w", evidenceErr))
	} else {
		if digestBytes(evidenceBytes) != migration.Baseline.EvidenceSHA256 {
			problems = append(problems, errors.New("migration: v2 evidence snapshot hash mismatch"))
		}
		var oldEvidence baselineEvidence
		if err := yaml.Unmarshal(evidenceBytes, &oldEvidence); err != nil {
			problems = append(problems, fmt.Errorf("migration: decode v2 evidence snapshot: %w", err))
		} else {
			if len(oldEvidence.Reviews) != migration.Baseline.EvidenceReviewCount {
				problems = append(problems, fmt.Errorf("migration: v2 evidence review count is %d, baseline declares %d", len(oldEvidence.Reviews), migration.Baseline.EvidenceReviewCount))
			}
			for _, review := range oldEvidence.Reviews {
				evidenceBySkill[review.SkillID] = append(evidenceBySkill[review.SkillID], review.ID)
			}
		}
	}

	targeted := map[string]bool{}
	mergedTargetSources := map[string]int{}
	expectedLegacy := map[string]MigrationTaskRef{}
	lastOldID := ""
	for _, entry := range migration.Entries {
		if entry.OldID == "" || index.byOld[entry.OldID].OldID != "" {
			problems = append(problems, fmt.Errorf("migration: duplicate or empty old_id %q", entry.OldID))
		}
		if lastOldID != "" && entry.OldID <= lastOldID {
			problems = append(problems, fmt.Errorf("migration: entries are not strictly sorted at %q", entry.OldID))
		}
		lastOldID = entry.OldID
		index.byOld[entry.OldID] = entry
		if !oldIDs[entry.OldID] {
			problems = append(problems, fmt.Errorf("migration: unknown v2 skill %q", entry.OldID))
		}
		if entry.Rationale == "" {
			problems = append(problems, fmt.Errorf("migration: %q requires rationale", entry.OldID))
		}
		if !strictlySorted(entry.NewIDs) {
			problems = append(problems, fmt.Errorf("migration: %q new_ids must be sorted and unique", entry.OldID))
		}
		for _, target := range entry.NewIDs {
			value, ok := skills[target]
			if !ok || !isAssignable(value) {
				problems = append(problems, fmt.Errorf("migration: %q targets unknown or non-assignable skill %q", entry.OldID, target))
			}
			targeted[target] = true
			mergedTargetSources[target]++
		}
		current, currentExists := skills[entry.OldID]
		switch entry.MigrationType {
		case "preserved":
			if len(entry.NewIDs) != 1 || entry.NewIDs[0] != entry.OldID || !currentExists || !isAssignable(current) || entry.CompatibilityBehavior != "active-id" {
				problems = append(problems, fmt.Errorf("migration: preserved %q must remain the same active id", entry.OldID))
			}
		case "split":
			if len(entry.NewIDs) < 2 {
				problems = append(problems, fmt.Errorf("migration: split %q requires at least two targets", entry.OldID))
			}
			problems = append(problems, validateDeprecatedMigrationSkill(entry, current, currentExists)...)
		case "renamed":
			if len(entry.NewIDs) != 1 || entry.NewIDs[0] == entry.OldID {
				problems = append(problems, fmt.Errorf("migration: renamed %q requires one different target", entry.OldID))
			}
			problems = append(problems, validateDeprecatedMigrationSkill(entry, current, currentExists)...)
		case "deprecated":
			if len(entry.NewIDs) != 0 {
				problems = append(problems, fmt.Errorf("migration: deprecated %q cannot have targets", entry.OldID))
			}
			problems = append(problems, validateDeprecatedMigrationSkill(entry, current, currentExists)...)
		case "merged":
			if len(entry.NewIDs) != 1 {
				problems = append(problems, fmt.Errorf("migration: merged %q requires one target", entry.OldID))
			}
			problems = append(problems, validateDeprecatedMigrationSkill(entry, current, currentExists)...)
		default:
			problems = append(problems, fmt.Errorf("migration: %q has invalid type %q", entry.OldID, entry.MigrationType))
		}
		actualEvidence := append([]string(nil), evidenceBySkill[entry.OldID]...)
		sort.Strings(actualEvidence)
		if !equalStringSlices(actualEvidence, entry.AffectedEvidence) {
			problems = append(problems, fmt.Errorf("migration: %q affected_evidence differs from v2 snapshot", entry.OldID))
		}
		if entry.MigrationType == "preserved" && entry.CompatibilityBehavior != "active-id" {
			problems = append(problems, fmt.Errorf("migration: preserved %q requires active-id compatibility", entry.OldID))
		}
		if entry.MigrationType != "preserved" {
			want := "reject-deprecated"
			if len(entry.AffectedTaskReferences) > 0 {
				want = "explicit-legacy-resolution"
			}
			if entry.CompatibilityBehavior != want {
				problems = append(problems, fmt.Errorf("migration: %q compatibility_behavior must be %q", entry.OldID, want))
			}
		}
		for _, ref := range entry.AffectedTaskReferences {
			key := taskRefKey(ref)
			if ref.OldID != entry.OldID || !taskIDPattern.MatchString(ref.TaskID) || !set("skills.primary", "skills.secondary", "prerequisites", "new_concepts")[ref.Field] || !sha256Pattern.MatchString(ref.ManifestSHA256) {
				problems = append(problems, fmt.Errorf("migration: %q has malformed affected task reference %s", entry.OldID, key))
			}
			if target, ok := skills[ref.ResolvedSkillID]; !ok || !isAssignable(target) {
				problems = append(problems, fmt.Errorf("migration: affected task reference %s has non-assignable resolution", key))
			}
			if entry.MigrationType == "preserved" {
				if ref.ResolvedSkillID != entry.OldID || ref.CoverageBehavior != "native" {
					problems = append(problems, fmt.Errorf("migration: preserved task reference %s must remain native", key))
				}
				continue
			}
			if ref.CoverageBehavior == "native" {
				problems = append(problems, fmt.Errorf("migration: changed task reference %s cannot use native coverage", key))
			}
			if _, duplicate := expectedLegacy[key]; duplicate {
				problems = append(problems, fmt.Errorf("migration: duplicate affected changed task reference %s", key))
			}
			expectedLegacy[key] = ref
		}
	}
	for _, oldID := range sortedOldIDs {
		if _, ok := index.byOld[oldID]; !ok {
			problems = append(problems, fmt.Errorf("migration: missing v2 skill %q", oldID))
		}
	}
	if len(migration.Entries) != len(sortedOldIDs) {
		problems = append(problems, fmt.Errorf("migration: entries=%d, v2 skills=%d", len(migration.Entries), len(sortedOldIDs)))
	}
	for _, entry := range migration.Entries {
		if entry.MigrationType == "merged" && len(entry.NewIDs) == 1 && mergedTargetSources[entry.NewIDs[0]] < 2 {
			problems = append(problems, fmt.Errorf("migration: merged %q target %q has no second source", entry.OldID, entry.NewIDs[0]))
		}
	}

	lastAddition := ""
	additions := map[string]bool{}
	for _, addition := range migration.Additions {
		if addition.ID == "" || addition.Rationale == "" || additions[addition.ID] {
			problems = append(problems, fmt.Errorf("migration: duplicate or incomplete addition %q", addition.ID))
		}
		if lastAddition != "" && addition.ID <= lastAddition {
			problems = append(problems, fmt.Errorf("migration: additions are not strictly sorted at %q", addition.ID))
		}
		lastAddition = addition.ID
		additions[addition.ID] = true
		if targeted[addition.ID] {
			problems = append(problems, fmt.Errorf("migration: addition %q is also a migrated target", addition.ID))
		}
		if value, ok := skills[addition.ID]; !ok || !isAssignable(value) {
			problems = append(problems, fmt.Errorf("migration: addition %q is not active and assignable", addition.ID))
		}
	}
	assignableCount := 0
	for id, value := range skills {
		if isAssignable(value) {
			assignableCount++
			if !targeted[id] && !additions[id] {
				problems = append(problems, fmt.Errorf("migration: active skill %q is neither migrated nor added", id))
			}
		} else if value.Status == "deprecated" {
			entry, ok := index.byOld[id]
			if !ok || entry.MigrationType == "preserved" {
				problems = append(problems, fmt.Errorf("migration: deprecated skill %q lacks a non-preserved v2 entry", id))
			}
		}
	}
	if assignableCount != migration.TargetAssignableSkillCount {
		problems = append(problems, fmt.Errorf("migration: assignable count %d differs from declared %d", assignableCount, migration.TargetAssignableSkillCount))
	}

	for _, ref := range migration.LegacyTaskReferences {
		key := taskRefKey(ref)
		if _, exists := index.legacyByKey[key]; exists {
			problems = append(problems, fmt.Errorf("migration: duplicate legacy task reference %s", key))
		}
		entry, ok := index.byOld[ref.OldID]
		if !ok || entry.MigrationType == "preserved" || !containsTaskRef(entry.AffectedTaskReferences, ref) {
			problems = append(problems, fmt.Errorf("migration: legacy task reference %s is not declared by a changed entry", key))
		}
		if target, ok := skills[ref.ResolvedSkillID]; !ok || !isAssignable(target) {
			problems = append(problems, fmt.Errorf("migration: legacy task reference %s has non-assignable resolution", key))
		}
		if ref.CoverageBehavior != "count-resolved" && ref.CoverageBehavior != "no-coverage" {
			problems = append(problems, fmt.Errorf("migration: legacy task reference %s has invalid coverage behavior", key))
		}
		index.legacyByKey[key] = ref
	}
	for key, expected := range expectedLegacy {
		actual, ok := index.legacyByKey[key]
		if !ok || !containsTaskRef([]MigrationTaskRef{actual}, expected) {
			problems = append(problems, fmt.Errorf("migration: changed task reference %s is missing from legacy_task_references", key))
		}
	}
	if len(index.legacyByKey) != len(expectedLegacy) {
		problems = append(problems, fmt.Errorf("migration: legacy_task_references has %d records, expected %d", len(index.legacyByKey), len(expectedLegacy)))
	}
	return index, problems
}

func validateDeprecatedMigrationSkill(entry MigrationEntry, current Skill, exists bool) []error {
	if !exists || current.Status != "deprecated" || isAssignable(current) || !equalStringSlices(current.ReplacedBy, entry.NewIDs) {
		return []error{fmt.Errorf("migration: changed v2 skill %q must remain a non-assignable deprecated stub with exact replacements", entry.OldID)}
	}
	return nil
}

func validateMigrationTaskReferences(root string, migration Migration, tasks []Task) []error {
	var problems []error
	taskByID := map[string]Task{}
	for _, task := range tasks {
		taskByID[task.ID] = task
	}
	for _, entry := range migration.Entries {
		for _, ref := range entry.AffectedTaskReferences {
			task, ok := taskByID[ref.TaskID]
			if !ok {
				problems = append(problems, fmt.Errorf("migration: affected task %q is missing", ref.TaskID))
				continue
			}
			path := filepath.Join(root, "tasks", "ready", ref.TaskID, "task.yaml")
			b, err := os.ReadFile(path)
			if err != nil || digestBytes(b) != ref.ManifestSHA256 {
				problems = append(problems, fmt.Errorf("migration: affected task %q manifest hash mismatch", ref.TaskID))
			}
			if !taskHasFieldValue(task, ref.Field, ref.OldID) {
				problems = append(problems, fmt.Errorf("migration: affected task %q no longer has %q in %s", ref.TaskID, ref.OldID, ref.Field))
			}
		}
	}
	return problems
}

func taskHasFieldValue(task Task, field, value string) bool {
	switch field {
	case "skills.primary":
		return task.Skills.Primary == value
	case "skills.secondary":
		return containsString(task.Skills.Secondary, value)
	case "prerequisites":
		return containsString(task.Prerequisites, value)
	case "new_concepts":
		return containsString(task.NewConcepts, value)
	default:
		return false
	}
}

func migrationResolution(index migrationIndex, taskID, field, rawID, manifestHash string) (string, bool) {
	ref, ok := index.legacyByKey[strings.Join([]string{taskID, field, rawID}, "\x00")]
	if !ok || ref.ManifestSHA256 != manifestHash {
		return "", false
	}
	return ref.ResolvedSkillID, true
}

func taskRefKey(ref MigrationTaskRef) string {
	return strings.Join([]string{ref.TaskID, ref.Field, ref.OldID}, "\x00")
}

func containsTaskRef(values []MigrationTaskRef, wanted MigrationTaskRef) bool {
	wantedKey := taskRefKey(wanted)
	for _, value := range values {
		if taskRefKey(value) == wantedKey && value.ManifestSHA256 == wanted.ManifestSHA256 && value.ResolvedSkillID == wanted.ResolvedSkillID && value.CoverageBehavior == wanted.CoverageBehavior {
			return true
		}
	}
	return false
}

func strictlySorted(values []string) bool {
	for index := 1; index < len(values); index++ {
		if values[index] <= values[index-1] {
			return false
		}
	}
	return true
}

func equalStringSlices(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

func digestBytes(value []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(value))
}

func isAssignable(skill Skill) bool {
	return skill.Assignable != nil && *skill.Assignable && skill.Status == "active"
}

func containsString(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}
