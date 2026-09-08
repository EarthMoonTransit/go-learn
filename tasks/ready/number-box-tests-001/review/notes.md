# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Call Number exactly once and format its result as "box " followed by one decimal place.
- The box may have any implementing concrete type and state; methods are safe and the interface is non-nil.
- Numbers are integers from -1000000 through 1000000, including zero; preserve the sign.
- Use the current method result on every call and add no newline or extra spaces.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Format signed integer with exactly one decimal | TestContractCase01/02 signed/zero/positive values | integer-format, absolute-number, constant-four | quoted box output and expected number |
| Call arbitrary Number implementer exactly once per call | TestContractCase01/02 with mutable and alternate types | twice | call count and quoted result |
| Use current result on every call | TestContractCase02 updates the same object | last-value | number, result and updated call count |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- integer-format
- absolute-number
- twice
- constant-four
- last-value

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, selected operation in `exercises/concept/sorting-room/.docs/instructions.md`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

A learner-controlled implementation verifies exact one-method calls, signed formatting and changing state as a transfer task.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 5 required defects, hints, route prerequisites and provenance. `.tmp/checks/number-box-tests-001.yaml` reports technical pass with 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
