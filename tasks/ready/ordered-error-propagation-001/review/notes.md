# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Call Amount with count, then Factor, each exactly once on success; return Amount times Factor divided by count.
- On the first error stop immediately and return zero with that exact error; do not call a later dependency.
- The calculator is non-nil, count is 1..1000, successful values are finite with magnitude <=1000. Allow absolute error 0.000000001 for numeric results.
- Preserve error identity even when the dependency also returns a nonzero numeric value.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Call Amount(count), then Factor once and compute the ratio | TestPublicCase01 inspects ordered calls, count and numeric result | wrong-count, double-amount, drop-division | divide value/error, calls and count |
| Stop at the first error; return original error and numeric zero | TestPublicCase02/03 use distinct sentinel errors and nonzero numeric/error pairs | ignore-amount-error, ignore-factor-error, fresh-error | first/second error result and recorded call sequence |
| Allow signed finite values with 1e-9 tolerance | Public fractional case and signed grader case reject NaN | drop-division and arithmetic result checks | actual numeric result and expected result |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- ignore-amount-error
- ignore-factor-error
- fresh-error
- drop-division
- wrong-count
- double-amount

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, selected operation in `exercises/concept/the-farm/.docs/instructions.md`. All declared consulted artifacts and SHA256 values were verified against the pinned local source. MIT notices and attribution were preserved; independently authored boundaries and defects extend the selected observable behavior. No generated canonical or reference-only material was reused. Private reference code and mutants are excluded from learner payloads.

## Independent review

The selected DivideFood operation retains ordered dependency calls and exact error propagation, with independently observable call counts and error/numeric pairs.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/ordered-error-propagation-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
