# ADR 0031: Foundation scope and additive taxonomy v4

- Status: accepted design for a separate migration batch; no live activation
- Date: 2026-10-18
- Extends: ADR 0003; retains all existing skill meanings and encounter minima

## Context

The review-package M2 asks for an explicit foundation/specialist coverage
boundary and identifies bit operations and recursion as taxonomy gaps. This
M2 is a scope/taxonomy work item; it is distinct from the first-content milestone
also called M2 in `docs/plans/foundation-next.md`.

The v3 baseline is `curriculum/skills.yaml` at commit
`0bc11985e9f67ab8677c3a7726f3239118e76442`, SHA-256
`8fc463c93634564e7972117ea7f7c1f90d4e5cb82659f2f631698ef68ba4ab64`.
It contains 142 active assignable skills and 35 deprecated compatibility stubs.
Its full-bank requirements are 409 primary and 693 total occurrences, with a
350–500 ready-task target. These are bank availability/depth requirements;
they do not prescribe that every learner complete that many tasks in sequence.

The executable `validateCoverage` currently checks primary/total occurrence
counts and the ready-task target. Taxonomy validation also checks declarations
of required modes and distinct exercise types. It does not establish that the
bank actually supplies those modes, types, or semantic outcomes. Independent
semantic review remains necessary; a numeric coverage pass cannot stand in for
that review or for human mastery.

## Decision and activation boundary

Accept the following scope and additive contracts as the design for taxonomy
v4. This documentation batch does not change `curriculum/skills.yaml`, schemas,
loaders, tasks, routes, or learner state. New ids cannot be assigned to tasks
until the separately reviewed schema/loader/migration batch is merged.
That batch must implement the complete finite migration below. It must not
silently reinterpret the v3 contract or retrofit new credit onto old work.

### Foundation and specialist bank scopes

Retain language fundamentals, package/toolchain boundaries, standard-library
contracts, testing/fuzzing/benchmark methodology, and ordinary concurrent
lifecycle/ownership in foundation. Classify four existing outcomes as
specialist continuations:

| Existing id | Reason for specialist scope |
|---|---|
| `sync.rwmutex-tradeoffs` | Selecting a readers/writer lock requires measured contention after correct mutex ownership and experimental design. |
| `sync.atomics` | A low-level independent state transition is a specialized coordination alternative after happens-before reasoning. |
| `sync.cond` | Lock-coupled predicate notification is a specialized coordination alternative; its existing mutex prerequisite remains unchanged. |
| `sync.pool` | Disposable allocation reuse is a performance technique with no durable cache or correctness guarantee. |

The other 138 existing active ids remain foundation. `sync.once-initialization`
remains an ordinary one-time lifecycle outcome with explicit failure semantics.
This is a product boundary, not a claim that official Go documentation specifies
a course curriculum. Encounter tier alone does not determine scope.

All existing definitions, prerequisites, mastery/outcome requirements,
misconceptions, exercise recommendations, and minima remain unchanged.
Foundation prerequisites must stay inside foundation; the reviewed classification
is closed under the existing prerequisite graph. No optional component family
is converted into a skill id. This decision creates no reflection/unsafe/cgo
skills and does not exclude context values or timers without a separate decision.
Fuzzing, benchmarks, packages and toolchain work remain foundation even while
their finite verification profiles await implementation.

### Three additive contracts

The complete author-designed records are in
`docs/plans/taxonomy-v4-additions.json`. Their `active-proposed-only` and
`assignable: false` markers describe this inactive planning document, not a new
runtime taxonomy state. The migration will create ordinary active assignable v4
records only after validating the complete contract.

| New id | Observable distinction | Prerequisites |
|---|---|---|
| `expressions.bitwise-masks` | Any/all selected-bit inspection, set/clear/toggle, preservation of unrelated bits, including zero and overlapping masks. | arithmetic, numeric conversions, comparison |
| `expressions.integer-shifts` | Constant representability versus runtime fixed-width behavior; direction, signedness and counts at zero, width and beyond. | arithmetic, constants, numeric conversions |
| `functions.recursion` | Terminal cases, a progressing decomposition and complete combination of required recursive subresults. | if, function signatures/results |

