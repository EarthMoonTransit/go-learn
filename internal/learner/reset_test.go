package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func resetFixture(t *testing.T) *Service {
	t.Helper()
	s, err := New(Config{Root: t.TempDir(), Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	progress := Progress{SchemaVersion: 1, CurrentTaskID: "sample-001", Tasks: map[string]*TaskProgress{}}
	for _, id := range []string{"sample-001", "other-001"} {
		progress.Tasks[id] = &TaskProgress{TaskID: id, Workspace: "workspaces/" + id}
		if err := os.MkdirAll(filepath.Join(s.home, "workspaces", id), 0700); err != nil {
			t.Fatal(err)
		}
	}
	if err := saveProgress(s.home, progress); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestResetRejectsRedirectedWorkspaceWithoutChangingProgress(t *testing.T) {
	for _, kind := range []string{"absolute-parent", "relative-parent", "inside-home-parent", "final-symlink"} {
		t.Run(kind, func(t *testing.T) {
			s := resetFixture(t)
			outside := t.TempDir()
			if kind == "inside-home-parent" {
				outside = filepath.Join(s.home, "unrelated")
			}
			protected := filepath.Join(outside, "sample-001")
			if err := os.MkdirAll(protected, 0700); err != nil {
				t.Fatal(err)
			}
			marker := filepath.Join(protected, "keep.txt")
			if err := os.WriteFile(marker, []byte("keep"), 0600); err != nil {
				t.Fatal(err)
			}
			link := filepath.Join(s.home, "workspaces")
			target := outside
			if kind == "final-symlink" {
				link = filepath.Join(link, "sample-001")
				target = protected
			}
			if err := os.Rename(link, link+"-saved"); err != nil {
				t.Fatal(err)
			}
			if kind == "relative-parent" || kind == "inside-home-parent" {
				var err error
				target, err = filepath.Rel(filepath.Dir(link), target)
				if err != nil {
					t.Fatal(err)
				}
			}
			if err := os.Symlink(target, link); err != nil {
				t.Fatal(err)
			}
			before, err := os.ReadFile(filepath.Join(s.home, "progress.json"))
			if err != nil {
				t.Fatal(err)
			}
			if err := s.ResetTask("", true); err == nil {
				t.Fatal("redirected reset succeeded")
			}
			after, err := os.ReadFile(filepath.Join(s.home, "progress.json"))
			if err != nil || !bytes.Equal(before, after) {
				t.Fatalf("progress changed: %v", err)
			}
			if content, err := os.ReadFile(marker); err != nil || string(content) != "keep" {
				t.Fatalf("outside file changed: %v", err)
			}
		})
	}
}

func TestResetPreservesSiblingsAndAllowsMissingOrStaleWorkspaces(t *testing.T) {
	for _, kind := range []string{"stale-marker", "missing-task", "missing-parent", "nested-symlink"} {
		t.Run(kind, func(t *testing.T) {
			s := resetFixture(t)
			workspace := filepath.Join(s.home, "workspaces", "sample-001")
			outside := t.TempDir()
			marker := filepath.Join(outside, "keep.txt")
			if err := os.WriteFile(marker, []byte("keep"), 0600); err != nil {
				t.Fatal(err)
			}
			switch kind {
			case "stale-marker":
				if err := os.WriteFile(filepath.Join(workspace, learnerWorkspaceMarker), []byte("stale"), 0600); err != nil {
					t.Fatal(err)
				}
			case "missing-task":
				if err := os.RemoveAll(workspace); err != nil {
					t.Fatal(err)
				}
			case "missing-parent":
				if err := os.RemoveAll(filepath.Dir(workspace)); err != nil {
					t.Fatal(err)
				}
			case "nested-symlink":
				if err := os.Symlink(outside, filepath.Join(workspace, "link")); err != nil {
					t.Fatal(err)
				}
			}
			if err := s.ResetTask("sample-001", false); err == nil {
				t.Fatal("reset without confirmation succeeded")
			}
			if err := s.ResetTask("../sample-001", true); err == nil {
				t.Fatal("traversal accepted")
			}
			if err := s.ResetTask("", true); err != nil {
				t.Fatal(err)
			}
			progress, err := s.Progress()
			if err != nil {
				t.Fatal(err)
			}
			if progress.Tasks["sample-001"] != nil || progress.Tasks["other-001"] == nil || progress.CurrentTaskID != "" {
				t.Fatalf("wrong progress after reset: %+v", progress)
			}
			if kind != "missing-parent" {
				if _, err := os.Stat(filepath.Join(s.home, "workspaces", "other-001")); err != nil {
					t.Fatal(err)
				}
			}
			if _, err := os.Stat(marker); err != nil {
				t.Fatal("reset followed nested symlink")
			}
		})
	}
}
