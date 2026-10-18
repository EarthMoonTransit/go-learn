# P18/P17/P19 preparation wave audit

Status: passed independent wave audit on canonical `8e08ea1f6b7997e8226f3d59a6ea69ab193e2313`, after all three local acceptances, exact-head reviews, green validate checks and actual merges. Reviewer: `/root/p19_final_review`; findings and blockers: none. No foundation release or export.

## Outcomes and assessment credit

| Preparation | Ready task and stable address | Mode / route role | Directly assessed distinction |
| --- | --- | --- | --- |
| P18 | `mapped-records-001`, 07.22 | implementation / practice | Declares two independent unrestricted type parameters in a callable API, preserving both through supplied input/output records. Find and Reduce already supplied their generic signatures; those tasks do not establish this declaration outcome. |
| P17 | `indexed-job-results-001`, 09.15 | implementation / introduction | Starts finite jobs concurrently, records each exact value and error identity by original index, and waits for all jobs even after an error. A void JoinAll contract does not establish result/error collection. |
| P19 | `gated-worker-arguments-001`, 09.16 | debugging / introduction | Repairs delayed reads of one existing range variable after the caller releases a start gate. Explicit argument binding, per-iteration declarations and index binding are accepted. The supplied defective capture is deterministic and race-free. |

P18 follows Reduce and precedes type sets. P17 and P19 follow the finite merge preparation and precede cancellation, in the order indexed collection → argument capture → cancel-worker. P17/P19 introduction roles refine the provisional PR85 transfer/practice planning because they assess distinct outcomes for the first time. The historical source decision remains intact; old completions gain no new credit.

The wave adds two implementation tasks and one debugging task. It adds no learner-authored-test task. `generics.type-parameters-inference` has three primary and three total occurrences against 3/5 minima; only P18 directly assesses declaring this API. The three task types are test-writing, debugging and implementation; the total remains below its minimum and no encounter has a transfer role. Numeric occurrences and type labels do not establish mastery of every observable outcome. One P17 introduction and one P19 introduction do not establish depth or mastery. No human learner completion evidence is used.

## Public feedback and valid alternatives

| Task | Required compiling defects rejected by public assertions | Main public evidence |
| --- | ---: | --- |
| P18 | 16 | `TestGenericDeclarationAndCallMatrix` checks declarations and explicit/inferred calls with named and non-comparable types using checker-owned go/types diagnostics. Runtime assertions check values, order, once-per-element calls, empty inputs and preservation. Public checks also cover the explicit no-erasure/reflection/goroutine requirements. |
| P17 | 13 | Indexed value/error identity, nonzero value with error, all jobs started before release, reverse completion with the error job released first, wait-all, empty batches and independent storage/calls. |
| P19 | 15 | Exact argument multisets with distinct/repeated/zero/negative values; return before start, no callback before start, all callbacks may enter while others are held, completion only after callbacks return, closure without payload, empty batches and independent calls. |

All 44 required defects compile and fail named public assertions with useful diagnostics. Compiler failures and timeouts are not detection evidence. Each task reference passes public and full verification. P17 additionally accepts an unbuffered-channel collector without a WaitGroup and a hybrid caller-plus-workers implementation. P19 accepts current per-iteration declarations and index binding with a shared already-closed empty completion handle. These alternatives verify that the tests enforce observable contracts rather than a single reference mechanism.

The reviewed public execution records contain 19 P18 rows, 20 P17 rows and 22 P19 rows, including the starters, references and alternatives. The canonical authored files match the disposable executed fixtures: 24 P18 files, 19 P17 files and 21 P19 files. Public synctest gates keep workers observable without sleep, capacity-limited callback observation channels or callback serialization. All task execution occurs through disposable checker/learner workspaces.

## Source decisions and provenance