All three belong to foundation and use the existing ordinary 3-primary/5-total
tier. Three primary opportunities reserve materially different implementation,
debugging, and learner-authored-test work; two further occurrences apply the
outcome in different contracts. This allocation is bank planning, not a claim
that those exercises already exist or a mandatory learner completion count.

`types.iota-enumeration` retains its existing bit-pattern declaration outcome.
Supplied mask manipulation needs neither iota nor runtime shifts. Integer shifts
include typed left-shift bit loss, signed versus unsigned right shifts, and
negative counts. A count at or beyond the operand width is legal and must not be
silently reduced modulo width. Negative constant counts are compiler facts;
negative signed runtime counts panic. Compile-negative demonstrations are
checker-owned facts, never compiling-mutant detection claims.

Recursion is not limited to trees. Terminal, minimal nonterminal and, where
applicable, unequal branching cases are general outcomes. Tree, pointer, struct
and callback prerequisites belong to the actual task contract when needed.
Writing tests for an existing recursive API is verification evidence; by itself
it does not establish learner-authored recursion.

P25 and every implementation/debugging task claiming this recursion outcome
remain blocked before authoring until an independent reviewer accepts a finite
assessment contract and disposable proof. Output equivalence plus self-call AST
presence is insufficient. The proof must accept the stated correct recursive
forms and reject iterative substitution, dead self-calls and unused dummy
recursion with useful public feedback. Accepted direct/helper/mutual forms and
any structural restrictions must be stated in the task. Runtime executed-path
evidence with bounded structural restrictions is only a candidate mechanism;
this ADR does not claim it has been implemented or proved.

### Exact numeric bank policy

| Inventory | Foundation skills | Specialist skills | Foundation primary/total | Specialist primary/total | Full-bank primary/total |
|---|---:|---:|---:|---:|---:|
| Existing v3 contracts | 138 | 4 | 401 / 677 | 8 / 16 | 409 / 693 |
| After three additions | 141 | 4 | 410 / 692 | 8 / 16 | 418 / 708 |

After migration there are 145 active skills and 35 unchanged deprecated stubs,
180 ids in total. The full-bank ready-task target remains 350–500.

Implement an explicit versioned `curriculum/coverage-scopes.yaml` sidecar,
validated by its own schema. It must classify every active v4 id exactly once
as foundation or specialist. There is no implicit default, wildcard or automatic
classification for an unknown future id. Omitted, duplicated, unknown or
deprecated entries are errors. An added id requires an explicit reviewed scope.

Existing `-coverage` continues to require full-bank numeric coverage, including
specialist ids and the unchanged bank-size target. A separate
`-foundation-coverage` flag will select the 141-id subset for numeric encounter
coverage. Supplying both flags is an error. The subset flag does not replace
ordinary validation of every task and the entire repository, and does not
require a newly invented smaller ready-task target. The 350–500 target remains
a full-bank property, always displayed with its actual full-bank status.

Both report forms must identify the selected scope and ids, actual/required
primary and total counts, per-skill deficits, remaining specialist/full-bank
requirements, and whether full-bank numeric coverage passes. Count only ready
primary and declared secondary occurrences under the existing rules;
prerequisites and supplied scaffolds do not become new credit. Separately label
the modes/distinct-types/outcomes as semantic review obligations. Do not present
them as automatically checked because occurrence counts passed. Any future
automatic evidence-coverage enforcement needs a separate finite design with
explicit activity/task-type and primary/secondary credit rules for both scopes.

The final foundation Goal still requires actual full-bank `-coverage` and
independent semantic acceptance. Scoped bank success does not establish complete
foundation delivery, a shorter learner route, or human mastery.

## Finite v3 to v4 migration

1. Freeze and hash the exact v3 taxonomy plus four versioned documents:
   `curriculum/mappings/exercism-go.yaml`,
   `curriculum/evidence/exercism-go.yaml`,
   `catalog/reviews/exercism-go.yaml`, and
   `catalog/seed-plans/exercism-foundation-pilot.yaml`.
   Record immutable snapshot paths, actual source commit, byte hashes, all 177
   ids and the sorted id-inventory digest. The unversioned concept inventory
   remains unchanged as well.
