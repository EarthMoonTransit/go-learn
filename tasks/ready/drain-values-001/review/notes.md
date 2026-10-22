# Introduction review

Contract/source decision approved by /root/reviewer; actual content/method and targeted execution independently approved 2026-10-22. First encounter only; no learner mastery or full-bank claim.

| Requirement | Public/reference assertion | Defect / method evidence |
| --- | --- | --- |
| wait through empty intervals until producer close | TestWaitForProducerClose/empty,gated-values | immediate-empty, first-only, empty-is-done; diagnostic names premature return phase |
| every value/order/zero/duplicates | TestWaitForProducerClose/gated-values; TestAlreadyClosedBufferedInput | stop-on-zero, sorted-values, reverse-values, deduplicate; actual/want sequence |
| non-nil empty | TestWaitForProducerClose/empty; TestAlreadyClosedBufferedInput | nil-empty; explicit nil diagnostic |
| independent input/result values | TestCallsKeepOwnResults | shared-pending; held A/B overlap, A completes while B remains open, actual/want own values |
| actual range and sole producer close path | TASK method checklist plus executed reference | manual receive loop is behavior-valid but method-incomplete; supplied receive-only restriction earns no credit |

Diagnostics identify case/input and actual versus expected result or lifecycle phase. Compile errors, timeouts and panics do not establish defect detection. Method review inspects actual code and the learner-visible checklist in TASK.md, separately from automated behavior.

Every behavioral defect must fail public-only checks in disposable task-check fixtures. A manual receive loop is a valid behavioral alternative but fails the separate channel-range method criterion.

Actual sample evidence: canonical and public-only9/9, race pass; manual-receive equivalent passes behavior9/9 but fails prescribed range method. Removing TestCallsKeepOwnResults leaves only shared-pending alive8/9, preserving reference and race pass. This requirement→public-assertion→mutant relationship is retained by ordinary task-check; no tooling extension.
