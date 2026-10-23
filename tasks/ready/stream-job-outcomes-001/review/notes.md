# I9 stream author precheck

Contract approved by sole reviewer in stream-pipe-contract-review.json; approved TASK wording unchanged; explicit constraint bullet yields152 words. Author metadata enables disposable checks, not independent acceptance. Earlier pool and one-job recovery tasks remain unchanged preparation at5781a387e119d1a29cbe13bee7670d0fc076f7cc.

| Basic requirement | Public assertion | Wrong variants / diagnostic |
| --- | --- | --- |
| Prompt nonnil output | HeldWorkersAndPromptOutput; CompletedResultStreamsBeforeOtherJob | inline-until-complete, nil-output: missing handle while callbacks held; output=nil |
| Bounded shared workers | held active count; fewer-jobs and empty | serial-workers, unbounded-workers: active held callbacks actual/want |
| Publish completed result before other job finishes | CompletedResultStreamsBeforeOtherJob job0 held/job1 finished | collect-before-publish: completed job1 not published while job0 held |
| Per-job recovery and continuation | WorkerContinuesAfterEveryPanic single-worker alternating panic/normal | recover-whole-worker: later calls/results missing; no uncaught panic |
| Exact panic policy/raw identity | indexed outcomes for pointer error, zero int, empty string | drop-panic-flag, replace-panic-value, panic-as-normal-error, stale-job-state |
| Exact normal pair/complete unique results | all held and sequence checks | drop-error, drop-value-with-error, duplicate-index, repeat-callback, drop-result |
| Closure after worker completion | held readability before release; collect after draining | premature-empty-stream, leave-output-open; early readable/open handle diagnostics |

All17 required mutants retain finite control flow and all panic cases recover. Entire-worker recovery exits safely and closes merged output, leaving only buffered undispatched indices. Premature visible close sends worker results to a separate sufficiently buffered sink, avoiding send-on-closed panic. Omitted-close mutant exits producers; collect gives an ordinary open-handle assertion. Callback gates always release before final collection; no Fatal bypasses release. No Sleep/Now/After, mutex-held Wait, select tie policy or timeout credit. No tests close learner output. Empty/zero values and nil normal errors are explicit. Grader varies counts up to8 and panic values without adding contract rules.

## Actual method and valid alternatives

Inspect actual shared finite queue, exactly min workers, synchronous runJob in each worker with deferred recover at job scope, next queued job after recovery, true merged stream and joined close. Child callback goroutines may preserve observed concurrency/results but violate this boundary; separate behavior-passing method-wrong control records that limit. Source ownership cannot be inferred from API or active counts alone. A fully buffered output and a completion-channel join are accepted controls; exact helper names, output capacity and WaitGroup spelling are not prescribed. Only fan-out-in and goroutine panic-boundary applications are proposed. No extra lifecycle/ownership/error/WaitGroup/component credit.

Actual public-only wrong diagnostics and correct full/public controls must be retained alongside canonical task-check. All task code executes only in checker disposable copies. Original ADR0008 creation/digest with sources=[]; final independent review, promotion and batch integration remain separate.

Sole actual source/method/evidence review approved2026-10-23 without findings;44accepted fixtures with32ordinary public defects across both owned units. Finalbatchdelivery separate.
