# Finite Go Learn v1 delivery

Owner-approved 2026-10-20; normative boundary: [ADR 0034](../adr/0034-finite-v1-course.md).
M0, M1 and M2.1–M2.3 are complete. On 2026-10-22 the owner authorized autonomous
reviewed batches through finite-v1 readiness. This replaces every historical
pause-after-five restriction below; no ordinary owner confirmation is required
after an accepted PR. One independent reviewer checks contracts and final diffs.
The historical delivery records remain evidence of their original authorizations.

| Stage | Deliverable | Exit |
| --- | --- | --- |
| M0 | Documentation PR: scope, assessment, allocation and stop policy | Independent review, exact-head green validate and merge |
| M1 | Five tasks: array indexing, nil/empty slices, promoted fields, masks, shifts | All five accepted, routed, reviewed and merged after green validate (completed) |
| M2, authorized continuation | Remaining encounters and independent applications | All 141 outcomes have two distinct reviewed encounters including independent application |
| M3, authorized continuation | Six integrated assessments and final readiness review | Complete evidence and honest automatic/manual/full-bank report |

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

## Operating limits (owner update, 2026-10-22)

- Normally 5–8 tasks per content batch, never over 10, each on its own
  `codex/<topic>` branch and PR. Tooling/taxonomy remain separate. Reuse suitable
  tasks/components; 105 remaining skill slots do not require 105 new files.
- Begin with the five introductions listed in the current checkpoint. While
  three introductions are available, include at least three in each batch;
  otherwise include every available introduction and its specific dependencies.
  All 28 first slots open at PR104 must be accepted by new content batch ten,
  or stop with a concrete blocker. Do not substitute convenient applications.
- By content batch two, execute and independently review a compliant minimal
  modules/workspaces specimen, including actual go.mod/go.work edits and a
  material wrong variant. Define benchmark evidence in the first two batches
  and execute one reference before the next ordinary series: equivalent results,
  setup/work boundaries, representative inputs, repeated runs and allocations;
  no machine-dependent ns/op threshold. Unexecuted instructions do not qualify.
- Review the 49 baseline candidate pairs by topic alongside content: two distinct
  learner actions, independent application, exact skill meaning and assessment.
  Record concrete decisions in the existing backlog; metadata alone earns no
  credit. Review/reuse existing components, including bounded-lru-cache-001.
- No new checker profiles, AST/rewriter/trace mechanisms, P21 engine, universal
  runners or platform extensions. An infrastructure fix must name a reproducible
  blocker of a current task and use a separate tooling PR. Diagnosis plus fixes
  are limited to 60 minutes per batch and two hours in any three consecutive
  batches; required checks are recorded separately. No splitting to reset limits.
- One independent reviewer checks the contract/source decision before authoring
  and the corrected final diff. Findings need a violated requirement and concrete
  evidence. After two unsuccessful repairs of one finding, reconsider the
  contract/assessment; the defect remains blocking. Method review inspects actual
  code and a finite checklist separately from behavioral pass.
- Use task-check and permitted disposable execution, task-promote, provenance,
  public/grader/reference/mutants and useful public feedback. Preserve all existing
  IDs, addresses, relative order, solutions and progress. After stable inputs,
  one full local CheckCollection feeds validator/report consumers. Keep all
  mandatory fast/race/vet/integration gates; changed inputs need fresh evidence.
  Merge only after independent final review and green validate on the final HEAD.
- Each batch has a four-hour available-execution checkpoint, excluding only
  documented CI-only idle time, host sleep and user wait. Local checks, review and
  content work count; overlapping reviewer time is not added twice. At the limit,
  stop expansion, finish an already running mandatory check, preserve/merge only
  fully accepted work, report the blocker and do not start another batch.
