package learner

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"
	"time"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

// Keep state regressions bounded without bypassing the real Service/CLI loaders.
// The full 151-task specification and its blockers are checked separately. This
// disposable fixture preserves actual task metadata and old addresses; omitted
// prerequisite introductions are explicit fixture-only preparation assumptions.
func boundedPedagogicalService(t *testing.T) *Service {
	t.Helper()
	s := uiTestService(t, "")
	route, err := loadRoute(s.root)
	if err != nil {
		t.Fatal(err)
	}
	plan, entries, err := loadPedagogicalPlan(s.root)
	if err != nil {
		t.Fatal(err)
	}
	skills, err := activeRouteSkills(s.root)
	if err != nil {
		t.Fatal(err)
	}
	selected := map[string]bool{}
	preparation := map[string]bool{}
	for _, id := range route.PreparationSkills {
		preparation[id] = true
	}
	root := t.TempDir()
	for _, dir := range []string{"config", "schemas", "sources", "curriculum"} {
		if err := isolation.CopyTree(filepath.Join(s.root, dir), filepath.Join(root, dir)); err != nil {
			t.Fatal(err)
		}
	}
	for _, id := range []string{"literal-bindings-001", "zero-results-001", "access-logic-001", "clamp-tests-001", "allocated-counter-001", "counter-read-tests-001", "pointer-chain-001", routeLRU} {
		selected[id] = true
		info, err := s.Show(id)
		if err != nil {
			t.Fatal(err)
		}
		for _, prerequisite := range info.Prerequisites {
			preparation[prerequisite] = true
		}
		if err := isolation.CopyTree(filepath.Join(s.root, "tasks/ready", id), filepath.Join(root, "tasks/ready", id)); err != nil {
			t.Fatal(err)
		}
	}
	var blocks []RouteBlock
	for _, block := range route.Blocks {
		var steps []RouteStep
		for _, step := range block.Steps {
			if selected[step.TaskID] {
				steps = append(steps, step)
				for _, id := range step.Skills {
					for _, prerequisite := range skills[id].Prerequisites {
						preparation[prerequisite] = true
					}
				}
			}
		}
		if len(steps) > 0 {
			block.Steps = steps
			blocks = append(blocks, block)
		}
	}
	route.Blocks = blocks
	route.PreparationSkills = nil
	for id := range preparation {
		route.PreparationSkills = append(route.PreparationSkills, id)
	}
	sort.Strings(route.PreparationSkills)
	route.Preparation = "docs/fixture-preparation.md"
	if err := writePrivateFile(filepath.Join(root, route.Preparation), []byte("Synthetic tooling fixture assumes exposure to these omitted prerequisites; no learner evidence is granted:\n"+strings.Join(route.PreparationSkills, "\n")+"\n"), 0600); err != nil {
		t.Fatal(err)
	}
	var edges []ExposureEdge
	for _, edge := range plan.ExposureEdges {
		if selected[edge.TaskID] && selected[edge.BeforeTaskID] {
			edges = append(edges, edge)
		}
	}
	plan.ExposureEdges = edges
	var pending []PendingPreparation
	for _, gate := range plan.PendingPreparation {
		if selected[gate.TaskID] {
			pending = append(pending, gate)
		}
	}
	plan.PendingPreparation = pending
	var mapping bytes.Buffer
	writer := csv.NewWriter(&mapping)
	if err := writer.Write([]string{"task_id", "old_number", "proposed_chapter", "proposed_subtopic", "proposed_order", "role", "disposition", "rationale"}); err != nil {
		t.Fatal(err)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].ProposedOrder < entries[j].ProposedOrder })
	position := 0
	for _, entry := range entries {
		if selected[entry.TaskID] {
			position++
			if err := writer.Write([]string{entry.TaskID, entry.OldNumber, entry.Chapter, entry.Subtopic, fmt.Sprint(position), entry.Role, entry.Disposition, entry.Rationale}); err != nil {
				t.Fatal(err)
			}
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}
	for path, value := range map[string]any{"curriculum/route.yaml": route, pedagogicalPlanPath: plan} {
		content, err := yaml.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, path), content, 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, plan.MappingFile), mapping.Bytes(), 0600); err != nil {
		t.Fatal(err)
	}
	s.root = root
	return s
}

