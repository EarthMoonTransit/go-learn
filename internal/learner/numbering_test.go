package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func numberedTestRoute() Route {
	r := testRoute()
	r.SchemaVersion = 2
	for i := range r.Blocks {
		r.Blocks[i].Number = i + 1
		for j := range r.Blocks[i].Steps {
			r.Blocks[i].Steps[j].Number = j + 1
		}
	}
	return r
}
func TestExplicitRouteNumbersRejectInvalidAssignments(t *testing.T) {
	skills := map[string]SkillMetadata{"setup": {}, "basic": {Prerequisites: []string{"setup"}}, "later": {Prerequisites: []string{"basic"}}}
	for _, test := range []struct {
		name   string
		change func(*Route)
	}{
		{"missing block", func(r *Route) { r.Blocks[0].Number = 0 }},
		{"duplicate block", func(r *Route) { r.Blocks[1].Number = 1 }},
		{"missing lesson", func(r *Route) { r.Blocks[0].Steps[0].Number = 0 }},
		{"duplicate lesson", func(r *Route) { r.Blocks[0].Steps[1].Number = 1 }},
		{"descending lesson", func(r *Route) { r.Blocks[0].Steps[0].Number = 3 }},
		{"unsafe block", func(r *Route) { r.Blocks[0].ID = "../basics" }},
		{"large block", func(r *Route) { r.Blocks[1].Number = 1000 }},
		{"v1 numbers", func(r *Route) { r.SchemaVersion = 1 }},
	} {
		t.Run(test.name, func(t *testing.T) {
			r := numberedTestRoute()
			test.change(&r)
			if validateRoute(r, skills) == nil {
				t.Fatal("invalid numbering accepted")
			}
		})
	}
	r := numberedTestRoute()
	r.Blocks[1].Number = 7
	r.Blocks[0].Steps[1].Number = 9
	if err := validateRoute(r, skills); err != nil {
		t.Fatal(err)
	}
	if got := lessonNumber(r.Blocks[0], r.Blocks[0].Steps[1]); got != "01.09" {
		t.Fatal(got)
	}
}
func TestNumberingPreservesExistingFoundationAssignments(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	r, err := loadRoute(root)
	if err != nil {
		t.Fatal(err)
	}
	if r.SchemaVersion != 2 || len(r.Blocks) != 9 {
		t.Fatalf("unexpected active route: %+v", r)
	}
	expected := [][]string{
		{"literal-bindings-001", "zero-results-001", "access-logic-001", "division-results-001", "exclusive-loop-001"},
		{"named-block-001", "array-copy-001", "slice-window-001", "append-value-001", "map-presence-001", "map-take-001", "detached-append-001"},
		{"bounded-value-001", "wide-sum-001", "encoded-width-001", "range-offset-001", "clamp-tests-001"},
		{"replace-byte-001", "exact-prefix-001", "typed-line-001", "join-values-001", "byte-window-tests-001"},
	}
	for i, b := range r.Blocks {
		if b.Number != i+1 {
			t.Fatalf("changed block number %s", b.ID)
		}
	}
	for i, ids := range expected {
		for j, id := range ids {
			step := r.Blocks[i].Steps[j]
			if step.TaskID != id || step.Number != j+1 {
				t.Fatalf("changed assignment for %s: %+v", id, step)
			}
		}
	}
}
func TestNumberingAgreesAcrossReadOnlyViewsAndCLI(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	home := t.TempDir()
	s, err := New(Config{Root: root, Home: home})
	if err != nil {
		t.Fatal(err)
	}
	course, err := s.CourseView()
	if err != nil {
		t.Fatal(err)
	}
	if course.Blocks[0].Title != "01. Variables, values, expressions and functions" || course.Blocks[0].Lessons[0].Number != "01.01" {
		t.Fatalf("course=%+v", course.Blocks[0])
	}
	view, err := s.TaskView("literal-bindings-001")
	if err != nil || view.Number != "01.01" {
		t.Fatalf("view=%+v err=%v", view, err)
	}
	for _, args := range [][]string{{"list"}, {"next"}, {"roadmap"}, {"show", "literal-bindings-001"}} {
		var out, errors bytes.Buffer
		if status := RunCLI(Config{Root: root, Home: home}, args, &out, &errors); status != 0 {
			t.Fatalf("%v: %s", args, errors.String())
		}
		if !strings.Contains(out.String(), "01.01") || !strings.Contains(out.String(), "literal-bindings-001") {
			t.Fatalf("%v: %s", args, out.String())
		}
	}
	if _, err := os.Stat(filepath.Join(home, "progress.json")); !os.IsNotExist(err) {
		t.Fatalf("reading numbers created progress: %v", err)
	}
}
func TestExplicitNumbersDoNotChangeNextOrCompletedHistory(t *testing.T) {
	r := numberedTestRoute()
	p := Progress{Tasks: map[string]*TaskProgress{"early-task-001": {Completed: true, Attempts: 4, UsedHints: []int{1}, ActivitySeconds: 320}}}
	next, err := routeNext(r, testRouteCatalog(), p)
	if err != nil || next.Task.ID != "later-task-001" || next.LessonNumber != "01.02" {
		t.Fatalf("next=%+v err=%v", next, err)
	}
	if p.Tasks["early-task-001"].Attempts != 4 || p.Tasks["early-task-001"].ActivitySeconds != 320 {
		t.Fatal("history changed")
	}
}
