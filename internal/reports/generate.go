package reports

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"go-practice-bank/internal/adapters/exercism"
	"go-practice-bank/internal/catalog"
	curriculumdata "go-practice-bank/internal/curriculum"
	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/sources"
	"gopkg.in/yaml.v3"
)

const maxReportBytes = 8 << 20

type Summary struct {
	Candidates          int
	Concept             int
	Practice            int
	WithStarter         int
	WithTests           int
	WithReference       int
	WithGeneratedData   int
	WithExactProvenance int
	ExactProvenanceRefs int
	WithWarnings        int
	Incomplete          int
}

type fileRef struct {
	candidate string
	path      string
}

type taxonomy struct {
	Version int `yaml:"version"`
	Domains []struct {
		ID     string `yaml:"id"`
		Title  string `yaml:"title"`
		Skills []struct {
			ID                 string `yaml:"id"`
			Status             string `yaml:"status"`
			Assignable         bool   `yaml:"assignable"`
			MinimumOccurrences struct {
				Primary int `yaml:"primary"`
				Total   int `yaml:"total"`
			} `yaml:"minimum_occurrences"`
		} `yaml:"skills"`
	} `yaml:"domains"`
}

type taxonomyMigration struct {
	FromTaxonomyVersion        int `yaml:"from_taxonomy_version"`
	ToTaxonomyVersion          int `yaml:"to_taxonomy_version"`
	TargetAssignableSkillCount int `yaml:"target_assignable_skill_count"`
	Entries                    []struct {
		OldID         string   `yaml:"old_id"`
		NewIDs        []string `yaml:"new_ids"`
		MigrationType string   `yaml:"migration_type"`
		Rationale     string   `yaml:"rationale"`
	} `yaml:"entries"`
	Additions []struct {
		ID        string `yaml:"id"`
		Rationale string `yaml:"rationale"`
	} `yaml:"additions"`
	LegacyTaskReferences []struct {
		TaskID string `yaml:"task_id"`
	} `yaml:"legacy_task_references"`
}

func Generate(root string, source sources.Source, candidates []catalog.Candidate, config exercism.Config) (Summary, error) {
	taxonomyBytes, err := safeio.ReadFile(root, "curriculum/skills.yaml", 4<<20)
	if err != nil {
		return Summary{}, fmt.Errorf("read taxonomy: %w", err)
	}
	var parsed taxonomy
	if err := yaml.Unmarshal(taxonomyBytes, &parsed); err != nil {
		return Summary{}, fmt.Errorf("decode taxonomy: %w", err)
	}
	skillIDs := map[string]curriculumdata.SkillContract{}
	for _, domain := range parsed.Domains {
		for _, skill := range domain.Skills {
			skillIDs[skill.ID] = curriculumdata.SkillContract{Status: skill.Status, Assignable: skill.Assignable}
		}
	}
	taxonomyHash := sha256.Sum256(taxonomyBytes)
	hashText := hex.EncodeToString(taxonomyHash[:])
	mapping, err := curriculumdata.LoadConceptMap(root)
	if err != nil {
		return Summary{}, err
	}
	evidenceFile, err := curriculumdata.LoadEvidence(root)
	if err != nil {
		return Summary{}, err
	}
	upstreamConcepts := make([]string, 0, len(config.Concepts))
	for _, concept := range config.Concepts {
		upstreamConcepts = append(upstreamConcepts, concept.Slug)
	}
	if err := curriculumdata.Validate(mapping, evidenceFile, source, parsed.Version, skillIDs, upstreamConcepts, candidates); err != nil {
		return Summary{}, fmt.Errorf("validate Exercism curriculum evidence: %w", err)
	}
	migrationBytes, err := safeio.ReadFile(root, "curriculum/migrations/v2-to-v3.yaml", 4<<20)
	if err != nil {
		return Summary{}, fmt.Errorf("read taxonomy migration: %w", err)
	}
	var migration taxonomyMigration
	if err := yaml.Unmarshal(migrationBytes, &migration); err != nil {
		return Summary{}, fmt.Errorf("decode taxonomy migration: %w", err)
	}
	v2Bytes, err := safeio.ReadFile(root, "curriculum/migrations/snapshots/skills-v2.yaml", 4<<20)
	if err != nil {
		return Summary{}, fmt.Errorf("read taxonomy v2 snapshot: %w", err)
	}
	var v2 taxonomy
	if err := yaml.Unmarshal(v2Bytes, &v2); err != nil {
		return Summary{}, fmt.Errorf("decode taxonomy v2 snapshot: %w", err)
	}
	registry, err := sources.LoadRegistry(root)
	if err != nil {
		return Summary{}, err
	}
	var companion *sources.Source
	if value, ok := registry.ExercismCompanion(source); ok {
		companion = &value
	}

	summary := summarize(candidates)
	outputs := map[string]string{
		"reports/import-exercism-go.md":         importReport(source, companion, candidates, summary),
		"reports/taxonomy-audit-exercism.md":    taxonomyReport(source, candidates, mapping, evidenceFile, hashText),
		"reports/exact-duplicates-exercism.md":  duplicateReport(source, candidates),
		"reports/unmapped-concepts-exercism.md": unmappedReport(source, candidates, mapping, evidenceFile, hashText),
		"reports/taxonomy-v3-migration.md":      taxonomyV3MigrationReport(parsed, v2, migration),
		"reports/taxonomy-v3-coverage-plan.md":  taxonomyV3CoverageReport(parsed, v2, evidenceFile),
		"reports/taxonomy-v3-exercism-map.md":   taxonomyV3ExercismReport(source, mapping, evidenceFile, hashText),
	}
	paths := make([]string, 0, len(outputs))
	for path := range outputs {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		content := []byte(strings.TrimRight(outputs[path], "\n") + "\n")
		if _, err := safeio.AtomicWrite(root, path, content, maxReportBytes); err != nil {
			return Summary{}, fmt.Errorf("write %s: %w", path, err)
		}
	}
	return summary, nil
}

