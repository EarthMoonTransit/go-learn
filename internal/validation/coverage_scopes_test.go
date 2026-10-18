package validation

import (
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/taxonomy"
	"gopkg.in/yaml.v3"
)

func TestCoverageScopeSchemaRejectsMalformedContract(t *testing.T) {
	root := repositoryRoot(t)
	schema, err := compileYAMLSchema(filepath.Join(root, filepath.FromSlash(taxonomy.ScopeSchemaPath)))
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(taxonomy.ScopePath))
	if err := validateYAMLSchema(path, schema, "coverage scopes"); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		change func(map[string]any)
	}{
		{"missing version", func(d map[string]any) { delete(d, "schema_version") }},
		{"wrong taxonomy", func(d map[string]any) { d["taxonomy_version"] = 3 }},
		{"unreviewed decision", func(d map[string]any) { d["decision"] = "unreviewed" }},
		{"unknown property", func(d map[string]any) { d["default"] = "foundation" }},
		{"missing assignment", func(d map[string]any) { d["foundation"] = d["foundation"].([]any)[1:] }},
		{"duplicate assignment", func(d map[string]any) { ids := d["foundation"].([]any); ids[1] = ids[0] }},
		{"wildcard assignment", func(d map[string]any) { d["specialist"].([]any)[0] = "*" }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var d map[string]any
			if err := decodeStrict(path, &d); err != nil {
				t.Fatal(err)
			}
			tc.change(d)
			data, err := yaml.Marshal(d)
			if err != nil {
				t.Fatal(err)
			}
			mutated := filepath.Join(t.TempDir(), "scopes.yaml")
			if err := os.WriteFile(mutated, data, 0600); err != nil {
				t.Fatal(err)
			}
			if err := validateYAMLSchema(mutated, schema, "coverage scopes"); err == nil {
				t.Fatal("malformed scope contract accepted")
			}
		})
	}
}
