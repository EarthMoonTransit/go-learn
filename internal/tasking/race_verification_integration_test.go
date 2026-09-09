//go:build taskintegration

package tasking

import (
	"go-practice-bank/internal/isolation"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The race changes no return value. Ordinary behavioral tests therefore cannot
// accidentally detect it through a second unrelated defect.
const racingIdentity = `package subject
import "sync"
func Identity(value int) int {
 var cell int
 var workers sync.WaitGroup
 workers.Add(2)
 go func(){cell=1;workers.Done()}()
 go func(){cell=2;workers.Done()}()
 workers.Wait()
 if cell == 1 {return value}; return value
}
`

func TestRaceOnlyDefectIsDetectedAndHasPublicFeedback(t *testing.T) {
	// A disposable session compiles a cold race-enabled standard library on CI.
	// Keep that infrastructure cost separate from the fixture's behavior checks.
	const commandTimeoutSeconds = 60
	root := fixtureRepository(t)
	id := "synthetic-fixed-v2-001"
	source := filepath.Join(repositoryRoot(t), "internal/testdata/tasking/v2/fixed", id)
	dest := filepath.Join(root, "tasks/drafts", id)
	if err := isolation.CopyTree(source, dest); err != nil {
		t.Fatal(err)
	}
	for _, state := range []State{StateInReview, StateReady} {
		if _, err := promoteForTest(t, root, PromoteOptions{ID: id, To: state, Actor: "fixture", Reason: "race fixture", Date: "2026-09-09"}); err != nil {
			t.Fatal(err)
		}
	}
	dest = filepath.Join(root, "tasks/ready", id)
	taskPath := filepath.Join(dest, "task.yaml")
	var task TaskV2
	if err := decodeStrictFile(taskPath, &task); err != nil {
		t.Fatal(err)
	}
	task.Verification.RaceRequired = true
	task.Verification.Limits.CommandTimeoutSeconds = commandTimeoutSeconds
	content, err := yaml.Marshal(task)
	if err != nil {
		t.Fatal(err)
	}
	mustWrite(t, taskPath, content)
	mutant := filepath.Join(dest, "solution/mutants/constant-zero/identity.go")
	mustWrite(t, mutant, []byte(racingIdentity))
	digest, err := TaskContentDigest(dest)
	if err != nil {
		t.Fatal(err)
	}
	var provenance ProvenanceV2
	provPath := filepath.Join(dest, "provenance.yaml")
	if err := decodeStrictFile(provPath, &provenance); err != nil {
		t.Fatal(err)
	}
	provenance.Project.ContentDigest = digest
	content, err = yaml.Marshal(provenance)
	if err != nil {
		t.Fatal(err)
	}
	mustWrite(t, provPath, content)
	session, err := NewCheckSession(root)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	loaded, err := loadLearnerPackage(root, id)
	if err != nil {
		t.Fatal(err)
	}
	ordinary := variantRequest{TaskID: id, GoVersion: learnerGoVersion(loaded), TimeoutSeconds: commandTimeoutSeconds, Implementation: []string{mutant}, Tests: appendCopy(loaded.layout.PublicTests, loaded.layout.GraderTests)}
	if result := session.runner.RunVariant(ordinary); !result.Passed {
		t.Fatalf("race-only defect must preserve ordinary behavior: %+v", result)
	}
	report, err := session.Check(id)
	if err != nil || !report.Passed || report.Mutants.Detected != 1 {
		t.Fatalf("race-only defect not detected: %+v %v", report, err)
	}
	workspace := t.TempDir()
	files, err := LearnerWorkspaceFiles(root, id)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		path := filepath.Join(workspace, filepath.FromSlash(file.Path))
		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		mustWrite(t, path, file.Content)
	}
	mustWrite(t, filepath.Join(workspace, "subject/identity.go"), []byte(racingIdentity))
	result, err := VerifyLearnerWorkspace(root, id, workspace, LearnerPublicChecks)
	if err != nil || result.Passed || !strings.Contains(result.Diagnostic, "DATA RACE") {
		t.Fatalf("public race feedback missing: %+v %v", result, err)
	}
	reference, err := os.ReadFile(filepath.Join(dest, "solution/reference/identity.go"))
	if err != nil {
		t.Fatal(err)
	}
	mustWrite(t, filepath.Join(workspace, "subject/identity.go"), reference)
	result, err = VerifyLearnerWorkspace(root, id, workspace, LearnerPublicChecks)
	if err != nil || !result.Passed {
		t.Fatalf("race-free public reference rejected: %+v %v", result, err)
	}
}
