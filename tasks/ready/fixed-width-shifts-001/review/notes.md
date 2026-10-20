# Content review

Status: independently approved on 2026-10-20 by reviewer /root/p21_admission_review.

One introduction encounter for `expressions.integer-shifts`; no whole-skill depth or learner-mastery claim. Source/original admission: PR100; no external content donor. All code is original.

| Requirement | Assertion/case | Realistic defect |
| --- | --- | --- |
| zero count | TestShiftByte count=0 | zero-count-zero |
| fixed-width left shift | TestShiftByte counts=1/7 | left-as-right, left-identity |
| unsigned right | TestShiftByte high-bit values | signed-unsigned-right, right-identity |
| signed right | TestShiftByte 129/128/255 | unsigned-signed-right |
| full count | TestShiftByte count=8/9/max | modulo-count, clamp-count |

Diagnostics identify the actual/expected value or representation. Mutants must compile; only assertion failures count. Disposable task-check evidence passes: reference and 8/8 compiling defects.

Every listed defect also fails the public suite alone: the disposable public-only fixture replaces grader assertions with an empty test package and passes all reference/defect gates.

Any result-equivalent implementation is permitted. No custom AST/method checker. Negative counts and constant representability are deliberately not assessed in this encounter.

Review removed a blanket ban on clamping: reducing counts above eight to eight is result-equivalent and passes the disposable alternative check. The clamp-count defect instead clamps to seven and changes the result.

The independent reviewer checked source admission, original provenance, license, API, prompt alignment, leakage boundaries and route prerequisites; all seven review flags are true. No authored code ran from source checkout and no real learner state was inspected. Full batch acceptance and exact-head CI are recorded separately; these fixtures do not establish learner mastery.
