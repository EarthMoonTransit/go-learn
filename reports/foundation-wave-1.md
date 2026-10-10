# First-wave checkpoint — implementation prepared, merge blocked by GitHub Actions

Updated: 2026-09-28 (historical first-wave evidence: 2026-09-11). Active plan: docs/plans/foundation-next.md.
Original checkout: clean master at 5f925a757f4bba63396d3b4b755776da75e68cec.
Work was isolated in codex branches/worktrees; no private learner history was read.
PR53 (fd9d536) and P0 PR55 (5643bee) were independently approved and merged only
after exact-head validate succeeded. Master then became 369bf2f.

## Latest user decision — retain Theory, stop its expansion

On 2026-09-11 the user stopped further Theory development. Preserve the P1
implementation and six P2 pages already prepared and visible in the local preview.
Do not add theory pages or features, and do not make a new page a prerequisite
for another exercise. The 6/47 figure below is an inventory ratio, not a remaining
theory-writing target. The next focus is settling the existing PRs and reviewing
executable exercises. Source, assessment and CI gates remain in force.

The user's main checkout was subsequently switched, at the user's request, to
codex/foundation-theory-local-preview at daef63a so the existing launcher shows
the prepared Theory feature. The original clean-master statement above records
the first-wave starting state, not the current local branch.

## What exists

