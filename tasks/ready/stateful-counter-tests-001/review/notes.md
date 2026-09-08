# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return a non-nil closure whose initial state is start.
- Each call adds delta to that closure state and returns the updated value.
- Closures returned by separate NewCounter calls have independent state, even when starts are equal.
- Start and accumulated results fit int; negative and zero deltas are valid. Calls are sequential.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Initialize and return updated state | TestContractCase01 plus Case02 | start-erased, previous-value | delta, actual and expected counter value |
| Persist signed/zero updates between calls | TestContractCase02 revisits the same closure | no-persistence, negative-ignored | delta, actual and expected counter value |
| Separate closure state even for equal starts | TestContractCase02 interleaves a/b at start5 and c at start-2 | shared-state | independent expected value for each closure call |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- no-persistence
- shared-state
- start-erased
- negative-ignored
- previous-value

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

Interleaved closures distinguish persistent per-instance state from a shared global counter, including equal starts, negative and zero deltas.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 5 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/stateful-counter-tests-001.yaml` reports technical pass and 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
