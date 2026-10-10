# LRU route placement

Date: 2026-10-10. Baseline: master `3c00d704ad5bb573858444d206478018f064a3b5`.
Scope: preparation route batch under ADR 0022; no task authoring or export.

Add the already reviewed `bounded-lru-cache-001` as **07.20**, a `transfer`
encounter for `generics.reusable-data-structures`, at the end of Libraries and
before section 08 (Verification). The user now requested route inclusion after
the first wave merged. This settles W6's optional-placement investigation in
`catalog/seed-plans/foundation-wave-next.json`: use an ordinary route step under
the existing contract, without adding optional-step semantics or a scheduler.

LRU combines generic storage, typed absence, mutation and callback traversal.
It follows the simpler generic stack and gives a spaced transfer encounter.
Placing it directly after 07.06 would require changing the existing increasing
integer addresses (ADR 0015); 07.20 preserves every old address and relative
order. Concurrency is not a prerequisite: this cache has a sequential contract
and no complexity or performance requirement.

## Earlier executable preparation

| Contract prerequisite / preparation | Earlier lesson | Observable evidence in the existing task |
| --- | --- | --- |
| `types.zero-values` | 01.02 `zero-results-001` | Scalar zero results publicly checked; nil slice/map results additionally covered by full tests. Generic typed absence is exercised again by 07.06. |
| `methods.receiver-semantics` | 06.07 `receiver-reset-001` | Public tests check pointer-alias-visible mutation, repeatable clearing and preservation of the old map. |
| `functions.function-values` | 06.08 `predicate-filter-001` | Public tests check callback invocation, input order and empty-input behavior. |
| `generics.constraints-type-sets` | 07.05 `named-numeric-constraint-001` | Public tests check the exact type set and preservation of named numeric types. LRU's supplied API already declares `K comparable, V any`. |
| Prior introduction to the primary skill | 07.06 `generic-stack-001` | Public tests check zero-valued generic storage, typed misses, noncomparable values and independent instances. |

These are available checked exercises, not claims about the user's completions
or mastery. Theory reading provides no prerequisite evidence.

## Compatibility and learner behavior

- Route v2, task IDs, all 150 previous addresses and their workspace paths stay
  unchanged. The route now has 151 lessons; the bank still has 154 ready tasks.
- `learn next` preserves any active unfinished lesson, including lessons in
  sections 08/09. Otherwise existing sequence/review rules apply. After 07.19,
  an unfinished LRU is available as 07.20; after it, the sequence reaches 08.01.
  The existing delayed-review policy may recommend LRU earlier once a generics
  review is due and its prerequisites have completion evidence.
- A learner who finished the old route can receive the new unfinished LRU;
  earlier completion of LRU by stable ID remains completed and is skipped.
- New numbered homes create `workspaces/07-libraries/20-bounded-lru-cache`.
  Existing unnumbered LRU workspaces remain valid at their recorded paths.
  Loading the route neither moves files nor runs a progress migration. No reset
  or organize command is required.
- Completion evidence, attempts, hints, assistance, current task and time are
  retained. Existing theory, task content and provenance remain untouched.

## Acceptance

Regression fixtures pin the 150 baseline addresses, block IDs and relative order;
exercise route boundaries, active later lessons, delayed review and pre-existing
LRU completion; compare progress bytes and workspace trees across read-only
UI/CLI access for progress versions 1–3 and an unnumbered LRU workspace.
Fixture progress is test data, never learner evidence.

Run fast tests → race → vet → integration. Reuse the successful full 163-package
collection from the unchanged bank inputs for `validate -phase1 -checks` and
`task-report -check -checks`; CI independently collects the bank once and reuses
its collection. Independent review and green exact-head `validate` are required
before merge. This batch adds zero packages, adaptations or originals and does
not claim full foundation coverage.
