package learner

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"go-practice-bank/internal/safeio"
	"gopkg.in/yaml.v3"
)

const pedagogicalPlanPath = "curriculum/plans/pedagogical-order.yaml"

// PedagogicalPlan is inactive planning input, never a scheduler or address registry.
type PedagogicalPlan struct {
	SchemaVersion      int                  `yaml:"schema_version"`
	Status             string               `yaml:"status"`
	BaselineCommit     string               `yaml:"baseline_commit"`
	MappingFile        string               `yaml:"mapping_file"`
	ExposureEdges      []ExposureEdge       `yaml:"exposure_edges"`
	PendingPreparation []PendingPreparation `yaml:"pending_preparation"`
}

type ExposureEdge struct {
	BeforeTaskID string `yaml:"before_task_id"`
	TaskID       string `yaml:"task_id"`
}

type PendingPreparation struct {
	TaskID   string `yaml:"task_id"`
	ReviewID string `yaml:"review_id"`
	Reason   string `yaml:"reason"`
}

type PedagogicalEntry struct {
	TaskID        string
	OldNumber     string
	Chapter       string
	Subtopic      string
	ProposedOrder int
	Role          string
	Disposition   string
	Rationale     string
}

type PedagogicalLesson struct {
	PedagogicalEntry
	LearnerStatus      string
	Prerequisites      []string
	FormalBlockers     []string
	PendingPreparation []PendingPreparation
}

type PedagogicalPreview struct {
	Plan             PedagogicalPlan
	RoutedCount      int
	StructuralIssues []string
	Lessons          []PedagogicalLesson
}

func (p PedagogicalPreview) Valid() bool { return len(p.StructuralIssues) == 0 }

// PreparationReady is a static planning result, not learner readiness or activation.
func (p PedagogicalPreview) PreparationReady() bool {
	if !p.Valid() {
		return false
	}
	for _, lesson := range p.Lessons {
		if len(lesson.FormalBlockers) > 0 || len(lesson.PendingPreparation) > 0 {
			return false
		}
	}
	return true
}

func loadPedagogicalPlan(root string) (PedagogicalPlan, []PedagogicalEntry, error) {
	content, err := safeio.ReadFile(root, pedagogicalPlanPath, 1<<20)
	if err != nil {
		return PedagogicalPlan{}, nil, fmt.Errorf("pedagogical plan: %w", err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)
	var plan PedagogicalPlan
	if err := decoder.Decode(&plan); err != nil {
		return plan, nil, fmt.Errorf("pedagogical plan: %w", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return plan, nil, errors.New("pedagogical plan must contain exactly one document")
	}
	if plan.SchemaVersion != 1 || plan.Status != "inactive" {
		return plan, nil, errors.New("pedagogical plan requires schema_version 1 and explicit status inactive; activation is unsupported")
	}
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(plan.BaselineCommit) {
		return plan, nil, errors.New("pedagogical plan requires an exact baseline commit")
	}
	content, err = safeio.ReadFile(root, plan.MappingFile, 1<<20)
	if err != nil {
		return plan, nil, fmt.Errorf("pedagogical mapping: %w", err)
	}
	entries, err := parsePedagogicalMapping(content)
	return plan, entries, err
}

func parsePedagogicalMapping(content []byte) ([]PedagogicalEntry, error) {
	reader := csv.NewReader(bytes.NewReader(bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))))
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("pedagogical mapping header: %w", err)
	}
	if !slices.Equal(header, []string{"task_id", "old_number", "proposed_chapter", "proposed_subtopic", "proposed_order", "role", "disposition", "rationale"}) {
		return nil, errors.New("pedagogical mapping requires the review CSV columns in their original order")
	}
	var entries []PedagogicalEntry
	for row := 2; ; row++ {
		record, err := reader.Read()
		if err == io.EOF {
			return entries, nil
		}
		if err != nil {
			return nil, fmt.Errorf("pedagogical mapping row %d: %w", row, err)
		}
		order, err := strconv.Atoi(record[4])
		if err != nil {
			return nil, fmt.Errorf("pedagogical mapping row %d: invalid proposed_order %q", row, record[4])
		}
		entries = append(entries, PedagogicalEntry{TaskID: record[0], OldNumber: record[1], Chapter: record[2], Subtopic: record[3], ProposedOrder: order, Role: record[5], Disposition: record[6], Rationale: record[7]})
	}
}

