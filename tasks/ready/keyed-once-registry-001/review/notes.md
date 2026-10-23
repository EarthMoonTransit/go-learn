# Per-key initialization author evidence

The approved I9 contract adds permanent entry identity per key and a compound lookup/install critical section beyond the single captured Once first and supplied Counter increment first. Only sync.once-initialization and sync.mutex-critical-sections applications are proposed. Ownership is a prerequisite/method boundary, without additional application credit.

| Requirement | Public assertion | Finite wrong variant | Useful feedback |
| --- | --- | --- | --- |
| lazy constructor | TestRegistryLazyConstruction and constructRegistry | eager-load | construction loader calls actual/want; constructor invoked loader before Get |
| exact key and one permanent entry per key | held/concurrent key result/counts, sequential lifetimes | wrong-key; new-entry-per-call; one-entry-all-keys | key actual/wanted complete pair or loader count |
| complete cached value/error, including failure/zero | held A nonzero error, B zero success, late calls, sequential bad key | drop-value; drop-error; retry-failure | complete pair/late result/count actual/want exact identity |
| wait for initialization publication | first Get observed alone while callback gate held | return-before-load | first Get returned before loader completed; no fabricated race |
| independent registries | TestRegistryInstanceIndependence same key/different loaders | shared-registry-state | independent results/counts actual/want |
| atomic permanent entry selection and lock/Once ownership | actual source plus finite concurrent counts/race | wrong entry lifecycle plus method controls | behavior is not proof of full lock scope |

The fixed public declaration `type Registry registryState` leaves the private registryState body entirely learner-owned and keeps the existing API checker stable. The starter supplies no mutex/map/Once/entry fields or initialization logic. Reference owns a registry mutex and map[string]*entry, atomically selects/installs a stable pointer, unlocks before Do/load and returns fields only after Once completion. All map reads/writes follow that mutex. No reset, copy, eviction, retry, panic or reentry feature is present. The alternative uses a private locked lookup helper, lazy map and result pair; its legitimate layout differs.

Safe test phases: constructor and first Get alone are observed with synctest.Wait before any contender exists. An eager constructor is released by fixture rescue. Additional callers acknowledge before Get; all callback gates (including fallback unknown-key gate) are closed unconditionally before the next Wait. Those acknowledgements do not prove actual contention inside Once. Exact per-key pairs/counts and race checking assess finite publication behavior; cross-key lock scope is inspected as method, not claimed as B-progress-while-A-held proof. No Wait holds a fixture mutex while workers contend; counters/logs are synchronized for the race-safe asynchronous defect. All async loads finish under bubble ownership after release.

The explicit-ready per-entry mutex cache has correct behavior but no sync.Once. The global-registry-lock control holds the lock through Do/load; its callbacks are independently released so it finishes and may pass behavior. Both are intentional method-wrong controls, not timeout/deadlock credit. A copied-Once or inaccessible lock-deadlock mutant is not required evidence. No Sleep/Now/After, timers, busy loops or custom runner.

Canonical and public-only fixtures must detect all nine ordinary wrong variants with useful messages. Correct alternative and both method controls must pass public/full; actual independent method review remains required. Original creation/provenance retained from official task-init; authored execution only via disposable task-check.

Sole actual source/method/evidence review approved2026-10-23 without findings; both Once and mutex-critical-section applications accepted as content. Officialpromotion/finalbatchdelivery separate.
