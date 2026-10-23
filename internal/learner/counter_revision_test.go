package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

const counterRevisionID = "concurrent-counter-read-001"

func counterRevisionService(t *testing.T) *Service {
	t.Helper()
	s := uiTestService(t, "")
	root := t.TempDir()
	for _, dir := range []string{"config", "schemas", "curriculum", "sources", "tasks/ready/" + counterRevisionID} {
		if err := isolation.CopyTree(filepath.Join(s.root, dir), filepath.Join(root, dir)); err != nil {
			t.Fatal(err)
		}
	}
	if err := isolation.CopyFile(filepath.Join(s.root, "docs/LEARNER_START.md"), filepath.Join(root, "docs/LEARNER_START.md")); err != nil {
		t.Fatal(err)
	}
	var err error
	s.root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	s.now = func() time.Time { return time.Date(2026, 10, 23, 11, 0, 0, 0, time.UTC) }
	return s
}

// The pinned old public bytes reconstruct historical content before Start;
// changing only a marker on a current workspace would not establish this.
func setCounterRevisionFixture(t *testing.T, root, fixture, want string) {
	t.Helper()
	dir := filepath.Join(root, "tasks/ready", counterRevisionID)
	before, err := tasking.TaskContentDigest(dir)
	if err != nil || (before != counterBeforeCommonStart && before != counterWithCommonStart) {
		t.Fatalf("unexpected counter revision %s: %v", before, err)
	}
	beforeFiles, err := tasking.LearnerWorkspaceFiles(root, counterRevisionID)
	if err != nil {
		t.Fatal(err)
	}
	public, err := os.ReadFile(filepath.Join("testdata", fixture))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "tests/public/public_test.go"), public, 0600); err != nil {
		t.Fatal(err)
	}
	digest, err := tasking.TaskContentDigest(dir)
	if err != nil || digest != want {
		t.Fatalf("counter fixture digest %s, want %s: %v", digest, want, err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "provenance.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var provenance tasking.ProvenanceV2
	if err := yaml.Unmarshal(data, &provenance); err != nil {
		t.Fatal(err)
	}
	provenance.Project.ContentDigest = digest
	writeYAMLFixture(t, filepath.Join(dir, "provenance.yaml"), provenance)
	afterFiles, err := tasking.LearnerWorkspaceFiles(root, counterRevisionID)
	if err != nil || !reflect.DeepEqual(beforeFiles, afterFiles) {
		t.Fatalf("counter revision changed learner-export paths, bytes or editable flags: %v", err)
	}
}

func historicalCounterWorkspace(t *testing.T, s *Service, completed, generation bool) (string, Progress, []byte, []byte) {
	t.Helper()
	setCounterRevisionFixture(t, s.root, "counter-public-before.go.txt", counterBeforeCommonStart)
	if err := saveProgress(s.home, Progress{SchemaVersion: 3, Tasks: map[string]*TaskProgress{}}); err != nil {
		t.Fatal(err)
	}
	workspace, err := s.Start(counterRevisionID)
	if err != nil {
		t.Fatal(err)
	}
	markerPath := filepath.Join(workspace, learnerWorkspaceMarker)
	markerBytes, err := os.ReadFile(markerPath)
	if err != nil {
		t.Fatal(err)
	}
	var marker WorkspaceManifest
	if err := json.Unmarshal(markerBytes, &marker); err != nil || marker.ContentDigest != counterBeforeCommonStart || marker.Generation == "" {
		t.Fatalf("workspace was not created from pinned historical content: %+v %v", marker, err)
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
		if err := os.WriteFile(markerPath, markerBytes, 0600); err != nil {
			t.Fatal(err)
		}
		if err := os.Chmod(markerPath, 0400); err != nil {
			t.Fatal(err)
		}
	}
	code, err := os.ReadFile(filepath.Join(s.root, "tasks/ready", counterRevisionID, "solution/mutants/unlock-before-read/subject.go"))
	if err != nil {
		t.Fatal(err)
	}
	code = append(code, []byte("\n// Saved learner edit must survive the public-test revision.\n")...)
	if err := os.WriteFile(filepath.Join(workspace, "subject/subject.go"), code, 0600); err != nil {
		t.Fatal(err)
	}
	p, err := loadProgress(s.home)
	if err != nil {
		t.Fatal(err)
	}
	item := p.Tasks[counterRevisionID]
	item.Attempts, item.ActivitySeconds = 7, 420
	item.UsedHints = []int{1}
	item.Feedback = []FeedbackEntry{{RecordedAt: "2026-10-22T10:00:00Z", Text: "saved observation"}}
	item.ActiveSessionStartedAt = ""
	if completed {
		item.Completed, item.Status = true, "completed"
		item.CompletedAt = "2026-10-22T10:01:00Z"
		item.SkillEncounters = []SkillEncounter{{SkillID: "sync.data-races-happens-before", Role: "primary", Count: 1}}
		p.Completions = []CompletionEvidence{{TaskID: counterRevisionID, CompletedAt: item.CompletedAt, Skills: []string{"sync.data-races-happens-before"}, RouteRole: "practice", Assisted: true}}
	}
	if err := saveProgress(s.home, p); err != nil {
		t.Fatal(err)
	}
	return workspace, p, markerBytes, code
}

func assertCounterSavedBytes(t *testing.T, workspace string, marker, code []byte) {
	t.Helper()
	for path, want := range map[string][]byte{learnerWorkspaceMarker: marker, "subject/subject.go": code} {
		got, err := os.ReadFile(filepath.Join(workspace, path))
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("counter compatibility rewrote saved %s: %v", path, err)
		}
	}
}

