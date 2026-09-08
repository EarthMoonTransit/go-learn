# Go Learn: two milestones to complete Go fundamentals

Status: accepted implementation plan, 2026-09-08. This is the authoritative
active roadmap. GoLand is the editor; the local UI and short `learn` commands
share the learning cycle. Statements, hints and reviews remain English.

## Active finish line

Cover Go and its fundamental standard library through executable practice.
Backend, HTTP/PostgreSQL application work, production services, Docker and CI
as lesson topics are deferred. The repository's engineering CI remains required.
Standalone theory, entry diagnostics, a custom editor and cloud distribution
are outside this delivery.

| Milestone | Product result | Definition of Done |
| --- | --- | --- |
| 1. Working course: 150 ready lessons | Numbered sections/workspaces, systematic source adaptation, sequential practice, independent coding/testing and repetitions | 150 learner-visible lessons have prerequisites and route positions, executable checks and independent review. Repetitions and mastery evidence work separately from completion. Existing solutions and history survive migrations. At least 76 lessons substantively adapt permitted exercise material; concept-only seeds do not qualify. |
| 2. Complete foundation: approximately 450 bank tasks | All 142 active skills, including concurrency, context, fuzzing and benchmarks, plus integrated components | Required `-coverage` passes, necessary verification profiles exist, and no required route segment is unprepared. Components assess combinations of skills. |

The full bank estimate respects the existing 350–500 target and 409 primary /
693 total encounter minima. The core learner route is provisionally 250–350
lessons; alternatives support remediation. Product readiness does not establish
human mastery. Learner evidence must be reported separately.

## Current checkpoint

Initial baseline: master `f41a602645718e276e8683d84c4f58f461476aa0`, 22 lessons.
The checkpoint below is the C02 result, effective after its PR passes `validate`
and is merged. Private learner state was not read or modified.

| Measure | C02 checkpoint |
| --- | --- |
| Learner-visible ready lessons | 32; another 118 are required for milestone 1 |
| Substantive source adaptations | 3 accepted exercise adaptations; at least 73 more required for the 76-task criterion |
| Original / concept-seed | 24 original, 5 concept-seed; seeds are excluded from substantive reuse |
| Covered skills | 28 distinct assessed primary/secondary skills; this is course coverage, not learner mastery |
| Activities | 16 implementation, 11 debugging, 5 test-writing; C02 adds 5/2/3 respectively |
| Other packages | Three legacy ready packages, eight drafts, one rejected; excluded from learner-visible count |
| Human completions / confirmed mastery | Unknown: private progress not inspected; automated fixtures are not human evidence |
| Delivered tooling | Explicit UI/CLI numbers, recoverable numbered workspaces, progress v3 completion evidence and delayed reviews |
| Source allocation | 128 additions mapped; 84 adaptation candidates and 44 original slots; C02 consumes 3 and 7 |
| Next result | Ten further adapted exercises, then continue toward 60, 100 and 150 |
| Obstacles | Remaining authoring/review work; unsupported later verification profiles require separate tooling first |

60 and 100 are reporting checkpoints inside milestone 1, not mandatory stops.
A successful batch authorizes continuation to the next within the accepted plan.

## Course organization

Keep the current foundation order, sections 01–09: first steps; values and
collections; control and verification; strings/bytes; Unicode; structs, pointers,
methods, embedding, interfaces and errors; packages, generics, I/O, encoding and
time; deeper testing; concurrency and context. Backend/production placeholders
have been removed from the active route.

Store section and lesson numbers explicitly in a versioned route. Display
`01.01` alongside the English title in UI/CLI. Stable task IDs remain progress,
provenance and command keys; their numeric suffix is not a lesson number.
Existing assigned numbers must not be silently regenerated when content grows.

Organize private workspaces by numbered section and lesson, for example
`workspaces/01-first-steps/01-literal-bindings/`. Canonical task packages retain
their lifecycle layout. ADRs 0016–0017 provide versioned progress,
backup, `learn organize --dry-run`, `learn organize` and interrupted-operation
recovery. Read-only commands never migrate. Conflicts never overwrite solutions.

## Source-first content delivery

