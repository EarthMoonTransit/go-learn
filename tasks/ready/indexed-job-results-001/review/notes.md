# P17 author notes

Source decision is PR85, merged b20ca9feacaa30b7a2ad20ca2761c4ee14c8c0f8.
The exact Ardan template2 fan-out/collect-all operation is retained, with random
number generation/printing replaced by supplied jobs and indexed value/error
records. Source has no completed body. All code/cases are local; Apache notice,
license and exact artifact hashes are preserved. No SHAD code or adjacent
Ardan answer/template content is copied. `api-shape` denotes adaptation of the
program's operation contract into a callable API, not copied implementation.

Primary is result/error collection; ownership, launch, errors, slices, function
values/closures and WaitGroup are already prepared prerequisites. JoinAll only
joins void callbacks. No particular synchronization primitive or exact goroutine
count is required: observe concurrent progress, indexed accounting and wait-all.
The supplied Result type keeps focus on collection. Timing estimates are
provisional, not human-calibrated. No new theory or learner state access.

## Public requirement feedback

| Requirement | Public assertion | Compiling wrong variants | Useful diagnostic |
| --- | --- | --- | --- |
| Each job once | TestEveryIndexedValueAndError, TestConcurrentStartAndWaitAll | duplicate-call, skip-last | job index actual calls, want exactly1 |
| All jobs can start while others are blocked | TestConcurrentStartAndWaitAll | serial | before release job[index] calls0, want1 |
| Input-index order and Index field | TestEveryIndexedValueAndError, TestConcurrentStartAndWaitAll | wrong-index, completion-order | result[index].Index actual/wanted; ordered values |
| Exact values and errors, including nonzero+error | TestEveryIndexedValueAndError | zero-value, zero-on-error, drop-error, clone-error | actual/wanted value; error identity/type versus original job error |
| Wait every job after early error | TestConcurrentStartAndWaitAll | no-wait, fail-fast | returned while jobs remain blocked, want every job returned first |
| Empty input gives zero records | TestEmptyBatches | empty-result | actual result length, want0; nil and allocated empty accepted |
| Independent returned records | TestIndependentResultRecords, TestConcurrentIndependentBatches | shared-output | first after second call and second after first-result edit actual/wanted values |

The gate test releases jobs in reverse order, beginning with a nonzero value and
error. It uses synctest and atomic call counters, never wall-clock sleeps.
Every test-owned gate is released on failure. Results are inspected only after
synctest synchronization, even for a wrongly premature implementation. The
fail-fast variant buffers N results so a blocked sender cannot masquerade as the
intended early-return defect. Required variants must compile and fail assertions;
timeout/compiler errors do not count as detection. Race is required for all
canonical, defect and public invocations under the current contract.

Public checks also exercise two concurrent independent batches. Grader expands
to repeated values, 19 jobs, typed-nil error and singleton. No panic recovery, cancellation, aggregate-error
format or universal performance threshold is imposed. Empty output may be nil
or allocated. Errors are passed through, not deep-copied or string-normalized.

## Verified evidence

The updated bytes passed the disposable fixture technical check, including
reference determinism, race, vet and all 13 compiled/detected required mutants.
Independent static review by `/root/p18_content_review` confirmed source hashes,
notices, contract/test alignment, gate cleanup and allowed synchronization choices.

Standalone learner public verification rejected all 13 required defects with
category `test`, named assertions and useful actual/wanted diagnostics. In
particular completion-order fails reverse-completion indexing, fail-fast and
no-wait report returning while jobs remain blocked, and serial reports missing
job starts before release. Reference passes public/full. Two valid alternatives
also pass both: an unbuffered indexed-result channel without WaitGroup, and a
hybrid running job0 in the caller with N-1 workers. Nil empty output is accepted.

Measured evidence is retained in `.tmp/p17-content/public-proof.json`,
`public-summary.json` and `review-check.yaml`; proof SHA-256 is
`f1bcece73c27750706c88880d9e7c16a50d59be49d7de1e845b03de636cc856e`.
Every authored fixture file equals the canonical draft bytes. Fixture review and
official promotion are isolated; no fixture completion is human evidence.
Independent canonical content APPROVE was granted by the same reviewer after
reading the 20 actual proof results, exact checksums and canonical byte equality.
Official task-promote completed draft → in-review → ready after that approval.
Route/address/report work after P18 merge, repository acceptance and exact-head
GitHub validate remain separate gates.
