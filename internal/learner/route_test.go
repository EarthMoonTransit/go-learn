package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

func testRoute() Route {
	return Route{SchemaVersion: 1, Title: "Course", Preparation: "Read the guide", PreparationSkills: []string{"setup"},
		Blocks: []RouteBlock{
			{ID: "basics", Title: "Basics", Steps: []RouteStep{
				{TaskID: "early-task-001", Role: "introduction", Skills: []string{"basic"}},
				{TaskID: "later-task-001", Role: "practice", Skills: []string{"later"}},
			}},
			{ID: "future", Title: "Future", Steps: []RouteStep{}},
		}}
}
func testRouteCatalog() []tasking.LearnerTaskInfo {
	return []tasking.LearnerTaskInfo{
		{ID: "early-task-001", PrimarySkill: "basic", Prerequisites: []string{"setup"}},
		{ID: "later-task-001", PrimarySkill: "later", Prerequisites: []string{"basic"}},
	}
}
func TestRouteKeepsSequenceAndExistingProgress(t *testing.T) {
	route := testRoute()
	completed := &TaskProgress{TaskID: "later-task-001", Completed: true, Attempts: 4, UsedHints: []int{1}, ActivitySeconds: 4000, SkillEncounters: []SkillEncounter{{SkillID: "later", Role: "primary", Count: 1}}}
	progress := Progress{SchemaVersion: 1, CurrentTaskID: "later-task-001", Tasks: map[string]*TaskProgress{"later-task-001": completed}}
	before := *completed
	choice, err := routeNext(route, testRouteCatalog(), progress)
	if err != nil || choice.Task.ID != "early-task-001" || !reflect.DeepEqual(choice.MissingPrerequisites, []string{"setup"}) {
		t.Fatalf("choice=%+v err=%v", choice, err)
	}
	if !reflect.DeepEqual(before, *completed) || progress.CurrentTaskID != "later-task-001" {
		t.Fatal("read-only selection changed progress")
	}
	progress.Tasks["early-task-001"] = &TaskProgress{Completed: true, SkillEncounters: []SkillEncounter{{SkillID: "basic", Count: 1}}}
	_, err = routeNext(route, testRouteCatalog(), progress)
	if err == nil || !strings.Contains(err.Error(), "next block is planned") {
		t.Fatalf("planned future block was skipped: %v", err)
	}
}
func TestRouteDoesNotSkipUnavailableNextLesson(t *testing.T) {
	_, err := routeNext(testRoute(), testRouteCatalog()[1:], Progress{})
	if err == nil || !strings.Contains(err.Error(), "early-task-001") {
		t.Fatalf("missing first lesson did not block recommendation: %v", err)
	}
}
func TestRouteRecommendsNextAfterCompletion(t *testing.T) {
	progress := Progress{Tasks: map[string]*TaskProgress{"early-task-001": {Completed: true, SkillEncounters: []SkillEncounter{{SkillID: "basic", Count: 1}}}}}
	choice, err := routeNext(testRoute(), testRouteCatalog(), progress)
	if err != nil || choice.Task.ID != "later-task-001" || len(choice.MissingPrerequisites) != 0 {
		t.Fatalf("choice=%+v err=%v", choice, err)
	}
}
func TestRouteRejectsMetadataDriftAndDependencyJumps(t *testing.T) {
	skills := map[string]SkillMetadata{"setup": {}, "basic": {Prerequisites: []string{"setup"}}, "later": {Prerequisites: []string{"basic"}}}
	if err := validateRoute(testRoute(), skills); err != nil {
		t.Fatal(err)
	}
	t.Run("duplicate", func(t *testing.T) {
		r := testRoute()
		r.Blocks[0].Steps[1].TaskID = r.Blocks[0].Steps[0].TaskID
		if validateRoute(r, skills) == nil {
			t.Fatal("duplicate accepted")
		}
	})
	t.Run("unknown skill", func(t *testing.T) {
		r := testRoute()
		r.Blocks[0].Steps[1].Skills = []string{"missing"}
		if validateRoute(r, skills) == nil {
			t.Fatal("unknown skill accepted")
		}
	})
	t.Run("order", func(t *testing.T) {
		r := testRoute()
		r.Blocks[0].Steps[0], r.Blocks[0].Steps[1] = r.Blocks[0].Steps[1], r.Blocks[0].Steps[0]
		if validateRoute(r, skills) == nil {
			t.Fatal("prerequisite jump accepted")
		}
	})
	t.Run("catalog", func(t *testing.T) {
		tasks := testRouteCatalog()
		tasks[0].PrimarySkill = "later"
		if _, err := routeCatalog(testRoute(), tasks); err == nil {
			t.Fatal("catalog drift accepted")
		}
	})
	t.Run("extra prerequisite", func(t *testing.T) {
		tasks := testRouteCatalog()
		tasks[0].Prerequisites = append(tasks[0].Prerequisites, "later")
		if _, err := routeCatalog(testRoute(), tasks); err == nil {
			t.Fatal("new earlier prerequisite silently accepted")
		}
	})
}
func TestRoadmapShowsReadinessWithoutImplyingMastery(t *testing.T) {
	var output bytes.Buffer
	catalog, _ := routeCatalog(testRoute(), testRouteCatalog()[:1])
	writeRoadmap(&output, Roadmap{Route: testRoute(), Tasks: catalog, Progress: Progress{}})
	for _, want := range []string{"Ready lessons: 1/2 defined; completed: 0", "Basics [current]", "later-task-001", "Future [planned]", "not proof of mastery"} {
		if !strings.Contains(output.String(), want) {
			t.Fatalf("missing %q in %s", want, output.String())
		}
	}
}
func TestRepositoryRouteMatchesReadyTaskMetadata(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	route, err := loadSchedulingRoute(root)
	if err != nil {
		t.Fatal(err)
	}
	if route == nil {
		t.Fatal("repository must ship the personal route")
	}
	tasks, err := tasking.ListReadyLearnerTasks(root)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := routeCatalog(*route, tasks); err != nil {
		t.Fatal(err)
	}
}
func TestAbsentRouteKeepsLegacyRepositoriesCompatible(t *testing.T) {
	route, err := loadRoute(t.TempDir())
	if err != nil || route != nil {
		t.Fatalf("route=%+v err=%v", route, err)
	}
}
func TestMalformedRouteIsNotSilentlyIgnored(t *testing.T) {
	for _, content := range []string{"schema_version: 1\nunknown: true\n", "schema_version: 1\n---\nschema_version: 1\n"} {
		root := t.TempDir()
		if err := os.Mkdir(filepath.Join(root, "curriculum"), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, "curriculum", "route.yaml"), []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := loadRoute(root); err == nil {
			t.Fatal("invalid route silently accepted")
		}
	}
}

