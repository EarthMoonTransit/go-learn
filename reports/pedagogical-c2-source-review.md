# S2 — source decisions for interfaces, defer, methods and keys

Planning baseline: `b9790106edc3d4c7f57e26d5f45dc1960f022c6c`, supplied
2026-10-11 review, C2/P05–P10. Repository baseline is master after PR #72,
`18a6b7cedb19b06fd0dcf59ec1a2944f113188d3`. Exact source and planning hashes,
paths, commits and decision boundaries are recorded in
`catalog/intake/2026-10-17/pedagogical-c2-source-review.json`.

Independent read-only source review by `/root/clamp_compatibility_review`
approved this bounded decision on 2026-10-17, including all eight source hashes
and commit blobs, notice boundaries, bounded inventory, semantic anchors,
planning hashes and actual candidate digests. This decision creates no tasks,
changes no source modes, promotes no package and grants no coverage or learner
evidence. Content review and normal lifecycle gates remain separate.

## Existing candidate lineage

P05 reuses `formatter-contract-001`. Its recorded external artifact is only
Exercism's pinned `exercises/concept/airport-robot/.meta/config.json`, used as a
concept seed. The exact file and root MIT LICENSE were checked against clean
checkout bytes and commit blobs. Its metadata identifies an interface topic;
it supplies no Formatter/Pair/Render statement, test, implementation or solution.
No upstream exercise body, test, exemplar or generated case was read. Preserve
the existing concept-only provenance and creation record; this decision grants
no broader adaptation permission. The selected ancestor paths contain only the
root LICENSE. Content remains independently written and needs a fresh review
before promotion of the existing task ID.

P10 reuses the original `resource-finalize-001` draft. Its identity, creation
record, actual content digest and empty external sources follow ADR 0008.
One write followed by one close, short-write detection and preservation of both
failures are a distinct combined-error ownership contract. Reusing that draft
avoids a second task with the same purpose. This is not approval of its current
tests or learner prerequisites; C2 must review them and use `task-promote`.

## Bounded permitted-artifact inspection

All three source caches are clean at their registry pins. Eight source files
were hashed and matched to the exact commit blobs. No upstream code, scripts
or tests were executed.

The same 14 Ardan language templates were scanned for defer, comparable keys,
method values/expressions and capture terms, with zero keyword matches. Two
bodies were read; answers and Playground solutions were excluded:

| Exact path below `topics/go/language/` | Observed contract | Decision |
| --- | --- | --- |
| `methods/exercises/template1/template1.go` | Define an average method and iterate struct values | Already pinned by ValueAverage; does not save method values across receiver mutation. No new adaptation is claimed. |
| `maps/exercises/template1/template1.go` | Construct and iterate a string-keyed map | Does not isolate equality and copy behavior of composite keys. |

The selected paths have Apache-2.0 headers and only the root LICENSE in their
ancestor notice boundary. This is a bounded negative finding, not a claim about
all available exercises or repositories. No reference-only task collection,
Go Tour support repository, LGWT material or deferred algorithm donor is used.

The pinned Go specification supplies semantic facts for defer registration and
LIFO execution, result assignment before deferred calls, saved method receivers,
explicit receiver arguments of method expressions, and comparable struct/array
map keys. The exact LICENSE/PATENTS and specification hashes and six anchors are
recorded. Specification examples, code and wording are not task content.

## Original authoring decisions

| Proposal | Isolated behavior | Reason |
| --- | --- | --- |
| P06 | Defer registration, argument values at registration, LIFO | Current Capture combines panic/recover; Resource combines cleanup/errors. The inspected template set has no matching small recording contract. |
| P07 | Deferred adjustment of a named scalar result after the return expression | Existing Resource has a different ownership policy; no admitted preparation contract supplies the caller-value and pending-panic boundaries. |
| P08 | Saved value receiver versus saved pointer receiver after mutation | The exact Ardan average template and ValueAverage's explicit-argument method expressions do not cover saved receiver identity. |
| P09 | Equal struct/array keys, changed components and insertion by value | Inspected maps use scalar keys; a generic comparable constraint alone does not assess composite equality and copying. |

After independent source approval, these four candidates may be authored from
their own technical contracts, examples, tests and defect models. Each requires
ADR 0008 project identity, a real creation record, actual content digest and
empty external task sources. Do not invent an original commit; release binding
remains deferred. Candidate notes must link this decision and receive separate
public/grader/reference/mutant and hidden-prerequisite review.

ValueAverage currently explicitly requests method expressions. That observation
informs the conditional M1 review; it does not automatically change assessed
tags, prerequisites, route credit or historical progress.

C2 remains at six changes: two existing candidate promotions and four new tasks.
It is preparation under ADR 0022, without a release/coverage claim. New addresses
are append-only and enter the explicit mapping only after actual readiness.
