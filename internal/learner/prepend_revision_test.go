package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"testing"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func prependRevisionService(t *testing.T) *Service {
	t.Helper()
	s := uiTestService(t, "")
	root := t.TempDir()
	for _, dir := range []string{"config", "schemas", "curriculum", "sources", "tasks/ready/prepend-values-tests-001"} {
		if err := isolation.CopyTree(filepath.Join(s.root, dir), filepath.Join(root, dir)); err != nil {
			t.Fatal(err)
		}
	}
	if err := isolation.CopyFile(filepath.Join(s.root, "docs/LEARNER_START.md"), filepath.Join(root, "docs/LEARNER_START.md")); err != nil {
		t.Fatal(err)
	}
	s.root = root
	return s
}

func historicalPrependFixture(t *testing.T, root string) {
	t.Helper()
	digest, err := tasking.TaskContentDigest(filepath.Join(root, "tasks/ready/prepend-values-tests-001"))
	if err != nil || (digest != prependBeforeVariadicBridge && digest != prependWithVariadicBridge) {
		t.Fatalf("unexpected canonical Prepend revision: %s, %v", digest, err)
	}
	setPrependRevisionFixture(t, root, "prepend-before-variadic-bridge", digest, prependBeforeVariadicBridge)
}

func applyPrependBridgeFixture(t *testing.T, root string) {
	t.Helper()
	setPrependRevisionFixture(t, root, "prepend-variadic-bridge", prependBeforeVariadicBridge, prependWithVariadicBridge)
}

