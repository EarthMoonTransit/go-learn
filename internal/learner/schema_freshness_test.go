package learner

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestSameServiceObservesSchemaErrorsAndRecovery(t *testing.T) {
	root := t.TempDir()
	repository, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	// Metadata and authored files are copied only as data; ListTasks executes no
	// task code. Use one unchanged real task and a synthetic learner home.
	for _, path := range []string{"config", "schemas", "sources", "curriculum", "tasks/ready/zero-results-001"} {
		if err := isolation.CopyTree(filepath.Join(repository, path), filepath.Join(root, path)); err != nil {
			t.Fatal(err)
		}
	}
	s, err := New(Config{Root: root, Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	if tasks, _, err := s.ListTasks(); err != nil || len(tasks) != 1 {
		t.Fatalf("initial catalog: %v", err)
	}
	before := homeSnapshot(t, s.home)
	path := filepath.Join(root, "schemas/task-v2.schema.json")
	original, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, invalid := range []string{"false", "{"} {
		if err := os.WriteFile(path, []byte(invalid), 0600); err != nil {
			t.Fatal(err)
		}
		if _, _, err := s.ListTasks(); err == nil || !strings.Contains(err.Error(), "task.yaml") {
			t.Fatalf("same Service hid schema error %q: %v", invalid, err)
		}
		if err := os.WriteFile(path, original, 0600); err != nil {
			t.Fatal(err)
		}
		if tasks, _, err := s.ListTasks(); err != nil || len(tasks) != 1 {
			t.Fatalf("schema recovery hidden by Service: %v", err)
		}
	}
	if !reflect.DeepEqual(before, homeSnapshot(t, s.home)) {
		t.Fatal("read-only schema validation changed learner state")
	}
}
