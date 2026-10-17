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
| deferred first calls and immediate third/checkpoint | TestRegistrationAndLIFO | eager-calls; defer-third | actual/want trace with checkpoint position |
| registration-time argument values | TestRegistrationAndLIFO | late-capture | actual/want saved values |
| LIFO order | TestRegistrationAndLIFO | reverse-registration | actual/want trace ordering |
| exactly four entries | TestRegistrationAndLIFO | missing-first | actual/want whole trace |
| declared defer mechanism in existing boundary | TestDeferredCallsRemainInFunction | manual-replay | actual/want deferred-call count |
| independent calls and varied inputs | TestRegistrationAndLIFO sequential seeds | eager-calls; late-capture | seed and actual/want trace |

The supplied immutable wrapper owns the slice, append and checkpoint. The learner
only changes the supplied runTrace body; no recorder closure or counter authoring
is required. The form assertion checks the expressly required defer mechanism,
without choosing the learner's local variable names or closure style.


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
full scopes. All 6 required compiling wrong variants failed public Check with
FailureCategory=test and useful actual/want diagnostics; no compile error,
timeout or hidden grader detection substitutes for this observation. Execution
used temporary learner payloads and fresh disposable checker copies, never the
authored source checkout or a real learner home. The one-off batch record is
`.tmp/pedagogical-c2/public-feedback.jsonl`. Persisted public-feedback harness
rows belong to the subsequent separate tooling batch, per docs/DEBUGGING.md;
that follow-up remains required before claiming the full rollout complete.