func setPrependRevisionFixture(t *testing.T, root, fixture, from, to string) {
	t.Helper()
	dir := filepath.Join(root, "tasks/ready/prepend-values-tests-001")
	before, err := tasking.TaskContentDigest(dir)
	if err != nil || before != from {
		t.Fatalf("Prepend fixture starts at %s, want %s: %v", before, from, err)
	}
	beforeFiles, err := tasking.LearnerWorkspaceFiles(root, "prepend-values-tests-001")
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"starter/subject_test.go", "hints/HINTS.md"} {
		content, err := os.ReadFile(filepath.Join("testdata", fixture, path))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, path), content, 0600); err != nil {
			t.Fatal(err)
		}
	}
	digest, err := tasking.TaskContentDigest(dir)
	if err != nil || digest != to {
		t.Fatalf("Prepend fixture digest %s, want %s: %v", digest, to, err)
	}
	content, err := os.ReadFile(filepath.Join(dir, "provenance.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var provenance tasking.ProvenanceV2
	if err := yaml.Unmarshal(content, &provenance); err != nil {
		t.Fatal(err)
	}
	provenance.Project.ContentDigest = digest
	writeYAMLFixture(t, filepath.Join(dir, "provenance.yaml"), provenance)
	content, err = os.ReadFile(filepath.Join(dir, "task.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var manifest tasking.TaskV2
	if err := yaml.Unmarshal(content, &manifest); err != nil {
		t.Fatal(err)
	}
	manifest.Prerequisites = []string{"functions.signatures-results", "testing.basics-diagnostics"}
	if to == prependWithVariadicBridge {
		manifest.Prerequisites = append(manifest.Prerequisites, "slices.header-model", "slices.len-cap-slicing")
	}
	writeYAMLFixture(t, filepath.Join(dir, "task.yaml"), manifest)
	afterFiles, err := tasking.LearnerWorkspaceFiles(root, "prepend-values-tests-001")
	if err != nil || len(beforeFiles) != len(afterFiles) {
		t.Fatal("workspace file set changed", err)
	}
	for i, before := range beforeFiles {
		after := afterFiles[i]
		if before.Path != after.Path || before.Editable != after.Editable || (!before.Editable && !bytes.Equal(before.Content, after.Content)) {
			t.Fatalf("read-only workspace contract changed: %s", before.Path)
		}
	}
}

func TestPrependRevisionPreservesHistoricalWorkspace(t *testing.T) {
	for _, status := range []string{"paused", "completed"} {
		for _, generation := range []bool{false, true} {
			for _, legacyPath := range []bool{false, true} {
				name := status + "/" + map[bool]string{false: "legacy-marker", true: "generation-marker"}[generation] + "/" + map[bool]string{false: "numbered-path", true: "legacy-path"}[legacyPath]
				t.Run(name, func(t *testing.T) {
					s := prependRevisionService(t)
					historicalPrependFixture(t, s.root)
					const id = "prepend-values-tests-001"
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
					if legacyPath {
						legacy := filepath.Join(s.home, "workspaces", id)
						if err := os.Rename(workspace, legacy); err != nil {
							t.Fatal(err)
						}
						workspace = legacy
						item.Workspace = filepath.Join("workspaces", id)
					}
					learnerPath := filepath.Join(workspace, "subject/subject_test.go")
					learnerCode, err := os.ReadFile(learnerPath)
					if err != nil {
						t.Fatal(err)
					}
					learnerCode = append(learnerCode, []byte("\n// My additional prefix cases must survive the revision.\n")...)
					if err := os.WriteFile(learnerPath, learnerCode, 0600); err != nil {
						t.Fatal(err)
					}
					item.Status, item.Completed = status, status == "completed"
					item.Attempts, item.ActivitySeconds = 7, 420
					item.ActiveSessionStartedAt = ""
					item.UsedHints = []int{1}
					if item.Completed {
						item.CompletedAt = "2026-10-01T00:00:00Z"
						item.SkillEncounters = []SkillEncounter{{SkillID: "functions.variadic", Role: "primary", Count: 1}}
						p.Completions = []CompletionEvidence{{TaskID: id, CompletedAt: item.CompletedAt, Skills: []string{"functions.variadic"}, RouteRole: "introduction", Assisted: true}}
					}
					if err := saveProgress(s.home, p); err != nil {
						t.Fatal(err)
					}
					markerPath := filepath.Join(workspace, learnerWorkspaceMarker)
					markerBytes, err := os.ReadFile(markerPath)
					if err != nil {
						t.Fatal(err)
					}
					var marker WorkspaceManifest
					if err := json.Unmarshal(markerBytes, &marker); err != nil {
						t.Fatal(err)
					}
					if marker.ContentDigest != prependBeforeVariadicBridge {
						t.Fatal("workspace is not historical")
					}
					if !generation {
						marker.Generation = ""
						markerBytes, err = json.Marshal(marker)
						if err != nil {
							t.Fatal(err)
						}
						if err := os.Chmod(markerPath, 0600); err != nil {
							t.Fatal(err)
						}
						if err := os.WriteFile(markerPath, markerBytes, 0444); err != nil {
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
					applyPrependBridgeFixture(t, s.root)
					view, err := s.TaskView(id)
					if err != nil || view.WorkspaceUnavailable || view.Workspace != workspace {
						t.Fatalf("historical workspace unavailable: %+v, %v", view, err)
					}
					if !slices.Contains(view.Info.Prerequisites, "slices.header-model") {
						t.Fatal("current metadata not applied")
					}
					after, err := workspaceTree(s.home, "")
					if err != nil || before != after {
						t.Fatal("view rewrote history, paths, marker, permissions or learner code", err)
					}
					if !item.Completed {
						if got, err := s.Resume(id); err != nil || got != workspace {
							t.Fatalf("resume requires reset: %s, %v", got, err)
						}
					}
					for path, want := range map[string][]byte{learnerPath: learnerCode, markerPath: markerBytes} {
						got, err := os.ReadFile(path)
						if err != nil || !bytes.Equal(got, want) {
							t.Fatalf("saved learner bytes changed: %s, %v", path, err)
						}
					}
					progress, err := s.Progress()
					if err != nil || !reflect.DeepEqual(progress.Completions, p.Completions) || !reflect.DeepEqual(progress.Tasks[id].SkillEncounters, item.SkillEncounters) || progress.Tasks[id].Attempts != 7 || progress.Tasks[id].ActivitySeconds != 420 || !slices.Equal(progress.Tasks[id].UsedHints, []int{1}) {
						t.Fatal("revision changed saved evidence, attempts or hints", err)
					}
				})
			}
		}
	}
}
