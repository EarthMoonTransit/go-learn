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
