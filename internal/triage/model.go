package triage

type ReviewFile struct {
	SchemaVersion   int               `yaml:"schema_version"`
	SourceID        string            `yaml:"source_id"`
	SourceCommit    string            `yaml:"source_commit"`
	TaxonomyVersion int               `yaml:"taxonomy_version"`
	ReviewedBy      string            `yaml:"reviewed_by"`
	ReviewedAt      string            `yaml:"reviewed_at"`
	Reviews         []CandidateReview `yaml:"reviews"`
}

type CandidateReview struct {
	CandidateID             string   `yaml:"candidate_id"`
	SourcePath              string   `yaml:"source_path"`
	SourceCommit            string   `yaml:"source_commit"`
	ExerciseKind            string   `yaml:"exercise_kind"`
	ProposedPrimarySkills   []string `yaml:"proposed_primary_skills"`
	ProposedSecondarySkills []string `yaml:"proposed_secondary_skills"`
	Prerequisites           []string `yaml:"prerequisites"`
	ConceptualMechanisms    []string `yaml:"conceptual_mechanisms"`
	LikelyMisconceptions    []string `yaml:"likely_misconceptions"`
	PossibleTaskForms       []string `yaml:"possible_task_forms"`
	DuplicateCluster        string   `yaml:"duplicate_cluster,omitempty"`
	ProvenanceStatus        string   `yaml:"provenance_status"`
	ReviewStatus            string   `yaml:"review_status"`
	Scores                  Scores   `yaml:"scores"`
	DecisionReason          string   `yaml:"decision_reason"`
	ReviewerRationale       string   `yaml:"reviewer_rationale"`
}

type Scores struct {
	GoSpecificity                 int `yaml:"go_specificity"`
	ConceptualFocus               int `yaml:"conceptual_focus"`
	TestQuality                   int `yaml:"test_quality"`
	EdgeCaseQuality               int `yaml:"edge_case_quality"`
	ProductionTransfer            int `yaml:"production_transfer"`
	RewriteBurden                 int `yaml:"rewrite_burden"`
	Uniqueness                    int `yaml:"uniqueness"`
	SuitabilityBeginners          int `yaml:"suitability_for_beginners"`
	SuitabilityDeliberatePractice int `yaml:"suitability_for_deliberate_practice"`
}

type ClusterFile struct {
	SchemaVersion int       `yaml:"schema_version"`
	SourceID      string    `yaml:"source_id"`
	SourceCommit  string    `yaml:"source_commit"`
	ReviewedBy    string    `yaml:"reviewed_by"`
	ReviewedAt    string    `yaml:"reviewed_at"`
	Clusters      []Cluster `yaml:"clusters"`
}

type Cluster struct {
	ID                  string         `yaml:"id"`
	Members             []string       `yaml:"members"`
	SignalTypes         []string       `yaml:"signal_types"`
	Signals             ClusterSignals `yaml:"signals"`
	Decision            string         `yaml:"decision"`
	PreferredCandidates []string       `yaml:"preferred_candidates"`
	ReviewerRationale   string         `yaml:"reviewer_rationale"`
}

type ClusterSignals struct {
	SharedFileHashes           []string `yaml:"shared_file_hashes"`
	SharedTestHashes           []string `yaml:"shared_test_hashes"`
	ConceptSets                []string `yaml:"concept_sets"`
	PublicAPISimilarity        string   `yaml:"public_api_similarity"`
	ASTSimilarity              string   `yaml:"ast_similarity"`
	ReviewedSemanticSimilarity string   `yaml:"reviewed_semantic_similarity"`
}

type SeedPlan struct {
	SchemaVersion   int             `yaml:"schema_version"`
	SourceID        string          `yaml:"source_id"`
	SourceCommit    string          `yaml:"source_commit"`
	TaxonomyVersion int             `yaml:"taxonomy_version"`
	ReviewedBy      string          `yaml:"reviewed_by"`
	ReviewedAt      string          `yaml:"reviewed_at"`
	SelectedSeeds   []SeedSelection `yaml:"selected_seeds"`
	Gaps            []SeedGap       `yaml:"gaps"`
}

type SeedSelection struct {
	CandidateID        string   `yaml:"candidate_id"`
	TargetSkills       []string `yaml:"target_skills"`
	IntendedTaskForms  []string `yaml:"intended_task_forms"`
	PilotRole          string   `yaml:"pilot_role"`
	SelectionRationale string   `yaml:"selection_rationale"`
}

type SeedGap struct {
	SkillIDs       []string `yaml:"skill_ids"`
	RequiredOrigin string   `yaml:"required_origin"`
	Rationale      string   `yaml:"rationale"`
}

type SkillContract struct {
	ID                       string
	Domain                   string
	Definition               string
	Prerequisites            []string
	CommonMisconceptions     []string
	RecommendedExerciseTypes []string
	Status                   string
	Assignable               bool
}

type Bundle struct {
	Reviews  ReviewFile
	Clusters ClusterFile
	Seeds    SeedPlan
}