Inventory → skill-based selection → substantive adaptation → verification → route.
Start with the saved 177 Exercism candidates and selected permitted Learn Go
with Tests / Ardan Labs artifacts. Use first-party Go material for specific gaps.
Inventory large sets; the ten-task limit applies to release/review batches.
Before mass release, record all 128 planned additions with primary skill,
prerequisites, candidate, exact source/path/commit or original-development reason.

At least 76 of the 150 ready lessons must substantively reuse permitted exercise
material. A topic name, metadata-only mapping or concept-seed does not qualify.
Adapt relevant behavior/code/checks to technical, story-free lessons while
preserving all artifact provenance and notices. Do not copy reference-only
material, including related generated cases. Source solutions remain private.
Original work fills demonstrated gaps; never relabel it to meet the target.
Expand permitted-source selection and report any real deficit.

Implementation lessons start with signatures and minimal stubs, not a supplied
algorithm. Test-writing requires independently authored checks. Progress toward
multiple functions/files and complete components. Per ten new lessons, target
five implementation, three test-writing and two debugging tasks; also enforce
existing per-domain diversity. Existing content is repaired gradually, separately.

## Execution ledger and batch policy

| Batch | Result | Acceptance / status |
| --- | --- | --- |
| D02 | This active roadmap, aligned vision and execution ledger | Delivered in PR26; exact-head validate passed |
| S01 | Source inventory and allocation for 128 additions | Delivered in PR31; independent source/hash/dependency review and validate passed |
| T03 | Explicit route numbering and UI/CLI navigation | Delivered in PR27; labels, stable IDs and sequence verified; validate passed |
| T04 | Numbered workspaces and migration | Delivered in PR29; backup/recovery/conflict checks and validate passed |
| T05 | Repetition and independent mastery evidence | Delivered in PR30; legacy/reset regressions and validate passed |
| C02 | Ten Unicode lessons: three substantive adaptations and seven original boundary tasks | Independent review approved; weak learner-test suites are rejected; all required local/PR gates apply |
| C03–C14 minimum | Remaining 118 lessons from the source allocation and gaps | At most ten new/substantially revised tasks per batch; public/full/reference/mutant checks and independent review |
| K02 | Milestone-1 acceptance at 150 | All milestone conditions, adaptation count, route, components and progress evidence verified |

Tooling needed for unsupported task profiles is delivered before that content.
At least thirteen content batches are required; repairs may add batches. Each
batch has its own codex branch and PR, merged after exact-head `validate` passes.
Tooling and content are separate. No new approval is needed after each ten tasks.

After five real learner completions use available feedback to adjust upcoming
content. Its absence does not stop preparation or authorize fake calibration.
After three content batches audit semantic duplicates, activity balance and tests.
Maintain at least one prepared next learning block as the route expands.

## Mastery and acceptance

First repetitions become due after three other tasks. Confirm a skill only after
two distinct independent practice/transfer successes without hints or solution
reveal. Assisted completion remains progress and creates a need for independent
follow-up. Automatic fixture completion is not learner evidence.

Map each basic prompt requirement to public assertions, realistic wrong variants
and useful diagnostics. Follow QUALITY_GATES.md and DEBUGGING.md. Execute authored
content only in disposable verification environments. Tooling acceptance runs
fast tests, race, vet, integration, then one full CheckCollection; validator and
report-check reuse it. Partial banks use `-phase1`; foundation completion requires
`-coverage`. Migrations cover interrupted/repeated runs, conflicts, UI/CLI overlap,
solutions, completion, attempts, hints and time.

After each accepted batch update one checkpoint: ready / substantively adapted /
original and concept-seed / covered skills / human completions / confirmed mastery /
next result / blockers. Do not substitute coverage for mastery or planned for ready.

## Historical delivery and deferred continuation

The opening 12 lessons, control/testing expansion, local UI and five strings/bytes
lessons are delivered. LOCAL_UI_GOAL.md records that completed bounded iteration;
it does not cap this 150-task run. Reports retain their historical evidence.

After foundation, a separate accepted plan may restore HTTP/PostgreSQL, background
services, production failure scenarios, Docker and CI lessons. Redis, gRPC and
brokers remain a later extension. None is required to complete the active scope.