func TestCounterRevisionPreservesHistoricalWorkspace(t *testing.T) {
	for _, completed := range []bool{false, true} {
		for _, generation := range []bool{false, true} {
			name := map[bool]string{false: "in-progress", true: "completed"}[completed] + "/" + map[bool]string{false: "legacy", true: "generation"}[generation]
			t.Run(name, func(t *testing.T) {
				s := counterRevisionService(t)
				workspace, saved, marker, code := historicalCounterWorkspace(t, s, completed, generation)
				before, err := workspaceTree(s.home, "")
				if err != nil {
					t.Fatal(err)
				}
				setCounterRevisionFixture(t, s.root, "counter-public-common-start.go.txt", counterWithCommonStart)
				view, err := s.TaskView(counterRevisionID)
				if err != nil || view.WorkspaceUnavailable || view.Workspace != workspace || view.Info.ContentDigest != counterWithCommonStart {
					t.Fatalf("historical counter unavailable: %+v %v", view, err)
				}
				after, err := workspaceTree(s.home, "")
				if err != nil || before != after {
					t.Fatalf("read changed solution, marker, permissions or history: %v", err)
				}
				if got, err := s.Resume(counterRevisionID); err != nil || got != workspace {
					t.Fatalf("resume: %s %v", got, err)
				}
				assertCounterSavedBytes(t, workspace, marker, code)
				progress, err := loadProgress(s.home)
				if err != nil {
					t.Fatal(err)
				}
				got, original := progress.Tasks[counterRevisionID], saved.Tasks[counterRevisionID]
				if got.Attempts != 7 || got.ActivitySeconds != 420 || !reflect.DeepEqual(got.UsedHints, original.UsedHints) || !reflect.DeepEqual(got.Feedback, original.Feedback) || !reflect.DeepEqual(got.SkillEncounters, original.SkillEncounters) || !reflect.DeepEqual(progress.Completions, saved.Completions) || got.CompletedAt != original.CompletedAt || got.Completed != completed {
					t.Fatalf("resume changed saved evidence: %+v", progress)
				}
				if got.LastInteractionAt != s.timestamp() || (got.ActiveSessionStartedAt != "") != !completed {
					t.Fatalf("ordinary resume state missing: %+v", got)
				}
			})
		}
	}
}
