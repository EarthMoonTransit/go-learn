# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Copy all five source strings into an independent array and return its pointer.
- Return the address of each copied array element in matching index order; addresses must refer to that returned array.
- Changing source or copied elements must not change the other array. Separate calls produce independent arrays.
- Source is non-nil; empty, repeated and Unicode strings are valid. Compare pointer identities, not numeric memory addresses.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Copy all five values with matching element pointers | TestContractCase01 | original-addresses, same-address, omit-final | copy contents and mismatching pointer index |
| Source and copied array are independent | TestContractCase02 mutates both arrays and dereferences returned pointer | source-alias | arrays share state / element pointer targets wrong storage |
| Separate calls return independent arrays | TestContractCase02 retains two returned arrays | shared-array | source and both copied arrays after mutation |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- source-alias
- original-addresses
- same-address
- omit-final
- shared-array

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/arrays/exercises/template1/template1.go`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the Apache-2.0 notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

The coherent source-alias defect now returns matching pointers into the source, so address consistency alone cannot replace the ownership checks.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 5 required defects, hints, route prerequisites and provenance. `.tmp/checks/array-copy-address-tests-001.yaml` reports technical pass with 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.

The saved `review/weak-submission.txt` was checked by the author in a disposable copy. `.tmp/checks/array-copy-address-tests-weak.yaml` reports public reference compatibility and the starter-full insufficiency gate passing. Its detected-mutant summary refers to the strong reference, not the weak suite. The function-only constant suite or address-only array suite therefore remains insufficient.
