# Finite-v1 introduction batch I1

Base: `ab880dea6f344f1e90ed777df16ed330dbd0303f` (PR104 merge).
Branch: `codex/v1-introductions-01`. Independent reviewer: `/root/reviewer`.
Content/method/native reviews accepted 2026-10-22; final batch gates and merge pending.

| Unit | Readiness evidence | Finite-v1 credit |
| --- | --- | --- |
| recursive-total-001, 03.40 | Fixed-tests reference and 8/8 defects; public-only fixture; actual recursive base/progress/combine checklist; iterative equivalent passes behavior only | functions.recursion first encounter |
| interval-table-tests-001, 08.07 | Mutation-tests reference detects 8/8 defects; actual named table and shared assertion reviewed separately | testing.table-driven first encounter |
| worker-completion-tests-001, existing 09.01 | Unchanged existing task; learner chooses premature-completion and overlapping-call assertions at controlled boundaries | testing.deterministic-boundaries first encounter only |
| owned-batch-channel-001, 09.17 | Fixed-tests, 11/11 defects, public-only/race and buffered equivalent; private sole producer/closer reviewed | channels.ownership first encounter |
| modules-build-001, manual preparation | Four executed disposable specimens, exact files/commands/status/output and independent checklist | modules/workspaces then build/run: two separate manual first encounters |

This is five units, three new automatic packages, one unchanged reuse and one
manual unit. Six first slots reduce the finite allocation 28 → 22; 77 application
slots and six integrated assessments remain. The automatic route grows 191 → 194;
197 ready packages include three legacy packages, and the checked inventory has
204 packages. Manual outcomes do not enter those counts or learner progress.
All previous addresses, role/skill tuples, relative order and task file bytes
are preserved. Ownership replaces proposed range-drain because ownership is its
missing direct prerequisite; range-drain remains scheduled, not discarded.

## Reuse decision

At base commit above, inspect
`tasks/ready/worker-completion-tests-001/{TASK.md,starter/subject_test.go,solution/reference/subject_test.go,provenance.yaml}`.
The supplied blocked/released scaffold does not inspect completion before release
or overlapping-call independence. The reference's TestWorker1 adds the former;
TestWorker2 chooses two active calls with a held gate and compares completion
channels. Required close-before-work and shared-active-completion defects assess
these learner-owned decisions. Synctest/helper internals themselves receive no
credit. This is a bounded first encounter at a controlled scheduling boundary;
a later independent application must introduce/select a seam. Original creation
identity and every task byte remain unchanged. No new external source or copied
reference-only content was used.

Twelve of the 49 old candidate rows were semantically reviewed: eleven pairs
accepted and arrays.value-semantics needs a real first encounter. Its two proposed
introductions supplied copying through their wrappers/signatures; only
array-copy-address-tests-001 independently assesses the outcome. The repair is
kept open, with one unit reallocated inside the original 302-unit ceiling (current
conservative allocation including that repair: 301). No old task bytes change.
For control-flow.if, decision-table-tests-001 replaces the rejected
independent-policies-001 candidate. For defer, defer-trace-001 then
resource-finalize-001 assess registration/argument evaluation and independent
cleanup. Concrete pinned paths/actions are in v1-backlog.json; 37 candidate rows
remain unreviewed. This review does not change automatic coverage.

## Corrections and evidence

- Public independence initially covered only empty calls. The added
  TestActiveCallIndependence holds different nonempty inputs, drains A while B
  stays held, then drains B. The race-safe shared-pending-values defect now fails
  public/full. Removing that public test in a disposable fixture preserves
  reference/race success but leaves exactly that mutant alive (10/11). Existing
  task-check therefore guards this public assertion without tooling changes.
- The native specimen uses the locally installed Go 1.27 and canonical temporary
  paths, avoiding the observed /var versus /private/var mismatch. GOWORK is
  switched immediately after work init, before work use. Four accepted runs used
  the final procedure, private offline environments and 60-second command limits.
- Native digest, actual normalized outputs, file-edit snapshots and finite review
  are retained under `docs/labs/modules-build-001/`. All disposable roots were
  removed. Correct controls run after missing-use/wrong-module repairs and after
  the wrong-package diagnostic. No timeout is counted as an expected failure.

## Acceptance and accounting

Targeted checks, public-only checks, valid alternatives and the single omission
experiment execute authored code only through disposable task-check fixtures.
Lifecycle transitions use task-promote. The stable-input full CheckCollection
passed 204/204 packages (12:09:23–12:27:03 UTC, 17m40s; child CPU 33m51s).
Quality and both numeric coverage reports were generated from that same file;
all later validator/report consumers must reuse it. Fast tests passed in 1m50s.
Mandatory race/vet and
integration, independent final review and exact-head green validate remain gates.
PR/HEAD/CI and final command intervals will be recorded before completion.

Startup 14:32:40–14:37:43 MSK (5m03s) is conservatively charged to NATIVE-01
diagnosis. Owner wait is excluded; work resumed at 14:48 MSK. Native recipe
selection/path diagnosis adds at most 5 minutes, conservatively counted in the
infrastructure budget: <=10m03s this batch and rolling three-batch window; no
production tooling changed. Local acceptance and review time count toward the
four-hour batch checkpoint; overlapping CI/reviewer intervals are not added twice.

Next batch must retain introduction priority, include range-drain after ownership,
and execute a benchmark reference with the agreed finite evidence criteria.
No full-bank minima, complete 141-row acceptance or learner mastery is claimed.
