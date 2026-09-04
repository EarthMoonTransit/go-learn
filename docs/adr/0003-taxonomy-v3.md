# ADR 0003: Taxonomy v3 — atomic skill contracts

- Status: accepted
- Date: 2026-09-04
- Supersedes: the taxonomy v2 coverage contract and the proposal in ADR 0002

## Context

Taxonomy v2 has 99 assignable skills and a 396-primary/594-total encounter
baseline. Its domain coverage is broad enough for Go fundamentals, but several
ids combine mechanisms that can fail independently, require different tests, or
need different prerequisites. Examples include slice storage and growth,
interface typed nil, error-chain construction and inspection, benchmark timing
and allocation reporting, channel close ownership, and context deadline
budgeting.

The repository already contains three ready tasks and reviewed source evidence.
A v3 change therefore needs an auditable migration rather than silent id
reinterpretation. Upstream Exercism labels are supporting evidence, not the
authority that determines taxonomy boundaries.

## Decision

Adopt taxonomy v3 with 28 domains, 142 active assignable skills, and 35
non-assignable deprecated compatibility stubs. Every active skill has a complete
teaching contract:

- a domain-local stable id and concise definition;
- active assignable state and taxonomy version;
- an acyclic graph of active prerequisites;
- at least two distinct misconceptions;
- at least three recommended exercise types;
- skill-specific mastery and observable evidence outcomes;
- all three evidence modes: implementation, debugging, and
  verification-or-integration;
- one of the reviewed encounter tiers: 2/4, 3/5, or 4/6.

The tiers produce a baseline of 409 primary and 693 total encounters. The
primary total remains inside the 350–500 ready-task target because a ready task
has exactly one primary skill. Total encounters include secondary skills and are
not a task-count target.

The authoritative migration is
`curriculum/migrations/v2-to-v3.yaml`. Immutable v2 taxonomy, mapping, and
evidence snapshots are stored under `curriculum/migrations/snapshots/` and
verified by SHA-256. Every one of the 99 v2 ids has exactly one migration entry.
Changed ids remain present as deprecated stubs whose `replaced_by` list exactly
matches the migration targets.

## Compatibility policy

Deprecated ids are not aliases and cannot be assigned to new tasks, mapping
targets, evidence, or prerequisites. A pre-existing task may use one only when
the migration contains an exact allowlist record keyed by task id, metadata
field, old id, and manifest SHA-256. Coverage behavior must also be explicit.
This prevents a compatibility exception from becoming a global alias.

The three existing task manifests use v2 ids that remain active with unchanged
meaning:

- `zero-state-001` → `types.zero-values`;
- `variadic-sum-001` → `functions.variadic`;
- `defined-types-001` → `types.named-underlying`.

They require no legacy resolution and remain byte-for-byte unchanged.

## Exercism mapping and evidence

Freeze the exact inventory of 38 upstream labels in
`curriculum/mappings/exercism-go-concepts.yaml`. Every label has one of four
states: `mapped`, `ambiguous`, `taxonomy-gap`, or `scope-deferred`. A mapped
label names exactly one active assignable skill; an ambiguous label lists at
least two active possibilities and never assigns coverage. `basics` and
`numbers` remain taxonomy gaps.

Reviewed evidence is re-pinned to taxonomy v3 with an explicit migration review
of the v2 evidence snapshot. Evidence remains `coverage_eligible: false`.
Accepted label-based evidence is allowed only for a one-to-one `mapped` label;
manual evidence must identify and hash the reviewed artifact.

## Deliberate non-splits

Not every paired phrase becomes multiple skills. The v3 review retains, among
others, array indexing/iteration, basic map operations/comma-ok, nil/delete/clear
maps, struct literals/zero values, struct copy/equality, field visibility/tags,
reader/writer contracts, file cleanup, test seams/doubles, and some toolchain
pairs. In these cases the v2 contract can still support three materially
different exercise forms without creating an independent learner-progress
decision. A future split needs new task evidence and a new ADR.

## Consequences

- Coverage planning is more precise, especially for ownership, lifetime,
  encoding, testing, and concurrency failures.
- The taxonomy file is larger because deprecated stubs preserve provenance.
- Full coverage remains intentionally red until later content phases.
- Reports and validators must distinguish active skills from compatibility
  stubs; counting every id is invalid.
- No task, draft, candidate record, source checkout, or UI is created by this
  decision.

The complete split/merge inventory and coverage deltas are generated in
`reports/taxonomy-v3-migration.md` and
`reports/taxonomy-v3-coverage-plan.md`.
