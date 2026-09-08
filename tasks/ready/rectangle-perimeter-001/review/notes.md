# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return twice the sum of Width and Height.
- Each dimension is finite between zero and 1000; zero and fractional dimensions are valid.
- Preserve the input value; Width and Height contribute independently.
- Allow absolute error 0.000000001 and reject NaN.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Twice the sum of independent fields | TestPublicCase01 square/asymmetric/individual-field cases | area, width-only, missing-double | rectangle, numeric actual/expected |
| Fractional and zero dimensions are valid | TestPublicCase01 includes fractional and omitted fields | integer-dimensions, zero-invalid | rectangle and expected perimeter |
| Preserve value and honor absolute tolerance | Public and grader reject NaN/use 1e-9; caller value comparison | numeric defects above | rectangle and numeric result |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- area
- width-only
- missing-double
- integer-dimensions
- zero-invalid

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, selected operation in `structs/v8/shapes.go`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

Independent and omitted fields distinguish field use and zero-value behavior; all numeric checks honor the declared tolerance.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 5 required defects, hints, route prerequisites and provenance. `.tmp/checks/rectangle-perimeter-001.yaml` reports technical pass with 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
