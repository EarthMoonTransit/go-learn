//go:build taskintegration

package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"go-practice-bank/internal/tasking"
)

func TestBrokenLearnerCodeRemainsExportableAcrossProfiles(t *testing.T) {
	repository := learnerFixtureRepository(t)
	for _, fixture := range []struct{ id, path string }{
		{"synthetic-fixed-v2-001", "subject/identity.go"},
		{"synthetic-mutation-v2-001", "subject/identity_test.go"},
		{"synthetic-workspace-v2-001", "core/apply.go"},
	} {
		t.Run(fixture.id, func(t *testing.T) {
			s, err := New(Config{Root: repository, Home: t.TempDir()})
			if err != nil {
				t.Fatal(err)
			}
			workspace, err := s.Start(fixture.id)
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(workspace, filepath.FromSlash(fixture.path))
			code, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			broken := append(code, []byte("\nfunc unfinished() {\n")...)
			if err := os.WriteFile(path, broken, 0600); err != nil {
				t.Fatal(err)
			}
			var stdout, stderr bytes.Buffer
			if exit := RunCLI(Config{Root: repository, Home: s.home}, []string{"check"}, &stdout, &stderr); exit != 1 {
				t.Fatalf("syntax error exit=%d stdout=%q stderr=%q", exit, stdout.String(), stderr.String())
			}
			if !strings.Contains(stdout.String(), "EOF") || strings.Contains(stdout.String(), workspace) || strings.Contains(stdout.String(), "changed the public API") {
				t.Fatalf("unhelpful syntax diagnostic: %s", stdout.String())
			}
			full, err := s.Complete()
			if err != nil || full.Passed || full.Diagnostic != "" {
				t.Fatalf("full result=%+v err=%v", full, err)
			}
			progress, err := s.Progress()
			if err != nil {
				t.Fatal(err)
			}
			item := progress.Tasks[fixture.id]
			if item.Attempts != 2 || item.Completed || item.LastPublicCheck == nil || !strings.Contains(item.LastPublicCheck.Diagnostic, "EOF") {
				t.Fatalf("lost failed attempt: %+v", item)
			}
			export, err := s.ExportReviewContext()
			if err != nil {
				t.Fatal(err)
			}
			assertReviewContextSafe(t, export)
			content, err := os.ReadFile(export)
			if err != nil || !bytes.Contains(content, broken) || bytes.Contains(content, []byte(workspace)) {
				t.Fatalf("broken code export: %v", err)
			}
			// Relaxing syntax validation must not relax fixed-file validation.
			mod := filepath.Join(workspace, "go.mod")
			if err := os.Chmod(mod, 0600); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(mod, []byte("module replacement\n"), 0600); err != nil {
				t.Fatal(err)
			}
			if _, err := s.ExportReviewContext(); err == nil {
				t.Fatal("modified go.mod was exported")
			}
			if _, err := s.Check(); err == nil {
				t.Fatal("modified go.mod was checked")
			}
		})
	}
}

