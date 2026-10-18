//go:build taskintegration

package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestNativeFuzzCLIPreservesExistingLearnerHistory(t *testing.T) {
	repository := learnerFixtureRepository(t)
	home := t.TempDir()
	service, err := New(Config{Root: repository, Home: home})
	if err != nil {
		t.Fatal(err)
	}
	service.verify = integrationVerifier(t, repository)
	const oldID = "synthetic-fixed-v2-001"
	oldWorkspace, err := service.Start(oldID)
	if err != nil {
		t.Fatal(err)
	}
	reference, err := os.ReadFile(filepath.Join(repository, "tasks/ready", oldID, "solution/reference/identity.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(oldWorkspace, "subject/identity.go"), reference, 0600); err != nil {
		t.Fatal(err)
	}
	if result, err := service.Complete(); err != nil || !result.Passed {
		t.Fatalf("old profile completion: %+v %v", result, err)
	}
	if err := service.Feedback("retain this existing completion"); err != nil {
		t.Fatal(err)
	}
	if err := service.Pause(); err != nil {
		t.Fatal(err)
	}
	before, err := service.Progress()
	if err != nil {
		t.Fatal(err)
	}
	oldTask, err := json.Marshal(before.Tasks[oldID])
	if err != nil {
		t.Fatal(err)
	}
	oldFiles := map[string][]byte{}
	for _, name := range workspacePaths(t, oldWorkspace) {
		path := filepath.Join(oldWorkspace, filepath.FromSlash(name))
		info, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		if info.IsDir() {
			continue
		}
		body, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		oldFiles[name] = body
	}
	var stdout, stderr bytes.Buffer
	run := func(want int, args ...string) {
		t.Helper()
		stdout.Reset()
		stderr.Reset()
		arguments := append([]string{"-root", repository, "-home", home}, args...)
		if code := RunCLI(Config{}, arguments, &stdout, &stderr); code != want {
			t.Fatalf("%v: code=%d want=%d stdout=%s stderr=%s", args, code, want, stdout.String(), stderr.String())
		}
	}
	const id = "synthetic-hex-v2-001"
	run(0, "show", id)
	if !strings.Contains(stdout.String(), "100,000-step") || !strings.Contains(stdout.String(), "8192 records") {
		t.Fatal("CLI omitted finite source contract")
	}
	run(0, "start", id)
	workspace := filepath.Join(home, "workspaces", id)
	assertLearnerWorkspaceSafe(t, workspace)
	contractPath := filepath.Join(workspace, "FUZZ_SOURCE.md")
	contract, err := os.ReadFile(contractPath)
	if err != nil {
		t.Fatal(err)
	}
	// Deliberately simulate an editor overriding the filesystem readonly bit;
	// content validation must still reject the changed contract.
	if err := os.Chmod(contractPath, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(contractPath, append(append([]byte{}, contract...), '\n'), 0600); err != nil {
		t.Fatal(err)
	}
	run(2, "check")
	if !strings.Contains(stderr.String(), "read-only") {
		t.Fatalf("missing readonly grammar rejection: %s", stderr.String())
	}
	if err := os.WriteFile(contractPath, contract, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(contractPath, 0444); err != nil {
		t.Fatal(err)
	}
	run(0, "check")
	run(1, "complete")
	if strings.Contains(stdout.String(), "paired-uppercase") || strings.Contains(stdout.String(), "false-status") {
		t.Fatal("CLI exposed private defect identities")
	}
	reference, err = os.ReadFile(filepath.Join(repository, "tasks/ready", id, "solution/reference/fuzz_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(workspace, "subject/fuzz_test.go"), reference, 0600); err != nil {
		t.Fatal(err)
	}
	run(0, "complete")
	after, err := service.Progress()
	if err != nil {
		t.Fatal(err)
	}
	oldAfter, err := json.Marshal(after.Tasks[oldID])
	if err != nil {
		t.Fatal(err)
	}
	if len(after.Completions) < len(before.Completions) || !bytes.Equal(oldTask, oldAfter) || !reflect.DeepEqual(before.Completions, after.Completions[:len(before.Completions)]) {
		t.Fatal("new native checks changed old completion/history")
	}
	if !after.Tasks[id].Completed || after.Tasks[id].Attempts != 3 {
		t.Fatalf("native attempts/completion lost: %+v", after.Tasks[id])
	}
	for name, want := range oldFiles {
		got, err := os.ReadFile(filepath.Join(oldWorkspace, filepath.FromSlash(name)))
		if err != nil || !bytes.Equal(want, got) {
			t.Fatalf("old workspace changed at %s: %v", name, err)
		}
	}
}