// These histories are synthetic; no personal learner state is read or executed.
func pedagogicalProgressFixtures(route Route, tasks map[string]tasking.LearnerTaskInfo) map[string]Progress {
	steps := map[string]RouteStep{}
	for _, b := range route.Blocks {
		for _, step := range b.Steps {
			steps[step.TaskID] = step
		}
	}
	empty := func() Progress { return Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}} }
	complete := func(p *Progress, id string) {
		info := tasks[id]
		stamp := fmt.Sprintf("2026-10-01T00:%02d:%02dZ", len(p.Completions)/60, len(p.Completions)%60)
		task := &TaskProgress{TaskID: id, Status: "completed", Workspace: "workspaces/" + id, Completed: true, CompletedAt: stamp, PrimarySkill: info.PrimarySkill, Activity: info.Activity}
		assessed := steps[id].Skills
		if id == "clamp-tests-001" {
			// These October 1 histories predate E01/M1. Preserve the original
			// assessed skills even though the current route has a new secondary.
			assessed = []string{"expressions.comparison"}
		}
		for i, skill := range assessed {
			role := "secondary"
			if i == 0 {
				role = "primary"
			}
			task.SkillEncounters = append(task.SkillEncounters, SkillEncounter{SkillID: skill, Role: role, Count: 1})
		}
		p.Tasks[id] = task
		p.Completions = append(p.Completions, CompletionEvidence{TaskID: id, CompletedAt: stamp, Skills: slices.Clone(assessed), RouteRole: steps[id].Role})
	}
	fixtures := map[string]Progress{"fresh": empty()}
	for _, status := range []string{"in-progress", "paused"} {
		p := empty()
		p.CurrentTaskID = routeLRU // position 128 in the full plan, last in this fixture; old address 07.20
		task := &TaskProgress{TaskID: routeLRU, Status: status, Workspace: "workspaces/07-libraries/20-bounded-lru-cache", Attempts: 3, ActivitySeconds: 321, StartedAt: "2026-10-01T00:00:00Z", LastInteractionAt: "2026-10-01T00:02:00Z"}
		if status == "in-progress" {
			task.ActiveSessionStartedAt = "2026-10-01T00:02:00Z"
		}
		p.Tasks[routeLRU] = task
		fixtures[status] = p
	}
	p := empty()
	complete(&p, "clamp-tests-001")
	p.CurrentTaskID = "clamp-tests-001"
	p.Tasks[p.CurrentTaskID].Attempts = 7
	p.Tasks[p.CurrentTaskID].UsedHints = []int{1, 2}
	p.Tasks[p.CurrentTaskID].SolutionRevealed = true
	p.Tasks[p.CurrentTaskID].SolutionRevealedAt = "2026-10-01T00:05:00Z"
	p.Tasks[p.CurrentTaskID].ActivitySeconds = 1234
	p.Tasks[p.CurrentTaskID].Feedback = []FeedbackEntry{{RecordedAt: "2026-10-01T00:06:00Z", Text: "synthetic feedback"}}
	p.Tasks[p.CurrentTaskID].LastPublicCheck = &PublicCheckRecord{CheckedAt: "2026-10-01T00:04:00Z", Passed: true, Summary: "synthetic check"}
	p.Assistance = map[string]bool{p.CurrentTaskID: true}
	p.Completions[0].Assisted = true
	fixtures["completed Clamp without new secondary"] = p
	for _, version := range []int{1, 2} {
		p := empty()
		complete(&p, "clamp-tests-001")
		p.SchemaVersion = version
		p.Completions = nil
		p.CurrentTaskID = "clamp-tests-001"
		p.Tasks[p.CurrentTaskID].Attempts = 4
		p.Tasks[p.CurrentTaskID].UsedHints = []int{1}
		p.Tasks[p.CurrentTaskID].ActivitySeconds = 900
		fixtures[fmt.Sprintf("legacy v%d", version)] = p
	}
	p = empty()
	// Current metadata permits PointerChain review without mutation-sharing.
	// Pin that existing behavior; T0b must change eligibility in its own PR.
	for _, id := range []string{"allocated-counter-001", "counter-read-tests-001", "literal-bindings-001", "zero-results-001", "access-logic-001"} {
		complete(&p, id)
	}
	fixtures["due PointerChain without mutation preparation"] = p
	p = empty()
	for _, b := range route.Blocks {
		for _, step := range b.Steps {
			if step.TaskID != "literal-bindings-001" {
				complete(&p, step.TaskID)
			}
		}
	}
	fixtures["due review without candidates"] = p
	p = empty()
	for _, b := range route.Blocks {
		for _, step := range b.Steps {
			complete(&p, step.TaskID)
		}
	}
	fixtures["all completed"] = p
	return fixtures
}

