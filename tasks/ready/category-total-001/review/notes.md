# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Sum Amount for exact matching Category and Day within inclusive From..To.
- If the category occurs anywhere in input, return the sum and nil even when no matching day lies in the period or From exceeds To.
- If no input record has that category, return zero and the original non-nil ErrCategory sentinel; preserve the sentinel and every input record.
- Empty categories, negative and zero amounts are valid; finite amounts have magnitude <=1000 and at most 100 records. Allow absolute sum error 0.000000001.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact category filtering and inclusive bounds; preserve signed totals | TestPublicCase01 includes another category inside the period and values on both endpoints | exclusive-boundaries, no-category-filter, all-categories-in-period, absolute-amount | total got versus 1.5,nil |
| Known category outside/reversed period is successful zero | TestPublicCase02 | known-only-in-period | known outside actual versus 0,nil |
| Unknown category returns original non-nil sentinel, empty category is valid | TestPublicCase03/04 | fresh-error, empty-category-rejected | original sentinel / empty category diagnostics |
| Preserve input and sentinel also on success; allow stated numerical tolerance | TestPublicCase01 cloned slice, every public/grader case captures original sentinel; all output assertions use 1e-9 and reject NaN | mutate-input, success-reassigned | input changed / sentinel identity changed / numeric mismatch |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- known-only-in-period
- exclusive-boundaries
- no-category-filter
- fresh-error
- absolute-amount
- empty-category-rejected
- mutate-input
- all-categories-in-period
- success-reassigned

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, operation in `exercises/concept/expenses/.docs/instructions.md`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

An in-period record from another category independently exposes filtering errors; original sentinel guards cover both success and failure.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 9 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/category-total-001.yaml` reports technical pass and 9/9 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.

Additional acceptance experiment passed in a disposable repository copy through task-check: saved weak-submission remains insufficient (test-writing), or saved +5e-10 tolerated-output passes both public and full reference checks (implementation). Canonical learner data was not accessed.
