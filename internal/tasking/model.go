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
	SchemaVersion int           `yaml:"schema_version" json:"schema_version"`
	TaskID        string        `yaml:"task_id" json:"task_id"`
	State         State         `yaml:"state" json:"state"`
	Passed        bool          `yaml:"passed" json:"passed"`
	Gates         []Gate        `yaml:"gates" json:"gates"`
	Mutants       MutantSummary `yaml:"mutants" json:"mutants"`
}

type AggregateTask struct {
	ID       string `yaml:"id" json:"id"`
	State    State  `yaml:"state" json:"state"`
	TaskType string `yaml:"task_type" json:"task_type"`
	Primary  string `yaml:"primary_skill" json:"primary_skill"`
	Legacy   bool   `yaml:"legacy" json:"legacy"`
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
	SchemaVersion int             `yaml:"schema_version" json:"schema_version"`
	Tasks         []AggregateTask `yaml:"tasks" json:"tasks"`
	States        []StateCount    `yaml:"states" json:"states"`
	TaskTypes     []TypeCount     `yaml:"task_types" json:"task_types"`
	Gates         []Gate          `yaml:"gates" json:"gates"`
}
