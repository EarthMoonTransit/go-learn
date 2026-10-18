package taxonomy

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestExplicitScopesPreserveApprovedInventoryAndArithmetic(t *testing.T) {
	root := fixture(t)
	var live document
	if _, err := readStrict(root, "curriculum/skills.yaml", &live); err != nil {
		t.Fatal(err)
	}
	all, err := skills(live)
	if err != nil {
		t.Fatal(err)
	}
	scopes, err := readCoverageScopes(root, all)
	if err != nil {
		t.Fatal(err)
	}
	for _, group := range []struct {
		name                  string
		ids                   []string
		count, primary, total int
	}{{"foundation", scopes.Foundation, 141, 410, 692}, {"specialist", scopes.Specialist, 4, 8, 16}} {
		primary, total := 0, 0
		for _, id := range group.ids {
			minimum := all[id]["minimum_occurrences"].(map[string]any)
			primary += minimum["primary"].(int)
			total += minimum["total"].(int)
		}
		if len(group.ids) != group.count || primary != group.primary || total != group.total {
			t.Fatalf("%s = %d skills / %d primary / %d total", group.name, len(group.ids), primary, total)
		}
	}
	for _, id := range []string{"expressions.bitwise-masks", "expressions.integer-shifts", "functions.recursion", "sync.once-initialization"} {
		if !slices.Contains(scopes.Foundation, id) {
			t.Fatalf("required foundation classification missing: %s", id)
		}
	}
}

func TestLiveAndReviewedLoadersRejectInvalidScopes(t *testing.T) {
	cases := []struct {
		name, want string
		change     func(*coverageScopes)
	}{
		{"schema version", "expected versioned", func(s *coverageScopes) { s.SchemaVersion = 2 }},
		{"taxonomy version", "expected versioned", func(s *coverageScopes) { s.TaxonomyVersion = 3 }},
		{"decision", "expected versioned", func(s *coverageScopes) { s.Decision = "unreviewed" }},
		{"missing foundation", "missing skill", func(s *coverageScopes) { s.Foundation = s.Foundation[1:] }},
		{"missing specialist", "missing skill", func(s *coverageScopes) { s.Specialist = s.Specialist[1:] }},
		{"duplicate within scope", "duplicate skill", func(s *coverageScopes) { s.Foundation = append(s.Foundation, s.Foundation[0]) }},
		{"duplicate across scopes", "duplicate skill", func(s *coverageScopes) { s.Specialist = append(s.Specialist, s.Foundation[0]) }},
		{"unknown id", "unknown skill", func(s *coverageScopes) { s.Foundation[0] = "functions.future" }},
		{"wildcard", "unknown skill", func(s *coverageScopes) { s.Foundation[0] = "*" }},
		{"unapproved swap", "unapproved scope", func(s *coverageScopes) { s.Foundation[0], s.Specialist[0] = s.Specialist[0], s.Foundation[0] }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			if _, err := ValidateTransition(root); err != nil {
				t.Fatal(err)
			}
			var scopes coverageScopes
			rewrite(t, root, ScopePath, &scopes, func() { tc.change(&scopes) })
			assertScopeRejected(t, root, tc.want)
		})
	}
	t.Run("deprecated stub", func(t *testing.T) {
		root := fixture(t)
		var live document
		if _, err := readStrict(root, "curriculum/skills.yaml", &live); err != nil {
			t.Fatal(err)
		}
		all, err := skills(live)
		if err != nil {
			t.Fatal(err)
		}
		for _, id := range sortedKeys(all) {
			if all[id]["status"] == "deprecated" {
				var scopes coverageScopes
				rewrite(t, root, ScopePath, &scopes, func() { scopes.Foundation[0] = id })
				assertScopeRejected(t, root, "deprecated/non-assignable")
				return
			}
		}
		t.Fatal("fixture lost deprecated stubs")
	})
	for _, tc := range []struct{ name, data, want string }{
		{"unknown field", "future: true\n", "field future not found"},
		{"trailing document", "---\n{}\n", "trailing data"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			var scopes coverageScopes
			// Marshal to ordinary YAML before appending another field/document.
			rewrite(t, root, ScopePath, &scopes, func() {})
			path := filepath.Join(root, filepath.FromSlash(ScopePath))
			before, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, append(before, tc.data...), 0600); err != nil {
				t.Fatal(err)
			}
			assertScopeRejected(t, root, tc.want)
		})
	}
	t.Run("absent sidecar", func(t *testing.T) {
		root := fixture(t)
		if err := os.Remove(filepath.Join(root, filepath.FromSlash(ScopePath))); err != nil {
			t.Fatal(err)
		}
		assertScopeRejected(t, root, ScopePath)
	})
}

func assertScopeRejected(t *testing.T, root, want string) {
	t.Helper()
	for name, validate := range map[string]func(string) (Transition, error){"live": ValidateLiveTransition, "reviewed": ValidateTransition} {
		if _, err := validate(root); err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("%s loader: got %v, want %q", name, err, want)
		}
	}
}

func TestFoundationScopeRequiresPrerequisiteClosure(t *testing.T) {
	root := fixture(t)
	var live document
	if _, err := readStrict(root, "curriculum/skills.yaml", &live); err != nil {
		t.Fatal(err)
	}
	all, err := skills(live)
	if err != nil {
		t.Fatal(err)
	}
	// Exercise the scope checker independently of immutable contract hash checks.
	all["functions.recursion"]["prerequisites"] = []any{"sync.atomics"}
	if _, err := readCoverageScopes(root, all); err == nil || !strings.Contains(err.Error(), "requires non-foundation prerequisite sync.atomics") {
		t.Fatalf("foundation depended on specialist scope: %v", err)
	}
}
