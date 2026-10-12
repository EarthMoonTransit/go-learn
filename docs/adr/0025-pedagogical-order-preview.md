# ADR 0025: inactive pedagogical order preview, stable lesson addresses

- Status: accepted for T0a implementation; activation requires a separate PR
- Date: 2026-10-12
- Phase: preparation/tooling (ADR 0022), not foundation release/export

## Decision

Keep three separate concepts: the stable address registry in route v2, thematic
chapters used for presentation, and a proposed pedagogical sequence. Preserve
every existing task ID, block ID, block number, step number, role and assessed
skill. Route v2 keeps its strict fields and increasing-number contract. A chapter
or proposed position is never a lesson address or a workspace path component.

Add `curriculum/plans/pedagogical-order.yaml`, a strict, single-document planning
sidecar with `schema_version: 1`, `status: inactive`, `baseline_commit`,
`mapping_file`, `exposure_edges` and `pending_preparation`. Only `inactive` is
supported in T0a; missing/active status and other versions are rejected.
`mapping_file` is a repository-relative path, read through the existing safe I/O
helper. CSV columns retain the supplied review specification verbatim:
`task_id,old_number,proposed_chapter,proposed_subtopic,proposed_order,role,disposition,rationale`.
The original CSV is stored without edits, including its UTF-8 BOM. Its baseline
is `b9790106edc3d4c7f57e26d5f45dc1960f022c6c`; this is a planning reference, not a
claim about the current master or test results. Review inputs are supplied in
`go-learn-review-package.zip/go-learn-review/`; they are planning evidence, not
new task sources or permission to author the proposed content.

Validation compares against the **current** route and ready v2 catalog, with no
hard-coded 151 limit. Report all unknown/duplicate/missing IDs, unavailable tasks,
changed old numbers/roles and noncontiguous or duplicate positions. A structurally
invalid mapping is displayed with diagnostics and a failing CLI status; it cannot
be treated as a pedagogically valid plan. CSV syntax/sidecar errors also fail.
Never discard new route lessons or silently repair a stale mapping.

For a structurally valid mapping, walk its explicit positions, starting with
route preparation skills. Check actual manifest prerequisites against earlier
assessed skills and direct taxonomy prerequisites against earlier or same-lesson
assessed skills, matching the existing route introduction contract. These are
static exposure checks, not mastery or evidence of a learner's completion.
Reviewed `exposure_edges` additionally check specific existing task placements.
`pending_preparation` records unreleased content/metadata actions by review ID,
target task ID and reason. Review/proposal IDs never enter the task sequence.
Pending actions cannot be marked released in this format; later content/tooling
PRs must review and update the plan. P13 is a scaffold revision of the existing
Worker task, not a new task ID. These checks do not certify unreviewed implicit
prerequisites or content correctness.

Extend the existing read-only CLI with `learn roadmap --pedagogical`. Show
proposed chapter/subtopic/position, old address, task ID, role, learner status,
actual manifest prerequisites, formal blockers and pending preparation. Valid
but blocked plans remain previewable (exit 0); structural errors exit 1. Always
print inactive status separately from structural validity and preparation
readiness. No UI redesign or second scheduler is introduced.

## Blockers and compatibility

Unchanged baseline metadata produces two manifest blockers:
`decision-table-tests-001` and `constant-value-tests-001` need
`testing.basics-diagnostics` before its current first assessed occurrence.
E01 plus the coordinated M1 Clamp secondary change must fix that in a separate
content/metadata batch. T0a preserves and displays both warnings. P05 before
Shape, P06/P07 before Capture and P11/P12/P13-scaffold before Worker are still
unreleased semantic gates. Other M1 reviews and E02 are also visible preparation
actions; `value-average-001` secondary remains a conditional review, not a new
mandatory skill or automatically applied metadata.

Only preview reads the sidecar. Production `Next`, `dueReviewChoice`, UI views,
lookup, GoLand targets and workspace checks continue to use route v2.
Read-only preview does not save/upgrade progress, start a timer, recover journals,
create/move workspaces, or change current selection, attempts, hints, assistance,
timestamps, encounters, completion history or CompletionEvidence. Numbered paths
still use the old block ID/number and step number; legacy `workspaces/<task-id>`
remains accepted. No real learner home is used in acceptance tests.

## Deferred T0b and rollback

T0b may pass a validated, prepared order into the **existing** scheduler with
priority unfinished current task → eligible due review → next lesson. Readiness
means introduced mechanisms, not mastery of the previous course or completion of
the review task itself. No eligible review leaves the review due and allows
ordinary progression. Preserve interval 3, mastery from two distinct independent
practice/transfer task IDs, assistance and the rule that reset adds no evidence.
PointerChain review eligibility needs the reviewed mutation-sharing gate.

One explicit future readiness equivalence may recognize old completed Clamp as
familiarity with basic Test API, after reviewing its old contract/scaffold. It
must not add a secondary to historical CompletionEvidence or grant independent
mastery. There is no generic retagging of historical completions. Ambiguous legacy
progress remains legacy/assisted under ADR 0017. None of this is implemented here.

Rollback removes the preview and planning sidecar; no user-data migration or
workspace rollback is necessary. Future activation needs its own ADR/acceptance,
released bridges, independent review and exact-head green `validate`.
