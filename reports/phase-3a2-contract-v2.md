# Phase 3A.2 — Task Contract v2

## Outcome

Implemented a backwards-compatible task/provenance contract split without adding
learning tasks or changing legacy ready packages.

## Contract surface

- Task schema v1 remains byte-identical and readable.
- Task v2 separates activity, submission kind and verification profile.
- Provenance v2 supports original, adapted and combined multi-source records.
- `task-init` defaults to v2; `-schema-version 1` retains the legacy scaffold.
- `task-check` auto-selects v1/v2 from `schema_version`.
- `task-promote` preserves and validates the selected schema.
- Check collections now use `task-check/v2`; v1 collections are stale by design.

## Executable evidence

Synthetic fixtures cover fixed implementation checks, learner-authored mutation
tests, and allowlisted multi-package workspace patches. Adversarial tests cover
unconditional test failure, grader-path modification, forbidden reference-only
copying and stale collections. Integration-tagged tests execute all three profiles
in checker-created temporary modules.

## Deferred intentionally

Fuzz and benchmark profiles have bounded schemas but no promotion-capable runner.
No pilot task, real draft/ready task, UI or new source bank was created.

## Acceptance

All commands passed on 2026-09-04:

- `go test ./...`
- `go test -race ./...`
- `go vet ./...`
- `go test -tags=taskintegration ./...`
- one `task-check -all` collection consumed by both `validate -phase1 -checks`
  and `task-report -check -checks`
- `go run ./cmd/validate -root . -phase1 -run-commands`

Protected inputs remained byte-identical: all three `tasks/ready` package tree
digests, `schemas/task.schema.json`, `curriculum/skills.yaml`, reviewed Exercism
candidates and the pilot seed plan match their pre-batch SHA-256 values.
