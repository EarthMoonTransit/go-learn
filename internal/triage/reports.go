package triage

import (
	"fmt"
	"sort"
	"strings"

	"go-practice-bank/internal/safeio"
)

const maxReportBytes = 16 << 20

var reportReviewStatuses = []string{
	"accept-as-seed", "accept-after-rewrite", "split-into-multiple-seeds", "merge-with-other-candidate",
	"idea-only", "deferred", "reject-duplicate", "reject-low-value", "reject-bad-tests",
	"reject-too-algorithmic", "reject-too-broad", "reject-provenance", "reject-license",
}

func GenerateReports(root string, bundle Bundle, skills map[string]SkillContract) error {
	outputs := map[string]string{
		"reports/exercism-candidate-scorecard.md":     scorecardReport(bundle.Reviews),
		"reports/exercism-seed-selection.md":          seedSelectionReport(bundle, skills),
		"reports/exercism-rejected-candidates.md":     rejectedReport(bundle.Reviews),
		"reports/exercism-taxonomy-gaps.md":           gapReport(bundle.Seeds, skills),
		"reports/exercism-task-form-opportunities.md": taskFormReport(bundle.Reviews),
		"reports/exercism-candidate-duplicates.md":    duplicateReport(bundle.Clusters),
	}
	paths := make([]string, 0, len(outputs))
	for path := range outputs {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		content := []byte(strings.TrimRight(outputs[path], "\n") + "\n")
		if _, err := safeio.AtomicWrite(root, path, content, maxReportBytes); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}

func scorecardReport(file ReviewFile) string {
	statusCounts := map[string]int{}
	distributions := map[string][6]int{}
	type ranked struct {
		id     string
		score  int
		status string
	}
	var ranking []ranked
	for _, review := range file.Reviews {
		statusCounts[review.ReviewStatus]++
		for label, score := range review.Scores.values() {
			values := distributions[label]
			values[score]++
			distributions[label] = values
		}
		quality := review.Scores.GoSpecificity + review.Scores.ConceptualFocus + review.Scores.TestQuality + review.Scores.EdgeCaseQuality + review.Scores.ProductionTransfer + review.Scores.Uniqueness + review.Scores.SuitabilityBeginners + review.Scores.SuitabilityDeliberatePractice + (5 - review.Scores.RewriteBurden)
		ranking = append(ranking, ranked{id: review.CandidateID, score: quality, status: review.ReviewStatus})
	}
	sort.Slice(ranking, func(i, j int) bool {
		if ranking[i].score != ranking[j].score {
			return ranking[i].score > ranking[j].score
		}
		return ranking[i].id < ranking[j].id
	})
	var b strings.Builder
	b.WriteString("# Exercism candidate scorecard\n\n")
	fmt.Fprintf(&b, "Reviewed candidates: **%d**. Source commit: `%s`. Scores are reviewed 0–5 judgments; rewrite burden 5 means the largest rewrite. Presence of upstream tests is not an acceptance signal.\n\n", len(file.Reviews), file.SourceCommit)
	b.WriteString("## Decision status breakdown\n\n| Status | Candidates |\n|---|---:|\n")
	for _, status := range reportReviewStatuses {
		fmt.Fprintf(&b, "| `%s` | %d |\n", status, statusCounts[status])
	}
	b.WriteString("\n## Score distributions\n\n| Dimension | 0 | 1 | 2 | 3 | 4 | 5 |\n|---|---:|---:|---:|---:|---:|---:|\n")
	for _, label := range sortedDistributionKeys(distributions) {
		values := distributions[label]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %d | %d |\n", label, values[0], values[1], values[2], values[3], values[4], values[5])
	}
	b.WriteString("\n## Highest combined pedagogical scores\n\nThe combined score rewards the eight positive dimensions and reverses rewrite burden. It is a report aid, not an automatic decision rule.\n\n| Candidate | Score / 45 | Status |\n|---|---:|---|\n")
	limit := 20
	if len(ranking) < limit {
		limit = len(ranking)
	}
	for _, item := range ranking[:limit] {
		fmt.Fprintf(&b, "| `%s` | %d | `%s` |\n", item.id, item.score, item.status)
	}
	return b.String()
}

func seedSelectionReport(bundle Bundle, skills map[string]SkillContract) string {
	covered := map[string]bool{}
	byDomain := map[string]map[string]bool{}
	for _, seed := range bundle.Seeds.SelectedSeeds {
		for _, id := range seed.TargetSkills {
			covered[id] = true
			domain := skills[id].Domain
			if byDomain[domain] == nil {
				byDomain[domain] = map[string]bool{}
			}
			byDomain[domain][id] = true
		}
	}
	activeByDomain := map[string]int{}
	for _, skill := range skills {
		if skill.Status == "active" && skill.Assignable {
			activeByDomain[skill.Domain]++
		}
	}
	var b strings.Builder
	b.WriteString("# Exercism foundation pilot seed selection\n\n")
	fmt.Fprintf(&b, "Selected source seeds: **%d**. Approximate distinct skill coverage: **%d** active skills. Seeds are planning inputs, not ready tasks and not permission to copy canonical statements or generated cases.\n\n", len(bundle.Seeds.SelectedSeeds), len(covered))
	b.WriteString("## Coverage by domain\n\n| Domain | Covered by selected seeds | Active skills |\n|---|---:|---:|\n")
	for _, domain := range sortedIntKeys(activeByDomain) {
		fmt.Fprintf(&b, "| `%s` | %d | %d |\n", domain, len(byDomain[domain]), activeByDomain[domain])
	}
	b.WriteString("\n## Selected seeds\n\n| Candidate | Target skills | Task forms | Pilot role | Rationale |\n|---|---|---|---|---|\n")
	for _, seed := range bundle.Seeds.SelectedSeeds {
		fmt.Fprintf(&b, "| `%s` | %s | %s | `%s` | %s |\n", seed.CandidateID, codeList(seed.TargetSkills), codeList(seed.IntendedTaskForms), seed.PilotRole, md(seed.SelectionRationale))
	}
	return b.String()
}

func rejectedReport(file ReviewFile) string {
	var rejected []CandidateReview
	var weak []CandidateReview
	for _, review := range file.Reviews {
		if isRejectedStatus(review.ReviewStatus) {
			rejected = append(rejected, review)
		}
		if review.Scores.TestQuality <= 2 || review.ReviewStatus == "reject-bad-tests" {
			weak = append(weak, review)
		}
	}
	var b strings.Builder
	b.WriteString("# Rejected Exercism candidates\n\n")
	fmt.Fprintf(&b, "Rejected candidates: **%d**. Rejection is a pedagogical/provenance decision for this bank, not a claim that the upstream exercise is universally poor.\n\n", len(rejected))
	b.WriteString("| Candidate | Status | Test score | Reason |\n|---|---|---:|---|\n")
	for _, review := range rejected {
		fmt.Fprintf(&b, "| `%s` | `%s` | %d | %s |\n", review.CandidateID, review.ReviewStatus, review.Scores.TestQuality, md(review.DecisionReason))
	}
	b.WriteString("\n## Candidates with weak tests\n\n| Candidate | Status | Test score | Edge-case score | Review |\n|---|---|---:|---:|---|\n")
	for _, review := range weak {
		fmt.Fprintf(&b, "| `%s` | `%s` | %d | %d | %s |\n", review.CandidateID, review.ReviewStatus, review.Scores.TestQuality, review.Scores.EdgeCaseQuality, md(review.ReviewerRationale))
	}
	return b.String()
}

func gapReport(plan SeedPlan, skills map[string]SkillContract) string {
	var b strings.Builder
	b.WriteString("# Exercism taxonomy gaps after pilot seed selection\n\n")
	total := 0
	for _, gap := range plan.Gaps {
		total += len(gap.SkillIDs)
	}
	fmt.Fprintf(&b, "Explicit uncovered active skills: **%d**. These gaps are assigned to an origin strategy instead of forcing an unsuitable Exercism candidate into coverage.\n\n", total)
	for _, gap := range plan.Gaps {
		byDomain := map[string][]string{}
		for _, id := range gap.SkillIDs {
			byDomain[skills[id].Domain] = append(byDomain[skills[id].Domain], id)
		}
		fmt.Fprintf(&b, "## %s (%d skills)\n\n%s\n\n", gap.RequiredOrigin, len(gap.SkillIDs), md(gap.Rationale))
		for _, domain := range sortedSliceKeys(byDomain) {
			fmt.Fprintf(&b, "- `%s`: %s\n", domain, codeList(byDomain[domain]))
		}
		b.WriteByte('\n')
	}
	b.WriteString("## Taxonomy representation findings\n\n")
	b.WriteString("- Bitwise expressions are not an active taxonomy v3 skill. `practice/allergies`, `practice/eliuds-eggs`, `practice/grains`, and `practice/secret-handshake` therefore remain deferred or idea-only rather than being mislabeled as arithmetic.\n")
	b.WriteString("- Regular-expression practice and documentation-comment authoring remain scope-deferred in the Exercism mapping; they are not counted as seed coverage.\n")
	b.WriteString("- This report records the gap without changing taxonomy v3; any new skill requires a separate ADR and migration plan.\n")
	return b.String()
}

func taskFormReport(file ReviewFile) string {
	counts := map[string]int{}
	var versatile []CandidateReview
	for _, review := range file.Reviews {
		for _, form := range review.PossibleTaskForms {
			counts[form]++
		}
		if len(review.PossibleTaskForms) >= 4 {
			versatile = append(versatile, review)
		}
	}
	var b strings.Builder
	b.WriteString("# Exercism task-form opportunities\n\n")
	b.WriteString("A form is listed only when it can test a distinct capability; cosmetic reskins are excluded. Rejected candidates have no proposed forms.\n\n")
	b.WriteString("| Task form | Candidate opportunities |\n|---|---:|\n")
	for _, form := range sortedKeys(counts) {
		fmt.Fprintf(&b, "| `%s` | %d |\n", form, counts[form])
	}
	b.WriteString("\n## Candidates supporting at least four distinct forms\n\n| Candidate | Status | Forms |\n|---|---|---|\n")
	for _, review := range versatile {
		fmt.Fprintf(&b, "| `%s` | `%s` | %s |\n", review.CandidateID, review.ReviewStatus, codeList(review.PossibleTaskForms))
	}
	return b.String()
}

func duplicateReport(file ClusterFile) string {
	fileHashGroups, testHashGroups := 0, 0
	var exactFileHashes []string
	for _, cluster := range file.Clusters {
		if len(cluster.Signals.SharedFileHashes) > 0 {
			fileHashGroups++
			exactFileHashes = append(exactFileHashes, cluster.Signals.SharedFileHashes...)
		}
		if len(cluster.Signals.SharedTestHashes) > 0 {
			testHashGroups++
		}
	}
	sort.Strings(exactFileHashes)
	var b strings.Builder
	b.WriteString("# Reviewed Exercism candidate duplicates\n\n")
	fmt.Fprintf(&b, "Reviewed semantic clusters: **%d**. Exact hashes, concept sets, public API shapes, and parser-only AST shapes are signals; the recorded decision is a human review and is never inferred from one signal. No upstream code was executed.\n\n", len(file.Clusters))
	fmt.Fprintf(&b, "Mechanical baseline: **%d** shared individual-file hash groups and **%d** identical non-empty test-hash groups. Shared file hashes represented below: %s.\n\n", fileHashGroups, testHashGroups, codeList(exactFileHashes))
	b.WriteString("| Cluster | Members | Signals | Decision | Preferred | Reviewed conclusion |\n|---|---|---|---|---|---|\n")
	for _, cluster := range file.Clusters {
		fmt.Fprintf(&b, "| `%s` | %s | %s | `%s` | %s | %s |\n", cluster.ID, codeList(cluster.Members), codeList(cluster.SignalTypes), cluster.Decision, codeList(cluster.PreferredCandidates), md(cluster.Signals.ReviewedSemanticSimilarity))
	}
	b.WriteString("\n## Detailed signal notes\n\n")
	for _, cluster := range file.Clusters {
		fmt.Fprintf(&b, "### %s\n\n- Public API: %s\n- AST shape: %s\n- Concept sets: %s\n- Shared file hashes: %s\n- Shared test hashes: %s\n- Rationale: %s\n\n", cluster.ID, valueOrNone(cluster.Signals.PublicAPISimilarity), valueOrNone(cluster.Signals.ASTSimilarity), valueOrNone(strings.Join(cluster.Signals.ConceptSets, ", ")), valueOrNone(strings.Join(cluster.Signals.SharedFileHashes, ", ")), valueOrNone(strings.Join(cluster.Signals.SharedTestHashes, ", ")), md(cluster.ReviewerRationale))
	}
	return b.String()
}

func sortedKeys(values map[string]int) []string {
	result := make([]string, 0, len(values))
	for value := range values {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func sortedDistributionKeys(values map[string][6]int) []string {
	result := make([]string, 0, len(values))
	for value := range values {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func sortedIntKeys(values map[string]int) []string { return sortedKeys(values) }

func sortedSliceKeys(values map[string][]string) []string {
	result := make([]string, 0, len(values))
	for value := range values {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func codeList(values []string) string {
	if len(values) == 0 {
		return "—"
	}
	parts := make([]string, len(values))
	for index, value := range values {
		parts[index] = "`" + md(value) + "`"
	}
	return strings.Join(parts, ", ")
}

func valueOrNone(value string) string {
	if value == "" {
		return "none"
	}
	return md(value)
}

func md(value string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "|", "\\|", "`", "\\`", "\r", " ", "\n", " ").Replace(value)
}
