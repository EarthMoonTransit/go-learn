# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Complete stock/item.go and subject/describe.go. New stores Name and Weight with both private counts zero.
- Each setter changes only its matching count; getters return current counts. Inputs are nonnegative integers.
- Describe returns "Name:Weight:OnHand:Sold" with decimal integers, using stock exports and methods.
- Preserve Name exactly, including empty text and spaces; separate Item values have independent state.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| New stores exported fields and initializes private counts | TestPublicCase01 | missing-fields | new fields and counts |
| Each setter updates only its count | TestPublicCase01/02 and TestOnHandPreservesSold | reset-onhand, reset-sold, copy-only-setter, swapped-accessors | SetOnHand unrelated state actual/expected; getters and summary |
| Two-package Describe uses exported state and exact name | TestPublicCase01/02 independent Item values and padded/empty names | wrong-summary | quoted summary and independent item output |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- missing-fields
- reset-onhand
- swapped-accessors
- copy-only-setter
- wrong-summary
- reset-sold

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/exporting/exercises/template1/template1.go`. Exact declared artifacts and SHA256 values were checked against the pinned local checkout; Apache-2.0 attribution and notice were verified. Source narrative/process output is replaced by a technical API and observable checks, with independent boundaries and defects. No reference-only or generated canonical material was reused. Private references and mutants remain excluded from learner payloads.

## Independent review

The learner edits two packages and uses exported fields/accessors. Both setter orders now have public preservation checks. Field visibility is introduced without claiming struct-tag coverage.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 6 required defects, hints, route prerequisites and provenance. `.tmp/checks/package-inventory-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.
