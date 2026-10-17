package learner

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func syntheticPedagogicalPlan() (PedagogicalPlan, []PedagogicalEntry, Roadmap, map[string]SkillMetadata) {
	plan := PedagogicalPlan{SchemaVersion: 1, Status: "inactive", BaselineCommit: strings.Repeat("a", 40), MappingFile: "mapping.csv"}
	entries := []PedagogicalEntry{
		{TaskID: "early-task-001", OldNumber: "01.01", Chapter: "C01", Subtopic: "Basic", ProposedOrder: 1, Role: "introduction", Disposition: "core"},
		{TaskID: "later-task-001", OldNumber: "01.02", Chapter: "C02", Subtopic: "Later", ProposedOrder: 2, Role: "practice", Disposition: "reinforcement"},
	}
	catalog, _ := routeCatalog(numberedTestRoute(), testRouteCatalog())
	roadmap := Roadmap{Route: numberedTestRoute(), Tasks: catalog}
	skills := map[string]SkillMetadata{"setup": {}, "basic": {Prerequisites: []string{"setup"}}, "later": {Prerequisites: []string{"basic"}}}
	return plan, entries, roadmap, skills
}

func TestPedagogicalMappingRejectsStructuralDrift(t *testing.T) {
	for _, tc := range []struct {
		name, diagnostic string
		change           func(*PedagogicalPlan, *[]PedagogicalEntry, *Roadmap)
	}{
		{"unknown", "unknown mapped task ID proposal-p01-001", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[0].TaskID = "proposal-p01-001" }},
		{"duplicate", "duplicate mapped task ID early-task-001", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[1].TaskID = "early-task-001" }},
		{"missing", "missing current route task ID later-task-001", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { *e = (*e)[:1] }},
		{"new route task", "missing current route task ID new-task-001", func(_ *PedagogicalPlan, _ *[]PedagogicalEntry, r *Roadmap) {
			r.Route.Blocks[0].Steps = append(r.Route.Blocks[0].Steps, RouteStep{TaskID: "new-task-001", Number: 3, Role: "practice", Skills: []string{"basic"}})
		}},
		{"old address", "old_number", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[0].OldNumber = "99.01" }},
		{"role", "differs from current role", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[0].Role = "transfer" }},
		{"zero position", "outside contiguous range", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[0].ProposedOrder = 0 }},
		{"gap", "missing proposed_order 2", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[1].ProposedOrder = 3 }},
		{"duplicate position", "duplicate proposed_order 1", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[1].ProposedOrder = 1 }},
		{"unavailable", "no ready v2 learner task", func(_ *PedagogicalPlan, _ *[]PedagogicalEntry, r *Roadmap) { delete(r.Tasks, "early-task-001") }},
		{"missing chapter", "proposed chapter and subtopic", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[0].Chapter = " " }},
		{"invalid disposition", "unsupported disposition", func(_ *PedagogicalPlan, e *[]PedagogicalEntry, _ *Roadmap) { (*e)[0].Disposition = "optional" }},
		{"proposal as edge", "invalid or duplicate exposure edge", func(p *PedagogicalPlan, _ *[]PedagogicalEntry, _ *Roadmap) {
			p.ExposureEdges = []ExposureEdge{{"P05", "later-task-001"}}
		}},
		{"self edge", "invalid or duplicate exposure edge", func(p *PedagogicalPlan, _ *[]PedagogicalEntry, _ *Roadmap) {
			p.ExposureEdges = []ExposureEdge{{"later-task-001", "later-task-001"}}
		}},
		{"duplicate edge", "invalid or duplicate exposure edge", func(p *PedagogicalPlan, _ *[]PedagogicalEntry, _ *Roadmap) {
			p.ExposureEdges = []ExposureEdge{{"early-task-001", "later-task-001"}, {"early-task-001", "later-task-001"}}
		}},
		{"missing gate reason", "invalid or duplicate pending preparation", func(p *PedagogicalPlan, _ *[]PedagogicalEntry, _ *Roadmap) {
			p.PendingPreparation = []PendingPreparation{{TaskID: "later-task-001", ReviewID: "P13"}}
		}},
		{"duplicate pending", "invalid or duplicate pending preparation", func(p *PedagogicalPlan, _ *[]PedagogicalEntry, _ *Roadmap) {
			p.PendingPreparation = []PendingPreparation{{"later-task-001", "P13", "scaffold revision"}, {"later-task-001", "P13", "scaffold revision"}}
		}},
		{"legacy route", "requires the stable numbered route v2", func(_ *PedagogicalPlan, _ *[]PedagogicalEntry, r *Roadmap) { r.Route.SchemaVersion = 1 }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p, e, r, skills := syntheticPedagogicalPlan()
			tc.change(&p, &e, &r)
			preview := validatePedagogicalOrder(p, e, r, skills)
			if preview.Valid() || preview.PreparationReady() || !strings.Contains(strings.Join(preview.StructuralIssues, "\n"), tc.diagnostic) {
				t.Fatalf("want %q; got %+v", tc.diagnostic, preview)
			}
			for _, lesson := range preview.Lessons {
				if len(lesson.FormalBlockers) != 0 {
					t.Fatal("invalid sequence was analyzed as a valid exposure plan")
				}
			}
		})
	}
}

