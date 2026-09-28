# P5 / C17 — one generic LRU adaptation

The six-slot maximum produced one substantive adaptation:
`bounded-lru-cache-001`, primary `generics.reusable-data-structures`, activity
implementation, fixed-tests profile. No originals were added. SHAD commit
6718601587ec77cfaa632bbd29a8ca0716fe81ec supplies the integer cache's recency,
eviction, ordered stoppable traversal and reusable Clear behavior; exact four
artifact hashes and the root MIT notice are retained. The generic K/V API,
usable zero Cache and deterministic tests are local changes under ADR0024 (PR62).

## Assessment and review

The prior integer draft was rejected because its supplied interface methods did
not assess the declared primary. ADR0024 and its scaffold were independently
accepted before authoring. `/root/review_intake` then approved the substantive
content at 8557cb332c619e0f970f4f85ca7f5fee083fe27c on 2026-09-28, with no blocking
findings. Generic behavior is observed using nonempty slices and structs with
slice fields, named/composite/pointer keys, zero/nil presence, typed zero misses,
zero Cache, reference assignment semantics and independent cache instances.

Public tests also cover the complete source-derived operation contract. The full
suite adds integer and generic deterministic traces using an independent logical
access-counter oracle. A linear reference is valid: no performance threshold,
sleep, random discovery or implementation-shape assertion is required. The source
hashes, complete MIT notice and distinction from existing exercises were rechecked.

Targeted task-check passes; all 20 required mutants compile and are detected.
A separate public-only probe through the checker-owned disposable runner accepts
the reference and rejects all 20 with useful actual/expected diagnostics. The
canonical content digest remained
f57610fc17ba67e99e5efc2a0923d199899d6855c16d9b7b4fdb7ee01e9cf9aa.
No authored code was executed in the canonical task directory. An initial
vet-only probe was discarded; reported public evidence comes from normal tests.

The accepted package uses normal task-promote transitions draft → in-review →
ready. The 150 route IDs/numbers and all existing ready packages are unchanged.
The new package has no mandatory route entry until placement is reviewed
separately. No private learner history, fixture completion or timing estimate is
claimed as learner evidence. The old draft remains in Git at 75df9ec.

## Remaining gates and scope

Final repository acceptance passed on 2026-09-28. The one full collection
passed for 163 packages and generated the reports before the ordered suite, as
required by DEBUGGING for content changes. Fast (42.08s), race (265.55s), vet
(1.05s), and integration with the accepted 30-minute budget (324.19s) all exited
zero. Validator -phase1 and task-report -check then both passed using that same
collection; no task code was rerun to generate or check reports.

Collection SHA-256:
e2f2ec448d13fb4ab9dd3780037f8f872e3f64e3789dbb40006f1a662b77cce2.
Generated report delta 325d6b2 was independently approved: only LRU changed,
with 154 ready, 8 draft and 1 rejected package; all other 162 rows are unchanged.
Evidence is in the managed foundation-lru worktree under .tmp/results
(task-checks.yaml, acceptance.json and individual gate logs). Preparation follows ADR0022 (-phase1), not final foundation release
or public export (-coverage). Source PR54/PR60 and decision PR62 must land before
this dependent content PR. GitHub's billing failure cannot be bypassed by local
results. Theory expansion is frozen and no new theory page is a dependency.

W1 exact ParseDuration lineage is unresolved; W2 adds little distinct practice;
W3/W4 canonical lineage remains restricted; W5 supplies no distinct nil
representation contract. Externalsort, treeiter, dupcall and later concurrency
components remain outside this batch. One admitted exercise is sufficient; six
was a maximum, not a quota.