func TestPedagogicalPreviewPreservesSyntheticHistoriesAndWorkspaceTrees(t *testing.T) {
	base := boundedPedagogicalService(t)
	roadmap, err := base.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	var tasks []tasking.LearnerTaskInfo
	for _, task := range roadmap.Tasks {
		tasks = append(tasks, task)
	}
	for name, fixture := range pedagogicalProgressFixtures(roadmap.Route, roadmap.Tasks) {
		t.Run(name, func(t *testing.T) {
			s := uiTestService(t, "")
			s.root = base.root
			s.now = func() time.Time { t.Fatal("read-only preview accessed the session clock"); return time.Time{} }
			s.verify = func(string, string, string, tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
				t.Fatal("preview invoked task verification")
				return tasking.LearnerVerificationResult{}, nil
			}
			s.launch = func(ideTarget) error { t.Fatal("preview launched the IDE"); return nil }
			if name != "fresh" {
				// Preserve deliberate whitespace too; decoding/re-encoding would fail
				// the home-tree comparison even if semantic fields stayed equal.
				content, err := json.MarshalIndent(fixture, "", "   ")
				if err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(s.home, "progress.json"), append(content, '\n', '\n'), 0600); err != nil {
					t.Fatal(err)
				}
				for _, task := range fixture.Tasks {
					if err := s.checkWorkspaceRelative(task.TaskID, task.Workspace); err != nil {
						t.Fatal(err)
					}
					if err := writePrivateFile(filepath.Join(s.home, task.Workspace, "learner.txt"), []byte("synthetic private content: "+task.TaskID+"\n"), 0600); err != nil {
						t.Fatal(err)
					}
				}
			}
			before, err := loadProgress(s.home)
			if err != nil {
				t.Fatal(err)
			}
			beforeTree, err := workspaceTree(s.home, "")
			if err != nil {
				t.Fatal(err)
			}
			beforeChoice, beforeErr := s.nextFor(tasks, before)
			preview, err := s.PedagogicalPreview()
			// This bounded subset has no remaining preparation actions after
			// PointerChain's M1 review. The full repository still has eight.
			if err != nil || !preview.Valid() || !preview.PreparationReady() {
				t.Fatalf("bounded preview must be valid and prepared: %+v, %v", preview, err)
			}
			after, err := loadProgress(s.home)
			if err != nil {
				t.Fatal(err)
			}
			afterTree, err := workspaceTree(s.home, "")
			if err != nil || beforeTree != afterTree || !reflect.DeepEqual(before, after) || !reflect.DeepEqual(completionHistory(before), completionHistory(after)) || !reflect.DeepEqual(learningView(before), learningView(after)) {
				t.Fatalf("preview changed bytes, paths, history, assistance or timer: %v", err)
			}
			afterChoice, afterErr := s.nextFor(tasks, after)
			if !reflect.DeepEqual(beforeChoice, afterChoice) || fmt.Sprint(beforeErr) != fmt.Sprint(afterErr) {
				t.Fatal("preview changed production Next/due review")
			}
			switch name {
			case "fresh", "completed Clamp without new secondary", "legacy v1", "legacy v2":
				if afterChoice.Task.ID != "literal-bindings-001" {
					t.Fatal(afterChoice, afterErr)
				}
			case "in-progress", "paused":
				if afterChoice.Task.ID != routeLRU || !afterChoice.Resume || afterChoice.LessonNumber != "07.20" {
					t.Fatal(afterChoice, afterErr)
				}
				last := preview.Lessons[len(preview.Lessons)-1]
				if last.TaskID != routeLRU || last.OldNumber != "07.20" || last.LearnerStatus != before.Tasks[routeLRU].Status {
					t.Fatal("fixture must retain the current task at the far end of the proposed plan", last)
				}
			case "due PointerChain without mutation preparation":
				if knownLearnerSkills(before)["pointers.mutation-sharing"] || afterChoice.Task.ID != "pointer-chain-001" || afterChoice.ReviewSkill != "pointers.address-deref-nil" {
					t.Fatal("fixture did not exercise the existing ungated PointerChain review", afterChoice, afterErr)
				}
			case "due review without candidates":
				if learningView(before).Due == 0 || afterChoice.ReviewSkill != "" || afterChoice.Task.ID != "literal-bindings-001" {
					t.Fatal(afterChoice, afterErr)
				}
			case "all completed":
				if afterErr == nil || !strings.Contains(afterErr.Error(), "all defined route lessons are completed") {
					t.Fatal(afterChoice, afterErr)
				}
			}
			if name == "completed Clamp without new secondary" {
				for _, evidence := range after.Completions {
					if slices.Contains(evidence.Skills, "testing.basics-diagnostics") {
						t.Fatal("preview retagged a historical Clamp completion")
					}
				}
			}
		})
	}
}

