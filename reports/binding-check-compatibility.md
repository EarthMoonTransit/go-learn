# Opening lesson public-check correction: workspace compatibility

Date: 2026-09-08. This tooling batch prepares safe delivery of a separate
literal-bindings-001 public-test correction. The learner reported public success
with enabled=false, then a generic full-check failure. The condition already
requires bool true; public tests omitted both label and enabled.

The content digest includes tests, so changing only public checks would otherwise
reject existing workspaces and suggest reset. ADR 0014 permits the exact old →
prepared-new digest pair for this task. Every other revision mismatch remains
rejected. Existing workspace confinement, marker metadata and checker inspection
remain intact. No solution, marker or progress migration is performed. The task
statement, API and all exported workspace files/flags stay byte-identical.

An isolated temporary repository/home reproduced the old public success, applied
only the prepared public test, compared LearnerWorkspaceFiles including editable
flags, then exercised TaskView, Resume and Check using the old marker. Public
checking now reports `enabled = false (bool), want bool true`; all four existing
wrong variants fail public checking, while reference public/full checking passes.
The original marker and edited fixture code remain byte-identical; ordinary
attempt accounting records seven checks. No real learner home was accessed.

Fast unit regressions also cover old and current revisions, reverse and unrelated
revisions/tasks, legacy and generation markers, preserved historical completion,
attempts/hints/time bytes, and invalid schema/task/submission identity. The UI
hidden-failure workspace-patch fixture now uses named-block-001's existing
wrong-element-type variant. It no longer requires keeping the opening defect.
The dedicated opening-route integration continues covering literal-bindings.

Independent read-only agent `/root/ui_goal_reviewer` approved the tooling and
prepared content correction with no blocking findings. Fast tests, race, vet and
tagged integration passed (learner integration: 110.436 seconds). One full
34-package CheckCollection was executed, then reused by phase1 validation and
deterministic report checking; both passed and reports remained current.
Collection SHA256: `723b94434eac253c03124b48d44f9c06c582dea6957d4268ae0cc9b0edcae308`.
Exact-head PR validate remains required before merge.
The content task itself remains unchanged in this tooling batch.