- After merge update the short report/backlog and this checkpoint, then continue.
  Stop for completion, owner instruction, a real blocker, exhausted limits or a
  necessary scope/rule change. One repeat is allowed for a confirmed transient
  external outage; the same repeated failure blocks. No new controller; handoff
  preserves the goal and does not mean completion.

## Current checkpoint: first introduction batch

- Owner/controller: current finite-v1 goal, sole independent reviewer
  `/root/reviewer`; repository `/Users/egor/Documents/ChatGPT/go-learn`.
- Branch: `codex/v1-introductions-01`, base
  `ab880dea6f344f1e90ed777df16ed330dbd0303f` (PR104 merge). PR104 validate
  [39726118510](https://github.com/EarthMoonTransit/go-learn/actions/runs/39726118510)
  succeeded on `5c38dcc84969cb41a6ff55574f00e3afea9a9801`. Current I1 content,
  method, native and route reviews are accepted; final local gates, PR/HEAD review
  and exact-head CI/merge remain pending. The stable full collection passed all
  204 packages in 17m40s; generated quality/coverage reports reuse that collection.
- Baseline reconciled: 191 routed lessons, 194 ready packages, 201 total checked
  packages; 28 first slots, 77 applications and six components remain. Automatic
  counts, manual acceptance and unknown learner mastery remain separate.
- Selected first five: functions.recursion (behavior plus real decomposition,
  base/progress/combine checklist); testing.table-driven (learner cases plus
  mutation checks and method review); testing.deterministic-boundaries (controlled
  seam/signals without sleep; first inspect worker-completion-tests-001 for reuse);
  channels.ownership; toolchain.build-run (actual build/run/test evidence).
  Reviewer confirmed that range-drain must follow its still-missing ownership
  prerequisite, so ownership replaces it under the approved substitution rule;
  range-drain remains a required later introduction. Build-run directly requires
  modules-workspaces; the existing preparation text supplies context but earns
  no finite-v1 evidence. Recursion/table-driven/deterministic-boundaries have
  available direct prerequisites. Worker reuse has been accepted for the bounded
  first encounter only; its future independent application remains open.
- Resolved rule blocker NATIVE-01, independently confirmed 2026-10-22: the previous
  AGENTS rule 9 required task-check without a native-lab exception.
  `internal/tasking/runner.go` assembles fixed go.mod and test/vet commands;
  `internal/isolation/legacy.go` rejects build/run and benchmark flags;
  `internal/isolation/isolation.go` forces GOWORK=off; task contract validation
  disallows editable go.mod/go.work. The reserved benchmark profile is not
  executable. Calling the low-level isolation API or subprocesses from authored
  tests would not establish an authorized assessment path.
- The owner explicitly approved the proposed finite manual procedure on
  2026-10-22; AGENTS rule 9 and ADR0034 now record that narrow exception.
  The combined modules/workspaces then build/run lab has now been executed and
  independently accepted with separate assessed actions. Next: finish the single
  full collection, required local gates, final HEAD review and PR validate.
  Content/method/native review has accepted the five units; final gates and merge
  remain pending. Resume timestamp: 14:48 MSK.
- Startup/diagnosis accounting: 14:32:40–14:37:43 MSK (5m03s); conservatively
  charged entirely to NATIVE-01 diagnosis in this batch and rolling three-batch
  infrastructure budget. Code/fix time is 0; never add that interval twice.
  Owner-wait after the blocked response is excluded from the four-hour checkpoint.
  Native recipe/path diagnosis adds a conservative 5-minute upper bound: cumulative
  infrastructure <=10m03s, no production tooling changes. [I1 report](../../reports/v1-introductions-01.md)
  records six proposed closed first slots, separate manual evidence and current gates.
- Old-pair review: 12/49 rows inspected, 11 accepted, arrays.value-semantics
  requires one first-encounter repair; 37 rows remain. The repair is allocated
  within the frozen ceiling (conservative current total 301/302), separately
  from the original 28-first-slot deadline. No supplied mechanism earns credit.

### Authorized finite manual procedure (owner decision, 2026-10-22)

This permits finite manual execution, not a runner or automatic assessment.
Each specimen still requires real execution and independent acceptance:

1. Permit only finite-v1 manual units outside tasks/ready to execute their exact
   reviewed standard commands in disposable copies without task-check. Keep all
   automatic task gates and source-checkout execution prohibitions unchanged.
2. Each unit names its complete file set, permitted edits, exact commands and
   expected outputs before execution. Copy only regular reviewed files into a
   fresh temporary directory; reject symlinks/escapes. Use local Go, offline
   stdlib/local modules, private home/config/cache/temp directories, telemetry off,
   and a clean child environment. Only the modules/workspaces specimen enables
   GOWORK pointing to that temporary go.work; never use the user's workspace.
3. The build specimen runs go build, its disposable executable, go run and go test;
   the module specimen creates/edits go.mod/go.work and runs the intended local
   packages. The benchmark specimen runs correctness tests and repeated native
   benchmarks with allocations. Each command is bounded (at most 60 seconds),
   with command/cwd, Go version, exit status and actual output retained. No new
   checker profile, generalized executor, AST/trace analysis or performance grade.
4. Execute reference and realistic wrong examples from independent disposable
   copies. Reviewer inspects actual files, outcomes, useful feedback and method
   checklist. Record source/original provenance and content digest; bind actual
   commits on release. Remove temporary execution artifacts after retaining
   evidence. A described but unexecuted or unreviewed unit stays unfinished.
5. Link accepted manual work in the existing learner index with prerequisites
   and sequence; record manual readiness separately, without automatic progress
   or coverage credit. Reuse the authorized procedure for analogous v1 labs,
   with each finite command list reviewed; no general arbitrary-command permission.

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

## M2.2 bounded application batch

Owner-authorized on 2026-10-21, limited to five applications on existing profiles:
`consume-slice-prefix-001`, `move-overlap-range-001`, `spare-view-tests-001`,
`guarded-stage-values-001`, `transfer-shared-values-001`. These change the
observable contracts relative to their introductions; source/design approval
precedes authoring. The pointer-dependent slice task follows pointer mutation,
and the stored-function task follows the closure introduction.

The route grows from 181 to 186, retaining all prior addresses and order.
Remaining allocation: 28 first encounters, 82 applications, six components.
Keep independent content/public-feedback/final-diff review, the single shared
collection, required local gates and green exact-head validate before merge.
Apply the same four-hour work checkpoint and separate overlapping local/CI
intervals and host sleep. Stop after this batch and report results/time; no
subsequent batch, P21 work or new checking engine is authorized.

## M2.3 bounded application batch

Owner-authorized on 2026-10-22, limited to five applications on existing profiles:
`reset-selected-state-001`, `named-step-conversion-001`, `iota-layout-tests-001`,
`replace-rune-range-001` and `first-pair-positions-001`. Independent source/design
review precedes authoring. The selected-state task adds selective ownership and
selection counting beyond the existing receiver reset. Named scalar conversions,
paired iota declarations with a skipped position and a new block, variable rune
splicing, and first-occurrence tuple classification change their introduction
contracts. Supplied declarations/observation wrappers are not learner-authored
mechanism evidence; acceptance covers the stated observable application.

The route grows from 186 to 191, retaining all prior addresses and order.
Remaining allocation: 28 first encounters, 77 applications, six components.
All five applications follow their introductions and prerequisites. Keep the
independent content/public-feedback/final-diff review, single shared collection,
required local gates and green exact-head validate before merge. A corrected
invalid-UTF-8 key distinction is protected by a tracked public assertion and
mutant; the grader cannot mask removal of that assertion.
Apply the same four-hour work checkpoint and separate overlapping local/CI
intervals and host sleep. Stop after this batch and report results/time; no
subsequent batch, P21 work or new checking engine is authorized.
