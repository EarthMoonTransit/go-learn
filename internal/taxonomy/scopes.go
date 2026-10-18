package taxonomy

import "fmt"

const (
	ScopePath       = "curriculum/coverage-scopes.yaml"
	ScopeSchemaPath = "schemas/coverage-scopes.schema.json"
)

// CoverageScopes is the explicit finite v4 classification. Public loaders
// return independently owned slices; no parsed taxonomy graphs are exposed.
type CoverageScopes struct {
	SchemaVersion   int      `yaml:"schema_version"`
	TaxonomyVersion int      `yaml:"taxonomy_version"`
	Decision        string   `yaml:"decision"`
	Foundation      []string `yaml:"foundation"`
	Specialist      []string `yaml:"specialist"`
}

// LoadCoverageScopes verifies the live transition and returns a fresh, complete
// classification. Historical v3 contexts do not acquire an implicit scope.
func LoadCoverageScopes(root string) (CoverageScopes, error) {
	_, live, err := readImmutableDocument(root, "curriculum/skills.yaml")
	if err != nil {
		return CoverageScopes{}, err
	}
	if live.Version != 4 {
		return CoverageScopes{}, fmt.Errorf("coverage scopes: taxonomy version %d is unsupported; require v4", live.Version)
	}
	if _, err := ValidateLiveTransition(root); err != nil {
		return CoverageScopes{}, err
	}
	all, err := skills(live)
	if err != nil {
		return CoverageScopes{}, err
	}
	return readCoverageScopes(root, all)
}

// readCoverageScopes validates the explicit classification of the already
// checked finite v4 inventory. It supplies no default for omitted or future IDs.
func readCoverageScopes(root string, live map[string]map[string]any) (CoverageScopes, error) {
	var scopes CoverageScopes
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
