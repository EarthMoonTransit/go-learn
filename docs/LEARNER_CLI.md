# Local learner CLI

`learn` is a local, text-only workflow for ready Task Contract v2 packages. It
does not expose legacy task packages, authoring lifecycle states, or hidden
verification material.

## Quick start

```text
go run ./cmd/learn list
go run ./cmd/learn show detached-append-001
go run ./cmd/learn start detached-append-001
# edit the printed workspace path
go run ./cmd/learn check
go run ./cmd/learn complete
```

Global `-root` and `-home` flags precede the command. `-root` defaults to the
current directory. Learner state defaults to the OS user-config directory under
`go-practice-bank/learner`; `GO_LEARN_HOME` or `-home` can override it.

## Commands

- `list` lists ready v2 tasks with local status, activity, difficulty, and
  primary skill.
- `show <task-id>` prints learner-visible metadata and `TASK.md`.
- `start <task-id>` creates a private workspace and refuses to replace existing
  progress.
- `resume <task-id>` selects and validates an existing workspace.
- `check` runs only public verification and records one attempt.
- `complete` runs full verification in temporary storage. A passing result
  records completion and skill encounters.
- `hint` reveals the next unused numbered hint.
- `review` is available only after completion and shows taxonomy-based review
  prompts, not the reference solution.
- `solution --reveal` is the separate, explicit solution-reveal action. Reveal
  is recorded but never marks completion.
- `next` selects one unfinished task with the deterministic policy below.
- `progress` prints attempts, hints, reveal/completion state, active duration,
  feedback count, and skill encounters.
- `feedback <text>` stores up to 2,000 characters locally for the current task.
- `reset-task [task-id] --yes` removes only that task's local workspace and
  progress. Confirmation is mandatory.
- `export-review-context` writes a private Markdown file under `exports/`.

`check` and unsuccessful `complete` use exit code 1. Invalid usage or unsafe
workspace state uses exit code 2.

## Persistent workspace contract

The canonical package under `tasks/ready` is read-only input. `start` creates
`<learner-home>/workspaces/<task-id>` atomically. It contains `TASK.md`, a fixed
`go.mod`, a private manifest, and only the code needed by the submission:

| Submission kind | Editable learner content | Fixed learner-visible content |
| --- | --- | --- |
| `implementation` | starter Go files under `subject/` | `TASK.md`, `go.mod` |
| `tests` | starter `*_test.go` files under `subject/` | `TASK.md`, `go.mod` |
| `workspace-patch` | exactly `submission.editable_paths` | `TASK.md`, `go.mod`, declared workspace skeleton |

For implementation and test submissions, additional Go files are allowed only
inside `subject/` and must preserve the declared submission shape and public
API. Workspace-patch submissions may change exactly the declared editable
paths. Symlinks, devices, undeclared directories, modified fixed files, and
path escapes are rejected.

Public tests are assembled by `check` in temporary storage; they are not copied
into the editable workspace. Grader tests, verification references, reference
solutions, and mutants are never copied there.

## Verification boundary

Every check creates a fresh temporary Go module with isolated build, module,
home, and temporary directories. `GOWORK=off`, `GOPROXY=off`,
`GOTELEMETRY=off`, and read-only module mode prevent workspace inheritance,
dependency downloads, and Go telemetry. The canonical task and learner
workspace are inputs; commands execute only in the temporary copy.

`check` returns bounded diagnostics produced by public tests. `complete` can use
the full grader, required defect models for test submissions, determinism runs,
`go vet`, and a race check when the task contract requires them. Its output is a
coarse pass/fail category and never includes hidden test names, cases, paths, or
mutant identifiers.

This is process and filesystem hygiene for code authored by the local learner,
not an operating-system sandbox for hostile code. The CLI must not be used to
execute submissions obtained from an untrusted party.

## Local progress

`progress.json` is a private, atomically replaced local file. It stores task
status, attempts, numbered hints used, solution reveal, completion, accumulated
active-session seconds, skill encounters, public-check output, and feedback. An
active session starts at `start`/`resume` and ends at `check`/`complete`, so the
duration is an explicit-work-session approximation. No grader, reference,
mutant, hidden expectation, telemetry identifier, or cloud state is stored.

## Deterministic `next`

Completed primary and secondary skill encounters satisfy prerequisites. The
scheduler excludes completed tasks, then sorts by:

1. all prerequisites encountered before missing prerequisites;
2. fewer missing prerequisites;
3. lower difficulty;
4. fewer occurrences of the activity among the three most recent tasks;
5. larger deficit against the primary-skill minimum;
6. task ID as a stable final tie-breaker.

If no task has all prerequisites recorded, `next` returns the closest candidate
and explicitly lists the missing prerequisite encounters.

## Review-context export

The Markdown export contains `TASK.md`, current learner code, the last public
check result, primary and secondary skill metadata, and only the hints already
used. Dynamic Markdown fences prevent learner code from escaping its code
block. The export omits grader tests, reference solutions, mutants, internal
paths, and hidden expectations. It is suitable for deliberate manual pasting;
the CLI itself performs no network request.
