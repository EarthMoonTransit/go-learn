# Introduction review

Contract/source decision approved by /root/reviewer; actual content/method and targeted execution independently approved 2026-10-22. First encounter only; no learner mastery or full-bank claim.

| Requirement | Public/reference assertion | Defect / method evidence |
| --- | --- | --- |
| deterministic sorted keys and literal key spelling | ExampleSummary Output | reverse-order, first-only, lowercase-keys, missing-key |
| signed/zero values and exact separator | ExampleSummary Output | absolute-negative, drop-zero, semicolon-separator |
| nil/empty contract | ExampleSummary_empty Output | wrong-empty |
| discoverable examples, intentional Output, reader utility | actual source plus native example events | ordinary tests may detect defects but fail method review; supplied empty example excluded |

Diagnostics identify case/input and actual versus expected result or lifecycle phase. Compile errors, timeouts and panics do not establish defect detection. Method review inspects actual code and the learner-visible checklist in TASK.md, separately from automated behavior.

Public checks require correct-implementation compatibility; full mutation detection assesses submitted tests. Discoverable examples, intentional output matching and documentation value still require manual review.

Actual sample evidence: actual reference examples pass correct implementation and detect8/8 required defects; ExampleSummary and ExampleSummary_empty naming, Output and documentation value reviewed. Starter empty example excluded from learner-action credit.
