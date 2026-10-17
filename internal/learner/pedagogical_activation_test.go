package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func TestExecutionOrderValidatesPreparationWithoutRenumbering(t *testing.T) {
	r := Route{SchemaVersion: 2, Title: "Course", Preparation: "guide", PreparationSkills: []string{"setup"}, Blocks: []RouteBlock{{Number: 1, ID: "basics", Title: "Basics", Steps: []RouteStep{
		{Number: 1, TaskID: "early-task-001", Role: "introduction", Skills: []string{"basic"}},
		{Number: 2, TaskID: "last-task-001", Role: "practice", Skills: []string{"last"}},
		{Number: 3, TaskID: "middle-task-001", Role: "introduction", Skills: []string{"middle"}},
	}}}}
	skills := map[string]SkillMetadata{"setup": {}, "basic": {Prerequisites: []string{"setup"}}, "middle": {Prerequisites: []string{"basic"}}, "last": {Prerequisites: []string{"middle"}}}
	tasks := []tasking.LearnerTaskInfo{{ID: "early-task-001", PrimarySkill: "basic", Prerequisites: []string{"setup"}}, {ID: "last-task-001", PrimarySkill: "last", Prerequisites: []string{"middle"}}, {ID: "middle-task-001", PrimarySkill: "middle", Prerequisites: []string{"basic"}}}
	if validateRoute(r, skills) == nil {
		t.Fatal("address-order prerequisite jump accepted")
	}
	if _, err := routeCatalog(r, tasks); err == nil {
		t.Fatal("actual manifest jump accepted")
	}
	addressed := executionLessons(r)
	r.execution = []routeLesson{addressed[0], addressed[2], addressed[1]}
	r.pedagogical = true
	if err := validateRoute(r, skills); err != nil {
		t.Fatal(err)
	}
	if _, err := routeCatalog(r, tasks); err != nil {
		t.Fatal(err)
	}
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{"early-task-001": {Completed: true, SkillEncounters: []SkillEncounter{{SkillID: "basic", Count: 1}}}}}
	choice, err := routeNext(r, tasks, p)
	if err != nil || choice.Task.ID != "middle-task-001" || choice.LessonNumber != "01.03" || len(choice.MissingPrerequisites) != 0 {
		t.Fatal(choice, err)
	}
	p.CurrentTaskID = "last-task-001"
	p.Tasks[p.CurrentTaskID] = &TaskProgress{Status: "paused", Attempts: 4, ActivitySeconds: 321, UsedHints: []int{1}}
	before, _ := json.Marshal(p)
	choice, err = routeNext(r, tasks, p)
	after, _ := json.Marshal(p)
	if err != nil || !choice.Resume || choice.Task.ID != "last-task-001" || choice.LessonNumber != "01.02" || !bytes.Equal(before, after) {
		t.Fatal(choice, err)
	}
}

func TestActiveReviewWaitsForSecondaryIntroductionWithoutRequiringMastery(t *testing.T) {
	r := Route{pedagogical: true, PreparationSkills: []string{"setup"}, Blocks: []RouteBlock{{Number: 1, ID: "practice", Title: "Practice", Steps: []RouteStep{
		{Number: 1, TaskID: "introduce-secondary-001", Role: "introduction", Skills: []string{"secondary"}},
		{Number: 2, TaskID: "review-task-001", Role: "practice", Skills: []string{"due", "secondary"}},
	}}}}
	tasks := []tasking.LearnerTaskInfo{{ID: "introduce-secondary-001", PrimarySkill: "secondary", Prerequisites: []string{"setup"}}, {ID: "review-task-001", PrimarySkill: "due", SecondarySkills: []string{"secondary"}, NewConcepts: []string{"secondary"}, Prerequisites: []string{"setup"}}}
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}, Completions: []CompletionEvidence{evidence("old-intro-001", "due", "introduction", false), evidence("other-a-001", "other", "introduction", false), evidence("other-b-001", "other", "introduction", false), evidence("other-c-001", "other", "introduction", false)}}
	before := learningView(p)
	choice, err := routeNext(r, tasks, p)
	if err != nil || choice.Task.ID != "introduce-secondary-001" || choice.ReviewSkill != "" || !strings.Contains(choice.Reason, "reviews remain due") || !reflect.DeepEqual(before, learningView(p)) {
		t.Fatal(choice, err)
	}
	p.Tasks["introduce-secondary-001"] = &TaskProgress{Completed: true, SkillEncounters: []SkillEncounter{{SkillID: "secondary", Count: 1}}}
	choice, err = routeNext(r, tasks, p)
	if err != nil || choice.Task.ID != "review-task-001" || choice.ReviewSkill != "due" || learningView(p).Mastered != 0 || !strings.Contains(choice.Reason, "introduced: due, secondary, setup") {
		t.Fatal(choice, err)
	}
	// A reset cannot turn this same task ID into another review opportunity.
	p.Completions = append(p.Completions, evidence("review-task-001", "due", "practice", true))
	if _, ok := dueReviewChoice(r, map[string]tasking.LearnerTaskInfo{tasks[1].ID: tasks[1]}, p, map[string]bool{"due": true, "secondary": true}); ok {
		t.Fatal("reset manufactured new review evidence")
	}
}