func summarize(candidates []catalog.Candidate) Summary {
	result := Summary{Candidates: len(candidates)}
	for _, candidate := range candidates {
		switch candidate.ExerciseKind {
		case "concept":
			result.Concept++
		case "practice":
			result.Practice++
		}
		if len(candidate.StarterFilePaths) > 0 {
			result.WithStarter++
		}
		if len(candidate.PublicTestPaths) > 0 {
			result.WithTests++
		}
		if len(candidate.ReferenceSolutionPaths) > 0 {
			result.WithReference++
		}
		if len(candidate.GeneratedTestDataPaths) > 0 {
			result.WithGeneratedData++
		}
		if len(candidate.ProvenanceRefs) > 0 {
			result.WithExactProvenance++
			result.ExactProvenanceRefs += len(candidate.ProvenanceRefs)
		}
		if len(candidate.ImportWarnings) > 0 {
			result.WithWarnings++
		}
		if len(candidate.StatementPaths) == 0 || len(candidate.StarterFilePaths) == 0 || len(candidate.PublicTestPaths) == 0 || len(candidate.ReferenceSolutionPaths) == 0 {
			result.Incomplete++
		}
	}
	return result
}

func importReport(source sources.Source, companion *sources.Source, candidates []catalog.Candidate, summary Summary) string {
	type counts struct{ total, starter, tests, reference, generated, exactProvenance, warnings, incomplete int }
	byKind := map[string]counts{"concept": {}, "practice": {}}
	statusCounts := map[string]int{}
	difficultyCounts := map[string]int{}
	warningCounts := map[string]int{}
	provenanceCommits := map[string]int{}
	var incomplete []string
	for _, candidate := range candidates {
		value := byKind[candidate.ExerciseKind]
		value.total++
		if len(candidate.StarterFilePaths) > 0 {
			value.starter++
		}
		if len(candidate.PublicTestPaths) > 0 {
			value.tests++
		}
		if len(candidate.ReferenceSolutionPaths) > 0 {
			value.reference++
		}
		if len(candidate.GeneratedTestDataPaths) > 0 {
			value.generated++
		}
		if len(candidate.ProvenanceRefs) > 0 {
			value.exactProvenance++
		}
		for _, reference := range candidate.ProvenanceRefs {
			provenanceCommits[reference.Commit]++
		}
		if len(candidate.ImportWarnings) > 0 {
			value.warnings++
		}
		if len(candidate.StatementPaths) == 0 || len(candidate.StarterFilePaths) == 0 || len(candidate.PublicTestPaths) == 0 || len(candidate.ReferenceSolutionPaths) == 0 {
			value.incomplete++
			incomplete = append(incomplete, candidate.Slug)
		}
		byKind[candidate.ExerciseKind] = value
		status := candidate.UpstreamStatus
		if status == "" {
			status = "active (implicit)"
		}
		statusCounts[status]++
		if candidate.UpstreamDifficulty == nil {
			difficultyCounts["not provided"]++
		} else {
			difficultyCounts[strconv.Itoa(*candidate.UpstreamDifficulty)]++
		}
		for _, warning := range candidate.ImportWarnings {
			warningCounts[warning]++
		}
	}
	var b strings.Builder
	fmt.Fprintf(&b, "# Exercism Go import report\n\n- Source: `%s`\n- Commit: `%s`\n- License: `%s`\n- Usage mode: `%s`\n\n", md(source.Repository), source.Commit, md(source.License.Expression), md(source.UsageMode))
	if companion != nil {
		fmt.Fprintf(&b, "Registered generated-data companion: `%s` at primary commit `%s` (`%s`, `%s`), with **%d** exact historical revisions approved for header resolution.\n\n", md(companion.Repository), companion.Commit, md(companion.License.Expression), md(companion.UsageMode), len(companion.ProvenanceCommits))
	} else {
		b.WriteString("Generated-data companion is not registered; generated cases remain provenance-incomplete.\n\n")
	}
	fmt.Fprintf(&b, "The report contains metadata and hashes only. No upstream statement, test, starter, or solution content is copied.\n\n")
	fmt.Fprintf(&b, "## Summary\n\n- Candidates: **%d**\n- Concept exercises: **%d**\n- Practice exercises: **%d**\n- With starter files: **%d**\n- With public tests: **%d**\n- With reference solutions: **%d**\n- With generated test data: **%d**\n- With exact generated-case provenance: **%d** candidates / **%d** file references\n- With import warnings: **%d**\n- Incomplete records: **%d**\n\n", summary.Candidates, summary.Concept, summary.Practice, summary.WithStarter, summary.WithTests, summary.WithReference, summary.WithGeneratedData, summary.WithExactProvenance, summary.ExactProvenanceRefs, summary.WithWarnings, summary.Incomplete)
	b.WriteString("## Breakdown\n\n| Kind | Total | Starter | Tests | Reference | Generated data | Exact provenance | Warnings | Incomplete |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, kind := range []string{"concept", "practice"} {
		value := byKind[kind]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %d | %d | %d | %d |\n", kind, value.total, value.starter, value.tests, value.reference, value.generated, value.exactProvenance, value.warnings, value.incomplete)
	}
	b.WriteString("\n## Exact generated-case provenance\n\n")
	writeCountTable(&b, "problem-specifications commit", provenanceCommits)
	b.WriteString("\n## Upstream status\n\n")
	writeCountTable(&b, "Status", statusCounts)
	b.WriteString("\n## Upstream difficulty\n\nConcept exercises do not declare difficulty; their value remains null.\n\n")
	writeCountTable(&b, "Difficulty", difficultyCounts)
	b.WriteString("\n## Warning distribution\n\n")
	if len(warningCounts) == 0 {
		b.WriteString("No warnings.\n")
	} else {
		writeCountTable(&b, "Warning", warningCounts)
	}
	b.WriteString("\n## Incomplete records\n\n")
	if len(incomplete) == 0 {
		b.WriteString("None under the mechanical completeness rule: statement, starter, tests, and reference solution are present.\n")
	} else {
		sort.Strings(incomplete)
		for _, slug := range incomplete {
			fmt.Fprintf(&b, "- `%s`\n", md(slug))
		}
	}
	return b.String()
}