func TestPedagogicalExposureUsesManifestTaxonomyAndReviewedEdges(t *testing.T) {
	p, e, r, skills := syntheticPedagogicalPlan()
	// CSV physical order is irrelevant; explicit proposed_order is the specification.
	slices.Reverse(e)
	if preview := validatePedagogicalOrder(p, e, r, skills); !preview.Valid() || !preview.PreparationReady() || preview.Lessons[0].TaskID != "early-task-001" {
		t.Fatal(preview)
	}
	for _, tc := range []struct {
		name, blocker string
		change        func(*PedagogicalPlan, *Roadmap, map[string]SkillMetadata)
	}{
		{"manifest", "manifest requires earlier later", func(_ *PedagogicalPlan, r *Roadmap, _ map[string]SkillMetadata) {
			task := r.Tasks["early-task-001"]
			task.Prerequisites = append(task.Prerequisites, "later")
			r.Tasks[task.ID] = task
		}},
		{"taxonomy", "taxonomy basic requires earlier later", func(_ *PedagogicalPlan, _ *Roadmap, skills map[string]SkillMetadata) {
			skills["basic"] = SkillMetadata{Prerequisites: []string{"later"}}
		}},
		{"reviewed edge", "reviewed exposure requires earlier task later-task-001", func(p *PedagogicalPlan, _ *Roadmap, _ map[string]SkillMetadata) {
			p.ExposureEdges = []ExposureEdge{{"later-task-001", "early-task-001"}}
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p, e, r, skills := syntheticPedagogicalPlan()
			tc.change(&p, &r, skills)
			preview := validatePedagogicalOrder(p, e, r, skills)
			if !preview.Valid() || preview.PreparationReady() || !slices.Contains(preview.Lessons[0].FormalBlockers, tc.blocker) {
				t.Fatal(preview)
			}
		})
	}
	// The route permits a dependent assessed skill and its introduction in one task.
	r.Route.Blocks[0].Steps[0].Skills = []string{"basic", "paired"}
	skills["basic"] = SkillMetadata{Prerequisites: []string{"paired"}}
	skills["paired"] = SkillMetadata{Prerequisites: []string{"setup"}}
	if preview := validatePedagogicalOrder(p, e, r, skills); !preview.PreparationReady() {
		t.Fatal(preview)
	}
	// Manifest prerequisites still require an earlier encounter, not a same-task tag.
	task := r.Tasks["early-task-001"]
	task.Prerequisites = []string{"paired"}
	r.Tasks[task.ID] = task
	if preview := validatePedagogicalOrder(p, e, r, skills); !slices.Contains(preview.Lessons[0].FormalBlockers, "manifest requires earlier paired") {
		t.Fatal(preview)
	}
}

