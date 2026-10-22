# Finite-v1 author evidence

Contract independently approved by /root/reviewer; owner authorized authoring during I3 CI. Actual independent content and method review accepted by /root/reviewer2026-10-22. Original implementation, test text and educational text were authored from the approved bounded contract. Only task-init and earlier packaging boilerplate were reused. No external material is copied or adapted.

| Requirement | Public assertion | Defect or method | Expected diagnostic or boundary |
| --- | --- | --- | --- |
| Each callback exactly once, including after failure | heldPool callback counters; TestSingleWorkerContinuesAfterError | repeat-callback; failfast-workers | job N callback calls=actual; want exactly 1 |
| Bounded concurrent progress with enough/fewer jobs | TestWorkersStartTogetherWithinBound; TestFewerJobsThanWorkers | serial-workers; unbounded-workers | active callbacks while held=actual; want expected |
| Every exact index/value/error, including value with error | checkPoolResults in every public scenario | drop-result; duplicate-index; drop-error; drop-value-with-error | missing/duplicate index, value actual/expected, exact callback error |
| Join all callbacks before returning | heldPool before safe release and after result collection | return-before-join | RunPool returned before held callbacks completed |
| No failfast; zero/negative values; empty input | TestSingleWorkerContinuesAfterError; TestEmptyPool | failfast-workers; starter empty result | missing job and call count; correct nil/empty both accepted |
| Order is unspecified | checkPoolResults indexes results; disposable reverse-order reference control | correct reverse-order variant must pass | no scheduling-order assertion |
| Actual bounded shared queue, merged result channel and WG ownership | independent actual method inspection required | reference queue/worker loop/Add/Done/Wait/close/drain | behavior alone cannot prove mechanism or exact idle-worker count |

All held callbacks use a channel release before collecting final outcomes. No Sleep/Now/After, timing threshold, mutex contender hold or panic/crash credit is used. Tests and task code run only through task-check disposable copies. Canonical and public-only copied fixtures are required; only the pool additionally uses a correct reversed-result control. Method inspection is separate from automated behavior evidence.

Independent actual method and distinct-action acceptance: Actual shared job queue, bounded worker goroutines, merged result channel, Add-before-launch/Done-per-worker/Wait-before-single-close and drain accepted. Result/error collection and WG applications are distinct from indexed one-job-per-goroutine fixed-slot first. Correct reversed-result control passes; unspecified result order preserved.
