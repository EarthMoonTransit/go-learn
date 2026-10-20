# Content review

Status: independently approved on 2026-10-20 by reviewer /root/p21_admission_review.

One introduction encounter for `arrays.indexing-iteration`; no whole-skill depth or learner-mastery claim. Source/original admission: PR100; no external content donor. All code is original.

| Requirement | Assertion/case | Realistic defect |
| --- | --- | --- |
| selected elements | middle/first/last/all | range-copy, skip-last, add-instead-replace |
| unselected elements | middle/first/last | include-end, prefix-instead-window, clear-outside |
| empty interval | empty-middle/empty-end | reject-empty |
| invalid bounds | negative/reversed/past-end/both-past | invalid-success |

Diagnostics identify the actual/expected value or representation. Mutants must compile; only assertion failures count. Disposable task-check evidence passes: reference and 8/8 compiling defects.

Every listed defect also fails the public suite alone: the disposable public-only fixture replaces grader assertions with an empty test package and passes all reference/defect gates.

Any result-equivalent implementation is permitted. No custom AST/method checker.

The independent reviewer checked source admission, original provenance, license, API, prompt alignment, leakage boundaries and route prerequisites; all seven review flags are true. No authored code ran from source checkout and no real learner state was inspected. Full batch acceptance and exact-head CI are recorded separately; these fixtures do not establish learner mastery.