P18 substantively adapts the admitted two-parameter per-element mapping operation from `birsi-go-training` at `3ae230e734644d637118f2a6818166922d5197c5`, `exercises/stage3/09_generics/generics.go`, SHA-256 `b043c993a83266c07510c4d7a5b3cba4cd1b27af0db8bda2394000c72105c4e5`, under MIT. The local record extension, prompt, tests and implementations were authored here. Admission: [PR83](https://github.com/EarthMoonTransit/go-learn/pull/83).

P17 substantively adapts finite fan-out/all-results from `ardanlabs-gotraining` at `4ae5a46bfc665909a7cb29e11c18775330c63401`, `topics/go/concurrency/channels/exercises/template2/template2.go`, SHA-256 `d24cc6c7dd8214386d2d75fc97eb2ad056167b6bf98e802b34af345a69d43f3a`, under Apache-2.0. The translation from a program to a callable job API and the indexed value/error extension are documented. Attribution preserves Kennedy/Walker credits. Admission and bounded comparison: [PR85](https://github.com/EarthMoonTransit/go-learn/pull/85).

P19 is an original Task Contract v2 task, with actual project identity, creation record, content digest and an empty external task-source list. PR85 records the specific comparison: the six selected Ardan templates do not provide the existing-variable delayed-capture contract, and reusing a fan-out skeleton alone would be nominal. Go specification/memory references supply semantic facts only. SHAD and Golings reference-only sources were not copied or closely paraphrased. The three provenance records retain deferred release-manifest binding; no source or future release commit is fabricated.

## Preservation and inventory

For P18, P17 and P19 respectively, all 166, 167 and 168 prior route tuples retain task id, block, stable number, role and skills. The corresponding existing pedagogical CSV rows preserve every field except execution order. The original pedagogical-order CSV is byte-identical across all three batches, and CRLF is preserved in the active mapping. All new prerequisites occur earlier or are explicitly prepared. All 3382 pre-existing task files are byte-identical to the pre-wave base; only the three new packages are added. Real learner home, stored progress and solution workspaces were not inspected or modified.

The final wave inventory is 169 learner lessons, 172 ready packages including legacy fixtures, and 179 lifecycle packages. It contains 87 substantive adaptations, 76 originals and six concept seeds. Direct primary/secondary assessment reaches 106 of 142 active skills, leaving 36 without an assessed occurrence. This count measures bank availability, not mastery, required modes/types, complete depth or foundation acceptance.

## Exact-head acceptance and remaining gates

P18 content: [PR86](https://github.com/EarthMoonTransit/go-learn/pull/86), reviewed head `8bbb53c8b1f2feeb9d138e4cab83e32c820f5346`, merge `0bc11985e9f67ab8677c3a7726f3239118e76442`, green exact-head [validate](https://github.com/EarthMoonTransit/go-learn/actions/runs/39077387431). P17 content: [PR87](https://github.com/EarthMoonTransit/go-learn/pull/87), reviewed head `0bc94aba2372ea978752731d324b1afa0e6d6065`, merge `3b899d97c39c045722fcd17f9f9145b167a705fe`, green exact-head [validate](https://github.com/EarthMoonTransit/go-learn/actions/runs/39080445818). Both have ordered fast, race, vet and integration success, phase1 validation and report-check using their single full collections.

P19 content: [PR89](https://github.com/EarthMoonTransit/go-learn/pull/89), reviewed head `39ecfca44ea651a7f7a97f5edf219b595c8f26c0`, merge `8e08ea1f6b7997e8226f3d59a6ea69ab193e2313`, green exact-head [validate](https://github.com/EarthMoonTransit/go-learn/actions/runs/39083441707). Ordered fast, race, vet, integration, phase1 validation and report-check pass. All six log hashes, the unchanged tested tree `7883a1f724fb8a5d6e32be5fad0a33282ac1e531` and collection hash were independently checked.

| Collection | Packages / failed | SHA-256 |
| --- | --- | --- |
| P18 | 177 / 0 | `8c3dedd2843fa4e98334ee3662c7e0108dbd5abc2c70a41ba5ea5ab064885260` |
| P17 | 178 / 0 | `207f8d097692556930465a71c3c93aba299143de13b6b0122bf5de0036769341` |
| P19 | 179 / 0 | `1ea635c0795450a6c9c3babfd5e69b095e8895e7c6466696e466d99ef462736b` |

Each batch's reports, validator and report-check share its collection. Review/lifecycle/notes participate in input freshness. Preparation phase1 acceptance does not satisfy the eventual full `-coverage` requirement. Foundation remains unfinished: M2 design and its separate migration/scoped numeric tooling, finite fuzz/benchmark/module profiles, reviewed source/content work, breadth/depth/components and independent final semantic/full-coverage acceptance remain open. Six Theory pages stay frozen.
