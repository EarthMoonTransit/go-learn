# Content review

Status: independently approved on 2026-10-20 by reviewer /root/p21_admission_review.

One introduction encounter for `expressions.bitwise-masks`; no whole-skill depth or learner-mastery claim. Source/original admission: PR100; no external content donor. All code is original.

| Requirement | Assertion/case | Realistic defect |
| --- | --- | --- |
| any/all and zero mask | TestMaskStates cases 0/0,165/60,195/3 | any-is-all, all-is-any, zero-mask-not-all |
| nonzero-mask search, true and false | 0/255,255/255,165/60,195/3,128/1 | any-false-nonzero, any-true-nonzero, all-false-nonzero, all-true-nonzero |
| setting | 195/3,128/1 | set-toggles, set-drops-unrelated |
| clearing | 165/60,85/170 | clear-keeps-selected, clear-drops-unrelated |
| toggling and unrelated bits | 255/255,195/3 | toggle-sets, toggle-drops-unrelated |

Diagnostics identify the actual/expected value or representation. Mutants must compile; only assertion failures count. Disposable task-check evidence passes: reference and 13/13 compiling defects.

The public mutation profile checks correct-implementation compatibility; full checks require detecting each defect. Hidden mutation coverage is not mislabelled as public rejection of the implementation. Starter retains required survivors.

Any result-equivalent implementation is permitted. No custom AST/method checker.

Two saved weak submissions exercise the reviewed omissions: zero-only predicates and ignoring unrelated bits when toggling. They must remain insufficient under full mutation checking while passing correct-implementation compatibility.

Both saved weak suites pass correct-implementation compatibility but remain insufficient under full checking. The 13/13 detected summary belongs to the strong reference, not the weak submissions. Nonzero true/false predicates and preservation of unrelated toggle bits now have independent defects.

The independent reviewer checked source admission, original provenance, license, API, prompt alignment, leakage boundaries and route prerequisites; all seven review flags are true. No authored code ran from source checkout and no real learner state was inspected. Full batch acceptance and exact-head CI are recorded separately; these fixtures do not establish learner mastery.
