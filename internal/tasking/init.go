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
	"go-practice-bank/internal/toolchainpolicy"
	"go-practice-bank/internal/triage"
	"gopkg.in/yaml.v3"
)

type InitOptions struct {
	SchemaVersion       int
	Slug                string
	PrimarySkill        string
	SourceID            string
	SourcePath          string
	TaskType            string
	TransformationType  string
	Actor               string
	Date                string
	Activity            string
	SubmissionKind      string
	VerificationProfile string
}

func Init(root string, options InitOptions) (LocatedTask, error) {
	if options.SchemaVersion == 0 {
		options.SchemaVersion = 2
	}
	if options.SchemaVersion == 2 {
		return initV2(root, options)
	}
	if options.SchemaVersion != 1 {
		return LocatedTask{}, fmt.Errorf("unsupported task schema version %d", options.SchemaVersion)
	}
	toolchain, err := toolchainpolicy.Load(root)
	if err != nil {
		return LocatedTask{}, err
	}
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
		"{{GO_LANGUAGE_VERSION}}": toolchain.GoLanguageVersion,
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

func initV2(root string, options InitOptions) (LocatedTask, error) {
	toolchain, err := toolchainpolicy.Load(root)
	if err != nil {
		return LocatedTask{}, err
	}
	identity, err := loadProjectIdentity(root)
	if err != nil {
		return LocatedTask{}, err
	}
	if !slugPattern.MatchString(options.Slug) || len(options.Slug) < 4 || len(options.Slug) > 76 {
		return LocatedTask{}, fmt.Errorf("invalid v2 task slug %q", options.Slug)
	}
	if options.Actor == "" {
		options.Actor = "task-init"
	}
	if options.Date == "" {
		options.Date = time.Now().Format("2006-01-02")
	}
	if _, err := time.Parse("2006-01-02", options.Date); err != nil {
		return LocatedTask{}, fmt.Errorf("invalid date: %w", err)
	}
	if len(options.Actor) > 100 {
		return LocatedTask{}, errors.New("actor exceeds 100 characters")
	}
	if options.VerificationProfile == "" {
		options.VerificationProfile = "fixed-tests"
	}
	if options.Activity == "" {
		options.Activity = map[string]string{
			"fixed-tests": "implementation", "mutation-tests": "test-writing", "workspace-tests": "integration",
			"fuzz-seed-tests": "fuzz-test-writing", "benchmark-contract": "performance",
		}[options.VerificationProfile]
	}
	if options.SubmissionKind == "" {
		options.SubmissionKind = map[string]string{
			"fixed-tests": "implementation", "mutation-tests": "tests", "workspace-tests": "workspace-patch",
			"fuzz-seed-tests": "tests", "benchmark-contract": "implementation",
		}[options.VerificationProfile]
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

	submission := Submission{Kind: options.SubmissionKind, EditablePaths: []string{}, WorkspaceSkeleton: []string{}}
	if options.VerificationProfile == "workspace-tests" {
		submission.EditablePaths = []string{"subject/implementation.go"}
		submission.WorkspaceSkeleton = []string{"subject/api.go"}
	}
	task := TaskV2{
		SchemaVersion: 2, ID: id, Title: "Draft: " + strings.ReplaceAll(options.Slug, "-", " "),
		Activity: options.Activity, Submission: submission,
		Skills:        TaskSkills{Primary: options.PrimarySkill, Secondary: []string{}},
		Prerequisites: append([]string(nil), primary.Prerequisites...), NewConcepts: []string{options.PrimarySkill},
		Difficulty: "basic", EstimatedTimeMinutes: TimeRange{Min: 15, Max: 30},
		Verification: VerificationPlan{
			Profile: options.VerificationProfile, Packages: []string{"./subject"}, PublicChecks: []string{"./subject"},
			GraderChecks: []string{"./subject"}, ReferenceChecks: []string{"./subject"}, RequiredMutants: []string{"required-behavior"},
			Limits: VerificationLimits{CommandTimeoutSeconds: 30, DeterminismRuns: 2, FuzzSeconds: 0, BenchmarkRuns: 1},
		},
		QualityStatus: "draft",
	}
	sort.Strings(task.Prerequisites)
	if options.VerificationProfile == "fuzz-seed-tests" {
		task.Verification.Fuzz = &FuzzPlan{Grammar: "hex-codec-v1", Target: "FuzzHex", CorrectVariants: []string{"loop", "empty"}}
	}
	if err := validateSchema(root, "task-v2.schema.json", task); err != nil {
		return LocatedTask{}, fmt.Errorf("generated v2 task metadata: %w", err)
	}
	lifecycle := Lifecycle{SchemaVersion: 1, State: StateDraft, UpdatedAt: options.Date, Transitions: []Transition{{From: "", To: StateDraft, At: options.Date, Actor: options.Actor, Reason: "initialized v2 authoring package"}}}
	review := Review{SchemaVersion: 1, Reviewer: "unassigned", ReviewedAt: options.Date, Decision: "request-changes", MinimumMutants: 1, Notes: "Complete independent pedagogical and provenance review before promotion."}
	for _, item := range []struct {
		schema string
		value  any
	}{{"task-lifecycle.schema.json", lifecycle}, {"task-review.schema.json", review}} {
		if err := validateSchema(root, item.schema, item.value); err != nil {
			return LocatedTask{}, fmt.Errorf("generated scaffold metadata: %w", err)
		}
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
	directories := []string{"starter", "solution/reference", "review", "hints"}
	switch options.VerificationProfile {
	case "fixed-tests":
		directories = append(directories, "tests/public", "tests/grader", "solution/mutants/required-behavior")
	case "mutation-tests":
		directories = append(directories, "verification/reference", "verification/mutants/required-behavior")
	case "fuzz-seed-tests":
		directories = append(directories, "verification/reference", "verification/mutants/required-behavior", "verification/correct/loop", "verification/correct/empty")
	case "workspace-tests":
		directories = append(directories, "workspace/subject", "tests/public/subject", "tests/grader/subject", "solution/mutants/required-behavior/subject", "starter/subject", "solution/reference/subject")
	default:
		directories = append(directories, "verification")
	}
	for _, directory := range directories {
		if err := os.MkdirAll(filepath.Join(staging, filepath.FromSlash(directory)), 0o755); err != nil {
			return LocatedTask{}, err
		}
	}
	for name, value := range map[string]any{"task.yaml": task, "lifecycle.yaml": lifecycle, "review/review.yaml": review} {
		content, err := yaml.Marshal(value)
		if err != nil {
			return LocatedTask{}, err
		}
		if err := os.WriteFile(filepath.Join(staging, filepath.FromSlash(name)), content, 0o644); err != nil {
			return LocatedTask{}, err
		}
	}
	replacements := map[string]string{"{{ID}}": id, "{{SOURCE_REPOSITORY}}": identity.Repository, "{{SOURCE_COMMIT}}": "deferred to release manifest", "{{SOURCE_LICENSE}}": identity.License, "{{GO_LANGUAGE_VERSION}}": toolchain.GoLanguageVersion}
	for name, templateName := range map[string]string{"go.mod": "go.mod.tmpl", "TASK.md": "TASK.md.tmpl", "review/notes.md": "review/notes.md.tmpl", "hints/HINTS.md": "hints/HINTS.md.tmpl", "ATTRIBUTION.md": "ATTRIBUTION.md.tmpl"} {
		content, err := os.ReadFile(filepath.Join(root, "tasks", "_template", filepath.FromSlash(templateName)))
		if err != nil {
			return LocatedTask{}, err
		}
		text := string(content)
		for old, replacement := range replacements {
			text = strings.ReplaceAll(text, old, replacement)
		}
		if err := os.WriteFile(filepath.Join(staging, filepath.FromSlash(name)), []byte(text), 0o644); err != nil {
			return LocatedTask{}, err
		}
	}
	if options.VerificationProfile == "fuzz-seed-tests" {
		contract, err := readRegularFile(filepath.Join(root, filepath.FromSlash(fuzzCanonicalContractPath)), 8<<20)
		if err != nil {
			return LocatedTask{}, fmt.Errorf("finite source contract: %w", err)
		}
		if err := os.WriteFile(filepath.Join(staging, filepath.FromSlash(fuzzSourceContractPath)), contract, 0644); err != nil {
			return LocatedTask{}, err
		}
		promptPath := filepath.Join(staging, "TASK.md")
		prompt, err := os.ReadFile(promptPath)
		if err != nil {
			return LocatedTask{}, err
		}
		prompt = append(prompt, []byte("\n## Finite native source contract\n\nRead the [complete source contract](FUZZ_SOURCE.md) before editing the native target.\n")...)
		if err := os.WriteFile(promptPath, prompt, 0644); err != nil {
			return LocatedTask{}, err
		}
	}
	digest, err := TaskContentDigest(staging)
	if err != nil {
		return LocatedTask{}, err
	}
	provenance := ProvenanceV2{
		SchemaVersion: 2, Kind: "original",
		Project:  ProvenanceProject{Identity: identity.ID, Repository: identity.Repository, License: identity.License, ContentDigest: digest},
		Creation: ProvenanceCreation{Author: options.Actor, CreatedAt: options.Date, Summary: "Original authoring scaffold created without external task content."},
		Sources:  []ProvenanceSourceV2{}, AttributionFile: "ATTRIBUTION.md",
		Release: ProvenanceRelease{Binding: "deferred-to-release-manifest", Manifest: filepath.ToSlash(filepath.Join(identity.ReleaseManifestDirectory, id+".yaml"))},
	}
	if options.SourceID != "" {
		if err := cleanRelativePath(options.SourcePath); err != nil {
			return LocatedTask{}, fmt.Errorf("source path: %w", err)
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
		role, usage := "adapted-code", "adapted"
		if source.UsageMode == "reference-only" {
			role, usage = "factual-reference", "factual-reference"
		} else if source.UsageMode == "concepts-only" {
			role, usage = "concept-seed", "concept-only"
		}
		provenance.Kind = "adapted"
		provenance.Creation.Summary = "Authoring scaffold created with one audited external source reference; artifact digest requires author review."
		provenance.Sources = []ProvenanceSourceV2{{RegistryID: source.ID, Repository: source.Repository, Commit: source.Commit, License: source.License.Expression, UsageMode: source.UsageMode, Role: role, Artifacts: []ProvenanceArtifactV2{{Path: options.SourcePath, SHA256: strings.Repeat("0", 64), Usage: usage}}}}
	}
	if err := validateSchema(root, "task-provenance-v2.schema.json", provenance); err != nil {
		return LocatedTask{}, fmt.Errorf("generated v2 provenance: %w", err)
	}
	content, err := yaml.Marshal(provenance)
	if err != nil {
		return LocatedTask{}, err
	}
	if err := os.WriteFile(filepath.Join(staging, "provenance.yaml"), content, 0o644); err != nil {
		return LocatedTask{}, err
	}
	if err := os.Rename(staging, destination); err != nil {
		return LocatedTask{}, fmt.Errorf("publish draft: %w", err)
	}
	remove = false
	return LocatedTask{ID: id, State: StateDraft, Path: destination}, nil
}
