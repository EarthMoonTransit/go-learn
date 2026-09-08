# Content audit: C11–C13

Independent content audit by `/root/ui_goal_reviewer`, 2026-09-09.

This report describes the candidate C13 checkpoint after acceptance: 142 routed
lessons. C13 lifecycle promotion and batch acceptance were in progress during the
audit. Its ten task-level approvals do not substitute for full collection,
required tooling checks, PR validation or merge. Three legacy ready packages
outside the route are excluded. No private learner progress was inspected.

## Recomputed evidence

Counts come from the exact IDs in `curriculum/route.yaml` and canonical
`task.yaml` / `provenance.yaml`, resolving ready, review or draft locations during
promotion. The thirty C11–C13 IDs were cross-checked against their independently
reviewed packages; stale authoring specifications were used only to identify the
batch membership, not to assess final content. The active assignable skill count
was independently read from `curriculum/skills.yaml`.

| Batch | Substantive adaptations | Originals | Implementation | Test-writing | Debugging |
|---|---:|---:|---:|---:|---:|
| C11 | 3 | 7 | 5 | 3 | 2 |
| C12 | 3 | 7 | 5 | 3 | 2 |
| C13 | 2 | 8 | 4 | 4 | 2 |
| Total | 8 | 22 | 14 | 10 | 6 |

The eight substantive adaptations comprise four Learn Go with Tests operations
(Find, Reduce, Stack and filesystem records), three Exercism operations
(FirstTurn, Schedule and IsAfternoonAppointment), and one Ardan named/anonymous
struct exercise. Their exact pinned artifacts, hashes, permissions and notices
were checked during independent task review. No reference-only material was
adapted. Originals fill specific generic, error, I/O, encoding, testing and
lifecycle gaps; they are not relabeled to satisfy the source quota.

The complete candidate route contains **82 substantive adaptations, 55 originals
and five concept seeds**: 142 lessons in total. The seeds are map-take,
detached-append, clamp-tests, exact-prefix and typed-line; they are excluded from
substantive reuse. Activity totals are **70 implementation, 39 test-writing and
33 debugging**. C13's actual 4/4/2 mix is recorded explicitly. RunNamed is an
implementation of an observable real-subtest runner, not a claim that a learner
must use an otherwise ungraded test-writing style.

There are **87 distinct assessed primary/secondary skill IDs**, up from 62 at the
C10 checkpoint. Prerequisites and preparatory reading are not assessed
occurrences. The taxonomy has 142 active assignable skills, so 55 currently have
no routed primary/secondary encounter; even the 87 encountered skills still need
their required breadth, recurrence and activity evidence.

## Distinct learning value and deliberate repetition

Generic Find requires tests across element types, including non-comparable
values, while Reduce carries an accumulator of a potentially different type in
callback order. The numeric-constraint repair checks the actual admitted type
set, including underlying named types and equivalent embedded constraints. Stack
adds a generic stateful LIFO protocol, empty results and reuse. These introduce
different generic obligations rather than renaming a numeric collection helper.

Error matching, typed extraction and panic containment assess different
boundaries: custom Is and identity, custom As with first-match traversal and
typed nils, and a controlled callback panic boundary preserving ordinary errors.
They build on earlier creation, propagation and wrapping work. Struct-copy tests
make scalar copying and shared slice storage coexist; the named/anonymous struct
repair checks real different type forms. Variadic aliasing revisits slice
sharing through expanded argument passing. Exact dynamic-type tests revisit
named types, aliases and typed nils as a focused transfer encounter.

Duration truncation, full timestamp parsing and afternoon classification share
the time package but require different observable results. Timestamp parsing
preserves the UTC instant and fractional seconds; the classifier implements a
half-open window with its own layout. Neither uses the current date or local
machine timezone. They are useful related practice, not three renamed parsers.

Bounded copying counts partial bytes and preserves errors and stream ownership.
Buffered flushing now receives a caller-owned *bufio.Writer with an already
pending prefix, so buffering and empty-input flushing are directly observable.
Scanner tests cover line semantics, bounded token sizes and original reader
errors. Filesystem records combine lexical traversal, header/body parsing,
collection failure and exactly-once cleanup. That multi-file component is a
substantive integration of earlier skills, with explicit source extensions.

JSON shape/tag repair compares decoded object meaning and retains private fields;
strict single-value decoding tests trailing input, unknown fields, partial-result
discard and explicitly permitted standard field matching/null behavior. Binary
and text marshaling share interface conventions but assess different wire
formats, valid ranges and atomic receiver updates. Independent output storage
and preserved input bytes are observable requirements in both.

External-window testing deliberately reuses simple interval behavior to isolate
the external package boundary and private representation. Diagnostic helpers
expose Helper ordering and raw Errorf arguments. Named cases create actual t.Run
children. Cleanup helpers separate acquisition from scope-owned lifetime, with
recorded callbacks and real testing integration. RecordingWriter reverses the
earlier writer-exercise role: the learner builds a controlled test double with a
total byte budget, snapshots and independent state. The new worker exercise uses
controlled overlapping callbacks and completion channels inside synctest.

