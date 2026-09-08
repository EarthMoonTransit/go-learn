# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- The supplied Rate constant is 0.05 and Days is the integer constant 365.
- Rate32 returns Rate converted to float32, and DayCount returns Days.
- AnnualFraction returns days divided by Days using floating-point division; days is 0..10000.
- Check exported constants as well as functions. Rate32 must equal float32(0.05); allow absolute error 0.000000001 for AnnualFraction and reject NaN.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Supplied exported constants have their stated values | TestContractCase01 and Case02 direct Rate/Days checks | integer-rate, wrong-days, wrong-rate-declaration, wrong-days-declaration | Rate actual/expected; Days constant and function values |
| Exported uses convert Rate32 and return DayCount | TestContractCase02 | rate-percentage, fixed-day-count | actual float32 or integer value |
| AnnualFraction uses fractional division with stated tolerance | TestContractCase03 days0/1/365/730, NaN guard | integer-fraction, fraction-inverted | days and numeric actual/expected |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- integer-rate
- wrong-days
- rate-percentage
- fixed-day-count
- integer-fraction
- fraction-inverted
- wrong-rate-declaration
- wrong-days-declaration

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, selected operation in `exercises/concept/savings-account/.docs/instructions.md`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

Independent declaration-only defects make direct constant assertions necessary even when all accessors and arithmetic functions remain correct. This tests supplied constant behavior, not learner authorship of declarations.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 8 required defects, hints, route prerequisites and provenance. `.tmp/checks/constant-value-tests-001.yaml` reports technical pass with 8/8 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.

The saved `review/weak-submission.txt` was checked by the author in a disposable copy. `.tmp/checks/constant-value-tests-weak.yaml` reports public reference compatibility and the starter-full insufficiency gate passing. Its detected-mutant summary refers to the strong reference, not the weak suite. The function-only constant suite or address-only array suite therefore remains insufficient.
