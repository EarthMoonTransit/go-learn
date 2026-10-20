# Content review

Status: independently approved on 2026-10-20 by reviewer /root/p21_admission_review.

One introduction encounter for `embedding.field-promotion`; no whole-skill depth or learner-mastery claim. Source/original admission: PR100; no external content donor. All code is original.

| Requirement | Assertion/case | Realistic defect |
| --- | --- | --- |
| outer field | TestSelectedFields outer Code | embedded-for-outer |
| shadowed embedded field | Header.Code | outer-for-embedded |
| promoted Name | promoted Name | label-for-name |
| qualified siblings | Left.Value/Right.Value | swap-siblings, left-for-both, missing-right |

Diagnostics identify the actual/expected value or representation. Mutants must compile; only assertion failures count. Disposable task-check evidence passes: reference and 6/6 compiling defects.

Every listed defect also fails the public suite alone: the disposable public-only fixture replaces grader assertions with an empty test package and passes all reference/defect gates.

Any result-equivalent implementation is permitted. No custom AST/method checker.

The independent reviewer checked source admission, original provenance, license, API, prompt alignment, leakage boundaries and route prerequisites; all seven review flags are true. No authored code ran from source checkout and no real learner state was inspected. Full batch acceptance and exact-head CI are recorded separately; these fixtures do not establish learner mastery.
