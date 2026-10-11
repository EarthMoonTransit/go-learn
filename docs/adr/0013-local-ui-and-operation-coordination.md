# ADR 0013: Local learner UI and coordinated operations

- Status: accepted
- Date: 2026-09-05

## Decision

`learn ui` serves an embedded local interface through net/http on 127.0.0.1.
GoLand remains the editor. HTTP handlers consume explicit public views and call
the same Service operations as the CLI; no second checker, scheduler or store
is introduced. Read-only views do not select tasks, reveal hints or advance
session accounting. The existing taxonomy-based review remains a set of review
prompts, not an automatic assessment of learner code.

All mutating Service operations acquire a short-lived exclusive per-home
`.operation.lock` for their complete duration, including verification. The
existing progress-file lock is acquired inside that lease for atomic writes.
The UI does not lock the home for its lifetime. Contending operations fail with
a busy message. Unknown/stale locks are never deleted automatically. The owner
keeps its file descriptor open and checks lock identity before removing it.

UI mutations carry the intended task and a revision of persisted progress plus
workspace marker hashes. These preconditions are checked under the operation
lease. Each newly created workspace marker contains a random optional
`generation`, so reset/start of the same ID within one second cannot authorize
an old tab against a replacement workspace. The generation is not a secret or
an assessment result. Existing markers need no rewrite. Confined, bounded marker
reads cannot expose file contents through the revision.

The UI retains at most 16 process-local jobs and admits one operation at a time.
A request ID is bound to its complete action/task/revision payload; duplicate
requests return the existing job while retained. Shutdown and admission share
one mutex. Shutdown closes the listener and drains the admitted worker, allowing
the checker to finish and persist its actual result before exit. Forced process
termination cannot report a completed job; after restart inspect persisted
progress and recover any abandoned lock explicitly. Jobs do not survive restart.

## HTTP and IDE boundaries

The printed URL carries a random session token in its fragment. The browser
removes the fragment and uses a Bearer header for private reads and mutations.
Exact Host and Origin validation, POST-only mutations, bounded requests,
no-store responses and restrictive CSP protect the local API. There is no CORS,
arbitrary file/command/root/home API or reference reveal endpoint. Statements and
diagnostics render as text nodes. Full-check diagnostics and raw internal errors
are withheld. Public diagnostics come from the existing safe learner checker.

GoLand receives only validated server-derived paths as separate process
arguments. The executable and application location are discovered locally,
never supplied by HTTP. Failure leaves the copy-path workflow available.

The shared-project launcher (2026-10-11) passes the private `workspaces` root and
the selected lesson's first declared editable Go file. It validates the lesson
marker, real ancestor directories and regular target file before launch. The
GoLand project/file CLI syntax reuses the same project window; macOS forwards
arguments through the documented `open -n -a … --args` launcher and GoLand's
single-instance handoff. It does not change global IDE window preferences.
No shared Go module or `go.work` file is generated. Task workspaces, legacy paths,
markers and progress are unchanged; opening code never selects the active task.
Rolling back only restores the former per-lesson launch behavior. The shared
project's `.idea` directory is outside task roots and verifier/export inputs.

Launcher references: [JetBrains CLI documentation](https://www.jetbrains.com/help/go/working-with-the-ide-features-from-command-line.html)
and the installed GoLand 2025.3.3 `--help` project-plus-file syntax. Desktop smoke
used a disposable learner home with two lessons: both files opened in the same
`workspaces` window without changing the user's individual lesson project.

## Compatibility, migration and rollback

Progress remains schema version 1; attempts, hints, completion and session time
retain their existing meaning. No mastery migration or automatic repetition is
included. Use matching versions of UI and CLI with a home: an older binary does
not participate in operation coordination and must not run concurrently.

Verifier infrastructure errors exclude their waiting time under ADR 0009 while
preserving active/paused session state. They do not create a successful result,
completion or a counted verification attempt.

The optional workspace generation is backward-readable by the new implementation:
old workspaces without it remain usable and unchanged. Earlier binaries reject
new generation-bearing markers because their decoder disallows unknown fields.
Consequently rollback must retain this marker-reading compatibility change when
continuing with those workspaces. Alternatively retain the newer CLI while
disabling the UI. Do not reset learner workspaces or strip markers as an automatic
rollback operation. There is no progress migration to undo.

## Acceptance

Service regressions cover CLI/UI contention, hint selection, same-clock workspace
recreation, legacy markers, read-only persistence and lock ownership. HTTP tests
cover admission/idempotency, authorization, malformed requests, private data,
launcher boundaries and restart semantics. Disposable verification and browser
acceptance are recorded separately in the delivery report; agent review is not
human learning calibration.
