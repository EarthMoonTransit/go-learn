# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Dynamic int and finite float64 values format as "number " plus one decimal place using fmt formatting.
- NumberBox formats "box " plus its Number result with one decimal place.
- FancyNumberBox formats "fancy " plus one decimal place: parse Text only for exact FancyNumber values; otherwise use zero. Invalid or overflowing integers use zero.
- NumberBox takes precedence when both interfaces match. Nil and unsupported types return "unsupported".
- Interface method receivers are non-nil. Do not call Value on unsupported FancyNumberBox implementations.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Dispatch int and float64 with one decimal | TestPublicCase01 | int-unsupported, float-integer | Describe dynamic type and quoted result |
| Use NumberBox before overlapping FancyNumberBox | TestPublicCase01 numberHolder and bothBox | no-box, wrong-interface-priority | Describe bothBox expected box 9.0 |
| Fancy branch accepts exact FancyNumber and handles parse failures | TestPublicCase01 overflow, bad text and alternate implementation | all-fancy-values, partial-error-result | Describe FancyNumber expected fancy 0.0; forbidden Value panic |
| Nil and unsupported dynamic types | TestPublicCase01 nil/string/int64 | nil-number, default-number | Describe dynamic type expected unsupported |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- int-unsupported
- float-integer
- no-box
- all-fancy-values
- nil-number
- default-number
- partial-error-result
- wrong-interface-priority

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, operation in `exercises/concept/sorting-room/.docs/instructions.md`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

Public overflow and overlapping-interface cases expose parse-error handling and switch priority; Extract and Describe retain different source operations.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 8 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/dynamic-description-001.yaml` reports technical pass and 8/8 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
