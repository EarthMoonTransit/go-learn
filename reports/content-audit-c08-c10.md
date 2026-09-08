# Content audit: C08–C10

Independent content audit by `/root/ui_goal_reviewer`, 2026-09-08.
This is the candidate checkpoint after C10 acceptance: 112 routed lessons,
including the ten independently reviewed C10 additions. At review time 102 were
ready and ten were drafts; the C10 PR and required validate gate must still pass.
Three legacy ready packages remain outside this route count. Task-level approval
does not substitute for batch acceptance.

## Recomputed counts

The thirty additions contain 21 substantive adaptations (six Exercism, four
Learn Go with Tests, eleven Ardan Labs) and nine originals. Each batch has five
implementation, three test-writing and two debugging tasks: 15/9/6 in total.

The complete candidate route has 74 substantive adaptations, 33 originals and
five earlier concept seeds. Its activity totals are 56 implementation, 29
test-writing and 27 debugging. There are 62 distinct primary/secondary assessed
skills; introductions may cover only part of a skill and these counts establish
neither complete foundation coverage nor learner mastery. The planned 150-lesson
checkpoint still needs 38 lessons and at least two more substantive adaptations
to meet the minimum 76-source quota. Future tasks require their own review.

Counts were computed from canonical route task IDs and task/provenance metadata,
not from stale authoring templates. Concept seeds were excluded from the
substantive adaptation count.

## Semantic distinctions and repetition

The slice exercises have different observable storage contracts: a capped shared
window, an overlapping in-place shift, an independent copy with reserved
capacity, a detached prefix, and generated sequences with a shared indexed view.
The array task additionally checks element-pointer identity and independent
arrays across calls. These are not changes of names around the same submission.

Stateful counters test captured state across independent closures; deposits test
shared pointer receivers; package inventory introduces a real two-package API
with exported fields and private counts. Record construction combines fresh
pointer results with nil errors. PointerChain distinguishes a target address
from the address of a pointer variable and tests pointer reassignment.

Constant-value tests assess supplied declarations and their uses. Month
enumeration requires real constants and an iota block; constant-kinds checks
actual declared kinds and values separately from returned values. Their roles
must not be described interchangeably as learner-authored declarations.

Formatted assignments retain width, precision and multiline formatting. The
writer greeting introduces one-call writer behavior, short-write handling and
caller ownership. NumberBox test-writing and receiver-speaker implementation
revisit interface dispatch in different activities, with mutable/custom
implementations and explicit value/pointer receiver forms. Typed-value
observations revisit inferred/zero values through interfaces as a transfer
encounter, rather than claiming another new introduction.

Map quantity removal, dictionary deletion, complete unordered enumeration and
an embedded cache differ in their state transitions and error policy. Cache
promotion/override is introduced without claiming ambiguity-resolution coverage.
Floating classification adds finite/NaN/infinity boundaries. Ordered dependency
errors, deferred cleanup and contextual error wrapping address distinct lifecycle
and error behaviors.

## Public feedback and test strength

Independent review repaired nonempty zero-reserve alias checking, mutations
outside a capped view, writer ownership on success/error/short writes, declared
constant values, both inventory setter orders and padded record labels. The
wrapping reference now honors Error() text even when a cause implements a
different fmt.Formatter, and preserves a direct cause link.

Mutation reviews separated requirements that previously shared an incidental
failure: single-receiver deposit tests remain insufficient; constant tests that
check only functions miss declaration-only defects; array address-consistency
checks alone miss a coherent source alias. Cache error identity now has a
fresh-error defect that otherwise preserves retries and zero documents.

The constant-kinds canonical reference uses a standard-library import and fixed
API declarations, so its acceptance checks guard against hidden single-file
restrictions. Mutation-tests tasks retain strong reference checks and required
realistic defects; they do not prescribe a particular test layout. Map checks
compare membership rather than iteration order. Slice and pointer tasks avoid
runtime-dependent growth capacities, GC timing and numeric addresses.

The separately reviewed T08 matrix exercises 13 defective submissions through
public Service.Check and three valid tolerance fixtures through Check/Complete
in temporary learner homes. Its omission experiment produced five outer
regression failures specifically because public verification incorrectly passed.
That tooling has its own acceptance gates. Subsequent C09/C10 feedback repairs
must retain their separate regression coverage; this audit does not declare
future tooling accepted. All execution evidence is technical, not calibration
from real learner sessions.

## Candidate primary-domain distribution

| Domain | Implementation | Test-writing | Debugging | Highest share |
|---|---:|---:|---:|---:|
| arrays | 1 | 1 | 0 | 50% |
| control-flow | 5 | 1 | 5 | 45.5% |
| embedding | 0 | 1 | 0 | 100% |
| errors | 3 | 2 | 3 | 37.5% |
| expressions | 4 | 2 | 0 | 66.7% |
| functions | 3 | 4 | 3 | 40% |
| interfaces | 4 | 3 | 1 | 50% |
| io | 1 | 0 | 0 | 100% |
| maps | 5 | 3 | 0 | 62.5% |
| methods | 2 | 1 | 0 | 66.7% |
| packages | 1 | 0 | 0 | 100% |
| pointers | 3 | 1 | 0 | 75% |
| slices | 5 | 3 | 6 | 42.9% |
| strings | 6 | 1 | 2 | 66.7% |
| structs | 1 | 2 | 1 | 50% |
| types | 8 | 1 | 4 | 61.5% |
| unicode | 4 | 3 | 2 | 44.4% |

Arrays, slices and structs now meet the per-domain activity share threshold in
this candidate route. Expressions, maps, methods, pointers, strings and types
still need balancing. The newly introduced embedding, I/O and packages domains
have only one encounter each. A repeated 5/3/2 batch split does not by itself
solve these domain deficits. Continue reporting them explicitly while using the
partial-bank phase1 gate; full foundation acceptance still requires coverage.
No learner attempts, completion or mastery were changed or inferred by this audit.
