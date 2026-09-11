# Reviewed skill evidence

Candidate metadata never contributes curriculum coverage by itself. A
candidate-to-skill claim becomes reviewed evidence only through a record in
`curriculum/evidence/<source-id>.yaml` that passes
`schemas/skill-evidence.schema.json` and the semantic validator.

## Evidence identity

Every record is pinned by source id, exact source commit, exercise kind, slug,
source path, upstream concept, taxonomy version, and stable skill id. A review
also records its basis, strength, decision, reviewer, date, and rationale.

For taxonomy migrations the evidence file also records a reviewed migration
decision and the SHA-256 of the previous evidence snapshot. Evidence may target
only an active assignable skill. Deprecated ids and unreviewed global aliases
are invalid.

`upstream-concept-label` means the source explicitly declares the concept.
`manual-artifact-review` means a reviewer inspected a bounded candidate package;
it must record the reviewed artifact path and SHA-256 and must never be inferred
from an example solution alone.

## Decisions

- `accepted`: the record may guide future task authoring and mapping reports;
- `rejected`: the apparent association was reviewed and found unsuitable;
- `needs-review`: no curriculum claim may be made yet.

`coverage_eligible` is fixed to `false`. Only a validated ready task contributes
primary or secondary coverage. Evidence does not create tasks, assign task
skills, or authorize copying source material.

Changing the source commit or taxonomy version invalidates the review until the
record is re-reviewed. Ambiguous concept mappings cannot create accepted direct
evidence without an exercise-specific manual review.

The frozen upstream inventory lives in
`curriculum/mappings/exercism-go-concepts.yaml`. A label-based accepted record is
valid only when that label has a one-to-one `mapped` entry. `ambiguous`,
`taxonomy-gap`, and `scope-deferred` labels cannot silently choose a skill.

## Theory and reconsidered candidates

Read-only topics and their task links provide navigation, never assessed skill
encounters or learner evidence. Reconsidering an algorithmic rejection records
the old decision, new pedagogical disposition and remaining lineage gate; it
does not automatically accept the candidate or permit reference-only artifacts.
See the [active plan](plans/foundation-next.md) and ADRs 0020–0021.