No blocking cosmetic duplicate was found among these thirty packages. The
repeated storage, exact-text and error-identity themes are deliberate transfer
practice; task IDs or renamed examples alone are not counted as new learning.

## Public feedback and verification quality

Review removed hidden syntax restrictions from generic Find and admitted valid
standard-library imports and semantically equivalent numeric constraints. Required
Find defects now distinguish string and non-comparable inputs. Typed error tests
need custom As and joined first-match behavior, not only a wrapped direct match.

Bounded-copy checks capture original sentinel identity and observe ownership on
successful positive-byte copies. Its corresponding close defect is restricted
to that path. Scanner's final-line defect preserves the normal reader/scanner
error behavior, so error tests cannot incidentally kill it. The saved weak suite
with only terminated valid lines remains insufficient in a disposable checker
fixture. Strong mutation-reference counts do not describe that weak suite.

JSON tests explicitly accept standard field-level null and case-insensitive
matching while rejecting top-level null. Binary tests replace an already
populated receiver instead of testing decoding only into zero values. Filesystem
closure checks compare per-file multiplicities rather than impose an unstated
close order; separate malformed-separator and fresh-format-error defects enforce
the format contract. Its original ErrFormat is captured before task calls.

Cleanup public checks observe Helper before acquisition and retain raw Errorf
format/arguments. WriterDouble exposes its initial Snapshot through public Check.
The typed-nil defect now misclassifies only genuinely nil values. Worker tests
hold overlapping callbacks before comparing completion channels, and separate
shared-active-channel and sent-value defects protect independence and closure.
The scaffold uses synctest.Wait and deferred gate release instead of sleeps or
wall-clock deadlines; the task records race-required verification.

T09 separately protects the earlier inventory, exact label and Error-versus-
Formatter public repairs. Independently reviewed T10 adds public-only checks for
bounded-copy sentinel/ownership, flushing a pending buffer with empty text, and
partial JSON result rejection. Their temporary learner homes and reference
rechecks do not create human progress. Separate T11 race-checker work and its
acceptance remain outside this content audit; no future tooling is approved here.

## Primary-domain distribution and remaining deficits

| Domain | Implementation | Test-writing | Debugging | Highest share |
|---|---:|---:|---:|---:|
| arrays | 1 | 1 | 0 | 50.0% |
| control-flow | 5 | 2 | 5 | 41.7% |
| embedding | 0 | 1 | 0 | 100.0% |
| encoding | 2 | 1 | 1 | 50.0% |
| errors | 5 | 3 | 3 | 45.5% |
| expressions | 4 | 2 | 0 | 66.7% |
| functions | 3 | 4 | 4 | 36.4% |
| generics | 1 | 1 | 2 | 50.0% |
| goroutines | 0 | 1 | 0 | 100.0% |
| interfaces | 4 | 4 | 1 | 44.4% |
| io | 3 | 2 | 0 | 60.0% |
| maps | 5 | 3 | 0 | 62.5% |
| methods | 2 | 1 | 0 | 66.7% |
| packages | 2 | 1 | 0 | 66.7% |
| pointers | 3 | 1 | 0 | 75.0% |
| slices | 5 | 3 | 6 | 42.9% |
| strings | 6 | 1 | 2 | 66.7% |
| structs | 1 | 3 | 2 | 50.0% |
| testing | 4 | 0 | 0 | 100.0% |
| time | 2 | 0 | 1 | 66.7% |
| types | 8 | 1 | 4 | 61.5% |
| unicode | 4 | 3 | 2 | 44.4% |

The overall activity mix is balanced, but the stricter curriculum policy permits
no activity above 50% of primary encounters in a domain. Embedding, expressions,
goroutines, I/O, maps, methods, packages, pointers, strings, testing, time and
types remain above that threshold. Small denominators are visible rather than
hidden. Testing has four implementation tasks because their primary outcomes
are reusable testing tools; many test-writing activities assess other domains.
They do not cancel the testing-domain imbalance. Channels currently appear as a
secondary introduction, not a separate mature primary-domain sequence.

All explicit prerequisites of these thirty tasks resolve at their numbered route
positions. Worker launch and basic blocking channel operations are introduced
together; the provided scaffold supports this first encounter. This does not
establish channel ownership, cancellation, synchronization or leak-prevention
coverage. Further concurrency/context encounters, fuzzing, benchmarks, toolchain
project work and the remaining per-skill tiers still require their own content
and, where needed, tooling review.

## Milestone boundary

Milestone 1 remains **150 accepted learner-visible lessons**, with eight lessons
still to deliver after this candidate checkpoint. The minimum 76 substantive
adaptations is met by the candidate 82, but the route, integrated components,
progress behavior and acceptance gates must also satisfy the milestone plan.

Milestone 2 remains complete foundation: all 142 skills with required encounter
and activity coverage, approximately 450 bank tasks, appropriate verification
profiles and a successful `-coverage` run. **142 lessons are not 142 covered
skills.** The current partial bank must not be presented as full foundation.

No human completions, independent repetitions or confirmed mastery were inferred
from fixtures, content counts or source reuse. The existing learner's mastery is
unknown to this audit. This report approves the aggregate instructional review,
not future lessons, runtime migrations, unfinished tooling or batch release gates.
