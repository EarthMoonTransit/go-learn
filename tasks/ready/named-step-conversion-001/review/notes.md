# Content review

Independent application of `types.named-underlying`. Source/design approved by /root/m2_2_review. Distinct contract: Named-array declaration becomes conversion and arithmetic across two different defined scalar types, with their identities visible through a supplied observation wrapper.

| Requirement | Public assertion / reference learner test | Realistic defect |
| --- | --- | --- |
| two distinct defined types with int16 underlying | TestDistinctTypesAndConversion | reading-alias, step-alias, same-defined-type, wide-reading, wide-step, unsigned-step |
| signed conversion and sum with reading result | TestDistinctTypesAndConversion | ignore-step, subtract-step, result-is-step |

Diagnostics name the affected input/position and actual versus expected values or types. Compile errors/timeouts are not semantic detection. Equivalent implementations are accepted; no algorithm/source-shape, performance, full-skill or human-mastery claim.

All declared defects are checked again by existing task-check in disposable public-only fixtures with the grader removed; hidden detection alone is insufficient.

Independent content review APPROVE on 2026-10-22. Targeted technical checks and all applicable public-only fixtures passed. Exact content digests are recorded in the independent batch receipt; route/final diff and repository acceptance remain separate.