func duplicateReport(source sources.Source, candidates []catalog.Candidate) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Deterministic duplicate signals: Exercism Go\n\nSource commit: `%s`\n\n", source.Commit)
	b.WriteString("These are mechanical signals only. Equal concept sets are weak evidence and are not semantic duplicate findings. Empty hash/test/concept sets are excluded. No AI or fuzzy comparison is used.\n\n")

	byHash := map[string][]fileRef{}
	for _, candidate := range candidates {
		for _, file := range candidate.FileHashes {
			if file.SHA256 != "" {
				byHash[file.SHA256] = append(byHash[file.SHA256], fileRef{candidate: candidate.Slug, path: file.Path})
			}
		}
	}
	var repeatedHashes []string
	for hash, refs := range byHash {
		candidateSet := map[string]bool{}
		for _, item := range refs {
			candidateSet[item.candidate] = true
		}
		if len(candidateSet) > 1 {
			repeatedHashes = append(repeatedHashes, hash)
		}
	}
	sort.Strings(repeatedHashes)
	fmt.Fprintf(&b, "## Shared individual file hashes (%d groups)\n\n", len(repeatedHashes))
	if len(repeatedHashes) == 0 {
		b.WriteString("None.\n\n")
	} else {
		for _, hash := range repeatedHashes {
			refs := byHash[hash]
			sort.Slice(refs, func(i, j int) bool {
				if refs[i].candidate != refs[j].candidate {
					return refs[i].candidate < refs[j].candidate
				}
				return refs[i].path < refs[j].path
			})
			fmt.Fprintf(&b, "- `%s`: ", hash)
			writeLimitedRefs(&b, refs, 12)
			b.WriteByte('\n')
		}
		b.WriteByte('\n')
	}

	testSets := map[string][]string{}
	for _, candidate := range candidates {
		hashes := hashesForPaths(candidate, candidate.PublicTestPaths)
		if len(hashes) > 0 {
			testSets[strings.Join(hashes, ",")] = append(testSets[strings.Join(hashes, ",")], candidate.Slug)
		}
	}
	writeGroupedCandidates(&b, "Identical non-empty test hash sets", testSets)

	slugs, paths := map[string][]string{}, map[string][]string{}
	for _, candidate := range candidates {
		slugs[candidate.Slug] = append(slugs[candidate.Slug], candidate.SourcePath)
		paths[candidate.SourcePath] = append(paths[candidate.SourcePath], candidate.Slug)
	}
	writeGroupedCandidates(&b, "Duplicate slug references", onlyRepeated(slugs))
	writeGroupedCandidates(&b, "Duplicate source-path references", onlyRepeated(paths))

	conceptSets := map[string][]string{}
	for _, candidate := range candidates {
		if len(candidate.UpstreamConcepts) > 0 {
			key := strings.Join(candidate.UpstreamConcepts, ",")
			conceptSets[key] = append(conceptSets[key], candidate.Slug)
		}
	}
	writeGroupedCandidates(&b, "Equal non-empty upstream concept sets (weak signal)", conceptSets)
	return b.String()
}

