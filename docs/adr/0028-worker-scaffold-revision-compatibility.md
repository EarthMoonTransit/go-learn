# ADR 0028: Preserve Worker workspaces across gated scaffolding

- Status: accepted for preparation tooling; P13 content is a separate batch
- Date: 2026-10-17
- Content integration base: C2 `fb0abb0408fde73f1e850654b9af6ab39676613c`
- Source decision: S3 `e7c0fa67c6ee8de43e59f202a10f509c3688ee55`

## Decision

P13 adds supplied failure-safe gated work and a named blocked/released/completed
example to the existing Worker test-writing task. Keep Start, all eight mandatory
mutants, assessed skills, role, editable paths and fixed exported files. Following
ADRs 0014, 0026 and 0027, accept exactly this directional content pair:

- Stored: `4c3339ca779815a6c174e242c5bf2e87c1e0e332eb8e5c0419e5bdbe70abeeb1`
- Current: `9b1ad8c690bb9bdfffa46027949fa76d2ad7f77bb04897bb1f5a48e5eede8c69`

Historical and proposed starter/hints are pinned under
`internal/learner/testdata/worker-before-gated-scaffold/` and
`internal/learner/testdata/worker-gated-scaffold/`. Canonical task content remains
unchanged in this tooling batch. Its later content revision must use these exact
bytes and refresh its actual provenance/review/report evidence.

The supplied example starts Start in a controlled goroutine so inline work
cannot trap the test's own release path. Every failure releases its gate and
waits for runnable work; phase assertions use synctest.Wait and useful named
actual/want feedback. It still accepts the premature-completion and shared-active
completion variants. The learner must add those independent cases. The existing
reference must continue detecting all eight variants. Supplied scaffolding adds
no independent evidence of the learner designing a channel test harness.

## Preservation boundary

The task remains an ADR 0008 original: retain its existing real creation record,
project identity and empty external sources. S3 supplies facts and original
revision reasons only; no new upstream task material is borrowed. No original
commit or release binding is invented.

Historical numbered and legacy paths, learner edits, marker bytes/permissions,
paused/completed state, attempts, hints, activity and completion/encounter history
remain. Resume and public/full verification continue in place. Unknown pairs,
reverse migration and other task IDs are rejected; existing schema, generation,
submission, fixed-file and symlink checks still apply. No real learner home is
used. Rollback leaves old workspaces usable; new workspaces need the new revision
or explicit reconciliation.

## Acceptance

Focused fast tests cover both marker formats, paused/completed progress and
numbered/legacy paths. Integration resumes old work through the real disposable
verifier, accepts both starters publicly while keeping incomplete mutation
coverage private, checks named phase diagnostics, and accepts the reference
against all eight variants. A bounded eight-variant public matrix checks actual
new-starter behavior and useful failures, including the controlled inline case;
the two intentionally uncovered defects remain public successes.

This separate tooling batch also persists the repaired C2 public feedback:
three Formatter cases and all eleven Resource variants, with concrete diagnostics
and passing reference counterparts through public Check. Assertion-removal
experiments use disposable repository copies and must make the relevant outer
regression fail on unexpected public success. Authored task code executes only
inside disposable verifier workspaces.

Ordered fast/race/vet/integration acceptance remains mandatory. Reuse the single
successful C2 full collection only after confirming all 171 task inputs and
checker/policy are identical; validator and report check use that same collection.
This is ADR 0022 preparation, with no route, scheduler, schema, taxonomy, source
permission, canonical task, learner state or release/export changes.
