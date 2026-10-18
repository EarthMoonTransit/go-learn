# P20 active-plan integration audit

Preparation evidence captured on 2026-10-18 in `codex/p20-fuzz-seeds-invariants`,
against base `5ef505c0776bced18c6b2e5dbd18ef04944877d7`. This is a branch-local
integration checkpoint, not release/export, foundation acceptance or human mastery.

## Problem and correction

The address registry contained 171 lessons while the active pedagogical CSV still
contained 169. Catalog/address preservation alone did not prove scheduling readiness.
The actual learner API reported both missing IDs and rejected `Next` before the fix.

The corrected v2 mapping places `hex-seed-classes-001` and
`hex-wire-invariants-001` at execution positions 151/152, directly after
`text-level-001` in C12 encoding and before C13 time. Their stable addresses are
08.05/08.06. The previous 151..169 execution positions shift by two; every old
address, ID, relative order, chapter, subtopic, role, disposition and rationale is
preserved. The original historical review CSV remains untouched. One reviewed
successful-exposure edge connects the first lesson to the second.

## Actual automatic regression

The trusted metadata-only drivers are retained as
`.tmp/p20-fuzz-content/active-plan-proof.go` and `late-current-proof.go`.
They call the real learner API with explicitly created temporary homes; no authored
task or test code runs from the source checkout. Completion data used for selection
scenarios is explicitly synthetic and assisted/legacy; it is not learner evidence.

| Probe | Observed result |
|---|---|
| Before: real PedagogicalPreview and Next | Exactly two missing current route IDs; active scheduler rejected |
| After: configured execution | 171 lessons, preparation-ready plan |
| Old route/catalog records | All 169 stable tuples and public task records unchanged |
| Existing unfinished workspace and progress | Exact file/mode/history/clock snapshots retained by readonly operations |
| Separate late current 09.16 while earlier lessons remain unfinished | Resume priority retained, ahead of ordinary earlier lessons |
| Synthetic old course completed | Next selects seeds at position151/address08.05 |
| First starter has only a passing public check | Next still selects the first lesson; no completion is inferred |
| First lesson has a synthetic completed encounter | Next selects invariants at position152/address08.06 |
| Cleanup | Every temporary regression home removed |

The existing structural unit regression for a new route task missing its active
mapping remains applicable. These actual-data before/after probes cover this
specific configuration and preserve replayable source/results in the goal evidence.

## Same collection, explicit metadata delta

The one full collection completed before this integration correction: 181 task
checks in 968.976 seconds, 2243 generic variant runs, 68 native fuzz runs,
2345 disposable workspaces and one runtime query. Original frozen inputs stayed
unchanged during collection and its three report-generation/four check consumers.

Collection SHA-256:
`c737cffdbd4656607a2929fd733f973045f7774f9828b386447fef0f611d4daf`.

Only three non-task metadata/document files changed afterward:

- `curriculum/plans/pedagogical-order-v2.csv`
- `curriculum/plans/pedagogical-order.yaml`
- `docs/ROADMAP.md`

No task package, authored content digest, verification code, schema, taxonomy,
source registry or toolchain policy changed. Actual Go `TaskInputDigest` equality
was checked for all 181 packages, both against the pre-correction snapshot and
the exact saved collection. The verifier/runtime identity stayed current.
Validator phase1, task-report check and both numeric coverage checks passed again
using that same file; no full task-bank execution was repeated. The original
71-file source review is not represented as review of this later metadata delta.

## Evidence hashes and remaining gates at capture

- `active-plan-proof.go`: `297d0dfd22fc7bcc15111f3511dcd839779878beb9ab5ae70492fbf1badc4a46`
- `active-plan-before.json`: `f448f9e665b3e59b4ba1490ac59c3abb23c5840b094e3c5a098549269305838b`
- `active-plan-after.json`: `c87929966de50fbe8ed37850f93c3f306fc63c8bacee21c959832469be4ccff3`
- `late-current-proof.go`: `1831183585daeb2aa89f01d71883d2f2dcd291608752b093000d6ba6a146fe47`
- `late-current-proof.json`: `6618b5bc67c6da11cea31579af790e6573011b912a00aac69f6540ae7d00bdbf`
- `post-plan-consumers.json`: `c21b31c96a1daccfd9f683c2a089546a3d73b4c705f66ee07ba445da08c893e8`

The active CSV retains its existing CRLF line endings. Whitespace lint uses
`git -c core.whitespace=cr-at-eol diff --check` to recognize that existing CSV
format; this does not bypass any commit/push hook or acceptance gate.

Final ordered repository fast/race/vet/integration, final exact-head independent
review and exact-head green GitHub validate were still required at this capture.
The collection checks numeric report freshness, not complete foundation coverage.
The reported foundation gaps remain 243 primary / 494 total occurrences.
