// Package coveragereport produces numeric encounter reports from an explicit
// current successful collection. It never runs authored task code.
package coveragereport

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/taxonomy"
	"go-practice-bank/internal/validation"
	"gopkg.in/yaml.v3"
)

const (
	ContractVersion = "numeric-coverage/v1"
	SchemaPath      = "schemas/coverage-report.schema.json"
	SemanticStatus  = "not_evaluated_by_numeric_coverage"
	artifactLimit   = 4 << 20
)

type Report struct {
	SchemaVersion         int                        `json:"schema_version" yaml:"schema_version"`
	ReportContract        string                     `json:"report_contract" yaml:"report_contract"`
	InputIdentity         string                     `json:"input_identity" yaml:"input_identity"`
	SelectedScope         string                     `json:"selected_scope" yaml:"selected_scope"`
	SelectedSkillIDs      []string                   `json:"selected_skill_ids" yaml:"selected_skill_ids"`
	SelectedNumericPassed bool                       `json:"selected_numeric_passed" yaml:"selected_numeric_passed"`
	Numeric               validation.NumericCoverage `json:"numeric" yaml:"numeric"`
	SemanticReviewStatus  string                     `json:"semantic_review_status" yaml:"semantic_review_status"`
}

func ValidateScope(scope string) error {
	if scope != "full" && scope != "foundation" {
		return fmt.Errorf("explicit -scope full|foundation is required")
	}
	return nil
}

// Build requires ordinary repository validity, current schema-valid checks and
// every package report passing. Numeric deficits do not invalidate a report.
// In particular, a partial validation.Result is never consumed after an error.
func Build(root, scope string, checks tasking.CheckCollection) (Report, error) {
	if err := ValidateScope(scope); err != nil {
		return Report{}, err
	}
	// Reject unsupported historical contexts before a whole-bank ordinary pass.
	// The complete ordinary/evidence guard below still applies to every report.
	if _, err := taxonomy.LoadCoverageScopes(root); err != nil {
		return Report{}, fmt.Errorf("coverage report: v4 scopes required: %w", err)
	}
	result, err := validation.ValidateRoot(root, validation.Options{Checks: &checks})
	if err != nil {
		return Report{}, fmt.Errorf("coverage report: ordinary validation: %w", err)
	}
	if err := tasking.RequirePassedCheckReports(checks); err != nil {
		return Report{}, err
	}
	if result.NumericCoverage == nil {
		return Report{}, fmt.Errorf("coverage report: %s", result.NumericUnavailable)
	}
	schema, err := safeio.ReadFile(root, SchemaPath, artifactLimit)
	if err != nil {
		return Report{}, err
	}
	n := *result.NumericCoverage
	ids, passed := n.Full.SkillIDs, n.FullNumericPassed
	if scope == "foundation" {
		ids, passed = n.Foundation.SkillIDs, n.Foundation.MinimaMet
	}
	identity, err := inputIdentity(scope, n, checks, schema)
	if err != nil {
		return Report{}, err
	}
	report := Report{SchemaVersion: 1, ReportContract: ContractVersion, InputIdentity: identity,
		SelectedScope: scope, SelectedSkillIDs: append([]string(nil), ids...), SelectedNumericPassed: passed,
		Numeric: n, SemanticReviewStatus: SemanticStatus}
	if err := validateSchema(schema, report); err != nil {
		return Report{}, err
	}
	return report, nil
}

// Identity uses the JSON encoding of these typed fields in declaration order.
// Collections must already be validated. Sorting here makes identity independent
// of incidental slice order; it does not relax collection validation ordering.
func inputIdentity(scope string, n validation.NumericCoverage, checks tasking.CheckCollection, schema []byte) (string, error) {
	type classification struct {
		SkillID string `json:"skill_id"`
		Scope   string `json:"scope"`
	}
	type task struct {
		ID          string        `json:"id"`
		State       tasking.State `json:"state"`
		InputDigest string        `json:"input_digest"`
	}
	input := struct {
		ReportContract        string           `json:"report_contract"`
		SelectedScope         string           `json:"selected_scope"`
		TaxonomyVersion       int              `json:"taxonomy_version"`
		Classification        []classification `json:"classification"`
		ReportSchemaSHA256    string           `json:"report_schema_sha256"`
		CheckContract         string           `json:"check_contract"`
		ToolchainPolicyDigest string           `json:"toolchain_policy_digest"`
		Tasks                 []task           `json:"tasks"`
	}{ReportContract: ContractVersion, SelectedScope: scope, TaxonomyVersion: n.TaxonomyVersion,
		ReportSchemaSHA256: digest(schema), CheckContract: checks.ContractVersion, ToolchainPolicyDigest: checks.ToolchainPolicyDigest}
	for _, row := range n.Skills {
		input.Classification = append(input.Classification, classification{row.SkillID, row.Scope})
	}
	sort.Slice(input.Classification, func(i, j int) bool { return input.Classification[i].SkillID < input.Classification[j].SkillID })
	for _, checked := range checks.Tasks {
		input.Tasks = append(input.Tasks, task{checked.TaskID, checked.State, checked.InputDigest})
	}
	sort.Slice(input.Tasks, func(i, j int) bool { return input.Tasks[i].ID < input.Tasks[j].ID })
	content, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return digest(content), nil
}

