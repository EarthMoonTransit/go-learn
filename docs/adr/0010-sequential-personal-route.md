# ADR 0010: Sequential personal route

Status: accepted. Date: 2026-09-05.

The user chose a sequential English course in GoLand, without an entry
diagnostic. The generic scheduler can otherwise choose an advanced task when
the bank is incomplete. A ready-task count also hides missing learning blocks.

Add version 1 curriculum/route.yaml, independently versioned from task and
progress contracts. Blocks contain ordered task IDs, explicit assessed skills
(primary first) and introduction/practice/transfer roles. Missing task packages
are planned content, never available lessons. Validation checks skill dependency
order and binds ready task metadata to route metadata. Empty blocks represent
later milestones; they cannot be silently skipped by next.

The first toolchain preparation is a named English guide. It makes the
unrecorded prerequisite visible and grants no task completion, encounter or
mastery. The route may recommend the first guided exercise with that explicit
preparation warning. Explicit start/resume remain available.

learn roadmap is a read-only view. With a configured route, learn next picks
the first unfinished lesson and reports a content gap if it is unavailable.
Existing completions, including tasks completed out of order, are preserved.
Without a route, the existing generic scheduler remains available for compatible
fixture repositories and older course checkouts.

No task schema, taxonomy IDs, encounter semantics or progress version changes.
No progress migration is required. The rollout adds the route file; removing
it restores the generic selection policy without changing user data.

This batch does not implement mastery or spaced repetition. Their acceptance
gate remains the working-course milestone in ROADMAP.md. A completion or route
role alone cannot be presented as independent mastery.
