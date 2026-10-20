# Content review

Independent-application slot for `arrays.indexing-iteration`, following its M1 introduction. Independent source/design review approved; task-content review approved on 2026-10-20; disposable technical checks passed. No whole-skill, full-bank minima or learner-mastery claim.

| Requirement | Public assertion / reference learner test case | Realistic defect |
| --- | --- | --- |
| snapshot/index direction | cycle | identity, in-place, scatter-instead-gather |
| repeated indexes | repeats | reject-repeated-indexes, skip-last |
| invalid map atomicity and status | negative-first/invalid-last/negative-last | invalid-zero-result, invalid-success, partial-on-invalid |

Diagnostics name inputs and actual/expected results. Compile errors and timeouts are not accepted as defect evidence. Any result-equivalent implementation is allowed; no AST or method checker.

Every listed defect must also fail a disposable public-only task-check fixture; hidden grader success does not replace public feedback.

Accepted content evidence: reference passes public/full compatibility and deterministic checks; all 8 required mutants compile and are detected. The same defects fail public assertions with an empty grader in a disposable fixture.
