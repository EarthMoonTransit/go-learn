# C2 author review

Original content under the approved S2 source decision, ADR 0008 identity,
creation record, actual digest, empty external task sources and deferred release
binding. No upstream exercise body, code, test, solution or example was used.
Estimates are not learner-calibrated. Public/full execution uses disposable
task-check and learner verification only; no real learner home is used.

Assessed primary/secondary are not repeated in prerequisites, as Task Contract v2
forbids overlap. Explicit practice placement and exposure edges provide the
already-introduced assessed mechanism without schema or taxonomy changes.

## Requirements → public assertions → wrong variants → diagnostics

| Requirement | Public assertion | Wrong variant | Diagnostic |
| --- | --- | --- | --- |
| equal composite keys share an entry and replace | TestCompositeEqualityAndReplacement | no-replacement; pointer-keys | equal/replacement actual/want string and presence |
| each array component and Class participates | TestCompositeEqualityAndReplacement | missing-first-part; missing-second-part; missing-class | changed component collision, actual/want |
| insertion copies caller's key | TestInsertionCopiesCallerKey | pointer-keys | original/mutated key actual/want after mutation |
| present empty versus absent | TestPresenceAndIndependentTables | empty-means-absent | actual/want string and presence |
| lookup does not mutate storage | repeated lookup in TestPresenceAndIndependentTables | destructive-lookup | second present-empty actual/want |
| independent empty tables | TestPresenceAndIndependentTables | shared-table | independent table actual/want |

The private storage type permits a compiling pointer-keyed defect without changing
the exported Key or function signatures. Its behavior loses the inserted value key
after caller mutation. No interface, method, closure or generic implementation is
required. Pointers are an explicit prerequisite because Put observes a caller-owned
Key through *Key; placement remains after the relevant C06 preparations.


Independent content/provenance review, normal lifecycle promotion and final
ordered repository acceptance remain required before batch completion.


## Independent content approval

`/root/clamp_e01_content_review` approved this exact content digest on
2026-10-17 after read-only TASK/API/public/grader/mutant/hint/prerequisite and
provenance review. Formatter approval is limited to the explicit P05 guided
introduction; unaided method-set design and independent mastery are not claimed.
P08 remains conditional on reviewed M1/placement before activation. Public/full
execution, normal promotion and ordered repository acceptance follow separately.


## Disposable public/full evidence

The existing learner verifier accepted the reference through both public and
full scopes. All 8 required compiling wrong variants failed public Check with
FailureCategory=test and useful actual/want diagnostics; no compile error,
timeout or hidden grader detection substitutes for this observation. Execution
used temporary learner payloads and fresh disposable checker copies, never the
authored source checkout or a real learner home. The one-off batch record is
`.tmp/pedagogical-c2/public-feedback.jsonl`. Persisted public-feedback harness
rows belong to the subsequent separate tooling batch, per docs/DEBUGGING.md;
that follow-up remains required before claiming the full rollout complete.
