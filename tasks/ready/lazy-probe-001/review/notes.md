# C1 author review

This original preparation task follows the approved, bounded S1 source decision.
No external task content was consulted or adapted. The technical contract differs
from the existing ready tasks as recorded in the S1 report. Project identity,
creation record, empty external sources and deferred release binding follow ADR 0008.
Only subject/implementation.go is editable; subject/api.go is supplied immutable
scaffolding. Tests are checker-owned. No learner home or authored source checkout
is executed. Estimates have not been calibrated with learners.

## Requirements → public assertions → wrong variants → diagnostics

| Requirement | Public assertion | Wrong variant | Diagnostic |
| --- | --- | --- | --- |
| disabled means false and zero effects | TestProbeGate disabled/false, disabled/true | eager-call; inverted-gate | actual/want result and Probe calls; eager-call explanation |
| enabled invokes exactly once | TestProbeGate enabled cases | double-call; never-probe | calls actual/want; exactly once |
| preserve both possible probe results | TestProbeGate enabled/false, enabled/true | discard-result; never-probe | configured Probe result, actual/want return |
| independent sequential observations | TestProbeGate resets for every case | eager-call; inverted-gate | per-case result and call count |

The learner uses the supplied Probe function, without writing a closure,
interface, counter or testing fixture. No syntax beyond the declared behavior is
graded: an equivalent if-based solution is accepted. Hidden checks only compose
the same four boolean cases into repeated calls.


All required variants compile; incorrect outputs or declared-form assertions must
fail with the named contract and actual/expected values. Reference public/full
success and mutant failures are measured with disposable task-check execution.
Independent review and repository acceptance remain required before promotion.

The primary skill is not repeated in prerequisites: Task Contract v2 forbids assessed/prerequisite overlap. Planned practice placement after its existing introduction records that exposure; direct taxonomy prerequisites remain explicit. No schema or skill semantics change.

Independent content/provenance approval by /root/clamp_e01_content_review on 2026-10-17 covers the final package, including the GCD rejection/tolerance fix where applicable. No further content blockers. Normal lifecycle promotion, final targeted checks, full ordered repository acceptance and exact-head CI remain separate gates.

Measured C1 public acceptance: the reference passes public and full disposable learner verification; every declared wrong variant compiles and fails public tests with an actual/want diagnostic (4, 5, 7 and 6 variants respectively). No real learner home was used. The focused GCD negative and compound-assignment tolerance integrations pass. In a disposable repository copy, disabling only the public remainder assertion makes the negative integration fail on unexpected public success, confirming the regression checks that assertion.
