# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Rectangle area is Width times Height; Circle area is math.Pi times Radius squared; Triangle area is Base times Height divided by two.
- AreaOf calls Area through Shape, including custom implementing types supplied by the caller; preserve their returned value.
- Dimensions are finite from 0 through 1000; allow absolute error 0.000000001 and reject NaN.
- Zero dimensions and fractional dimensions are valid; results do not depend on previous calls.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Three geometric formulas | TestContractCase01/02 via AreaOf | rectangle-sum, circle-linear, approximate-pi, triangle-no-half | rectangle/shape index got versus expected |
| Zero and fractional dimensions with tolerance | TestContractCase02 rejects NaN and allows 1e-9 | fraction-truncated, zero-fallback | shape index and numeric result |
| Open Shape contract, including caller-defined implementations | TestCustomAreaContract with a custom Shape and varied returned values | custom-unsupported | custom dispatch got versus expected |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- rectangle-sum
- circle-linear
- approximate-pi
- triangle-no-half
- fraction-truncated
- zero-fallback
- custom-unsupported

## Source boundary

Substantive adaptation of `learn-go-with-tests` at commit `4675d96ad50b815d3e9fdc2122ab3dbc14636252`, operation in `structs/v8/shapes.go`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

AreaOf and the custom-unsupported defect require exercising an open interface with a learner-defined implementation, independently of geometric arithmetic.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 7 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/shape-contract-tests-001.yaml` reports technical pass and 7/7 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
