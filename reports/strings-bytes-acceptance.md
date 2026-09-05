# Strings and bytes — five-task content batch

Date: 2026-09-05. C01 follows the accepted local UI and IDE compatibility work.
The batch adds exactly five learner-visible Task Contract v2 lessons to the
previous 17. Legacy packages and drafts are not included in that increase.

| Task | Primary skill | Activity | Required defects |
| --- | --- | --- | ---: |
| replace-byte-001 | strings.immutable-bytes | implementation | 6 |
| exact-prefix-001 | strings.standard-library | debugging | 6 |
| typed-line-001 | strings.formatting | debugging | 6 |
| join-values-001 | strings.building | implementation | 6 |
| byte-window-tests-001 | strings.immutable-bytes | test-writing | 6 |

## Content and independent review

Each English statement defines a narrow observable API contract, constraints,
public checks and progressive hints. Four implementation submissions use
fixed-tests; the supplied test-writing scaffold uses the real mutation-tests
profile with a correct verification implementation and six defect variants.
String indexing and slicing operate on bytes, including opaque byte values;
advanced Unicode knowledge is not a prerequisite. Construction accepts either
concatenation or a builder without timing/allocation requirements. Test-writing
assesses boundary behavior, not table structure or diagnostic style.

Independent read-only agent reviewer `/root/ui_goal_reviewer` approved all five
packages and the actual route after two repairs: an unverified no-console-output
prohibition became guidance, and an artificial input-corruption variant became
plausible destructive decimal conversion while constructing the correct output.
The reviewer checked prompt/test alignment, prerequisites, distinct mechanisms,
public APIs, license/provenance, clean-room evidence and solution boundaries.
Each package records that approval and six required mutants. This is independent
agent review, not human calibration; difficulty/time estimates remain provisional.

All lifecycle transitions used task-promote: initialized draft, in-review, ready.
No existing task was moved or overwritten. The first 17 route entries retain
exactly their original order, roles and skills. The five-task block follows them
without an empty intervening block. Three other exercises separate initial byte
practice from independent byte-window verification. Unicode remains a separate
planned block afterward; neither Unicode completion nor mastery is implied.

## Provenance

exact-prefix-001 uses welcome-to-tech-palace as a concept seed; typed-line-001
uses party-robot. Both are pinned to Exercism Go commit
`ae1477abfeeb9fbeaf971caee220e91f9e4df790` and exact
`exercises/concept/<seed>/.meta/config.json` paths. Author and independent reviewer
read only those metadata files and LICENSE, verified their stored SHA256 values,
authors/contributors and MIT notice. The author independently wrote the APIs,
text, implementations, tests and variants; the reviewer inspected that work
without authoring it. No upstream statements, solutions, generated
tests or problem-specifications canonical cases were read or copied.

The other three tasks are original under ADR 0008: project identity, creation
record, authored-content digest and an empty external sources list. Digests
remain subject to factual VCS binding at release time; no commit is fabricated.
No source adapter, taxonomy, schema or verification profile changed.

## Acceptance

Targeted disposable task-checks passed all five references and detected all
30 required defect variants. The repaired formatting/construction tasks were
rechecked before promotion; promotion reran the applicable gates.

A temporary-home service acceptance checks 22 ready v2 lessons, selects each of
the five new tasks in order after synthetic completion of the original 17, then
reports planned Unicode. Next/Roadmap leave the persisted fixture bytes unchanged,
including existing completion, hints, attempts and time. The original 17 route
entries are also compared directly with the accepted pre-C01 master. This fixture
is not a user completion record. No real learner home was read, copied or changed.

The final ordered repository checks passed: `go test ./...`,
`go test -race ./...`, `go vet ./...`, and
`go test -tags=taskintegration ./...`. The initial fast run identified a stale
29-package report; the one full 34-package CheckCollection was then produced and
used to regenerate the report before the successful ordered run. No test or gate
was weakened and no second full bank execution was needed.

The collection at `.tmp/goal-ui/c01/task-checks.yaml` has SHA256
`723b94434eac253c03124b48d44f9c06c582dea6957d4268ae0cc9b0edcae308`:
one session, 34 task checks, 314 disposable workspaces and 299 variant runs.
`validate -root . -phase1 -checks <collection>` and
`task-report -root . -check -checks <collection>` reused it and passed. All five
new ready packages have passing technical/review status and six of six required
defects detected. These are partial-bank gates, not foundation coverage release.
Exact-head PR validate and post-merge validate remain required before this batch
is considered accepted in master.

## Continuation

The prepared bank has 22 ready v2 lessons plus three legacy ready packages;
eight pre-existing drafts and one rejected package remain outside the learner
route. Local UI launch remains `learn ui` or `go run ./cmd/learn ui` from the
checkout; GoLand remains the editor. UI/browser/IDE acceptance is recorded in
local-ui-acceptance.md and ide-metadata-acceptance.md, not repeated or counted as
human use here. The historical human checkpoint remains one assisted completion
and zero confirmed mastered skills. After five real solutions, use learner
feedback to choose and calibrate the next Unicode block toward the working course.