func unmappedReport(source sources.Source, candidates []catalog.Candidate, mapping curriculumdata.ConceptMap, evidenceFile curriculumdata.EvidenceFile, taxonomyHash string) string {
	var mapped, ambiguous, gaps, deferred []curriculumdata.ConceptMapping
	for _, item := range mapping.Concepts {
		switch item.Status {
		case "mapped":
			mapped = append(mapped, item)
		case "ambiguous":
			ambiguous = append(ambiguous, item)
		case "taxonomy-gap":
			gaps = append(gaps, item)
		case "scope-deferred":
			deferred = append(deferred, item)
		}
	}
	reviewed := curriculumdata.AcceptedEvidenceByConcept(evidenceFile)
	var b strings.Builder
	fmt.Fprintf(&b, "# Unmapped Exercism concepts\n\n- Source commit: `%s`\n- Taxonomy SHA-256: `%s`\n\n", source.Commit, taxonomyHash)
	fmt.Fprintf(&b, "Frozen concept inventory: **%d**; mapped: **%d**; ambiguous: **%d**; taxonomy gaps: **%d**; scope-deferred: **%d**. Fuzzy matching, stemming, and topic normalization are disabled.\n\n", len(mapping.Concepts), len(mapped), len(ambiguous), len(gaps), len(deferred))
	b.WriteString("The authoritative map is `curriculum/mappings/exercism-go.yaml`. Candidate evidence is separately reviewed and never contributes ready-task coverage.\n\n")
	b.WriteString("## Taxonomy gaps\n\n| Upstream concept | Rationale | Candidate evidence |\n|---|---|---|\n")
	for _, item := range gaps {
		fmt.Fprintf(&b, "| `%s` | %s | %s |\n", md(item.Upstream), md(item.Rationale), evidence(candidates, item.Upstream))
	}
	b.WriteString("\n## Scope-deferred concepts\n\n| Upstream concept | Rationale | Candidate evidence |\n|---|---|---|\n")
	for _, item := range deferred {
		fmt.Fprintf(&b, "| `%s` | %s | %s |\n", md(item.Upstream), md(item.Rationale), evidence(candidates, item.Upstream))
	}
	b.WriteString("\n## Ambiguous or partial correspondence\n\nThese labels are broader than one atomic skill and must not assign coverage automatically.\n\n| Upstream concept | Candidate skills | Rationale |\n|---|---|---|\n")
	for _, item := range ambiguous {
		fmt.Fprintf(&b, "| `%s` | %s | %s |\n", md(item.Upstream), codeList(item.SkillIDs), md(item.Rationale))
	}
	b.WriteString("\n## Mapped labels\n\n| Upstream concept | Taxonomy skill | Accepted reviewed candidate evidence |\n|---|---|---|\n")
	for _, item := range mapped {
		fmt.Fprintf(&b, "| `%s` | `%s` | %s |\n", md(item.Upstream), md(item.SkillIDs[0]), reviewedEvidence(reviewed[item.Upstream]))
	}
	b.WriteString("\nTopics are retained in candidate metadata as low-confidence inventory labels and never converted directly to skills. Prerequisites are dependency evidence, not skill encounters.\n")
	return b.String()
}

