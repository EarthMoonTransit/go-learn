# Learner CLI acceptance

Date: 2026-09-04  
Scope: Phase 3D local learner CLI only

## Result

The repository now provides `go run ./cmd/learn ...` for ready Task Contract v2
packages. No task package, taxonomy record, source registry entry, schema, or UI
was changed. The learner state is local and contains no telemetry or cloud
integration.

## Command surface

The implemented commands are `list`, `show`, `start`, `resume`, `check`,
`complete`, `hint`, `review`, `solution`, `next`, `progress`, `feedback`,
`reset-task`, and `export-review-context`. Integration coverage exercises every
command, including explicit gates for solution reveal, post-completion review,
and destructive reset.

## Submission profile matrix

| Synthetic profile | Learner submission | Public baseline | Full baseline | Reference-equivalent submission |
| --- | --- | --- | --- | --- |
| `fixed-tests` | implementation files | fails | fails safely | passes |
| `mutation-tests` | learner test files | passes correct behavior | fails safely when a defect model survives | passes |
| `workspace-tests` | allowlisted workspace patch | fails | fails safely | passes |

The integration fixtures are created in test temporary directories from
`internal/testdata`; no fixture enters `tasks/ready` in the source checkout.

## Boundary assertions

- canonical task content digests remain unchanged across start, public check,
  full completion, review, feedback, and export;
- persistent workspaces contain no grader, reference, mutant, or verification
  artifacts;
- a solution reveal returns content only after `--reveal`, records the reveal,
  does not persist the solution in the workspace, and does not complete a task;
- full-check results contain no diagnostic body or mutant identifier;
- review exports contain the required learner-visible sections and no known
  hidden fixture identifiers;
- fixed workspace files, path boundaries, and symlinked private directories are
  rejected;
- progress round-trips locally without checker-owned fields;
- repeated scheduler calls over unchanged state select the same task;
- Markdown export fences remain valid when learner code contains backticks.

## Acceptance commands

The following commands are the Phase 3D acceptance graph:

```text
go test ./...
go test -race ./...
go vet ./...
go test -tags taskintegration ./...
go run ./cmd/task-check -root . -all -output <temporary-checks.yaml>
go run ./cmd/validate -root . -phase1 -checks <temporary-checks.yaml>
go run ./cmd/task-report -root . -check -checks <temporary-checks.yaml>
```

Final local results:

| Command | Result |
| --- | --- |
| `go test ./...` | pass |
| `go test -race ./...` | pass |
| `go vet ./...` | pass |
| `go test -tags taskintegration ./...` | pass |
| one full `task-check -all` collection | pass; 15 lifecycle packages checked |
| `validate -phase1` reusing that collection | pass; 5 ready manifests, 142 assignable skills, 12 audited sources, 177 reviewed candidates |
| deterministic `task-report -check` reusing that collection | pass |

The source checkout still contains 5 ready manifests. The temporary learner
fixtures and their three submission profiles were never added to that count.

## Known boundary

The disposable Go environment prevents source-checkout execution, workspace
inheritance, dependency downloads, and Go telemetry. It is not a kernel-level
sandbox. The intended input is code written by the local learner, not hostile
third-party code.
