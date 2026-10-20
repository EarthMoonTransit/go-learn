# Content review

Independent-application slot for `embedding.field-promotion`, following its M1 introduction. Independent source/design review approved; task-content review approved on 2026-10-20; disposable technical checks passed. No whole-skill, full-bank minima or learner-mastery claim.

| Requirement | Public assertion / reference learner test case | Realistic defect |
| --- | --- | --- |
| three independent additive updates | nonzero/zero/negative cases | swap-rejected-targets, only-outer-rejected, only-embedded-rejected, assign-accepted, assign-embedded, assign-outer |
| untouched fields | nonzero/zero cases | clear-kept, clear-tag |

Diagnostics name inputs and actual/expected results. Compile errors and timeouts are not accepted as defect evidence. Any result-equivalent implementation is allowed; no AST or method checker.

Every listed defect must also fail a disposable public-only task-check fixture; hidden grader success does not replace public feedback.

Accepted content evidence: reference passes public/full compatibility and deterministic checks; all 8 required mutants compile and are detected. The same defects fail public assertions with an empty grader in a disposable fixture.