func TestChangedLearnerAPIIsAFailedAttemptAndCanBeExported(t *testing.T) {
	repository := learnerFixtureRepository(t)
	s, err := New(Config{Root: repository, Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	workspace, err := s.Start("synthetic-fixed-v2-001")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(workspace, "subject", "identity.go")
	if err := os.WriteFile(path, []byte("package subject\nfunc Changed(value int) int { return value }\n"), 0600); err != nil {
		t.Fatal(err)
	}
	result, err := s.Check()
	if err != nil || result.Passed || !strings.Contains(result.Diagnostic, "changed the public API") {
		t.Fatalf("API result=%+v err=%v", result, err)
	}
	if _, err := s.ExportReviewContext(); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(repository, "README.md"), path); err != nil {
		t.Fatal(err)
	}
	if _, err := s.ExportReviewContext(); err == nil {
		t.Fatal("symlinked learner file was exported")
	}
}

func TestLearnerSessionsCountRetriesAndPauseOnSwitch(t *testing.T) {
	repository := learnerFixtureRepository(t)
	clock := time.Date(2026, 9, 4, 10, 0, 0, 0, time.UTC)
	s, err := New(Config{Root: repository, Home: t.TempDir(), Now: func() time.Time { return clock }})
	if err != nil {
		t.Fatal(err)
	}
	workspace, err := s.Start("synthetic-fixed-v2-001")
	if err != nil {
		t.Fatal(err)
	}
	// A syntax error takes the real failed-attempt path without executing code.
	if err := os.WriteFile(filepath.Join(workspace, "subject", "identity.go"), []byte("package subject\nfunc broken(\n"), 0600); err != nil {
		t.Fatal(err)
	}
	clock = clock.Add(time.Minute)
	if _, err := s.Check(); err != nil {
		t.Fatal(err)
	}
	clock = clock.Add(9 * time.Minute)
	if _, err := s.Check(); err != nil {
		t.Fatal(err)
	}
	assertTime := func(id string, want int64) {
		t.Helper()
		progress, err := s.Progress()
		if err != nil {
			t.Fatal(err)
		}
		if got := progress.Tasks[id].ActivitySeconds; got != want {
			t.Fatalf("%s seconds=%d want=%d", id, got, want)
		}
	}
	assertTime("synthetic-fixed-v2-001", 600)
	clock = clock.Add(time.Minute)
	if _, err := s.Resume("synthetic-fixed-v2-001"); err != nil {
		t.Fatal(err)
	}
	assertTime("synthetic-fixed-v2-001", 660)
	if err := s.Pause(); err != nil {
		t.Fatal(err)
	}
	clock = clock.Add(time.Hour)
	if _, err := s.Check(); err != nil {
		t.Fatal(err)
	}
	assertTime("synthetic-fixed-v2-001", 660)
	if _, err := s.Resume("synthetic-fixed-v2-001"); err != nil {
		t.Fatal(err)
	}
	clock = clock.Add(time.Minute)
	if _, err := s.Start("synthetic-workspace-v2-001"); err != nil {
		t.Fatal(err)
	}
	clock = clock.Add(time.Minute)
	assertTime("synthetic-fixed-v2-001", 720)
	assertTime("synthetic-workspace-v2-001", 60)
	if _, err := s.Resume("synthetic-fixed-v2-001"); err != nil {
		t.Fatal(err)
	}
	clock = clock.Add(time.Minute)
	assertTime("synthetic-fixed-v2-001", 780)
	assertTime("synthetic-workspace-v2-001", 60)
}

func TestBothReadyTrialTasksCompleteWithoutChangingCanonicalContent(t *testing.T) {
	repository, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"map-take-001", "detached-append-001"} {
		t.Run(id, func(t *testing.T) {
			canonical := filepath.Join(repository, "tasks", "ready", id)
			before, err := tasking.TaskContentDigest(canonical)
			if err != nil {
				t.Fatal(err)
			}
			s, err := New(Config{Root: repository, Home: t.TempDir()})
			if err != nil {
				t.Fatal(err)
			}
			workspace, err := s.Start(id)
			if err != nil {
				t.Fatal(err)
			}
			public, err := s.Check()
			if err != nil || public.Passed {
				t.Fatalf("starter result=%+v err=%v", public, err)
			}
			if _, err := s.ExportReviewContext(); err != nil {
				t.Fatal(err)
			}
			// Canonical code is read as data; execution stays in the runner's copy.
			reference, err := os.ReadFile(filepath.Join(canonical, "solution", "reference", "subject.go"))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, "subject", "subject.go"), reference, 0600); err != nil {
				t.Fatal(err)
			}
			public, err = s.Check()
			if err != nil || !public.Passed {
				t.Fatalf("public result=%+v err=%v", public, err)
			}
			full, err := s.Complete()
			if err != nil || !full.Passed || full.Diagnostic != "" {
				t.Fatalf("full result=%+v err=%v", full, err)
			}
			if err := s.Feedback("automated smoke, not human feedback"); err != nil {
				t.Fatal(err)
			}
			progress, err := s.Progress()
			if err != nil {
				t.Fatal(err)
			}
			item := progress.Tasks[id]
			if !item.Completed || item.Attempts != 3 || item.ActiveSessionStartedAt != "" || len(item.Feedback) != 1 {
				t.Fatalf("invalid completion: %+v", item)
			}
			assertLearnerWorkspaceSafe(t, workspace)
			after, err := tasking.TaskContentDigest(canonical)
			if err != nil || before != after {
				t.Fatalf("canonical task changed: %v", err)
			}
		})
	}
}
