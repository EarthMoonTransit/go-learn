package triage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
)

const testCommit = "0123456789abcdef0123456789abcdef01234567"

func TestValidateReviewInvariants(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Bundle)
		want   string
	}{
		{
			name: "candidate reviewed exactly once",
			mutate: func(bundle *Bundle) {
				bundle.Reviews.Reviews = bundle.Reviews.Reviews[1:]
			},
			want: "is not reviewed",
		},
		{
			name: "known skills only",
			mutate: func(bundle *Bundle) {
				bundle.Reviews.Reviews[0].ProposedPrimarySkills = []string{"missing.skill"}
			},
			want: "unknown or non-assignable primary skill",
		},
		{
			name: "valid statuses only",
			mutate: func(bundle *Bundle) {
				bundle.Reviews.Reviews[0].ReviewStatus = "looks-good"
			},
			want: "invalid status",
		},
		{
			name: "cluster references are bidirectional",
			mutate: func(bundle *Bundle) {
				bundle.Reviews.Reviews[0].DuplicateCluster = "missing-cluster"
			},
			want: "membership is",
		},
		{
			name: "seed references a reviewed candidate",
			mutate: func(bundle *Bundle) {
				bundle.Seeds.SelectedSeeds[len(bundle.Seeds.SelectedSeeds)-1].CandidateID = "practice/z99"
			},
			want: "unknown reviewed candidate",
		},
		{
			name: "rejected seed is forbidden",
			mutate: func(bundle *Bundle) {
				bundle.Reviews.Reviews[0].ReviewStatus = "reject-low-value"
				bundle.Reviews.Reviews[0].PossibleTaskForms = nil
			},
			want: "ineligible candidate",
		},
		{
			name: "provenance must match catalog",
			mutate: func(bundle *Bundle) {
				bundle.Reviews.Reviews[0].ProvenanceStatus = "generated-data-unresolved"
			},
			want: "invalid provenance status",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bundle, candidates, skills, source := validFixture()
			test.mutate(&bundle)
			err := Validate(bundle, candidates, skills, 3, source)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, test.want)
			}
		})
	}
}

func TestValidateLicenseAndProvenancePolicy(t *testing.T) {
	t.Run("source must permit adaptation", func(t *testing.T) {
		bundle, candidates, skills, source := validFixture()
		source.UsageMode = "reference-only"
		err := Validate(bundle, candidates, skills, 3, source)
		if err == nil || !strings.Contains(err.Error(), "not approved for adaptation") {
			t.Fatalf("Validate() error = %v", err)
		}
	})

	t.Run("unresolved generated data cannot be accepted", func(t *testing.T) {
		bundle, candidates, skills, source := validFixture()
		candidates[0].GeneratedTestDataPaths = []string{"exercises/practice/c00/.meta/tests.toml"}
		bundle.Reviews.Reviews[0].ProvenanceStatus = "generated-data-unresolved"
		err := Validate(bundle, candidates, skills, 3, source)
		if err == nil || !strings.Contains(err.Error(), "cannot be accepted while generated-data provenance is unresolved") {
			t.Fatalf("Validate() error = %v", err)
		}
	})
}

func TestGenerateReportsIsDeterministic(t *testing.T) {
	bundle, _, skills, _ := validFixture()
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "reports"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := GenerateReports(root, bundle, skills); err != nil {
		t.Fatal(err)
	}
	first := readReports(t, root)
	if err := GenerateReports(root, bundle, skills); err != nil {
		t.Fatal(err)
	}
	second := readReports(t, root)
	if fmt.Sprint(first) != fmt.Sprint(second) {
		t.Fatalf("report output changed between identical runs")
	}
}

func readReports(t *testing.T, root string) map[string]string {
	t.Helper()
	result := map[string]string{}
	paths, err := filepath.Glob(filepath.Join(root, "reports", "*.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range paths {
		contents, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		result[filepath.Base(path)] = string(contents)
	}
	return result
}

func validFixture() (Bundle, []catalog.Candidate, map[string]SkillContract, sources.Source) {
	source := sources.Source{ID: "exercism-go", Commit: testCommit, UsageMode: "adapt", License: sources.SourceLicense{Expression: "MIT"}}
	skill := SkillContract{ID: "types.zero-values", Domain: "types", Status: "active", Assignable: true}
	skills := map[string]SkillContract{skill.ID: skill}
	bundle := Bundle{
		Reviews:  ReviewFile{SchemaVersion: 1, SourceID: source.ID, SourceCommit: source.Commit, TaxonomyVersion: 3, ReviewedBy: "test", ReviewedAt: "2026-09-04"},
		Clusters: ClusterFile{SchemaVersion: 1, SourceID: source.ID, SourceCommit: source.Commit, ReviewedBy: "test", ReviewedAt: "2026-09-04"},
		Seeds:    SeedPlan{SchemaVersion: 1, SourceID: source.ID, SourceCommit: source.Commit, TaxonomyVersion: 3, ReviewedBy: "test", ReviewedAt: "2026-09-04"},
	}
	var candidates []catalog.Candidate
	for index := 0; index < 30; index++ {
		id := fmt.Sprintf("practice/c%02d", index)
		path := "exercises/" + id
		candidate := catalog.Candidate{SourceID: source.ID, SourceCommit: source.Commit, SourcePath: path, Slug: strings.TrimPrefix(id, "practice/"), ExerciseKind: "practice", License: "MIT", UsageMode: "adapt"}
		candidates = append(candidates, candidate)
		bundle.Reviews.Reviews = append(bundle.Reviews.Reviews, CandidateReview{
			CandidateID: id, SourcePath: path, SourceCommit: source.Commit, ExerciseKind: "practice",
			ProposedPrimarySkills: []string{skill.ID}, ConceptualMechanisms: []string{"observable zero-value behavior"},
			LikelyMisconceptions: []string{"assuming uninitialized variables contain arbitrary data"}, PossibleTaskForms: []string{"implementation"},
			ProvenanceStatus: "source-exact", ReviewStatus: "accept-as-seed",
			Scores: Scores{3, 3, 3, 3, 3, 3, 3, 3, 3}, DecisionReason: "Focused seed for fixture validation.", ReviewerRationale: "The fixture deliberately isolates one known skill.",
		})
		bundle.Seeds.SelectedSeeds = append(bundle.Seeds.SelectedSeeds, SeedSelection{
			CandidateID: id, TargetSkills: []string{skill.ID}, IntendedTaskForms: []string{"implementation"}, PilotRole: "practice",
			SelectionRationale: "The fixture keeps every selected reference reviewable.",
		})
	}
	return bundle, candidates, skills, source
}
