# Finite-v1 introduction batch I4 — in progress

Started2026-10-22T18:10:42UTC on `codex/v1-introductions-04`, base
`e9e9ca854de593047e2790565a83c53397d21ec9`. Owner authorized parallel authoring
while I3 CI runs. Dependency PR108 passed all gates and merged as `0aeef21b1dd6dd28ebf97ac59614b63e5715ea47`; PR107 also passed exact-head validate39828282621 and merged as `2efeab7a05596a6591a35e649c332ba8e14153ac`;
Actual method/content review is accepted. I4 still needs one stable collection, all gates,
independent final review and exact-head green validate. No readiness credit yet.

Six units: four original automatic drafts (bounded-job-pool, recover-worker,
pipe-exchange, idle-stream) and two unchanged reuses (cancelable-forward,
indexed-job-results). Contract/source plan approved by sole reviewer/root/reviewer.
Two authors work on separate task files; controller owns route/reuse/integration.

Five original firsts are planned. Result/error collection and WG applications in
the pool and panic policy application in Launch require separate actual evidence.
WG first repair asks for actual Add/Done/Wait in indexed-results; original valid
non-WG automatic solutions retain their validity. Cancellation first is bounded
to one stage. Timer first does not close the later ticker application.

Existing task bytes, provenance, addresses, route role/skills and relative order
must remain unchanged. Manual method supplements do not add automatic coverage
or learner progress. New sources use ADR0008 originals, creation/digest and empty
external sources; exact reuse paths/hashes are in the source-review record.

Four-hour checkpoint starts18:10:42UTC; all active author/review/check work counts,
overlap once. I3 overlap is reported explicitly. Infrastructure fixes, if any,
retain the per-batch and three-batch limits. No new tooling is planned for I4.

## Targeted author evidence

Four original packages were checked through task-check in disposable copies.
Canonical and public-only reference, repeated observations, vet and race pass.
Public wrong-reference probes retain explicit assertion diagnostics: pool9,
worker recovery12, Pipe9, timer9 (39 total). No compile error, process panic,
race or timeout is credited as a behavioral defect detection. Correct reversed
pool results and Pipe wrappers hiding io.Closer pass; these alternatives remain
within the contract. Sole independent reviewer accepted actual content/method/provenance and all39 diagnostic observations on2026-10-22.

| Unit | New assessed action | Evidence boundary |
| --- | --- | --- |
| cancelable-forward reuse | Cancellation at both blocking operations and downstream abandonment | One-stage first, unchanged9-mutant automatic scope; whole-chain application open |
| bounded-job-pool | Fixed worker pool over shared queue and merged indexed outcomes | Actual queue/worker/WG method reviewed separately; output order unspecified |
| recover-worker | Worker-local recovery publishes one deliberate outcome then closes | String/int panic contract; all12 mutants retain recover; no crash-only credit |
| pipe-exchange | Synchronous producer/consumer close causes and producer join | Actual Pipe and ownership inspected; correct endpoint wrappers accepted |
| idle-stream | Initial and reset quiet deadline with owned Timer | Stop is method evidence; no co-ready priority or timer-leak proof; ticker application open |
| indexed-job-results reuse | Actual Add-before-launch/Done/Wait first assessment | Existing13-mutant automatic contract and non-WG alternatives remain unchanged |

Sole independent reviewer accepted both unchanged reuse methods and exact49pinned
files on2026-10-22. Pool's new WaitGroup application changes registration from
jobs to persistent workers and completion from fixed slots to safe closure of a
merged channel; the result/error application preserves every indexed outcome in
that unordered stream. Worker recovery changes the earlier synchronous boundary
into an asynchronous failure policy. These three applications were accepted from the actual source; complete batch
gates and merge remain required.

Proposed addresses are09.22–09.25. Old199 route objects and their relative
pedagogical order must remain exact; source-path hashes preserve the two reuses,
including the adapted Ardan provenance chain on indexed-job-results. The four
original packages have no external source. All8 official lifecycle transitions passed. Route integration preserves all199
old objects and relative order; all4146 old task files match the pinned base.
The four new addresses are09.22–09.25; route now has203 units. One stable213-task
collection passed213/213 in19m07s (18:32:00–18:51:07UTC) and generated task-quality
and both numeric reports. Fast/race/vet/integration and final consumer checks are
running on stable inputs; their results and final CI/merge remain required.

## Stable collection and delivery boundary

Automatic inventory is206ready,213checked; route203 and cumulative5manual units.
After accepted I4 the original first-slot remainder is5, independent applications71
and components5, with explicit semantic repairs retained. The old49pair audit is
unchanged; WG first repair is separate and now has actual source/method acceptance.
Numeric full minima still fail (primary206/418, deficit219; total236/708,
deficit472); foundation minima still fail (206/410, deficit211;236/692,
deficit456). Manual methods do not inflate these numbers. Human mastery is unknown.

Sole reviewer accepted the contract, actual content/methods and39public diagnostic
observations. Final exact-HEAD diff review and all local/CI results must precede
merge. I5 proceeds in a separate worktree during checks; this does not waive I4
acceptance. No infrastructure work was added in I4. Four-hour clock18:10:42UTC
continues through local gates/review and productive overlap, with no CI-only idle
exclusion. This report records the state at content freeze; final acceptance is
recorded in the PR and next postmerge checkpoint.
