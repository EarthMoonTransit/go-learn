# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Average is Hits divided by Attempts using floating-point division; zero attempts returns zero.
- Inputs satisfy 0 <= Hits <= Attempts <= 1000000. Label is unrelated to the result.
- MethodFunctions returns non-nil functions corresponding to Sample.Average and (*Sample).Average method expressions. Each computes Average from its current argument; calls are independent.
- Do not change any field of the receiver. Absolute error up to 0.000000001 is allowed.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Fractional ratio and zero attempts | TestPublicCase01/02/03; grader all-hits case | integer-division, reversed-ratio, zero-one, percentage, label-filter | average/zero/value-pointer got versus expected |
| Independent value and pointer function arguments; preserve receiver | TestMethodFunctions varies one pointee and a second value | value-argument-ignored, pointer-argument-ignored, pointer-cached | non-nil functions, method function hits, changed argument |
| Absolute tolerance 1e-9 and finite results | Every public/grader numerical assertion rejects NaN and uses the stated tolerance | Numerical defects above | actual and expected numeric values |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- integer-division
- reversed-ratio
- zero-one
- percentage
- label-filter
- value-argument-ignored
- pointer-argument-ignored
- pointer-cached

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, operation in `topics/go/language/methods/exercises/template1/template1.go`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream Apache-2.0 notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

MethodFunctions makes both method-set function signatures and fresh argument dispatch observable; all numerical checks use the documented tolerance.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 8 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/value-average-001.yaml` reports technical pass and 8/8 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.

Additional acceptance experiment passed in a disposable repository copy through task-check: saved weak-submission remains insufficient (test-writing), or saved +5e-10 tolerated-output passes both public and full reference checks (implementation). Canonical learner data was not accessed.
