package learner

import (
	"bytes"
	"encoding/csv"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func TestAdditiveTaxonomyPreservesLearnerHistoryAndWorkspace(t *testing.T) {
	repository, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	root := boundedTaxonomyHistoryRoot(t, repository)
	// Copy authored files only as data. Verification below is the explicit fast
	// service fixture; no task source or tests are executed in this regression.
	// All reviewed source-document bytes are deliberately absent below, including
	// the retained snapshots, so consumers cannot acquire a snapshot dependency.
	for _, relative := range []string{
		"curriculum/mappings/exercism-go.yaml",
		"curriculum/evidence/exercism-go.yaml",
		"catalog/reviews/exercism-go.yaml",
		"catalog/seed-plans/exercism-foundation-pilot.yaml",
		"curriculum/migrations/snapshots/exercism-map-v3.yaml",
		"curriculum/migrations/snapshots/exercism-evidence-v3.yaml",
		"curriculum/migrations/snapshots/exercism-reviews-v3.yaml",
		"curriculum/migrations/snapshots/exercism-seeds-v3.yaml",
	} {
		path := filepath.Join(root, filepath.FromSlash(relative))
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
		if _, err := os.Lstat(path); !os.IsNotExist(err) {
			t.Fatalf("reviewed bytes %s must be absent from learner fixture: %v", relative, err)
		}
	}
	live, err := os.ReadFile(filepath.Join(root, "curriculum/skills.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	old, err := os.ReadFile(filepath.Join(root, "curriculum/migrations/snapshots/skills-v3.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "curriculum/skills.yaml"), old, 0600); err != nil {
		t.Fatal(err)
	}
	home := t.TempDir()
	newService := func() *Service {
		s, err := New(Config{Root: root, Home: home, Now: func() time.Time { return time.Date(2026, 10, 18, 0, 0, 0, 0, time.UTC) }})
		if err != nil {
			t.Fatal(err)
		}
		s.verify = func(_, _, _ string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
			return tasking.LearnerVerificationResult{Passed: true, Scope: scope, Summary: "Synthetic service-policy fixture; no task execution"}, nil
		}
		return s
	}
	beforeService := newService()
	if _, err := beforeService.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	if _, err := beforeService.Check(); err != nil {
		t.Fatal(err)
	}
	if _, err := beforeService.Complete(); err != nil {
		t.Fatal(err)
	}
	workspace, err := beforeService.Start("zero-results-001")
	if err != nil {
		t.Fatal(err)
	}
	edited := false
	if err := filepath.WalkDir(workspace, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if edited || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		edited = true
		return os.WriteFile(path, append(b, []byte("\n// Preserve this learner edit byte for byte.\n")...), 0600)
	}); err != nil {
		t.Fatal(err)
	}
	if !edited {
		t.Fatal("no learner Go file edited in workspace fixture")
	}
	if err := beforeService.Pause(); err != nil {
		t.Fatal(err)
	}
	beforeChoice, err := beforeService.Next()
	if err != nil {
		t.Fatal(err)
	}
	beforeRoadmap, err := beforeService.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	if beforeChoice.Task.ID != "zero-results-001" || beforeChoice.LessonNumber != "01.02" || !beforeChoice.Resume || beforeChoice.PedagogicalPosition != 2 {
		t.Fatalf("fixture must resume the original second lesson: %+v", beforeChoice)
	}
	schedulingRoute, err := loadSchedulingRoute(root)
	if err != nil {
		t.Fatal(err)
	}
	if schedulingRoute == nil || !schedulingRoute.pedagogical || len(executionLessons(*schedulingRoute)) != 2 || len(beforeRoadmap.Tasks) != 2 || beforeRoadmap.Tasks["literal-bindings-001"].ID == "" || beforeRoadmap.Tasks["zero-results-001"].ID == "" {
		t.Fatal("fixture must use the active execution plan and exactly two original tasks")
	}
	beforeProgress, err := beforeService.Progress()
	if err != nil {
		t.Fatal(err)
	}
	if len(beforeProgress.Completions) != 1 || beforeProgress.CurrentTaskID != "zero-results-001" {
		t.Fatal("fixture must contain a historical completion and a different unfinished current task")
	}
	beforeHome := homeSnapshot(t, home)
	beforeMarker, err := os.ReadFile(filepath.Join(workspace, learnerWorkspaceMarker))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "curriculum/skills.yaml"), live, 0600); err != nil {
		t.Fatal(err)
	}
	afterService := newService()
	afterChoice, err := afterService.Next()
	if err != nil {
		t.Fatal(err)
	}
	afterRoadmap, err := afterService.Roadmap()
	if err != nil {
		t.Fatal(err)
	}
	afterProgress, err := afterService.Progress()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(beforeChoice, afterChoice) || !reflect.DeepEqual(beforeRoadmap, afterRoadmap) {
		t.Fatal("additive taxonomy changed current priority, route addresses/roles, readiness or recommendation")
	}
	if !reflect.DeepEqual(beforeProgress, afterProgress) || !reflect.DeepEqual(beforeHome, homeSnapshot(t, home)) {
		t.Fatal("read-only v4 load migrated progress, historical encounters/completions or learner files")
	}
	resumed, err := afterService.Resume("zero-results-001")
	if err != nil || resumed != workspace {
		t.Fatalf("resume=%q %v; want oldworkspace %q", resumed, err, workspace)
	}
	afterMarker, err := os.ReadFile(filepath.Join(workspace, learnerWorkspaceMarker))
	if err != nil || !bytes.Equal(beforeMarker, afterMarker) {
		t.Fatalf("unchanged task content acquired a new marker/digest: %v", err)
	}
	afterResumeHome := homeSnapshot(t, home)
	for path, content := range beforeHome {
		if strings.HasPrefix(path, "workspaces/") {
			if afterResumeHome[path] != content {
				t.Fatalf("resume rewrote learner workspace %s", path)
			}
		}
	}
	resumedProgress, err := afterService.Progress()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(beforeProgress.Completions, resumedProgress.Completions) || resumedProgress.CurrentTaskID != beforeProgress.CurrentTaskID {
		t.Fatal("resume retagged historical completion or changed current task")
	}
}

