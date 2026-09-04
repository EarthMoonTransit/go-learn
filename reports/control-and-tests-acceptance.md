# Control flow and comparison verification — batch 3

Date: 2026-09-05. Four original lessons and one substantially repaired draft.

| Task | Primary skill | Activity | Required defects detected |
| --- | --- | --- | ---: |
| bounded-value-001 | control-flow.if | implementation | 4/4 |
| wide-sum-001 | types.numeric-conversions | debugging | 4/4 |
| encoded-width-001 | control-flow.switch | debugging | 4/4 |
| range-offset-001 | control-flow.range | debugging | 4/4 |
| clamp-tests-001 | expressions.comparison | test-writing | 6/6 |

Independent read-only review approved the final packages. Its findings repaired
the clamp interior defect so exact-endpoint assertions cannot substitute for an
interior case, corrected its provisional difficulty/time metadata and placed
three intervening tasks before comparison verification. Direct prerequisites
and all 22 required defects passed targeted disposable checks. The new block
follows the existing twelve lessons without reordering them.

The clamp package retains its Exercism counter concept-seed record and earlier
rejection history. Its metadata artifact SHA, exact commit/path, contributors
and MIT notice were checked against the pinned cache. Only metadata and LICENSE
were authoring/review inputs; no upstream exercise statement, implementation,
tests, canonical cases or solution was used. The remaining four tasks are
original under ADR 0008. The repaired task no longer claims table-driven layout,
named cases, diagnostic quality or testing-skill mastery. It assesses expected
results for inclusive comparisons through learner-written tests with a supplied
scaffold. No schema or taxonomy identifier changed.

## Third-batch cumulative audit

The 17 learner-visible tasks contain 9 implementation, 7 debugging and 1
test-writing activities, using 14 implementation submissions, 2 workspace
patches and 1 test submission. Independent review found the semantic progressions
defensible: append grows into storage isolation, map presence into deletion,
and interval decisions into independently written verification. None of the
existing twelve requires a content change from this review.

Task-form balance within each domain is still incomplete. Types/maps favor
implementation and control flow favors debugging. Future blocks should include
test-writing and executable prediction, with later practice in other forms for
maps, conversions and range. Existing Unicode drafts should be repaired when
their prerequisites are taught. These observations are curriculum debt, not a
claim of complete coverage or learner mastery.

The separate learner export issue for test submissions must be checked before
delivery: the exported test workspace currently lacks its checker-supplied API
declaration for the editor. Runtime verification uses disposable assembly; an
IDE support repair belongs in a tooling batch.

## Local acceptance

Fast tests, race, vet and tagged integration passed. One full 29-package
task-check collection was used to regenerate reports; phase1 validation and
the deterministic report check reused that same collection. Promotions used
task-promote after independent review. These are partial-bank checks, not a
foundation coverage release. Human progress remains one assisted completion;
no trial or authoring check is counted as learner mastery.
