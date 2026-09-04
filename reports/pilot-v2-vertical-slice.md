# Foundation pilot v2: vertical slice

Generated from the Phase 3B authoring inventory on 2026-09-04. These are
authoring packages, not released tasks: all 12 remain in `tasks/drafts`, every
review decision is `request-changes`, and no package was promoted to `ready`.

## Profile

| Task | Activity | Submission | Verification | Primary skill | Secondary skills |
| --- | --- | --- | --- | --- | --- |
| `effective-rate-001` | implementation | implementation | fixed-tests | `types.numeric-conversions` | `expressions.arithmetic` |
| `detached-append-001` | debugging | implementation | fixed-tests | `slices.backing-array-aliasing` | `slices.append-growth` |
| `utf8-prefix-001` | debugging | implementation | fixed-tests | `unicode.range-decoding` | `unicode.utf8-encoding` |
| `map-take-001` | implementation | implementation | fixed-tests | `maps.nil-delete-clear` | `maps.operations-comma-ok` |
| `counter-receiver-001` | refactoring | implementation | fixed-tests | `methods.receiver-semantics` | `methods.method-sets` |
| `formatter-contract-001` | implementation | implementation | fixed-tests | `interfaces.implicit-satisfaction` | `interfaces.small-contracts` |
| `typed-nil-resource-001` | debugging | implementation | fixed-tests | `interfaces.typed-nil` | — |
| `error-chain-query-001` | implementation | implementation | fixed-tests | `errors.wrapping` | `errors.is`, `errors.as` |
| `stream-normalize-001` | integration | workspace-patch | workspace-tests | `io.reader-writer-contracts` | `io.scanning` |
| `clamp-tests-001` | test-writing | tests | mutation-tests | `testing.table-driven` | `testing.basics-diagnostics` |
| `rune-offset-tests-001` | test-repair | tests | mutation-tests | `unicode.range-decoding` | `testing.basics-diagnostics` |
| `resource-finalize-001` | refactoring | implementation | fixed-tests | `io.files-cleanup` | `functions.defer`, `errors.propagation-context` |

Activity totals: 4 implementation, 3 debugging, 2 test submissions, 2
refactoring, and 1 integration. Verification totals: 9 fixed-tests, 2
mutation-tests, and 1 workspace-tests.

## Authoring outcome

- Every `TASK.md` has one context sentence, an explicit Go API, 4–6 observable
  requirements, constraints, and fewer than 160 words.
- Every package has progressive hints, review notes, provenance v2, a weak
  starter/submission baseline, a passing reference, and at least four defects.
- No task uses the network, filesystem, clocks, randomness, sleeps, goroutines,
  or external dependencies.
- Public and grader checks are stored separately for fixed/workspace profiles.
  Test-submission profiles expose only learner tests and checker-owned
  implementations.
- No race run is required because concurrency is not assessed.

## Deliberate review questions

1. `error-chain-query-001` covers construction, `errors.Is`, and `errors.As`;
   reviewers must decide whether the secondary surface is too broad.
2. `typed-nil-resource-001` supports every nil-capable dynamic kind, not only
   pointers; this is precise but reflection-heavy.
3. `stream-normalize-001` requires emission of completed lines before a later
   read error; this is a useful streaming discriminator but a strong contract.
4. `counter-receiver-001` uses an unexported interface to make method-set
   behavior observable; reviewers must confirm that the extra mechanism earns
   its complexity.
5. `utf8-prefix-001` preserves invalid source bytes rather than normalizing
   `utf8.RuneError`; this must remain explicit during review.
6. `resource-finalize-001` preserves simultaneous write and close errors with
   `errors.Join`; the policy is intentional but more advanced than basic defer.

## Contract v2 findings

- `fuzz-seed-tests` is schema-valid but reserved and non-executable in
  `task-check/v2`; the Unicode verification task therefore uses
  `mutation-tests`.
- A quality report says `starter_status: pass` when the gate correctly proves
  that the starter fails. The field names gate success, not learner code
  success, and is easy to misread.
- Draft reports set `promotion_eligible: true` after technical success even
  while review evidence fails. Eligibility is really “eligible to enter
  review,” but the field has no destination qualifier.
- Validator cannot accept `-checks` together with `-run-commands`, so reusable
  v2 checks and legacy command validation require two invocations.
- Repository tests require the committed aggregate report to match the current
  task inventory; adding drafts makes fast tests fail until `task-report`
  regenerates the report.
- Go's test cache does not observe task files read dynamically by integration
  tests. A plain tagged run may reuse an earlier result after content changes;
  Phase 3B also ran the suite with `-count=1` to force current evidence.

These findings are documented only. Phase 3B does not change Contract v2.
