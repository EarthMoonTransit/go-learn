# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Positive counts return a nil error.
- Negative counts return a non-nil *InvalidCountError with the original Count and Reason "negative count"; zero uses Reason "empty count".
- Error formats "<Count> items are invalid: <Reason>" using decimal Count and exact Reason, including for manually constructed errors.
- Check the concrete pointer type, structured fields and displayed text; a generic error with identical text is insufficient.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Positive values return actual nil | TestContractCase01 | typed-nil-success | positive error actual versus nil |
| Negative and zero errors expose exact pointer type, Count and Reason | TestContractCase02 checks type, non-nil pointer and both fields | zero-valid, negative-valid, erased-count, generic-error, same-reason | count, dynamic type, fields and quoted text |
| Error uses arbitrary manually supplied Reason | TestContractCase03 constructs Count 7 with custom reason | fixed-reason-text, reason-derived-from-count | quoted custom Error output |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- zero-valid
- negative-valid
- erased-count
- generic-error
- same-reason
- fixed-reason-text
- typed-nil-success
- reason-derived-from-count

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, operation in `exercises/concept/the-farm/.docs/instructions.md`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

The dedicated reason-derived-from-count defect makes a manually constructed custom Reason necessary beyond checking ValidateCount failures.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 8 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/structured-error-tests-001.yaml` reports technical pass and 8/8 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
