# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-08. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Bounded count and early EOF success | TestPublicCase01 | copy-all, copy-exact | copy got count/error/data want … |
| Negative and zero limits do no I/O; preserve original non-nil ErrLimit | TestPublicCase02 and TestPreservedLimitSentinel | fresh-limit-error, reassign-sentinel | limit/read/write counts; ErrLimit sentinel identity |
| Bytes plus read error, partial write count, original write error and ErrShortWrite | TestPublicCase03–04 | swallow-error | read error; partial write |
| Caller retains stream ownership on success, error and zero limit | TestPublicCase02–04 and TestSuccessfulCopyKeepsStreamsOpen | close-reader, close-writer, close-on-success | closes got … want0,0 |

## Instructional and technical evidence

Primary skill: `io.copy-limit-composition`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `io.copy-limit-composition`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `copy-all`, `copy-exact`, `swallow-error`, `fresh-limit-error`, `close-reader`, `close-writer`, `reassign-sentinel`, `close-on-success`. The author-provided `.tmp/checks/bounded-copy-001.yaml` records technical PASS, reference PASS and 8/8 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
