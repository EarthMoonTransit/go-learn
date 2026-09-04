# Learner trial readiness

Date: 2026-09-04  
Outcome: **READY_FOR_HUMAN_TRIAL** after the fixes in this batch.

## Changes and audit outcomes

- **Reset boundary — fixed.** `ResetTask` delegates deletion to
  `removeLearnerWorkspace` in `internal/learner/store.go`. The helper rejects
  symlinked home, parent, and task entries, verifies directory identity while
  opening, and removes the task through the opened `os.Root`. A replaced parent
  pathname cannot redirect the deletion. Missing/stale workspaces remain
  resettable; unrelated workspaces and progress remain intact. Rejected resets
  leave progress unchanged.
- **Failed code and export — fixed.** `internal/tasking/learner.go` separates
  filesystem inspection from code-contract validation. Syntax/API errors are
  failed attempts with bounded public diagnostics and exit code 1. Full results
  remain coarse. Export accepts broken learner code while rejecting modified
  fixed files, unexpected paths, and symlinks.
- **Session time — fixed.** Retries continue the active work session; repeated
  resume preserves elapsed time. Verification time is excluded. Pause,
  completion, and task switching end the interval. ADR 0009 describes progress
  v1 compatibility; previously unrecorded time is not reconstructed.
- **Original provenance — aligned.** AGENTS.md now distinguishes the legacy v1
  internal registry rule from the v2 project-identity contract already accepted
  in ADR 0008. No task or source records changed.
- **Stale lock — recovery documented.** The error identifies the lock path and
  the guide explains recovery after all learner commands stop. Automatic stale
  lock recovery is not implemented.

The default scheduler is unchanged. The trial explicitly starts with
`map-take-001`, then `detached-append-001`.

## Verification

Focused tests cover absolute/relative parent links, links inside learner home,
final and nested links, missing directories, stale markers, explicit reset
confirmation, traversal rejection, sibling preservation, and unchanged progress
on rejection. The original external-directory deletion no longer reproduces.
A separate read-only security review found no surviving concrete bypass or
compatibility regression in the reset boundary.

Integration tests cover broken code, useful saved diagnostics, safe export,
fixed-file rejection, and coarse full results for implementation, tests, and
workspace-patch submissions. They also exercise both actual ready tasks through
starter failure, passing public/full checks, feedback, completion records, and
canonical content-digest preservation. All learner state is temporary;
automated completions do not count as human calibration.

Local acceptance passed on Go 1.27.0, Darwin/arm64, in this order:

```text
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/task-check -root . -all -output <temporary-checks.yaml>
go run ./cmd/validate -root . -phase1 -checks <temporary-checks.yaml>
go run ./cmd/task-report -root . -check -checks <temporary-checks.yaml>
```

The single collection contains 15 packages, one session, 107 variant runs and
15 legacy commands (122 child Go runs). Deterministic quality reports are
current. The bank remains 5 ready packages, of which 2 are learner-visible v2.

Linux validation runs in the PR's required GitHub Actions `validate` job. No
claim of Windows or hostile-code sandbox validation is made. Concurrent path
replacement was reviewed against `os.Root` guarantees; the regression suite
uses deterministic path substitutions. Work time remains an explicit session
approximation with manual pause, not an idle detector.

## Next user action

Follow `docs/LEARNER_CLI.md`'s first human trial. Solve the two tasks independently,
record active time and feedback, and retain both review-context exports. No
further mandatory infrastructure phase precedes this trial. Repair the next
small content batch after the user feedback is available.
