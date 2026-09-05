# GoLand workspace metadata — T01 acceptance

Date: 2026-09-05. Baseline: `1820c2c84a6a3988bfbf635cdbb23d8fcc177099`.

## Reproduction and boundary

A synthetic workspace accepted for each submission kind was rejected after adding
root `.idea/workspace.xml`. The baseline reproduction and the new regression
suite both observed the undeclared-directory error before the change.

The fix permits only root `.idea/` metadata. Inspection still walks its contents
and rejects symlinks and nonregular files; metadata never enters selected
implementation, test or workspace-patch files. Root `.idea` must be a directory,
and task declarations cannot use that reserved path. Fixed files, unknown
paths/directories, additional workspace-patch files, and optional editor API
rules retain their existing boundaries. Inspection does not rewrite IDE files.

## Executed evidence

- New all-kind tests compare submission selection with/without nested metadata,
  repeat inspection and compare metadata bytes.
- Adversarial cases cover root/nested symlinks, nonregular metadata on Unix,
  `.idea` as a file, similar-prefix and nested unknown directories, changed fixed
  files, undeclared code and reserved fixed/editable task paths.
- Existing optional `learn_api.go` and workspace directory tests pass.
- The disposable learner integration tests run all three submission kinds with
  deliberately invalid Go text under `.idea`, complete corrected fixture code,
  assert metadata is absent from review-context export and remains unchanged.
- Fast `go test ./...`, race, vet and tagged integration pass.
- One full `task-check -all` collection covers 29 lifecycle packages. Phase1
  validation and deterministic report checking reuse it; generated reports
  remain unchanged. These are partial-bank checks, not full coverage acceptance.

Independent read-only agent review (`ui_goal_reviewer`) found no blocking
findings in the final boundary, regression tests, integration changes and ADR.
This is agent review, not human calibration.

GoLand.app was found on the development machine, but this batch uses metadata
fixtures and does not claim an actual IDE launch. Browser/launcher acceptance is
part of T02/K01. All test workspaces/homes are freshly temporary; no real learner
home, solutions or progress were accessed or changed for this batch. Task
inventory stays at 17 learner-visible ready v2 tasks plus three legacy ready
packages; eight drafts and one rejected package remain outside learner content.

Compatibility and non-destructive rollback are recorded in ADR 0012. No progress
schema, task contract, taxonomy or task content changes are included.
