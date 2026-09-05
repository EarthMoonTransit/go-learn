# Local learner CLI

For test-writing exercises with supported simple APIs, new workspaces include
a read-only `subject/learn_api.go` so GoLand can resolve the tested functions.
Write your tests in the editable test file and run `learn check` or
`learn complete`. The editor declarations deliberately have no working
implementation and are excluded from verification. Older workspaces without
this file still work. Complex APIs retain the checker-only workflow described
in ADR 0011 until their declaration support is added.

`learn` currently provides a local CLI workflow for ready Task Contract v2 packages. It
does not expose legacy task packages, authoring lifecycle states, or hidden
verification material.

## Sequential personal route

The repository now ships curriculum/route.yaml. Use learn roadmap to see the
current block, ready lessons, completed lessons and planned future content.
Use learn next for the first unfinished route lesson and its start command.
See [LEARNER_START.md](LEARNER_START.md) for the initial GoLand preparation.

The preparation guide grants no completion or mastery. Existing completions
remain recorded even when their task appears later in the route. Missing lessons
stop automatic recommendations at the gap; explicit start/resume still work.
Read-only route commands do not select a task or change session time.

## Quick start

```text
go run ./cmd/learn list
go run ./cmd/learn next
go run ./cmd/learn show literal-bindings-001
go run ./cmd/learn start literal-bindings-001
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
- `pause` records elapsed work and pauses the timer before a break.
- `check` runs only public verification and records one attempt, including syntax
  and public-API errors in learner code.
- `complete` runs full verification in temporary storage. A passing result
  records completion and skill encounters.
- `hint` reveals the next unused numbered hint.
- `review` is available only after completion and shows taxonomy-based review
  prompts, not the reference solution.
- `solution --reveal` is the separate, explicit solution-reveal action. Reveal
  is recorded but never marks completion.
- `next` selects the first unfinished route lesson; repositories without a route use the fallback policy below.
- `roadmap` shows the sequential course and local completion state.
- `progress` prints attempts, hints, reveal/completion state, active duration,
  feedback count, and skill encounters.
- `feedback <text>` stores up to 2,000 characters locally for the current task.
- `reset-task [task-id] --yes` removes only that task's local workspace and
  progress. Confirmation is mandatory.
- `export-review-context` writes a private Markdown file under `exports/`.

`check` and unsuccessful `complete` use exit code 1. Syntax and public-API
errors are failed attempts with bounded public diagnostics. Invalid usage or
unsafe workspace state uses exit code 2.

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
path escapes are rejected. Root `.idea/` settings created by GoLand are allowed
as private IDE metadata: they are inspected for symlinks/nonregular files but
never submitted, executed, exported or rewritten. `.idea` must be a directory;
other unknown directories remain forbidden. Task content cannot declare paths
inside this reserved metadata directory. See ADR 0012.

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
active session starts at `start`/`resume`, continues across checks and failed
completion attempts, and ends at `pause`, successful completion, or selection
of another task. Repeating `resume` for the active task preserves elapsed work.
Verification time is excluded. A check while paused leaves the timer paused.
`progress` includes elapsed time in the active session without closing it.

Use `pause` before a break and `resume <task-id>` when returning. There is no
automatic idle detection. Existing progress remains readable; old missing time
is not reconstructed. See ADR 0009 for the compatibility contract. No grader,
reference, mutant, hidden expectation, telemetry identifier, or cloud state is
stored.

Mutating commands share a per-home `.operation.lock` with the UI, including
the whole duration of verification. A conflicting command reports busy; retry
after the active command finishes. The UI does not hold this lock while idle.
Read-only commands and page views do not change selection or work-session time.

If a process is forcibly interrupted, `.operation.lock` or `.progress.lock`
can remain. Nothing removes an unknown lock automatically. First stop every
UI/CLI process using this home and verify that no check is still running. Only
then remove the named lock file and retry. Its PID is a clue, not proof of safe
removal (PIDs can be reused). Keep `progress.json` and all workspaces. Never
remove a lock while another process may own it.

## Generic fallback without a route

When curriculum/route.yaml is absent, completed primary and secondary skill encounters satisfy prerequisites. The
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
the CLI itself performs no network request. Syntax and API errors do not block
exporting learner-owned code; filesystem boundaries and fixed files are still
validated.

## Historical first human trial

This is the original two-task pilot, retained as historical evidence. New
learners follow `next` and the sequential route described above. Automated
UI acceptance uses a separate temporary home and does not require a human trial.

Original pilot instructions: use `map-take-001` first, then `detached-append-001`. The scheduler is a generic
policy; with an empty history it may recommend the second task first because it
has fewer unrecorded prerequisites.

For each task, run `show`, then `start` (or `resume` if already started), and edit
only the printed workspace. Iterate with `check`; use `hint` when needed and
`complete` when public checks pass. After completion run `review`, then:

```text
go run ./cmd/learn feedback "clarity=.../5; difficulty=.../5; actual=...m; public-feedback=.../5; learned=...; frustration=...; next=..."
go run ./cmd/learn export-review-context
```

Measure active time independently during this first trial. Keep both exports and
the feedback. Do not inspect grader tests or reveal a reference solution before
your attempt. A task you cannot complete is still useful trial evidence: record
where you stopped and export the current code. Automated smoke-test completions
are not human feedback and do not count as calibration.

## Local UI

From the repository checkout:

```bash
go run ./cmd/learn ui
```

With the existing checkout-based short launcher, use `learn ui`. If your `learn`
is a separately installed binary, rebuild that binary from this version before
using it; a local build that does not overwrite any launcher is:

```bash
go build -o .tmp/learn ./cmd/learn
.tmp/learn ui
```

The foreground process prints its effective root, home and a URL. Open the full
URL, including the session fragment. The browser removes the fragment from the
address bar. Keep the process running; Ctrl+C closes the listener and waits for
an active check to persist its actual result. Reopen the newly printed URL after
restarting. Jobs and their display results are temporary; progress, completion
and used hints persist. For a fixed local port use `learn ui -port 8080`.
Global `-root`, `-home` and `GO_LEARN_HOME` follow the same rules as other commands.

Select a lesson and click Start, then Open in GoLand or Copy path. Open the
printed workspace as a project and edit the files named in the task. Check runs
public tests; Complete runs full verification and records completion on success.
The Review prompts appear after completion. Next lesson follows the route;
unprepared content is labelled explicitly. Browsing another lesson does not
resume it. Use Resume to switch the active workspace and Pause before breaks.

Used hints are saved. Reference reveal remains the explicit CLI command
`learn solution --reveal`; the UI has no solution endpoint. Review prompts are
based on the task's skills, not an automated analysis of your code.

Use matching UI and CLI versions with the same home. Progress remains v1 and
existing workspaces remain readable without rewriting. Newly created workspace
markers include a generation that older binaries cannot read. For rollback,
keep the new CLI (you can stop using the UI), or retain the marker compatibility
change. Do not reset workspaces as an update/rollback procedure. Details:
[ADR 0013](adr/0013-local-ui-and-operation-coordination.md).

The UI listens only on 127.0.0.1. Keep the session URL private. No accounts,
hosting, mastery migration or dynamic repetition are included in this delivery.
