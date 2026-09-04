package triage

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
)

var (
	candidateIDPattern = regexp.MustCompile(`^(concept|practice)/[a-z0-9]+(?:-[a-z0-9]+)*$`)
	clusterIDPattern   = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
)

var reviewStatuses = stringSet(
	"accept-as-seed", "accept-after-rewrite", "split-into-multiple-seeds", "merge-with-other-candidate",
	"idea-only", "deferred", "reject-duplicate", "reject-low-value", "reject-bad-tests",
	"reject-too-algorithmic", "reject-too-broad", "reject-provenance", "reject-license",
)

var taskForms = stringSet(
	"implementation", "debugging", "write-tests", "repair-tests", "refactoring", "API-design",
	"code-review", "integration", "fuzzing", "benchmark", "concurrency-debugging",
)

var provenanceStatuses = stringSet("source-exact", "multi-hop-exact", "generated-data-unresolved", "blocked")
var clusterSignalTypes = stringSet("file-hash", "test-hash", "concept-set", "public-api", "ast-shape", "semantic")
var clusterDecisions = stringSet("keep-distinct", "prefer-one", "merge-as-seed", "reject-redundant")
var pilotRoles = stringSet("introduction", "practice", "transfer", "diagnostic")
var gapOrigins = stringSet("original-tasks", "permissive-sources", "debugging-tasks", "production-derived-tasks")