func taxonomyReport(source sources.Source, candidates []catalog.Candidate, mapping curriculumdata.ConceptMap, evidenceFile curriculumdata.EvidenceFile, taxonomyHash string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Taxonomy v3 audit against Exercism Go\n\n- Source commit: `%s`\n- Taxonomy SHA-256: `%s`\n\n", source.Commit, taxonomyHash)
	fmt.Fprintf(&b, "Taxonomy version: **%d**. Explicit concept mappings: **%d**. Accepted reviewed candidate evidence records: **%d**. Candidate evidence is not ready-task coverage.\n\n", mapping.TaxonomyVersion, len(mapping.Concepts), acceptedEvidenceCount(evidenceFile))
	b.WriteString("## Main findings\n\n")
	b.WriteString("### Broad mechanisms resolved in taxonomy v3\n\nTaxonomy v3 applies the reviewed decompositions for slice storage, rune and UTF-8 semantics, interfaces, errors, testing, goroutine lifecycle, channels, synchronization, and context. Floating-point semantics, string formatting, `fmt.Stringer`, and the standard-library `strings` package now have explicit assignable skills.\n\n")
	b.WriteString("### Deliberate non-mappings\n\n`basics` and `numbers` remain taxonomy gaps because each is an upstream umbrella rather than an atomic learner capability. `comments`, `randomness`, and `regular-expressions` remain scope-deferred.\n\n")

	type finding struct {
		area, current, proposal string
		labels                  []string
	}
	findings := []finding{
		{"Slices", "`slices.model`, `capacity-bounds`, `append-aliasing`, `growth-allocation` overlap across nil/empty state, slicing, allocation and aliasing.", "Consider construction/nil/empty; len/cap/slicing; append/reallocation; aliasing/defensive-copy; copy/overlap; preallocation.", []string{"slices"}},
		{"Strings / bytes / runes / UTF-8", "`strings.construction-conversion` and `unicode.utf8-runes` each cover several independently testable mechanics.", "Split construction/allocation from string↔[]byte↔[]rune conversion; add strings package and formatting; separate rune model from UTF-8 encode/decode.", []string{"strings", "strings-package", "string-formatting", "runes", "stringers"}},
		{"Methods / pointers", "`methods.receivers-method-sets` combines declaration, receiver mutation, method sets and addressable-call convenience; `pointers.receiver-choice` overlaps it.", "Split receiver semantics from method sets; merge receiver-choice into receiver semantics; require manual evidence for method values/expressions.", []string{"methods", "pointers"}},
		{"Interfaces", "Dynamic interface representation is coupled with typed-nil behavior; assertions are coupled with switches.", "Split dynamic type/value from typed nil, and type assertions from type switches.", []string{"interfaces", "type-assertion"}},
		{"Errors", "Creation is coupled with propagation, wrapping with `Is`, and typed-error design with `As`.", "Separate creation/sentinels, propagation/context, wrapping-chain construction, chain inspection, typed errors, and panic/recover boundaries.", []string{"errors", "error-wrapping"}},
		{"Testing", "Exercism has tests but does not label testing as a learning objective; current table/subtests, helper/cleanup, parallel/race and examples/golden pairs are broad.", "Add basic test discovery/`testing.T`/failure diagnostics, then consider splitting the existing paired skills.", nil},
		{"Goroutines / channels / sync", "Only low-confidence topic `concurrency` exists; it cannot select a specific concurrency skill. Several current skills combine ownership/results, leaks/panics, operations/close, select/nil and pipeline/fan/cancellation.", "Split those mechanics; consider directional channels and data-race/happens-before skills through a separate ADR.", []string{"concurrency"}},
		{"Context", "No Exercism concept or candidate provides context evidence.", "Keep current decomposition unchanged from this audit; consider cancellation causes only through an independent curriculum decision.", nil},
	}
	b.WriteString("## Audit trail for the applied decomposition\n\nThe detailed accepted operations are frozen by `docs/adr/0003-taxonomy-v3.md` and `curriculum/migrations/v2-to-v3.yaml`. The observations below explain the original pressure for the split.\n\n| Area | v2 observation | v3 treatment | Candidate evidence |\n|---|---|---|---|\n")
	for _, item := range findings {
		example := "No direct Exercism curriculum evidence"
		if len(item.labels) > 0 {
			example = evidenceMany(candidates, item.labels)
		}
		fmt.Fprintf(&b, "| %s | %s | %s | %s |\n", md(item.area), md(item.current), md(item.proposal), example)
	}
	b.WriteString("\n## Mapping policy\n\n1. Version mappings by source id and exact commit.\n2. Sort labels and candidates lexicographically.\n3. Count only `concepts`/`practices` as declared targets.\n4. Keep prerequisites as dependency evidence and topics as low-confidence inventory.\n5. Do not normalize hyphens, underscores, plurals, or use fuzzy/stemming matches.\n6. Mark one-to-many labels ambiguous instead of assigning every possible skill.\n7. Permit only reviewed overrides keyed by source commit, exercise kind, and slug.\n8. Report WIP/deprecated exercises separately.\n")
	return b.String()
}

