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
| non-nil output, all occurrences including repeats/zero | finite multiset cases | zero-on-close; skip-second-input | occurrences actual/want; non-nil output |
| preserve each input order | disjoint tagged stream projections | reverse-left-order | stream order actual/want |
| progress from either open input | controlled interleaved availability | serial-inputs; buffer-after-right-close | available input output absent, want progress |
| select/nil mechanism, closed input no longer selected | explicit public construct check plus held remaining input | keep-closed-case | nil disabling missing; one input closed, want pending |
| close only after both inputs | closed-left/live-right phase and finite closure | close-on-first-input | pending/closed actual/want |

Select and nil removal are explicit construction requirements, checked in the
learner's MergeTwo and referenced local helpers in addition to behavior. The check accepts local channel
aliases/ordinary equivalent syntax; it does not claim general program analysis.
The source is independently authored, not an adaptation of an upstream example.
Required variants are finite: the realistic keep-closed-case mistake decrements
its remaining count repeatedly instead of removing a case. Cleanup closes both
owned fixture inputs and drains bounded output; no required defect is detected
only by timeout, spin or uncontrolled scheduling. No fairness/early-abandonment
promise is made. Public observations follow synctest.Wait boundaries.


Independent content/provenance review, normal lifecycle promotion and final
ordered repository acceptance remain required before batch completion.


Review repairs: controlled closure is tested symmetrically, including a finite
buffer-after-right-close defect. The public construction guard follows local
function/method references and accepts parenthesized/typed nil expressions.
review/tolerated-output.txt records private-helper layout and both nil forms;
public/full tolerance acceptance and persisted C3 regressions remain required.

Independent exact-content review approved by /root/clamp_e01_content_review on 2026-10-17. Actual targeted reference, 7/7 required mutants and race passed; accepted alternative solutions also passed where recorded. Final ready-state checks and batch acceptance follow this metadata freeze. Separate persisted public-feedback regressions remain an explicit rollout gate.

Final public-only verification: the reference passed both public and full learner checks; every required compiling wrong variant failed public with FailureCategory=test and useful actual/want diagnostics. The five new C3 packages cover 41 such variants in total. These checks ran only through disposable learner verification; grader-only kills or command timeouts do not substitute for this evidence.
