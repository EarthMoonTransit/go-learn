# Foundation pilot v2: mutation results

Technical checks were run through `task-check/v2` in disposable workspaces.
The 12 draft packages detect all 50 required defects. Human review remains
pending: every `review/review.yaml` has `decision: request-changes`.

| Task | Baseline evidence | Detected | Required defects |
| --- | --- | ---: | --- |
| `effective-rate-001` | starter fails fixed tests | 5/5 | fixed full rate; integer operation order; missing percentage scale; rounded result; early percentage truncation |
| `detached-append-001` | starter fails fixed tests | 4/4 | alias after restoration; direct append reuse; dropped prefix; input mutation |
| `utf8-prefix-001` | starter fails fixed tests | 4/4 | byte limit; off-by-one rune; invalid-byte re-encoding; stop at `RuneError` |
| `map-take-001` | starter fails fixed tests | 4/4 | unconditional presence; whole-map clear; retained entry; zero-as-missing |
| `counter-receiver-001` | starter fails fixed tests | 4/4 | copied receiver; ignored delta; replacement; value receiver |
| `formatter-contract-001` | starter fails fixed tests | 4/4 | concrete-only dispatch; omitted field; stringification bypass; swapped fields |
| `typed-nil-resource-001` | starter fails fixed tests | 4/4 | interface-only nil; unchecked `IsNil`; pointer-only kind; pointer type/value confusion |
| `error-chain-query-001` | starter fails fixed tests | 4/4 | shallow type assertion; dropped cause; message comparison; wrapping nil |
| `stream-normalize-001` | starter workspace fails fixed tests | 4/4 | ignored scan error; ignored write error; retained blank lines; whole-input read on error |
| `clamp-tests-001` | weak tests leave mutants alive | 4/4 | missing lower clamp; missing upper clamp; shifted lower result; shifted upper result |
| `rune-offset-tests-001` | weak tests leave mutants alive | 4/4 | every-byte offsets; missing terminal; rune ordinals; stop at invalid byte |
| `resource-finalize-001` | starter fails fixed tests | 5/5 | close before write; ignored close error; ignored short write; overwritten write error; skipped close |

## Verification interpretation

- For fixed and workspace profiles, `starter-public` and `starter-full` gates
  pass only when the starter itself fails the required checks.
- For mutation profiles, starter tests pass against the canonical
  implementation and leave at least one required mutant alive; reference test
  submissions pass the canonical implementation and kill every mutant.
- Reference submissions pass two deterministic runs, public API fingerprints
  match all variants, and `go vet` passes in isolated workspaces.
- No mutant survives. Compilation failures do not count as detections; all 50
  were detected by test failures.
- Public test files contain no grader files. Grader and reference artifacts
  remain checker-owned and are not release-bound.

The aggregate is technical evidence, not pedagogical approval.
