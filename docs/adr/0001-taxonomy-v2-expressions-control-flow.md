# ADR 0001: Taxonomy v2 — expressions and control flow

- Status: accepted
- Date: 2026-09-04
- Supersedes: taxonomy version 1 coverage baseline

## Context

Taxonomy v1 contained 92 skills but jumped from declarations/functions directly
to data structures. It had no primary coverage target for arithmetic and boolean
expressions, comparison, `if`, `switch`, the general `for` statement, or the
general mechanics of `range`. The pinned Exercism Go concept inventory exposed
the gap, but upstream labels are evidence rather than the authority for this
decision.

## Decision

Set `curriculum/skills.yaml` to version 2 without renaming or changing the
meaning of any existing skill id. Add two domains and seven additive skills:

- `expressions.arithmetic`
- `expressions.boolean-logic`
- `expressions.comparison`
- `control-flow.if`
- `control-flow.switch`
- `control-flow.for`
- `control-flow.range`

Every new skill keeps the existing 4-primary/6-total minimum. The taxonomy now
contains 28 domains and 99 skills, with a baseline of 396 primary and 594 total
encounters. This remains compatible with the 350–500 task target because a
ready task contributes exactly one primary encounter.

Exercism concept mapping moves from hard-coded report logic to
`curriculum/mappings/exercism-go.yaml`. Mapping never creates task coverage.
Candidate-to-skill evidence counts as reviewed only when an explicit record in
`curriculum/evidence/` has `decision: accepted`.

## Migration

1. Change only the taxonomy version and add the two domains; preserve all v1
   ids and meanings.
2. Update validator baselines and tests from 26/92/368/552 to
   28/99/396/594.
3. Validate concept mappings and reviewed evidence against registry commit and
   taxonomy version.
4. Regenerate reports from versioned mapping data.
5. Keep existing task manifests unchanged. New skills receive coverage only
   through future, separately reviewed content batches.

`schemas/task.schema.json` is unchanged, so no task-manifest migration is
required.

## Consequences

The prerequisite graph becomes more explicit and future authoring can measure
control-flow coverage. Full `-coverage` remains red until the bank reaches its
target. Broad upstream labels such as `basics` and `numbers` remain unmapped;
they are not silently assigned to every plausible skill.
