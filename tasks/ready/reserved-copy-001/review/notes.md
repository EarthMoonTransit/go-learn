# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return an independent copy with length len(in) and capacity len(in)+extra.
- Preserve input values and order; later output changes and appends must not affect input.
- Extra is 0..1000, input length <=1000. Nil and empty input are valid, including positive extra capacity.
- The exact capacity is observable API behavior, not a timing or allocation benchmark.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact copied values, length and capacity including extra reservation | TestPublicCase01 | extra-as-length, no-reserve, missing-copy | copy contents/length/capacity |
| Independent storage with and without extra capacity | TestPublicCase01 mutates/appends; TestPublicZeroExtraCase01 mutates nonempty extra=0 output | zero-extra-alias | input changed / zero extra shares input |
| Nil/empty input and positive reservation | TestPublicCase02 for extra0 and3 | empty-no-reserve | empty extra, length and capacity |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- extra-as-length
- no-reserve
- missing-copy
- zero-extra-alias
- empty-no-reserve

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

The extra=0 alias defect is now exposed by a public nonempty-copy check, alongside exact spare-capacity behavior.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 5 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/reserved-copy-001.yaml` reports technical pass and 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
