# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- The first result is true exactly for identifier "alpha" or "beta".
- The second result is true exactly when version is at least 18, independently of the identifier.
- Identifiers are case-sensitive and are not trimmed; empty text is allowed.
- Version may be any int, including negative values. Do not combine the two decisions into one result.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact two-entry allowlist independent of threshold | TestPublicCase01 alpha/beta and unlisted identifiers | only-alpha, dependent-list | quoted identifier/version and bool pair actual/expected |
| Version >=18 independent of identifier, including negatives | TestPublicCase01 boundary/negative cases | threshold-strict, dependent-version | version and actual/expected bool pair |
| No trimming or case folding | TestPublicCase01 Alpha and padded beta | trim-id, case-fold | quoted identifier and classification |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- threshold-strict
- only-alpha
- dependent-version
- dependent-list
- trim-id
- case-fold

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/control_flow/exercises/template1/template1.go`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the Apache-2.0 notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

The two independent outcomes retain both source classification operations rather than reducing the task to only a threshold or allowlist.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 6 required defects, hints, route prerequisites and provenance. `.tmp/checks/independent-policies-001.yaml` reports technical pass with 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