2. Add `schemas/taxonomy-v4.schema.json`, a dedicated v3-to-v4 migration schema,
   and a coverage-scope schema. Preserve `schemas/taxonomy-v3.schema.json`, the
   v2 snapshots and the v2-to-v3 migration. The new migration must account for
   all 177 old ids: 142 preserved active ids and 35 preserved deprecated stubs,
   with unchanged replacement targets and rejection semantics. Compare entire
   old teaching contracts, allowing only the explicit taxonomy-version advance.
   List exactly the three additions separately. Reject missing, duplicate or
   unknown entries and unreviewed contract changes.
3. Validate the historical v2-to-v3 migration against the immutable v3 target,
   rather than asking it to explain the live v4 additions. Keep old v3 report
   generation bound to the historical v3 inputs; create distinctly named v4
   migration and scope reports. Do not rewrite historical reports with v4 totals
   under a v3 heading.
4. Keep the four reviewed v3 documents byte-for-byte unchanged. A new explicit
   compatibility record binds each exact path/hash and referenced-id inventory
   to the preserved v3 contracts. Validate original provenance and decisions
   against frozen v3, then verify all referenced skills retain their exact
   reviewed active/assignable or deprecated semantics in live v4. This applies
   to curriculum mapping/evidence and triage review/seed loaders and validators.
   Do not relax version equality globally, silently repin old review decisions,
   or automatically map new skills. Unknown versions or changed reviewed bytes
   require a new reviewed record, not an alias. Existing `coverage_eligible:
   false` evidence remains ineligible.
5. Update only the finite loader, validation, report and task-check collection
   contract paths needed by the new files. Include relevant v4 schema, scope,
   migration and compatibility inputs in collection freshness checks. A changed
   contract must invalidate an old collection. Generate one new full collection
   after inputs stabilize; use that same file for reports, validator and
   report-check. No old successful collection is represented as current.

No learner-data migration is needed for these additive unchanged-id contracts.
Do not edit any task manifest, route row, display address, role, learner solution,
workspace marker, historical encounter or CompletionEvidence. Existing content
digests and exact revision-compatibility rules remain intact. No new skill gets
retroactive credit. The current all-bank route, current-task priority, readiness
and resume behavior remain active. Optional route selection would require its
own finite design and review before any activation.

## Implementation batches and acceptance

This ADR is a documentation decision. Complete the independent audit of the
P18/P17/P19 content wave before merging it. Deliver taxonomy/schema/compatibility
tooling and scoped numeric coverage tooling as separate reviewable batches;
neither batch contains new tasks or Theory pages. New content follows only
after its required profile and source decisions are actually accepted.

Migration checks must cover the complete old/new inventory; byte-exact old
contracts and reviewed documents; deprecated-stub rejection; prerequisite
closure; all arithmetic above; missing/duplicate/unknown scope ids; tampered
baselines/compatibility records; and stale collection rejection. Scope fixtures
must demonstrate foundation numeric pass with specialist deficit while full
coverage fails, then complete full-bank pass, plus missing foundation encounter
and unchanged bank-size failures. Semantic evidence is reported honestly.

Temporary learner-home fixtures must preserve existing task ids, addresses,
role/skill tuples, stored progress and solution bytes; current task, Next,
readiness, resume and historical completion interpretation must stay unchanged.
Read-only commands must not migrate data. Do not inspect the real learner home.

Each tooling batch runs fast tests, race tests, vet, then one integration run;
one full task-check collection drives generated reports, phase1 validation and
report-check. Final release still requires full `-coverage`, independent review
and green exact-head GitHub `validate`. The documentation batch can reuse an
unchanged-input collection and does not claim the future tooling checks passed.

## Alternatives

Keeping one undifferentiated scope leaves the requested product boundary
unexpressed. Zeroing specialist minima or removing ids would weaken full-bank
acceptance and erase teaching contracts. Changing v3 in place would make its
historical migration and evidence misleading. A label-only policy would not
provide a distinct, inspectable bank-coverage result. These alternatives are
rejected; the explicit additive v4 migration preserves history and keeps its
remaining work visible.
