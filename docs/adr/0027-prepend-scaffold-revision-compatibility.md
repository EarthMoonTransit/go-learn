# ADR 0027: Preserve Prepend workspaces across the variadic bridge

- Status: accepted for preparation tooling; E02 content is a separate batch
- Date: 2026-10-17
- Source baseline: master `49c6202c1c09f7462a98746a0ace883192a24b49`
- Integration base: M1 PR #70, `7b958d84bfa299e501fdcbc414d8631350be546d`

## Decision

E02 supplies both individual variadic arguments and slice expansion in
`prepend-values-tests-001`. Replacing starter or hints changes TaskContentDigest,
while the API, accepted submissions, editable paths and six mutation models stay
the same. Following ADRs 0014 and 0026, accept exactly this directional pair:

- Stored: `334ffccf8a0b227463a7ef227e02709e0278961f074d32bef639cbcc4aa050f6`
- Current: `bc2c7c7028ae1e42ea2bf5144995a5f9698e095b81f86fa1e6f8192e5bff98b8`

The target is pinned by `internal/learner/testdata/prepend-variadic-bridge/`.
The historical starter/hints are preserved byte-for-byte at
`internal/learner/testdata/prepend-before-variadic-bridge/`, from the base commit's
`tasks/ready/prepend-values-tests-001`. Both variants are test fixtures here;
canonical task content and metadata remain unchanged until E02 review/release.

The revised examples retain uncovered behavior so completing the starter still
fails full verification. Neither example adds input-mutation, allocation or
nilness guarantees beyond TASK.md. Both forms are guidance, not enforced test
organization. The content batch also records the actual existing slice-header
and slicing prerequisites, without adding assessed skills.

## Source and preservation boundary

Reuse the previously admitted Prepend package's provenance and creation record.
Its adapted source remains registry `exercism-go`, commit
`ae1477abfeeb9fbeaf971caee220e91f9e4df790`, exact paths
`exercises/concept/card-tricks/.docs/instructions.md`,
`exercises/concept/card-tricks/card_tricks_test.go`, metadata and LICENSE, with
the artifact hashes and MIT notice retained in its existing provenance. This
fixture revision uses the admitted local API/contract; no new external statement,
canonical cases, implementation or tests were consulted. It is not a fresh
upstream audit or broader source admission.

An old workspace retains learner file bytes, marker, fixed files, path and
permissions. All read-only exported files and editable flags remain identical.
Paused work resumes in place. Completed history, skills, hints, attempts and
activity are preserved. New prerequisites describe exposure, not new completion
evidence. No actual learner home is inspected or changed.

Unknown source/target digests, another task ID and reverse migration are rejected.
Existing schema, submission identity, generation, symlink and fixed-file checks
remain in force. This is one pinned transition, not automatic compatibility for
all future scaffold revisions. Rollback preserves old workspaces; workspaces
created from the new revision require that revision or explicit reconciliation.

## Acceptance

Fast regressions cover paused/completed status, both marker formats and
numbered/legacy paths, with personal learner comments and preserved evidence.
Integration resumes the old workspace through the real disposable verifier,
accepts old/new starters publicly while rejecting their incomplete coverage
privately, verifies useful public failure feedback for each supplied example,
and accepts the unchanged reference against all six mutants.

Required ordered repository suites, one full collection and shared validator/
report evidence remain mandatory. This is preparation tooling under ADR 0022.
The route, 151-row mapping, scheduler, taxonomy, schemas and canonical task
packages remain unchanged. After merge, E02 separately applies these exact files,
refreshes provenance/review/reports and removes only its pending preparation.