func taxonomyV3MigrationReport(current, previous taxonomy, migration taxonomyMigration) string {
	active, deprecated, _, _ := taxonomyTotals(current)
	previousSkills, _, oldPrimary, oldTotal := taxonomyTotals(previous)
	typeCounts := map[string]int{}
	var changed []struct {
		oldID, operation, targets, rationale string
	}
	for _, entry := range migration.Entries {
		typeCounts[entry.MigrationType]++
		if entry.MigrationType != "preserved" {
			changed = append(changed, struct{ oldID, operation, targets, rationale string }{entry.OldID, entry.MigrationType, codeList(entry.NewIDs), entry.Rationale})
		}
	}
	_, _, newPrimary, newTotal := taxonomyTotals(current)
	var b strings.Builder
	b.WriteString("# Taxonomy v3 migration\n\n")
	fmt.Fprintf(&b, "Migration: taxonomy v%d → v%d. Domains: **%d**. Assignable skills: **%d**. Deprecated compatibility stubs: **%d**. v2 skills inventoried: **%d**.\n\n", migration.FromTaxonomyVersion, migration.ToTaxonomyVersion, len(current.Domains), active, deprecated, previousSkills)
	fmt.Fprintf(&b, "Operations: **%d preserved**, **%d split**, **%d merged**, **%d renamed**, **%d deprecated-without-replacement**; additive v3 skills: **%d**. Legacy task resolutions: **%d**.\n\n", typeCounts["preserved"], typeCounts["split"], typeCounts["merged"], typeCounts["renamed"], typeCounts["deprecated"], len(migration.Additions), len(migration.LegacyTaskReferences))
	fmt.Fprintf(&b, "Coverage contract changes from **%d primary / %d total** to **%d primary / %d total** encounters.\n\n", oldPrimary, oldTotal, newPrimary, newTotal)
	b.WriteString("## Split, merge, rename, and deprecate operations\n\n| v2 skill | Operation | v3 assignable targets | Rationale |\n|---|---|---|---|\n")
	for _, item := range changed {
		targets := item.targets
		if targets == "" {
			targets = "—"
		}
		fmt.Fprintf(&b, "| `%s` | %s | %s | %s |\n", md(item.oldID), item.operation, targets, md(item.rationale))
	}
	b.WriteString("\n## Additive v3 skills\n\n| Skill | Rationale |\n|---|---|\n")
	for _, item := range migration.Additions {
		fmt.Fprintf(&b, "| `%s` | %s |\n", md(item.ID), md(item.Rationale))
	}
	b.WriteString("\nEvery v2 id appears exactly once in the machine-readable migration. Changed ids remain as non-assignable deprecated stubs; they are not global aliases.\n")
	return b.String()
}