func Validate(bundle Bundle, candidates []catalog.Candidate, skills map[string]SkillContract, taxonomyVersion int, source sources.Source) error {
	var problems []error
	if bundle.Reviews.SchemaVersion != 1 || bundle.Reviews.SourceID != source.ID || bundle.Reviews.SourceCommit != source.Commit || bundle.Reviews.TaxonomyVersion != taxonomyVersion || bundle.Reviews.ReviewedBy == "" || !validDate(bundle.Reviews.ReviewedAt) {
		problems = append(problems, errors.New("candidate reviews: metadata does not match source/taxonomy or lacks review identity"))
	}
	if bundle.Clusters.SchemaVersion != 1 || bundle.Clusters.SourceID != source.ID || bundle.Clusters.SourceCommit != source.Commit || bundle.Clusters.ReviewedBy == "" || !validDate(bundle.Clusters.ReviewedAt) {
		problems = append(problems, errors.New("candidate clusters: metadata does not match source or lacks review identity"))
	}
	if bundle.Seeds.SchemaVersion != 1 || bundle.Seeds.SourceID != source.ID || bundle.Seeds.SourceCommit != source.Commit || bundle.Seeds.TaxonomyVersion != taxonomyVersion || bundle.Seeds.ReviewedBy == "" || !validDate(bundle.Seeds.ReviewedAt) {
		problems = append(problems, errors.New("seed plan: metadata does not match source/taxonomy or lacks review identity"))
	}
	if source.UsageMode != "adapt" || source.License.Expression == "" {
		problems = append(problems, fmt.Errorf("candidate reviews: source %q is not approved for adaptation", source.ID))
	}

	candidateByID := map[string]catalog.Candidate{}
	for _, candidate := range candidates {
		id := CandidateID(candidate)
		if _, exists := candidateByID[id]; exists {
			problems = append(problems, fmt.Errorf("candidate reviews: duplicate catalog candidate %q", id))
		}
		candidateByID[id] = candidate
	}

	reviewByID := map[string]CandidateReview{}
	lastReviewID := ""
	for _, review := range bundle.Reviews.Reviews {
		if !candidateIDPattern.MatchString(review.CandidateID) || review.CandidateID <= lastReviewID {
			problems = append(problems, fmt.Errorf("candidate reviews: ids must be valid, unique, and strictly sorted at %q", review.CandidateID))
		}
		lastReviewID = review.CandidateID
		if _, exists := reviewByID[review.CandidateID]; exists {
			problems = append(problems, fmt.Errorf("candidate reviews: duplicate review %q", review.CandidateID))
		}
		reviewByID[review.CandidateID] = review
		candidate, exists := candidateByID[review.CandidateID]
		if !exists {
			problems = append(problems, fmt.Errorf("candidate reviews: unknown candidate %q", review.CandidateID))
			continue
		}
		if review.SourcePath != candidate.SourcePath || review.SourceCommit != candidate.SourceCommit || review.ExerciseKind != candidate.ExerciseKind {
			problems = append(problems, fmt.Errorf("candidate reviews: %q provenance metadata differs from catalog", review.CandidateID))
		}
		if candidate.SourceID != source.ID || candidate.SourceCommit != source.Commit || candidate.License != source.License.Expression || candidate.UsageMode != source.UsageMode {
			problems = append(problems, fmt.Errorf("candidate reviews: %q catalog provenance differs from registry", review.CandidateID))
		}
		if !reviewStatuses[review.ReviewStatus] {
			problems = append(problems, fmt.Errorf("candidate reviews: %q has invalid status %q", review.CandidateID, review.ReviewStatus))
		}
		if !provenanceStatuses[review.ProvenanceStatus] || review.ProvenanceStatus != expectedProvenance(candidate) {
			problems = append(problems, fmt.Errorf("candidate reviews: %q has invalid provenance status %q", review.CandidateID, review.ProvenanceStatus))
		}
		if review.ProvenanceStatus == "generated-data-unresolved" && isAcceptedStatus(review.ReviewStatus) {
			problems = append(problems, fmt.Errorf("candidate reviews: %q cannot be accepted while generated-data provenance is unresolved", review.CandidateID))
		}
		if review.ReviewStatus == "reject-provenance" && review.ProvenanceStatus != "generated-data-unresolved" && review.ProvenanceStatus != "blocked" {
			problems = append(problems, fmt.Errorf("candidate reviews: %q rejects provenance that is exact", review.CandidateID))
		}
		if review.ReviewStatus == "reject-license" && candidate.License == source.License.Expression && source.UsageMode == "adapt" {
			problems = append(problems, fmt.Errorf("candidate reviews: %q rejects the registry-approved license without candidate-specific evidence", review.CandidateID))
		}
		for label, ids := range map[string][]string{
			"primary": review.ProposedPrimarySkills, "secondary": review.ProposedSecondarySkills, "prerequisite": review.Prerequisites,
		} {
			if !strictlySorted(ids) {
				problems = append(problems, fmt.Errorf("candidate reviews: %q %s skills must be sorted and unique", review.CandidateID, label))
			}
			for _, id := range ids {
				contract, ok := skills[id]
				if !ok || contract.Status != "active" || !contract.Assignable {
					problems = append(problems, fmt.Errorf("candidate reviews: %q references unknown or non-assignable %s skill %q", review.CandidateID, label, id))
				}
			}
		}
		if overlap(review.ProposedPrimarySkills, review.ProposedSecondarySkills) || overlap(review.ProposedPrimarySkills, review.Prerequisites) || overlap(review.ProposedSecondarySkills, review.Prerequisites) {
			problems = append(problems, fmt.Errorf("candidate reviews: %q skill roles overlap", review.CandidateID))
		}
		if len(review.ConceptualMechanisms) == 0 || len(review.LikelyMisconceptions) == 0 || len(review.DecisionReason) < 15 || len(review.ReviewerRationale) < 20 {
			problems = append(problems, fmt.Errorf("candidate reviews: %q lacks a reviewed pedagogical rationale", review.CandidateID))
		}
		if !strictlySorted(review.PossibleTaskForms) {
			problems = append(problems, fmt.Errorf("candidate reviews: %q task forms must be sorted and unique", review.CandidateID))
		}
		for _, form := range review.PossibleTaskForms {
			if !taskForms[form] {
				problems = append(problems, fmt.Errorf("candidate reviews: %q has invalid task form %q", review.CandidateID, form))
			}
		}
		if isRejectedStatus(review.ReviewStatus) && len(review.PossibleTaskForms) != 0 {
			problems = append(problems, fmt.Errorf("candidate reviews: rejected candidate %q cannot propose task forms", review.CandidateID))
		}
		for label, score := range review.Scores.values() {
			if score < 0 || score > 5 {
				problems = append(problems, fmt.Errorf("candidate reviews: %q score %s=%d is outside 0..5", review.CandidateID, label, score))
			}
		}
	}
	for id := range candidateByID {
		if _, ok := reviewByID[id]; !ok {
			problems = append(problems, fmt.Errorf("candidate reviews: candidate %q is not reviewed", id))
		}
	}
	if len(reviewByID) != len(candidateByID) {
		problems = append(problems, fmt.Errorf("candidate reviews: reviewed=%d catalog=%d", len(reviewByID), len(candidateByID)))
	}

	clusterByID := map[string]Cluster{}
	memberCluster := map[string]string{}
	lastClusterID := ""
	for _, cluster := range bundle.Clusters.Clusters {
		if !clusterIDPattern.MatchString(cluster.ID) || cluster.ID <= lastClusterID {
			problems = append(problems, fmt.Errorf("candidate clusters: ids must be valid, unique, and strictly sorted at %q", cluster.ID))
		}
		lastClusterID = cluster.ID
		clusterByID[cluster.ID] = cluster
		if len(cluster.Members) < 2 || !strictlySorted(cluster.Members) || len(cluster.SignalTypes) == 0 || !strictlySorted(cluster.SignalTypes) || !clusterDecisions[cluster.Decision] || len(cluster.ReviewerRationale) < 20 {
			problems = append(problems, fmt.Errorf("candidate clusters: %q has an incomplete reviewed contract", cluster.ID))
		}
		for _, signal := range cluster.SignalTypes {
			if !clusterSignalTypes[signal] {
				problems = append(problems, fmt.Errorf("candidate clusters: %q has invalid signal %q", cluster.ID, signal))
			}
		}
		if !strictlySorted(cluster.PreferredCandidates) {
			problems = append(problems, fmt.Errorf("candidate clusters: %q preferred candidates must be sorted and unique", cluster.ID))
		}
		for _, id := range cluster.Members {
			if _, ok := reviewByID[id]; !ok {
				problems = append(problems, fmt.Errorf("candidate clusters: %q references unknown member %q", cluster.ID, id))
			}
			if previous, duplicate := memberCluster[id]; duplicate {
				problems = append(problems, fmt.Errorf("candidate clusters: member %q appears in %q and %q", id, previous, cluster.ID))
			}
			memberCluster[id] = cluster.ID
		}
		for _, id := range cluster.PreferredCandidates {
			if !contains(cluster.Members, id) {
				problems = append(problems, fmt.Errorf("candidate clusters: %q prefers non-member %q", cluster.ID, id))
			}
		}
	}
	for id, review := range reviewByID {
		actual := memberCluster[id]
		if review.DuplicateCluster != actual {
			problems = append(problems, fmt.Errorf("candidate clusters: review %q declares %q, membership is %q", id, review.DuplicateCluster, actual))
		}
		if actual != "" {
			if _, ok := clusterByID[actual]; !ok {
				problems = append(problems, fmt.Errorf("candidate clusters: review %q references unknown cluster %q", id, actual))
			}
		}
	}

	coveredSkills := map[string]bool{}
	selectedCandidates := map[string]bool{}
	lastSeedID := ""
	for _, seed := range bundle.Seeds.SelectedSeeds {
		if seed.CandidateID <= lastSeedID || selectedCandidates[seed.CandidateID] {
			problems = append(problems, fmt.Errorf("seed plan: candidates must be unique and strictly sorted at %q", seed.CandidateID))
		}
		lastSeedID = seed.CandidateID
		selectedCandidates[seed.CandidateID] = true
		review, ok := reviewByID[seed.CandidateID]
		if !ok {
			problems = append(problems, fmt.Errorf("seed plan: unknown reviewed candidate %q", seed.CandidateID))
			continue
		}
		if isRejectedStatus(review.ReviewStatus) || review.ReviewStatus == "idea-only" || review.ReviewStatus == "deferred" || review.ProvenanceStatus == "generated-data-unresolved" || review.ProvenanceStatus == "blocked" {
			problems = append(problems, fmt.Errorf("seed plan: ineligible candidate %q with status %q", seed.CandidateID, review.ReviewStatus))
		}
		if len(seed.TargetSkills) == 0 || !strictlySorted(seed.TargetSkills) || len(seed.IntendedTaskForms) == 0 || !strictlySorted(seed.IntendedTaskForms) || !pilotRoles[seed.PilotRole] || len(seed.SelectionRationale) < 20 {
			problems = append(problems, fmt.Errorf("seed plan: candidate %q has an incomplete selection contract", seed.CandidateID))
		}
		allowedSkills := append(append([]string{}, review.ProposedPrimarySkills...), review.ProposedSecondarySkills...)
		for _, id := range seed.TargetSkills {
			if !contains(allowedSkills, id) {
				problems = append(problems, fmt.Errorf("seed plan: candidate %q targets unreviewed skill %q", seed.CandidateID, id))
			}
			coveredSkills[id] = true
		}
		for _, form := range seed.IntendedTaskForms {
			if !contains(review.PossibleTaskForms, form) {
				problems = append(problems, fmt.Errorf("seed plan: candidate %q uses unreviewed task form %q", seed.CandidateID, form))
			}
		}
	}
	if len(bundle.Seeds.SelectedSeeds) < 30 || len(bundle.Seeds.SelectedSeeds) > 50 {
		problems = append(problems, fmt.Errorf("seed plan: selected seed count %d is outside 30..50", len(bundle.Seeds.SelectedSeeds)))
	}

	gapSkills := map[string]bool{}
	lastGapKey := ""
	for _, gap := range bundle.Seeds.Gaps {
		key := gap.RequiredOrigin + "\x00" + strings.Join(gap.SkillIDs, "\x00")
		if key <= lastGapKey || len(gap.SkillIDs) == 0 || !strictlySorted(gap.SkillIDs) || !gapOrigins[gap.RequiredOrigin] || len(gap.Rationale) < 20 {
			problems = append(problems, fmt.Errorf("seed plan: invalid or unsorted gap group %q", gap.RequiredOrigin))
		}
		lastGapKey = key
		for _, id := range gap.SkillIDs {
			contract, ok := skills[id]
			if !ok || contract.Status != "active" || !contract.Assignable {
				problems = append(problems, fmt.Errorf("seed plan: gap references unknown or non-assignable skill %q", id))
			}
			if gapSkills[id] {
				problems = append(problems, fmt.Errorf("seed plan: skill %q appears in multiple gap groups", id))
			}
			if coveredSkills[id] {
				problems = append(problems, fmt.Errorf("seed plan: covered skill %q also appears as a gap", id))
			}
			gapSkills[id] = true
		}
	}
	for id, contract := range skills {
		if contract.Status == "active" && contract.Assignable && !coveredSkills[id] && !gapSkills[id] {
			problems = append(problems, fmt.Errorf("seed plan: active skill %q is neither covered nor an explicit gap", id))
		}
	}
	return errors.Join(problems...)
}

