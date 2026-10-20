# Content review

Independent-application slot for `expressions.bitwise-masks`, following its M1 introduction. Independent source/design review approved; task-content review approved on 2026-10-20; disposable technical checks passed. No whole-skill, full-bank minima or learner-mastery claim.

| Requirement | Public assertion / reference learner test case | Realistic defect |
| --- | --- | --- |
| selected bits can become zero | 255/0/15 | or-only, ignore-replacement |
| preserve unselected bits | 128/1/1 and zero mask | drop-unselected, ignore-mask |
| exclude unselected replacement bits | 0/255/15 | take-all-replacement |
| correct operand/destination and overlap | 165/60/90,195/3/3,full mask | swap-selection, xor-instead-replace |

Diagnostics name inputs and actual/expected results. Compile errors and timeouts are not accepted as defect evidence. Any result-equivalent implementation is allowed; no AST or method checker.

Every listed defect must also fail a disposable public-only task-check fixture; hidden grader success does not replace public feedback.

Accepted content evidence: reference passes public/full compatibility and deterministic checks; all 7 required mutants compile and are detected. The same defects fail public assertions with an empty grader in a disposable fixture.
