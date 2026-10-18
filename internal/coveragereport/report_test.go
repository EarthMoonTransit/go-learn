package coveragereport

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/taxonomy"
	"go-practice-bank/internal/testrepo"
	"go-practice-bank/internal/validation"
	"gopkg.in/yaml.v3"
)

func repositoryRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func TestReportsRequireEligibleEvidenceAndKeepBothScopes(t *testing.T) {
	root := testrepo.Copy(t, repositoryRoot(t))
	checks := testrepo.Collection(t, root)
	testrepo.WriteArtifacts(t, root, checks)
	reports := map[string]Report{}
	for _, scope := range []string{"full", "foundation"} {
		r, err := Generate(root, scope, checks, false)
		if err != nil {
			t.Fatal(err)
		}
		reports[scope] = r
		if r.SelectedNumericPassed || r.Numeric.FullNumericPassed || r.Numeric.Foundation.MinimaMet || r.SemanticReviewStatus != SemanticStatus || len(r.Numeric.Skills) != 145 || r.Numeric.ReadyTarget.Min != 350 || r.Numeric.ReadyTarget.Max != 500 || r.Numeric.Full.PrimaryDeficit == 0 || r.Numeric.Specialist.PrimaryDeficit == 0 {
			t.Fatalf("incomplete-bank report hid deficits: %+v", r)
		}
		want := 145
		if scope == "foundation" {
			want = 141
		}
		if len(r.SelectedSkillIDs) != want {
			t.Fatalf("%s selection count %d", scope, len(r.SelectedSkillIDs))
		}
		md, yp, _ := Paths(scope)
		content, err := os.ReadFile(filepath.Join(root, yp))
		if err != nil {
			t.Fatal(err)
		}
		var decoded Report
		if err := yaml.Unmarshal(content, &decoded); err != nil || !reflect.DeepEqual(r, decoded) {
			t.Fatalf("YAML differs from report: %v", err)
		}
		markdown, err := os.ReadFile(filepath.Join(root, md))
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(markdown, MarshalMarkdown(decoded)) {
			t.Fatal("Markdown and YAML disagree")
		}
		for _, id := range r.SelectedSkillIDs {
			if !bytes.Contains(markdown, []byte("- `"+id+"`")) {
				t.Fatalf("Markdown lost selected %s", id)
			}
		}
	}
	if !reflect.DeepEqual(reports["full"].Numeric, reports["foundation"].Numeric) || reports["full"].InputIdentity == reports["foundation"].InputIdentity {
		t.Fatal("scope selection changed full arithmetic or lost identity distinction")
	}
	// A check succeeds for a valid, intentionally incomplete bank. It also
	// demonstrates that generated outputs are not their own collection inputs.
	if _, err := Generate(root, "full", checks, true); err != nil {
		t.Fatal(err)
	}
	schema, err := os.ReadFile(filepath.Join(root, SchemaPath))
	if err != nil {
		t.Fatal(err)
	}
	bad := reports["full"]
	bad.SemanticReviewStatus = "pass"
	if err := validateSchema(schema, bad); err == nil {
		t.Fatal("schema accepted fabricated semantic pass")
	}
	bad = reports["full"]
	bad.Numeric.Skills = bad.Numeric.Skills[:144]
	if err := validateSchema(schema, bad); err == nil {
		t.Fatal("schema accepted missing active row")
	}

	// Snapshot both scopes. Every invalid-input call below must fail before
	// writing, including when ordinary validation returned a partial result.
	before := map[string][]byte{}
	for _, scope := range []string{"full", "foundation"} {
		md, yp, _ := Paths(scope)
		for _, p := range []string{md, yp} {
			before[p], err = os.ReadFile(filepath.Join(root, p))
			if err != nil {
				t.Fatal(err)
			}
		}
	}
	assertUnchanged := func() {
		t.Helper()
		for p, want := range before {
			got, err := os.ReadFile(filepath.Join(root, p))
			if err != nil || !bytes.Equal(got, want) {
				t.Fatalf("invalid evidence modified %s: %v", p, err)
			}
		}
	}
	// Exercise both stale artifact paths using the already validated report.
	// No additional full-bank metadata pass is needed for byte comparison.
	md, yp, _ := Paths("full")
	for _, p := range []string{md, yp} {
		changed := append(slices.Clone(before[p]), '\n')
		if err := os.WriteFile(filepath.Join(root, p), changed, 0600); err != nil {
			t.Fatal(err)
		}
		if err := publish(root, reports["full"], true); err == nil || !strings.Contains(err.Error(), p+" is stale") {
			t.Fatalf("stale output accepted: %v", err)
		}
		for other, want := range before {
			if other == p {
				want = changed
			}
			got, err := os.ReadFile(filepath.Join(root, other))
			if err != nil || !bytes.Equal(got, want) {
				t.Fatalf("check modified %s", other)
			}
		}
		if err := publish(root, reports["full"], false); err != nil {
			t.Fatal(err)
		}
		assertUnchanged()
	}
	mod, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module invalid\n\ngo 1.27\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := Generate(root, "foundation", checks, false); err == nil || !strings.Contains(err.Error(), "root module must be") {
		t.Fatalf("ordinary error consumed a partial result: %v", err)
	}
	assertUnchanged()
	if err := os.WriteFile(filepath.Join(root, "go.mod"), mod, 0600); err != nil {
		t.Fatal(err)
	}
	failed := checks
	failed.Tasks = slices.Clone(checks.Tasks)
	r := &failed.Tasks[0].Report
	r.Passed, r.PromotionEligible = false, false
	r.TechnicalStatus, r.ReferenceStatus = "fail", "fail"
	r.FailingGateIDs = []string{"reference-full"}
	r.Gates = []tasking.Gate{{ID: "reference-full", Status: "fail", Details: []string{"synthetic current failure"}}}
	path := testrepo.WriteArtifacts(t, root, failed)
	loaded, err := tasking.LoadCheckCollection(root, path)
	if err != nil {
		t.Fatal(err)
	}
	// The paired validation-package fixture separately asserts ordinary nil;
	// Build's failed-package diagnostic proves this path reached that boundary.
	if _, err := Generate(root, "full", loaded, false); err == nil || !strings.Contains(err.Error(), "failed package reports:") {
		t.Fatalf("fresh failed collection accepted or masked by unrelated errors: %v", err)
	}
	assertUnchanged()
	if err := os.WriteFile(filepath.Join(root, SchemaPath), append(schema, '\n'), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := Generate(root, "foundation", checks, false); err == nil || !strings.Contains(err.Error(), "stale check collection: task input digest changed") {
		t.Fatalf("own-schema edit did not invalidate evidence: %v", err)
	}
	assertUnchanged()
}

func TestV3OrdinaryAndFullCompatibilityRejectNewConsumers(t *testing.T) {
	root := testrepo.Copy(t, repositoryRoot(t))
	for _, scope := range []string{"full", "foundation"} {
		md, yp, _ := Paths(scope)
		for _, p := range []string{md, yp} {
			if err := os.Remove(filepath.Join(root, p)); err != nil && !os.IsNotExist(err) {
				t.Fatal(err)
			}
		}
	}
	v3, err := os.ReadFile(filepath.Join(root, taxonomy.V3Path))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "curriculum/skills.yaml"), v3, 0600); err != nil {
		t.Fatal(err)
	}
	checks := testrepo.Collection(t, root)
	testrepo.WriteArtifacts(t, root, checks)
	ordinary, err := validation.ValidateRoot(root, validation.Options{Checks: &checks})
	if err != nil || len(ordinary.Coverage) != 142 || ordinary.NumericCoverage != nil {
		t.Fatalf("v3 ordinary contract changed: %v", err)
	}
	full, err := validation.ValidateRoot(root, validation.Options{EnforceCoverage: true, Checks: &checks})
	if err == nil || !strings.Contains(err.Error(), "ready task count") || len(full.Coverage) != 142 || strings.Contains(err.Error(), "requires taxonomy v4") {
		t.Fatalf("v3 full enforcement changed: %v", err)
	}
	_, err = validation.ValidateRoot(root, validation.Options{EnforceFoundationCoverage: true, Checks: &checks})
	if err == nil || !strings.Contains(err.Error(), "requires taxonomy v4") {
		t.Fatalf("v3 foundation accepted: %v", err)
	}
	if _, err := Generate(root, "full", checks, false); err == nil || !strings.Contains(err.Error(), "v4 scopes required") {
		t.Fatalf("v3 report accepted: %v", err)
	}
	md, yp, _ := Paths("full")
	for _, p := range []string{md, yp} {
		if _, err := os.Lstat(filepath.Join(root, p)); !os.IsNotExist(err) {
			t.Fatalf("v3 report wrote %s", p)
		}
	}
}

