package curriculum

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/sources"
	"gopkg.in/yaml.v3"
)

var reviewDatePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var evidenceIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:[./-][a-z0-9]+)*$`)

type SkillContract struct {
	Status     string
	Assignable bool
}

type ConceptInventory struct {
	SchemaVersion int      `yaml:"schema_version"`
	SourceID      string   `yaml:"source_id"`
	SourceCommit  string   `yaml:"source_commit"`
	Concepts      []string `yaml:"concepts"`
}

type ConceptMap struct {
	SchemaVersion   int              `yaml:"schema_version"`
	SourceID        string           `yaml:"source_id"`
	SourceCommit    string           `yaml:"source_commit"`
	TaxonomyVersion int              `yaml:"taxonomy_version"`
	Review          MapReview        `yaml:"review"`
	Concepts        []ConceptMapping `yaml:"concepts"`
}

type MapReview struct {
	Status     string `yaml:"status"`
	ReviewedBy string `yaml:"reviewed_by"`
	ReviewedAt string `yaml:"reviewed_at"`
	Decision   string `yaml:"decision"`
}

type ConceptMapping struct {
	Upstream  string   `yaml:"upstream"`
	Status    string   `yaml:"status"`
	SkillIDs  []string `yaml:"skill_ids"`
	Rationale string   `yaml:"rationale"`
}

type EvidenceFile struct {
	SchemaVersion   int              `yaml:"schema_version"`
	SourceID        string           `yaml:"source_id"`
	SourceCommit    string           `yaml:"source_commit"`
	TaxonomyVersion int              `yaml:"taxonomy_version"`
	MigrationReview MigrationReview  `yaml:"migration_review"`
	Reviews         []EvidenceReview `yaml:"reviews"`
}

type MigrationReview struct {
	Status                 string `yaml:"status"`
	ReviewedBy             string `yaml:"reviewed_by"`
	ReviewedAt             string `yaml:"reviewed_at"`
	Decision               string `yaml:"decision"`
	PreviousEvidenceSHA256 string `yaml:"previous_evidence_sha256"`
}

type EvidenceReview struct {
	ID               string `yaml:"id"`
	ExerciseKind     string `yaml:"exercise_kind"`
	Slug             string `yaml:"slug"`
	SourcePath       string `yaml:"source_path"`
	UpstreamConcept  string `yaml:"upstream_concept"`
	SkillID          string `yaml:"skill_id"`
	Basis            string `yaml:"basis"`
	Strength         string `yaml:"strength"`
	Decision         string `yaml:"decision"`
	CoverageEligible *bool  `yaml:"coverage_eligible"`
	ReviewedBy       string `yaml:"reviewed_by"`
	ReviewedAt       string `yaml:"reviewed_at"`
	Rationale        string `yaml:"rationale"`
	ArtifactPath     string `yaml:"artifact_path,omitempty"`
	ArtifactSHA256   string `yaml:"artifact_sha256,omitempty"`
}

func LoadConceptInventory(root string) (ConceptInventory, error) {
	var value ConceptInventory
	if err := decodeStrict(root, "curriculum/mappings/exercism-go-concepts.yaml", &value); err != nil {
		return ConceptInventory{}, err
	}
	return value, nil
}

func LoadConceptMap(root string) (ConceptMap, error) {
	var value ConceptMap
	if err := decodeStrict(root, "curriculum/mappings/exercism-go.yaml", &value); err != nil {
		return ConceptMap{}, err
	}
	return value, nil
}

func LoadEvidence(root string) (EvidenceFile, error) {
	var value EvidenceFile
	if err := decodeStrict(root, "curriculum/evidence/exercism-go.yaml", &value); err != nil {
		return EvidenceFile{}, err
	}
	return value, nil
}

func Validate(mapping ConceptMap, evidence EvidenceFile, source sources.Source, taxonomyVersion int, skillContracts map[string]SkillContract, upstreamConcepts []string, candidates []catalog.Candidate) error {
	var problems []error
	if mapping.SchemaVersion != 1 || mapping.SourceID != source.ID || mapping.SourceCommit != source.Commit || mapping.TaxonomyVersion != taxonomyVersion {
		problems = append(problems, errors.New("concept map version or source provenance does not match taxonomy/registry"))
	}
	if mapping.Review.Status != "reviewed" || mapping.Review.ReviewedBy == "" || !validDate(mapping.Review.ReviewedAt) || mapping.Review.Decision != "docs/adr/0003-taxonomy-v3.md" {
		problems = append(problems, errors.New("concept map requires reviewed status, reviewer, date, and decision"))
	}
	byConcept := map[string]ConceptMapping{}
	lastConcept := ""
	for _, item := range mapping.Concepts {
		if item.Upstream == "" || item.Rationale == "" {
			problems = append(problems, errors.New("concept mapping requires upstream id and rationale"))
		}
		if _, exists := byConcept[item.Upstream]; exists {
			problems = append(problems, fmt.Errorf("duplicate concept mapping %q", item.Upstream))
		}
		if lastConcept != "" && item.Upstream <= lastConcept {
			problems = append(problems, fmt.Errorf("concept mappings are not strictly sorted at %q", item.Upstream))
		}
		lastConcept = item.Upstream
		byConcept[item.Upstream] = item
		lastSkillID := ""
		for _, skillID := range item.SkillIDs {
			contract, ok := skillContracts[skillID]
			if !ok {
				problems = append(problems, fmt.Errorf("concept %q references unknown skill %q", item.Upstream, skillID))
			} else if contract.Status != "active" || !contract.Assignable {
				problems = append(problems, fmt.Errorf("concept %q references non-assignable skill %q", item.Upstream, skillID))
			}
			if lastSkillID != "" && skillID <= lastSkillID {
				problems = append(problems, fmt.Errorf("concept %q skill_ids are not strictly sorted", item.Upstream))
			}
			lastSkillID = skillID
		}
		switch item.Status {
		case "mapped":
			if len(item.SkillIDs) != 1 {
				problems = append(problems, fmt.Errorf("mapped concept %q must map to one skill", item.Upstream))
			}
		case "ambiguous":
			if len(item.SkillIDs) < 2 {
				problems = append(problems, fmt.Errorf("ambiguous concept %q must list at least two candidate skills", item.Upstream))
			}
		case "taxonomy-gap", "scope-deferred":
			if len(item.SkillIDs) != 0 {
				problems = append(problems, fmt.Errorf("concept %q with status %q cannot assign skills", item.Upstream, item.Status))
			}
		default:
			problems = append(problems, fmt.Errorf("concept %q has invalid status %q", item.Upstream, item.Status))
		}
	}
	if upstreamConcepts != nil {
		expected := uniqueSorted(upstreamConcepts)
		actual := make([]string, 0, len(byConcept))
		for concept := range byConcept {
			actual = append(actual, concept)
		}
		sort.Strings(actual)
		if !equalStrings(actual, expected) {
			problems = append(problems, fmt.Errorf("concept map inventory differs from pinned root config: mapped=%v upstream=%v", actual, expected))
		}
	}

	if evidence.SchemaVersion != 1 || evidence.SourceID != source.ID || evidence.SourceCommit != source.Commit || evidence.TaxonomyVersion != taxonomyVersion {
		problems = append(problems, errors.New("evidence version or source provenance does not match taxonomy/registry"))
	}
	if evidence.MigrationReview.Status != "reviewed" || evidence.MigrationReview.ReviewedBy == "" || !validDate(evidence.MigrationReview.ReviewedAt) || evidence.MigrationReview.Decision != "docs/adr/0003-taxonomy-v3.md" || !regexp.MustCompile(`^[0-9a-f]{64}$`).MatchString(evidence.MigrationReview.PreviousEvidenceSHA256) {
		problems = append(problems, errors.New("evidence requires an explicit reviewed v2-to-v3 migration decision"))
	}
	candidateByPath := map[string]catalog.Candidate{}
	for _, candidate := range candidates {
		if _, exists := candidateByPath[candidate.SourcePath]; exists {
			problems = append(problems, fmt.Errorf("duplicate candidate source path %q", candidate.SourcePath))
		}
		candidateByPath[candidate.SourcePath] = candidate
	}
	seenReview := map[string]bool{}
	seenNaturalKey := map[string]bool{}
	for _, review := range evidence.Reviews {
		if !evidenceIDPattern.MatchString(review.ID) || seenReview[review.ID] {
			problems = append(problems, fmt.Errorf("duplicate or empty evidence review id %q", review.ID))
		}
		seenReview[review.ID] = true
		naturalKey := strings.Join([]string{review.SourcePath, review.UpstreamConcept, review.SkillID, review.Basis}, "\x00")
		if seenNaturalKey[naturalKey] {
			problems = append(problems, fmt.Errorf("duplicate evidence natural key for %q", review.ID))
		}
		seenNaturalKey[naturalKey] = true
		if review.SourcePath != path.Join("exercises", review.ExerciseKind, review.Slug) {
			problems = append(problems, fmt.Errorf("evidence %q source path does not match kind/slug", review.ID))
		}
		if review.CoverageEligible == nil || *review.CoverageEligible {
			problems = append(problems, fmt.Errorf("evidence %q must explicitly set coverage_eligible to false", review.ID))
		}
		contract, ok := skillContracts[review.SkillID]
		if !ok {
			problems = append(problems, fmt.Errorf("evidence %q references unknown skill %q", review.ID, review.SkillID))
		} else if contract.Status != "active" || !contract.Assignable {
			problems = append(problems, fmt.Errorf("evidence %q references non-assignable skill %q", review.ID, review.SkillID))
		}
		mappingItem, ok := byConcept[review.UpstreamConcept]
		if !ok {
			problems = append(problems, fmt.Errorf("evidence %q references unmapped concept %q", review.ID, review.UpstreamConcept))
		} else if review.Decision == "accepted" && review.Basis == "upstream-concept-label" && (mappingItem.Status != "mapped" || !contains(mappingItem.SkillIDs, review.SkillID)) {
			problems = append(problems, fmt.Errorf("evidence %q cannot accept a non-mapped label", review.ID))
		}
		if review.Basis != "upstream-concept-label" && review.Basis != "manual-artifact-review" {
			problems = append(problems, fmt.Errorf("evidence %q has invalid basis %q", review.ID, review.Basis))
		}
		if review.Basis == "manual-artifact-review" && (review.ArtifactPath == "" || !regexp.MustCompile(`^[0-9a-f]{64}$`).MatchString(review.ArtifactSHA256)) {
			problems = append(problems, fmt.Errorf("evidence %q manual review requires artifact_path and artifact_sha256", review.ID))
		}
		if review.Strength != "direct" && review.Strength != "supporting" {
			problems = append(problems, fmt.Errorf("evidence %q has invalid strength %q", review.ID, review.Strength))
		}
		if review.Decision != "accepted" && review.Decision != "rejected" && review.Decision != "needs-review" {
			problems = append(problems, fmt.Errorf("evidence %q has invalid decision %q", review.ID, review.Decision))
		}
		if review.ReviewedBy == "" || !validDate(review.ReviewedAt) || len(review.Rationale) < 10 {
			problems = append(problems, fmt.Errorf("evidence %q lacks review metadata", review.ID))
		}
		if candidates != nil {
			candidate, ok := candidateByPath[review.SourcePath]
			if !ok || candidate.Slug != review.Slug || candidate.ExerciseKind != review.ExerciseKind {
				problems = append(problems, fmt.Errorf("evidence %q does not resolve to a candidate", review.ID))
			} else if review.Basis == "upstream-concept-label" && !contains(candidate.UpstreamConcepts, review.UpstreamConcept) {
				problems = append(problems, fmt.Errorf("evidence %q concept is not declared by candidate", review.ID))
			}
		}
	}
	return errors.Join(problems...)
}

func ValidateInventory(inventory ConceptInventory, source sources.Source) error {
	var problems []error
	if inventory.SchemaVersion != 1 || inventory.SourceID != source.ID || inventory.SourceCommit != source.Commit {
		problems = append(problems, errors.New("concept inventory source provenance does not match registry"))
	}
	if len(inventory.Concepts) != 38 {
		problems = append(problems, fmt.Errorf("concept inventory must contain 38 labels, got %d", len(inventory.Concepts)))
	}
	last := ""
	for _, concept := range inventory.Concepts {
		if concept == "" || (last != "" && concept <= last) {
			problems = append(problems, fmt.Errorf("concept inventory is empty, duplicated, or unsorted at %q", concept))
		}
		last = concept
	}
	return errors.Join(problems...)
}

func MappingByConcept(mapping ConceptMap) map[string]ConceptMapping {
	result := make(map[string]ConceptMapping, len(mapping.Concepts))
	for _, item := range mapping.Concepts {
		result[item.Upstream] = item
	}
	return result
}

func AcceptedEvidenceByConcept(evidence EvidenceFile) map[string][]EvidenceReview {
	result := map[string][]EvidenceReview{}
	for _, review := range evidence.Reviews {
		if review.Decision == "accepted" {
			result[review.UpstreamConcept] = append(result[review.UpstreamConcept], review)
		}
	}
	for concept := range result {
		sort.Slice(result[concept], func(i, j int) bool { return result[concept][i].ID < result[concept][j].ID })
	}
	return result
}

func decodeStrict(root, relative string, destination any) error {
	b, err := safeio.ReadFile(root, relative, 4<<20)
	if err != nil {
		return fmt.Errorf("read %s: %w", relative, err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(b))
	decoder.KnownFields(true)
	if err := decoder.Decode(destination); err != nil {
		return fmt.Errorf("decode %s: %w", relative, err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err == nil {
		return fmt.Errorf("decode %s: multiple YAML documents are not allowed", relative)
	} else if !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode %s trailer: %w", relative, err)
	}
	return nil
}

func uniqueSorted(values []string) []string {
	set := map[string]bool{}
	for _, value := range values {
		set[value] = true
	}
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func equalStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func validDate(value string) bool {
	if !reviewDatePattern.MatchString(value) {
		return false
	}
	parsed, err := time.Parse("2006-01-02", value)
	return err == nil && parsed.Format("2006-01-02") == value
}
