# Content review

Independent-application slot for `expressions.integer-shifts`, following its M1 introduction. Independent source/design review approved; task-content review approved on 2026-10-20; disposable technical checks passed. No whole-skill, full-bank minima or learner-mastery claim.

| Requirement | Public assertion / reference learner test case | Realistic defect |
| --- | --- | --- |
| ordinary shifted bit window | 4/8,8/4,12/4,1/16 | left-shift, mask-before-shift, offset-plus-one, always-zero |
| zero width and finite mask | 0/0,0/4 | zero-means-full, ignore-width |
| large widths and unsigned fill | 0/15,0/16,0/17,0/max,1/16 | width-modulo, cap-at-fifteen, signed-fill |
| large offset | 16/16,32/8,max/max | offset-modulo |

Diagnostics name inputs and actual/expected results. Compile errors and timeouts are not accepted as defect evidence. Any result-equivalent implementation is allowed; no AST or method checker.

Public checks establish compatibility with the correct implementation; full mutation checks establish defect detection. This is not public rejection of a hidden implementation. The weak starter leaves surviving defects.

Accepted content evidence: reference passes public/full compatibility and deterministic checks; all 10 required mutants compile and are detected. Mutation detection belongs to full checks; weak starter is insufficient.
