// Package taxonomy validates the finite additive v3-to-v4 transition without
// reinterpreting the reviewed v3 documents or their original decisions.
package taxonomy

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	Decision      = "docs/adr/0031-foundation-scope-and-taxonomy-v4.md"
	MigrationPath = "curriculum/migrations/v3-to-v4.yaml"
	V3Path        = "curriculum/migrations/snapshots/skills-v3.yaml"
	V3SHA256      = "8fc463c93634564e7972117ea7f7c1f90d4e5cb82659f2f631698ef68ba4ab64"
	V3Commit      = "0bc11985e9f67ab8677c3a7726f3239118e76442"
)

type Baseline struct {
	SourceCommit    string `yaml:"source_commit"`
	TaxonomyPath    string `yaml:"taxonomy_path"`
	TaxonomySHA256  string `yaml:"taxonomy_sha256"`
	SkillCount      int    `yaml:"skill_count"`
	ActiveCount     int    `yaml:"active_count"`
	DeprecatedCount int    `yaml:"deprecated_count"`
	SkillIDsSHA256  string `yaml:"skill_ids_sha256"`
}

type Entry struct {
	ID             string `yaml:"id"`
	Status         string `yaml:"status,omitempty"`
	ContractSHA256 string `yaml:"contract_sha256"`
}

type ReviewedDocument struct {
	Path         string   `yaml:"path"`
	SnapshotPath string   `yaml:"snapshot_path"`
	SHA256       string   `yaml:"sha256"`
	SkillIDs     []string `yaml:"skill_ids"`
}

type Transition struct {
	SchemaVersion       int                `yaml:"schema_version"`
	FromTaxonomyVersion int                `yaml:"from_taxonomy_version"`
	ToTaxonomyVersion   int                `yaml:"to_taxonomy_version"`
	Decision            string             `yaml:"decision"`
	Baseline            Baseline           `yaml:"baseline"`
	Entries             []Entry            `yaml:"entries"`
	Additions           []Entry            `yaml:"additions"`
	ReviewedDocuments   []ReviewedDocument `yaml:"reviewed_documents"`
}

type document struct {
	Version int       `yaml:"version"`
	Track   string    `yaml:"track"`
	Target  yaml.Node `yaml:"target_ready_tasks"`
	Policy  yaml.Node `yaml:"coverage_policy"`
	Domains []struct {
		ID     string           `yaml:"id"`
		Title  string           `yaml:"title"`
		Skills []map[string]any `yaml:"skills"`
	} `yaml:"domains"`
}

var reviewedV3 = map[string]struct{ snapshot, sha string }{
	"curriculum/mappings/exercism-go.yaml":              {"curriculum/migrations/snapshots/exercism-map-v3.yaml", "28238a0d041da4d283ae9d0873ceedd1dcf4ac3b83a68d1ee006158ce839a9f8"},
	"curriculum/evidence/exercism-go.yaml":              {"curriculum/migrations/snapshots/exercism-evidence-v3.yaml", "98f93350520bf6319d8a348a8a016fac3262d57c227b8b82c7993308427487c2"},
	"catalog/reviews/exercism-go.yaml":                  {"curriculum/migrations/snapshots/exercism-reviews-v3.yaml", "05bf10d9bdbc282ca66c8b0c5bc9129deab1475be5a53bab3c303d51ec698f03"},
	"catalog/seed-plans/exercism-foundation-pilot.yaml": {"curriculum/migrations/snapshots/exercism-seeds-v3.yaml", "0cbbf900b5da2ba75460bd0f2d54c832f64e1debeaa2f0daadfe2025c134a2c5"},
}

// These are canonical teaching-contract hashes from the independently reviewed
// ADR0031 addition records, excluding only the taxonomy_version field.
var approvedAdditions = map[string]string{
	"expressions.bitwise-masks":  "74406b8f2afcfb663c21c75cecd7b84aceeb9f9430eb2a734efd1e7fa5d79b8c",
	"expressions.integer-shifts": "2d437a42847d3e2babc2288bd75ddddf01366f7a0470ee21c6f971b1ffe1923b",
	"functions.recursion":        "dec90f41de0b538fc44e9adc95f39139eb84ac197b5b48f531d9801004f8d5d6",
}

