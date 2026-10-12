//go:build taskintegration

package tasking

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLearnerSessionKeepsFreshWorkspacesAndInputs(t *testing.T) {
	root := learnerCatalogFixture(t)
	const id = "zero-results-001"
	session, err := NewCheckSession(root)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := session.Close(); err != nil {
			t.Error(err)
		}
	})
	runner := session.runner.(*realExecutionRunner)
	environment := runner.environment.Root
	canonical := filepath.Join(root, "tasks/ready", id)
	before, err := TaskContentDigest(canonical)
	if err != nil {
		t.Fatal(err)
	}
	workspace := t.TempDir()
	files, err := LearnerWorkspaceFiles(root, id)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		ideaWrite(t, workspace, file.Path, file.Content)
	}

	check := func(scope LearnerVerificationScope, want bool) {
		t.Helper()
		result, err := session.VerifyLearnerWorkspace(id, workspace, scope)
		if err != nil || result.Passed != want {
			t.Fatalf("%s want pass=%t: %+v %v", scope, want, result, err)
		}
		if scope == LearnerFullChecks && result.Diagnostic != "" {
			t.Fatal("full verification leaked diagnostics")
		}
	}
	check(LearnerPublicChecks, false)
	// Record all old workspaces and plant a test that fails if a later run
	// reuses their directories. These are already disposable checker copies.
	seen := map[string]bool{}
	poisonOldCopies := func() {
		t.Helper()
		entries, err := os.ReadDir(filepath.Join(environment, "workspaces"))
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			seen[entry.Name()] = true
			path := filepath.Join(environment, "workspaces", entry.Name(), "subject/stale_test.go")
			mustWrite(t, path, []byte("package subject_test\nimport \"testing\"\nfunc TestStale(t *testing.T) { t.Fatal(\"reused workspace\") }\n"))
		}
	}
	poisonOldCopies()
	reference := mustRead(t, filepath.Join(canonical, "solution/reference/subject.go"))
	mustWrite(t, filepath.Join(workspace, "subject/subject.go"), reference)
	check(LearnerFullChecks, true)
	poisonOldCopies()
	check(LearnerPublicChecks, true)
	entries, err := os.ReadDir(filepath.Join(environment, "workspaces"))
	if err != nil {
		t.Fatal(err)
	}
	fresh := 0
	for _, entry := range entries {
		if seen[entry.Name()] {
			continue
		}
		fresh++
		// No private grader may survive the transition from full to public.
		paths, err := filepath.Glob(filepath.Join(environment, "workspaces", entry.Name(), "subject", "*"))
		if err != nil {
			t.Fatal(err)
		}
		if len(paths) != 2 {
			t.Fatalf("public workspace has unexpected files: %v", paths)
		}
		for _, path := range paths {
			if strings.Contains(filepath.Base(path), "grader") || strings.Contains(filepath.Base(path), "stale") {
				t.Fatalf("previous verification material leaked: %s", path)
			}
		}
	}
	if fresh != 1 {
		t.Fatalf("public run created %d workspaces, want 1", fresh)
	}
	metrics := runner.Metrics()
	if metrics.Sessions != 1 || metrics.ChildGoRuns != 5 || metrics.Workspaces != len(entries) || metrics.ChildGoRuns != metrics.Workspaces || metrics.VariantRuns != metrics.ChildGoRuns {
		t.Fatalf("execution was skipped or reused: %+v directories=%d", metrics, len(entries))
	}
	// The same session must reject changed read-only files without execution.
	goMod := mustRead(t, filepath.Join(workspace, "go.mod"))
	mustWrite(t, filepath.Join(workspace, "go.mod"), []byte("module changed\n"))
	if _, err := session.VerifyLearnerWorkspace(id, workspace, LearnerPublicChecks); err == nil {
		t.Fatal("accepted changed go.mod")
	}
	mustWrite(t, filepath.Join(workspace, "go.mod"), goMod)
	metadata := filepath.Join(canonical, "review/review.yaml")
	original := mustRead(t, metadata)
	mustWrite(t, metadata, []byte("unknown_field: true\n"))
	if _, err := session.VerifyLearnerWorkspace(id, workspace, LearnerPublicChecks); err == nil {
		t.Fatal("accepted stale review metadata")
	}
	mustWrite(t, metadata, original)
	if _, err := session.VerifyLearnerWorkspace(id, workspace, "invalid"); err == nil {
		t.Fatal("accepted invalid scope")
	}
	if runner.Metrics() != metrics {
		t.Fatal("static rejection executed code")
	}
	after, err := TaskContentDigest(canonical)
	if err != nil || before != after {
		t.Fatal("canonical content changed")
	}
	if err := session.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(environment); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("session files survived close: %v", err)
	}
	if _, err := session.VerifyLearnerWorkspace(id, workspace, LearnerPublicChecks); err == nil {
		t.Fatal("closed session accepted work")
	}
}
