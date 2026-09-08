package learner

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/triage"
	"gopkg.in/yaml.v3"
)

// Route is curriculum metadata, never a source of learner completion records.
type Route struct {
	SchemaVersion     int          `yaml:"schema_version"`
	Title             string       `yaml:"title"`
	Preparation       string       `yaml:"preparation"`
	PreparationSkills []string     `yaml:"preparation_skills"`
	Blocks            []RouteBlock `yaml:"blocks"`
}
type RouteBlock struct {
	Number int         `yaml:"number,omitempty"`
	ID     string      `yaml:"id"`
	Title  string      `yaml:"title"`
	Steps  []RouteStep `yaml:"steps"`
}
type RouteStep struct {
	Number int      `yaml:"number,omitempty"`
	TaskID string   `yaml:"task_id"`
	Role   string   `yaml:"role"`
	Skills []string `yaml:"skills"`
}
type Roadmap struct {
	Route    Route
	Tasks    map[string]tasking.LearnerTaskInfo
	Progress Progress
}

func loadRoute(root string) (*Route, error) {
	content, err := os.ReadFile(filepath.Join(root, "curriculum", "route.yaml"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)
	var route Route
	if err := decoder.Decode(&route); err != nil {
		return nil, fmt.Errorf("curriculum route: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, errors.New("curriculum route must contain exactly one document")
	}
	skills, err := activeRouteSkills(root)
	if err != nil {
		return nil, err
	}
	if err := validateRoute(route, skills); err != nil {
		return nil, err
	}
	if route.Preparation != "" {
		if filepath.IsAbs(route.Preparation) || !filepath.IsLocal(route.Preparation) {
			return nil, errors.New("preparation guide must be repository-relative")
		}
		guide := filepath.Join(root, route.Preparation)
		if info, err := os.Stat(guide); err != nil || !info.Mode().IsRegular() {
			return nil, errors.New("preparation guide is missing")
		}
		route.Preparation = guide
	}
	return &route, nil
}

func activeRouteSkills(root string) (map[string]SkillMetadata, error) {
	_, contracts, err := triage.LoadSkillContracts(root)
	if err != nil {
		return nil, err
	}
	skills := map[string]SkillMetadata{}
	for id, contract := range contracts {
		if contract.Status == "active" && contract.Assignable {
			skills[id] = SkillMetadata{ID: id, Prerequisites: contract.Prerequisites}
		}
	}
	return skills, nil
}

func validateRoute(route Route, skills map[string]SkillMetadata) error {
	if (route.SchemaVersion != 1 && route.SchemaVersion != 2) || strings.TrimSpace(route.Title) == "" || len(route.Blocks) == 0 {
		return errors.New("curriculum route requires version 1 or 2, a title and blocks")
	}
	known := map[string]bool{}
	for _, id := range route.PreparationSkills {
		if _, ok := skills[id]; !ok {
			return fmt.Errorf("unknown preparation skill %s", id)
		}
		if known[id] {
			return fmt.Errorf("duplicate preparation skill %s", id)
		}
		known[id] = true
	}
	if len(known) > 0 && strings.TrimSpace(route.Preparation) == "" {
		return errors.New("preparation skills require learner guidance")
	}
	ids, blocks := map[string]bool{}, map[string]bool{}
	lastBlockNumber := 0
	for _, block := range route.Blocks {
		if route.SchemaVersion == 2 {
			if block.Number <= lastBlockNumber || block.Number > 999 || !routeBlockIDPattern.MatchString(block.ID) {
				return errors.New("route v2 requires safe block IDs and increasing explicit numbers in 1..999")
			}
			lastBlockNumber = block.Number
		} else if block.Number != 0 {
			return errors.New("explicit route numbers require version 2")
		}
		if block.ID == "" || strings.TrimSpace(block.Title) == "" || blocks[block.ID] {
			return errors.New("route blocks require unique IDs and titles")
		}
		blocks[block.ID] = true
		lastLessonNumber := 0
		for _, step := range block.Steps {
			if route.SchemaVersion == 2 {
				if step.Number <= lastLessonNumber || step.Number > 999 {
					return errors.New("route v2 requires increasing explicit lesson numbers in 1..999")
				}
				lastLessonNumber = step.Number
			} else if step.Number != 0 {
				return errors.New("explicit route numbers require version 2")
			}
			if !learnerTaskIDPattern.MatchString(step.TaskID) || ids[step.TaskID] {
				return fmt.Errorf("invalid or repeated route task %s", step.TaskID)
			}
			ids[step.TaskID] = true
			if step.Role != "introduction" && step.Role != "practice" && step.Role != "transfer" {
				return fmt.Errorf("invalid route role for %s", step.TaskID)
			}
			if len(step.Skills) == 0 {
				return fmt.Errorf("route task %s requires declared skills", step.TaskID)
			}
			seen := map[string]bool{}
			// Check the entire step before adding its outcomes. A second skill may
			// deliberately introduce a dependency of the first in the same exercise.
			for _, id := range step.Skills {
				skill, ok := skills[id]
				if !ok || seen[id] {
					return fmt.Errorf("unknown or repeated route skill %s", id)
				}
				seen[id] = true
				for _, prerequisite := range skill.Prerequisites {
					if !known[prerequisite] && !hasRouteSkill(step.Skills, prerequisite) {
						return fmt.Errorf("route task %s lacks earlier preparation for %s", step.TaskID, prerequisite)
					}
				}
			}
			for id := range seen {
				known[id] = true
			}
		}
	}
	return nil
}
func hasRouteSkill(skills []string, id string) bool {
	for _, skill := range skills {
		if skill == id {
			return true
		}
	}
	return false
}
func routeCatalog(route Route, tasks []tasking.LearnerTaskInfo) (map[string]tasking.LearnerTaskInfo, error) {
	catalog := map[string]tasking.LearnerTaskInfo{}
	for _, task := range tasks {
		catalog[task.ID] = task
	}
	known := map[string]bool{}
	for _, id := range route.PreparationSkills {
		known[id] = true
	}
	for _, block := range route.Blocks {
		for _, step := range block.Steps {
			if task, ok := catalog[step.TaskID]; ok {
				assessed := append([]string{task.PrimarySkill}, task.SecondarySkills...)
				if len(assessed) != len(step.Skills) || task.PrimarySkill != step.Skills[0] {
					return nil, fmt.Errorf("route skill contract differs for %s", step.TaskID)
				}
				for _, skill := range assessed {
					if !hasRouteSkill(step.Skills, skill) {
						return nil, fmt.Errorf("route skill contract differs for %s", step.TaskID)
					}
				}
				for _, skill := range task.Prerequisites {
					if !known[skill] {
						return nil, fmt.Errorf("route task %s requires %s before this position", step.TaskID, skill)
					}
				}
			}
			for _, skill := range step.Skills {
				known[skill] = true
			}
		}
	}
	return catalog, nil
}
func (s *Service) Roadmap() (Roadmap, error) {
	route, err := loadRoute(s.root)
	if err != nil {
		return Roadmap{}, err
	}
	if route == nil {
		return Roadmap{}, errors.New("no curriculum route is configured")
	}
	tasks, progress, err := s.ListTasks()
	if err != nil {
		return Roadmap{}, err
	}
	catalog, err := routeCatalog(*route, tasks)
	if err != nil {
		return Roadmap{}, err
	}
	return Roadmap{Route: *route, Tasks: catalog, Progress: progress}, nil
}
func knownLearnerSkills(progress Progress) map[string]bool {
	known := map[string]bool{}
	for _, task := range progress.Tasks {
		if task == nil || !task.Completed {
			continue
		}
		for _, encounter := range task.SkillEncounters {
			if encounter.Count > 0 {
				known[encounter.SkillID] = true
			}
		}
	}
	return known
}
func routeNext(route Route, tasks []tasking.LearnerTaskInfo, progress Progress) (SchedulerChoice, error) {
	catalog, err := routeCatalog(route, tasks)
	if err != nil {
		return SchedulerChoice{}, err
	}
	known := knownLearnerSkills(progress)
	if current := progress.Tasks[progress.CurrentTaskID]; current != nil && !current.Completed {
		for _, block := range route.Blocks {
			for _, step := range block.Steps {
				if step.TaskID == progress.CurrentTaskID {
					if task, ok := catalog[step.TaskID]; ok {
						return SchedulerChoice{Resume: true, Task: task, LessonNumber: lessonNumber(block, step), RouteBlock: numberedBlockTitle(block), RouteRole: step.Role, Preparation: route.Preparation}, nil
					}
				}
			}
		}
	}
	if choice, ok := dueReviewChoice(route, catalog, progress, known); ok {
		return choice, nil
	}
	for _, block := range route.Blocks {
		if len(block.Steps) == 0 {
			return SchedulerChoice{}, fmt.Errorf("next block is planned: %s; no lessons are ready yet", block.Title)
		}
		for _, step := range block.Steps {
			if stored := progress.Tasks[step.TaskID]; stored != nil && stored.Completed {
				continue
			}
			task, ok := catalog[step.TaskID]
			if !ok {
				return SchedulerChoice{}, fmt.Errorf("next lesson is not ready: %s (%s); see learn roadmap", step.TaskID, block.Title)
			}
			choice := SchedulerChoice{Resume: progress.Tasks[task.ID] != nil, Task: task, RouteBlock: numberedBlockTitle(block), RouteRole: step.Role, Preparation: route.Preparation, LessonNumber: lessonNumber(block, step)}
			for _, skill := range task.Prerequisites {
				if !known[skill] {
					choice.MissingPrerequisites = append(choice.MissingPrerequisites, skill)
				}
			}
			sort.Strings(choice.MissingPrerequisites)
			return choice, nil
		}
	}
	return SchedulerChoice{}, errors.New("all defined route lessons are completed; see learn roadmap for the next milestone")
}
func writeRoadmap(output io.Writer, roadmap Roadmap) {
	fmt.Fprintln(output, roadmap.Route.Title)
	available, completed, total := 0, 0, 0
	for _, block := range roadmap.Route.Blocks {
		for _, step := range block.Steps {
			total++
			if _, ok := roadmap.Tasks[step.TaskID]; ok {
				available++
			}
			if p := roadmap.Progress.Tasks[step.TaskID]; p != nil && p.Completed {
				completed++
			}
		}
	}
	fmt.Fprintf(output, "Ready lessons: %d/%d defined; completed: %d\n", available, total, completed)
	writeLearning(output, roadmap.Progress)
	if roadmap.Route.Preparation != "" {
		fmt.Fprintf(output, "Preparation: %s\n", roadmap.Route.Preparation)
	}
	currentShown := false
	for _, block := range roadmap.Route.Blocks {
		done := 0
		for _, step := range block.Steps {
			if p := roadmap.Progress.Tasks[step.TaskID]; p != nil && p.Completed {
				done++
			}
		}
		state := "planned"
		for _, step := range block.Steps {
			if _, available := roadmap.Tasks[step.TaskID]; available {
				state = "ready"
				break
			}
		}
		if len(block.Steps) > 0 && done == len(block.Steps) {
			state = "completed"
		} else if !currentShown {
			state = "current"
			currentShown = true
		}
		fmt.Fprintf(output, "\n%s [%s] — %d/%d completed\n", numberedBlockTitle(block), state, done, len(block.Steps))
		if len(block.Steps) == 0 {
			fmt.Fprintln(output, "  Content planned; no task completion is implied.")
		}
		for _, step := range block.Steps {
			status := "planned"
			if _, ok := roadmap.Tasks[step.TaskID]; ok {
				status = "ready"
			}
			if p := roadmap.Progress.Tasks[step.TaskID]; p != nil {
				if p.Completed {
					status = "completed"
				} else {
					status = "in progress"
				}
			}
			title := roadmap.Tasks[step.TaskID].Title
			fmt.Fprintf(output, "  %-12s %s %s (%s; id: %s)\n", status, lessonNumber(block, step), title, step.Role, step.TaskID)
		}
	}
}
