package tasking

type State string

const (
	StateDraft    State = "draft"
	StateInReview State = "in-review"
	StateReady    State = "ready"
	StateRejected State = "rejected"
	StateRetired  State = "retired"
)

type Task struct {
	SchemaVersion        int        `yaml:"schema_version" json:"schema_version"`
	ID                   string     `yaml:"id" json:"id"`
	Title                string     `yaml:"title" json:"title"`
	TaskType             string     `yaml:"task_type" json:"task_type"`
	Skills               TaskSkills `yaml:"skills" json:"skills"`
	Prerequisites        []string   `yaml:"prerequisites" json:"prerequisites"`
	NewConcepts          []string   `yaml:"new_concepts" json:"new_concepts"`
	Difficulty           string     `yaml:"difficulty" json:"difficulty"`
	EstimatedTimeMinutes TimeRange  `yaml:"estimated_time_minutes" json:"estimated_time_minutes"`
	Source               TaskSource `yaml:"source" json:"source"`
	TransformationType   string     `yaml:"transformation_type" json:"transformation_type"`
	ValidationCommands   []string   `yaml:"validation_commands" json:"validation_commands"`
	QualityStatus        string     `yaml:"quality_status" json:"quality_status"`
}

type TaskSkills struct {
	Primary   string   `yaml:"primary" json:"primary"`
	Secondary []string `yaml:"secondary" json:"secondary"`
}

type TimeRange struct {
	Min int `yaml:"min" json:"min"`
	Max int `yaml:"max" json:"max"`
}

type TaskSource struct {
	ID         string `yaml:"id" json:"id"`
	Repository string `yaml:"repository" json:"repository"`
	Path       string `yaml:"path" json:"path"`
	Commit     string `yaml:"commit" json:"commit"`
	License    string `yaml:"license" json:"license"`
}

// TaskV2 separates what the learner does, what files are submitted, and how
// the submission is verified. It deliberately contains no executable command
// strings; the checker owns the finite verification profiles.
type TaskV2 struct {
	SchemaVersion        int              `yaml:"schema_version" json:"schema_version"`
	ID                   string           `yaml:"id" json:"id"`
	Title                string           `yaml:"title" json:"title"`
	Activity             string           `yaml:"activity" json:"activity"`
	Submission           Submission       `yaml:"submission" json:"submission"`
	Skills               TaskSkills       `yaml:"skills" json:"skills"`
	Prerequisites        []string         `yaml:"prerequisites" json:"prerequisites"`
	NewConcepts          []string         `yaml:"new_concepts" json:"new_concepts"`
	Difficulty           string           `yaml:"difficulty" json:"difficulty"`
	EstimatedTimeMinutes TimeRange        `yaml:"estimated_time_minutes" json:"estimated_time_minutes"`
	Verification         VerificationPlan `yaml:"verification" json:"verification"`
	QualityStatus        string           `yaml:"quality_status" json:"quality_status"`
}

type Submission struct {
	Kind              string   `yaml:"kind" json:"kind"`
	EditablePaths     []string `yaml:"editable_paths" json:"editable_paths"`
	WorkspaceSkeleton []string `yaml:"workspace_skeleton" json:"workspace_skeleton"`
}

type VerificationPlan struct {
	Profile         string             `yaml:"profile" json:"profile"`
	Packages        []string           `yaml:"packages" json:"packages"`
	PublicChecks    []string           `yaml:"public_checks" json:"public_checks"`
	GraderChecks    []string           `yaml:"grader_checks" json:"grader_checks"`
	ReferenceChecks []string           `yaml:"reference_checks" json:"reference_checks"`
	RequiredMutants []string           `yaml:"required_mutants" json:"required_mutants"`
	RaceRequired    bool               `yaml:"race_required" json:"race_required"`
	Vet             bool               `yaml:"vet" json:"vet"`
	Limits          VerificationLimits `yaml:"limits" json:"limits"`
}

type VerificationLimits struct {
	CommandTimeoutSeconds int `yaml:"command_timeout_seconds" json:"command_timeout_seconds"`
	DeterminismRuns       int `yaml:"determinism_runs" json:"determinism_runs"`
	FuzzSeconds           int `yaml:"fuzz_seconds" json:"fuzz_seconds"`
	BenchmarkRuns         int `yaml:"benchmark_runs" json:"benchmark_runs"`
}

type Lifecycle struct {
	SchemaVersion int          `yaml:"schema_version" json:"schema_version"`
	State         State        `yaml:"state" json:"state"`
	UpdatedAt     string       `yaml:"updated_at" json:"updated_at"`
	Transitions   []Transition `yaml:"transitions" json:"transitions"`
}

