# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return one Entry for every input key with its exact associated value. Output order is unspecified.
- Do not omit zero, negative or empty-key entries; do not duplicate entries.
- Preserve the input map. Changing the returned entries must not change input values.
- Nil or empty maps return zero entries; keys may contain Unicode and spaces.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| One exact entry for every key, unordered and without duplicates | TestContractCase02 checks membership/seen keys and values | first-only, erase-values, duplicate-entry | length and missing/duplicate/wrong entry |
| Include empty keys and signed/zero values; handle empty input | TestContractCase02/03 | positive-only, empty-key-ignored, empty-placeholder | entry membership/length and empty result |
| Preserve input and avoid mutable aliasing | TestContractCase02 clones map and modifies output entry | destructive | input changed / result shares mutable state |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- first-only
- positive-only
- empty-key-ignored
- erase-values
- destructive
- duplicate-entry
- empty-placeholder

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/maps/exercises/template1/template1.go`. Exact declared artifacts and SHA256 values were checked against the pinned local checkout; Apache-2.0 attribution and notice were verified. Source narrative/process output is replaced by a technical API and observable checks, with independent boundaries and defects. No reference-only or generated canonical material was reused. Private references and mutants remain excluded from learner payloads.

## Independent review

Membership and uniqueness checks avoid depending on map iteration order while protecting complete entries and input preservation.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 7 required defects, hints, route prerequisites and provenance. `.tmp/checks/map-entry-tests-001.yaml` reports technical pass and 7/7 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.