func TestPedagogicalPreviewCLIShowsStableAddressesAndBlockers(t *testing.T) {
	s := uiTestService(t, "")
	beforeTree, err := workspaceTree(s.home, "")
	if err != nil {
		t.Fatal(err)
	}
	var output, diagnostics bytes.Buffer
	if code := RunCLI(Config{Root: s.root, Home: s.home}, []string{"roadmap", "--pedagogical"}, &output, &diagnostics); code != 0 {
		t.Fatalf("code=%d: %s", code, &diagnostics)
	}
	for _, text := range []string{
		"inactive (activation disabled)", "Mapping: 151 rows; current route: 151 tasks", "Structural validity: true; preparation ready: false",
		"Position 15 | C02", "Address 03.30 | decision-table-tests-001 | transfer", "Address 03.26 | constant-value-tests-001 | practice",
		"PENDING P05", "PENDING P06", "PENDING P07",
		"Position 143 | C14", "Address 09.01 | worker-completion-tests-001 | introduction", "PENDING P11", "PENDING P12", "PENDING P13",
	} {
		if !strings.Contains(output.String(), text) {
			t.Fatalf("missing %q in preview", text)
		}
	}
	for _, resolved := range []string{"BLOCKED:", "PENDING E01/M1", "PENDING E02"} {
		if strings.Contains(output.String(), resolved) {
			t.Fatalf("resolved blocker remains in preview: %s", resolved)
		}
	}
	afterTree, err := workspaceTree(s.home, "")
	if err != nil || beforeTree != afterTree {
		t.Fatal("CLI preview wrote learner data", err)
	}
	// A bounded disposable catalog is intentionally incomplete: structural errors
	// must be displayed, exit nonzero, and never mutate progress/workspaces.
	s = goLandTestService(t)
	output.Reset()
	diagnostics.Reset()
	beforeTree, err = workspaceTree(s.home, "")
	if err != nil {
		t.Fatal(err)
	}
	if code := RunCLI(Config{Root: s.root, Home: s.home}, []string{"roadmap", "--pedagogical"}, &output, &diagnostics); code != 1 || !strings.Contains(output.String(), "Structural validity: false") || !strings.Contains(output.String(), "no ready v2 learner task") {
		t.Fatalf("code=%d: %s %s", code, &output, &diagnostics)
	}
	afterTree, err = workspaceTree(s.home, "")
	if err != nil || beforeTree != afterTree {
		t.Fatal("invalid CLI preview wrote learner data", err)
	}
}

