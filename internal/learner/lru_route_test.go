package learner

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const routeLRU = "bounded-lru-cache-001"

func TestLRURoutePreservesMilestone150Addresses(t *testing.T) {
	s := uiTestService(t, "")
	r, err := s.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	// Captured from master 3c00d704ad5bb573858444d206478018f064a3b5.
	baseline, err := os.ReadFile("testdata/foundation-150-addresses.tsv")
	if err != nil {
		t.Fatal(err)
	}
	positions := map[string]int{}
	index := 0
	for _, b := range r.Route.Blocks {
		for _, step := range b.Steps {
			positions[fmt.Sprintf("%s\t%s\t%s", lessonNumber(b, step), b.ID, step.TaskID)] = index
			index++
		}
	}
	last := -1
	for _, line := range strings.Split(strings.TrimSpace(string(baseline)), "\n") {
		position, ok := positions[line]
		if !ok || position <= last {
			t.Fatalf("old address/block/order changed: %s", line)
		}
		last = position
		fields := strings.Split(line, "\t")
		numbers := strings.Split(fields[0], ".")
		want := fmt.Sprintf("workspaces/%s-%s/%s-%s", numbers[0], fields[1], numbers[1], strings.TrimSuffix(fields[2], "-001"))
		if got, err := s.numberedWorkspaceRelative(fields[2]); err != nil || got != want {
			t.Fatalf("old workspace assignment %s: %s, %v; want %s", fields[2], got, err, want)
		}
	}
	if _, ok := positions["07.20\tlibraries\t"+routeLRU]; !ok {
		t.Fatal("LRU must retain address 07.20 in libraries")
	}
}

func TestLRURouteRecommendationsPreserveProgress(t *testing.T) {
	s := uiTestService(t, "")
	r, err := s.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	tasks, _, err := s.ListTasks()
	if err != nil {
		t.Fatal(err)
	}
	// Synthetic completed-prefix histories exercise policy, not authored code.
	prefix := func(stop string) Progress {
		p := Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}
		for _, b := range r.Route.Blocks {
			for _, step := range b.Steps {
				if step.TaskID == stop {
					return p
				}
				task := &TaskProgress{TaskID: step.TaskID, Completed: true}
				for _, skill := range step.Skills {
					task.SkillEncounters = append(task.SkillEncounters, SkillEncounter{SkillID: skill, Count: 1})
				}
				p.Tasks[step.TaskID] = task
				p.Completions = append(p.Completions, CompletionEvidence{TaskID: step.TaskID, Skills: step.Skills, RouteRole: step.Role})
			}
		}
		return p
	}
	for _, tc := range []struct {
		name, stop, current, want, number string
		oldRouteDone                      bool
	}{
		{name: "section boundary", stop: routeLRU, want: routeLRU, number: "07.20"},
		{name: "after LRU", stop: "diagnostic-helper-001", want: "diagnostic-helper-001", number: "08.01"},
		{name: "active later lesson", stop: routeLRU, current: "closed-receive-001", want: "closed-receive-001", number: "09.02"},
		{name: "old route completed", oldRouteDone: true, want: routeLRU, number: "07.20"},
		{name: "existing delayed review", stop: "parse-timestamp-001", want: routeLRU, number: "07.20"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := prefix(tc.stop)
			if tc.oldRouteDone {
				delete(p.Tasks, routeLRU)
				for i, e := range p.Completions {
					if e.TaskID == routeLRU {
						p.Completions = append(p.Completions[:i], p.Completions[i+1:]...)
						break
					}
				}
			}
			if tc.current != "" {
				p.CurrentTaskID = tc.current
				p.Tasks[tc.current] = &TaskProgress{TaskID: tc.current, Attempts: 4, UsedHints: []int{1}, ActivitySeconds: 321}
			}
			if err := saveProgress(s.home, p); err != nil {
				t.Fatal(err)
			}
			before := persistedBytes(t, s)
			choice, err := s.Next()
			if err != nil || choice.Task.ID != tc.want || choice.LessonNumber != tc.number || len(choice.MissingPrerequisites) != 0 {
				t.Fatalf("choice=%+v err=%v", choice, err)
			}
			if tc.name == "existing delayed review" && choice.ReviewSkill != "generics.reusable-data-structures" {
				t.Fatalf("expected existing generics review policy: %+v", choice)
			}
			if !bytes.Equal(before, persistedBytes(t, s)) {
				t.Fatal("recommendation changed saved progress")
			}
		})
	}
	// Earlier off-route completion has no route role. Enrolling it must not
	// manufacture a transfer success or replay an already completed task.
	p := prefix("diagnostic-helper-001")
	p.Completions[len(p.Completions)-1].RouteRole = ""
	before := learningView(p)
	choice, err := routeNext(r.Route, tasks, p)
	if err != nil || choice.Task.ID != "diagnostic-helper-001" || !reflect.DeepEqual(before, learningView(p)) || p.Completions[len(p.Completions)-1].RouteRole != "" {
		t.Fatalf("off-route completion changed: %+v %v", choice, err)
	}
}

