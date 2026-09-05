# Local learner UI: T02 / K01 acceptance

Date: 2026-09-05. Tooling branch: `codex/local-learner-ui`.
Base: `11212a10bb9e56ebee64199e7c92c655692deb01`, after the accepted GoLand
metadata change in PR20. No content tasks are added by this batch.

## Delivered behavior

`go run ./cmd/learn ui` starts a foreground loopback UI. The existing local
checkout-based `learn` launcher also supports `learn ui` without modification.
The browser supports route selection, formatted statements, start/resume,
workspace copy and GoLand launch, public checks, full completion, hints, pause,
review prompts and the next lesson. UI and CLI use one Service/checker/store.

Progress stays v1. Existing workspaces remain unchanged; new markers include an
optional generation to invalidate old views after same-ID recreation. Matching
CLI/UI versions are required. Old binaries cannot read new generation-bearing
markers; rollback guidance is in ADR0013 and LEARNER_CLI.md.

## Evidence

| Requirement | Evidence and scope |
| --- | --- |
| Same learner rules and read-only views | Service coordination tests exercise public wrappers and CLI invocation while a verification is held; conflicting mutations fail busy, views preserve persisted bytes. |
| Same-ID replacement and legacy compatibility | A fixed-clock reset/start recreates identical progress but changes the revision; stale actions fail. A marker without generation remains usable and byte-identical. |
| Hints, time, locks | Concurrent hint requests choose distinct hints; replaced/unknown locks remain intact. Verifier infrastructure failures exclude waiting time for active/paused public/full cases without creating a result or completion. |
| HTTP boundary | Tests reject bad Host/Origin/token, cross-site requests, non-POST actions, extra fields, large bodies/URIs, traversal and private-artifact routes. Private responses use no-store; the public shell contains no home/root/token/task data. |
| Jobs and concurrency | Retained request IDs bind the complete payload, duplicate clicks return one job, conflicting tabs/CLI fail busy, retention is capped at 16, and a restart rejects old tokens/jobs while preserving progress. |
| Shutdown | A real loopback listener is cancelled during held full verification. New admission is rejected/closed; the server waits for the worker, persists its actual failed result and releases the lease. No persistent-job promise is made. |
| Actual checker paths | Tagged HTTP integration uses existing `literal-bindings-001` workspace-patch, `zero-results-001` implementation and `clamp-tests-001` test submission through disposable verification. |
| Safe hidden failure | For each kind, valid weak submissions pass public/reference checks and fail hidden grader or required-mutant checks. HTTP omits full diagnostics and internal identifiers; correct fixtures subsequently pass and complete. Compiler errors are checked separately. |
| IDE metadata | All three HTTP integration workspaces contain deliberately invalid Go under `.idea`; checks ignore it and preserve its bytes. T01 additionally verifies export exclusion and filesystem guards. |
| Unavailable launcher | A fake unavailable launcher receives only the validated server-derived path and returns copy-path guidance. Invalid workspaces never invoke it. |

All integration and browser work uses explicitly temporary learner homes. Test
solutions are fixtures copied into those homes, never the user's work. Authored
code runs only through the existing disposable verifier. This is automated
engineering evidence, not human learning calibration.

## Browser and GoLand smoke

A real Chrome browser in an isolated context completed:

1. Route → first statement → Start → workspace path.
2. Open in GoLand → the expected temporary project opened.
3. Public failure → corrected temporary fixture → public pass.
4. A double click produced one additional verification attempt.
5. Hint → full completion → visible taxonomy-based review → next lesson.
6. Copy path produced the browser's successful copy feedback.
7. Two tabs shared completion; starting the next task disabled check controls in
   the other tab's old-task preview until an explicit Resume.
8. An XSS-like Markdown fixture rendered literal text without active elements
   or script execution. No browser console errors/warnings appeared in the
   normal pre-restart flow.
9. Same-port server restart rejected the previous token. Opening the new URL in
   the same tab consumed the new fragment, reconnected and preserved completion,
   the used hint, review and the next task's paused state.

The restart smoke found a same-document fragment-navigation bug; a hashchange
reload fixed it and the scenario was repeated successfully. Authentication
failure now immediately disables action controls. The earlier UI smoke also
corrected estimated-time JSON field casing.

GoLand was actually launched on macOS and its project title and exact temporary
workspace path were observed. The project was opened in **Safe Mode**; trusted
project configuration, indexing/completion and IDE Run were not accepted by
this smoke. Checks ran through the UI's disposable verifier. The temporary IDE
window, browser tabs and UI process were closed after verification.

## Independent review and remaining gates

Independent read-only agent `ui_goal_reviewer` reviewed runtime boundaries,
tests, compatibility and the frontend delta. Same-ID stale views and misleading
route readiness were fixed. The requested real hidden-failure evidence was
added. Final agent review reported no blocking findings. This is not human
calibration.

The final ordered local run passed:

```text
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/task-check -root . -all -output .tmp/goal-ui/t02/task-checks.yaml
go run ./cmd/validate -root . -phase1 -checks .tmp/goal-ui/t02/task-checks.yaml
go run ./cmd/task-report -root . -check -checks .tmp/goal-ui/t02/task-checks.yaml
```

The full CheckCollection contains 29 lifecycle packages, SHA256
`4ea57ea27a5efde9acb8dd1ca31ab01a548f9155f6e3fcba725541a02113e9e4`.
It was executed once for this final tooling acceptance and reused by both
consumers; its deterministic result matches T01 because the task bank did not
change. The existing task-quality report remains current without regeneration.
The tagged learner integration package passed in 97.056 seconds. Partial-bank
validation was used; no full-foundation coverage or new task count is claimed.
Exact-head PR `validate` and post-merge `validate` remain mandatory process gates.

## Scope after this tooling batch

The bank remains 17 learner-visible ready v2 tasks plus three legacy ready
packages; eight drafts and one rejected package remain separate. C01 is pending:
exactly five strings/bytes tasks after UI acceptance. No mastery, progress v2,
dynamic repetition, theory, backend or second content batch is included.
The real learner home, solutions and progress were neither read nor changed for
this acceptance.