func CandidateID(candidate catalog.Candidate) string {
	return candidate.ExerciseKind + "/" + candidate.Slug
}

func expectedProvenance(candidate catalog.Candidate) string {
	if len(candidate.GeneratedTestDataPaths) > 0 {
		if len(candidate.ProvenanceRefs) > 0 {
			return "multi-hop-exact"
		}
		return "generated-data-unresolved"
	}
	return "source-exact"
}

func isAcceptedStatus(status string) bool {
	return status == "accept-as-seed" || status == "accept-after-rewrite" || status == "split-into-multiple-seeds" || status == "merge-with-other-candidate"
}

func isRejectedStatus(status string) bool {
	return strings.HasPrefix(status, "reject-")
}

func validDate(value string) bool {
	parsed, err := time.Parse("2006-01-02", value)
	return err == nil && parsed.Format("2006-01-02") == value
}

func strictlySorted(values []string) bool {
	for index := 1; index < len(values); index++ {
		if values[index] <= values[index-1] {
			return false
		}
	}
	return true
}

func overlap(left, right []string) bool {
	seen := map[string]bool{}
	for _, value := range left {
		seen[value] = true
	}
	for _, value := range right {
		if seen[value] {
			return true
		}
	}
	return false
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func stringSet(values ...string) map[string]bool {
	result := make(map[string]bool, len(values))
	for _, value := range values {
		result[value] = true
	}
	return result
}

func (scores Scores) values() map[string]int {
	return map[string]int{
		"go_specificity":                      scores.GoSpecificity,
		"conceptual_focus":                    scores.ConceptualFocus,
		"test_quality":                        scores.TestQuality,
		"edge_case_quality":                   scores.EdgeCaseQuality,
		"production_transfer":                 scores.ProductionTransfer,
		"rewrite_burden":                      scores.RewriteBurden,
		"uniqueness":                          scores.Uniqueness,
		"suitability_for_beginners":           scores.SuitabilityBeginners,
		"suitability_for_deliberate_practice": scores.SuitabilityDeliberatePractice,
	}
}