func TestLRURouteReadOnlyViewsPreserveWorkspaces(t *testing.T) {
	for _, version := range []int{1, 2, 3} {
		t.Run(fmt.Sprint(version), func(t *testing.T) {
			s := uiTestService(t, "")
			// An unnumbered LRU workspace is also the pre-enrollment layout for
			// numbered homes. Create it through the legacy workspace policy.
			if err := saveProgress(s.home, Progress{SchemaVersion: 1, Tasks: map[string]*TaskProgress{}}); err != nil {
				t.Fatal(err)
			}
			workspace, err := s.Start(routeLRU)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, "subject", "subject.go"), []byte("package subject\n// preserved learner draft\n"), 0600); err != nil {
				t.Fatal(err)
			}
			p, err := s.Progress()
			if err != nil {
				t.Fatal(err)
			}
			p.SchemaVersion = version
			p.Tasks[routeLRU].Attempts = 4
			p.Tasks[routeLRU].UsedHints = []int{1}
			p.Tasks[routeLRU].ActivitySeconds = 321
			p.Assistance = map[string]bool{routeLRU: true}
			if err := saveProgress(s.home, p); err != nil {
				t.Fatal(err)
			}
			before := persistedBytes(t, s)
			tree, err := workspaceTree(workspace, "")
			if err != nil {
				t.Fatal(err)
			}
			view, err := s.TaskView(routeLRU)
			if err != nil || view.WorkspaceUnavailable || view.Number != "07.20" {
				t.Fatalf("view=%+v err=%v", view, err)
			}
			course, err := s.CourseView()
			if err != nil || course.NextTaskID != routeLRU {
				t.Fatalf("course next=%s err=%v", course.NextTaskID, err)
			}
			for _, args := range [][]string{{"list"}, {"next"}, {"roadmap"}, {"show", routeLRU}} {
				var out, errors bytes.Buffer
				if status := RunCLI(Config{Root: s.root, Home: s.home}, args, &out, &errors); status != 0 || !strings.Contains(out.String(), "07.20") {
					t.Fatalf("%v: status=%d output=%s errors=%s", args, status, &out, &errors)
				}
			}
			if !bytes.Equal(before, persistedBytes(t, s)) {
				t.Fatal("read-only views rewrote history or migrated progress")
			}
			if after, err := workspaceTree(workspace, ""); err != nil || after != tree {
				t.Fatalf("learner workspace changed: %v", err)
			}
			if _, err := os.Stat(filepath.Join(s.home, "workspaces/07-libraries/20-bounded-lru-cache")); !os.IsNotExist(err) {
				t.Fatalf("read-only access created a numbered replacement: %v", err)
			}
		})
	}
}
