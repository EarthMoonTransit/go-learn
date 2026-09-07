package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestPublicCheckRevisionPreservesExistingWorkspace(t *testing.T) {
	for _, generation := range []bool{false, true} {
		t.Run(map[bool]string{false: "legacy", true: "generation"}[generation], func(t *testing.T) {
			s := uiTestService(t, "")
			const id = "literal-bindings-001"
			workspace, err := s.Start(id)
			if err != nil {
				t.Fatal(err)
			}
			p, err := s.Progress()
			if err != nil {
				t.Fatal(err)
			}
			p.Tasks[id].Completed = true
			p.Tasks[id].Status = "completed"
			p.Tasks[id].CompletedAt = "2026-09-05T01:00:00Z"
			p.Tasks[id].Attempts = 4
			p.Tasks[id].UsedHints = []int{1}
			p.Tasks[id].ActivitySeconds = 420
			p.Tasks[id].ActiveSessionStartedAt = ""
			if err := saveProgress(s.home, p); err != nil {
				t.Fatal(err)
			}
			info, err := s.Show(id)
			if err != nil {
				t.Fatal(err)
			}
			markerPath := filepath.Join(workspace, learnerWorkspaceMarker)
			data, err := os.ReadFile(markerPath)
			if err != nil {
				t.Fatal(err)
			}
			var marker WorkspaceManifest
			if err := json.Unmarshal(data, &marker); err != nil {
				t.Fatal(err)
			}
			if !generation {
				marker.Generation = ""
			}
			marker.ContentDigest = bindingsBeforePublicChecks
			data, err = json.Marshal(marker)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.Chmod(markerPath, 0600); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(markerPath, data, 0600); err != nil {
				t.Fatal(err)
			}
			if err := os.Chmod(markerPath, 0444); err != nil {
				t.Fatal(err)
			}
			solutionPath := filepath.Join(workspace, "subject", "bindings.go")
			solution := []byte("package subject\n// learner work must remain byte-identical\nvar count = 12\nvar scale = 1.25\nvar label = \"G\"\nvar enabled = false\n")
			if err := os.WriteFile(solutionPath, solution, 0600); err != nil {
				t.Fatal(err)
			}
			beforeProgress := persistedBytes(t, s)
			info.ContentDigest = bindingsWithPublicChecks
			if got, err := s.validateWorkspace(p.Tasks[id], info); err != nil || got != workspace {
				t.Fatalf("compatible workspace=%s err=%v", got, err)
			}
			for path, want := range map[string][]byte{markerPath: data, solutionPath: solution, filepath.Join(s.home, "progress.json"): beforeProgress} {
				got, err := os.ReadFile(path)
				if err != nil || !bytes.Equal(got, want) {
					t.Fatalf("compatibility rewrote %s: %v", filepath.Base(path), err)
				}
			}
			for name, mutate := range map[string]func(*WorkspaceManifest){
				"schema":         func(m *WorkspaceManifest) { m.SchemaVersion++ },
				"task":           func(m *WorkspaceManifest) { m.TaskID = "other-task-001" },
				"submission":     func(m *WorkspaceManifest) { m.SubmissionKind = "tests" },
				"unknown digest": func(m *WorkspaceManifest) { m.ContentDigest = "unrecognized" },
			} {
				t.Run(name, func(t *testing.T) {
					bad := marker
					mutate(&bad)
					encoded, err := json.Marshal(bad)
					if err != nil {
						t.Fatal(err)
					}
					if err := os.Chmod(markerPath, 0600); err != nil {
						t.Fatal(err)
					}
					if err := os.WriteFile(markerPath, encoded, 0600); err != nil {
						t.Fatal(err)
					}
					if err := os.Chmod(markerPath, 0444); err != nil {
						t.Fatal(err)
					}
					if _, err := s.validateWorkspace(p.Tasks[id], info); err == nil {
						t.Fatal("invalid marker accepted")
					}
				})
			}
		})
	}
}

func TestPublicCheckRevisionIsExactAndDirectional(t *testing.T) {
	for _, tc := range []struct {
		name, id, stored, current string
		want                      bool
	}{
		{"old current", "literal-bindings-001", bindingsBeforePublicChecks, bindingsBeforePublicChecks, true},
		{"new current", "literal-bindings-001", bindingsWithPublicChecks, bindingsWithPublicChecks, true},
		{"approved update", "literal-bindings-001", bindingsBeforePublicChecks, bindingsWithPublicChecks, true},
		{"rollback", "literal-bindings-001", bindingsWithPublicChecks, bindingsBeforePublicChecks, false},
		{"other task", "zero-results-001", bindingsBeforePublicChecks, bindingsWithPublicChecks, false},
		{"third target", "literal-bindings-001", bindingsBeforePublicChecks, "other", false},
		{"third source", "literal-bindings-001", "other", bindingsWithPublicChecks, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := compatibleContentRevision(tc.id, tc.stored, tc.current); got != tc.want {
				t.Fatalf("compatible=%v, want %v", got, tc.want)
			}
		})
	}
}
