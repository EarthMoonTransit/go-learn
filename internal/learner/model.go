package learner

import "go-practice-bank/internal/tasking"

const progressSchemaVersion = 3

type Progress struct {
	SchemaVersion int                      `json:"schema_version"`
	Assistance    map[string]bool          `json:"assistance,omitempty"`
	Completions   []CompletionEvidence     `json:"completions,omitempty"`
	CurrentTaskID string                   `json:"current_task_id,omitempty"`
	Tasks         map[string]*TaskProgress `json:"tasks"`
}

type TaskProgress struct {
	TaskID                 string             `json:"task_id"`
	Status                 string             `json:"status"`
	Workspace              string             `json:"workspace"`
	Activity               string             `json:"activity"`
	PrimarySkill           string             `json:"primary_skill"`
	Attempts               int                `json:"attempts"`
	UsedHints              []int              `json:"used_hints"`
	SolutionRevealed       bool               `json:"solution_revealed"`
	SolutionRevealedAt     string             `json:"solution_revealed_at,omitempty"`
	Completed              bool               `json:"completed"`
	CompletedAt            string             `json:"completed_at,omitempty"`
	StartedAt              string             `json:"started_at"`
	LastInteractionAt      string             `json:"last_interaction_at"`
	ActiveSessionStartedAt string             `json:"active_session_started_at,omitempty"`
	ActivitySeconds        int64              `json:"activity_seconds"`
	SkillEncounters        []SkillEncounter   `json:"skill_encounters"`
	Feedback               []FeedbackEntry    `json:"feedback"`
	LastPublicCheck        *PublicCheckRecord `json:"last_public_check,omitempty"`
	LastCompletionSummary  string             `json:"last_completion_summary,omitempty"`
}

type SkillEncounter struct {
	SkillID string `json:"skill_id"`
	Role    string `json:"role"`
	Count   int    `json:"count"`
}

type FeedbackEntry struct {
	RecordedAt string `json:"recorded_at"`
	Text       string `json:"text"`
}

type PublicCheckRecord struct {
	CheckedAt  string `json:"checked_at"`
	Passed     bool   `json:"passed"`
	Summary    string `json:"summary"`
	Diagnostic string `json:"diagnostic,omitempty"`
}

type WorkspaceManifest struct {
	SchemaVersion  int      `json:"schema_version"`
	Generation     string   `json:"generation,omitempty"`
	TaskID         string   `json:"task_id"`
	ContentDigest  string   `json:"content_digest"`
	SubmissionKind string   `json:"submission_kind"`
	EditablePaths  []string `json:"editable_paths"`
	ReadOnlyPaths  []string `json:"read_only_paths"`
}

type SkillMetadata struct {
	ID                   string
	Definition           string
	Prerequisites        []string
	CommonMisconceptions []string
	Mastery              []string
	MinimumPrimary       int
	MinimumTotal         int
}

type SchedulerChoice struct {
	Reason               string
	ReviewSkill          string
	LessonNumber         string
	Resume               bool
	RouteBlock           string
	RouteRole            string
	Preparation          string
	Task                 tasking.LearnerTaskInfo
	MissingPrerequisites []string
	RecentTypeCount      int
	PrimaryDeficit       int
}
