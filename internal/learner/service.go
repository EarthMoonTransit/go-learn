package learner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

var learnerTaskIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*-[0-9]{3}$`)

const learnerWorkspaceMarker = ".go-learn-workspace.json"

type Config struct {
	Root string
	Home string
	Now  func() time.Time
}

type Service struct {
	root string
	home string
	now  func() time.Time
}

func New(config Config) (*Service, error) {
	if config.Root == "" {
		config.Root = "."
	}
	if config.Home == "" {
		var err error
		config.Home, err = DefaultHome()
		if err != nil {
			return nil, err
		}
	}
	root, err := filepath.Abs(config.Root)
	if err != nil {
		return nil, err
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return nil, err
	}
	info, err := os.Lstat(root)
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return nil, errors.New("repository root must be a real directory")
	}
	home, err := prepareDirectory(config.Home)
	if err != nil {
		return nil, err
	}
	now := config.Now
	if now == nil {
		now = time.Now
	}
	return &Service{root: root, home: home, now: now}, nil
}

func (s *Service) Home() string { return s.home }

func (s *Service) ListTasks() ([]tasking.LearnerTaskInfo, Progress, error) {
	tasks, err := tasking.ListReadyLearnerTasks(s.root)
	if err != nil {
		return nil, Progress{}, err
	}
	progress, err := loadProgress(s.home)
	return tasks, progress, err
}

func (s *Service) Show(id string) (tasking.LearnerTaskInfo, error) {
	return tasking.LoadReadyLearnerTask(s.root, id)
}

func (s *Service) Start(id string) (string, error) {
	info, err := tasking.LoadReadyLearnerTask(s.root, id)
	if err != nil {
		return "", err
	}
	files, err := tasking.LearnerWorkspaceFiles(s.root, id)
	if err != nil {
		return "", err
	}
	workspaceRoot, err := preparePrivateSubdirectory(s.home, "workspaces")
	if err != nil {
		return "", err
	}
	target := filepath.Join(workspaceRoot, id)
	now := s.timestamp()
	createdTarget := false
	err = withProgressLock(s.home, func(progress *Progress) error {
		if _, exists := progress.Tasks[id]; exists {
			return fmt.Errorf("task %q already has progress; use resume or reset-task", id)
		}
		if _, err := os.Lstat(target); err == nil {
			return fmt.Errorf("learner workspace already exists: %s", target)
		} else if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		staging, err := os.MkdirTemp(workspaceRoot, "."+id+"-*")
		if err != nil {
			return err
		}
		published := false
		defer func() {
			if !published {
				_ = os.RemoveAll(staging)
			}
		}()
		manifest := WorkspaceManifest{
			SchemaVersion: 1, TaskID: id, ContentDigest: info.ContentDigest,
			SubmissionKind: info.SubmissionKind,
		}
		for _, file := range files {
			mode := os.FileMode(0o400)
			if file.Editable {
				mode = 0o600
				manifest.EditablePaths = append(manifest.EditablePaths, file.Path)
			} else {
				manifest.ReadOnlyPaths = append(manifest.ReadOnlyPaths, file.Path)
			}
			if err := writePrivateFile(filepath.Join(staging, filepath.FromSlash(file.Path)), file.Content, mode); err != nil {
				return err
			}
		}
		sort.Strings(manifest.EditablePaths)
		sort.Strings(manifest.ReadOnlyPaths)
		marker, err := json.MarshalIndent(manifest, "", "  ")
		if err != nil {
			return err
		}
		marker = append(marker, '\n')
		if err := writePrivateFile(filepath.Join(staging, learnerWorkspaceMarker), marker, 0o400); err != nil {
			return err
		}
		if err := os.Rename(staging, target); err != nil {
			return err
		}
		published = true
		createdTarget = true
		pauseOtherSessions(progress, id, now)
		progress.Tasks[id] = &TaskProgress{
			TaskID: id, Status: "in-progress", Workspace: filepath.ToSlash(filepath.Join("workspaces", id)),
			Activity: info.Activity, PrimarySkill: info.PrimarySkill,
			UsedHints: []int{}, SkillEncounters: []SkillEncounter{}, Feedback: []FeedbackEntry{},
			StartedAt: now, LastInteractionAt: now, ActiveSessionStartedAt: now,
		}
		progress.CurrentTaskID = id
		return nil
	})
	if err != nil {
		if createdTarget {
			if cleanupErr := os.RemoveAll(target); cleanupErr != nil {
				return "", errors.Join(err, fmt.Errorf("rollback learner workspace: %w", cleanupErr))
			}
		}
		return "", err
	}
	return target, nil
}

func (s *Service) Resume(id string) (string, error) {
	info, err := tasking.LoadReadyLearnerTask(s.root, id)
	if err != nil {
		return "", err
	}
	now := s.timestamp()
	var workspace string
	err = withProgressLock(s.home, func(progress *Progress) error {
		task, ok := progress.Tasks[id]
		if !ok {
			return fmt.Errorf("task %q has not been started", id)
		}
		workspace, err = s.validateWorkspace(task, info)
		if err != nil {
			return err
		}
		pauseOtherSessions(progress, id, now)
		if !task.Completed && task.ActiveSessionStartedAt == "" {
			task.ActiveSessionStartedAt = now
		}
		task.LastInteractionAt = now
		progress.CurrentTaskID = id
		return nil
	})
	return workspace, err
}

func (s *Service) Pause() error {
	now := s.timestamp()
	return withProgressLock(s.home, func(progress *Progress) error {
		if progress.CurrentTaskID == "" {
			return errors.New("no current task; use start or resume")
		}
		task, err := requireSameCurrent(progress, progress.CurrentTaskID)
		if err != nil {
			return err
		}
		endActiveSession(task, now)
		task.LastInteractionAt = now
		return nil
	})
}

func (s *Service) Check() (tasking.LearnerVerificationResult, error) {
	progress, _, info, workspace, err := s.currentTask()
	if err != nil {
		return tasking.LearnerVerificationResult{}, err
	}
	verificationStarted := s.timestamp()
	result, err := tasking.VerifyLearnerWorkspace(s.root, info.ID, workspace, tasking.LearnerPublicChecks)
	if err != nil {
		return tasking.LearnerVerificationResult{}, err
	}
	now := s.timestamp()
	err = withProgressLock(s.home, func(current *Progress) error {
		stored, err := requireSameCurrent(current, progress.CurrentTaskID)
		if err != nil {
			return err
		}
		stored.Attempts++
		finishVerificationSession(stored, verificationStarted, now, !stored.Completed)
		stored.LastInteractionAt = now
		stored.LastPublicCheck = &PublicCheckRecord{
			CheckedAt: now, Passed: result.Passed, Summary: result.Summary, Diagnostic: result.Diagnostic,
		}
		return nil
	})
	return result, err
}

func (s *Service) Complete() (tasking.LearnerVerificationResult, error) {
	progress, task, info, workspace, err := s.currentTask()
	if err != nil {
		return tasking.LearnerVerificationResult{}, err
	}
	if task.Completed {
		return tasking.LearnerVerificationResult{}, errors.New("task is already completed")
	}
	verificationStarted := s.timestamp()
	result, err := tasking.VerifyLearnerWorkspace(s.root, info.ID, workspace, tasking.LearnerFullChecks)
	if err != nil {
		return tasking.LearnerVerificationResult{}, err
	}
	now := s.timestamp()
	err = withProgressLock(s.home, func(current *Progress) error {
		stored, err := requireSameCurrent(current, progress.CurrentTaskID)
		if err != nil {
			return err
		}
		stored.Attempts++
		finishVerificationSession(stored, verificationStarted, now, !result.Passed && !stored.Completed)
		stored.LastInteractionAt = now
		stored.LastCompletionSummary = result.Summary
		if result.Passed && !stored.Completed {
			stored.Completed = true
			stored.Status = "completed"
			stored.CompletedAt = now
			stored.SkillEncounters = []SkillEncounter{{SkillID: info.PrimarySkill, Role: "primary", Count: 1}}
			for _, id := range info.SecondarySkills {
				stored.SkillEncounters = append(stored.SkillEncounters, SkillEncounter{SkillID: id, Role: "secondary", Count: 1})
			}
		}
		return nil
	})
	return result, err
}

func (s *Service) Hint() (int, string, error) {
	progress, task, info, _, err := s.currentTask()
	if err != nil {
		return 0, "", err
	}
	hints, err := tasking.LearnerHints(s.root, info.ID)
	if err != nil {
		return 0, "", err
	}
	used := map[int]bool{}
	for _, index := range task.UsedHints {
		used[index] = true
	}
	next := 0
	for index := range hints {
		if !used[index+1] {
			next = index + 1
			break
		}
	}
	if next == 0 {
		return 0, "", errors.New("all hints have already been used")
	}
	now := s.timestamp()
	err = withProgressLock(s.home, func(current *Progress) error {
		stored, err := requireSameCurrent(current, progress.CurrentTaskID)
		if err != nil {
			return err
		}
		stored.UsedHints = append(stored.UsedHints, next)
		sort.Ints(stored.UsedHints)
		stored.LastInteractionAt = now
		return nil
	})
	return next, hints[next-1], err
}

func (s *Service) Solution(reveal bool) ([]tasking.LearnerWorkspaceFile, error) {
	if !reveal {
		return nil, errors.New("solution reveal requires the explicit --reveal flag")
	}
	progress, _, info, _, err := s.currentTask()
	if err != nil {
		return nil, err
	}
	files, err := tasking.LearnerSolutionFiles(s.root, info.ID)
	if err != nil {
		return nil, err
	}
	now := s.timestamp()
	err = withProgressLock(s.home, func(current *Progress) error {
		stored, err := requireSameCurrent(current, progress.CurrentTaskID)
		if err != nil {
			return err
		}
		stored.SolutionRevealed = true
		stored.SolutionRevealedAt = now
		stored.LastInteractionAt = now
		return nil
	})
	return files, err
}

func (s *Service) Review() (string, error) {
	_, task, info, workspace, err := s.currentTask()
	if err != nil {
		return "", err
	}
	if !task.Completed {
		return "", errors.New("review is available only after successful completion")
	}
	skills, err := loadSkillMetadata(s.root)
	if err != nil {
		return "", err
	}
	files, err := tasking.LearnerSubmissionFiles(s.root, info.ID, workspace)
	if err != nil {
		return "", err
	}
	var output strings.Builder
	fmt.Fprintf(&output, "# Review: %s\n\n", info.Title)
	fmt.Fprintf(&output, "Status: completed\nAttempts: %d\nActive duration: %s\nSolution revealed: %t\n\n", task.Attempts, formatDuration(task.ActivitySeconds), task.SolutionRevealed)
	writeSkillReview(&output, "Primary skill", skills[info.PrimarySkill])
	for _, id := range info.SecondarySkills {
		writeSkillReview(&output, "Secondary skill", skills[id])
	}
	output.WriteString("## Learner files\n\n")
	for _, file := range files {
		fmt.Fprintf(&output, "- `%s`\n", file.Path)
	}
	return output.String(), nil
}

func (s *Service) Next() (SchedulerChoice, error) {
	tasks, progress, err := s.ListTasks()
	if err != nil {
		return SchedulerChoice{}, err
	}
	skills, err := loadSkillMetadata(s.root)
	if err != nil {
		return SchedulerChoice{}, err
	}
	known := map[string]bool{}
	primaryCounts := map[string]int{}
	for _, item := range progress.Tasks {
		if !item.Completed {
			continue
		}
		for _, encounter := range item.SkillEncounters {
			known[encounter.SkillID] = true
			if encounter.Role == "primary" {
				primaryCounts[encounter.SkillID] += encounter.Count
			}
		}
	}
	recent := recentActivities(progress, 3)
	var choices []SchedulerChoice
	for _, task := range tasks {
		if item := progress.Tasks[task.ID]; item != nil && item.Completed {
			continue
		}
		choice := SchedulerChoice{Task: task, RecentTypeCount: recent[task.Activity]}
		for _, prerequisite := range task.Prerequisites {
			if !known[prerequisite] {
				choice.MissingPrerequisites = append(choice.MissingPrerequisites, prerequisite)
			}
		}
		sort.Strings(choice.MissingPrerequisites)
		choice.PrimaryDeficit = skills[task.PrimarySkill].MinimumPrimary - primaryCounts[task.PrimarySkill]
		if choice.PrimaryDeficit < 0 {
			choice.PrimaryDeficit = 0
		}
		choices = append(choices, choice)
	}
	if len(choices) == 0 {
		return SchedulerChoice{}, errors.New("all ready Task Contract v2 tasks are completed")
	}
	sort.Slice(choices, func(i, j int) bool {
		left, right := choices[i], choices[j]
		leftEligible, rightEligible := len(left.MissingPrerequisites) == 0, len(right.MissingPrerequisites) == 0
		if leftEligible != rightEligible {
			return leftEligible
		}
		if len(left.MissingPrerequisites) != len(right.MissingPrerequisites) {
			return len(left.MissingPrerequisites) < len(right.MissingPrerequisites)
		}
		if difficultyRank(left.Task.Difficulty) != difficultyRank(right.Task.Difficulty) {
			return difficultyRank(left.Task.Difficulty) < difficultyRank(right.Task.Difficulty)
		}
		if left.RecentTypeCount != right.RecentTypeCount {
			return left.RecentTypeCount < right.RecentTypeCount
		}
		if left.PrimaryDeficit != right.PrimaryDeficit {
			return left.PrimaryDeficit > right.PrimaryDeficit
		}
		return left.Task.ID < right.Task.ID
	})
	return choices[0], nil
}

func (s *Service) Progress() (Progress, error) {
	progress, err := loadProgress(s.home)
	if err != nil {
		return Progress{}, err
	}
	// Show elapsed work in the snapshot without persisting or closing the session.
	now := s.timestamp()
	for _, task := range progress.Tasks {
		started := task.ActiveSessionStartedAt
		endActiveSession(task, now)
		task.ActiveSessionStartedAt = started
	}
	return progress, nil
}

func (s *Service) Feedback(text string) error {
	text = strings.TrimSpace(text)
	if text == "" || len([]rune(text)) > 2000 || strings.ContainsRune(text, 0) {
		return errors.New("feedback must contain 1..2000 characters")
	}
	progress, _, _, _, err := s.currentTask()
	if err != nil {
		return err
	}
	now := s.timestamp()
	return withProgressLock(s.home, func(current *Progress) error {
		stored, err := requireSameCurrent(current, progress.CurrentTaskID)
		if err != nil {
			return err
		}
		stored.Feedback = append(stored.Feedback, FeedbackEntry{RecordedAt: now, Text: text})
		stored.LastInteractionAt = now
		return nil
	})
}

func (s *Service) ResetTask(id string, confirmed bool) error {
	if !confirmed {
		return errors.New("reset-task requires the explicit --yes flag")
	}
	if id != "" && !learnerTaskIDPattern.MatchString(id) {
		return fmt.Errorf("invalid task id %q", id)
	}
	return withProgressLock(s.home, func(progress *Progress) error {
		if id == "" {
			id = progress.CurrentTaskID
		}
		if id == "" || !learnerTaskIDPattern.MatchString(id) {
			return errors.New("no current task to reset")
		}
		if _, ok := progress.Tasks[id]; !ok {
			return fmt.Errorf("task %q has no local progress", id)
		}
		if err := removeLearnerWorkspace(s.home, id); err != nil {
			return err
		}
		delete(progress.Tasks, id)
		if progress.CurrentTaskID == id {
			progress.CurrentTaskID = ""
		}
		return nil
	})
}

func (s *Service) ExportReviewContext() (string, error) {
	_, task, info, workspace, err := s.currentTask()
	if err != nil {
		return "", err
	}
	files, err := tasking.LearnerSubmissionFiles(s.root, info.ID, workspace)
	if err != nil {
		return "", err
	}
	skills, err := loadSkillMetadata(s.root)
	if err != nil {
		return "", err
	}
	hints, err := tasking.LearnerHints(s.root, info.ID)
	if err != nil {
		return "", err
	}
	content := renderReviewContext(info, task, files, skills, hints)
	exportRoot, err := preparePrivateSubdirectory(s.home, "exports")
	if err != nil {
		return "", err
	}
	path := filepath.Join(exportRoot, info.ID+".md")
	if err := atomicPrivateWrite(path, []byte(content)); err != nil {
		return "", err
	}
	return path, nil
}

func (s *Service) currentTask() (Progress, *TaskProgress, tasking.LearnerTaskInfo, string, error) {
	progress, err := loadProgress(s.home)
	if err != nil {
		return Progress{}, nil, tasking.LearnerTaskInfo{}, "", err
	}
	if progress.CurrentTaskID == "" {
		return Progress{}, nil, tasking.LearnerTaskInfo{}, "", errors.New("no current task; use start or resume")
	}
	task := progress.Tasks[progress.CurrentTaskID]
	if task == nil {
		return Progress{}, nil, tasking.LearnerTaskInfo{}, "", errors.New("current task progress is missing")
	}
	info, err := tasking.LoadReadyLearnerTask(s.root, task.TaskID)
	if err != nil {
		return Progress{}, nil, tasking.LearnerTaskInfo{}, "", err
	}
	workspace, err := s.validateWorkspace(task, info)
	return progress, task, info, workspace, err
}

func (s *Service) validateWorkspace(progress *TaskProgress, info tasking.LearnerTaskInfo) (string, error) {
	if progress.Workspace != filepath.ToSlash(filepath.Join("workspaces", progress.TaskID)) {
		return "", errors.New("progress contains an invalid workspace path")
	}
	workspace := filepath.Join(s.home, filepath.FromSlash(progress.Workspace))
	workspaceRoot := filepath.Join(s.home, "workspaces")
	relative, err := filepath.Rel(workspaceRoot, workspace)
	if err != nil || relative != progress.TaskID {
		return "", errors.New("learner workspace escapes the local workspace root")
	}
	workspaceInfo, err := os.Lstat(workspace)
	if err != nil || workspaceInfo.Mode()&os.ModeSymlink != 0 || !workspaceInfo.IsDir() {
		return "", errors.New("learner workspace must be a real directory")
	}
	resolved, err := filepath.EvalSymlinks(workspace)
	if err != nil || resolved != workspace {
		return "", errors.New("learner workspace contains a symlink escape")
	}
	content, err := readPrivateRegular(filepath.Join(workspace, learnerWorkspaceMarker), 64<<10)
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	var manifest WorkspaceManifest
	if err := decoder.Decode(&manifest); err != nil {
		return "", err
	}
	if manifest.SchemaVersion != 1 || manifest.TaskID != progress.TaskID || manifest.ContentDigest != info.ContentDigest || manifest.SubmissionKind != info.SubmissionKind {
		return "", errors.New("learner workspace is stale or belongs to another task; reset and start again")
	}
	return workspace, nil
}

func (s *Service) timestamp() string { return s.now().UTC().Truncate(time.Second).Format(time.RFC3339) }

func requireSameCurrent(progress *Progress, id string) (*TaskProgress, error) {
	if progress.CurrentTaskID != id {
		return nil, errors.New("current task changed while the command was running")
	}
	task := progress.Tasks[id]
	if task == nil {
		return nil, errors.New("current task progress disappeared")
	}
	return task, nil
}

func endActiveSession(progress *TaskProgress, now string) {
	if progress.ActiveSessionStartedAt == "" {
		return
	}
	started, startErr := time.Parse(time.RFC3339, progress.ActiveSessionStartedAt)
	ended, endErr := time.Parse(time.RFC3339, now)
	if startErr == nil && endErr == nil && !ended.Before(started) {
		progress.ActivitySeconds += int64(ended.Sub(started) / time.Second)
	}
	progress.ActiveSessionStartedAt = ""
}

func pauseOtherSessions(progress *Progress, selectedID, now string) {
	for id, task := range progress.Tasks {
		if id != selectedID {
			endActiveSession(task, now)
		}
	}
}

func finishVerificationSession(progress *TaskProgress, started, ended string, continueWorking bool) {
	wasActive := progress.ActiveSessionStartedAt != ""
	endActiveSession(progress, started)
	if wasActive && continueWorking {
		progress.ActiveSessionStartedAt = ended
	}
}

func loadSkillMetadata(root string) (map[string]SkillMetadata, error) {
	content, err := os.ReadFile(filepath.Join(root, "curriculum", "skills.yaml"))
	if err != nil {
		return nil, err
	}
	var document struct {
		Domains []struct {
			Skills []struct {
				ID                   string   `yaml:"id"`
				Definition           string   `yaml:"definition"`
				Prerequisites        []string `yaml:"prerequisites"`
				CommonMisconceptions []string `yaml:"common_misconceptions"`
				Mastery              []string `yaml:"mastery"`
				MinimumOccurrences   struct {
					Primary int `yaml:"primary"`
					Total   int `yaml:"total"`
				} `yaml:"minimum_occurrences"`
			} `yaml:"skills"`
		} `yaml:"domains"`
	}
	if err := yaml.Unmarshal(content, &document); err != nil {
		return nil, err
	}
	result := map[string]SkillMetadata{}
	for _, domain := range document.Domains {
		for _, skill := range domain.Skills {
			result[skill.ID] = SkillMetadata{
				ID: skill.ID, Definition: skill.Definition,
				Prerequisites:        append([]string(nil), skill.Prerequisites...),
				CommonMisconceptions: append([]string(nil), skill.CommonMisconceptions...),
				Mastery:              append([]string(nil), skill.Mastery...),
				MinimumPrimary:       skill.MinimumOccurrences.Primary, MinimumTotal: skill.MinimumOccurrences.Total,
			}
		}
	}
	return result, nil
}

func recentActivities(progress Progress, limit int) map[string]int {
	var items []*TaskProgress
	for _, item := range progress.Tasks {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].LastInteractionAt != items[j].LastInteractionAt {
			return items[i].LastInteractionAt > items[j].LastInteractionAt
		}
		return items[i].TaskID < items[j].TaskID
	})
	if len(items) > limit {
		items = items[:limit]
	}
	result := map[string]int{}
	for _, item := range items {
		result[item.Activity]++
	}
	return result
}

func difficultyRank(value string) int {
	switch value {
	case "introductory":
		return 0
	case "basic":
		return 1
	case "intermediate":
		return 2
	case "advanced":
		return 3
	default:
		return 4
	}
}
