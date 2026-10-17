package learner

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"go-practice-bank/internal/tasking"
)

// Evidence is captured at completion, so viewing a solution afterwards cannot
// rewrite an earlier independent result. Only distinct task IDs count.
type CompletionEvidence struct {
	TaskID      string   `json:"task_id"`
	CompletedAt string   `json:"completed_at"`
	Skills      []string `json:"skills"`
	RouteRole   string   `json:"route_role"`
	Assisted    bool     `json:"assisted"`
	Legacy      bool     `json:"legacy"`
}
type SkillLearning struct {
	ID          string `json:"id"`
	Independent int    `json:"independent"`
	Mastered    bool   `json:"mastered"`
	ReviewDue   bool   `json:"review_due"`
	Remaining   int    `json:"other_lessons_remaining"`
}
type LearningView struct {
	Mastered        int             `json:"mastered"`
	Due             int             `json:"due"`
	UpgradeRequired bool            `json:"upgrade_required"`
	Skills          []SkillLearning `json:"skills"`
}

// Legacy completion dates order scheduling evidence only. They do not establish
// assistance-at-completion, hence never confirm mastery retrospectively.
func completionHistory(p Progress) []CompletionEvidence {
	seen := map[string]bool{}
	for _, e := range p.Completions {
		seen[e.TaskID] = true
	}
	history := []CompletionEvidence{}
	for id, t := range p.Tasks {
		if t == nil || !t.Completed || seen[id] {
			continue
		}
		skills := []string{}
		for _, e := range t.SkillEncounters {
			if e.Count > 0 {
				skills = append(skills, e.SkillID)
			}
		}
		if len(skills) == 0 && t.PrimarySkill != "" {
			skills = append(skills, t.PrimarySkill)
		}
		history = append(history, CompletionEvidence{TaskID: id, CompletedAt: t.CompletedAt, Skills: skills, Assisted: true, Legacy: true})
	}
	sort.Slice(history, func(i, j int) bool {
		if history[i].CompletedAt != history[j].CompletedAt {
			return history[i].CompletedAt < history[j].CompletedAt
		}
		return history[i].TaskID < history[j].TaskID
	})
	for _, e := range p.Completions {
		history = append(history, e)
	}
	return history
}
func (s *Service) recordCompletion(p *Progress, t *TaskProgress, info tasking.LearnerTaskInfo, now string) error {
	if p.SchemaVersion < 3 {
		return nil
	}
	history := completionHistory(*p)
	for _, e := range history {
		if e.TaskID == info.ID {
			return nil
		}
	}
	role := ""
	route, err := loadRoute(s.root)
	if err != nil {
		return err
	}
	if route != nil {
		for _, b := range route.Blocks {
			for _, step := range b.Steps {
				if step.TaskID == info.ID {
					role = step.Role
				}
			}
		}
	}
	skills := append([]string{info.PrimarySkill}, info.SecondarySkills...)
	p.Completions = append(history, CompletionEvidence{TaskID: info.ID, CompletedAt: now, Skills: skills, RouteRole: role, Assisted: p.Assistance[info.ID] || len(t.UsedHints) > 0 || t.SolutionRevealed})
	return nil
}
func learningView(p Progress) LearningView {
	view := LearningView{UpgradeRequired: p.SchemaVersion < 3, Skills: []SkillLearning{}}
	history := completionHistory(p)
	last := map[string]int{}
	independent := map[string]map[string]bool{}
	seen := map[string]bool{}
	index := 0
	for _, e := range history {
		if seen[e.TaskID] {
			continue
		}
		seen[e.TaskID] = true
		for _, skill := range e.Skills {
			last[skill] = index
			if !e.Legacy && !e.Assisted && (e.RouteRole == "practice" || e.RouteRole == "transfer") {
				if independent[skill] == nil {
					independent[skill] = map[string]bool{}
				}
				independent[skill][e.TaskID] = true
			}
		}
		index++
	}
	for id, at := range last {
		n := len(independent[id])
		remaining := 3 - (index - 1 - at)
		if remaining < 0 {
			remaining = 0
		}
		item := SkillLearning{ID: id, Independent: n, Mastered: n >= 2, Remaining: remaining, ReviewDue: n < 2 && remaining == 0}
		if item.Mastered {
			view.Mastered++
		}
		if item.ReviewDue {
			view.Due++
		}
		view.Skills = append(view.Skills, item)
	}
	sort.Slice(view.Skills, func(i, j int) bool { return view.Skills[i].ID < view.Skills[j].ID })
	return view
}
func dueReviewChoice(route Route, catalog map[string]tasking.LearnerTaskInfo, p Progress, known map[string]bool) (SchedulerChoice, bool) {
	if p.SchemaVersion < 3 {
		return SchedulerChoice{}, false
	}
	if t := p.Tasks[p.CurrentTaskID]; t != nil && !t.Completed {
		return SchedulerChoice{}, false
	}
	due := map[string]bool{}
	for _, skill := range learningView(p).Skills {
		if skill.ReviewDue {
			due[skill.ID] = true
		}
	}
	for _, id := range route.PreparationSkills {
		known[id] = true
	}
	completed := completedReadinessTasks(p)
	for _, lesson := range executionLessons(route) {
		step := lesson.step
		if step.Role != "practice" && step.Role != "transfer" {
			continue
		}
		task, ok := catalog[step.TaskID]
		if !ok {
			continue
		}
		if old := p.Tasks[task.ID]; old != nil && old.Completed {
			continue
		}
		already := false
		for _, e := range p.Completions {
			if e.TaskID == task.ID {
				already = true
			}
		}
		if already {
			continue
		}
		required := append([]string{}, task.Prerequisites...)
		if route.pedagogical {
			required = append(required, step.Skills...)
			required = append(required, task.NewConcepts...)
		}
		prepared := true
		for _, skill := range required {
			if !known[skill] {
				prepared = false
			}
		}
		if !prepared || (route.pedagogical && !reviewedExposureReady(route.exposureEdges, task.ID, completed)) {
			continue
		}
		for _, skill := range step.Skills {
			if due[skill] {
				choice := lessonChoice(route, lesson, task, p)
				choice.ReviewSkill = skill
				introduced := map[string]bool{}
				for _, id := range required {
					introduced[id] = true
				}
				names := make([]string, 0, len(introduced))
				for id := range introduced {
					names = append(names, id)
				}
				sort.Strings(names)
				choice.Reason = fmt.Sprintf("Delayed review: %s; introduced: %s. Three distinct other lessons completed.", skill, strings.Join(names, ", "))
				return choice, true
			}
		}
	}
	return SchedulerChoice{}, false
}
func writeLearning(out io.Writer, p Progress) {
	v := learningView(p)
	fmt.Fprintln(out, "A completion is not proof of mastery; two independent practice/transfer tasks are required.")
	fmt.Fprintf(out, "Confirmed skills: %d; delayed reviews due: %d\n", v.Mastered, v.Due)
	if v.UpgradeRequired {
		fmt.Fprintln(out, "Enable completion evidence with learn organize --dry-run, then learn organize. Existing completions remain progress.")
	}
	for _, s := range v.Skills {
		state := "practice needed"
		if s.Mastered {
			state = "confirmed"
		} else if s.ReviewDue {
			state = "review due (next ready practice/transfer lesson)"
		}
		fmt.Fprintf(out, "  %s: %d/2 independent tasks; %s\n", s.ID, s.Independent, state)
	}
}

func rememberAssistance(p *Progress, id string) {
	if p.SchemaVersion < 3 {
		return
	}
	if p.Assistance == nil {
		p.Assistance = map[string]bool{}
	}
	p.Assistance[id] = true
}
func upgradeLearningProgress(p *Progress, target int) {
	if p.SchemaVersion < 3 && target >= 3 {
		p.Completions = completionHistory(*p)
		p.SchemaVersion = target
		for id, t := range p.Tasks {
			if t != nil && (len(t.UsedHints) > 0 || t.SolutionRevealed) {
				rememberAssistance(p, id)
			}
		}
	}
	p.SchemaVersion = target
}
