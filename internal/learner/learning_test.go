package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/tasking"
)

func evidence(id, skill, role string, assisted bool) CompletionEvidence {
	return CompletionEvidence{TaskID: id, Skills: []string{skill}, RouteRole: role, Assisted: assisted}
}
func TestLearningDistinguishesIndependentPracticeFromProgress(t *testing.T) {
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}, Completions: []CompletionEvidence{
		evidence("intro-001", "skill", "introduction", false),
		evidence("helped-001", "skill", "practice", true),
		evidence("first-001", "skill", "practice", false),
		evidence("first-001", "skill", "practice", false),
	}}
	v := learningView(p)
	if v.Mastered != 0 || v.Skills[0].Independent != 1 {
		t.Fatal(v)
	}
	p.Completions = append(p.Completions, evidence("second-001", "skill", "transfer", false))
	if v = learningView(p); v.Mastered != 1 || v.Skills[0].Independent != 2 {
		t.Fatal(v)
	}
	// Legacy records have no trustworthy assistance-at-completion snapshot.
	p = Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{"old-001": {Completed: true, PrimarySkill: "skill"}, "old-002": {Completed: true, PrimarySkill: "skill"}}}
	if v = learningView(p); v.Mastered != 0 || v.Skills[0].Independent != 0 {
		t.Fatal(v)
	}
}
func TestReviewWaitsForThreeDistinctOtherTasks(t *testing.T) {
	p := Progress{SchemaVersion: 3, Completions: []CompletionEvidence{evidence("base-001", "skill", "introduction", false)}}
	for i, id := range []string{"other-a-001", "other-b-001", "other-c-001"} {
		if v := learningView(p); v.Skills[0].ReviewDue {
			t.Fatalf("due after %d others", i)
		}
		p.Completions = append(p.Completions, evidence(id, "z-other", "introduction", false))
	}
	if v := learningView(p); !v.Skills[0].ReviewDue {
		t.Fatal(v)
	}
	p.Completions = append(p.Completions, evidence("review-001", "skill", "practice", true))
	if v := learningView(p); v.Skills[0].ReviewDue || v.Skills[0].Remaining != 3 {
		t.Fatal(v)
	}
}
func TestDueReviewChoosesPreparedDistinctPractice(t *testing.T) {
	r := Route{Blocks: []RouteBlock{{ID: "review", Number: 2, Steps: []RouteStep{{TaskID: "unready-001", Role: "practice", Skills: []string{"skill"}}, {TaskID: "ready-001", Number: 2, Role: "transfer", Skills: []string{"skill"}}}}}}
	catalog := map[string]tasking.LearnerTaskInfo{"unready-001": {ID: "unready-001", Prerequisites: []string{"missing"}}, "ready-001": {ID: "ready-001", Prerequisites: []string{"skill"}}}
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}, Completions: []CompletionEvidence{evidence("intro-001", "skill", "introduction", false), evidence("one-001", "other", "introduction", false), evidence("two-001", "other", "introduction", false), evidence("three-001", "other", "introduction", false)}}
	choice, ok := dueReviewChoice(r, catalog, p, map[string]bool{"skill": true})
	if !ok || choice.Task.ID != "ready-001" || choice.ReviewSkill != "skill" || choice.LessonNumber != "02.02" {
		t.Fatal(choice, ok)
	}
	p.Tasks["current-001"] = &TaskProgress{}
	p.CurrentTaskID = "current-001"
	if _, ok = dueReviewChoice(r, catalog, p, map[string]bool{"skill": true}); ok {
		t.Fatal("interrupted an ongoing lesson")
	}
}
func TestCompletionSnapshotPreservesAssistanceAtSuccess(t *testing.T) {
	s := uiTestService(t, "")
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}
	task := &TaskProgress{TaskID: "zero-results-001"}
	p.Tasks[task.TaskID] = task
	info, err := s.Show(task.TaskID)
	if err != nil {
		t.Fatal(err)
	}
	if err = s.recordCompletion(&p, task, info, "2026-09-08T12:00:00Z"); err != nil {
		t.Fatal(err)
	}
	task.SolutionRevealed = true
	if p.Completions[0].Assisted || p.Completions[0].RouteRole == "" {
		t.Fatal(p.Completions)
	}
	if err = s.recordCompletion(&p, task, info, "later"); err != nil || len(p.Completions) != 1 {
		t.Fatal(p.Completions, err)
	}
}
func TestV2LearningUpgradeBacksUpWithoutMovingNumberedFiles(t *testing.T) {
	s := uiTestService(t, "")
	workspace, err := s.Start("literal-bindings-001")
	if err != nil {
		t.Fatal(err)
	}
	p, err := loadProgress(s.home)
	if err != nil {
		t.Fatal(err)
	}
	p.SchemaVersion = 2
	if err = saveProgress(s.home, p); err != nil {
		t.Fatal(err)
	}
	before := persistedBytes(t, s)
	if _, err = s.CourseView(); err != nil {
		t.Fatal(err)
	}
	plan, err := s.Organize(true)
	if err != nil || plan.ToVersion != 3 || len(plan.Moves) != 0 || !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal(plan, err)
	}
	plan, err = s.Organize(false)
	if err != nil {
		t.Fatal(err)
	}
	backup, err := os.ReadFile(filepath.Join(s.home, plan.Backup, "progress.json"))
	if err != nil || !bytes.Equal(backup, before) {
		t.Fatal("backup differs", err)
	}
	after, err := loadProgress(s.home)
	if err != nil || after.SchemaVersion != 3 {
		t.Fatal(after, err)
	}
	after.SchemaVersion = 2
	a, _ := json.Marshal(after)
	b, _ := json.Marshal(p)
	if !bytes.Equal(a, b) {
		t.Fatal("upgrade changed history")
	}
	if _, err = os.Stat(workspace); err != nil {
		t.Fatal(err)
	}
}

