# Introduction review

Contract/source decision approved by /root/reviewer; actual content and execution approved 2026-10-22. First encounter only; no learner mastery or full-bank claim.

| Requirement | Public/reference assertion | Defect / method evidence |
| --- | --- | --- |
| lower, upper and outside boundaries | TestNamedPartitions/below,lower,upper,above | exclude-lower, include-upper, ignore-lower, ignore-upper |
| interior, empty, reversed, negative partitions | TestNamedPartitions/interior,empty,reversed,negative-interior | lower-only, empty-is-point, swap-reversed, nonnegative-only |
| learner-selected named table with readable shared assertion | actual reference code + learner method checklist | duplicated hand-written tests may detect defects but do not satisfy table method |

Diagnostics identify case/input and actual versus expected result or lifecycle phase. Compile errors, timeouts and panics do not establish defect detection. Method review inspects actual code and the learner-visible checklist in TASK.md, separately from automated behavior.

Public checks require correct-implementation compatibility; full mutation detection assesses submitted tests. Named data partitions and shared execution still require manual review.

Independent actual-content approval by /root/reviewer on 2026-10-22. Original creation/provenance, API, story-free contract, public feedback and method checklist verified. Targeted technical checks pass; all required defects detected, with applicable race passing. Batch/route/CI acceptance remains separate.

Method specimen: solution/reference/subject_test.go defines nine named rows spanning distinct endpoint/outside/interior/empty/reversed/negative partitions. One loop executes the predicate and a shared assertion reports name, inputs, actual and expected booleans. All eight defects are detected. These rows were authored as a test submission, not provided in the starter.
