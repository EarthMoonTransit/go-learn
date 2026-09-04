package learner

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
)

func TestProgressRoundTripContainsOnlyLearnerState(t *testing.T) {
	home := t.TempDir()
	progress := Progress{SchemaVersion: progressSchemaVersion, CurrentTaskID: "sample-001", Tasks: map[string]*TaskProgress{
		"sample-001": {
			TaskID: "sample-001", Status: "in-progress", Workspace: "workspaces/sample-001",
			UsedHints: []int{1}, Feedback: []FeedbackEntry{{RecordedAt: "2026-09-04T10:00:00Z", Text: "needs another example"}},
			SkillEncounters: []SkillEncounter{}, StartedAt: "2026-09-04T10:00:00Z", LastInteractionAt: "2026-09-04T10:00:00Z",
		},
	}}
	if err := saveProgress(home, progress); err != nil {
		t.Fatal(err)
	}
	loaded, err := loadProgress(home)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.CurrentTaskID != progress.CurrentTaskID || loaded.Tasks["sample-001"].Feedback[0].Text != "needs another example" {
		t.Fatalf("round trip changed progress: %+v", loaded)
	}
	content, err := os.ReadFile(filepath.Join(home, "progress.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"grader", "reference_solution", "mutants", "hidden_expectations"} {
		if bytes.Contains(bytes.ToLower(content), []byte(forbidden)) {
			t.Fatalf("progress contains forbidden checker field %q", forbidden)
		}
	}
}

func TestPreparePrivateSubdirectoryRejectsSymlink(t *testing.T) {
	home := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(home, "workspaces")); err != nil {
		t.Fatal(err)
	}
	if _, err := preparePrivateSubdirectory(home, "workspaces"); err == nil {
		t.Fatal("symlinked learner subdirectory was accepted")
	}
}

func TestPrepareDirectoryRejectsSymlinkHome(t *testing.T) {
	parent := t.TempDir()
	outside := t.TempDir()
	home := filepath.Join(parent, "learner-home")
	if err := os.Symlink(outside, home); err != nil {
		t.Fatal(err)
	}
	if _, err := prepareDirectory(home); err == nil {
		t.Fatal("symlinked learner home was accepted")
	}
}

func TestRenderReviewContextEscapesLearnerCodeFence(t *testing.T) {
	info := tasking.LearnerTaskInfo{ID: "sample-001", Prompt: "# Task\n\nImplement the API.", PrimarySkill: "functions.signatures-results"}
	progress := &TaskProgress{UsedHints: []int{1}, LastPublicCheck: &PublicCheckRecord{Passed: false, Summary: "public checks found a problem", Diagnostic: "public failure"}}
	files := []tasking.LearnerWorkspaceFile{{Path: "subject/value.go", Content: []byte("package subject\n```\n")}}
	content := renderReviewContext(info, progress, files, map[string]SkillMetadata{
		"functions.signatures-results": {ID: "functions.signatures-results", Definition: "Declares callable APIs."},
	}, []string{"Inspect the result type."})
	for _, wanted := range []string{"# Learner review context", "subject/value.go", "public failure", "functions.signatures-results", "Inspect the result type."} {
		if !strings.Contains(content, wanted) {
			t.Fatalf("review context missing %q", wanted)
		}
	}
	if !strings.Contains(content, "````go") {
		t.Fatal("learner code did not select a safe outer Markdown fence")
	}
}

func TestCLIArgumentContracts(t *testing.T) {
	if _, ok, err := resetArguments([]string{"sample-001", "--yes"}); err != nil || !ok {
		t.Fatalf("valid reset arguments rejected: ok=%t err=%v", ok, err)
	}
	if _, _, err := resetArguments([]string{"sample-001", "other-001", "--yes"}); err == nil {
		t.Fatal("multiple reset task ids accepted")
	}
	var stdout, stderr bytes.Buffer
	if code := RunCLI(Config{Root: t.TempDir(), Home: t.TempDir()}, []string{"unknown"}, &stdout, &stderr); code != 2 || !strings.Contains(stderr.String(), "unknown command") {
		t.Fatalf("unknown command result: code=%d stderr=%q", code, stderr.String())
	}
}
