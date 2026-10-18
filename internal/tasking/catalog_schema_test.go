package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

func TestCatalogCompilesOnceButValidatesEveryPackage(t *testing.T) {
	root := learnerCatalogFixture(t)
	compiled := map[string]int{}
	validated := 0
	context := newCatalogSchemaValidator(func(name string) (*jsonschema.Schema, error) {
		compiled[name]++
		return compileSchema(root, name)
	})
	validate := func(name string, value any) error {
		validated++
		return context(name, value)
	}
	for _, id := range []string{"clamp-tests-001", "zero-results-001"} {
		located, err := FindTask(root, id)
		if err != nil {
			t.Fatal(err)
		}
		if _, problems := loadPackageWithSchemaValidator(root, located, validate); len(problems) != 0 {
			t.Fatalf("%s: %v", id, problems)
		}
	}
	if len(compiled) != 4 || validated != 8 {
		t.Fatalf("expected four schemas and all eight document validations: compiled=%v validated=%d", compiled, validated)
	}
	for name, count := range compiled {
		if count != 1 {
			t.Fatalf("%s compiled %d times", name, count)
		}
	}
	// A later package cannot inherit the result of an earlier valid document.
	loader := newLearnerPackageLoader(root)
	if _, err := loader.load("clamp-tests-001"); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "tasks/ready/zero-results-001/task.yaml")
	var task TaskV2
	if err := decodeStrictFile(path, &task); err != nil {
		t.Fatal(err)
	}
	task.Activity = "invalid-activity"
	data, err := yaml.Marshal(task)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := loader.load("zero-results-001"); err == nil || !strings.Contains(err.Error(), "zero-results-001") || !strings.Contains(err.Error(), "task-v2.schema.json violation") {
		t.Fatalf("later invalid package bypassed its schema: %v", err)
	}
}

func TestCatalogSchemaContextKeepsConstraintsAndDiagnostics(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "schemas"), 0700); err != nil {
		t.Fatal(err)
	}
	name := "catalog-fixture.schema.json"
	data := []byte(`{"$schema":"https://json-schema.org/draft/2020-12/schema","type":"object","additionalProperties":false,"required":["kind","email"],"properties":{"kind":{"enum":["one"]},"email":{"type":"string","format":"email"}}}`)
	if err := os.WriteFile(filepath.Join(root, "schemas", name), data, 0600); err != nil {
		t.Fatal(err)
	}
	validate := newLearnerPackageLoader(root).schema
	for _, tc := range []struct {
		name  string
		value map[string]any
		valid bool
	}{
		{"valid", map[string]any{"kind": "one", "email": "member@example.org"}, true},
		{"missing", map[string]any{"kind": "one"}, false},
		{"unknown", map[string]any{"kind": "one", "email": "member@example.org", "extra": true}, false},
		{"enum", map[string]any{"kind": "other", "email": "member@example.org"}, false},
		{"format", map[string]any{"kind": "one", "email": "invalid"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, fresh := validate(name, tc.value), validateSchema(root, name, tc.value)
			if (got == nil) != tc.valid || (fresh == nil) != tc.valid {
				t.Fatalf("cached=%v fresh=%v, valid=%t", got, fresh, tc.valid)
			}
			if got != nil && got.Error() != fresh.Error() {
				t.Fatalf("diagnostics changed: cached=%v fresh=%v", got, fresh)
			}
		})
	}
}

func TestCatalogSchemaContextsDoNotCrossRoots(t *testing.T) {
	a, b := learnerCatalogFixture(t), learnerCatalogFixture(t)
	if _, err := ListReadyLearnerTasks(a); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(b, "schemas/task-v2.schema.json"), []byte("false"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := ListReadyLearnerTasks(b); err == nil || !strings.Contains(err.Error(), "task.yaml") {
		t.Fatalf("root B reused A's schema: %v", err)
	}
	if _, err := ListReadyLearnerTasks(a); err != nil {
		t.Fatalf("root B poisoned A: %v", err)
	}
}

func TestCatalogSchemaCompileErrorsKeepDiagnostics(t *testing.T) {
	for _, data := range []string{"{", `{"$schema":"https://json-schema.org/draft/2020-12/schema","$ref":"#/$defs/missing"}`} {
		root := t.TempDir()
		if err := os.MkdirAll(filepath.Join(root, "schemas"), 0700); err != nil {
			t.Fatal(err)
		}
		name := "invalid-fixture.schema.json"
		if err := os.WriteFile(filepath.Join(root, "schemas", name), []byte(data), 0600); err != nil {
			t.Fatal(err)
		}
		got := newLearnerPackageLoader(root).schema(name, map[string]any{})
		fresh := validateSchema(root, name, map[string]any{})
		if got == nil || fresh == nil || got.Error() != fresh.Error() || !strings.Contains(got.Error(), name) {
			t.Fatalf("compile/ref error context changed: cached=%v fresh=%v", got, fresh)
		}
	}
}

func TestCatalogLazySchemasKeepEmptyLegacyAndStandaloneBehavior(t *testing.T) {
	root := t.TempDir()
	if tasks, err := ListReadyLearnerTasks(root); err != nil || len(tasks) != 0 {
		t.Fatalf("empty catalog required schemas: %v", err)
	}
	path := filepath.Join(root, "tasks/ready/legacy-only-001")
	if err := os.MkdirAll(path, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(path, "task.yaml"), []byte("schema_version: 1\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if tasks, err := ListReadyLearnerTasks(root); err != nil || len(tasks) != 0 {
		t.Fatalf("legacy-only catalog required v2 schemas: %v", err)
	}
	root = learnerCatalogFixture(t)
	if err := os.MkdirAll(filepath.Join(root, "tasks/drafts/unrelated-invalid-entry"), 0700); err != nil {
		t.Fatal(err)
	}
	if _, err := ListReadyLearnerTasks(root); err == nil {
		t.Fatal("catalog ignored malformed lifecycle inventory")
	}
	if _, err := LoadReadyLearnerTask(root, "zero-results-001"); err != nil {
		t.Fatalf("standalone lookup acquired unrelated inventory dependency: %v", err)
	}
}
