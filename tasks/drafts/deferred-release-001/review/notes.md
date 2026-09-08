# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Call work once and return its integer result unchanged.
- Call release exactly once after work finishes and before WithRelease returns.
- If work panics, still call release and let the original panic continue unchanged. Do not recover it.
- Both callbacks are non-nil; release does not panic. Deferred calls run during normal return and panic unwinding.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Return work result and invoke each callback once, work before release | TestPublicCase01 checks integer result and exact event sequence | release-first, double-release, result-erased, double-work | normal result and events |
| Release during panic unwinding and preserve original panic | TestPublicCase02 uses pointer panic identity and checks events | normal-only | panic original token and work,release sequence |
| Callback prerequisites and lifecycle scope | Function-values prerequisite; negative return grader case | result-erased | negative result and release count |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- normal-only
- release-first
- double-release
- result-erased
- double-work

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

The callback lifecycle is observable on both normal return and panic unwinding. Function-value prerequisites are explicit. This introductory encounter does not claim complete LIFO or deferred-argument-evaluation mastery.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 5 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/deferred-release-001.yaml` reports technical pass and 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