| Batch | Result | PR | Independent review | Status |
| --- | --- | --- | --- | --- |
| Planning | 47-topic map, source research and executable first wave | [53](https://github.com/EarthMoonTransit/go-learn/pull/53) | Approved fd9d536 | Merged |
| P0 / D03 | Aligned roadmap/vision/rules; ADR0020–0022 for Theory, source policy and preparation/release gates | [55](https://github.com/EarthMoonTransit/go-learn/pull/55) | Approved 5643bee | Merged |
| Intake | Pinned metadata inventories; stale aggregate digests repaired | [54](https://github.com/EarthMoonTransit/go-learn/pull/54) | Approved a25e92e | Awaiting CI |
| P1 / T14 | Optional topic loader/schema/validation, CLI/API/UI, safe links and read-only behavior | [56](https://github.com/EarthMoonTransit/go-learn/pull/56) | Approved 3c338b9 after null-list fix | Awaiting CI |
| P2 / C15 | Six local English pages and real task links | [58](https://github.com/EarthMoonTransit/go-learn/pull/58) | Approved daef63a | Draft depends on P1 |
| P3 / S02 | Artifact hashes, five initial bounded unit reviews, five Koans files, all 46 reconsiderations; narrow LRU selection | [60](https://github.com/EarthMoonTransit/go-learn/pull/60) | Approved 3e28648 after artifact/history and nil-wording corrections | Draft depends on intake |
| P4T / T15 | Nil public regression and exact content compatibility, ADR0023 | [57](https://github.com/EarthMoonTransit/go-learn/pull/57) | Approved be5d299 | Awaiting CI |
| P4C / C16 | Public nil len/cap assertion, seventh mutant, provenance and review mapping | [59](https://github.com/EarthMoonTransit/go-learn/pull/59) | Approved 5a656c2 with generated report | Draft depends on P4T |
| P5 / C17 | One generic LRU adaptation, 20 defects | [61](https://github.com/EarthMoonTransit/go-learn/pull/61) | Approved 8557cb3; normal lifecycle at 74e149e | Depends on P3 and ADR0024 in P6 |
| P6 / K03 | Checkpoint, Theory freeze, ADR0024 and next allocation | [62](https://github.com/EarthMoonTransit/go-learn/pull/62) | ADR0024 authoring decision approved ee008b4; final checkpoint separately reviewed | Awaiting CI; not foundation release acceptance |

Each PR's generated-report-only follow-up retains the reviewed behavior; its exact
head and collection evidence must be checked before merge. Draft dependency PRs
must be retargeted to master after their prerequisites land to trigger validate.
PR60 includes PR54 while intake remains unmerged; do not merge around that gate.

## Changed file groups

| Batch | Principal files |
| --- | --- |
| P0 | docs/{ROADMAP,PROJECT_VISION,CURRICULUM_RULES,SKILL_EVIDENCE,LEARNER_CLI,QUALITY_GATES}.md; docs/adr/0020–0022; reports/foundation-wave-baseline.md |
| P1 | internal/learner/{theory.go,theory_test.go,cli.go,ui.go,ui/app.js,ui/app.test.cjs,ui/index.html}; internal/validation/validate.go; schemas/topic.schema.json; ADR0020 and CLI documentation; reports/theory-tooling.md |
| P2 | curriculum/topics.yaml; theory/en/{workspace,first-tests,slice-views,slice-ownership,interface-contracts,cancellation}.md; reports/theory-pages.md |
| Intake/P3 | catalog/intake/2026-09-10; sources/registry.yaml; reports/task-quality.{md,yaml}; reports/foundation-source-review.md |
| P4T | internal/learner/{content_revisions.go,content_revisions_test.go,nil_feedback_integration_test.go}; ADR0023/index; reports/bounded-view-nil-tooling.md |
| P4C | tasks/ready/bounded-slice-view-001/{tests/public,solution/mutants/nil-spare-cap,task.yaml,provenance.yaml,review}; reports/task-quality.{md,yaml}; reports/bounded-view-nil-content.md |
| P5 | tasks/ready/bounded-lru-cache-001 (TASK, starter, public/grader, reference, 20 mutants, provenance/attribution, hints, lifecycle/review); reports/foundation-first-content.md; generated task-quality reports |
| P6 | reports/foundation-wave-1.md; catalog/seed-plans/foundation-wave-next.json; docs/adr/0024 and index; existing Theory freeze in vision/roadmap/plan |

## Verified learner loop

Browser smoke used only /tmp/go-learn-p2/.tmp/ui-smoke. It displayed all 150 old
lessons, opened local Theory from lesson 01.01, followed its Practice link,
created a private numbered workspace, sent it through the existing GoLand
launcher, ran public Check on the starter and showed useful type/value diagnostics,
then returned to Theory. Comparing progress.json immediately before/after return
showed identical bytes. No hints, completion or skill evidence came from reading.
At first-wave verification, all ten worktrees preserved the same 150 (section, task ID, number) tuples and
byte-identical taxonomy. Route SHA-256:
a1ae9dfd067f8f8f4847acd50741a83bc86b6a8aeebd2f31205552593a53b7e9.
The one browser attempt and all automated completions are fixtures, not the user's
achievements. The temporary UI session was separate from the normal learner home and stopped
after the smoke test.

Six page examples passed in disposable Go 1.27 modules with module networking off.
Official reading URLs were checked; pages themselves load locally without remote
assets. B2 uses only a two-index example; L1 does not implement StartCancelable.
P4 independently passed an omission experiment: removing only TestPublicNilView
retaining the content review marker causes unexpected public success and fails
the persistent regression. Non-nil empty zero-capacity results remain valid.

## Counts and limits

- Routed tasks remain 150: 84 substantive adaptations, 61 originals, 5 concept seeds.
- Historical 84 adaptations: 55 tasks from 24 Exercism source exercises, 13 Ardan and 16 LGWT.
- Master has 153 lifecycle ready packages. P5 now prepares 154, including the new LRU outside the route.
- New ready tasks in the content branch: 1 SHAD adaptation from 1 source exercise. New originals: 0; none merged yet.
- P2 supplies 6 of 47 planned topic explanations; theory does not add task coverage.
- Assessed coverage remains 99 of 142 skills; phase1 is preparation, not full coverage.
- Intake counts remain 62 SHAD schedule units, 43 Golings entries, 12 Birsi packages,
  20 Koans thematic files and 6 Loong challenges; 143 mixed records are not 143 tasks.
- Source review joins 24 queued units plus 5 Koans files; unreviewed units remain
  needs-artifact-review. LRU is the single admitted source exercise, now approved and promoted in its
  content branch; it is not merged, released or routed.
- Human completion, mastery, timing calibration and retention evidence: unknown.

## Historical local acceptance — September 11

Ordered fast tests, race, vet and taskintegration passed for the planning baseline,
intake, P1, P2, P3, P4T, P4C and P5. Each completed with exit status zero.
P1's final null-list validation change additionally passed focused normal/race Go
tests and the Node renderer checks; bank inputs remained identical. Validator
uses -phase1 under ADR0022 and task-report -check uses the same saved collection.
No final -coverage or public release is claimed.

| Inputs | Collection SHA-256 | Reuse |
| --- | --- | --- |
| Baseline / P1 | 8bb9d8e0187a1adc77f1ad4b30470e81d70a8e67f7a876aa6a619e0a53924213 | P2 and P4T validation/report reuse baseline bank inputs |
| PR54 intake registry | 4f4fe2f95e5de155b9257fe74cc00c338c12d609bade7371472ddbaa83ad853b | Intake validator/report |
| P3 narrow source permission | 3174a727df33867d8292263659299441a8b8daeac0f67d1024b2e6ac7a85fe15 | P3 validator/report |
| P4C nil amendment | 936372b8a0719ceb7447682e5255f8bb0b441eecb88438f8c969a3772398cf47 | P4C validator/report |
| P5 LRU draft plus P3 registry | b10f63b663597e0ace3236961323fe7aab97733bb53448251ef0b2a77b77c6c2 | P5 validator/report passed; 163 packages, draft review remains fail |

The September 11 worktrees and local collections were under /tmp/go-learn-*;
the OS has since removed those temporary directories. The hashes above record
historical verified runs; the underlying local files are no longer available.
Do not claim they are currently reusable. Git commits and remote PRs are intact.
Work resumed in managed worktrees foundation-settle and foundation-lru under
/Users/egor/.codex/worktrees. New generic-content evidence is saved in
foundation-lru/go-learn/.tmp/results. No learner history is used.

## Fresh local acceptance — September 28

The baseline full collection was rebuilt once because its old local file was lost.
It reproduces SHA-256
8bb9d8e0187a1adc77f1ad4b30470e81d70a8e67f7a876aa6a619e0a53924213
exactly; validator -phase1 and report-check both pass using that same file
(162 packages). P6 is a documentation batch; bank inputs remain unchanged.

P5's final ready package has content digest
f57610fc17ba67e99e5efc2a0923d199899d6855c16d9b7b4fdb7ee01e9cf9aa.
Independent content review approved 8557cb3 and lifecycle delta 74e149e.
One full collection passes for 163 packages (154 ready, 8 drafts, 1 rejected):
e2f2ec448d13fb4ab9dd3780037f8f872e3f64e3789dbb40006f1a662b77cce2.
Generated reports were committed at 325d6b2. Ordered fast/race/vet/integration all passed (42.08s / 265.55s / 1.05s /
324.19s), followed by successful validator -phase1 and report-check using the
same collection. Generated report delta 325d6b2 was independently approved;
all other 162 package rows are unchanged. No authored code was rerun for reports.

The current main checkout remains clean on the existing Theory preview. Its
route hash above is unchanged; this work did not read or rewrite learner history.

## Exact external blocker and continuation

GitHub Actions run 34632696791 for P1 head 3c338b9 was rerun on September 28
(attempt 2) and again failed without starting steps. New P6 run 36476056635 at
35932af failed for the same reason, independently confirming the blocker.
The check annotation says recent account payments failed or the spending limit
needs increasing. The user was asked to restore Billing & plans. This is not a
Go test failure, and local success cannot replace the required GitHub validate.
No payment/settings change, direct master push or merge bypass was attempted.

After service restoration: rerun failed validate at the current immutable heads;
verify exact-head success and independent reviews, merge prerequisite PRs,
retarget their dependent drafts and run the same gate there. Preserve saved
CheckCollections for unchanged inputs. After registry/content merges, reconcile
combined generated reports from a single collection for the resulting inputs;
never declare stale independent reports current on a different combined tree.

The integer LRU draft's assessment rejection was resolved at the decision level
by ADR0024, independently approved at ee008b4 on September 28. The source
behavior remains SHAD LRU; the local generic API assesses the existing
`generics.reusable-data-structures` skill, with usable zero Cache, typed misses,
noncomparable values and named/composite/pointer keys. No taxonomy/profile or
learner migration is introduced. The previous rejected draft remains in Git.
Twenty compiling defects were detected by full tests and by a separate public-only
disposable probe at 8557cb3. Independent content review approved 8557cb3; normal task-promote moved the
package through in-review to ready at 74e149e. Final repository acceptance passed, as recorded in the fresh evidence section. No new theory dependency
is attached to this exercise.

The next allocation is bounded to settlement/placement of this one adaptation.
No new task quota, scheduler, progress reset or Theory expansion is authorized.
Tabletest awaits exact Go lineage; structs2 offers too little distinct practice;
binary-search/pangram canonical data stay restricted. Nil-representation work is
not inferred from the separate nil-input len/cap repair. Externalsort, treeiter,
dupcall and subsequent concurrency components stay outside this batch.

This checkpoint does not declare the first wave or foundation complete. Public
release/export and final foundation acceptance still require -coverage.
