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

// Reconstruct the pre-E01 package from an explicit historical fixture before
// creating old workspaces. This stays meaningful after canonical content changes.
func prepareHistoricalClampFixture(t *testing.T, root string) {
	t.Helper()
	setClampRevisionFixture(t, root, "clamp-before-testing-bridge", clampWithTestingBridge, clampBeforeTestingBridge, []string{})
}

func applyClampRevisionFixture(t *testing.T, root string) {
	t.Helper()
	setClampRevisionFixture(t, root, "clamp-testing-bridge", clampBeforeTestingBridge, clampWithTestingBridge, []string{"testing.basics-diagnostics"})
}

func setClampRevisionFixture(t *testing.T, root, fixture, from, to string, secondary []string) {
	t.Helper()
	dir := filepath.Join(root, "tasks/ready/clamp-tests-001")
	before, err := tasking.TaskContentDigest(dir)
	if err != nil || before != from {
		t.Fatalf("unexpected Clamp baseline: %s; want %s, %v", before, from, err)
	}
	beforeFiles, err := tasking.LearnerWorkspaceFiles(root, "clamp-tests-001")
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"starter/subject_test.go", "hints/HINTS.md"} {
		data, err := os.ReadFile(filepath.Join("testdata", fixture, path))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, path), data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	digest, err := tasking.TaskContentDigest(dir)
	if err != nil || digest != to {
		t.Fatalf("fixture drifted from its exact compatibility target: %s; want %s, %v", digest, to, err)
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
	manifestPath := filepath.Join(dir, "task.yaml")
	data, err = os.ReadFile(manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	var manifest tasking.TaskV2
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		t.Fatal(err)
	}
	manifest.Skills.Secondary = secondary
	writeYAMLFixture(t, manifestPath, manifest)
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
				step.Skills = append([]string{"expressions.comparison"}, secondary...)
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
				prepareHistoricalClampFixture(t, s.root)
				const id = "clamp-tests-001"
				if err := saveProgress(s.home, Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}); err != nil {
					t.Fatal(err)
				}
				workspace, err := s.Start(id)
				if err != nil {
					t.Fatal(err)
				}
				learnerPath := filepath.Join(workspace, "subject/subject_test.go")
				learnerCode, err := os.ReadFile(learnerPath)
				if err != nil {
					t.Fatal(err)
				}
				learnerCode = append(learnerCode, []byte("\n// Personal boundary-case notes must survive the content update.\n")...)
				if err := os.WriteFile(learnerPath, learnerCode, 0600); err != nil {
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
				for path, want := range map[string][]byte{learnerPath: learnerCode, markerPath: data} {
					got, err := os.ReadFile(path)
					if err != nil || !bytes.Equal(got, want) {
						t.Fatalf("view/resume rewrote saved learner data: %s, %v", path, err)
					}
				}
			})
		}
	}
}