type Transition struct {
	From   State  `yaml:"from" json:"from"`
	To     State  `yaml:"to" json:"to"`
	At     string `yaml:"at" json:"at"`
	Actor  string `yaml:"actor" json:"actor"`
	Reason string `yaml:"reason" json:"reason"`
}

type Review struct {
	SchemaVersion                  int    `yaml:"schema_version" json:"schema_version"`
	Reviewer                       string `yaml:"reviewer" json:"reviewer"`
	ReviewedAt                     string `yaml:"reviewed_at" json:"reviewed_at"`
	Decision                       string `yaml:"decision" json:"decision"`
	StoryFree                      bool   `yaml:"story_free" json:"story_free"`
	PromptTestAlignment            bool   `yaml:"prompt_test_alignment" json:"prompt_test_alignment"`
	PublicAPIStable                bool   `yaml:"public_api_stable" json:"public_api_stable"`
	ProvenanceVerified             bool   `yaml:"provenance_verified" json:"provenance_verified"`
	LicenseVerified                bool   `yaml:"license_verified" json:"license_verified"`
	ReferenceOnlyCleanRoomVerified bool   `yaml:"reference_only_clean_room_verified" json:"reference_only_clean_room_verified"`
	SolutionLeakageChecked         bool   `yaml:"solution_leakage_checked" json:"solution_leakage_checked"`
	MinimumMutants                 int    `yaml:"minimum_mutants" json:"minimum_mutants"`
	Notes                          string `yaml:"notes" json:"notes"`
}

type Provenance struct {
	SchemaVersion        int                 `yaml:"schema_version" json:"schema_version"`
	Source               ProvenanceSource    `yaml:"source" json:"source"`
	UpstreamAuthors      []string            `yaml:"upstream_authors" json:"upstream_authors"`
	UpstreamContributors []string            `yaml:"upstream_contributors" json:"upstream_contributors"`
	Transformation       ProvenanceTransform `yaml:"transformation" json:"transformation"`
	AttributionFile      string              `yaml:"attribution_file" json:"attribution_file"`
}

type ProvenanceSource struct {
	ID         string `yaml:"id" json:"id"`
	Repository string `yaml:"repository" json:"repository"`
	Path       string `yaml:"path" json:"path"`
	Commit     string `yaml:"commit" json:"commit"`
	License    string `yaml:"license" json:"license"`
	UsageMode  string `yaml:"usage_mode" json:"usage_mode"`
}

type ProvenanceTransform struct {
	Type                  string   `yaml:"type" json:"type"`
	Summary               string   `yaml:"summary" json:"summary"`
	IndependentlyAuthored []string `yaml:"independently_authored" json:"independently_authored"`
	Excluded              []string `yaml:"excluded" json:"excluded"`
}

type ProvenanceV2 struct {
	SchemaVersion   int                  `yaml:"schema_version" json:"schema_version"`
	Kind            string               `yaml:"kind" json:"kind"`
	Project         ProvenanceProject    `yaml:"project" json:"project"`
	Creation        ProvenanceCreation   `yaml:"creation" json:"creation"`
	Sources         []ProvenanceSourceV2 `yaml:"sources" json:"sources"`
	AttributionFile string               `yaml:"attribution_file" json:"attribution_file"`
	Release         ProvenanceRelease    `yaml:"release" json:"release"`
}

type ProvenanceProject struct {
	Identity      string `yaml:"identity" json:"identity"`
	Repository    string `yaml:"repository" json:"repository"`
	License       string `yaml:"license" json:"license"`
	ContentDigest string `yaml:"content_digest" json:"content_digest"`
}

type ProvenanceCreation struct {
	Author    string `yaml:"author" json:"author"`
	CreatedAt string `yaml:"created_at" json:"created_at"`
	Summary   string `yaml:"summary" json:"summary"`
}

type ProvenanceSourceV2 struct {
	RegistryID string                 `yaml:"registry_id" json:"registry_id"`
	Repository string                 `yaml:"repository" json:"repository"`
	Commit     string                 `yaml:"commit" json:"commit"`
	License    string                 `yaml:"license" json:"license"`
	UsageMode  string                 `yaml:"usage_mode" json:"usage_mode"`
	Role       string                 `yaml:"role" json:"role"`
	Artifacts  []ProvenanceArtifactV2 `yaml:"artifacts" json:"artifacts"`
}

type ProvenanceArtifactV2 struct {
	Path   string `yaml:"path" json:"path"`
	SHA256 string `yaml:"sha256" json:"sha256"`
	Usage  string `yaml:"usage" json:"usage"`
}

type ProvenanceRelease struct {
	Binding  string `yaml:"binding" json:"binding"`
	Manifest string `yaml:"manifest" json:"manifest"`
}

type Gate struct {
	ID      string   `yaml:"id" json:"id"`
	Status  string   `yaml:"status" json:"status"`
	Details []string `yaml:"details" json:"details"`
}

