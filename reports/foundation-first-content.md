# P5 / C17 — one executable LRU draft; no ready promotion

The six-slot maximum produced one authored substantive adaptation:
`tasks/drafts/bounded-lru-cache-001`. SHAD's four pinned lrucache artifacts
supply the interface and cross-operation behavior; provenance records exact
hashes, complete MIT notice and the local deterministic adaptation boundary.
No originals or canonical Exercism cases were introduced.

Public tests cover independent/empty/zero-capacity instances, arbitrary signed
integers, lookup presence, Get/Set recency, least-recent eviction, ordered early
termination and Clear/reuse. An independent logical-clock/map oracle extends
mixed traces. A linear slice reference is deliberately valid; no complexity,
allocation, machine timing or random discovery threshold is imposed.

Targeted task-check passes reference and detects 14/14 compiling defects. A
public-only probe through the existing isolation runner also rejects all 14 with
useful diagnostics and accepts the reference. Tests never ran in the canonical
task directory. Exact content digest is in provenance.yaml. Full repository
acceptance uses one saved collection and remains a separate PR gate.

## Independent review outcome

`/root/review_intake` verified all five source hashes and MIT notice, reviewed the
contract/reference/dedup and requested changes on instructional assessment.
The provisional primary interfaces.implicit-satisfaction is not actually tested:
the compiling starter supplies every interface method, and the defects preserve
satisfaction. Replacing the label with receiver semantics is also unsupported by
this scaffold. A found public zero-key/negative-value gap was corrected with two
explicit mutants and public assertions. The assessment issue remains recorded in
review/review.yaml as request-changes.

The task therefore remains draft. It contributes no ready count, assessed skill
coverage, route address or learner evidence. Source permission does not guarantee
that an exercise fits the current assessment contract. P6 should prepare a bounded
assessment decision before promotion, not silently change taxonomy or manufacture
an unrelated task to fill the quota. No learner solutions or progress were read.

W1 exact ParseDuration lineage is unresolved; W2 adds little distinct practice;
W3/W4 canonical lineage remains restricted; W5 supplies no distinct nil
representation contract. Externalsort, treeiter, dupcall and subsequent concurrency
components remain outside this content batch. P3/PR60 and P2/PR58 remain merge
prerequisites; GitHub's billing failure cannot be bypassed by local evidence.
