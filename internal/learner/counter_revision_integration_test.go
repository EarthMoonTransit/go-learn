//go:build taskintegration

package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

func TestCounterRevisionUsesExistingWorkspaceAndRealVerifier(t *testing.T) {
	canonical, err := filepath.Abs("../../tasks/ready/" + counterRevisionID)
	if err != nil {
		t.Fatal(err)
	}
	before, err := tasking.TaskContentDigest(canonical)
	if err != nil {
		t.Fatal(err)
	}
	for _, generation := range []bool{false, true} {
		t.Run(map[bool]string{false: "legacy", true: "generation"}[generation], func(t *testing.T) {
			s := counterRevisionService(t)
			workspace, saved, marker, code := historicalCounterWorkspace(t, s, false, generation)
			setCounterRevisionFixture(t, s.root, "counter-public-common-start.go.txt", counterWithCommonStart)
			s.verify = integrationVerifier(t, s.root)
			if got, err := s.Resume(counterRevisionID); err != nil || got != workspace {
				t.Fatalf("old workspace resume: %s %v", got, err)
			}
			assertCounterSavedBytes(t, workspace, marker, code)
			result, err := s.Check()
			if err != nil || result.Passed || !strings.Contains(result.Diagnostic, "DATA RACE") {
				t.Fatalf("current public race diagnostic: %+v %v", result, err)
			}
			assertCounterSavedBytes(t, workspace, marker, code)
			progress, err := loadProgress(s.home)
			if err != nil || progress.Tasks[counterRevisionID].Attempts != 8 || progress.Tasks[counterRevisionID].Completed || len(progress.Completions) != 0 {
				t.Fatalf("ordinary failed Check effects: %+v %v", progress, err)
			}
			reference, err := os.ReadFile(filepath.Join(s.root, "tasks/ready", counterRevisionID, "solution/reference/subject.go"))
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(workspace, "subject/subject.go"), reference, 0600); err != nil {
				t.Fatal(err)
			}
			result, err = s.Check()
			if err != nil || !result.Passed {
				t.Fatalf("reference current public: %+v %v", result, err)
			}
			result, err = s.Complete()
			if err != nil || !result.Passed {
				t.Fatalf("reference current full: %+v %v", result, err)
			}
			assertCounterSavedBytes(t, workspace, marker, reference)
			progress, err = loadProgress(s.home)
			if err != nil {
				t.Fatal(err)
			}
			item := progress.Tasks[counterRevisionID]
			if item.Attempts != 10 || item.ActivitySeconds != 420 || !item.Completed || item.CompletedAt != s.timestamp() || !reflect.DeepEqual(item.UsedHints, saved.Tasks[counterRevisionID].UsedHints) || !reflect.DeepEqual(item.Feedback, saved.Tasks[counterRevisionID].Feedback) || len(progress.Completions) != 1 || progress.Completions[0].TaskID != counterRevisionID {
				t.Fatalf("ordinary Complete effects: %+v", progress)
			}
			prior := persistedBytes(t, s)
			if _, err := s.Complete(); err == nil || !strings.Contains(err.Error(), "already completed") {
				t.Fatalf("completed history must not be duplicated: %v", err)
			}
			if !bytes.Equal(prior, persistedBytes(t, s)) {
				t.Fatal("rejected repeat Complete rewrote history")
			}
		})
	}
	after, err := tasking.TaskContentDigest(canonical)
	if err != nil || before != after {
		t.Fatalf("canonical task changed: %s -> %s: %v", before, after, err)
	}
}
