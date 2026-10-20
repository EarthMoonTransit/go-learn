# Finite Go Learn v1 delivery

Owner-approved 2026-10-20; normative boundary: [ADR 0034](../adr/0034-finite-v1-course.md).
M0 and M1 are complete. The owner authorized one further five-task batch (M2.1),
then another pause and result/time report; no subsequent batch is authorized.

| Stage | Deliverable | Exit |
| --- | --- | --- |
| M0 | Documentation PR: scope, assessment, allocation and stop policy | Independent review, exact-head green validate and merge |
| M1 | Five tasks: array indexing, nil/empty slices, promoted fields, masks, shifts | All five accepted, routed, reviewed and merged after green validate; pause and report |
| M2, later authorization | Remaining encounters and independent applications | All 141 outcomes have two distinct reviewed encounters including independent application |
| M3, later authorization | Six integrated assessments and final readiness review | Complete evidence and honest automatic/manual/full-bank report |

## Finite allocation

[Backlog](v1-backlog.json), baseline `21d89d59939c3ee1f7ea891c136798decce9e9da`:
171 lessons plus up to 33 first encounters, 92 independent applications and six
components. M1 consumes five first-encounter slots. These are planning slots,
not admitted tasks. Source admission, deduplication and quality precede authoring
and acceptance. Existing suitable components and reviewed combined credit can
reduce work; 302 is a ceiling, not a quota.

Keep the six theory pages, toolchain, registry, UI, scheduler, profiles and
historical learner state. Defer specialist sync outcomes, deeper practice and
new platform features. If review invalidates baseline assumptions, report the
specific gap and repair/reallocate; do not silently expand or claim completion.

## First five outcomes

Bounded source/design review precedes authoring:

1. `arrays.indexing-iteration`: update indexed array elements, not range copies;
   respect bounds and unselected elements.
2. `slices.nil-empty`: preserve observable nil versus non-nil empty results,
   including non-empty inputs with no matches.
3. `embedding.field-promotion`: choose promoted or explicitly qualified fields
   when outer and embedded ownership differ.
4. `expressions.bitwise-masks`: inspect/update selected bits and preserve others;
   zero and overlapping masks matter.
5. `expressions.integer-shifts`: explicit fixed-width runtime shifts at zero,
   width and beyond without modulo reduction.

Each is a first encounter, not full-skill depth or every mode. State any
unassessed facets and supplied mechanisms. Use existing implementation,
debugging or learner-test profiles only when they assess the intended outcome.

## Operating limits

- Normally 5–8 tasks per content batch, never over 10. Tooling/taxonomy remain
  separate. Preserve bounded source-first decisions.
- No new verification profile or general checking project. A tooling fix names
  the current task and failing condition.
- Review content/contract and then the corrected final diff. Extra review is
  warranted by changed evidence or unresolved findings, not new design stages.
- Debug affected tasks with task-check in disposable workspaces. After stable
  inputs, one full CheckCollection feeds all consumers. Keep required gates;
  changed inputs/failures, not elapsed time, justify rerunning.
- M1 has a four-hour agent-work checkpoint excluding CI-only idle time. If
  exceeded, stop expansion and report results/blocker; never waive quality.
- Record start/end and command/CI intervals. Distinguish wall time, available
  goal accounting, local checks and CI-only waiting; do not add overlapping
  intervals together.
- After M1 pause the goal and report task IDs/outcomes, checks/review, PRs,
  defects and time. No automatic M2. Estimate later work from measured M1 pace.

## Ready means

All frozen outcomes have the two distinct encounters and six components assess
meaningful combinations. Automatic tasks pass existing quality gates; manual
units have fulfilled finite execution/review criteria. No required future
engine or placeholder remains. Prerequisites and existing state are preserved.
Applicable checks and exact-head GitHub validate pass.

Report automatic ready packages, manual units, remaining full-bank minima and
unknown human mastery separately. This milestone is not full-bank `-coverage`
or a full-bank release/export.

## M1 delivery record

Five original task packages consume the five selected first-encounter slots:
`array-window-update-001`, `nil-preserving-filter-001`,
`promoted-field-selection-001`, `mask-state-tests-001` and
`fixed-width-shifts-001`. Their task reviews record the contracts, defect
matrices and limits; generated reports remain the automatic acceptance evidence.
The route grows from 171 to 176 without renumbering previous lessons. This is
five introductions, not five mastered skills. The historical allocation baseline
above stays fixed; `v1-backlog.json` links delivered slots separately.

A preliminary wall-clock alert during M1 was corrected after inspecting host
sleep intervals: at 16:44 UTC, roughly 8h44 of 9h57 elapsed had been host sleep.
The remaining 1h13 is an upper bound on active work, not measured CPU/model time;
CI overlaps some sleep intervals. The four-hour agent-work limit was therefore
not shown to be exceeded. The final report separates wall time, observed host
sleep, command intervals and CI without adding overlapping intervals.

## M2.1 bounded application batch

Owner-authorized on 2026-10-20 after the M1 checkpoint. This batch selects the
five independent-application slots corresponding to M1: arrays, slice presence,
promoted fields, masks and shifts. It changes the observable contracts and uses
implementation, debugging and learner-written tests on existing profiles.
Source/design approval precedes authoring; each application must follow its
introductions and all prerequisites. Existing addresses/order are preserved.

Exit remains five accepted and routed tasks, independent content/final-diff
review, mandatory local checks and exact-head green GitHub validate before
merge, then stop and report. Apply the same four-hour work checkpoint and
separate overlapping check/CI intervals and host sleep. No new verifier,
taxonomy, skill identifiers or automatic completion semantics are authorized.
The first-encounter backlog remains visible; this batch consumes application
slots only and does not establish full-bank coverage or human mastery.

M2.1 delivery: `scatter-snapshot-001`, `presence-concat-tests-001`,
`embedded-field-update-001`, `replace-mask-field-001`,
`shift-boundaries-tests-001`. The five reviewed application slots are linked
in the backlog; first-encounter slots are unchanged. The route has 181 lessons
and preserves all 176 prior tuples/order. Generated reports and PR acceptance
are the executable evidence; no claim that the entire v1 course is complete.
