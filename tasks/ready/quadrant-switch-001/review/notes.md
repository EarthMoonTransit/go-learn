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
| origin wins overlapping boundaries | TestNineRegions (0,0) | axis-priority | Locate arguments, actual/want token; boundary priority |
| signed horizontal/vertical axes | TestNineRegions four axes | axis-sign; axis-priority | Locate arguments and actual/want token |
| four quadrant tokens | TestNineRegions four quadrants | swapped-quadrant; missing-default | Locate arguments and actual/want token |
| expressionless switch | TestExpressionlessSwitch | tagged-switch; starter | tag present=true or count=0, required form |
| no fallthrough | TestExpressionlessSwitch; TestNineRegions origin | fallthrough-origin | fallthrough present=true, want false; origin token actual/want |

The fixed matrix in the full suite is independent of the reference dispatch and
includes int extremes, so overflowing arithmetic sign tests are not required.
Only the explicitly declared switch/no-fallthrough form is graded. Public AST
code remains checker-owned; the learner need not know AST APIs.


All required variants compile; incorrect outputs or declared-form assertions must
fail with the named contract and actual/expected values. Reference public/full
success and mutant failures are measured with disposable task-check execution.
Independent review and repository acceptance remain required before promotion.

The primary skill is not repeated in prerequisites: Task Contract v2 forbids assessed/prerequisite overlap. Planned practice placement after its existing introduction records that exposure; direct taxonomy prerequisites remain explicit. No schema or skill semantics change.

Independent content/provenance approval by /root/clamp_e01_content_review on 2026-10-17 covers the final package, including the GCD rejection/tolerance fix where applicable. No further content blockers. Normal lifecycle promotion, final targeted checks, full ordered repository acceptance and exact-head CI remain separate gates.

Measured C1 public acceptance: the reference passes public and full disposable learner verification; every declared wrong variant compiles and fails public tests with an actual/want diagnostic (4, 5, 7 and 6 variants respectively). No real learner home was used. The focused GCD negative and compound-assignment tolerance integrations pass. In a disposable repository copy, disabling only the public remainder assertion makes the negative integration fail on unexpected public success, confirming the regression checks that assertion.
