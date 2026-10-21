# Content review

Independent application of `functions.function-values`. Source/design approved by /root/m2_2_review. Distinct contract: predicate-filter-001 calls one predicate on independent records; this folds multiple stored functions, handles nil values and early rejection with callback input/call-order observability.

| Requirement | Public assertion / reference learner test | Realistic defect |
| --- | --- | --- |
| ordered calls with current committed input | TestStageOrderCommitAndStop | reuse-initial, reverse-order, double-call |
| reject without commit/count or later calls | TestStageOrderCommitAndStop / TestStageNilRepeatedAndFirstRejection | commit-rejected, continue-after-reject, count-rejected |
| nil entries and repeated function values | TestStageNilRepeatedAndFirstRejection | stop-at-nil, count-nil, double-call |
| function entries remain usable | TestStageEntriesRetained / TestStageOrderCommitAndStop | clear-entries |

Diagnostics report actual and expected values or a named state/alias violation. Compile errors and timeouts are not defect detection. All behaviorally equivalent implementations are accepted; no source-shape, allocation-timing or human-mastery claim.

Every declared defect is also checked in a disposable public-only fixture with the grader removed, using the existing task-check. Hidden detection alone does not establish useful public feedback.

Independent content review approved 2026-10-21 by m2_2_review at the provenance content digest. Targeted technical checks passed; 9/9 required defects detected. Public-only fixtures independently detected every declared defect with assertion diagnostics and accepted the reference.
