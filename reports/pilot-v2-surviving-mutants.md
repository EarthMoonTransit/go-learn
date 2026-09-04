# Pilot v2 surviving-mutant analysis

Task Contract v2 reported **50/50 required mutant IDs detected** before red-team
review. That number was technically correct but pedagogically incomplete:
`detached-append-001` and `counter-receiver-001` each contained two mutants with
the same observable behavior, and several stated requirements had no mutant or
grader evidence.

## Accepted mutation result

| Task | Before review | After review | Survived |
| --- | --- | --- | ---: |
| `detached-append-001` | 4 IDs / 3 distinct behaviors | 4/4 distinct detected | 0 |
| `map-take-001` | 4/4 distinct detected | 4/4 distinct detected | 0 |
| **Promoted total** | **8 IDs / 7 distinct behaviors** | **8/8 distinct detected** | **0** |

The slice mutant `alias-after-restore` copied values onto themselves and was
equivalent to direct append reuse. It was replaced by `reuse-when-empty`, which
returns the input alias only on the empty variadic path. Existing grader evidence
detects that separate defect.

## Why 42 historical detections were not accepted as quality evidence

Phase 3B detected all 42 declared mutant IDs belonging to the nine returned and
one rejected package. The final collection rechecks 38/38 in returned drafts;
rejected packages are audited structurally and are no longer executed. Those
results do not justify promotion because mutation coverage cannot compensate for
the following surviving defect classes:

| Task | Surviving or unmodelled defect class |
| --- | --- |
| `clamp-tests-001` | Ad-hoc tests can pass without table structure, named cases, or actionable diagnostics. |
| `counter-receiver-001` | Direct update bypasses the private receiver path; copy-before-interface and value-receiver mutants remain equivalent. |
| `effective-rate-001` | Very large integers and zero capacity are outside current evidence. |
| `error-chain-query-001` | A manually constructed `FieldError{Err:nil}` can panic; mutant count does not repair excessive scope. |
| `formatter-contract-001` | Implicit satisfaction is already implemented; Pair formatting mutants do not assess it. |
| `resource-finalize-001` | Raw write errors without context, empty-write behavior, and combined short-write/close failure are not distinguished. |
| `rune-offset-tests-001` | Adjacent invalid bytes, truncated encodings, and valid encoded U+FFFD are not distinguished. |
| `stream-normalize-001` | Correct whole-input buffering, nil-error short writes, and ASCII-only trimming can escape the current suite. |
| `typed-nil-resource-001` | Named nil maps, channels, and functions can satisfy the broad interface contract but are not covered. |
| `utf8-prefix-001` | Negative-only limit defects and valid U+FFFD handling are not independently covered. |

These are reviewed semantic conclusions, not executable mutant directories.
Phase 3C does not add mutants to returned tasks or create a replacement for the
rejected task.

## Before and after

- Before review: 50/50 declared mutants detected; at most 48 distinct observed
  behaviors after equivalence review; zero pedagogically approved tasks.
- After review: 8/8 distinct required mutants are detected in two approved
  tasks; 38/38 are rechecked in returned drafts; the rejected package retains
  historical 4/4 evidence but is not executed; zero required mutants survive in
  ready tasks.
- The aggregate quality report must therefore be read with review status and
  lifecycle state, not as a standalone quality signal.
