package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

// Apply the proposed E01/M1 revision only to a disposable repository. Canonical
// content stays on the old revision until its separate content review and PR.
func applyClampRevisionFixture(t *testing.T, root string) {
	t.Helper()
	dir := filepath.Join(root, "tasks/ready/clamp-tests-001")
	before, err := tasking.TaskContentDigest(dir)
	if err != nil || before != clampBeforeTestingBridge {
		t.Fatalf("unexpected Clamp baseline: %s, %v", before, err)
	}
	beforeFiles, err := tasking.LearnerWorkspaceFiles(root, "clamp-tests-001")
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"starter/subject_test.go", "hints/HINTS.md"} {
		data, err := os.ReadFile(filepath.Join("testdata/clamp-testing-bridge", path))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, path), data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	digest, err := tasking.TaskContentDigest(dir)
	if err != nil || digest != clampWithTestingBridge {
		t.Fatalf("proposed content drifted from the exact compatibility target: %s, %v", digest, err)
	}
	provenancePath := filepath.Join(dir, "provenance.yaml")
	data, err := os.ReadFile(provenancePath)
	if err != nil {
		t.Fatal(err)
	}
	var provenance tasking.ProvenanceV2
	if err := yaml.Unmarshal(data, &provenance); err != nil {
		t.Fatal(err)
	}
	provenance.Project.ContentDigest = digest
	writeYAMLFixture(t, provenancePath, provenance)
	// Exercise future metadata without changing any historical learner record.
	manifestPath := filepath.Join(dir, "task.yaml")
	data, err = os.ReadFile(manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	data = bytes.Replace(data, []byte("secondary: []"), []byte("secondary: [testing.basics-diagnostics]"), 1)
	if err := os.WriteFile(manifestPath, data, 0600); err != nil {
		t.Fatal(err)
	}
	routePath := filepath.Join(root, "curriculum/route.yaml")
	data, err = os.ReadFile(routePath)
	if err != nil {
		t.Fatal(err)
	}
	var route Route
	if err := yaml.Unmarshal(data, &route); err != nil {
		t.Fatal(err)
	}
	for b := range route.Blocks {
		for i := range route.Blocks[b].Steps {
			step := &route.Blocks[b].Steps[i]
			if step.TaskID == "clamp-tests-001" {
				step.Skills = []string{"expressions.comparison", "testing.basics-diagnostics"}
			}
		}
	}
	writeYAMLFixture(t, routePath, route)
	afterFiles, err := tasking.LearnerWorkspaceFiles(root, "clamp-tests-001")
	if err != nil || len(beforeFiles) != len(afterFiles) {
		t.Fatal("revision changed the workspace file set", err)
	}
	for i, before := range beforeFiles {
		after := afterFiles[i]
		if before.Path != after.Path || before.Editable != after.Editable || (!before.Editable && !bytes.Equal(before.Content, after.Content)) {
			t.Fatalf("revision changed a path, editable flag or read-only file: %s", before.Path)
		}
	}
}

func writeYAMLFixture(t *testing.T, path string, value any) {
	t.Helper()
	data, err := yaml.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
}

func TestClampRevisionPreservesWorkspaceAndHistoricalEvidence(t *testing.T) {
	for _, status := range []string{"paused", "completed"} {
		for _, generation := range []bool{false, true} {
			name := status + "/" + map[bool]string{false: "legacy-marker", true: "generation-marker"}[generation]
			t.Run(name, func(t *testing.T) {
				s := goLandTestService(t)
				const id = "clamp-tests-001"
				if err := saveProgress(s.home, Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}); err != nil {
					t.Fatal(err)
				}
				workspace, err := s.Start(id)
				if err != nil {
					t.Fatal(err)
				}
				p, err := s.Progress()
				if err != nil {
					t.Fatal(err)
				}
				item := p.Tasks[id]
				item.Status, item.Completed = status, status == "completed"
				item.Attempts, item.ActivitySeconds = 7, 420
				item.ActiveSessionStartedAt = ""
				item.UsedHints = []int{1}
				if item.Completed {
					item.CompletedAt = "2026-10-01T00:00:00Z"
					item.SkillEncounters = []SkillEncounter{{SkillID: "expressions.comparison", Role: "primary", Count: 1}}
					p.Completions = []CompletionEvidence{{TaskID: id, CompletedAt: item.CompletedAt, Skills: []string{"expressions.comparison"}, RouteRole: "practice", Assisted: true}}
				}
				if err := saveProgress(s.home, p); err != nil {
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
				if marker.ContentDigest != clampBeforeTestingBridge {
					t.Fatal("workspace was not created from the historical content")
				}
				if !generation {
					marker.Generation = ""
					data, err = json.Marshal(marker)
					if err != nil {
						t.Fatal(err)
					}
					if err := os.Chmod(markerPath, 0600); err != nil {
						t.Fatal(err)
					}
					if err := os.WriteFile(markerPath, data, 0444); err != nil {
						t.Fatal(err)
					}
					if err := os.Chmod(markerPath, 0444); err != nil {
						t.Fatal(err)
					}
				}
				before, err := workspaceTree(s.home, "")
				if err != nil {
					t.Fatal(err)
				}
				applyClampRevisionFixture(t, s.root)
				view, err := s.TaskView(id)
				if err != nil || view.WorkspaceUnavailable || view.Workspace != workspace {
					t.Fatalf("old workspace unavailable: %+v, %v", view, err)
				}
				if !reflect.DeepEqual(view.Info.SecondarySkills, []string{"testing.basics-diagnostics"}) {
					t.Fatal("fixture must expose the new secondary while history remains old")
				}
				after, err := workspaceTree(s.home, "")
				if err != nil || before != after {
					t.Fatal("revision rewrote workspace, marker, progress, paths or permissions", err)
				}
				progress, err := s.Progress()
				if err != nil || !reflect.DeepEqual(progress.Completions, p.Completions) || !reflect.DeepEqual(progress.Tasks[id].SkillEncounters, item.SkillEncounters) {
					t.Fatal("new secondary changed historical evidence", err)
				}
				if !item.Completed {
					if got, err := s.Resume(id); err != nil || got != workspace {
						t.Fatalf("resume requires reset: %s, %v", got, err)
					}
				}
			})
		}
	}
}