func TestActivePointerReviewNeedsActualMutationSharingPreparation(t *testing.T) {
	s := uiTestService(t, "")
	task, err := s.Show("pointer-chain-001")
	if err != nil {
		t.Fatal(err)
	}
	r := Route{pedagogical: true, Blocks: []RouteBlock{{Number: 6, ID: "contracts", Title: "Pointers", Steps: []RouteStep{{Number: 28, TaskID: task.ID, Role: "practice", Skills: []string{task.PrimarySkill}}}}}}
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}, Completions: []CompletionEvidence{evidence("pointer-intro-001", task.PrimarySkill, "introduction", false), evidence("one-001", "other", "introduction", false), evidence("two-001", "other", "introduction", false), evidence("three-001", "other", "introduction", false)}}
	known := map[string]bool{task.PrimarySkill: true, "types.zero-values": true}
	catalog := map[string]tasking.LearnerTaskInfo{task.ID: task}
	if _, ok := dueReviewChoice(r, catalog, p, known); ok {
		t.Fatal("PointerChain offered before mutation-sharing was introduced")
	}
	known["pointers.mutation-sharing"] = true
	if choice, ok := dueReviewChoice(r, catalog, p, known); !ok || choice.Task.ID != task.ID {
		t.Fatal(choice, ok)
	}
}

func TestReadinessCompatibilityDoesNotRetagHistoricalEvidence(t *testing.T) {
	r := Route{pedagogical: true}
	catalog := map[string]tasking.LearnerTaskInfo{"clamp-tests-001": {ContentDigest: clampWithTestingBridge}, "value-average-001": {ContentDigest: "5b7a7b79dd6b2d355ad5901c0f81853b26d6d52cfbe7d8b5f59b708e8a012b68"}}
	for _, id := range []string{"clamp-tests-001", "value-average-001"} {
		skill := "testing.basics-diagnostics"
		if id == "value-average-001" {
			skill = "methods.values-expressions"
		}
		for _, completed := range []bool{false, true} {
			p := Progress{SchemaVersion: 1, Tasks: map[string]*TaskProgress{id: {TaskID: id, Completed: completed, PrimarySkill: "old-skill", Status: "paused", Attempts: 3, ActivitySeconds: 412, LastPublicCheck: &PublicCheckRecord{Passed: true}}}}
			before, _ := json.Marshal(p)
			clock := learningView(p)
			known := introducedLearnerSkills(r, catalog, p)
			after, _ := json.Marshal(p)
			if known[skill] != completed || !bytes.Equal(before, after) || !reflect.DeepEqual(clock, learningView(p)) {
				t.Fatal(id, completed, known)
			}
		}
		// A retained completion remains familiarity after reset, with its old skills.
		p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{id: {TaskID: id}}, Completions: []CompletionEvidence{evidence(id, "old-skill", "practice", true)}}
		before, _ := json.Marshal(p)
		clock := learningView(p)
		if !introducedLearnerSkills(r, catalog, p)[skill] {
			t.Fatal("lost recorded familiarity on reset", id)
		}
		after, _ := json.Marshal(p)
		if !bytes.Equal(before, after) || !reflect.DeepEqual(clock, learningView(p)) {
			t.Fatal("retagged history or changed mastery/clock", id)
		}
		changed := map[string]tasking.LearnerTaskInfo{id: {ContentDigest: "unreviewed-revision"}}
		if introducedLearnerSkills(r, changed, p)[skill] {
			t.Fatal("equivalence extended to an unknown revision", id)
		}
	}
}

