//go:build taskintegration

package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
)

func TestLearnerWorkflowSupportsEveryV2SubmissionKind(t *testing.T) {
	repository := learnerFixtureRepository(t)
	fixtures := []struct {
		profile             string
		id                  string
		learnerPath         string
		referencePath       string
		starterPublicPasses bool
	}{
		{"fixed", "synthetic-fixed-v2-001", "subject/identity.go", "solution/reference/identity.go", false},
		{"mutation", "synthetic-mutation-v2-001", "subject/identity_test.go", "solution/reference/identity_test.go", true},
		{"workspace", "synthetic-workspace-v2-001", "core/apply.go", "solution/reference/core/apply.go", false},
	}
	for _, fixture := range fixtures {
		t.Run(fixture.profile, func(t *testing.T) {
			clock := time.Date(2026, 9, 4, 10, 0, 0, 0, time.UTC)
			home := t.TempDir()
			service, err := New(Config{Root: repository, Home: home, Now: func() time.Time { return clock }})
			if err != nil {
				t.Fatal(err)
			}
			taskDirectory := filepath.Join(repository, "tasks", "ready", fixture.id)
			before, err := tasking.TaskContentDigest(taskDirectory)
			if err != nil {
				t.Fatal(err)
			}
			workspace, err := service.Start(fixture.id)
			if err != nil {
				t.Fatal(err)
			}
			assertLearnerWorkspaceSafe(t, workspace)
			if fixture.profile == "mutation" {
				api, err := os.ReadFile(filepath.Join(workspace, "subject", "learn_api.go"))
				if err != nil || !strings.Contains(string(api), "func Identity(") || !strings.Contains(string(api), "panic(") {
					t.Fatalf("missing editor API declaration: %q err=%v", api, err)
				}
			}
			if _, err := service.Solution(false); err == nil {
				t.Fatal("solution was revealed without an explicit action")
			}
			clock = clock.Add(90 * time.Second)
			public, err := service.Check()
			if err != nil {
				t.Fatal(err)
			}
			if public.Passed != fixture.starterPublicPasses {
				t.Fatalf("starter public result = %t, want %t: %+v", public.Passed, fixture.starterPublicPasses, public)
			}
			full, err := service.Complete()
			if err != nil {
				t.Fatal(err)
			}
			if full.Passed || full.Diagnostic != "" || strings.Contains(strings.ToLower(full.Summary), "mutant") {
				t.Fatalf("starter full result leaked or passed: %+v", full)
			}
			if _, _, err := service.Hint(); err != nil {
				t.Fatal(err)
			}
			contextPath, err := service.ExportReviewContext()
			if err != nil {
				t.Fatal(err)
			}
			assertReviewContextSafe(t, contextPath)

			reference, err := os.ReadFile(filepath.Join(taskDirectory, filepath.FromSlash(fixture.referencePath)))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, filepath.FromSlash(fixture.learnerPath)), reference, 0o600); err != nil {
				t.Fatal(err)
			}
			clock = clock.Add(30 * time.Second)
			if _, err := service.Resume(fixture.id); err != nil {
				t.Fatal(err)
			}
			clock = clock.Add(60 * time.Second)
			full, err = service.Complete()
			if err != nil {
				t.Fatal(err)
			}
			if !full.Passed || full.Diagnostic != "" {
				t.Fatalf("reference-equivalent learner submission failed or leaked: %+v", full)
			}
			if _, err := service.Review(); err != nil {
				t.Fatal(err)
			}
			if err := service.Feedback("public diagnostic was useful"); err != nil {
				t.Fatal(err)
			}
			progress, err := service.Progress()
			if err != nil {
				t.Fatal(err)
			}
			item := progress.Tasks[fixture.id]
			if !item.Completed || item.Attempts != 3 || item.ActivitySeconds != 180 || len(item.SkillEncounters) == 0 || len(item.Feedback) != 1 {
				t.Fatalf("incomplete learner progress: %+v", item)
			}
			after, err := tasking.TaskContentDigest(taskDirectory)
			if err != nil || after != before {
				t.Fatalf("canonical task changed: before=%s after=%s err=%v", before, after, err)
			}
		})
	}
}

