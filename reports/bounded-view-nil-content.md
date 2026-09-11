# P4C / C16 — public nil-input correction

bounded-slice-view-001 now publicly checks the existing nil-input requirement:
valid zero bounds produce zero length and capacity. The diagnostic reports both
actual and expected dimensions. Nil and non-nil empty zero-capacity outputs are
both valid. A nil-only spare-capacity mutant joins the six existing defects.

TASK, starter, reference, hints, task ID, route number and exported workspace
payload remain unchanged. Provenance records the corrected digest; P4T carries
its exact directional compatibility. No migration/reset or new ready ID is used.

Targeted task-check passed and Service.Check regression passed on this content.
Independent P4T/P4C review recomputed both digests, ran reference/public/full and
confirmed that removing only the nil public assertion causes unexpected public
success and breaks the regression. Requirement/assertion/defect/diagnostic mapping
is in review/notes.md. All task execution used disposable verifier copies.

This is an incomplete-bank preparation amendment. Full collection/report refresh,
local acceptance and exact-head CI are required. GitHub Actions is currently
billing-blocked, so this PR cannot be merged. P4T must land first.