func activateBoundedFixture(t *testing.T, s *Service) {
	t.Helper()
	plan, _, err := loadPedagogicalPlan(s.root)
	if err != nil {
		t.Fatal(err)
	}
	plan.SchemaVersion = 2
	plan.Status = "active"
	plan.PendingPreparation = nil
	raw, err := yaml.Marshal(plan)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(s.root, pedagogicalPlanPath), raw, 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := loadSchedulingRoute(s.root); err != nil {
		t.Fatal(err)
	}
}

func TestActivationKeepsPausedAndLegacyWorkspaceBytesAndIDETargets(t *testing.T) {
	for _, version := range []int{1, 3} {
		for _, paused := range []bool{false, true} {
			t.Run(strings.Join([]string{string(rune('0' + version)), map[bool]string{true: "paused", false: "active"}[paused]}, "/"), func(t *testing.T) {
				s := boundedPedagogicalService(t)
				if err := saveProgress(s.home, Progress{SchemaVersion: version, Tasks: map[string]*TaskProgress{}}); err != nil {
					t.Fatal(err)
				}
				workspace, err := s.Start(routeLRU)
				if err != nil {
					t.Fatal(err)
				}
				source := filepath.Join(workspace, "subject/subject.go")
				f, err := os.OpenFile(source, os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					t.Fatal(err)
				}
				if _, err = f.WriteString("\n// retained learner draft\n"); err != nil {
					t.Fatal(err)
				}
				if err = f.Close(); err != nil {
					t.Fatal(err)
				}
				if paused {
					if err = s.Pause(); err != nil {
						t.Fatal(err)
					}
				}
				before, err := loadProgress(s.home)
				if err != nil {
					t.Fatal(err)
				}
				tree, err := workspaceTree(s.home, "")
				if err != nil {
					t.Fatal(err)
				}
				target, err := s.goLandTarget(routeLRU, workspace)
				if err != nil {
					t.Fatal(err)
				}
				activateBoundedFixture(t, s)
				choice, err := s.Next()
				if err != nil || !choice.Resume || choice.Task.ID != routeLRU || choice.LessonNumber != "07.20" {
					t.Fatal(choice, err)
				}
				course, err := s.CourseView()
				if err != nil || course.NextTaskID != routeLRU || !strings.Contains(course.NextMessage, choice.Reason) {
					t.Fatal(course, err)
				}
				view, err := s.TaskView(routeLRU)
				if err != nil || view.WorkspaceUnavailable || view.Workspace != workspace || view.Number != "07.20" {
					t.Fatal(view, err)
				}
				if afterTarget, err := s.goLandTarget(routeLRU, workspace); err != nil || afterTarget != target {
					t.Fatal(afterTarget, err)
				}
				if _, err = s.PedagogicalPreview(); err != nil {
					t.Fatal(err)
				}
				if _, err = s.Roadmap(); err != nil {
					t.Fatal(err)
				}
				var out, diagnostics bytes.Buffer
				if code := RunCLI(Config{Root: s.root, Home: s.home}, []string{"next"}, &out, &diagnostics); code != 0 || !strings.Contains(out.String(), choice.Reason) {
					t.Fatal(code, out.String(), diagnostics.String())
				}
				after, err := loadProgress(s.home)
				if err != nil || !reflect.DeepEqual(before, after) || !reflect.DeepEqual(completionHistory(before), completionHistory(after)) || !reflect.DeepEqual(learningView(before), learningView(after)) {
					t.Fatal("activation rewrote history/timer/assistance", err)
				}
				afterTree, err := workspaceTree(s.home, "")
				if err != nil || tree != afterTree {
					t.Fatal("activation changed workspace or progress bytes", err)
				}
			})
		}
	}
}

