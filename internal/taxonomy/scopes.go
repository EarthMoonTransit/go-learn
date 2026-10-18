package taxonomy

import "fmt"

const (
	ScopePath       = "curriculum/coverage-scopes.yaml"
	ScopeSchemaPath = "schemas/coverage-scopes.schema.json"
)

type coverageScopes struct {
	SchemaVersion   int      `yaml:"schema_version"`
	TaxonomyVersion int      `yaml:"taxonomy_version"`
	Decision        string   `yaml:"decision"`
	Foundation      []string `yaml:"foundation"`
	Specialist      []string `yaml:"specialist"`
}

// readCoverageScopes validates the explicit classification of the already
// checked finite v4 inventory. It supplies no default for omitted or future IDs.
func readCoverageScopes(root string, live map[string]map[string]any) (coverageScopes, error) {
	var scopes coverageScopes
	if _, err := readStrict(root, ScopePath, &scopes); err != nil {
		return scopes, err
	}
	if scopes.SchemaVersion != 1 || scopes.TaxonomyVersion != 4 || scopes.Decision != Decision {
		return scopes, fmt.Errorf("coverage scopes: expected versioned ADR0031 v4 classification")
	}
	assignments := map[string]string{}
	for _, group := range []struct {
		name string
		ids  []string
	}{{"foundation", scopes.Foundation}, {"specialist", scopes.Specialist}} {
		for _, id := range group.ids {
			skill := live[id]
			if skill == nil {
				return scopes, fmt.Errorf("coverage scopes: unknown skill %q", id)
			}
			if skill["status"] != "active" || skill["assignable"] != true {
				return scopes, fmt.Errorf("coverage scopes: deprecated/non-assignable skill %q", id)
			}
			if assignments[id] != "" {
				return scopes, fmt.Errorf("coverage scopes: duplicate skill %q", id)
			}
			assignments[id] = group.name
		}
	}
	for _, id := range sortedKeys(live) {
		skill := live[id]
		if skill["status"] != "active" || skill["assignable"] != true {
			continue
		}
		scope := assignments[id]
		if scope == "" {
			return scopes, fmt.Errorf("coverage scopes: missing skill %q", id)
		}
		specialist := false
		switch id {
		case "sync.rwmutex-tradeoffs", "sync.atomics", "sync.cond", "sync.pool":
			specialist = true
		}
		if (scope == "specialist") != specialist {
			return scopes, fmt.Errorf("coverage scopes: unapproved scope %q for %s", scope, id)
		}
		if scope != "foundation" {
			continue
		}
		prerequisites, ok := skill["prerequisites"].([]any)
		if !ok {
			return scopes, fmt.Errorf("coverage scopes: invalid prerequisites for %s", id)
		}
		for _, raw := range prerequisites {
			prerequisite, ok := raw.(string)
			if !ok || assignments[prerequisite] != "foundation" {
				return scopes, fmt.Errorf("coverage scopes: foundation skill %s requires non-foundation prerequisite %v", id, raw)
			}
		}
	}
	return scopes, nil
}
