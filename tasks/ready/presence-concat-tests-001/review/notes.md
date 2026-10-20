# Content review

Independent-application slot for `slices.nil-empty`, following its M1 introduction. Independent source/design review approved; task-content review approved on 2026-10-20; disposable technical checks passed. No whole-skill, full-bank minima or learner-mastery claim.

| Requirement | Public assertion / reference learner test case | Realistic defect |
| --- | --- | --- |
| all absent vs any present | no-parts/all-absent/one-present-empty/empty-after-nil/empty-before-nil | absent-becomes-present, empty-presence-lost |
| all parts in order and all values | later-data/ordered-parts | always-nil, stop-at-nil, first-part-only, reverse-parts, drop-zero, duplicate |
| input element preservation | ordered-parts | mutate-input |

Diagnostics name inputs and actual/expected results. Compile errors and timeouts are not accepted as defect evidence. Any result-equivalent implementation is allowed; no AST or method checker.

Public checks establish compatibility with the correct implementation; full mutation checks establish defect detection. This is not public rejection of a hidden implementation. The weak starter leaves surviving defects.

Review correction: first-part-presence and last-part-presence collect values correctly but derive an empty result's nilness from only one argument. The empty-after-nil and empty-before-nil assertions reject them. Saved weak-single-part-presence.txt omits both multi-part empty cases; its measured insufficiency must remain visible. No new checker.

Accepted content evidence: reference passes public/full compatibility and deterministic checks; all 11 required mutants compile and are detected. Mutation detection belongs to full checks; weak starter is insufficient.
