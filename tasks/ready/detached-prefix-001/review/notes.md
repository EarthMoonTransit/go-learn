# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return the first n values with length and capacity exactly n; preserve their order.
- The result owns independent storage: mutations through input or result must not affect the other.
- Bounds satisfy 0 <= n <= len(in). Nil input and n==0 are valid, with zero capacity.
- Preserve input during the call; no garbage-collector timing or absolute memory measurement is required.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact ordered prefix with length and capacity n | TestPublicCase01 | copy-entire, reverse | prefix contents/capacity and input |
| Input/result mutations must be independent | TestPublicCase01 changes both directions and checks input preserved at return | shared-prefix, capped-but-shared | output shares input / input shares output |
| Zero prefix and nil input have zero capacity | TestPublicCase02 starts with excess input capacity; nil grader | empty-retained | empty retained capacity |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- shared-prefix
- capped-but-shared
- copy-entire
- reverse
- empty-retained

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

Bidirectional mutation checks distinguish detached storage from merely capped shared storage; exact prefix capacity is observable without GC timing.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 5 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/detached-prefix-001.yaml` reports technical pass and 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
