# Beginner collections batch 2

Date: 2026-09-05. Five original Task Contract v2 lessons complete the
12-lesson opening inventory.

| Task | Primary skill | Additional assessment | Required defects detected |
| --- | --- | --- | ---: |
| named-block-001 | types.named-underlying | arrays.value-semantics | 4/4 |
| array-copy-001 | arrays.value-semantics | none | 4/4 |
| slice-window-001 | slices.header-model | slices.len-cap-slicing | 5/5 |
| append-value-001 | slices.append-growth | none | 4/4 |
| map-presence-001 | maps.operations-comma-ok | none | 4/4 |

Independent read-only content review approved the final packages. It identified
and required two repairs: a compile-invalid always-fallback defect, and missing
input-preservation checks on rejected slice bounds. The latter now has a public
example and a dedicated clear-invalid defect. The slicing hint explicitly
distinguishes Go's capacity bound from this API's length restriction.

The named array introduction assesses both distinct type identity and assignment
copy behavior, unlike the legacy numeric-unit refactoring. Array replacement is
subsequent implementation practice. Append does not require storage isolation;
map lookup does not remove entries. The existing detached-append and map-take
tasks therefore remain distinct later exercises.

References, starter failures and all 21 required defects passed disposable checks.
Promotions used task-promote after independent approval. Full local acceptance
passed fast tests, race, vet and taskintegration; one 25-package collection was
reused by phase1 validation and deterministic report checks.

No human completion, hint, session or mastery record was fabricated. The initial
assisted map-take completion remains the only observed human completion.
Estimated times and the expanded route await human calibration. This batch is
partial content, not a release claiming full foundation coverage.
