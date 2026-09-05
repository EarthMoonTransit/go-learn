# ADR 0012: IDE metadata outside learner submissions

Status: accepted. Date: 2026-09-05. Goal batch: T01.

## Evidence and decision

A synthetic workspace accepted for each of implementation, tests and
workspace-patch becomes invalid as soon as GoLand creates `.idea/workspace.xml`:
inspection rejects the undeclared `.idea` directory before verification. The
learner must be able to retain ordinary project settings without changing the
submitted code or checker inputs.

Allow only the root `.idea/` metadata subtree in persistent learner workspaces.
Walk its entries and reject symlinks and nonregular files as for all workspace
content; `.idea` itself must be a directory. Do not follow links or skip the
subtree without inspection. No other unknown directories are allowed.

IDE metadata is ignored for submission selection and learner-code export, even
when a metadata filename ends in `.go`. It is never copied into the disposable
verification module, interpreted as code, or modified by workspace inspection.
Declared fixed files and editable paths retain their existing checks. A task
package cannot use this reserved workspace directory to carry a submitted file.

## Compatibility, migration and rollback

Progress stays at schema v1, and task schemas/provenance/check profiles are
unchanged. Existing workspaces without metadata or the optional editor API aid
remain valid. No workspace rewrite, progress migration or automatic cleanup is
needed. New GoLand metadata is permitted in place.

A code rollback restores the old rejection: users with `.idea` would have to
move that directory aside manually while using the old binary. Never delete IDE
settings as an automatic rollback or recovery action. Learner code and progress
remain usable and byte-compatible. Use matching application versions for CLI
and any later local UI.

## Acceptance

All-kind regression tests compare selected files before/after nested IDE
metadata and ensure the bytes stay intact. Adversarial tests retain unknown
path, fixed-file, root/nested symlink and nonregular-file rejection. Existing
optional API declaration compatibility remains covered. The disposable learner
integration scenarios verify all three submission kinds with invalid Go text
inside `.idea`, and ensure it is absent from review-context exports.

These tests establish metadata compatibility, not human calibration or proof of
an actual desktop GoLand launch. Launcher/browser acceptance belongs to T02/K01.
