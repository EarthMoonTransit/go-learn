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
| asynchronous return | startHeld controlled blocked phase | inline-runner | blocked, want return before Release |
| exactly one supplied runner | held/released Calls | skip-runner; duplicate-runner | actual calls, want 1 |
| return exact supplied handle | handle equality | wrong-handle; shared-job | returned supplied handle false, want true |
| pending before release, closed afterward | Done phase probes | premature-release; skip-runner | completed before Release, want pending; pending, want closed |
| independent jobs and release | two held handles, sequential releases | shared-job | per-job calls/identity and held state |

All channels, sync.Once and atomic counts are supplied immutable scaffolding,
not learner-implemented mechanisms or independent mastery evidence. Cleanup
releases all runners and waits at a synctest boundary, including inline and
skipped/duplicate defects; no required mutant is detected by timeout alone.


Independent content/provenance review, normal lifecycle promotion and final
ordered repository acceptance remain required before batch completion.


Independent exact-content/provenance approval by `/root/clamp_e01_content_review`,
2026-10-17. Six required variants, race/reference and starter reports passed;
public/full learner evidence and final ordered batch acceptance follow.

Final public-only verification: the reference passed both public and full learner checks; every required compiling wrong variant failed public with FailureCategory=test and useful actual/want diagnostics. The five new C3 packages cover 41 such variants in total. These checks ran only through disposable learner verification; grader-only kills or command timeouts do not substitute for this evidence.
