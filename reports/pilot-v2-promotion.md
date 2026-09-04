# Pilot v2 promotion outcome

Phase 3C ends with **2 promoted**, **9 returned to draft**, and **1 rejected**.
All 12 packages passed through `in-review`; every exit transition was performed
by `task-promote`. No task was moved or copied directly, and no existing legacy
ready package was edited.

## Lifecycle result

| Outcome | Tasks |
| --- | --- |
| ready | `detached-append-001`, `map-take-001` |
| returned-to-draft | `clamp-tests-001`, `counter-receiver-001`, `effective-rate-001`, `error-chain-query-001`, `formatter-contract-001`, `resource-finalize-001`, `rune-offset-tests-001`, `stream-normalize-001`, `utf8-prefix-001` |
| rejected | `typed-nil-resource-001` |

The rejected seed was not replaced. Returned packages retain their complete
review rationale and `draft -> in-review -> draft` history. The rejected package
retains `draft -> in-review -> rejected` history.

## Promoted profiles

| Task | Activity | Submission | Verification | Primary | Secondary |
| --- | --- | --- | --- | --- | --- |
| `detached-append-001` | debugging | implementation | fixed-tests | `slices.backing-array-aliasing` | `slices.append-growth` |
| `map-take-001` | implementation | implementation | fixed-tests | `maps.nil-delete-clear` | — |

Both references pass full checks. Both starters fail an obligatory check. All
eight required promoted mutants are distinct, compile, and are detected by test
failure rather than compilation failure. Public tests remain separate from the
grader and do not expose checker-owned cases.

## Pilot activity distribution

| Activity | Total | Ready | Returned | Rejected |
| --- | ---: | ---: | ---: | ---: |
| implementation | 4 | 1 | 3 | 0 |
| debugging | 3 | 1 | 1 | 1 |
| test-writing | 1 | 0 | 1 | 0 |
| test-repair | 1 | 0 | 1 | 0 |
| refactoring | 2 | 0 | 2 | 0 |
| integration | 1 | 0 | 1 | 0 |

The pilot still contains the requested 4 implementation, 3 debugging, 2 test
submissions, 2 refactoring, and 1 integration packages, but only the two ready
profiles count as accepted curriculum evidence.

## Ready coverage delta

Before Phase 3C, the three legacy ready tasks supplied 3 primary and 3 total
encounters across `types` and `functions`. The two promoted tasks raise this to
5 primary and 6 total encounters:

| Domain | Primary before -> after | Total before -> after |
| --- | --- | --- |
| `maps` | 0 -> 1 | 0 -> 1 |
| `slices` | 0 -> 1 | 0 -> 2 |

No returned or rejected package contributes mastery coverage. This is a quality
decision, not an attempt to meet minimum taxonomy counts during the pilot.

## Acceptance evidence

- each promoted task passed a separate `task-check` in ready state;
- one full `task-check -all` collection was reused by validator and report;
- deterministic `task-report -check` passed against that collection;
- `go test ./...`, `go test -race ./...`, `go vet ./...`, and
  `go test -tags taskintegration ./...` passed;
- GitHub required `validate` passed for the promotion batch before merge.

The final branch repeated the full collection-reuse and CI graph after recording
all return/reject decisions and regenerating deterministic quality reports.
