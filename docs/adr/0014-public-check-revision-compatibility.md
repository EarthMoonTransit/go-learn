# ADR 0014: Compatibility for the opening lesson's public-check correction

- Status: accepted
- Date: 2026-09-08

## Problem and decision

literal-bindings-001 asks for four specific variable types and values, but its
public tests check only count and scale. A learner reported public success with
enabled=false followed by an opaque full-check failure. A separate content batch
adds label/enabled assertions and actionable public diagnostics.

Public tests do not enter the persistent learner workspace. Nevertheless,
TaskContentDigest includes verification files, so this correction would otherwise
invalidate existing workspaces and suggest a destructive reset. Permit only this
exact, directional compatibility pair for literal-bindings-001:

- Stored: f427ca2b6b99603eab0f689a70e9955066c974f5b91d1521b795e5a14292947b
- Current: f96c2cdcdd423a608cc7c02214e595bc02fdd3aa00ce9fc03478df2ef2afe3d5

The only authored-content delta is tests/public/subject/public_test.go. TASK,
starter, fixed files, hints, API, submission layout and editable flags remain
identical. Keep full authored digests and all other revision mismatches rejected.
This is a bounded compatibility exception, not a general revision migration.

## Preservation and rollback

No marker, solution or progress file is rewritten for compatibility. Legacy and
generation-bearing markers are supported. Schema, task/submission identity,
workspace confinement, fixed-file and symlink checks remain enforced. Check and
Complete retain their ordinary attempt/time behavior; compatibility itself grants
no completion and does not change a prior completion or hint record.

Old markers also work after rolling the content back. New markers deliberately
remain incompatible with the old content revision: retain the corrected content
and compatible CLI for such workspaces. Do not strip markers or reset learner
work automatically. No progress or task schema changes are introduced.

## Acceptance

Unit tests cover the exact directional pair, unrelated tasks/revisions, legacy
markers, generation markers, unchanged stored bytes and invalid marker metadata.
A disposable-home acceptance compares exported workspace files and editable flags
before/after the prepared revision, then exercises old-workspace view/resume,
public diagnostics and full completion with the existing disposable checker.
The UI hidden-failure workspace-patch fixture uses named-block-001, whose existing
wrong-element-type variant still passes public checks and fails full checking.
It no longer depends on preserving the reported opening-lesson defect.

## Counter public race workload (2026-10-23)

COUNTER-WORKSPACE-REVISION-01 applies this same bounded policy to
`concurrent-counter-read-001`. The reviewed content correction adds a finite
common-start reader/writer public workload after CI retained an
`unlock-before-read` survivor. It changes no learner-exported material:
TASK.md, go.mod, starter implementation, paths and editable flags stay identical.
Permit only this additional directional pair:

- Stored: abff4eff554adc02cda1c3179cbd49291414e4e665f801c07d32844e3ad3f41a
- Current: 0896d577d779512d5e773cd92130702056624033e3f3b138ab23e2cd49319c03

The pinned old public fixture comes from accepted `c7a178b8` material; the new
fixture is the independently accepted I12 common-start public source. Both are
inert testdata. Tests reconstruct the exact historical content digest in a
private repository before creating workspaces, then project the new public bytes
only into that disposable repository. Full learner-export path/content/editable
comparisons must be equal. The canonical task is never edited by this tooling
change or by its tests.

Legacy and generation markers, in-progress and completed histories, solution
bytes, hints, feedback, accumulated time and existing completion evidence remain
intact on read/compatibility. Resume and real disposable Check/Complete retain
their ordinary selection, session, attempt and completion effects; a repeated
Complete stays rejected. The old marker is never upgraded implicitly. Exact-pair
unit cases reject rollback, unknown source/target and another task ID. Existing
schema, identity, fixed-file, confinement and symlink guards are unchanged.

This is an application of ADR0014, not a general content migration or a new
verification mechanism. The source adaptation pins and the dynamic race workload's
finite observation limits belong to the separate content decision. Separate
review, local gates and exact-head green validate remain required for delivery.
