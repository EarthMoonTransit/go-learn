# Finite-v1 author evidence

Contract independently approved by /root/reviewer; owner authorized authoring during I3 CI. Actual independent content and method review accepted by /root/reviewer2026-10-22. Original implementation, test text and educational text were authored from the approved bounded contract. Only task-init and earlier packaging boilerplate were reused. No external material is copied or adapted.

| Requirement | Public assertion | Defect or method | Expected diagnostic or boundary |
| --- | --- | --- | --- |
| Prompt nonnil output; callback runs in worker | TestLaunchReturnsWhileWorkerIsHeld; drainWorker nil assertion | inline-callback; nil-output | Launch did not return while callback held; output=nil |
| Callback exactly once | held/completed/panicking atomic call counters | repeat-callback | callback calls=actual; want exactly 1 |
| One result after completion followed by closure | held-output check and drainWorker for every scenario | duplicate-result; never-close; silence-panic | result count=actual; want exactly 1; channel remains open |
| Normal exact value, nil error, false flag | TestLaunchReturnsWhileWorkerIsHeld; TestIndependentNormalAndPanickingCalls | wrong-normal-value; normal-error | worker value=actual; want expected; normal error=actual; want nil |
| Real string/int panics, including empty/zero, exact context and zero value | TestPanicPolicyForStringAndIntegerValues | drop-panic-flag; drop-context; wrong-panic-value; silence-panic | Panicked actual/expected; exact quoted error text; value actual/expected |
| Independent normal and panic calls | TestIndependentNormalAndPanickingCalls launches a normal callback after draining a panic | shared-panic-state | normal error and Panicked must remain nil/false after a different worker panic |
| Actual recovery in launched goroutine | independent actual method inspection required | every declared mutant also contains recover | absence-of-recover crash earns no detection credit; caller fully drains finite output |

All held callbacks use a channel release before collecting final outcomes. No Sleep/Now/After, timing threshold, mutex contender hold or panic/crash credit is used. Tests and task code run only through task-check disposable copies. Canonical and public-only copied fixtures are required; only the pool additionally uses a correct reversed-result control. Method inspection is separate from automated behavior evidence.

Independent actual method and distinct-action acceptance: Actual worker-local deferred recover maps string/int panic to the published outcome, sends exactly one result and closes. Normal/error values and independent calls accepted. This applies earlier synchronous recovery in a separate worker boundary; no missing-recover crash detection claimed.
