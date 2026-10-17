# ADR 0030: active pedagogical order with stable lesson addresses

- Status: accepted for T0b implementation, subject to independent review and CI
- Date: 2026-10-17
- Phase: preparation/tooling under ADR 0022, not complete foundation coverage

## Decision and migration

Route v2 remains the strict address registry: preserve all 151 original task ID,
block ID, block number, lesson number and role tuples. Append 15 addresses for
actually ready C1/C2/C3 packages; proposal IDs never become task IDs. Preserve the
original review CSV bytes. A new explicit v2 mapping extends that sequence with
the reviewed preparations; it is not generated from current learner progress.

The existing pedagogical sidecar gains schema version 2 with explicit status
active or inactive. Version 1 remains inactive-only. Active loading requires
structural validity, every mapped ready v2 package, matching assessed skills,
contiguous explicit positions, all actual manifest/taxonomy prerequisites and
reviewed exposure edges, and no pending preparation. Invalid active configuration
fails closed for scheduler selection; the read-only preview still explains its
structural/preparation errors. Explicit stable-address lookup, workspace resume
and IDE targets remain readable independently of scheduler-plan readiness. Missing/inactive sidecars preserve address-order execution.

Separate structural address validation from exposure validation, and validate
exposure in the selected execution order. Pass that validated sequence into the
existing route scheduler: unfinished current task, then eligible delayed review,
then ordinary next. Addresses, lookup, deep links, GoLand targets and numbered or
legacy workspace paths continue to use route v2. No second scheduler, progress
schema upgrade, workspace rewrite or learner-home migration is introduced.

Delayed review readiness requires introduced actual manifest prerequisites,
assessed primary/secondary mechanisms and declared new concepts. It does not
require mastery or previous completion of the selected practice. Missing
preparation, including PointerChain mutation-sharing, postpones that candidate;
review remains due and ordinary progression continues. Show a concrete selection
reason in CLI and UI. Keep the interval of three distinct other completions,
mastery from two distinct independent practice/transfer IDs, durable assistance,
and reset-without-new-evidence exactly as ADR 0017 defines them.

Read introduced mechanisms from stored completion evidence, without retagging it.
Two narrowly reviewed compatibility equivalences affect readiness only:
completed Clamp establishes basic Test API familiarity for the exact current
ADR 0026 content, and completed ValueAverage establishes explicit-receiver method
expression familiarity for unchanged content digest
`5b7a7b79dd6b2d355ad5901c0f81853b26d6d52cfbe7d8b5f59b708e8a012b68`.
The latter content is unchanged since its first ready commit
`3e7f2671ef8892245308f074df8ae5bea45b71cf`. Started/public-passing tasks do not
qualify. Neither equivalence writes CompletionEvidence, SkillEncounters,
assistance, clocks, attempts or mastery counts. There is no generic historical
secondary retagging. Supplied scaffolding is not independent harness/API design.

Apply the reviewed DetachedAppend variadic prerequisite and the bounded
ValueAverage methods.values-expressions secondary consistently with its route
step. Their authored content and workspace content digests stay unchanged.
Existing roles and primary skills stay unchanged. No taxonomy ID, schema field
meaning or skill definition changes.

## Acceptance and rollback

Exercise fresh, legacy, paused/current, assisted/reset, completed old-route,
postponed and exhausted-review states in disposable learner homes. Compare exact
progress/history/timer/workspace bytes, stable lookup and IDE targets. Audit all
18 near-term review changes and actual placement, plus separately merged source,
content and C3 public-feedback gates. Later proposals, M2 and T-FUZZ/T-BENCH/
T-MODULE remain explicitly deferred.

Run ordered repository fast/race/vet/integration gates, one full collection for
the final task inputs, and official report/validator checks against that same
collection. Merge only after independent exact-head review and green validate.

Rollback reverts the activation/metadata batch together. Reverting only status
would restore address-order prerequisite validation and correctly reject the
DetachedAppend variadic dependency at its old 02.07 address. Never delete history,
renumber old addresses or move workspaces to make rollback appear valid. Tasks
completed after activation retain their IDs, evidence and existing workspace paths.
