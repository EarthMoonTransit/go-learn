# I12 journal accepted source and evidence

Approved156-word TASK and original source decision in channel-state-contract-review.json. Official task-init/ADR0008 creation/digest/sources=[]; author precheck is not independent actual acceptance. Earlier join-all/Counter/registry are unchanged preparation, not source donors. Only goroutines.state-ownership and sync.mutex-ownership applications are proposed.

| Requirement | Public assertion | Ordinary wrong specimens | Diagnostic |
| --- | --- | --- | --- |
| empty independent new state | JournalValueHandles | initial-value; shared-constructor-state | new/independent sequence actual/want |
| exactly one append and index | JournalValueHandles mixed/repeated data | omit-append; double-append; reset-on-append; one-based-index; constant-index | alias append index and full sequence actual/want |
| copied handles retain one state | JournalValueHandles copy after use | receiver-detaches-state | original/copy lost update, sequence actual/want |
| detached snapshots from both object and other returned values | DetachedJournalSnapshots | expose-storage; reuse-snapshot-buffer | second changed through first; journal changed; retained snapshot changed |
| non-destructive complete ordered snapshot | two reads in journalSequence | snapshot-drains; reverse-snapshot; omit-last-snapshot | read/sequence actual/want |
| bounded concurrent alias updates/reads | ConcurrentJournalAliases; grader FullBound | finite wrongs above plus index/value/prefix assertions | duplicate index, mismatched stored argument, non-prefix snapshot |

Learner-private journalHandle and journalState contain pointer/mutex/slice allocation and all access logic; public value signatures alone earn no credit. Reference Append locks through len+append and Snapshot locks through len+copy. Handle copies retain the same state pointer; no used mutex is copied. Returned mutable slices are detached. Only constructor handles are in domain. The fixed-array and RWMutex alternatives are correct; nil/empty, field names and helper layout are not prescribed. Public/full passing global-lock and channel-token controls are separate method rejects: lock must belong with each state and be a mutex. No actor goroutine is created by any specimen.

All thirteen required mutants are finite, synchronized ordinary value/alias defects. The receiver-local detach makes a fresh state from a locked snapshot rather than copying a used mutex. Alias writes happen after joined operations, never as intentional racy concurrent writes. Concurrent tests use independently owned result slots and join before inspection; no fixture-held callback/lock and no synctest.Wait on mutex contenders. Exact interleaving is unconstrained. Mandatory reference race supplements, not replaces, actual inspection and finite ordinary feedback. No panic, compile, timeout or DATA RACE is credited as required detection. Actual source/evidence/method acceptance is recorded in independent-review.json; Official task-promote transitions are complete; final batch gates remain required.

Sole actual approval on2026-10-23 binds19transfer and13journal canonical mutants,
48current public/control scopes and32ordinary public diagnostics. All73attempted
roots were removed;25superseded transfer scopes earn no credit. Accepted concrete
method decisions and current hashes are retained in independent-review.json.
