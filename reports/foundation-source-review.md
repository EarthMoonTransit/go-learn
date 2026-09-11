# P3 / S02 — bounded source review and next allocation

Source-sync fetched the pinned SHAD, Golings, Birsi, Koans and Exercism commits
with detached checkout manifests. No upstream code/tests/scripts were executed.
The initial intake remains a historical manifest audit; its reference-only status
is not retroactively rewritten. The current decision is recorded in
catalog/intake/2026-09-10/artifact-review.json, including exact file SHA-256.

The 24 queued source units and five Koans files have joined artifact inventories.
The first five units received bounded behavioral/lineage review; other queue
entries remain needs-artifact-review. The five Koans files remain remediation
references, not five new exercises. All 46 historical Exercism rejections have
an append-only reconsideration record; native historical reviews stay unchanged.

## First-wave decisions

| Slot | Decision | Reason |
| --- | --- | --- |
| W1 tabletest | Deferred | ParseDuration's precise Go origin and notices are unresolved; root MIT alone cannot clear that chain. |
| W2 structs2 | Rejected for this wave | A compile-only embedding hole adds little beyond existing executable embedding practice. |
| W3 binary-search | Deferred | Generated canonical cases and external statement lineage are not admitted; algorithmic policy is no longer the rejection. |
| W4 pangram | Deferred | Same canonical-data boundary; completeness differs from uniqueness, but that does not confer copying permission. |
| W5 nil/empty | No new task | Birsi collections does not require distinct nil representation; existing bounded-slice-view already promises it and P4 repairs public feedback. |
| W6 lrucache | Accepted for adaptation, subject to this PR's review/gates | A bounded container with cross-operation recency/eviction, ordered traversal and clear is distinct from existing map updates and stacks. |

Only SHAD's four lrucache files enter a narrow adapt allowlist. MIT notice must
travel with the adaptation. Dependencies, random workload, GC settings, benchmark,
all other SHAD files and all other new sources remain outside copying permission.
The GitLab module identity is preserved as lineage context. No commercial course,
external article or solution repository is inferred to share that permission.

P5 allocation is one substantive adaptation, not a six-task quota. It needs an
intermediate implementation/transfer contract, public empty/capacity/update/Get
refresh/eviction/Range/clear assertions, independently authored deterministic
checks, a reference and realistic defects. Prerequisites: maps, methods, interfaces
and callbacks. Keep the package outside required sequence if existing route
placement cannot prove preparation; never renumber old lessons.

Externasort, treeiter, dupcall and subsequent concurrency components stay in the
next wave. Recover tabletest provenance or select another permitted contract
before adding that slot. No original is authored to evade a restricted source.

No ready packages, route entries or learner state change in P3. Source units,
adaptations and learner achievements remain distinct; human evidence is unknown.
