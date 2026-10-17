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
| callback exactly once | TestAdjustedResultAndCaller; TestPendingPanicIsPreserved | double-work | work calls actual/want |
| adjust after explicit return assigns named result | TestAdjustedResultAndCaller | eager-adjustment; overwrite-result; local-shadow | Adjust arguments and actual/want result |
| unchanged caller in success and panic | both behavioral tests | caller-mutated | caller actual/want value |
| preserve pending panic identity | TestPendingPanicIsPreserved | recover-panic | actual/want original panic |
| named result and deferred mechanism | TestNamedDeferredResultForm | manual-adjustment | named presence/defer count actual/want |

The starter already supplies the deferred function and callback call. The learner
repairs its target; no panic recovery implementation is requested. A pending
panic is injected by checker-owned work. Inputs bound arithmetic safely.


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
full scopes. All 7 required compiling wrong variants failed public Check with
FailureCategory=test and useful actual/want diagnostics; no compile error,
timeout or hidden grader detection substitutes for this observation. Execution
used temporary learner payloads and fresh disposable checker copies, never the
authored source checkout or a real learner home. The one-off batch record is
`.tmp/pedagogical-c2/public-feedback.jsonl`. Persisted public-feedback harness
rows belong to the subsequent separate tooling batch, per docs/DEBUGGING.md;
that follow-up remains required before claiming the full rollout complete.