func TestPedagogicalPreviewKeepsUIWorkspaceAndGoLandTargets(t *testing.T) {
	for _, version := range []int{1, 3} {
		t.Run(fmt.Sprint(version), func(t *testing.T) {
			s := boundedPedagogicalService(t)
			if err := saveProgress(s.home, Progress{SchemaVersion: version, Tasks: map[string]*TaskProgress{}}); err != nil {
				t.Fatal(err)
			}
			workspace, err := s.Start(routeLRU) // copies learner material as data only
			if err != nil {
				t.Fatal(err)
			}
			beforeNumber, err := s.LessonNumber(routeLRU)
			if err != nil {
				t.Fatal(err)
			}
			beforeTask, err := s.TaskView(routeLRU)
			if err != nil || beforeTask.WorkspaceUnavailable {
				t.Fatal(beforeTask, err)
			}
			beforeCourse, err := s.CourseView()
			if err != nil {
				t.Fatal(err)
			}
			beforeTarget, err := s.goLandTarget(routeLRU, workspace)
			if err != nil {
				t.Fatal(err)
			}
			beforeTree, err := workspaceTree(s.home, "")
			if err != nil {
				t.Fatal(err)
			}
			preview, err := s.PedagogicalPreview()
			if err != nil || !preview.Valid() {
				t.Fatal(preview, err)
			}
			afterTask, err := s.TaskView(routeLRU)
			if err != nil || !reflect.DeepEqual(beforeTask, afterTask) {
				t.Fatal("UI task view changed", err)
			}
			afterCourse, err := s.CourseView()
			if err != nil || !reflect.DeepEqual(beforeCourse, afterCourse) {
				t.Fatal("UI course changed", err)
			}
			afterTarget, err := s.goLandTarget(routeLRU, workspace)
			if err != nil || afterTarget != beforeTarget || afterTarget.Project != filepath.Join(s.home, "workspaces") || afterTarget.File != filepath.Join(workspace, "subject/subject.go") {
				t.Fatal("GoLand project or file target changed", err)
			}
			if number, err := s.LessonNumber(routeLRU); err != nil || number != beforeNumber || number != "07.20" {
				t.Fatal("lookup changed", number, err)
			}
			if expected, err := s.numberedWorkspaceRelative(routeLRU); err != nil || expected != "workspaces/07-libraries/20-bounded-lru-cache" {
				t.Fatal(expected, err)
			}
			p, err := loadProgress(s.home)
			if err != nil || s.checkWorkspaceRelative(routeLRU, p.Tasks[routeLRU].Workspace) != nil {
				t.Fatal("workspace lookup no longer accepts the original path", err)
			}
			afterTree, err := workspaceTree(s.home, "")
			if err != nil || afterTree != beforeTree {
				t.Fatal("UI/preview changed private files", err)
			}
		})
	}
}

func TestProductionReadOnlyPathsIgnoreMalformedPedagogicalPlan(t *testing.T) {
	s := goLandTestService(t)
	if err := os.WriteFile(filepath.Join(s.root, pedagogicalPlanPath), []byte("schema_version: 99\nstatus: active\n"), 0600); err != nil {
		t.Fatal(err)
	}
	before, err := workspaceTree(s.home, "")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.PedagogicalPreview(); err == nil {
		t.Fatal("unsupported active plan accepted")
	}
	for _, args := range [][]string{{"next"}, {"roadmap"}, {"show", "literal-bindings-001"}} {
		var out, diagnostics bytes.Buffer
		if code := RunCLI(Config{Root: s.root, Home: s.home}, args, &out, &diagnostics); code != 0 || !strings.Contains(out.String(), "01.01") {
			t.Fatalf("production %v: code=%d %s", args, code, &diagnostics)
		}
	}
	if view, err := s.CourseView(); err != nil || view.NextTaskID != "literal-bindings-001" {
		t.Fatal(view, err)
	}
	after, err := workspaceTree(s.home, "")
	if err != nil || before != after {
		t.Fatal("read-only failure/production paths changed home", err)
	}
}