func TestCanonicalIdentityUsesOnlyDeclaredStableInputs(t *testing.T) {
	n := validation.NumericCoverage{TaxonomyVersion: 4, Skills: []validation.NumericSkillCoverage{{SkillID: "b.two", Scope: "specialist"}, {SkillID: "a.one", Scope: "foundation"}}}
	checks := tasking.CheckCollection{ContractVersion: "contract", ToolchainPolicyDigest: "policy", Tasks: []tasking.CheckedTask{{TaskID: "b-001", State: tasking.StateReady, InputDigest: "b"}, {TaskID: "a-001", State: tasking.StateDraft, InputDigest: "a"}}}
	schema := []byte("schema bytes")
	identity, err := inputIdentity("full", n, checks, schema)
	if err != nil {
		t.Fatal(err)
	}
	if identity != "fe014bc5a62444aaeaf6a9e0bb676a8901c85561b0066df38f5f6b1107c67e1d" {
		t.Fatalf("canonical encoding drift: %s", identity)
	}
	slices.Reverse(n.Skills)
	slices.Reverse(checks.Tasks)
	checks.Execution = tasking.ExecutionMetrics{Sessions: 1, Workspaces: 900, ChildGoRuns: 900, VariantRuns: 900}
	checks.Tasks[0].Report.Gates = []tasking.Gate{{ID: "details", Status: "pass", Details: []string{"/absolute/temporary/path", "timestamp"}}}
	again, err := inputIdentity("full", n, checks, schema)
	if err != nil || again != identity {
		t.Fatal("permutation/execution details changed canonical identity")
	}
	for _, tc := range []struct {
		name   string
		change func(*validation.NumericCoverage, *tasking.CheckCollection, *[]byte, *string)
	}{
		{"scope", func(_ *validation.NumericCoverage, _ *tasking.CheckCollection, _ *[]byte, s *string) {
			*s = "foundation"
		}},
		{"classification", func(n *validation.NumericCoverage, _ *tasking.CheckCollection, _ *[]byte, _ *string) {
			n.Skills[0].Scope = "specialist"
		}},
		{"taxonomy", func(n *validation.NumericCoverage, _ *tasking.CheckCollection, _ *[]byte, _ *string) {
			n.TaxonomyVersion = 5
		}},
		{"schema", func(_ *validation.NumericCoverage, _ *tasking.CheckCollection, s *[]byte, _ *string) {
			*s = []byte("changed")
		}},
		{"contract", func(_ *validation.NumericCoverage, c *tasking.CheckCollection, _ *[]byte, _ *string) {
			c.ContractVersion = "changed"
		}},
		{"policy", func(_ *validation.NumericCoverage, c *tasking.CheckCollection, _ *[]byte, _ *string) {
			c.ToolchainPolicyDigest = "changed"
		}},
		{"task state", func(_ *validation.NumericCoverage, c *tasking.CheckCollection, _ *[]byte, _ *string) {
			c.Tasks[0].State = tasking.StateRejected
		}},
		{"task digest", func(_ *validation.NumericCoverage, c *tasking.CheckCollection, _ *[]byte, _ *string) {
			c.Tasks[0].InputDigest = "changed"
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			value, c, s, scope := n, checks, slices.Clone(schema), "full"
			value.Skills = slices.Clone(n.Skills)
			c.Tasks = slices.Clone(checks.Tasks)
			tc.change(&value, &c, &s, &scope)
			got, err := inputIdentity(scope, value, c, s)
			if err != nil || got == identity {
				t.Fatalf("declared input excluded: %v", err)
			}
		})
	}
}