func TestCompletedOriginalRouteGetsNewPreparationsWithoutHistoricalRetagging(t *testing.T) {
	s := uiTestService(t, "")
	baseline, err := os.ReadFile("testdata/review-baseline-addresses.tsv")
	if err != nil {
		t.Fatal(err)
	}
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}
	for _, line := range strings.Split(strings.TrimSpace(string(baseline)), "\n") {
		fields := strings.Split(line, "\t")
		id, role := fields[2], fields[3]
		task, err := s.Show(id)
		if err != nil {
			t.Fatal(err)
		}
		oldSkills := append([]string{task.PrimarySkill}, task.SecondarySkills...)
		if id == "clamp-tests-001" || id == "value-average-001" {
			oldSkills = []string{task.PrimarySkill}
		}
		stored := &TaskProgress{TaskID: id, Completed: true, PrimarySkill: task.PrimarySkill, Attempts: 2, ActivitySeconds: 91}
		for _, skill := range oldSkills {
			stored.SkillEncounters = append(stored.SkillEncounters, SkillEncounter{SkillID: skill, Count: 1})
		}
		p.Tasks[id] = stored
		p.Completions = append(p.Completions, CompletionEvidence{TaskID: id, Skills: oldSkills, RouteRole: role, Assisted: true})
	}
	if err := saveProgress(s.home, p); err != nil {
		t.Fatal(err)
	}
	before := persistedBytes(t, s)
	clock := learningView(p)
	choice, err := s.Next()
	if err != nil || choice.Task.ID != "conditional-binding-001" || choice.LessonNumber != "03.31" {
		t.Fatal(choice, err)
	}
	course, err := s.CourseView()
	if err != nil || course.Completed != 151 || course.Ready != 166 {
		t.Fatal(course, err)
	}
	after, err := loadProgress(s.home)
	if err != nil || !bytes.Equal(before, persistedBytes(t, s)) || !reflect.DeepEqual(clock, learningView(after)) || !reflect.DeepEqual(p.Completions, after.Completions) {
		t.Fatal("old complete route was retagged", err)
	}
}

func TestActiveOrderCannotFallBackWhenTheAddressRegistryIsMissing(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, pedagogicalPlanPath)
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("schema_version: 2\nstatus: active\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := loadSchedulingRoute(root); err == nil {
		t.Fatal("active plan silently fell back to a different scheduler")
	}
	if route, err := loadRoute(root); err != nil || route != nil {
		t.Fatal("missing address metadata is not readable", route, err)
	}
}

func TestActiveReviewCannotJumpPastReviewedClampExposure(t *testing.T) {
	s := uiTestService(t, "")
	route, err := loadSchedulingRoute(s.root)
	if err != nil {
		t.Fatal(err)
	}
	tasks, _, err := s.ListTasks()
	if err != nil {
		t.Fatal(err)
	}
	byID := map[string]tasking.LearnerTaskInfo{}
	roles := map[string]string{}
	for _, task := range tasks {
		byID[task.ID] = task
	}
	for _, lesson := range executionLessons(*route) {
		roles[lesson.step.TaskID] = lesson.step.Role
	}
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}
	for _, id := range []string{"literal-bindings-001", "zero-results-001", "access-logic-001", "division-results-001", "wide-sum-001", "grouped-unit-cost-001", "boolean-policies-001", "remaining-constant-001", "whole-minute-output-001", "percentage-output-001", "constant-kinds-001", "month-enumeration-001", "bounded-value-001", "conditional-binding-001"} {
		task := byID[id]
		skills := append([]string{task.PrimarySkill}, task.SecondarySkills...)
		p.Tasks[id] = &TaskProgress{TaskID: id, Completed: true}
		p.Completions = append(p.Completions, CompletionEvidence{TaskID: id, Skills: skills, RouteRole: roles[id]})
	}
	choice, err := routeNext(*route, tasks, p)
	if err != nil || choice.Task.ID != "clamp-tests-001" {
		t.Fatalf("review skipped required Clamp introduction: %+v %v", choice, err)
	}
	// All broad LazyProbe mechanisms are familiar, but its reviewed Clamp gate
	// still requires the actual prior completion; mere public success is not it.
	p.Tasks["clamp-tests-001"] = &TaskProgress{LastPublicCheck: &PublicCheckRecord{Passed: true}}
	known := introducedLearnerSkills(*route, byID, p)
	if reviewedExposureReady(route.exposureEdges, "lazy-probe-001", completedReadinessTasks(p)) {
		t.Fatal("started/public-passing Clamp satisfied exposure")
	}
	if !known["control-flow.if"] {
		t.Fatal("fixture lacks broad if familiarity")
	}
}

