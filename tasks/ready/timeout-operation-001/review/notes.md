# C3 author review

Original content under the approved S3 source decision, ADR 0008 identity,
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
| derived child, propagated values/cancellation | fast value/deadline and already canceled/gated cases | background-parent; skip-canceled-parent | child value/parent Done/calls actual/want |
| correct relative and earlier-parent budget | exact virtual deadline/expiry | no-timeout; doubled-budget | deadline/virtual elapsed actual/want |
| nonpositive budget already expired | zero/negative callback observations | no-timeout; doubled-budget | nonpositive error/deadline actual/want |
| operation exactly once and result/error unchanged | fast success/error/nonpositive/expiry result and call count | call-twice; zero-result; replace-operation-error; successful-result-canceled | operation calls/result actual/want |
| release on all ordinary returns, parent remains owned | fast success and error return child Done/parent Err | no-release; release-only-on-error; cancel-before-operation | child live after return / canceled too early |
| join cooperative cleanup after Done | owned operation gate, parent cancel before release | return-before-cleanup | returned before cleanup, want pending |

All timing observations are inside synctest.Test. The pinned fake-clock origin
is constructed with time.Date, not an ambient clock. Virtual time.Since and a
six-second fake Timer bound the absent/wrong-deadline defect; it produces a named
actual/expected assertion, never a wall-clock or command-timeout kill. The gated
operation also has an independent owned stop channel so background-parent or
early-return defects can always be released and joined by fixture cleanup.
Deadline races do not choose an arbitrary result winner: the callback owns its
returned result, and the wrapper always joins it. Manual Timer.Stop/Reset is not
learner-owned; its false prerequisite is removed separately by ADR0029.


Independent content/provenance review, normal lifecycle promotion and final
ordered repository acceptance remain required before batch completion.

Public success with nil error now checks unchanged nonzero result, child release
and live parent; successful-result-canceled and release-only-on-error isolate
this repaired public gap. Persisted C3 public-feedback regressions remain required.

The public nil-success case keeps an earlier-deadline parent live and requires
a distinct child with that same earlier deadline. reuse-earlier-parent isolates
false ownership optimization; both child release and unchanged parent state are
checked before the parent expires.

Independent exact-content review approved by /root/clamp_e01_content_review on 2026-10-17. Actual targeted reference, 13/13 required mutants and race passed; accepted alternative solutions also passed where recorded. Final ready-state checks and batch acceptance follow this metadata freeze. Separate persisted public-feedback regressions remain an explicit rollout gate.

Final public-only verification: the reference passed both public and full learner checks; every required compiling wrong variant failed public with FailureCategory=test and useful actual/want diagnostics. The five new C3 packages cover 41 such variants in total. These checks ran only through disposable learner verification; grader-only kills or command timeouts do not substitute for this evidence.
