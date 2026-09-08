# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return nil when err is nil, regardless of operation text.
- Otherwise return an error with text operation + ": " + err.Error(). Preserve spaces and empty operation text.
- The returned error must unwrap directly to the original err, preserving its identity.
- Repeated wrapping adds one context layer each time; do not flatten earlier layers or modify the original error.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Nil remains nil; nonnil gains exact operation context | TestPublicCase01/02/03 | nil-wrapped, no-context, trim-operation | nil result or wrapped/layer text |
| Direct original Unwrap identity and one layer per wrapping | TestPublicCase01/02 | text-only, flatten-chain | wrapped cause and retained layer identity |
| Use Error text even when fmt.Formatter differs | TestUsesErrorText custom cause | formatter-instead-of-error | Error text actual versus read: error text with original cause |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- text-only
- no-context
- nil-wrapped
- trim-operation
- flatten-chain
- formatter-instead-of-error

## Source boundary

Original task under ADR0008: canonical project identity, dated creation record and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest.

## Independent review

The reference preserves exact Error text and direct cause identity even for fmt.Formatter errors. Uniform helper method signatures keep the existing checker API contract stable; errors.typed is an explicit prerequisite.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 6 required defects, hints, route prerequisites and provenance. `.tmp/checks/error-context-wrap-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.