type MutantSummary struct {
	Required int      `yaml:"required" json:"required"`
	Total    int      `yaml:"total" json:"total"`
	Detected int      `yaml:"detected" json:"detected"`
	Survived []string `yaml:"survived" json:"survived"`
}

type QualityReport struct {
	SchemaVersion     int           `yaml:"schema_version" json:"schema_version"`
	TaskID            string        `yaml:"task_id" json:"task_id"`
	State             State         `yaml:"state" json:"state"`
	Passed            bool          `yaml:"passed" json:"passed"`
	TechnicalStatus   string        `yaml:"technical_status" json:"technical_status"`
	ReviewStatus      string        `yaml:"review_status" json:"review_status"`
	PromotionEligible bool          `yaml:"promotion_eligible" json:"promotion_eligible"`
	Legacy            bool          `yaml:"legacy" json:"legacy"`
	FailingGateIDs    []string      `yaml:"failing_gate_ids" json:"failing_gate_ids"`
	StarterStatus     string        `yaml:"starter_status" json:"starter_status"`
	ReferenceStatus   string        `yaml:"reference_status" json:"reference_status"`
	RaceStatus        string        `yaml:"race_status" json:"race_status"`
	Gates             []Gate        `yaml:"gates" json:"gates"`
	Mutants           MutantSummary `yaml:"mutants" json:"mutants"`
}

const CheckContractVersion = "task-check/v3"

type ExecutionMetrics struct {
	Sessions       int `yaml:"sessions" json:"sessions"`
	TaskChecks     int `yaml:"task_checks" json:"task_checks"`
	Workspaces     int `yaml:"workspaces" json:"workspaces"`
	VariantRuns    int `yaml:"variant_runs" json:"variant_runs"`
	LegacyCommands int `yaml:"legacy_commands" json:"legacy_commands"`
	ChildGoRuns    int `yaml:"child_go_runs" json:"child_go_runs"`
}

type CheckedTask struct {
	TaskID      string        `yaml:"task_id" json:"task_id"`
	State       State         `yaml:"state" json:"state"`
	InputDigest string        `yaml:"input_digest" json:"input_digest"`
	Report      QualityReport `yaml:"report" json:"report"`
}

type CheckCollection struct {
	SchemaVersion         int              `yaml:"schema_version" json:"schema_version"`
	ContractVersion       string           `yaml:"contract_version" json:"contract_version"`
	ToolchainPolicyDigest string           `yaml:"toolchain_policy_digest" json:"toolchain_policy_digest"`
	Tasks                 []CheckedTask    `yaml:"tasks" json:"tasks"`
	Execution             ExecutionMetrics `yaml:"execution" json:"execution"`
}

type AggregateTask struct {
	ID                string        `yaml:"id" json:"id"`
	State             State         `yaml:"state" json:"state"`
	InputDigest       string        `yaml:"input_digest" json:"input_digest"`
	TaskType          string        `yaml:"task_type" json:"task_type"`
	Primary           string        `yaml:"primary_skill" json:"primary_skill"`
	TechnicalStatus   string        `yaml:"technical_status" json:"technical_status"`
	ReviewStatus      string        `yaml:"review_status" json:"review_status"`
	PromotionEligible bool          `yaml:"promotion_eligible" json:"promotion_eligible"`
	FailingGateIDs    []string      `yaml:"failing_gate_ids" json:"failing_gate_ids"`
	StarterStatus     string        `yaml:"starter_status" json:"starter_status"`
	ReferenceStatus   string        `yaml:"reference_status" json:"reference_status"`
	RaceStatus        string        `yaml:"race_status" json:"race_status"`
	Mutants           MutantSummary `yaml:"mutants" json:"mutants"`
	Legacy            bool          `yaml:"legacy" json:"legacy"`
}

type StateCount struct {
	State State `yaml:"state" json:"state"`
	Count int   `yaml:"count" json:"count"`
}

type TypeCount struct {
	TaskType string `yaml:"task_type" json:"task_type"`
	Count    int    `yaml:"count" json:"count"`
}

type AggregateReport struct {
	SchemaVersion         int             `yaml:"schema_version" json:"schema_version"`
	SourceCheckContract   string          `yaml:"source_check_contract" json:"source_check_contract"`
	ToolchainPolicyDigest string          `yaml:"toolchain_policy_digest" json:"toolchain_policy_digest"`
	Tasks                 []AggregateTask `yaml:"tasks" json:"tasks"`
	States                []StateCount    `yaml:"states" json:"states"`
	TaskTypes             []TypeCount     `yaml:"task_types" json:"task_types"`
	Gates                 []Gate          `yaml:"gates" json:"gates"`
}
