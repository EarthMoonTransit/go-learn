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
| live forwarding before EOF | TestForwardProgressBeforeInputClose | collect-until-eof | live input output pending, want value before EOF |
| no retained receiver after cancellation | unbuffered sentinel-send probe before cleanup | retain-input-receiver | stage still receives caller input, want no retained receiver |
| order and exact occurrences | finite EOF sequence with zero/duplicates | drop-zero; reverse-input | forwarded order/occurrences actual/want |
| EOF and cancellation close output | EOF plus three cancellation phases | leave-output-open; background-context | pending, want closed |
| cancel while waiting input | waiting-input, already canceled parent | missing-receive-cancel | pending, want closed |
| cancel at output backpressure | capacity-aware full output plus one pending send | missing-send-cancel | extra output/pending, want closed |
| caller retains input | ownership probe after canceled stage | close-caller-input | caller input closed, want open |

Input is deliberately bidirectional in this ownership exercise, allowing a
compiling close-caller-input defect; the output remains receive-only. Parent
cancellation authority is not passed to the stage. The output-backpressure case
adapts to cap(out), accepts values queued before cancellation and does not impose
zero-post-cancel priority. Every producer has an independent cleanup stop signal;
cleanup cancels, joins that producer, closes only owned input and drains the
finite remaining output. Required defects have named assertion failures before
cleanup and do not rely on timeout, leaked workers or ambient scheduling.


Independent content/provenance review, normal lifecycle promotion and final
ordered repository acceptance remain required before batch completion.

A saved tolerated-output fixture uses capacity two, preserving the explicitly
unspecified output capacity; acceptance must pass public/full for that fixture.

Independent review found and repaired a retained-receiver false accept: closure
of output alone is not termination. A sentinel send probes the single unbuffered
caller input after cancellation/Wait and before cleanup closes it. The dedicated
finite mutant closes output then drains input; the assertion fails first and
cleanup subsequently releases that exact wrong receiver. A live open-input phase
also rejects cancelable collect-until-EOF, without imposing an output capacity.
Persisted tooling rows/removal probes for these C3 repairs remain required.

Independent exact-content review approved by /root/clamp_e01_content_review on 2026-10-17. Actual targeted reference, 9/9 required mutants and race passed; accepted alternative solutions also passed where recorded. Final ready-state checks and batch acceptance follow this metadata freeze. Separate persisted public-feedback regressions remain an explicit rollout gate.

Final public-only verification: the reference passed both public and full learner checks; every required compiling wrong variant failed public with FailureCategory=test and useful actual/want diagnostics. The five new C3 packages cover 41 such variants in total. These checks ran only through disposable learner verification; grader-only kills or command timeouts do not substitute for this evidence.