func digest(content []byte) string {
	value := sha256.Sum256(content)
	return hex.EncodeToString(value[:])
}

func validateSchema(content []byte, report Report) error {
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("coverage-report.schema.json", bytes.NewReader(content)); err != nil {
		return err
	}
	schema, err := compiler.Compile("coverage-report.schema.json")
	if err != nil {
		return err
	}
	b, err := json.Marshal(report)
	if err != nil {
		return err
	}
	var value any
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	if err := schema.Validate(value); err != nil {
		return fmt.Errorf("coverage report schema: %w", err)
	}
	return nil
}

func Paths(scope string) (markdown, machine string, err error) {
	if err := ValidateScope(scope); err != nil {
		return "", "", err
	}
	return "reports/coverage-" + scope + ".md", "reports/coverage-" + scope + ".yaml", nil
}

func MarshalMarkdown(r Report) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Numeric coverage: %s\n\n", r.SelectedScope)
	fmt.Fprintf(&b, "Generated by `coverage-report`, contract `%s`. Input identity: `%s`.\n\n", r.ReportContract, r.InputIdentity)
	fmt.Fprintf(&b, "Taxonomy: v%d. Selected numeric status: **%s** (%d skills).\n\n", r.Numeric.TaxonomyVersion, status(r.SelectedNumericPassed), len(r.SelectedSkillIDs))
	b.WriteString("Counts establish encounters only. Required modes, distinct task types, observable outcomes and learner mastery require independent review.\n\n")
	fmt.Fprintf(&b, "Semantic review status: `%s`.\n\n", r.SemanticReviewStatus)
	b.WriteString("| Scope | Actual/required primary | Primary deficit | Actual/required total | Total deficit | Minima |\n| --- | ---: | ---: | ---: | ---: | --- |\n")
	for _, group := range []struct {
		name string
		data validation.NumericScopeCoverage
	}{{"foundation", r.Numeric.Foundation}, {"specialist", r.Numeric.Specialist}, {"full", r.Numeric.Full}} {
		n := group.data
		fmt.Fprintf(&b, "| %s | %d/%d | %d | %d/%d | %d | %s |\n", group.name, n.ActualPrimary, n.RequiredPrimary, n.PrimaryDeficit, n.ActualTotal, n.RequiredTotal, n.TotalDeficit, status(n.MinimaMet))
	}
	t := r.Numeric.ReadyTarget
	fmt.Fprintf(&b, "\nWhole-bank ready target: **%s**, %d tasks, required %d..%d. Full-bank numeric status: **%s** (all minima and ready target). Foundation selection uses its minima only.\n\n", status(t.Met), t.Actual, t.Min, t.Max, status(r.Numeric.FullNumericPassed))
	b.WriteString("## Selected skill IDs\n\n")
	for _, id := range r.SelectedSkillIDs {
		fmt.Fprintf(&b, "- `%s`\n", id)
	}
	b.WriteString("\n## Complete active inventory\n\nDeficits are clipped per skill before summing; surplus cannot compensate another skill.\n\n| Skill | Domain | Scope | Actual/required primary | Primary deficit | Actual/required total | Total deficit |\n| --- | --- | --- | ---: | ---: | ---: | ---: |\n")
	for _, row := range r.Numeric.Skills {
		fmt.Fprintf(&b, "| %s | %s | %s | %d/%d | %d | %d/%d | %d |\n", row.SkillID, row.DomainID, row.Scope, row.ActualPrimary, row.RequiredPrimary, row.PrimaryDeficit, row.ActualTotal, row.RequiredTotal, row.TotalDeficit)
	}
	return b.Bytes()
}

func status(passed bool) string {
	if passed {
		return "pass"
	}
	return "fail"
}

// Generate validates all inputs and renders both formats before any write. Check
// compares exact bytes, and succeeds for valid reports with numeric deficits.
func Generate(root, scope string, checks tasking.CheckCollection, check bool) (Report, error) {
	report, err := Build(root, scope, checks)
	if err != nil {
		return Report{}, err
	}
	if err := publish(root, report, check); err != nil {
		return Report{}, err
	}
	return report, nil
}

// publish is called only after Build; it has no validation or execution bypass
// in the exported API. Both formats are rendered before the first write.
func publish(root string, report Report, check bool) error {
	machine, err := yaml.Marshal(report)
	if err != nil {
		return err
	}
	markdownPath, yamlPath, err := Paths(report.SelectedScope)
	if err != nil {
		return err
	}
	for _, item := range []struct {
		path    string
		content []byte
	}{{markdownPath, MarshalMarkdown(report)}, {yamlPath, machine}} {
		if check {
			actual, err := safeio.ReadFile(root, item.path, artifactLimit)
			if err != nil {
				return err
			}
			if !bytes.Equal(actual, item.content) {
				return fmt.Errorf("%s is stale; run coverage-report -scope %s -checks <collection>", item.path, report.SelectedScope)
			}
		} else if _, err := safeio.AtomicWrite(root, item.path, item.content, artifactLimit); err != nil {
			return err
		}
	}
	return nil
}