// The complete bank's task bytes, addresses and exposure rules have separate
// inventory/route regressions. This composed state case needs two real tasks:
// a historical completion and a different unfinished workspace. Keep the full
// taxonomy and its transition, but bound repeated catalog reads to those tasks.
func boundedTaxonomyHistoryRoot(t *testing.T, repository string) string {
	t.Helper()
	root := t.TempDir()
	for _, path := range []string{"curriculum", "schemas", "config", "sources", "docs", "theory"} {
		if err := isolation.CopyTree(filepath.Join(repository, path), filepath.Join(root, path)); err != nil {
			t.Fatal(err)
		}
	}
	selected := map[string]bool{"literal-bindings-001": true, "zero-results-001": true}
	for id := range selected {
		relative := filepath.Join("tasks", "ready", id)
		if err := isolation.CopyTree(filepath.Join(repository, relative), filepath.Join(root, relative)); err != nil {
			t.Fatal(err)
		}
		original, err := tasking.TaskContentDigest(filepath.Join(repository, relative))
		if err != nil {
			t.Fatal(err)
		}
		copied, err := tasking.TaskContentDigest(filepath.Join(root, relative))
		if err != nil || copied != original {
			t.Fatalf("fixture changed original task %s: %v", id, err)
		}
	}
	route, err := loadRouteMetadata(root)
	if err != nil || route == nil {
		t.Fatalf("original route: %v", err)
	}
	// The runtime loader resolves this path for display; the fixture file must
	// retain its original repository-relative guide reference.
	guide, err := filepath.Rel(root, route.Preparation)
	if err != nil {
		t.Fatal(err)
	}
	route.Preparation = filepath.ToSlash(guide)
	var blocks []RouteBlock
	for _, block := range route.Blocks {
		var steps []RouteStep
		for _, step := range block.Steps {
			if selected[step.TaskID] {
				steps = append(steps, step)
			}
		}
		if len(steps) > 0 {
			block.Steps = steps
			blocks = append(blocks, block)
		}
	}
	route.Blocks = blocks
	if len(blocks) != 1 || blocks[0].Number != 1 || len(blocks[0].Steps) != 2 || blocks[0].Steps[0].Number != 1 || blocks[0].Steps[1].Number != 2 {
		t.Fatal("fixture must retain the original first two addresses")
	}
	// Preserve the real preparation skills; these two adjacent early tasks need
	// no invented exposure assumptions. Keep schema v2 and active plan semantics.
	plan, _, err := loadPedagogicalPlan(root)
	if err != nil {
		t.Fatal(err)
	}
	if plan.SchemaVersion != 2 || plan.Status != "active" {
		t.Fatal("expected active v2 plan")
	}
	var edges []ExposureEdge
	for _, edge := range plan.ExposureEdges {
		if selected[edge.TaskID] && selected[edge.BeforeTaskID] {
			edges = append(edges, edge)
		}
	}
	plan.ExposureEdges = edges
	var pending []PendingPreparation
	for _, gate := range plan.PendingPreparation {
		if selected[gate.TaskID] {
			pending = append(pending, gate)
		}
	}
	plan.PendingPreparation = pending
	for path, value := range map[string]any{"curriculum/route.yaml": route, pedagogicalPlanPath: plan} {
		data, err := yaml.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, path), data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	path := filepath.Join(root, plan.MappingFile)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	rows, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil || len(rows) == 0 {
		t.Fatalf("active mapping: %v", err)
	}
	kept := [][]string{rows[0]}
	for _, row := range rows[1:] {
		if selected[row[0]] {
			kept = append(kept, row)
		}
	}
	if len(kept) != 3 || kept[1][0] != "literal-bindings-001" || kept[1][1] != "01.01" || kept[1][4] != "1" || kept[2][0] != "zero-results-001" || kept[2][1] != "01.02" || kept[2][4] != "2" {
		t.Fatal("fixture must retain the exact original mapping rows without renumbering")
	}
	var output bytes.Buffer
	writer := csv.NewWriter(&output)
	writer.UseCRLF = true
	if err := writer.WriteAll(kept); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, output.Bytes(), 0600); err != nil {
		t.Fatal(err)
	}
	return root
}
