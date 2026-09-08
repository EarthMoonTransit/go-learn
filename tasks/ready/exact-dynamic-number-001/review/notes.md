# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Only a dynamic FancyNumber value is accepted; convert its Text with strconv.Atoi and return the integer.
- Return zero for invalid or overflowing integers, nil interfaces, pointers to FancyNumber and all other implementing types.
- Do not trim whitespace or parse fractional text. Leading signs and leading zeroes follow strconv.Atoi.
- Do not call Value on unsupported types; dynamic type identity determines eligibility.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Accept exact FancyNumber and preserve signed integers | TestPublicCase01 | absolute-value | Extract quoted input, actual integer, expected integer |
| Return zero on invalid syntax and overflow | TestPublicCase01 includes spaces, fraction and overflowing decimal | trim-space, partial-error-result | Extract input and integer mismatch |
| Reject nil, pointers and other implementers without calling Value | TestPublicCase02 includes typed nil pointer and panic-on-call alternate | all-implementers, pointer-accepted, empty-nil-panic | unsupported dynamic type and result; forbidden call panic |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- all-implementers
- pointer-accepted
- trim-space
- absolute-value
- partial-error-result
- empty-nil-panic

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, operation in `exercises/concept/sorting-room/.docs/instructions.md`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

Exact concrete assertion is distinguished from calling every implementer, pointer acceptance and overflow partial values.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/exact-dynamic-number-001.yaml` reports technical pass and 6/6 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
