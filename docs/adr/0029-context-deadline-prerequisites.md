# ADR 0029: Deadline composition does not require manual timer ownership

- Status: accepted for a separate taxonomy preparation batch
- Date: 2026-10-17
- Baseline: master `d71294257d1b8f43c62f5d5f77c63282827bd98f`
- Planning evidence: supplied review and S3 source decision, PR #76

## Decision and reason

Remove only the active taxonomy prerequisite edge
`context.deadlines -> time.timers-tickers`. Keep the deadline skill's
`context.cancellation` prerequisite and both timeout prerequisites
`context.deadlines` and `time.duration-arithmetic`. No skill ID, definition,
minimum, evidence requirement, schema field or field meaning changes.

P16 introduces timeout budgets and absolute deadline composition with the parent
budget preserved. A caller derives a context and owns its CancelFunc; that API
releases the associated timer resources. The caller need not implement Timer
Stop/Reset lifecycle, which is the distinct outcome of time.timers-tickers.
Requiring that lower-level mechanism before using context timeouts is a false
exposure dependency. The supplied review places timers/tickers after P16, rather
than asking for retrospective credit or an additional manual-timer introduction.

Semantic evidence is registry golang-go, pinned commit
`2c31595783c74564c1299a5401dc6e606f379cc7`, exact
`src/context/context.go`, public WithDeadline/WithTimeout/CancelFunc contract,
with the hash verified in
`catalog/intake/2026-10-17/pedagogical-c3-source-review.json`. Facts only; no
upstream exercise, implementation, tests or example is copied. Independent
read-only prerequisite review by `/root/clamp_compatibility_review` confirmed
this narrow correction and the supplied plan's ordering.

## Migration and rollback

This changes the active graph, leaving immutable v2 snapshots byte-identical.
Existing task manifests, authored content, primary/secondary skills and route
entries are unchanged. None of the 171 baseline packages assesses or declares
these three timer/deadline/timeout skills; therefore no task loses a necessary
introduction and the current route needs no reordering. Future P16 keeps both
its planned assessments and actual cancellation/duration prerequisites.

Do not add timers familiarity, secondary tags, encounters, completions or mastery
to old progress. Workspace content digests, addresses, learner files/markers,
assistance and review clocks remain unchanged. There is no user-data migration.
Rollback restores this one prerequisite edge; once P16 is activated it would
require a new preparation decision before that stricter order could be used.

Taxonomy bytes enter every TaskInputDigest. The previous collection is stale
for this branch and cannot certify the new graph. Bootstrap aggregate input
metadata from unchanged task records only, without fabricating a collection or
claiming new runtime results. Then run ordered fast/race/vet/integration and one
new full task-check collection. The official report must reproduce bootstrap
bytes; validator and report-check must use that same new collection. Independent
exact-head review and green validate are required before merge.

This is a preparation taxonomy batch under ADR 0022. Manual timers/tickers remain
future content; no complete-coverage, foundation release or public export is
claimed. C3 content and T0b activation remain separate batches.
