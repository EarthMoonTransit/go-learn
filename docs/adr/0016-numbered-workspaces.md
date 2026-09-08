# ADR 0016: numbered private workspaces and recoverable organization

Status: accepted. Date: 2026-09-08. Implements the accepted foundation roadmap.

Progress v2 retains every learning field and adds no mastery claim. Its version
permits numbered workspace paths assigned by route v2: workspaces/<section number>
-<block ID>/<lesson number>-<task slug without its final numeric suffix>. IDs in
progress and workspace manifests remain unchanged. Missing/v1 routes use legacy
paths. Fresh homes use v2; v1 homes remain readable and writable at legacy paths
until explicit organization. Read-only commands never upgrade data.

`learn organize --dry-run` returns version/path changes without writing data.
`learn organize` acquires the shared UI/CLI operation lease, validates the full
move plan, snapshots original progress bytes and complete affected workspace
trees under a private backups/organize-* directory, then publishes a versioned
journal. Files are read as data, never executed. Backup rejects symlinks, special
files and individual workspaces over 128 MiB before moving them. File bytes and
permissions are preserved, including IDE metadata and workspace markers.

Each rename and the eventual progress save are recoverable. Reinvocation checks
before/after progress, source or destination tree digests, backup integrity and
current route assignments. Conflicts, edited files or changed history stop the
operation without deleting either copy. The journal is cleared only after the
new progress has been saved. Ordinary mutations refuse while it exists; read-only
views may report temporarily unavailable workspaces. Already organized homes are
no-ops. The same confinement/identity rules cover new paths, reset and UI revisions.

After an actual process kill, existing operation/progress lock recovery rules
still apply: establish that the process is stopped before removing stale locks,
then rerun organize. Never automatically steal a lock. Reopen moved workspaces in
GoLand. The tool does not silently move real learner data during development.

Rollback is explicit and offline: stop learner processes, retain current copies,
restore original progress.json and workspace trees from the named backup, then
use compatible code. Backups are never automatically deleted. A backup created
before a pre-journal failure may remain as recovery evidence.

Tests cover completed/in-progress history, byte-identical backups and markers,
preview/no-op behavior, partial renames, progress-save interruption, changed
recovery data, conflicts, symlink ancestors and UI/CLI lease exclusion. Existing
real UI integration flows run all three submission kinds after organizing legacy
workspaces; fresh-workspace integration exercises new paths as well.
