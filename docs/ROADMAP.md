# Go Learn: personal route from Go fundamentals to production backend

Status: accepted personal learning roadmap, 2026-09-05.

## Finish line

The learner independently builds and tests a backend service with HTTP,
PostgreSQL, background work, cancellation, graceful shutdown, diagnostics,
Docker and CI. The first completed route excludes Redis, gRPC and brokers;
these belong to a later distributed-systems extension.

Learning starts sequentially from the fundamentals, without an entry diagnostic.
GoLand is the editor; `learn` is the local command interface. Task statements,
hints and reviews are in English. Milestones depend on completed learning blocks,
not calendar deadlines. A custom editor, public distribution and a thousand-task
catalog are not prerequisites for completing the personal route.

## Milestones

| Milestone | Learner-visible ready tasks | Acceptance |
| --- | ---: | --- |
| Sequential beginning | 12 | A beginner can follow the opening blocks without unexplained prerequisites, file/command discovery or hidden requirements. |
| Working course | 60 | Types, expressions, flow, functions, collections and introductory tests; real learner evidence for implementation, debugging, test submissions and workspace patches. |
| Deeper language practice | 150 | Structs, pointers, methods, interfaces, errors, Unicode, generics, packages and I/O; one finished tested component; spaced practice. |
| Complete foundation | About 450 in the bank | All 142 active skills satisfy required coverage; independent learner practice confirms mastery. The core route is provisionally 250–350 tasks, with alternatives used for remediation. |
| Backend | 12-task slice, then 60 | HTTP/JSON, errors, PostgreSQL/transactions, clients and repeatable integration tests against a real local database. |
| Production project | Capability-based | Independently completed service passes normal-operation and failure scenarios. |

The foundation estimate respects the existing 350–500 target and 409 primary /
693 total encounter minima. Estimates are not permission to weaken quality or
generate cosmetic variants. Backend coverage must be separated from foundation
before backend tasks are introduced.

## Learning sequence

1. Toolchain and project structure; types; expressions; control flow; functions.
2. Arrays; slices; maps; strings and Unicode.
3. Structs; pointers; methods; embedding; interfaces; errors.
4. Packages and generics; I/O; files; encoding; time.
5. Deeper testing; fuzzing; benchmarks.
6. Goroutines; channels; sync; context; lifetime management.
7. HTTP; PostgreSQL; background work; operating the service.

Testing starts early and becomes more demanding. Atomic skill prerequisites
determine ordering within these broad blocks. Each skill follows introduction,
practice, transfer and spaced repetition. First repetitions become due after
three other tasks. Mastery requires two distinct independent practice/transfer
successes without hints or solution reveal; assisted completion remains progress
and leads to another independent check. Major blocks end with a component or an
extension of the final service.

## Delivery policy

- Content batches contain 5–10 tasks, never more than 10 substantially changed
  tasks. Tooling, taxonomy and adapters are separate batches.
- Maintain at least one next learning block ready. After every five learner
  completions, use feedback to adjust the next block. After three content
  batches, review semantic duplicates, task-form balance and verification quality.
- Repair existing advanced drafts when their prerequisites have been taught.
- Introduce sources only for a demonstrated content gap, preserving exact
  provenance and independent review. No reference-only material is copied.
- Keep canonical tasks separate from learner workspaces. Authored code and
  task tests execute only through disposable verification environments.
- Every batch uses a feature branch and PR, merged only after `validate` passes.
- Tooling acceptance runs fast tests, race, vet and integration; one full
  CheckCollection is reused by validation and report checks. Partial banks use
  `-phase1`; completion of foundation requires `-coverage`.

## Engine changes, when needed

Implemented for the opening course: a versioned route description and
`learn roadmap` show the current block, completed tasks and planned material.
`learn next` follows the route and explains missing preparation instead of
silently jumping ahead. Existing task IDs and learner completion records remain
intact. ADR 0010 records this change; later engine work follows below.

Separate task completion from skill mastery and support repetitions. Any change
to progress storage needs a version, backup and tested migration preserving
attempts, hints, completion and session time. Coverage is not learner mastery.

Add narrow toolchain/project checks in disposable workspaces for real module
exercises. Before teaching fuzzing/benchmarks, implement their reserved profiles
with bounded, stable evidence rather than absolute wall-clock thresholds.
Before backend, introduce a separate track and managed local HTTP/database
verification. Each contract extension needs its own ADR and tooling batch;
foundation checks remain reproducible and offline.

## Final acceptance

The service must demonstrate ordinary operation, invalid requests, transaction
failure and rollback, database unavailability, cancellation, background-worker
termination and restart with durable data. The learner explains design decisions
and repairs an introduced defect. Automatic checks alone are not human mastery.

## Current checkpoint

| Measure | Verified state on 2026-09-05 |
| --- | --- |
| Learner-visible ready tasks | 17 (the opening twelve plus four new lessons and one repaired test-writing draft) |
| Human-completed tasks | 1 (`map-take-001`: four attempts, one hint, no solution reveal) |
| Confirmed mastered skills | 0; one assisted completion is not mastery |
| Next product result | Human trial of the prepared blocks; 43 more ready tasks reach the 60-task milestone |
| Current gap | Human feedback on the opening block; mastery and spaced repetition remain scheduled for the working-course milestone |

The stored session duration includes any time without an explicit pause and is
not claimed as independently measured active learning time. The first feedback
and review-context export are stored privately in learner home.

## Opening work delivered

1. Recorded this roadmap and aligned the product vision with the personal finish line.
2. Prepared two reviewed English content batches of five tasks, with explicit
   prerequisites and positions in the learning route.
3. Added `learn roadmap`, sequential selection and an English GoLand preparation
   guide. The opening set contains 12 available tasks. Automated acceptance is
   complete; the beginner usability criterion still needs real learner feedback.

The next practical block adds numeric conversion, interval decisions, switch
dispatch, range updates and learner-written interval tests. The repaired
`clamp-tests-001` retains its Exercism concept-seed provenance and prior rejection
history. Its revised contract assesses comparison verification; it does not
claim table structure, diagnostic quality or testing mastery. The original
twelve lessons retain their order and learner completion records.

## Next batches

Use the first five real opening-course solutions to adjust the next block.
Continue toward 60 tasks in batches of 5–10: fill out the initial types,
expressions, control flow, functions and collections, and introduce test-writing
tasks. Add independent-skill evidence and repetition through a separate ADR and
tooling batch, including a versioned, backed-up progress migration. Do not count
an automated trial or the existing assisted completion as independent mastery.

After every accepted batch update this checkpoint: ready tasks, human-completed
tasks, confirmed skills, next result and blockers. Planned, draft and legacy-only
packages never count as available learner content. Human calibration gates later
milestones; automated fixture completions cannot satisfy them.

## Deferred final plan item: theory companion

User decision, 2026-09-05: keep building the practical course now; schedule this
item at the end of the plan. It does not block the next content batches.

Add short English introductions before new topics, with separate worked
examples and common mistakes. Link specific relevant sections of Go Tour or
other approved sources, then connect reading to the corresponding exercises.
After completion, explain the reasoning, edge cases and alternative approaches;
link later repetition to the same topic. Worked examples must not disclose the
pending exercise's solution, and source reuse follows the existing registry.

Acceptance: the learner can see what to read before a block and where to find
its post-completion explanation from the normal `learn` workflow. At present,
the course has task statements, hints and a toolchain/workspace preparation
guide; a systematic theory companion is not implemented.
