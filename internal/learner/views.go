package learner

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"go-practice-bank/internal/tasking"
)

type LessonView struct {
	Number    string `json:"number"`
	ID        string `json:"id"`
	Title     string `json:"title"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	Available bool   `json:"available"`
}
type BlockView struct {
	ID      string       `json:"id"`
	Title   string       `json:"title"`
	Status  string       `json:"status"`
	Lessons []LessonView `json:"lessons"`
}
type CourseView struct {
	Learning      LearningView `json:"learning"`
	Revision      string       `json:"revision"`
	CurrentTaskID string       `json:"current_task_id"`
	NextTaskID    string       `json:"next_task_id"`
	NextMessage   string       `json:"next_message"`
	Blocks        []BlockView  `json:"blocks"`
	Ready         int          `json:"ready"`
	Completed     int          `json:"completed"`
	IDEAvailable  bool         `json:"ide_available"`
	Preparation   string       `json:"preparation"`
}
type TaskView struct {
	Number               string                  `json:"number"`
	Revision             string                  `json:"revision"`
	Info                 tasking.LearnerTaskInfo `json:"info"`
	Started              bool                    `json:"started"`
	Current              bool                    `json:"current"`
	Completed            bool                    `json:"completed"`
	Attempts             int                     `json:"attempts"`
	SessionActive        bool                    `json:"session_active"`
	Workspace            string                  `json:"workspace"`
	IDEProject           string                  `json:"ide_project"`
	WorkspaceUnavailable bool                    `json:"workspace_unavailable"`
	UsedHints            []UsedHint              `json:"used_hints"`
	HintsRemaining       int                     `json:"hints_remaining"`
	Review               string                  `json:"review"`
}

func (s *Service) CourseView() (CourseView, error) {
	tasks, progress, err := s.ListTasks()
	if err != nil {
		return CourseView{}, err
	}
	revision, err := s.stateRevision(progress)
	if err != nil {
		return CourseView{}, err
	}
	view := CourseView{Revision: revision, CurrentTaskID: progress.CurrentTaskID, Blocks: []BlockView{}, Ready: len(tasks)}
	view.Learning = learningView(progress)
	view.IDEAvailable = s.ideAvailable
	catalog := map[string]tasking.LearnerTaskInfo{}
	for _, task := range tasks {
		catalog[task.ID] = task
		if p := progress.Tasks[task.ID]; p != nil && p.Completed {
			view.Completed++
		}
	}
	route, err := loadRoute(s.root)
	if err != nil {
		return CourseView{}, err
	}
	if route == nil {
		route = &Route{Blocks: []RouteBlock{{ID: "catalog", Title: "Available lessons"}}}
		for _, task := range tasks {
			route.Blocks[0].Steps = append(route.Blocks[0].Steps, RouteStep{TaskID: task.ID})
		}
	} else {
		if _, err := routeCatalog(*route, tasks); err != nil {
			return CourseView{}, err
		}
		if route.Preparation != "" {
			// The guide is a fixed repository input, never an HTTP filesystem path.
			relative, err := filepath.Rel(s.root, route.Preparation)
			if err != nil || !filepath.IsLocal(relative) {
				return CourseView{}, errors.New("invalid preparation path")
			}
			root, err := os.OpenRoot(s.root)
			if err != nil {
				return CourseView{}, err
			}
			info, err := root.Stat(relative)
			if err == nil && info.Mode().IsRegular() && info.Size() <= 64<<10 {
				data, readErr := root.ReadFile(relative)
				err = readErr
				view.Preparation = string(data)
			} else if err == nil {
				err = errors.New("invalid preparation guide")
			}
			closeErr := root.Close()
			if err != nil {
				return CourseView{}, err
			}
			if closeErr != nil {
				return CourseView{}, closeErr
			}
		}
	}
	choice, err := s.nextFor(tasks, progress)
	if err == nil {
		view.NextTaskID = choice.Task.ID
		if choice.ReviewSkill != "" {
			view.NextMessage = "Delayed review: " + choice.ReviewSkill
		}
	} else {
		view.NextMessage = "The next lesson is not prepared yet. See the planned blocks below."
	}
	currentShown := false
	for _, block := range route.Blocks {
		b := BlockView{ID: block.ID, Title: numberedBlockTitle(block), Status: "planned", Lessons: []LessonView{}}
		completed, prepared := 0, 0
		for _, step := range block.Steps {
			task, available := catalog[step.TaskID]
			lesson := LessonView{ID: step.TaskID, Title: task.Title, Role: step.Role, Status: "planned", Available: available, Number: lessonNumber(block, step)}
			if !available {
				lesson.Title = step.TaskID
			} else {
				lesson.Status = "ready"
				prepared++
			}
			if p := progress.Tasks[step.TaskID]; p != nil {
				if p.Completed {
					lesson.Status = "completed"
					completed++
				} else {
					lesson.Status = "in progress"
				}
			}
			b.Lessons = append(b.Lessons, lesson)
		}
		if prepared > 0 {
			b.Status = "ready"
		}
		if len(block.Steps) > 0 && completed == len(block.Steps) {
			b.Status = "completed"
		} else if !currentShown {
			b.Status = "current"
			currentShown = true
		}
		view.Blocks = append(view.Blocks, b)
	}
	return view, nil
}

// TaskView is a read-only snapshot. Looking at another task never resumes it,
// consumes a hint or updates the learner clock.
func (s *Service) TaskView(id string) (TaskView, error) {
	info, err := s.Show(id)
	if err != nil {
		return TaskView{}, err
	}
	progress, err := loadProgress(s.home)
	if err != nil {
		return TaskView{}, err
	}
	revision, err := s.stateRevision(progress)
	if err != nil {
		return TaskView{}, err
	}
	view := TaskView{Info: info, Revision: revision, Current: progress.CurrentTaskID == id, UsedHints: []UsedHint{}}
	view.Number, err = s.LessonNumber(id)
	if err != nil {
		return TaskView{}, err
	}
	hints, err := tasking.LearnerHints(s.root, id)
	if err != nil {
		return TaskView{}, err
	}
	view.HintsRemaining = len(hints)
	if task := progress.Tasks[id]; task != nil {
		view.Started = true
		view.Completed = task.Completed
		view.Attempts = task.Attempts
		view.SessionActive = task.ActiveSessionStartedAt != ""
		workspace, err := s.validateWorkspace(task, info)
		if err != nil {
			view.WorkspaceUnavailable = true
		} else {
			view.Workspace = workspace
			view.IDEProject = filepath.Join(s.home, "workspaces")
		}
		indices := append([]int(nil), task.UsedHints...)
		sort.Ints(indices)
		seen := map[int]bool{}
		for _, number := range indices {
			if number > 0 && number <= len(hints) && !seen[number] {
				view.UsedHints = append(view.UsedHints, UsedHint{number, hints[number-1]})
				seen[number] = true
			}
		}
		view.HintsRemaining -= len(view.UsedHints)
		if task.Completed && !view.WorkspaceUnavailable {
			review, err := s.reviewFor(task, info, workspace)
			if err != nil {
				return TaskView{}, err
			}
			view.Review = strings.TrimSpace(review)
		}
	}
	return view, nil
}
