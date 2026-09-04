package tasking

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"go-practice-bank/internal/sources"
	"go-practice-bank/internal/triage"
	"gopkg.in/yaml.v3"
)

type InitOptions struct {
	Slug               string
	PrimarySkill       string
	SourceID           string
	SourcePath         string
	TaskType           string
	TransformationType string
	Actor              string
	Date               string
}

func Init(root string, options InitOptions) (LocatedTask, error) {
	if !slugPattern.MatchString(options.Slug) {
		return LocatedTask{}, fmt.Errorf("invalid slug %q", options.Slug)
	}
	if len(options.Slug) < 4 || len(options.Slug) > 76 {
		return LocatedTask{}, errors.New("slug length must be 4..76 so the generated id satisfies task schema v1")
	}
	if err := cleanRelativePath(options.SourcePath); err != nil {
		return LocatedTask{}, err
	}
	if len(options.SourcePath) > 300 {
		return LocatedTask{}, errors.New("source path exceeds 300 characters")
	}
	if options.TaskType == "" {
		options.TaskType = "implement"
	}
	if !contains([]string{"implement", "debug", "predict", "refactor", "test-writing", "performance"}, options.TaskType) {
		return LocatedTask{}, fmt.Errorf("invalid task type %q", options.TaskType)
	}
	if options.Actor == "" {
		options.Actor = "task-init"
	}
	if len(options.Actor) > 100 {
		return LocatedTask{}, errors.New("actor exceeds 100 characters")
	}
	if options.Date == "" {
		options.Date = time.Now().Format("2006-01-02")
	}
	if _, err := time.Parse("2006-01-02", options.Date); err != nil {
		return LocatedTask{}, fmt.Errorf("invalid date: %w", err)
	}

	registry, err := sources.LoadRegistry(root)
	if err != nil {
		return LocatedTask{}, err
	}
	source, err := registry.Find(options.SourceID)
	if err != nil {
		return LocatedTask{}, err
	}
	if source.UsageMode == "blocked" {
		return LocatedTask{}, fmt.Errorf("source %q is blocked", source.ID)
	}
	if options.TransformationType == "" {
		options.TransformationType = "adapted"
		if source.UsageMode == "reference-only" || source.UsageMode == "concepts-only" {
			options.TransformationType = "metadata-only"
		}
	}
	if !contains([]string{"original", "adapted", "translated-adaptation", "combined", "metadata-only"}, options.TransformationType) {
		return LocatedTask{}, fmt.Errorf("invalid transformation type %q", options.TransformationType)
	}
	if (source.UsageMode == "reference-only" || source.UsageMode == "concepts-only") && options.TransformationType != "metadata-only" {
		return LocatedTask{}, fmt.Errorf("source mode %s permits only metadata-only provenance", source.UsageMode)
	}

	_, contracts, err := triage.LoadSkillContracts(root)
	if err != nil {
		return LocatedTask{}, err
	}
	primary, ok := contracts[options.PrimarySkill]
	if !ok || !primary.Assignable || primary.Status != "active" {
		return LocatedTask{}, fmt.Errorf("primary skill %q is not active and assignable", options.PrimarySkill)
	}

	ids, err := listTaskIDs(root)
	if err != nil {
		return LocatedTask{}, err
	}
	idSet := map[string]bool{}
	for _, id := range ids {
		idSet[id] = true
	}
	id := ""
	for index := 1; index <= 999; index++ {
		candidate := fmt.Sprintf("%s-%03d", options.Slug, index)
		if !idSet[candidate] {
			id = candidate
			break
		}
	}
	if id == "" {
		return LocatedTask{}, fmt.Errorf("all numeric ids for slug %q are occupied", options.Slug)
	}

	drafts := filepath.Join(root, "tasks", stateDirectories[StateDraft])
	if err := os.MkdirAll(drafts, 0o755); err != nil {
		return LocatedTask{}, err
	}
	destination := filepath.Join(drafts, id)
	if _, err := os.Lstat(destination); err == nil {
		return LocatedTask{}, fmt.Errorf("draft %q already exists", id)
	} else if !errors.Is(err, os.ErrNotExist) {
		return LocatedTask{}, err
	}
	staging, err := os.MkdirTemp(drafts, "."+id+"-init-")
	if err != nil {
		return LocatedTask{}, err
	}
	remove := true
	defer func() {
		if remove {
			_ = os.RemoveAll(staging)
		}
	}()
	for _, directory := range []string{"starter", "tests/public", "tests/grader", "solution/reference", "solution/mutants", "review", "hints"} {
		if err := os.MkdirAll(filepath.Join(staging, filepath.FromSlash(directory)), 0o755); err != nil {
			return LocatedTask{}, err
		}
	}

	task := Task{
		SchemaVersion: 1,
		ID:            id,
		Title:         "Draft: " + strings.ReplaceAll(options.Slug, "-", " "),
		TaskType:      options.TaskType,
		Skills:        TaskSkills{Primary: options.PrimarySkill, Secondary: []string{}},
		Prerequisites: append([]string(nil), primary.Prerequisites...),
		NewConcepts:   []string{options.PrimarySkill},
		Difficulty:    "basic",
		EstimatedTimeMinutes: TimeRange{
			Min: 15,
			Max: 30,
		},
		Source: TaskSource{
			ID:         source.ID,
			Repository: source.Repository,
			Path:       options.SourcePath,
			Commit:     source.Commit,
			License:    source.License.Expression,
		},
		TransformationType: options.TransformationType,
		ValidationCommands: []string{
			"go test -run '^$' ./...",
			"! go test ./tests/public",
			"! go test ./tests/grader",
			"go test ./solution/reference",
			"! go test ./solution/mutants/...",
		},
		QualityStatus: "draft",
	}
	sort.Strings(task.Prerequisites)
	lifecycle := Lifecycle{
		SchemaVersion: 1,
		State:         StateDraft,
		UpdatedAt:     options.Date,
		Transitions: []Transition{{
			From: "", To: StateDraft, At: options.Date, Actor: options.Actor, Reason: "initialized authoring package",
		}},
	}
	review := Review{
		SchemaVersion:  1,
		Reviewer:       "unassigned",
		ReviewedAt:     options.Date,
		Decision:       "request-changes",
		MinimumMutants: 1,
		Notes:          "Complete independent pedagogical and provenance review before promotion.",
	}
	provenance := Provenance{
		SchemaVersion: 1,
		Source: ProvenanceSource{
			ID: source.ID, Repository: source.Repository, Path: options.SourcePath,
			Commit: source.Commit, License: source.License.Expression, UsageMode: source.UsageMode,
		},
		UpstreamAuthors:      []string{},
		UpstreamContributors: []string{},
		Transformation: ProvenanceTransform{
			Type:                  options.TransformationType,
			Summary:               "TODO: record the reviewed transformation boundary.",
			IndependentlyAuthored: []string{"TODO: list independently authored artifacts."},
			Excluded:              []string{"TODO: list excluded upstream material."},
		},
		AttributionFile: "ATTRIBUTION.md",
	}
	for _, item := range []struct {
		schema string
		value  any
	}{
		{"task.schema.json", task},
		{"task-lifecycle.schema.json", lifecycle},
		{"task-review.schema.json", review},
		{"task-provenance.schema.json", provenance},
	} {
		if err := validateSchema(root, item.schema, item.value); err != nil {
			return LocatedTask{}, fmt.Errorf("generated scaffold metadata: %w", err)
		}
	}
	for name, value := range map[string]any{
		"task.yaml": task, "lifecycle.yaml": lifecycle, "review/review.yaml": review, "provenance.yaml": provenance,
	} {
		content, err := yaml.Marshal(value)
		if err != nil {
			return LocatedTask{}, err
		}
		if err := os.WriteFile(filepath.Join(staging, filepath.FromSlash(name)), content, 0o644); err != nil {
			return LocatedTask{}, err
		}
	}
	replacements := map[string]string{
		"{{ID}}": id, "{{SOURCE_REPOSITORY}}": source.Repository,
		"{{SOURCE_COMMIT}}": source.Commit, "{{SOURCE_LICENSE}}": source.License.Expression,
	}
	textFiles := map[string]string{
		"go.mod":          "go.mod.tmpl",
		"TASK.md":         "TASK.md.tmpl",
		"review/notes.md": "review/notes.md.tmpl",
		"hints/HINTS.md":  "hints/HINTS.md.tmpl",
		"ATTRIBUTION.md":  "ATTRIBUTION.md.tmpl",
	}
	for name, templateName := range textFiles {
		content, err := os.ReadFile(filepath.Join(root, "tasks", "_template", filepath.FromSlash(templateName)))
		if err != nil {
			return LocatedTask{}, fmt.Errorf("read authoring template %s: %w", templateName, err)
		}
		text := string(content)
		for old, replacement := range replacements {
			text = strings.ReplaceAll(text, old, replacement)
		}
		if err := os.WriteFile(filepath.Join(staging, filepath.FromSlash(name)), []byte(text), 0o644); err != nil {
			return LocatedTask{}, err
		}
	}
	if err := os.Rename(staging, destination); err != nil {
		return LocatedTask{}, fmt.Errorf("publish draft: %w", err)
	}
	remove = false
	return LocatedTask{ID: id, State: StateDraft, Path: destination}, nil
}
