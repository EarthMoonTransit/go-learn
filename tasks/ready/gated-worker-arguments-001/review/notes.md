# P19 author notes

Original gap decision is merged PR85 with exact registry/pin/artifact records.
The existing-variable '=' starter delays reading until caller closes start after
StartVisits returns. Loop assignments happen before close/receive; wrong values
are deterministic and race-free. This avoids obsolete pre1.22 folklore about
':=' loops. Every element, including duplicates, is an independent invocation;
callback order is deliberately unspecified. Input is held unchanged by caller
until completion. Tests do not require snapshot behavior after caller mutation.

Primary is goroutines.capture-arguments, not generic closure construction or
worker launch alone. Channel signals, joining, function values and slice/range
mechanics are already prepared. No exact goroutine count or synchronization
primitive is prescribed. Timing is provisional, not human-calibrated. No theory
page or real learner-state access is introduced.

## Public requirement feedback

| Requirement | Public assertion | Compiling defect | Useful diagnostic |
| --- | --- | --- | --- |
| Return nonnil handle before start | TestCapturedArguments | wait-for-start-before-return, nil-completion | launcher not returned or nil handle; want nonnil completion channel before release |
| Correct argument once for each element | TestCapturedArguments | existing-variable-capture, index-instead-value, zero-argument, skip-last, duplicate-callback | actual/want multiset including repeated elements; callback starts actual/want |
| No callback before start closes | TestCapturedArguments | no-start-gate | before start callback starts actual, want0 |
| All callbacks can start while others blocked | TestCapturedArguments | serial-callbacks | held callback starts actual, wantN |
| Completion only after all callbacks, no sends | TestCapturedArguments, TestIndependentConcurrentBatches | early-completion, drop-completion, payload-completion | completion closed while pending, still open after callbacks, or sent value; want correct state |
| Nil/empty complete without start | TestEmptyBatches | empty-waits-start | completion still open, want closed |
| Input preserved and batches independent | TestCapturedArguments, TestIndependentConcurrentBatches | input-mutated, shared-completion | input actual/want unchanged; second completion closed while callbacks pending |

Callbacks append under a mutex, release it before waiting on callback gate, and
never send into a capacity-limited observation channel. Duplicate callbacks can
finish after cleanup rather than producing blocked-sender timeout. Return handle
is sent into a buffered channel of capacity 1. Every test-owned start/callback gate is closed
on every exit, then synctest.Wait joins the bubble. Only parent code waits through
synctest. The shared-completion mutant closes through sync.Once, so its bug is
premature cross-batch completion rather than a second-close panic. Empty batches
are allowed to share an already-closed handle; channel identity is not asserted.

Grader expands to 23 repeated values and machine-int extremes. No timing, random,
compiler-error or timeout-only kill is accepted. Reference argument binding plus
per-iteration ':=' and index-binding/shared-empty-handle alternatives must pass
public/full before promotion.

## Measured evidence

Updated bytes pass the disposable fixture technical check, determinism, race,
vet and all 15 compiled/detected required mutants. Standalone public verification
rejects each of 15 defects with category test and named useful actual/wanted
diagnostics. The core existing-variable capture fails its argument multiset;
no DATA RACE diagnostic appears, consistent with the gate happens-before proof.
Reference, per-iteration declaration and index-binding/shared-empty-handle
alternatives all pass public/full.

Actual 22 results, exact authored fixture/canonical byte equality and diagnostic
needles are retained in .tmp/p19-content/public-proof.json, public-summary.json
and review-check.yaml. Public proof SHA-256: `95ee3fcbbf88247711bd16b17de9b8fc90388fe3cf8e1d9e901b34176f1f9073`.
Canonical independent content APPROVE was granted by `/root/p18_content_review`
after reading both evidence hashes, actual 22 results and canonical byte equality.
Official task-promote completed draft → in-review → ready after that approval;
route/docs/reports/full collection wait for P17 canonical. Automated fixture
execution is not learner completion or mastery evidence.
