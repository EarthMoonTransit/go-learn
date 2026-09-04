package validation

import "go-practice-bank/internal/sources"

type Curriculum struct {
	Version          int      `yaml:"version"`
	Track            string   `yaml:"track"`
	TargetReadyTasks Target   `yaml:"target_ready_tasks"`
	CoveragePolicy   Policy   `yaml:"coverage_policy"`
	Domains          []Domain `yaml:"domains"`
}

type Target struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

type Policy struct {
	CountedStatus                string `yaml:"counted_status"`
	PrimarySkillPerTask          int    `yaml:"primary_skill_per_task"`
	PrerequisiteCountsOccurrence bool   `yaml:"prerequisite_counts_as_occurrence"`
	Note                         string `yaml:"note"`
}

type Domain struct {
	ID     string  `yaml:"id"`
	Title  string  `yaml:"title"`
	Skills []Skill `yaml:"skills"`
}

type Skill struct {
	ID                       string               `yaml:"id"`
	Domain                   string               `yaml:"domain"`
	Definition               string               `yaml:"definition"`
	Prerequisites            []string             `yaml:"prerequisites"`
	CommonMisconceptions     []string             `yaml:"common_misconceptions"`
	RecommendedExerciseTypes []string             `yaml:"recommended_exercise_types"`
	MinimumOccurrences       Minimum              `yaml:"minimum_occurrences"`
	Mastery                  []string             `yaml:"mastery"`
	EvidenceRequirements     EvidenceRequirements `yaml:"evidence_requirements"`
	Status                   string               `yaml:"status"`
	Assignable               *bool                `yaml:"assignable"`
	TaxonomyVersion          int                  `yaml:"taxonomy_version"`
	ReplacedBy               []string             `yaml:"replaced_by,omitempty"`
}

type EvidenceRequirements struct {
	MinimumDistinctTaskTypes int      `yaml:"minimum_distinct_task_types"`
	RequiredModes            []string `yaml:"required_modes"`
	ObservableOutcomes       []string `yaml:"observable_outcomes"`
}

type Migration struct {
	SchemaVersion              int                 `yaml:"schema_version"`
	FromTaxonomyVersion        int                 `yaml:"from_taxonomy_version"`
	ToTaxonomyVersion          int                 `yaml:"to_taxonomy_version"`
	Decision                   string              `yaml:"decision"`
	Baseline                   MigrationBaseline   `yaml:"baseline"`
	TargetAssignableSkillCount int                 `yaml:"target_assignable_skill_count"`
	Entries                    []MigrationEntry    `yaml:"entries"`
	Additions                  []MigrationAddition `yaml:"additions"`
	LegacyTaskReferences       []MigrationTaskRef  `yaml:"legacy_task_references"`
}

type MigrationBaseline struct {
	TaxonomyPath        string `yaml:"taxonomy_path"`
	TaxonomySHA256      string `yaml:"taxonomy_sha256"`
	SkillCount          int    `yaml:"skill_count"`
	SkillIDsSHA256      string `yaml:"skill_ids_sha256"`
	MappingPath         string `yaml:"mapping_path"`
	MappingSHA256       string `yaml:"mapping_sha256"`
	EvidencePath        string `yaml:"evidence_path"`
	EvidenceSHA256      string `yaml:"evidence_sha256"`
	EvidenceReviewCount int    `yaml:"evidence_review_count"`
}

type MigrationEntry struct {
	OldID                  string             `yaml:"old_id"`
	NewIDs                 []string           `yaml:"new_ids"`
	MigrationType          string             `yaml:"migration_type"`
	Rationale              string             `yaml:"rationale"`
	AffectedEvidence       []string           `yaml:"affected_evidence"`
	AffectedTaskReferences []MigrationTaskRef `yaml:"affected_task_references"`
	CompatibilityBehavior  string             `yaml:"compatibility_behavior"`
}

type MigrationAddition struct {
	ID        string `yaml:"id"`
	Rationale string `yaml:"rationale"`
}

type MigrationTaskRef struct {
	TaskID           string `yaml:"task_id"`
	Field            string `yaml:"field"`
	OldID            string `yaml:"old_id"`
	ManifestSHA256   string `yaml:"manifest_sha256"`
	ResolvedSkillID  string `yaml:"resolved_skill_id"`
	CoverageBehavior string `yaml:"coverage_behavior"`
}

type V2Curriculum struct {
	Version          int        `yaml:"version"`
	Track            string     `yaml:"track"`
	TargetReadyTasks Target     `yaml:"target_ready_tasks"`
	CoveragePolicy   Policy     `yaml:"coverage_policy"`
	Domains          []V2Domain `yaml:"domains"`
}

type V2Domain struct {
	ID     string    `yaml:"id"`
	Title  string    `yaml:"title"`
	Skills []V2Skill `yaml:"skills"`
}

type V2Skill struct {
	ID                       string   `yaml:"id"`
	Prerequisites            []string `yaml:"prerequisites"`
	CommonMisconceptions     []string `yaml:"common_misconceptions"`
	RecommendedExerciseTypes []string `yaml:"recommended_exercise_types"`
	MinimumOccurrences       Minimum  `yaml:"minimum_occurrences"`
	Mastery                  []string `yaml:"mastery"`
}

type Minimum struct {
	Primary int `yaml:"primary"`
	Total   int `yaml:"total"`
}

type Registry = sources.Registry
type Source = sources.Source
type SourceLicense = sources.SourceLicense

type Task struct {
	SchemaVersion        int        `yaml:"schema_version"`
	ID                   string     `yaml:"id"`
	Title                string     `yaml:"title"`
	TaskType             string     `yaml:"task_type"`
	Skills               TaskSkills `yaml:"skills"`
	Prerequisites        []string   `yaml:"prerequisites"`
	NewConcepts          []string   `yaml:"new_concepts"`
	Difficulty           string     `yaml:"difficulty"`
	EstimatedTimeMinutes TimeRange  `yaml:"estimated_time_minutes"`
	Source               TaskSource `yaml:"source"`
	TransformationType   string     `yaml:"transformation_type"`
	ValidationCommands   []string   `yaml:"validation_commands"`
	QualityStatus        string     `yaml:"quality_status"`
}

type TaskSkills struct {
	Primary   string   `yaml:"primary"`
	Secondary []string `yaml:"secondary"`
}

type TimeRange struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

type TaskSource struct {
	ID         string `yaml:"id"`
	Repository string `yaml:"repository"`
	Path       string `yaml:"path"`
	Commit     string `yaml:"commit"`
	License    string `yaml:"license"`
}

type Coverage struct {
	SkillID         string
	DomainID        string
	RequiredPrimary int
	RequiredTotal   int
	ActualPrimary   int
	ActualTotal     int
}

type Result struct {
	Curriculum Curriculum
	Migration  Migration
	Registry   Registry
	Tasks      []Task
	Coverage   []Coverage
	ReadyTasks int
}

type Options struct {
	EnforceCoverage bool
	AllowEmpty      bool
	RunCommands     bool
}
