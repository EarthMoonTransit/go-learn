# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Helper registration precedes exactly one acquisition | TestPublicCase01–02 inspect helper inside callback | missing-helper, late-helper | Helper must precede acquire/failed acquire |
| Acquisition failure returns original error without cleanup or closing | TestPublicCase02 | discard-acquire-error | failed acquisition with calls/cleanups/closes |
| Success returns original resource and defers exactly one close to scope cleanup | TestPublicCase01–03 plus real-scope grader integration | close-immediately, no-registration, double-close | setup; cleanup closes; LIFO order |
| Original close error passed in exact Errorf call once | TestPublicCase03 records format and raw arguments | ignore-close-error, rewritten-close-error | Errorf must preserve format and original close error |

## Instructional and technical evidence

Primary skill: `testing.cleanup`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `testing.cleanup`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `close-immediately`, `no-registration`, `ignore-close-error`, `discard-acquire-error`, `double-close`, `missing-helper`, `late-helper`, `rewritten-close-error`. The author-provided `.tmp/checks/cleanup-order-review.yaml` records technical PASS, reference PASS and 8/8 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