// ContractPaths returns the additional v4 freshness inputs. V3 collections keep
// their original byte contract; a v4 collection binds its complete transition.
func ContractPaths() []string {
	paths := []string{MigrationPath, V3Path, ScopePath, ScopeSchemaPath, "schemas/taxonomy-v4.schema.json", "schemas/taxonomy-v3-to-v4.schema.json", "schemas/exercism-concept-map.schema.json", "schemas/skill-evidence.schema.json", "schemas/candidate-review.schema.json", "schemas/seed-plan.schema.json", "curriculum/migrations/v2-to-v3.yaml", "schemas/taxonomy-migration.schema.json", "curriculum/migrations/snapshots/skills-v2.yaml", "curriculum/migrations/snapshots/exercism-map-v2.yaml", "curriculum/migrations/snapshots/exercism-evidence-v2.yaml"}
	for path, pin := range reviewedV3 {
		paths = append(paths, path, pin.snapshot)
	}
	sort.Strings(paths)
	return paths
}

// ReviewedPath selects a historical taxonomy only after verifying the exact
// finite compatibility record. It is not a general version alias.
func ReviewedPath(root string) (string, int, error) {
	_, current, err := readImmutableDocument(root, "curriculum/skills.yaml")
	if err != nil {
		return "", 0, err
	}
	if current.Version == 3 {
		return "curriculum/skills.yaml", 3, nil
	}
	if current.Version != 4 {
		return "", 0, fmt.Errorf("taxonomy: unsupported version %d", current.Version)
	}
	if _, err := ValidateTransition(root); err != nil {
		return "", 0, err
	}
	return V3Path, 3, nil
}

// ValidateTransition checks all old contracts, all additions, and every reviewed
// document against the actual immutable baseline, including deprecated stubs.
func ValidateTransition(root string) (Transition, error) {
	return validateTransition(root, true)
}

// ValidateLiveTransition checks the live teaching contracts without requiring
// internal source-review bytes in a learner-only checkout. Reviewed source
// consumers must use ValidateTransition or ReviewedPath for the full check.
func ValidateLiveTransition(root string) (Transition, error) {
	return validateTransition(root, false)
}

func validateTransition(root string, requireReviewedBytes bool) (Transition, error) {
	var t Transition
	if _, err := readStrict(root, MigrationPath, &t); err != nil {
		return t, err
	}
	if t.SchemaVersion != 1 || t.FromTaxonomyVersion != 3 || t.ToTaxonomyVersion != 4 || t.Decision != Decision {
		return t, fmt.Errorf("taxonomy migration: expected finite v3-to-v4 ADR0031 transition")
	}
	b := t.Baseline
	if b.SourceCommit != V3Commit || b.TaxonomyPath != V3Path || b.TaxonomySHA256 != V3SHA256 || b.SkillCount != 177 || b.ActiveCount != 142 || b.DeprecatedCount != 35 {
		return t, fmt.Errorf("taxonomy migration: immutable v3 baseline identity/counts mismatch")
	}
	data, old, err := readImmutableDocument(root, V3Path)
	if err != nil {
		return t, err
	}
	if hash(data) != V3SHA256 || old.Version != 3 {
		return t, fmt.Errorf("taxonomy migration: immutable v3 taxonomy hash/version mismatch")
	}
	_, live, err := readImmutableDocument(root, "curriculum/skills.yaml")
	if err != nil {
		return t, err
	}
	if live.Version != 4 || live.Track != old.Track || !nodeEqual(live.Target, old.Target) {
		return t, fmt.Errorf("taxonomy migration: v4 version/track/task target mismatch")
	}
	if !policyEqual(live.Policy, old.Policy) {
		return t, fmt.Errorf("taxonomy migration: coverage policy mismatch")
	}
	before, err := skills(old)
	if err != nil {
		return t, err
	}
	after, err := skills(live)
	if err != nil {
		return t, err
	}
	ids := sortedKeys(before)
	if len(ids) != 177 || b.SkillIDsSHA256 != hash([]byte(strings.Join(ids, "\n")+"\n")) || len(t.Entries) != 177 || len(after) != 180 {
		return t, fmt.Errorf("taxonomy migration: complete177-old/180-live inventory required")
	}
	active, deprecated := 0, 0
	for i, id := range ids {
		skill, current := before[id], after[id]
		status, _ := skill["status"].(string)
		if status == "active" {
			active++
		} else if status == "deprecated" {
			deprecated++
		} else {
			return t, fmt.Errorf("taxonomy migration: invalid v3 status for %s", id)
		}
		if current == nil || skill["taxonomy_version"] != 3 || current["taxonomy_version"] != 4 || contractHash(skill) != contractHash(current) {
			return t, fmt.Errorf("taxonomy migration: old contract changed or missing: %s", id)
		}
		entry := t.Entries[i]
		if entry.ID != id || entry.Status != status || entry.ContractSHA256 != contractHash(skill) {
			return t, fmt.Errorf("taxonomy migration: incorrect preserved entry for %s", id)
		}
	}
	if active != 142 || deprecated != 35 {
		return t, fmt.Errorf("taxonomy migration: active/deprecated baseline counts mismatch")
	}
	expected := []string{"expressions.bitwise-masks", "expressions.integer-shifts", "functions.recursion"}
	if len(t.Additions) != len(expected) {
		return t, fmt.Errorf("taxonomy migration: exactly3 approved additions required")
	}
	for i, id := range expected {
		skill, entry := after[id], t.Additions[i]
		if before[id] != nil || skill == nil || skill["status"] != "active" || skill["assignable"] != true || skill["taxonomy_version"] != 4 || entry.ID != id || entry.Status != "" || entry.ContractSHA256 != contractHash(skill) || entry.ContractSHA256 != approvedAdditions[id] {
			return t, fmt.Errorf("taxonomy migration: invalid approved addition %s", id)
		}
	}
	if _, err := readCoverageScopes(root, after); err != nil {
		return t, err
	}
	if len(t.ReviewedDocuments) != len(reviewedV3) {
		return t, fmt.Errorf("taxonomy migration: exactly4 reviewed compatibility documents required")
	}
	last := ""
	for _, record := range t.ReviewedDocuments {
		pin, ok := reviewedV3[record.Path]
		if !ok || record.Path <= last || record.SnapshotPath != pin.snapshot || record.SHA256 != pin.sha {
			return t, fmt.Errorf("taxonomy migration: invalid reviewed document identity %q", record.Path)
		}
		last = record.Path
		if !requireReviewedBytes {
			continue
		}
		var snapshot, current yaml.Node
		saved, err := readStrict(root, record.SnapshotPath, &snapshot)
		if err != nil {
			return t, err
		}
		actual, err := readStrict(root, record.Path, &current)
		if err != nil {
			return t, err
		}
		if hash(saved) != pin.sha || !bytes.Equal(saved, actual) {
			return t, fmt.Errorf("taxonomy migration: reviewed v3 bytes changed: %s", record.Path)
		}
		set := map[string]bool{}
		collectReferences(&snapshot, set)
		if !reflect.DeepEqual(record.SkillIDs, sortedKeys(set)) {
			return t, fmt.Errorf("taxonomy migration: reviewed skill inventory mismatch: %s", record.Path)
		}
		for id := range set {
			if before[id] == nil || contractHash(before[id]) != contractHash(after[id]) {
				return t, fmt.Errorf("taxonomy migration: incompatible reviewed skill %s", id)
			}
		}
	}
	return t, nil
}

