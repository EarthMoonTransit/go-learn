package tasking

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func learnerCatalogFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	for _, relative := range []string{"config", "schemas", "sources", "curriculum", "tasks/ready/zero-results-001", "tasks/ready/clamp-tests-001"} {
		if err := isolation.CopyTree(filepath.Join(repositoryRoot(t), relative), filepath.Join(root, relative)); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func TestLearnerCatalogMatchesIndividualValidation(t *testing.T) {
	root := learnerCatalogFixture(t)
	tasks, err := ListReadyLearnerTasks(root)
	if err != nil || len(tasks) != 2 {
		t.Fatalf("catalog: %+v %v", tasks, err)
	}
	for _, task := range tasks {
		individual, err := LoadReadyLearnerTask(root, task.ID)
		if err != nil || !reflect.DeepEqual(task, individual) {
			t.Fatalf("catalog changed public metadata for %s: %v", task.ID, err)
		}
	}
	// The returned slice and its nested fields are not shared with later reads.
	tasks[0].Prerequisites = append(tasks[0].Prerequisites, "invented.skill")
	next, err := ListReadyLearnerTasks(root)
	if err != nil || reflect.DeepEqual(tasks, next) {
		t.Fatalf("catalog results leaked across reads: %v", err)
	}
}

func TestLearnerCatalogReloadsSharedPolicyAndTaskInputs(t *testing.T) {
	root := learnerCatalogFixture(t)
	for _, fixture := range []struct{ path, replacement, diagnostic string }{
		{"curriculum/skills.yaml", "unknown_field: true\n", "cannot load taxonomy"},
		{"sources/registry.yaml", "version: -1\n", "source registry"},
		{"config/project-identity.yaml", "unknown_field: true\n", "project identity"},
		{"config/outbound-license.yaml", "unknown_field: true\n", "outbound license policy"},
		{"schemas/task-v2.schema.json", "false", "task.yaml"},
		{"tasks/ready/zero-results-001/TASK.md", "changed content", "content_digest"},
	} {
		t.Run(fixture.path, func(t *testing.T) {
			if _, err := ListReadyLearnerTasks(root); err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(root, fixture.path)
			original, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if err := os.WriteFile(path, original, 0600); err != nil {
					t.Error(err)
				}
			})
			replacement := []byte(fixture.replacement)
			if strings.HasSuffix(fixture.path, "TASK.md") {
				replacement = append(original, []byte("\nAdditional technical requirement.\n")...)
			}
			if err := os.WriteFile(path, replacement, 0600); err != nil {
				t.Fatal(err)
			}
			if _, err := ListReadyLearnerTasks(root); err == nil || !strings.Contains(err.Error(), fixture.diagnostic) {
				t.Fatalf("stale catalog accepted changed %s: %v", fixture.path, err)
			}
		})
	}
}

func TestLearnerCatalogRejectsDuplicateAndSymlinkAfterRead(t *testing.T) {
	for _, kind := range []string{"duplicate", "symlink"} {
		t.Run(kind, func(t *testing.T) {
			root := learnerCatalogFixture(t)
			if _, err := ListReadyLearnerTasks(root); err != nil {
				t.Fatal(err)
			}
			if kind == "duplicate" {
				if err := os.MkdirAll(filepath.Join(root, "tasks/drafts/zero-results-001"), 0700); err != nil {
					t.Fatal(err)
				}
			} else {
				path := filepath.Join(root, "tasks/ready/zero-results-001/untrusted.go")
				if err := os.Symlink(filepath.Join(t.TempDir(), "outside"), path); err != nil {
					t.Fatal(err)
				}
			}
			if _, err := ListReadyLearnerTasks(root); err == nil {
				t.Fatal("catalog reused stale validation")
			}
		})
	}
}

// Reproducible metadata-only benchmark; task source is read as data.
func BenchmarkReadyLearnerCatalog(b *testing.B) {
	root, err := filepath.Abs("../..")
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if _, err := ListReadyLearnerTasks(root); err != nil {
			b.Fatal(err)
		}
	}
}