func TestPedagogicalLoaderIsStrictAndInactive(t *testing.T) {
	p, _, _, _ := syntheticPedagogicalPlan()
	content, err := yaml.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	validCSV := "\xef\xbb\xbftask_id,old_number,proposed_chapter,proposed_subtopic,proposed_order,role,disposition,rationale\r\nearly-task-001,01.01,C01,Basic,1,introduction,core,\"quoted, rationale\"\r\n"
	for _, tc := range []struct {
		name, yaml, csv string
		valid           bool
	}{
		{"BOM CRLF quoted CSV", string(content), validCSV, true},
		{"active", strings.Replace(string(content), "inactive", "active", 1), validCSV, false},
		{"missing status", strings.Replace(string(content), "status: inactive\n", "", 1), validCSV, false},
		{"version", strings.Replace(string(content), "schema_version: 1", "schema_version: 2", 1), validCSV, false},
		{"short commit", strings.Replace(string(content), p.BaselineCommit, "aaaa", 1), validCSV, false},
		{"unknown YAML field", string(content) + "scheduler: enabled\n", validCSV, false},
		{"duplicate YAML field", string(content) + "status: inactive\n", validCSV, false},
		{"multiple documents", string(content) + "---\nstatus: inactive\n", validCSV, false},
		{"trailing malformed YAML", string(content) + "---\n[", validCSV, false},
		{"mapping traversal", strings.Replace(string(content), "mapping.csv", "../mapping.csv", 1), validCSV, false},
		{"unknown CSV column", string(content), strings.Replace(validCSV, "rationale", "scheduler", 1), false},
		{"short CSV row", string(content), strings.Replace(validCSV, "core,\"quoted, rationale\"", "core", 1), false},
		{"invalid order", string(content), strings.Replace(validCSV, ",1,introduction", ",one,introduction", 1), false},
		{"invalid CSV quote", string(content), validCSV + "\"unterminated\n", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			if err := os.MkdirAll(filepath.Join(root, "curriculum/plans"), 0700); err != nil {
				t.Fatal(err)
			}
			for file, value := range map[string]string{pedagogicalPlanPath: tc.yaml, "mapping.csv": tc.csv} {
				if err := os.WriteFile(filepath.Join(root, file), []byte(value), 0600); err != nil {
					t.Fatal(err)
				}
			}
			_, entries, err := loadPedagogicalPlan(root)
			if (err == nil) != tc.valid {
				t.Fatalf("entries=%+v err=%v", entries, err)
			}
			if tc.valid && (len(entries) != 1 || entries[0].Rationale != "quoted, rationale") {
				t.Fatal(entries)
			}
		})
	}
}