func TestNextPrintsResumeForExistingUnfinishedWorkspace(t *testing.T) {
	progress := Progress{Tasks: map[string]*TaskProgress{"early-task-001": {Status: "paused", Attempts: 2}}}
	choice, err := routeNext(testRoute(), testRouteCatalog(), progress)
	if err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	writeNext(&output, choice)
	if !strings.Contains(output.String(), "learn resume early-task-001") || strings.Contains(output.String(), "learn start") {
		t.Fatal(output.String())
	}
	if progress.Tasks["early-task-001"].Status != "paused" || progress.Tasks["early-task-001"].Attempts != 2 {
		t.Fatal("next mutated existing session")
	}
}

func TestRouteRejectsDeprecatedPreparationAndLessonSkills(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	skills, err := activeRouteSkills(root)
	if err != nil {
		t.Fatal(err)
	}
	const deprecated = "types.constants-iota"
	if _, ok := skills[deprecated]; ok {
		t.Fatal("deprecated taxonomy entry entered route skill set")
	}
	r := Route{SchemaVersion: 1, Title: "Invalid", Preparation: "guide", PreparationSkills: []string{deprecated}, Blocks: []RouteBlock{{ID: "planned", Title: "Planned"}}}
	if validateRoute(r, skills) == nil {
		t.Fatal("deprecated preparation accepted")
	}
	r.PreparationSkills = nil
	r.Blocks[0].Steps = []RouteStep{{TaskID: "planned-task-001", Role: "introduction", Skills: []string{deprecated}}}
	if validateRoute(r, skills) == nil {
		t.Fatal("deprecated lesson accepted")
	}
}
