# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Create an integer initialized to 20 and a pointer variable referring to it.
- Return the integer pointer and the address of that pointer variable; neither result is nil and the second dereferences to the first.
- Mutating the integer through either path is shared; replacing the pointer variable must not change the original integer.
- Calls allocate independent integer storage and pointer variables. Compare identities without numeric addresses.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Return address of value20 and address of its pointer variable | TestPublicCase01 checks nil/identity/value | wrong-initial, nil-pointer-variable, parallel-values | invalid chain / dereference paths disagree |
| Shared target via dereference; pointer replacement preserves original target | TestPublicCase01 writes through pp then replaces *pp | parallel-values | pointer replacement actual values |
| Independent target and pointer-variable storage across calls | TestPublicCase02 | shared-target, shared-pointer-variable | calls share pointer or value storage |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- wrong-initial
- nil-pointer-variable
- parallel-values
- shared-target
- shared-pointer-variable

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/pointers/exercises/template1/template1.go`. Exact declared artifacts and SHA256 values were checked against the pinned local checkout; Apache-2.0 attribution and notice were verified. Source narrative/process output is replaced by a technical API and observable checks, with independent boundaries and defects. No reference-only or generated canonical material was reused. Private references and mutants remain excluded from learner payloads.

## Independent review

Tests distinguish the integer address from the pointer-variable address and verify reassignment and independent call storage without numeric addresses.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 5 required defects, hints, route prerequisites and provenance. `.tmp/checks/pointer-chain-001.yaml` reports technical pass and 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.
