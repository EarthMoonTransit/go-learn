# Opening lesson public checks

Date: 2026-09-08. A learner reported public success followed by a generic full
failure in literal-bindings-001. The condition already specifies the type and
value of all four variables, but public tests checked only count and scale.

Public checks now also assert label is string G and enabled is bool true, with
actual value/type diagnostics. Every existing wrong variant is detected publicly;
correct reference code passes both public and full checking. Statement, API,
starter, fixed files, hints, grading expectations and task ID remain unchanged.
Original provenance retains its creation record and updates the authored digest.
Review metadata records the independent review of this ready-task correction;
no lifecycle transition or replacement package is fabricated.

ADR 0014 provides exact directional compatibility for existing workspace markers.
A temporary repository/home acceptance reproduced the omission, proved exported
workspace bytes/editable flags unchanged, exercised the old workspace after the
public-only update and verified actionable enabled diagnostics, all four wrong
variants, correct public/full completion and unchanged marker/solution bytes.
Real learner solutions/progress were not read, copied or modified. Existing UI
servers must be restarted to use the compatible learner code; no task reset is
needed. A prior completed task remains completed.

Independent read-only agent ui_goal_reviewer approved the prepared content and
compatibility evidence with no blocking findings. Fast tests, race, vet and
tagged integration passed. One full 34-package CheckCollection regenerated the
report; phase1 validation and deterministic report-check reused it and passed.
Collection SHA256: `d9b7d966ed812d8fc6bd13eef1719f4aec0f5dd754ff6109b60880170b4ecedc`.
All four required variants are detected; independent temporary-home acceptance
also verifies they fail publicly. Exact-head PR validate is required before merge.
This is one corrected existing task; learner-visible inventory remains 22 v2,
with three legacy ready packages, eight drafts and one rejected package separate.
