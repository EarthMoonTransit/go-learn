//go:build taskintegration

package learner

import (
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/tasking"
)

func TestRealOpeningLessonRunsThroughLearnerWorkspace(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	service, err := New(Config{Root: root, Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	first, err := service.Next()
	if err != nil || first.Task.ID != "literal-bindings-001" {
		t.Fatalf("first=%+v err=%v", first, err)
	}
	canonical := filepath.Join(root, "tasks", "ready", first.Task.ID)
	before, err := tasking.TaskContentDigest(canonical)
	if err != nil {
		t.Fatal(err)
	}
	workspace, err := service.Start(first.Task.ID)
	if err != nil {
		t.Fatal(err)
	}
	next, err := service.Next()
	if err != nil || !next.Resume {
		t.Fatalf("next=%+v err=%v", next, err)
	}
	result, err := service.Check()
	if err != nil || result.Passed {
		t.Fatalf("starter=%+v err=%v", result, err)
	}
	// Read canonical reference as data; execution occurs only in the learner
	// runner's disposable copy, never in the canonical source tree.
	content, err := os.ReadFile(filepath.Join(canonical, "solution", "reference", "subject", "bindings.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(workspace, "subject", "bindings.go"), content, 0600); err != nil {
		t.Fatal(err)
	}
	result, err = service.Complete()
	if err != nil || !result.Passed {
		t.Fatalf("complete=%+v err=%v", result, err)
	}
	next, err = service.Next()
	if err != nil || next.Task.ID != "zero-results-001" {
		t.Fatalf("next=%+v err=%v", next, err)
	}
	progress, err := service.Progress()
	if err != nil || !progress.Tasks[first.Task.ID].Completed || progress.Tasks[first.Task.ID].Attempts != 2 {
		t.Fatalf("progress=%+v err=%v", progress, err)
	}
	after, err := tasking.TaskContentDigest(canonical)
	if err != nil || before != after {
		t.Fatal("canonical task changed")
	}
}
