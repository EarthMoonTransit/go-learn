# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return ten integers 0..9 in order, five strings "a".."e", and a view containing strings at indices one and two.
- The view shares storage with the five-string slice; view index zero corresponds to original index one.
- Separate calls return independent number and string storage.
- Check values and lengths; do not require an implementation-dependent capacity for the appended integer slice.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact numbers/words and indices1..2 view | TestContractCase01/02 compares values and lengths | missing-last, wrong-start, short-view | numbers and sequence contents |
| Bidirectional shared view into words | TestContractCase02 mutates view and original | detached-view | subslice alias / reverse alias missing |
| Independent allocations across calls without capacity assumptions | TestContractCase02 retains two calls and mutates numbers/words | shared-numbers, shared-words | number storage shared / independence incorrect |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- missing-last
- wrong-start
- short-view
- detached-view
- shared-numbers
- shared-words

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/slices/exercises/template1/template1.go`. Exact declared artifacts and SHA256 values were checked against the pinned local checkout; Apache-2.0 attribution and notice were verified. Source narrative/process output is replaced by a technical API and observable checks, with independent boundaries and defects. No reference-only or generated canonical material was reused. Private references and mutants remain excluded from learner payloads.

## Independent review

The retained append/subslice operations are checked by values, bidirectional sharing and independent calls without a runtime capacity assumption.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 6 required defects, hints, route prerequisites and provenance. `.tmp/checks/slice-sequence-tests-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.
