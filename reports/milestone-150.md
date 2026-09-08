# Milestone 1: 150-lesson acceptance

This checkpoint becomes delivered when the C14 content PR and final tooling
acceptance PR pass exact-head `validate` and are merged. It is course readiness,
not evidence that the learner completed or mastered these lessons.

## Inventory and lineage

The route contains 150 ready Task Contract v2 lessons across the nine explicitly
numbered foundation sections. All 128 additions from the 22-lesson baseline are
bound to actual task IDs in `catalog/seed-plans/foundation-150.json`; prerequisite
closure includes 101 skills, of which 99 are assessed and two are preparation.
The current taxonomy has 142 skills, so full foundation coverage is later work.

The ready set has 84 substantive adaptations, 61 originals and five concept
seeds. Only the 84 adaptations count toward the minimum 76. Exact source commits,
paths, artifact digests and license notices are retained. The saved 177 Exercism
candidates and explicit Learn Go with Tests/Ardan selections informed delivery;
reference-only sources and generated canonical cases were not copied.

Activities: 75 implementation, 40 test-writing and 35 debugging. Components
include a generic stack, a multi-package inventory, file record parsing and
cleanup, a recording writer, worker joining and context composition. Reference
solutions are excluded from learner workspace payloads.

## Product and preservation evidence

The isolated final smoke check reports `150/150 defined` through the CLI and
`150 ready lessons` through the rendered local UI. Sections 01–09 are present;
new learner progress starts with zero completions, zero confirmed skills and no
reviews due. CLI list, roadmap, next, progress and organize dry-run all succeed
without creating files in a fresh learner home. Existing live learner state is
not read or modified for this acceptance.

Numbered workspace organization and progress v3 have separate migration,
backup/recovery, conflict and repeated-run regressions. Stable task IDs preserve
solutions, completion, attempts, assistance and time. Delayed reviews and two
independent practice/transfer encounters are separate from completion; synthetic
fixtures never count as human calibration.

## Verification and remaining scope

Each content batch has independent review and isolated public/full/reference/
defect checks. Cross-batch audits are recorded for C02–C04, C05–C07, C08–C10 and
C11–C13. Reviewed public-feedback gaps have dedicated learner Check regressions.
ADR0019 makes race-required tasks run public and defect verification under the
race detector; stale v2 collections cannot certify the v3 execution contract.

Acceptance uses fast tests, race, vet, integration and one complete collection
reused by `validate -phase1 -checks` and `task-report -check -checks`. Full
foundation `-coverage`, fuzzing/benchmark profiles, remaining skills and required
encounter depth belong to milestone 2. Backend/production lesson topics remain
deferred. Human completion/mastery counts remain unknown rather than inferred.