func TestRepositoryPedagogicalMappingAndKnownBlockers(t *testing.T) {
	s := uiTestService(t, "")
	p, entries, err := loadPedagogicalPlan(s.root)
	if err != nil {
		t.Fatal(err)
	}
	// Supplied route-mapping.csv from the review, copied byte-for-byte.
	content, err := os.ReadFile(filepath.Join(s.root, p.MappingFile))
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%x", sha256.Sum256(content)) != "d0454a41c8c0a4d4cd7c6fb4c36e14568a15b89444f4157871efcba921367d75" {
		t.Fatal("review mapping changed; update its planning evidence explicitly")
	}
	roadmap, err := s.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	skills, err := activeRouteSkills(s.root)
	if err != nil {
		t.Fatal(err)
	}
	preview := validatePedagogicalOrder(p, entries, roadmap, skills)
	if !preview.Valid() || preview.PreparationReady() || len(entries) != 151 || preview.RoutedCount != 151 || p.Status != "inactive" {
		t.Fatalf("preview=%+v", preview)
	}
	formal := map[string][]string{}
	pending := map[string][]string{}
	ids := map[string]bool{}
	for i, lesson := range preview.Lessons {
		if ids[lesson.TaskID] || lesson.ProposedOrder != i+1 {
			t.Fatal("mapping lost unique explicit positions")
		}
		ids[lesson.TaskID] = true
		if len(lesson.FormalBlockers) > 0 {
			formal[lesson.TaskID] = lesson.FormalBlockers
		}
		for _, gate := range lesson.PendingPreparation {
			pending[lesson.TaskID] = append(pending[lesson.TaskID], gate.ReviewID)
		}
	}
	if len(formal) != 0 {
		t.Fatalf("E01/M1 must resolve the early testing blockers: %v", formal)
	}
	if !slices.Contains(roadmap.Tasks["clamp-tests-001"].SecondarySkills, "testing.basics-diagnostics") {
		t.Fatal("Clamp's manifest must match its reviewed route secondary")
	}
	for id, want := range map[string][]string{
		"clamp-tests-001": nil, "prepend-values-tests-001": {"E02"}, "shape-contract-tests-001": {"P05"},
		"pointer-chain-001": nil, "diagnostic-helper-001": nil, "named-cases-001": nil,
		"detached-append-001": {"M1"},
		"capture-panic-001":   {"P06", "P07"}, "worker-completion-tests-001": {"P11", "P12", "P13"},
	} {
		if !reflect.DeepEqual(pending[id], want) {
			t.Fatalf("%s gates=%v want=%v", id, pending[id], want)
		}
	}
	for _, id := range []string{"defined-types-001", "zero-state-001", "variadic-sum-001", "formatter-contract-001", "resource-finalize-001", "typed-nil-resource-001", "P13"} {
		if ids[id] {
			t.Fatalf("legacy/draft/rejected/proposal entered mapping: %s", id)
		}
	}
	// A real route addition is reported, never discarded to force the baseline count.
	roadmap.Route.Blocks[8].Steps = append(roadmap.Route.Blocks[8].Steps, RouteStep{Number: 10, TaskID: "added-task-001", Role: "practice", Skills: []string{"types.zero-values"}})
	changed := validatePedagogicalOrder(p, entries, roadmap, skills)
	if changed.RoutedCount != 152 || changed.Valid() || !slices.Contains(changed.StructuralIssues, "missing current route task ID added-task-001") {
		t.Fatal(changed)
	}
	// Check all stable block IDs/numbers, lesson numbers, roles and IDs independently
	// of the mapping; these tuples are captured from the exact review baseline.
	baseline, err := os.ReadFile("testdata/review-baseline-addresses.tsv")
	if err != nil {
		t.Fatal(err)
	}
	var current bytes.Buffer
	for _, block := range roadmap.Route.Blocks {
		for _, step := range block.Steps {
			if step.TaskID != "added-task-001" {
				fmt.Fprintf(&current, "%s\t%s\t%s\t%s\n", lessonNumber(block, step), block.ID, step.TaskID, step.Role)
			}
		}
	}
	if !bytes.Equal(baseline, current.Bytes()) {
		t.Fatal("stable review-baseline address registry changed")
	}
}

func TestPedagogicalMappingExcludesNonRoutedLifecycleTasks(t *testing.T) {
	for _, id := range []string{"formatter-contract-001", "typed-nil-resource-001", "defined-types-001", "zero-state-001", "variadic-sum-001"} {
		p, entries, r, skills := syntheticPedagogicalPlan()
		entries[0].TaskID = id
		// Even a ready catalog entry does not give a nonrouted task a course position.
		r.Tasks[id] = tasking.LearnerTaskInfo{ID: id}
		preview := validatePedagogicalOrder(p, entries, r, skills)
		if preview.Valid() || !slices.Contains(preview.StructuralIssues, "unknown mapped task ID "+id+": not in current route") {
			t.Fatal(preview)
		}
	}
}