// PedagogicalPreview reuses the current read-only roadmap and actual catalog.
// Only this explicit entry point loads the sidecar; Next and UI remain unchanged.
func (s *Service) PedagogicalPreview() (PedagogicalPreview, error) {
	plan, entries, err := loadPedagogicalPlan(s.root)
	if err != nil {
		return PedagogicalPreview{}, err
	}
	roadmap, err := s.Roadmap()
	if err != nil {
		return PedagogicalPreview{}, err
	}
	skills, err := activeRouteSkills(s.root)
	if err != nil {
		return PedagogicalPreview{}, err
	}
	return validatePedagogicalOrder(plan, entries, roadmap, skills), nil
}

func validatePedagogicalOrder(plan PedagogicalPlan, entries []PedagogicalEntry, roadmap Roadmap, skills map[string]SkillMetadata) PedagogicalPreview {
	preview := PedagogicalPreview{Plan: plan}
	issue := func(format string, args ...any) {
		preview.StructuralIssues = append(preview.StructuralIssues, fmt.Sprintf(format, args...))
	}
	if roadmap.Route.SchemaVersion != 2 {
		issue("pedagogical preview requires the stable numbered route v2")
	}
	type address struct {
		number string
		step   RouteStep
	}
	routed := map[string]address{}
	for _, block := range roadmap.Route.Blocks {
		for _, step := range block.Steps {
			routed[step.TaskID] = address{lessonNumber(block, step), step}
			preview.RoutedCount++
		}
	}
	positions := map[string]int{}
	orders := map[int]string{}
	for _, entry := range entries {
		if _, ok := positions[entry.TaskID]; ok {
			issue("duplicate mapped task ID %s", entry.TaskID)
		}
		positions[entry.TaskID] = entry.ProposedOrder
		if other, ok := orders[entry.ProposedOrder]; ok {
			issue("duplicate proposed_order %d for %s and %s", entry.ProposedOrder, other, entry.TaskID)
		}
		orders[entry.ProposedOrder] = entry.TaskID
		if entry.ProposedOrder < 1 || entry.ProposedOrder > len(entries) {
			issue("%s: proposed_order %d outside contiguous range 1..%d", entry.TaskID, entry.ProposedOrder, len(entries))
		}
		old, ok := routed[entry.TaskID]
		if !ok {
			issue("unknown mapped task ID %s: not in current route", entry.TaskID)
		} else {
			if entry.OldNumber != old.number {
				issue("%s: old_number %q differs from current address %s", entry.TaskID, entry.OldNumber, old.number)
			}
			if entry.Role != old.step.Role {
				issue("%s: role %q differs from current role %s", entry.TaskID, entry.Role, old.step.Role)
			}
		}
		if _, ok := roadmap.Tasks[entry.TaskID]; !ok {
			issue("%s: no ready v2 learner task (draft/rejected/legacy tasks cannot enter this plan)", entry.TaskID)
		}
		if strings.TrimSpace(entry.Chapter) == "" || strings.TrimSpace(entry.Subtopic) == "" {
			issue("%s: proposed chapter and subtopic are required", entry.TaskID)
		}
		if entry.Disposition != "core" && entry.Disposition != "reinforcement" {
			issue("%s: unsupported disposition %q", entry.TaskID, entry.Disposition)
		}
	}
	// Deterministic diagnostics retain every missing current ID, including additions.
	missing := []string{}
	for id := range routed {
		if _, ok := positions[id]; !ok {
			missing = append(missing, id)
		}
	}
	sort.Strings(missing)
	for _, id := range missing {
		issue("missing current route task ID %s", id)
	}
	for order := 1; order <= len(entries); order++ {
		if _, ok := orders[order]; !ok {
			issue("missing proposed_order %d", order)
		}
	}
	gateKeys := map[string]bool{}
	for _, edge := range plan.ExposureEdges {
		key := edge.BeforeTaskID + "/" + edge.TaskID
		_, beforeOK := routed[edge.BeforeTaskID]
		_, targetOK := routed[edge.TaskID]
		if !beforeOK || !targetOK || edge.BeforeTaskID == edge.TaskID || gateKeys[key] {
			issue("invalid or duplicate exposure edge %s -> %s; both must be distinct current route task IDs", edge.BeforeTaskID, edge.TaskID)
		}
		gateKeys[key] = true
	}
	for _, pending := range plan.PendingPreparation {
		key := pending.TaskID + "/" + pending.ReviewID
		_, targetOK := routed[pending.TaskID]
		if !targetOK || strings.TrimSpace(pending.ReviewID) == "" || strings.TrimSpace(pending.Reason) == "" || gateKeys[key] {
			issue("invalid or duplicate pending preparation %s for %s", pending.ReviewID, pending.TaskID)
		}
		gateKeys[key] = true
	}
	ordered := slices.Clone(entries)
	sort.SliceStable(ordered, func(i, j int) bool { return ordered[i].ProposedOrder < ordered[j].ProposedOrder })
	known := map[string]bool{}
	for _, id := range roadmap.Route.PreparationSkills {
		known[id] = true
	}
	for _, entry := range ordered {
		lesson := PedagogicalLesson{PedagogicalEntry: entry, LearnerStatus: "not-started"}
		if progress := roadmap.Progress.Tasks[entry.TaskID]; progress != nil {
			lesson.LearnerStatus = progress.Status
			if progress.Completed {
				lesson.LearnerStatus = "completed"
			}
		}
		lesson.Prerequisites = slices.Clone(roadmap.Tasks[entry.TaskID].Prerequisites)
		for _, pending := range plan.PendingPreparation {
			if pending.TaskID == entry.TaskID {
				lesson.PendingPreparation = append(lesson.PendingPreparation, pending)
			}
		}
		// Do not compute reassuring exposure results from a malformed sequence.
		if preview.Valid() {
			for _, id := range lesson.Prerequisites {
				if !known[id] {
					lesson.FormalBlockers = append(lesson.FormalBlockers, "manifest requires earlier "+id)
				}
			}
			assessed := routed[entry.TaskID].step.Skills
			for _, id := range assessed {
				for _, prerequisite := range skills[id].Prerequisites {
					if !known[prerequisite] && !hasRouteSkill(assessed, prerequisite) {
						lesson.FormalBlockers = append(lesson.FormalBlockers, "taxonomy "+id+" requires earlier "+prerequisite)
					}
				}
			}
			for _, edge := range plan.ExposureEdges {
				if edge.TaskID == entry.TaskID && positions[edge.BeforeTaskID] >= entry.ProposedOrder {
					lesson.FormalBlockers = append(lesson.FormalBlockers, "reviewed exposure requires earlier task "+edge.BeforeTaskID)
				}
			}
			for _, id := range assessed {
				known[id] = true
			}
		}
		preview.Lessons = append(preview.Lessons, lesson)
	}
	return preview
}

