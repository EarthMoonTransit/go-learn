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
| actual greatest common divisor | TestGCDValues coprime/common-factor/equal pairs | single-reduction; wrong-remainder; return-second | GCD arguments and actual/want |
| zero rules | TestGCDValues (0,0), (0,9), (9,0) | zero-first; return-second | GCD arguments and actual/want |
| condition-only loop | TestConditionOnlyLoop | post-clause; recursive | actual loop form or loop count; required form |
| no recursion | TestConditionOnlyLoop | recursive | recursion present=true, want false |
| argument exchange | TestGCDValues both orders | single-reduction; wrong-remainder | GCD arguments and actual/want |

All mandatory mutants terminate. No timeout is treated as useful instructional
feedback. The form assertion covers the declared loop and direct recursion,
without grading variable names or assignment style. The full independent oracle
enumerates 0..64 and adds fixed large inputs within the stated bounds.


All required variants compile; incorrect outputs or declared-form assertions must
fail with the named contract and actual/expected values. Reference public/full
success and mutant failures are measured with disposable task-check execution.
Independent review and repository acceptance remain required before promotion.

The primary skill is not repeated in prerequisites: Task Contract v2 forbids assessed/prerequisite overlap. Planned practice placement after its existing introduction records that exposure; direct taxonomy prerequisites remain explicit. No schema or skill semantics change.

Independent review found that numeric correctness plus a condition-only loop admitted subtraction-only reduction despite the explicit remainder requirement. TestConditionOnlyLoop now checks for % inside the loop body; the terminating subtraction-only variant returns correct numbers but fails with `GCD remainder reduction in loop = false, want true`. The public-only integration matrix installs this variant and requires that diagnostic, plus reference public success, through the disposable learner verifier. This content regression does not change checker behavior.

The public form contract explicitly permits both % and %= in the loop body. A valid compound-assignment alternative is retained in review/tolerated-output.txt and the public/full tolerance integration matrix; assignment spelling is not graded.

Independent content/provenance approval by /root/clamp_e01_content_review on 2026-10-17 covers the final package, including the GCD rejection/tolerance fix where applicable. No further content blockers. Normal lifecycle promotion, final targeted checks, full ordered repository acceptance and exact-head CI remain separate gates.

Measured C1 public acceptance: the reference passes public and full disposable learner verification; every declared wrong variant compiles and fails public tests with an actual/want diagnostic (4, 5, 7 and 6 variants respectively). No real learner home was used. The focused GCD negative and compound-assignment tolerance integrations pass. In a disposable repository copy, disabling only the public remainder assertion makes the negative integration fail on unexpected public success, confirming the regression checks that assertion.