func skills(d document) (map[string]map[string]any, error) {
	result := map[string]map[string]any{}
	for _, domain := range d.Domains {
		for _, s := range domain.Skills {
			id, ok := s["id"].(string)
			if !ok || id == "" || result[id] != nil {
				return nil, fmt.Errorf("taxonomy: empty/duplicate skill id %q", id)
			}
			result[id] = s
		}
	}
	return result, nil
}

func contractHash(skill map[string]any) string {
	m := map[string]any{}
	for key, value := range skill {
		if key != "taxonomy_version" {
			m[key] = value
		}
	}
	b, err := json.Marshal(m)
	if err != nil {
		return "invalid-contract"
	}
	return hash(b)
}

func nodeEqual(a, b yaml.Node) bool {
	var x, y any
	if a.Decode(&x) != nil || b.Decode(&y) != nil {
		return false
	}
	return reflect.DeepEqual(x, y)
}

func policyEqual(a, b yaml.Node) bool {
	var x, y map[string]any
	if a.Decode(&x) != nil || b.Decode(&y) != nil {
		return false
	}
	delete(x, "note")
	delete(y, "note")
	return reflect.DeepEqual(x, y)
}
func hash(b []byte) string { value := sha256.Sum256(b); return hex.EncodeToString(value[:]) }
func sortedKeys[V any](m map[string]V) []string {
	ids := make([]string, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func collectReferences(n *yaml.Node, set map[string]bool) {
	keys := map[string]bool{"skill_id": true, "skill_ids": true, "proposed_primary_skills": true, "proposed_secondary_skills": true, "prerequisites": true, "target_skills": true}
	if n.Kind == yaml.MappingNode {
		for i := 0; i < len(n.Content); i += 2 {
			key, value := n.Content[i].Value, n.Content[i+1]
			if keys[key] {
				if value.Kind == yaml.ScalarNode {
					set[value.Value] = true
				} else {
					for _, v := range value.Content {
						set[v.Value] = true
					}
				}
			}
			collectReferences(value, set)
		}
	} else {
		for _, c := range n.Content {
			collectReferences(c, set)
		}
	}
}