func TestAssistanceSurvivesResetThroughServiceOperations(t *testing.T) {
	for _, mode := range []string{"hint", "solution", "legacy"} {
		t.Run(mode, func(t *testing.T) {
			s := uiTestService(t, "")
			s.verify = func(string, string, string, tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
				return tasking.LearnerVerificationResult{Passed: true}, nil
			}
			if mode == "legacy" {
				if err := saveProgress(s.home, Progress{SchemaVersion: 1, Tasks: map[string]*TaskProgress{}}); err != nil {
					t.Fatal(err)
				}
			}
			id := "append-value-001"
			if _, err := s.Start(id); err != nil {
				t.Fatal(err)
			}
			if mode == "solution" {
				if _, err := s.Solution(true); err != nil {
					t.Fatal(err)
				}
			} else {
				if _, _, err := s.Hint(); err != nil {
					t.Fatal(err)
				}
			}
			if mode == "legacy" {
				if _, err := s.Complete(); err != nil {
					t.Fatal(err)
				}
				if _, err := s.Organize(false); err != nil {
					t.Fatal(err)
				}
			}
			if err := s.ResetTask(id, true); err != nil {
				t.Fatal(err)
			}
			if _, err := s.Start(id); err != nil {
				t.Fatal(err)
			}
			if _, err := s.Complete(); err != nil {
				t.Fatal(err)
			}
			p, err := loadProgress(s.home)
			if err != nil {
				t.Fatal(err)
			}
			if len(p.Completions) != 1 || !p.Completions[0].Assisted || !p.Assistance[id] {
				t.Fatal(p.Completions, p.Assistance)
			}
			if mode == "legacy" && !p.Completions[0].Legacy {
				t.Fatal("legacy identity lost")
			}
			if v := learningView(p); v.Mastered != 0 || v.Skills[0].Independent != 0 {
				t.Fatal(v)
			}
		})
	}
}
func TestRoutePreservesStartedReviewBeyondEarlierGap(t *testing.T) {
	p := Progress{SchemaVersion: 3, CurrentTaskID: "later-task-001", Tasks: map[string]*TaskProgress{"later-task-001": {TaskID: "later-task-001"}}}
	choice, err := routeNext(testRoute(), testRouteCatalog(), p)
	if err != nil || choice.Task.ID != "later-task-001" || !choice.Resume {
		t.Fatal(choice, err)
	}
}
func TestOldV2JournalRecoveryThenSeparateV3Upgrade(t *testing.T) {
	for _, stage := range []int{0, 1, 2} {
		t.Run(string(rune('0'+stage)), func(t *testing.T) {
			s := legacyOrganizeService(t)
			j := stageOrganizeRecovery(t, s)
			j.Plan.ToVersion = 2
			j.After.SchemaVersion = 2
			j.After.Completions = nil
			j.After.Assistance = nil
			data, _ := json.Marshal(j)
			if err := os.WriteFile(filepath.Join(s.home, organizeJournalName), data, 0600); err != nil {
				t.Fatal(err)
			}
			for _, m := range j.Plan.Moves[:stage] {
				if err := ensurePrivateDirectory(s.home, filepath.Dir(m.To)); err != nil {
					t.Fatal(err)
				}
				if err := os.Rename(filepath.Join(s.home, m.From), filepath.Join(s.home, m.To)); err != nil {
					t.Fatal(err)
				}
			}
			if stage == 2 {
				if err := saveProgress(s.home, j.After); err != nil {
					t.Fatal(err)
				}
			}
			result, err := s.Organize(false)
			if err != nil || result.ToVersion != 2 || !result.Recovery {
				t.Fatal(result, err)
			}
			p, err := loadProgress(s.home)
			if err != nil || !sameProgress(p, j.After) {
				t.Fatal(p, err)
			}
			before := persistedBytes(t, s)
			result, err = s.Organize(false)
			if err != nil || result.ToVersion != 3 || result.Recovery || result.Backup == j.Plan.Backup {
				t.Fatal(result, err)
			}
			backup, err := os.ReadFile(filepath.Join(s.home, result.Backup, "progress.json"))
			if err != nil || !bytes.Equal(backup, before) {
				t.Fatal("new exact backup missing", err)
			}
			p, err = loadProgress(s.home)
			if err != nil || len(p.Completions) != 1 || !p.Completions[0].Legacy || !p.Assistance["literal-bindings-001"] {
				t.Fatal(p, err)
			}
		})
	}
}
