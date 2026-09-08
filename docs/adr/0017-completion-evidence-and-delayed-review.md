# ADR 0017: completion evidence and delayed skill review

Status: accepted for implementation, 2026-09-08.

## Decision

Progress schema v3 adds an append-only first-completion evidence list. Each
entry captures stable task ID, completion time, assessed skills, route role and
assistance at the successful full check. A durable per-task assistance map
survives reset, including hints or reveals before any successful completion. Introduction, hinted, solution-revealed
and legacy results remain progress; mastery requires two distinct independent
practice/transfer task IDs for the skill. Viewing a solution afterwards does not
rewrite an earlier result. Resetting a task retains its first evidence and cannot
manufacture another independent encounter; task workspace/history reset keeps
its existing explicit confirmation requirement.

After an encounter, the first review is due after three other distinct task
completions. `next` prefers a ready, uncompleted practice/transfer task assessing
a due skill, only when its actual prerequisites are already known. An ongoing
current lesson is preserved. No suitable review means normal route progression
continues and the due count stays visible; missing content is not fabricated.
An assisted review schedules another independent opportunity. Two qualifying
tasks confirm the skill; this phase does not impose an endless review queue.
UI, roadmap and progress show completion and confirmed mastery separately.

## Migration and compatibility

Fresh homes use v3. Existing v1/v2 files stay byte-for-byte unchanged during
read-only views. `learn organize --dry-run` previews v1/v2→v3; `learn organize`
uses ADR 0016's exact progress backup, workspace backup (when moves are needed),
journal and replay. Already-numbered v2 homes require zero directory moves.
No old completion is retroactively marked independent: legacy dates order review
scheduling evidence only, because assistance at historical completion is unknown.
Migration freezes legacy completion identities and known assistance before
any reset can remove old task progress. Until explicit migration, v1/v2 commands preserve old behavior;
UI/CLI show the upgrade instruction and make no new mastery claim.

An unfinished ADR 0016 v2 journal remains recoverable under v3 code. Recovery
first finishes its declared target v2 without rewriting history; a subsequent
organize invocation upgrades to v3. Read-only commands never recover or migrate.
Rollback before further learning restores the exact backed-up progress and,
when applicable, workspace paths using ADR 0016's documented procedure. An old
binary must not open v3 state; it rejects the unsupported version.

## Verification

Tests distinguish independent/assisted/introduction/legacy results, reject
same-task double counting, check exactly three intervening completions and
review prerequisite eligibility, snapshot assistance at success, and migrate a
numbered v2 home without moving files. Existing interruption, conflict, backup,
read-only, progress-history and three-profile UI integration checks remain.
Real learner progress is never simulated or rewritten by course acceptance.
