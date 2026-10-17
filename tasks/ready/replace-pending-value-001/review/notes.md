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
| receive before send | full-buffer synctest boundary and cleanup drain | send-before-receive | blocked, want receive before send |
| return original integer | zero/negative/positive/equal pairs | return-next; return-zero | returned actual, want old |
| next remains pending exactly once | pending receive and extra-value probe | drop-replacement; restore-old | absent/wrong pending, want next |
| caller channel remains open | post-drain nonblocking receive | close-channel | closed, want open |

The public harness owns all asynchronous observation and bounded cleanup.
The send-first mutant is released by draining the original pending value and
joined at synctest.Wait; its named failure is not a command timeout. Concurrency
and race checks run only in disposable workspaces. Nil, closed, empty and larger
capacity channels are outside the input domain; no extra select design is asked.


Independent content/provenance review, normal lifecycle promotion and final
ordered repository acceptance remain required before batch completion.


Independent exact-content/provenance approval by `/root/clamp_e01_content_review`,
2026-10-17. Six required variants, race/reference and starter reports passed;
public/full learner evidence and final ordered batch acceptance follow.

Final public-only verification: the reference passed both public and full learner checks; every required compiling wrong variant failed public with FailureCategory=test and useful actual/want diagnostics. The five new C3 packages cover 41 such variants in total. These checks ran only through disposable learner verification; grader-only kills or command timeouts do not substitute for this evidence.
