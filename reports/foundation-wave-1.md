# First-wave checkpoint — implementation prepared, merge blocked by GitHub Actions

Date: 2026-09-11. Active plan: docs/plans/foundation-next.md.
Original checkout: clean master at 5f925a757f4bba63396d3b4b755776da75e68cec.
Work was isolated in codex branches/worktrees; no private learner history was read.
PR53 (fd9d536) and P0 PR55 (5643bee) were independently approved and merged only
after exact-head validate succeeded. Master then became 369bf2f.

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
| P5 / C17 | One executable LRU draft, 14 defects; no promotion | [61](https://github.com/EarthMoonTransit/go-learn/pull/61) | 8ea26eb approved for preserving draft; assessment request-changes retained | Draft depends on P3/P2; ready blocked by assessment |
| P6 / K03 | This checkpoint and prepared next allocation | This PR | Pending report review | Partial wave, not foundation acceptance |

Each PR's generated-report-only follow-up retains the reviewed behavior; its exact
head and collection evidence must be checked before merge. Draft dependency PRs
must be retargeted to master after their prerequisites land to trigger validate.
PR60 includes PR54 while intake remains unmerged; do not merge around that gate.

## Verified learner loop

Browser smoke used only /tmp/go-learn-p2/.tmp/ui-smoke. It displayed all 150 old
lessons, opened local Theory from lesson 01.01, followed its Practice link,
created a private numbered workspace, sent it through the existing GoLand
launcher, ran public Check on the starter and showed useful type/value diagnostics,
then returned to Theory. Comparing progress.json immediately before/after return
showed identical bytes. No hints, completion or skill evidence came from reading.
The one browser attempt and all automated completions are fixtures, not the user's
achievements. The temporary UI session was separate from the normal learner home and stopped after the smoke test.

Six page examples passed in disposable Go 1.27 modules with module networking off.
Official reading URLs were checked; pages themselves load locally without remote
assets. B2 uses only a two-index example; L1 does not implement StartCancelable.
P4 independently passed an omission experiment: removing only TestPublicNilView
retaining the content review marker causes unexpected public success and fails
the persistent regression. Non-nil empty zero-capacity results remain valid.

## Counts and limits

- Routed tasks remain 150: 84 substantive adaptations, 61 originals, 5 concept seeds.
- Historical 84 adaptations: 55 tasks from 24 Exercism source exercises, 13 Ardan and 16 LGWT.
- Total lifecycle ready remains 153, including 3 legacy packages outside the route.
- New ready tasks: 0. New originals: 0. Authored substantive adaptations: 1 draft, 0 ready.
- P2 supplies 6 of 47 planned topic explanations; theory does not add task coverage.
- Assessed coverage remains 99 of 142 skills; phase1 is preparation, not full coverage.
- Intake counts remain 62 SHAD schedule units, 43 Golings entries, 12 Birsi packages,
  20 Koans thematic files and 6 Loong challenges; 143 mixed records are not 143 tasks.
- Source review joins 24 queued units plus 5 Koans files; unreviewed units remain
  needs-artifact-review. LRU is the single admitted source behavior and now an executable draft, not ready.
- Human completion, mastery, timing calibration and retention evidence: unknown.

## Local acceptance evidence

Ordered fast tests, race, vet and taskintegration passed for the planning baseline,
intake, P1, P2, P3 and P4T. P4C and P5 are finishing their own ordered acceptance.
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

Local evidence is under /tmp/go-learn-{wave,intake,p1,p2,p3,p4t,p4c,p5}/.tmp/results
in isolated worktrees. Collections and logs are local check artifacts, not learner
history. P6 changes only this report/allocation; it validates against the baseline
collection rather than rerunning unchanged task code.

## Exact external blocker and continuation

GitHub Actions run 34632696791 for P1 head 3c338b9 failed without starting steps.
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

P5 preserved bounded-lru-cache-001 as a draft in PR61 with public/full tests,
reference and 14 realistic defects. Independent review verified the source but
found that the provisional primary skill is not assessed: the compiling starter
already supplies the complete interface. Pointer receivers are also supplied,
so receiver semantics is not a justified substitute. A public zero-key and
negative-value gap was fixed; all 14 defects then failed a public-only isolated
probe with useful diagnostics, and the reference passed. No promotion or route
entry was made. This is a specific assessment blocker, separate from CI billing.

catalog/seed-plans/foundation-wave-next.json prepares a maximum one-task follow-up
on that existing draft, with no new task quota. It requires an independently
reviewed observable assessment mapping; any activity/profile or taxonomy change
must be explicit and use its applicable ADR/migration gate. It does not authorize
cosmetic relabeling, a scheduler, progress reset or a broad content expansion.
Tabletest awaits exact Go lineage; structs2 offers too little distinct practice;
binary-search/pangram canonical data stay restricted. Nil-representation work is
not inferred from the separate nil-input len/cap repair. Externalsort, treeiter,
dupcall and subsequent concurrency components stay outside this batch.

This checkpoint does not declare the first wave or foundation complete. Public
release/export and final foundation acceptance still require -coverage.
