package learner

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
)

func TestAdditiveTaxonomyPreservesLearnerHistoryAndWorkspace(t *testing.T) {
	repository, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	// Copy authored files only as data. Verification below is the explicit fast
	// service fixture; no task source or tests are executed in this regression.
	// All reviewed source-document bytes are deliberately absent below, including
	// the retained snapshots, so consumers cannot acquire a snapshot dependency.
	for _, path := range []string{"curriculum", "schemas", "config", "sources", "docs", "theory", "tasks/ready"} {
		if err := isolation.CopyTree(filepath.Join(repository, path), filepath.Join(root, path)); err != nil {
			t.Fatal(err)
		}
	}
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
