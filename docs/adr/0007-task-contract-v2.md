# ADR 0007: Task Contract v2

- Status: accepted
- Date: 2026-09-04

## Context

Task schema v1 combines learner activity, submitted artifact and verification
commands. A `test-writing` task can therefore be checked accidentally as an
implementation task, while free-form command strings make the execution boundary
larger than the declared contract.

## Decision

Task Contract v2 separates three dimensions:

1. `activity` describes the capability exercised by the learner;
2. `submission.kind` defines the files the learner may change;
3. `verification.profile` selects checker-owned semantics.

The manifest is declarative. It may name only repository-relative Go files,
Go package patterns and bounded numeric limits. It cannot contain shell text,
external executables, network access, environment mutation or wall-clock
performance thresholds.

`fixed-tests`, `mutation-tests` and `workspace-tests` are executable profiles.
`fuzz-seed-tests` and `benchmark-contract` are reserved extension points: their
shape and limits are stable, but promotion fails until a checker implementation
is accepted separately.

The checker always assembles a fresh temporary module. Fixed and mutation
profiles use one generated `subject` package. Workspace verification copies only
the exact `submission.workspace_skeleton`, applies only `editable_paths`, then
adds checker-owned tests. `go.work`, nested modules, symlinks, native code and
submission changes to metadata, provenance or grader tests are rejected.

`task-check/v2` invalidates v1 check collections because they do not bind the new
schemas and policies. Task schema v1 remains supported and unchanged.

## Consequences

- Activity diversity can be reported without inferring it from commands.
- Test-writing submissions are run against a reference and required mutants.
- Workspace patches have an auditable write boundary.
- Adding an executable fuzz or benchmark profile requires a new reviewed tooling
  decision, not arbitrary commands in content.