func TestActiveTraversalRespectsReviewedGatesWithAndWithoutAssistance(t *testing.T) {
	s := uiTestService(t, "")
	route, err := loadSchedulingRoute(s.root)
	if err != nil {
		t.Fatal(err)
	}
	tasks, _, err := s.ListTasks()
	if err != nil {
		t.Fatal(err)
	}
	for _, assisted := range []bool{false, true} {
		t.Run(map[bool]string{false: "independent", true: "assisted"}[assisted], func(t *testing.T) {
			p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}
			completed := map[string]bool{}
			for len(completed) < len(route.execution) {
				choice, err := routeNext(*route, tasks, p)
				if err != nil {
					t.Fatal(err)
				}
				id := choice.Task.ID
				if completed[id] {
					t.Fatal("replayed completed ID", id)
				}
				for _, edge := range route.exposureEdges {
					if edge.TaskID == id && !completed[edge.BeforeTaskID] {
						t.Fatalf("%s selected before reviewed exposure %s (review=%s)", id, edge.BeforeTaskID, choice.ReviewSkill)
					}
				}
				if choice.PedagogicalChapter == "" || choice.PedagogicalSubtopic == "" || choice.PedagogicalPosition < 1 || choice.LessonNumber == "" {
					t.Fatal("selection lost display-only context", choice)
				}
				if choice.ReviewSkill != "" && !strings.Contains(choice.Reason, "introduced:") {
					t.Fatal("review reason lacks introduced mechanisms", choice)
				}
				skills := append([]string{choice.Task.PrimarySkill}, choice.Task.SecondarySkills...)
				p.Tasks[id] = &TaskProgress{TaskID: id, Completed: true}
				p.Completions = append(p.Completions, CompletionEvidence{TaskID: id, Skills: skills, RouteRole: choice.RouteRole, Assisted: assisted})
				completed[id] = true
			}
			if _, err := routeNext(*route, tasks, p); err == nil || !strings.Contains(err.Error(), "all defined route lessons are completed") {
				t.Fatal("exhausted route replayed a review", err)
			}
			if assisted && learningView(p).Mastered != 0 {
				t.Fatal("assisted traversal manufactured mastery")
			}
		})
	}
}

func TestActiveDisplayShowsPedagogicalContextAtTheStableAddress(t *testing.T) {
	s := uiTestService(t, "")
	p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}
	for _, id := range []string{"literal-bindings-001", "zero-results-001", "access-logic-001", "division-results-001"} {
		task, err := s.Show(id)
		if err != nil {
			t.Fatal(err)
		}
		p.Tasks[id] = &TaskProgress{TaskID: id, Completed: true}
		p.Completions = append(p.Completions, CompletionEvidence{TaskID: id, Skills: append([]string{task.PrimarySkill}, task.SecondarySkills...), RouteRole: "introduction", Assisted: true})
	}
	if err := saveProgress(s.home, p); err != nil {
		t.Fatal(err)
	}
	before := persistedBytes(t, s)
	choice, err := s.Next()
	if err != nil || choice.Task.ID != "wide-sum-001" || choice.LessonNumber != "03.02" || choice.PedagogicalPosition != 5 || !strings.HasPrefix(choice.PedagogicalChapter, "C01 ") {
		t.Fatal(choice, err)
	}
	course, err := s.CourseView()
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"C01 ", "position 5", "address 03.02", choice.PedagogicalSubtopic, choice.Reason} {
		if !strings.Contains(course.NextMessage, want) {
			t.Fatalf("UI context %q lacks %q", course.NextMessage, want)
		}
	}
	for _, block := range course.Blocks {
		if block.ID == "control-and-verification" && block.Status != "current" {
			t.Fatal("selected address block is not current", block)
		}
		if block.ID == "first-steps" && block.Status == "current" {
			t.Fatal("earlier address gap overrode selected lesson", block)
		}
	}
	var output, diagnostics bytes.Buffer
	if code := RunCLI(Config{Root: s.root, Home: s.home}, []string{"next"}, &output, &diagnostics); code != 0 {
		t.Fatal(code, diagnostics.String())
	}
	for _, want := range []string{"C01 ", "position 5", "address 03.02"} {
		if !strings.Contains(output.String(), want) {
			t.Fatalf("CLI context lacks %q: %s", want, &output)
		}
	}
	roadmap, err := s.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	output.Reset()
	writeRoadmap(&output, roadmap)
	if !strings.Contains(output.String(), "03. Control flow, numeric conversions and verification [current]") {
		t.Fatal("roadmap current does not follow selection", output.String())
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("display context rewrote progress")
	}
}