func taxonomyV3CoverageReport(current, previous taxonomy, evidenceFile curriculumdata.EvidenceFile) string {
	active, deprecated, primary, total := taxonomyTotals(current)
	_, _, oldPrimary, oldTotal := taxonomyTotals(previous)
	evidenced := map[string]bool{}
	for _, review := range evidenceFile.Reviews {
		if review.Decision == "accepted" {
			evidenced[review.SkillID] = true
		}
	}
	type domainTotal struct {
		id                         string
		assignable, primary, total int
		missing                    []string
	}
	var domains []domainTotal
	var largest []string
	for _, domain := range current.Domains {
		value := domainTotal{id: domain.ID}
		for _, skill := range domain.Skills {
			if skill.Status != "active" || !skill.Assignable {
				continue
			}
			value.assignable++
			value.primary += skill.MinimumOccurrences.Primary
			value.total += skill.MinimumOccurrences.Total
			if skill.MinimumOccurrences.Primary == 4 && skill.MinimumOccurrences.Total == 6 {
				largest = append(largest, skill.ID)
			}
			if !evidenced[skill.ID] {
				value.missing = append(value.missing, skill.ID)
			}
		}
		domains = append(domains, value)
	}
	sort.Strings(largest)
	var b strings.Builder
	b.WriteString("# Taxonomy v3 coverage plan\n\n")
	fmt.Fprintf(&b, "Domains: **%d**. Assignable skills: **%d**. Deprecated stubs: **%d**. Required encounters: **%d primary / %d total**, a change of **%+d / %+d** from v2 (**%d / %d**).\n\n", len(current.Domains), active, deprecated, primary, total, primary-oldPrimary, total-oldTotal, oldPrimary, oldTotal)
	b.WriteString("Minimum tiers are 2/4 for specialist mechanisms, 3/5 for the standard baseline, and 4/6 for high-risk ownership, lifetime, encoding, and concurrency mechanisms. A ready task contributes one primary encounter; evidence records contribute none.\n\n")
	b.WriteString("## Domain totals\n\n| Domain | Assignable skills | Primary minimum | Total minimum |\n|---|---:|---:|---:|\n")
	for _, value := range domains {
		fmt.Fprintf(&b, "| `%s` | %d | %d | %d |\n", md(value.id), value.assignable, value.primary, value.total)
	}
	b.WriteString("\n## Largest minimum tier (4 primary / 6 total)\n\n")
	b.WriteString(codeList(largest) + "\n\n")
	b.WriteString("## Assignable skills without accepted Exercism evidence\n\nThis is a source-evidence gap, not a taxonomy defect and not task coverage.\n\n")
	for _, value := range domains {
		if len(value.missing) > 0 {
			fmt.Fprintf(&b, "- `%s` (%d): %s\n", md(value.id), len(value.missing), codeList(value.missing))
		}
	}
	b.WriteString("\n## Over-complexity risks\n\n- Paired toolchain and library skills retained from v2 may still hide independent commands or API contracts; split only after task evidence shows separate failure modes.\n- Very small specialist skills can create artificial sequencing and reporting overhead; the three-distinct-task-form rule is the guardrail.\n- Concurrency decomposition is intentionally detailed because ownership and termination bugs are independently observable, but future content review must reject near-duplicate exercises.\n- Deprecated stubs increase file size and must be removed only through a later explicit migration after all external references are inventoried.\n")
	return b.String()
}

func taxonomyV3ExercismReport(source sources.Source, mapping curriculumdata.ConceptMap, evidenceFile curriculumdata.EvidenceFile, taxonomyHash string) string {
	counts := map[string]int{}
	reviewed := curriculumdata.AcceptedEvidenceByConcept(evidenceFile)
	for _, item := range mapping.Concepts {
		counts[item.Status]++
	}
	var b strings.Builder
	b.WriteString("# Taxonomy v3 Exercism map\n\n")
	fmt.Fprintf(&b, "- Source commit: `%s`\n- Taxonomy SHA-256: `%s`\n\n", source.Commit, taxonomyHash)
	fmt.Fprintf(&b, "All **%d** frozen upstream labels are explicit: **%d mapped**, **%d ambiguous**, **%d taxonomy-gap**, and **%d scope-deferred**. Accepted reviewed evidence records: **%d**; coverage-eligible records: **0**.\n\n", len(mapping.Concepts), counts["mapped"], counts["ambiguous"], counts["taxonomy-gap"], counts["scope-deferred"], acceptedEvidenceCount(evidenceFile))
	b.WriteString("| Upstream label | Status | Skill ids | Accepted evidence |\n|---|---|---|---|\n")
	for _, item := range mapping.Concepts {
		skills := codeList(item.SkillIDs)
		if skills == "" {
			skills = "—"
		}
		fmt.Fprintf(&b, "| `%s` | %s | %s | %s |\n", md(item.Upstream), item.Status, skills, reviewedEvidence(reviewed[item.Upstream]))
	}
	b.WriteString("\n`basics` and `numbers` deliberately remain gaps because they do not identify one atomic skill. Ambiguous labels never assign encounters automatically.\n")
	return b.String()
}