func TestLearnerSolutionRevealDoesNotCompleteAndResetIsExplicit(t *testing.T) {
	repository := learnerFixtureRepository(t)
	service, err := New(Config{Root: repository, Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	workspace, err := service.Start("synthetic-fixed-v2-001")
	if err != nil {
		t.Fatal(err)
	}
	before := workspacePaths(t, workspace)
	files, err := service.Solution(true)
	if err != nil || len(files) == 0 {
		t.Fatalf("explicit solution reveal failed: files=%d err=%v", len(files), err)
	}
	if after := workspacePaths(t, workspace); strings.Join(after, "\n") != strings.Join(before, "\n") {
		t.Fatalf("solution reveal persisted checker material: before=%v after=%v", before, after)
	}
	progress, err := service.Progress()
	if err != nil {
		t.Fatal(err)
	}
	if item := progress.Tasks["synthetic-fixed-v2-001"]; !item.SolutionRevealed || item.Completed {
		t.Fatalf("solution reveal changed completion incorrectly: %+v", item)
	}
	if err := service.ResetTask("synthetic-fixed-v2-001", false); err == nil {
		t.Fatal("reset succeeded without explicit confirmation")
	}
	if err := service.ResetTask("synthetic-fixed-v2-001", true); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Lstat(workspace); !os.IsNotExist(err) {
		t.Fatalf("reset left learner workspace: %v", err)
	}
}

func TestLearnerCLICommandSurfaceAndStableScheduler(t *testing.T) {
	repository := learnerFixtureRepository(t)
	home := t.TempDir()
	var stdout, stderr bytes.Buffer
	run := func(want int, command ...string) {
		t.Helper()
		arguments := append([]string{"-root", repository, "-home", home}, command...)
		stdout.Reset()
		stderr.Reset()
		if code := RunCLI(Config{}, arguments, &stdout, &stderr); code != want {
			t.Fatalf("command %v: code=%d want=%d stdout=%q stderr=%q", command, code, want, stdout.String(), stderr.String())
		}
	}
	for _, command := range [][]string{
		{"list"},
		{"show", "synthetic-fixed-v2-001"},
		{"next"},
		{"start", "synthetic-fixed-v2-001"},
		{"pause"},
		{"resume", "synthetic-fixed-v2-001"},
	} {
		run(0, command...)
	}
	run(1, "check")
	run(1, "complete")
	run(2, "solution")
	run(0, "solution", "--reveal")
	run(2, "review")
	reference, err := os.ReadFile(filepath.Join(repository, "tasks", "ready", "synthetic-fixed-v2-001", "solution", "reference", "identity.go"))
	if err != nil {
		t.Fatal(err)
	}
	learnerPath := filepath.Join(home, "workspaces", "synthetic-fixed-v2-001", "subject", "identity.go")
	if err := os.WriteFile(learnerPath, reference, 0o600); err != nil {
		t.Fatal(err)
	}
	run(0, "resume", "synthetic-fixed-v2-001")
	run(0, "complete")
	run(0, "review")
	run(0, "progress")
	run(0, "feedback", "clear", "diagnostic")
	run(0, "hint")
	run(0, "export-review-context")
	service, err := New(Config{Root: repository, Home: home})
	if err != nil {
		t.Fatal(err)
	}
	first, err := service.Next()
	if err != nil {
		t.Fatal(err)
	}
	second, err := service.Next()
	if err != nil {
		t.Fatal(err)
	}
	if first.Task.ID != second.Task.ID {
		t.Fatalf("scheduler is not stable: %s then %s", first.Task.ID, second.Task.ID)
	}
	run(2, "reset-task", "synthetic-fixed-v2-001")
	run(0, "reset-task", "synthetic-fixed-v2-001", "--yes")
}

func learnerFixtureRepository(t *testing.T) string {
	t.Helper()
	repository, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	for _, directory := range []string{"config", "schemas", "curriculum", "sources"} {
		if err := isolation.CopyTree(filepath.Join(repository, directory), filepath.Join(root, directory)); err != nil {
			t.Fatalf("copy %s: %v", directory, err)
		}
	}
	// Synthetic repositories exercise the generic fallback; they do not
	// contain the real course lessons or its preparation guide.
	if err := os.Remove(filepath.Join(root, "curriculum", "route.yaml")); err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	for _, file := range []string{"LICENSE", "THIRD_PARTY_NOTICES"} {
		if err := isolation.CopyFile(filepath.Join(repository, file), filepath.Join(root, file)); err != nil {
			t.Fatalf("copy %s: %v", file, err)
		}
	}
	for _, state := range []string{"drafts", "review", "ready", "rejected", "retired"} {
		if err := os.MkdirAll(filepath.Join(root, "tasks", state), 0o700); err != nil {
			t.Fatal(err)
		}
	}
	for _, fixture := range []struct{ profile, id string }{
		{"fixed", "synthetic-fixed-v2-001"},
		{"mutation", "synthetic-mutation-v2-001"},
		{"workspace", "synthetic-workspace-v2-001"},
	} {
		source := filepath.Join(repository, "internal", "testdata", "tasking", "v2", fixture.profile, fixture.id)
		destination := filepath.Join(root, "tasks", "ready", fixture.id)
		if err := isolation.CopyTree(source, destination); err != nil {
			t.Fatal(err)
		}
		replaceTestText(t, filepath.Join(destination, "task.yaml"), "quality_status: draft", "quality_status: ready")
		lifecycle := `schema_version: 1
state: ready
updated_at: "2026-09-04"
transitions:
  - from: ""
    to: draft
    at: "2026-09-04"
    actor: synthetic-fixture
    reason: initialized synthetic learner fixture
  - from: draft
    to: in-review
    at: "2026-09-04"
    actor: synthetic-reviewer
    reason: reviewed synthetic learner fixture
  - from: in-review
    to: ready
    at: "2026-09-04"
    actor: synthetic-reviewer
    reason: enabled synthetic learner integration fixture
`
		if err := os.WriteFile(filepath.Join(destination, "lifecycle.yaml"), []byte(lifecycle), 0o600); err != nil {
			t.Fatal(err)
		}
		hints := "# Progressive hints\n\n1. Inspect the function contract.\n2. Add a nonzero boundary case.\n"
		if err := os.WriteFile(filepath.Join(destination, "hints", "HINTS.md"), []byte(hints), 0o600); err != nil {
			t.Fatal(err)
		}
		digest, err := tasking.TaskContentDigest(destination)
		if err != nil {
			t.Fatal(err)
		}
		updateProvenanceDigest(t, filepath.Join(destination, "provenance.yaml"), digest)
	}
	return root
}

func updateProvenanceDigest(t *testing.T, path, digest string) {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	updated := false
	for index, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "content_digest:") {
			indent := line[:len(line)-len(strings.TrimLeft(line, " "))]
			lines[index] = indent + "content_digest: " + digest
			updated = true
		}
	}
	if !updated {
		t.Fatal("fixture provenance has no content_digest")
	}
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o600); err != nil {
		t.Fatal(err)
	}
}

