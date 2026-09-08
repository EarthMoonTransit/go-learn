package learner

import (
	"fmt"
	"go-practice-bank/internal/tasking"
	"io"
	"regexp"
	"text/tabwriter"
)

var routeBlockIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func numberedBlockTitle(block RouteBlock) string {
	if block.Number == 0 {
		return block.Title
	}
	return fmt.Sprintf("%02d. %s", block.Number, block.Title)
}

func lessonNumber(block RouteBlock, step RouteStep) string {
	if block.Number == 0 || step.Number == 0 {
		return ""
	}
	return fmt.Sprintf("%02d.%02d", block.Number, step.Number)
}

// LessonNumber resolves presentation metadata without changing task identity or
// learner state. Legacy routes have no explicit numbers.
func (s *Service) LessonNumber(id string) (string, error) {
	route, err := loadRoute(s.root)
	if err != nil || route == nil {
		return "", err
	}
	for _, block := range route.Blocks {
		for _, step := range block.Steps {
			if step.TaskID == id {
				return lessonNumber(block, step), nil
			}
		}
	}
	return "", nil
}

func writeNumberedTaskList(output io.Writer, route Route, tasks []tasking.LearnerTaskInfo, progress Progress) {
	writer := tabwriter.NewWriter(output, 0, 4, 2, ' ', 0)
	fmt.Fprintln(writer, "LESSON\tTITLE\tSTATUS\tACTIVITY\tID")
	catalog := map[string]tasking.LearnerTaskInfo{}
	for _, task := range tasks {
		catalog[task.ID] = task
	}
	for _, block := range route.Blocks {
		for _, step := range block.Steps {
			task, ok := catalog[step.TaskID]
			if !ok {
				continue
			}
			status := "not-started"
			if p := progress.Tasks[task.ID]; p != nil {
				status = p.Status
			}
			fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", lessonNumber(block, step), task.Title, status, task.Activity, task.ID)
			delete(catalog, task.ID)
		}
	}
	// Keep ready tasks outside a route visible without inventing a course number.
	for _, task := range tasks {
		if _, ok := catalog[task.ID]; ok {
			status := "not-started"
			if p := progress.Tasks[task.ID]; p != nil {
				status = p.Status
			}
			fmt.Fprintf(writer, "—\t%s\t%s\t%s\t%s\n", task.Title, status, task.Activity, task.ID)
		}
	}
	_ = writer.Flush()
}
