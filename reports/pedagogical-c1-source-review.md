# S1 — source decisions for the first pedagogical content batch

Planning baseline: `b9790106edc3d4c7f57e26d5f45dc1960f022c6c`, supplied
2026-10-11 review, sections D/E. Integration base is M1 PR #70 at
`7b958d84bfa299e501fdcbc414d8631350be546d`. Exact planning-file hashes, source
commits, paths and SHA-256 values are recorded in
`catalog/intake/2026-10-17/pedagogical-c1-source-review.json`.

This is a bounded authoring-source decision. It creates no exercises, promotes
no packages and grants no coverage or learner evidence. Independent read-only source
review by `/root/clamp_e01_content_review` approved this decision on 2026-10-17.
It independently checked exact source/commit blobs, planning hashes, license
boundaries, bounded search and original reasons. Technical/content acceptance
remains required for C1.

## Existing packages

E01 reuses the admitted local Clamp contract and lineage released in PR #69.
E02 reuses the admitted Prepend contract and exact source chain in its existing
provenance. Their recorded source decisions remain historical; this review does
not claim a new audit of the upstream Exercism artifacts or generated cases.
No new external material is used for the two scaffold revisions. E02 remains
conditional on its exact workspace-compatibility preparation in PR #71, followed
by its own content/provenance/feedback acceptance.

## Bounded source inspection

Existing `source-sync` verified clean detached registry commits for
`ardanlabs-gotraining` (`4ae5a46bfc665909a7cb29e11c18775330c63401`) and
`golang-go` (`2c31595783c74564c1299a5401dc6e606f379cc7`). Upstream code, scripts
and tests were not executed.

The Ardan language-template inventory contains 14 files. A bounded keyword scan
for scope/shadow/GCD/short-circuit/expressionless-switch terms had no matches.
Three nearest template bodies were inspected; answer directories were excluded:

| Exact template path below `topics/go/language/` | Existing contract | Selection |
| --- | --- | --- |
| `variables/exercises/template1/template1.go` | Zero/literal variable declarations and conversion | Does not isolate returning an outer binding after a conditional block. |
| `functions/exercises/template1/template1.go` | Struct/pointer/error results | Does not supply scope repair or an observable lazy-operand contract. |
| `control_flow/exercises/template1/template1.go` | If-based input classification | Does not supply GCD reduction or expressionless-switch case priority. |

Selected files carry Apache-2.0 headers; the exact root LICENSE was hashed.
The selected ancestor paths contain no additional NOTICE/PATENTS files. No new
Ardan adaptation is admitted and no source text, code, tests or cases will enter
these originals. This negative result is limited to the inspected set; it makes
no assertion that no such exercise exists anywhere else.

The exact pinned Go `doc/go_spec.html` was inspected for declarations/scope,
short declarations, logical operands, condition-only for and expression switches.
Its LICENSE/PATENTS hashes are recorded. This is semantic fact checking only;
no spec examples or exercise material are copied into learner content.

## Original authoring reasons

| Proposal | Distinct requirement | Concrete reason for original authoring |
| --- | --- | --- |
| P01 | Repair an outer binding through a supplied nested conditional block | Current literal bindings assess types; the inspected variable/function templates do not provide this scope-repair contract. |
| P02 | Observe whether a supplied probe runs, and how often | Current Access/BooleanPolicies take already-computed booleans. The inspected flow template has no lazy operand or supplied countable probe. |
| P03 | Condition-only for with decreasing remainder state for bounded GCD | Current CountEven repairs a counted-loop boundary. The bounded permitted template set has no matching numeric reduction contract. Restricted binary-search/algorithm material is not reused. |
| P04 | Construct an expressionless switch with overlapping-case priority | Current Width permits equivalent dispatch; TokenValue maps exact string tokens. The inspected flow template only requires if statements. |

These reasons authorize preparing independently authored candidates after
independent review; they do not approve unreviewed task bodies. They do not relax
the historical LGWT freeze, reference-only restrictions or existing donor
deferrals. No registry usage mode or allowlist is broadened.

Each original will use ADR 0008 project identity, a real creation record, an
actual authored-content digest and empty external task sources. Do not invent
an original commit. Actual release binding remains a later release-manifest
operation. Candidate notes must link this decision and retain their independent
public/grader/reference/mutant and hidden-prerequisite review.

C1 remains a bounded batch: E02 plus P01–P04, with E01 already released.
New task IDs are assigned only during authoring. Addresses are append-only,
and mapping extensions preserve every old address/role tuple. Source decisions
and content batches remain separate PRs; no real learner home is accessed.