func writePedagogicalPreview(output io.Writer, preview PedagogicalPreview) {
	fmt.Fprintf(output, "Pedagogical plan: %s (activation disabled)\nBaseline reference: %s\nMapping: %d rows; current route: %d tasks\n", preview.Plan.Status, preview.Plan.BaselineCommit, len(preview.Lessons), preview.RoutedCount)
	fmt.Fprintf(output, "Structural validity: %t; preparation ready: %t\n", preview.Valid(), preview.PreparationReady())
	fmt.Fprintln(output, "Exposure checks use the proposed sequence, not learner mastery. Next still uses the existing route.")
	for _, issue := range preview.StructuralIssues {
		fmt.Fprintf(output, "INVALID: %s\n", issue)
	}
	for _, lesson := range preview.Lessons {
		fmt.Fprintf(output, "\nPosition %d | %s | %s\n  Address %s | %s | %s | %s\n", lesson.ProposedOrder, lesson.Chapter, lesson.Subtopic, lesson.OldNumber, lesson.TaskID, lesson.Role, lesson.LearnerStatus)
		if len(lesson.Prerequisites) > 0 {
			fmt.Fprintf(output, "  Manifest prerequisites: %s\n", strings.Join(lesson.Prerequisites, ", "))
		}
		for _, blocker := range lesson.FormalBlockers {
			fmt.Fprintf(output, "  BLOCKED: %s\n", blocker)
		}
		for _, pending := range lesson.PendingPreparation {
			fmt.Fprintf(output, "  PENDING %s: %s\n", pending.ReviewID, pending.Reason)
		}
	}
}
