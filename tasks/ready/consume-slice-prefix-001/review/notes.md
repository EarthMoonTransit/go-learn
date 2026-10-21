# Content review

Independent application of `slices.header-model`. Source/design approved by /root/m2_2_review. Distinct contract: slice-window-001 only returns a view; this API changes one caller-owned header across successive operations while copied headers retain their length/capacity.

| Requirement | Public assertion / reference learner test | Realistic defect |
| --- | --- | --- |
| prefix/remaining values and exact capacities | TestTakePrefixHeaders | header-not-advanced, prefix-cap-trimmed, remaining-cap-trimmed |
| shared values and independent copied header | TestTakePrefixHeaders / TestTakePrefixSuccessiveCalls | prefix-detached, clear-consumed-storage, header-not-advanced |
| validation by length, atomic rejection | TestTakePrefixInvalidAndEmpty | accept-spare-values, clear-invalid-header, invalid-detaches-header |
| zero/nil and repeated transitions | TestTakePrefixInvalidAndEmpty / TestTakePrefixSuccessiveCalls | reject-empty-prefix, nil-becomes-present |

Diagnostics report actual and expected values or a named state/alias violation. Compile errors and timeouts are not defect detection. All behaviorally equivalent implementations are accepted; no source-shape, allocation-timing or human-mastery claim.

Every declared defect is also checked in a disposable public-only fixture with the grader removed, using the existing task-check. Hidden detection alone does not establish useful public feedback.

Independent content review approved 2026-10-21 by m2_2_review at the provenance content digest. Targeted technical checks passed; 10/10 required defects detected. Public-only fixtures independently detected every declared defect with assertion diagnostics and accepted the reference.