func replaceTestText(t *testing.T, path, old, replacement string) {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	result := strings.Replace(string(content), old, replacement, 1)
	if result == string(content) {
		t.Fatalf("%q not found in %s", old, path)
	}
	if err := os.WriteFile(path, []byte(result), 0o600); err != nil {
		t.Fatal(err)
	}
}

func assertLearnerWorkspaceSafe(t *testing.T, workspace string) {
	t.Helper()
	for _, relative := range workspacePaths(t, workspace) {
		lower := strings.ToLower(relative)
		for _, forbidden := range []string{"grader", "solution", "mutant", "verification", "reference"} {
			if strings.Contains(lower, forbidden) {
				t.Fatalf("persistent learner workspace contains %q: %s", forbidden, relative)
			}
		}
	}
}

func workspacePaths(t *testing.T, workspace string) []string {
	t.Helper()
	var paths []string
	err := filepath.WalkDir(workspace, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == workspace {
			return nil
		}
		relative, err := filepath.Rel(workspace, path)
		if err != nil {
			return err
		}
		paths = append(paths, filepath.ToSlash(relative))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(paths)
	return paths
}

func assertReviewContextSafe(t *testing.T, path string) {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lower := strings.ToLower(string(content))
	for _, forbidden := range []string{"testidentitygrader", "constant-zero", "solution/reference", "verification/mutants"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("review context contains hidden marker %q", forbidden)
		}
	}
	for _, required := range []string{"## Task", "## Learner code", "## Last public check", "## Skill metadata", "## Used hints"} {
		if !bytes.Contains(content, []byte(required)) {
			t.Fatalf("review context missing %s", required)
		}
	}
}
