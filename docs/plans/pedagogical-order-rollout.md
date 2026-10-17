# Pedagogical order rollout

This executes the near-term S1/C1, S2/C2, S3/C3 and T0b plan from the supplied
2026-10-11 review at baseline
`b9790106edc3d4c7f57e26d5f45dc1960f022c6c`. Proposal IDs are planning labels,
never task IDs or learner evidence. Source admission, content review, current
quality gates and exact-head green CI are required before each merge.

The goal is a prepared, active pedagogical order with preserved old addresses,
workspaces and history. It does not declare complete foundation coverage or a
public export. The review's explicitly deferred Later proposals, M2 specialist
policy and T-FUZZ/T-BENCH/T-MODULE stay subsequent work, as specified in that plan.

## Completed

- T0a read-only preview, validation and compatibility fixtures: PR #66.
- Test execution tooling with race, integration, grader and isolation retained:
  PR #67.
- Exact Clamp content-revision compatibility: PR #68.
- E01/M1 Clamp starter/hints and coordinated secondary: PR #69, merged at
  `49c6202c1c09f7462a98746a0ace883192a24b49`.
- M1 actual PointerChain/DiagnosticHelper/NamedCases prerequisites: PR #70,
  merged at `63a0a8f47573959c7b307c9e50fbf504d013ac34`.

- E02 exact old-workspace compatibility: PR #71, merged at
  `e80a1de34264aef98d646e0e8654ba2f077fec0b`.
- S1 bounded source decisions: PR #72, merged at
  `18a6b7cedb19b06fd0dcf59ec1a2944f113188d3`.

## Remaining execution checklist

- [x] M1: PointerChain mutation-sharing prerequisite; DiagnosticHelper interface
  and variadic prerequisites; NamedCases function-value/struct/slice prerequisites.
- [ ] M1/T0b: DetachedAppend variadic prerequisite. Its 02.07 stable address
  precedes Prepend 03.10; apply with execution-order separation, preserving the
  current route until activation. Review the conditional ValueAverage secondary
  explicitly; it is not an automatic tag addition.
- [x] S1: bounded source decisions for E02 and P01–P04. Reuse E01's admitted
  lineage; record exact artifacts or specific original reasons before authoring.
  Independent source review approved the decisions recorded in
  `reports/pedagogical-c1-source-review.md`; future content acceptance is separate.
- [x] C1/E02: both public variadic call examples in Prepend, with old-workspace
  compatibility before changing authored content.
- [x] C1/P01: scope exercise. Ready package: `conditional-binding-001`.
- [x] C1/P02: short-circuit effects exercise. Ready package: `lazy-probe-001`.
- [x] C1/P03: GCD loop exercise. Ready package: `gcd-reduction-001`.
- [x] C1/P04: expressionless switch exercise. Ready package: `quadrant-switch-001`.
- [x] S2: Formatter concept-seed provenance, Resource original, exact Ardan
  method artifact and bounded defer/key source decisions; PR #74, merged at
  `49fbd60b7aec9d2f5d7fdf4611afb7c0be8bfbfd`.
- [x] C2/P05: promote existing `formatter-contract-001` before Shape.
- [x] C2/P06: defer registration, argument evaluation and LIFO exercise.
- [x] C2/P07: named result modified by defer before Capture.
- [x] C2/P08: method capture exercise.
- [x] C2/P09: comparable map-key exercise.
- [x] C2/P10: promote existing `resource-finalize-001`.
- [x] Tooling: persist C2 public-feedback regressions and verify assertion-removal
  probes in a disposable repository copy: PR #77.
- [x] S3: bounded concurrency source decisions, original reasons and existing
  Worker lineage review: PR #76.
- [x] C3/P11: launch with supplied handle.
- [x] C3/P12: receive then send on an open full single-slot channel.
- [x] C3/P13: failure-safe controlled Worker scaffold, preserving the task ID
  and old workspaces.
- [x] C3/P14: finite channel merge.
- [x] C3/P15: cancelable forwarding.
- [x] C3/P16: timeout/context exercise.
- [ ] Tooling: persist C3 public-only feedback and tolerated-alternative regressions;
  verify omitted-assertion false accepts fail their outer regression in disposable
  copies (held launch, remaining merge stream, retained receiver/live forwarding,
  nil-success result/release and distinct-child ownership).
- [ ] Add append-only addresses for actually ready new/promoted packages and
  explicitly extend the pedagogical mapping. Preserve every original tuple.
- [ ] T0b ADR, existing scheduler activation, readiness filter and concrete
  review explanations. Current unfinished task retains priority; due eligible
  review precedes ordinary next; interval 3 and two distinct independent
  practice/transfer mastery evidence remain unchanged.
- [ ] T0b regression audit: old Clamp familiarity without retagging, legacy
  progress, numbered/legacy workspaces, paused/current tasks, postponed and
  exhausted reviews, assistance, reset and fully completed old route.
- [ ] Final audit against all items, merged exact-head CI evidence and the
  actual active route; retain incomplete Later scope explicitly.

Keep tooling, source decisions, metadata and bounded content batches reviewable.
No content batch exceeds ten new or substantially changed tasks. Use disposable
execution only, one full collection per unchanged batch inputs, and reuse it in
validator/report checks. New preparation tasks are not additional evidence for
already completed tasks; no learner home is used in acceptance.

C1 releases four reviewed original packages and the exact E02 revision. Their
append-only addresses and explicit insertion into the pedagogical mapping remain
the separate unchecked placement step above; the original 151-row mapping and
all stable address tuples are unchanged in this content batch.

C3 releases five original ready packages (launch-provided-job-001,
replace-pending-value-001, merge-two-streams-001, cancelable-forward-001 and
timeout-operation-001) and revises the existing Worker scaffold. The supplied
synctest/gate helpers are guided support, not independent harness mastery. All
five new references pass public/full checks and all 41 compiling defects have
useful public test diagnostics. Append-only placement, activation and persisted
C3 feedback/tolerance regressions remain separate unchecked steps above.
