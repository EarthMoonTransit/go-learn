# Introduction review

Contract/source decision approved by /root/reviewer; actual content and execution approved 2026-10-22. First encounter only; no learner mastery or full-bank claim.

| Requirement | Public/reference assertion | Defect / method evidence |
| --- | --- | --- |
| return before release, no sends before signal/close | TestOwnedOutput / controlled pre-release assertions | wait-before-return, ignore-gate |
| non-nil independent channels | TestOwnedOutput; TestIndependentOutputs | nil-output, shared-output |
| active calls keep separate data and release boundaries | TestActiveCallIndependence | shared-pending-values |
| all values exactly once/order including zero and negative | TestOwnedOutput/mixed,closed-gate | skip-last, duplicate-values, zero-ends-stream, reverse-values |
| close owned output after all sends, including empty | TestOwnedOutput/nil,empty,mixed | never-close |
| preserve input until output closes | TestOwnedOutput/mixed | erase-input |
| sole producer ownership explanation | actual reference source + TASK.md method checklist | API restriction on ready is supplied, no separate credit |

Diagnostics identify case/input and actual versus expected result or lifecycle phase. Compile errors, timeouts and panics do not establish defect detection. Method review inspects actual code and the learner-visible checklist in TASK.md, separately from automated behavior.

Every behavioral defect must fail public-only checks in disposable task-check fixtures. An iterative recursive-total equivalent is a valid behavioral alternative but fails the separate method criterion.

Public feedback correction: TestActiveCallIndependence holds two different nonempty inputs, releases/drains A while B stays held, then releases/drains B. The required race-safe shared-pending-values defect is detected by this public assertion. A disposable omission experiment must leave exactly that defect alive, so the existing task-check mutant gate protects the public regression without a tooling change.

Independent actual-content approval by /root/reviewer on 2026-10-22. Original creation/provenance, API, story-free contract, public feedback and method checklist verified. Targeted technical checks pass; all required defects detected, with applicable race passing. Batch/route/CI acceptance remains separate.

Method specimen: solution/reference/subject.go allocates one private output per call; its single producer waits for ready, sends each value, then closes output. Caller input lifetime is explicit. Buffered equivalent passes public/full/race. Removing TestActiveCallIndependence in a disposable copy gives passing reference/race but exactly shared-pending-values survives (10/11), making the declared mutant an automatic regression of this public assertion.