func taxonomyTotals(value taxonomy) (skills, deprecated, primary, total int) {
	for _, domain := range value.Domains {
		for _, skill := range domain.Skills {
			if value.Version < 3 || (skill.Status == "active" && skill.Assignable) {
				skills++
				primary += skill.MinimumOccurrences.Primary
				total += skill.MinimumOccurrences.Total
			} else if skill.Status == "deprecated" {
				deprecated++
			}
		}
	}
	return skills, deprecated, primary, total
}

func writeCountTable(b *strings.Builder, label string, values map[string]int) {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Fprintf(b, "| %s | Count |\n|---|---:|\n", label)
	for _, key := range keys {
		fmt.Fprintf(b, "| %s | %d |\n", md(key), values[key])
	}
}

func codeList(values []string) string {
	items := append([]string(nil), values...)
	sort.Strings(items)
	for index := range items {
		items[index] = "`" + md(items[index]) + "`"
	}
	return strings.Join(items, ", ")
}

func reviewedEvidence(reviews []curriculumdata.EvidenceReview) string {
	if len(reviews) == 0 {
		return "definition-only; no candidate review"
	}
	values := make([]string, 0, len(reviews))
	for _, review := range reviews {
		values = append(values, "`"+md(review.Slug)+"`")
	}
	sort.Strings(values)
	return strings.Join(values, ", ")
}

func acceptedEvidenceCount(evidence curriculumdata.EvidenceFile) int {
	count := 0
	for _, review := range evidence.Reviews {
		if review.Decision == "accepted" {
			count++
		}
	}
	return count
}

func hashesForPaths(candidate catalog.Candidate, paths []string) []string {
	wanted := map[string]bool{}
	for _, path := range paths {
		wanted[path] = true
	}
	var hashes []string
	for _, file := range candidate.FileHashes {
		if wanted[file.Path] && file.SHA256 != "" {
			hashes = append(hashes, file.SHA256)
		}
	}
	sort.Strings(hashes)
	return hashes
}

func writeLimitedRefs(b *strings.Builder, refs []fileRef, limit int) {
	for index, item := range refs {
		if index == limit {
			fmt.Fprintf(b, " … +%d more", len(refs)-limit)
			break
		}
		if index > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(b, "`%s:%s`", md(item.candidate), md(item.path))
	}
}

func writeGroupedCandidates(b *strings.Builder, title string, groups map[string][]string) {
	keys := make([]string, 0, len(groups))
	for key, candidates := range groups {
		if key != "" && len(unique(candidates)) > 1 {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	fmt.Fprintf(b, "## %s (%d groups)\n\n", title, len(keys))
	if len(keys) == 0 {
		b.WriteString("None.\n\n")
		return
	}
	for _, key := range keys {
		values := unique(groups[key])
		fmt.Fprintf(b, "- `%s`: ", md(key))
		for index, value := range values {
			if index > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(b, "`%s`", md(value))
		}
		b.WriteByte('\n')
	}
	b.WriteByte('\n')
}

func onlyRepeated(groups map[string][]string) map[string][]string {
	result := map[string][]string{}
	for key, values := range groups {
		if len(unique(values)) > 1 {
			result[key] = values
		}
	}
	return result
}

func evidence(candidates []catalog.Candidate, concept string) string {
	var slugs []string
	for _, candidate := range candidates {
		if contains(candidate.UpstreamConcepts, concept) || contains(candidate.UpstreamPrerequisites, concept) {
			slugs = append(slugs, candidate.Slug)
		}
		if contains(candidate.UpstreamTopics, concept) {
			slugs = append(slugs, candidate.Slug)
		}
	}
	slugs = unique(slugs)
	if len(slugs) == 0 {
		return "none in registered exercises"
	}
	const limit = 6
	shown := slugs
	if len(shown) > limit {
		shown = shown[:limit]
	}
	parts := make([]string, 0, len(shown))
	for _, slug := range shown {
		parts = append(parts, "`"+md(slug)+"`")
	}
	result := strings.Join(parts, ", ")
	if len(slugs) > limit {
		result += fmt.Sprintf(" (+%d more)", len(slugs)-limit)
	}
	return result
}

func evidenceMany(candidates []catalog.Candidate, concepts []string) string {
	var slugs []string
	for _, concept := range concepts {
		value := evidence(candidates, concept)
		if value != "none in registered exercises" {
			slugs = append(slugs, value)
		}
	}
	if len(slugs) == 0 {
		return "none"
	}
	return strings.Join(slugs, "; ")
}

func unique(values []string) []string {
	seen := map[string]bool{}
	var result []string
	for _, value := range values {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func md(value string) string {
	replacer := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "|", "\\|", "`", "\\`", "\r", " ", "\n", " ")
	return replacer.Replace(value)
}
