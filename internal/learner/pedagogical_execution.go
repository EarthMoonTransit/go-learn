package learner

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

// A lesson carries its original address even when execution uses another order.
type routeLesson struct {
	block RouteBlock
	step  RouteStep
}

func executionLessons(route Route) []routeLesson {
	if route.execution != nil {
		return route.execution
	}
	var lessons []routeLesson
	for _, block := range route.Blocks {
		if len(block.Steps) == 0 {
			lessons = append(lessons, routeLesson{block: block})
			continue
		}
		for _, step := range block.Steps {
			lessons = append(lessons, routeLesson{block: block, step: step})
		}
	}
	return lessons
}

func configuredPedagogicalStatus(root string) (string, bool, error) {
	content, err := safeio.ReadFile(root, pedagogicalPlanPath, 1<<20)
	if errors.Is(err, os.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", true, err
	}
	var header struct {
		Status string `yaml:"status"`
	}
	if err := yaml.Unmarshal(content, &header); err != nil {
		return "", true, fmt.Errorf("pedagogical plan: %w", err)
	}
	return header.Status, true, nil
}

func configurePedagogicalExecution(root string, route *Route, skills map[string]SkillMetadata) error {
	status, exists, err := configuredPedagogicalStatus(root)
	if err != nil {
		return err
	}
	// Inactive planning data does not control production selection. Its strict
	// schema and mapping diagnostics remain available through the preview.
	if !exists || status == "inactive" {
		return nil
	}
	plan, entries, err := loadPedagogicalPlan(root)
	if err != nil {
		return err
	}
	if plan.SchemaVersion != 2 || plan.Status != "active" {
		return errors.New("pedagogical execution requires schema_version 2 and status active")
	}
	tasks, err := tasking.ListReadyLearnerTasks(root)
	if err != nil {
		return err
	}
	catalog, err := routeCatalogMetadata(*route, tasks)
	if err != nil {
		return err
	}
	preview := validatePedagogicalOrder(plan, entries, Roadmap{Route: *route, Tasks: catalog}, skills)
	if !preview.PreparationReady() {
		issues := append([]string{}, preview.StructuralIssues...)
		for _, lesson := range preview.Lessons {
			for _, blocker := range lesson.FormalBlockers {
				issues = append(issues, lesson.TaskID+": "+blocker)
			}
			for _, pending := range lesson.PendingPreparation {
				issues = append(issues, lesson.TaskID+": pending "+pending.ReviewID)
			}
		}
		return fmt.Errorf("active pedagogical order is not prepared: %s", strings.Join(issues, "; "))
	}
	addressed := map[string]routeLesson{}
	for _, lesson := range executionLessons(*route) {
		addressed[lesson.step.TaskID] = lesson
	}
	route.execution = make([]routeLesson, 0, len(preview.Lessons))
	for _, lesson := range preview.Lessons {
		route.execution = append(route.execution, addressed[lesson.TaskID])
	}
	route.pedagogical = true
	_, err = routeCatalog(*route, tasks)
	return err
}

// This is a read-only readiness view. Historical completion evidence, mastery,
// review clocks and assistance are never rewritten from today's task metadata.
func introducedLearnerSkills(route Route, catalog map[string]tasking.LearnerTaskInfo, p Progress) map[string]bool {
	known := knownLearnerSkills(p)
	completed := map[string]bool{}
	for id, task := range p.Tasks {
		if task != nil && task.Completed {
			completed[id] = true
		}
	}
	for _, e := range completionHistory(p) {
		completed[e.TaskID] = true
		for _, skill := range e.Skills {
			known[skill] = true
		}
	}
	for _, skill := range route.PreparationSkills {
		known[skill] = true
	}
	if completed["clamp-tests-001"] && catalog["clamp-tests-001"].ContentDigest == clampWithTestingBridge {
		known["testing.basics-diagnostics"] = true
	}
	if completed["value-average-001"] && catalog["value-average-001"].ContentDigest == "5b7a7b79dd6b2d355ad5901c0f81853b26d6d52cfbe7d8b5f59b708e8a012b68" {
		known["methods.values-expressions"] = true
	}
	return known
}
