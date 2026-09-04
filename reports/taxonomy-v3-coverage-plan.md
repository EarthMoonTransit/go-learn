# Taxonomy v3 coverage plan

Domains: **28**. Assignable skills: **142**. Deprecated stubs: **35**. Required encounters: **409 primary / 693 total**, a change of **+13 / +99** from v2 (**396 / 594**).

Minimum tiers are 2/4 for specialist mechanisms, 3/5 for the standard baseline, and 4/6 for high-risk ownership, lifetime, encoding, and concurrency mechanisms. A ready task contributes one primary encounter; evidence records contribute none.

## Domain totals

| Domain | Assignable skills | Primary minimum | Total minimum |
|---|---:|---:|---:|
| `toolchain` | 4 | 12 | 20 |
| `types` | 7 | 21 | 35 |
| `functions` | 5 | 15 | 25 |
| `expressions` | 3 | 9 | 15 |
| `control-flow` | 4 | 12 | 20 |
| `arrays` | 2 | 5 | 9 |
| `slices` | 9 | 27 | 45 |
| `strings` | 5 | 13 | 23 |
| `unicode` | 4 | 12 | 20 |
| `maps` | 4 | 12 | 20 |
| `structs` | 3 | 9 | 15 |
| `methods` | 3 | 8 | 14 |
| `pointers` | 2 | 6 | 10 |
| `embedding` | 2 | 5 | 9 |
| `interfaces` | 7 | 21 | 35 |
| `errors` | 7 | 21 | 35 |
| `generics` | 4 | 10 | 18 |
| `packages` | 5 | 11 | 21 |
| `io` | 6 | 17 | 29 |
| `encoding` | 4 | 10 | 18 |
| `time` | 3 | 9 | 15 |
| `testing` | 11 | 32 | 54 |
| `fuzzing` | 3 | 8 | 14 |
| `benchmarks` | 3 | 7 | 13 |
| `goroutines` | 6 | 19 | 31 |
| `channels` | 11 | 33 | 55 |
| `sync` | 9 | 25 | 43 |
| `context` | 6 | 20 | 32 |

## Largest minimum tier (4 primary / 6 total)

`channels.cancellation`, `channels.close-semantics`, `channels.ownership`, `context.cancellation`, `context.deadlines`, `context.propagation-api`, `goroutines.leak-prevention`, `interfaces.typed-nil`, `slices.backing-array-aliasing`, `slices.subslice-retention`, `sync.data-races-happens-before`, `sync.mutex-critical-sections`, `sync.waitgroup-lifecycle`, `testing.deterministic-boundaries`, `unicode.invalid-utf8`

## Assignable skills without accepted Exercism evidence

This is a source-evidence gap, not a taxonomy defect and not task coverage.

- `toolchain` (4): `toolchain.build-run`, `toolchain.docs-environment`, `toolchain.formatting-vetting`, `toolchain.modules-workspaces`
- `types` (3): `types.declarations-inference`, `types.iota-enumeration`, `types.numeric-conversions`
- `functions` (3): `functions.closure-capture`, `functions.defer`, `functions.signatures-results`
- `arrays` (2): `arrays.indexing-iteration`, `arrays.value-semantics`
- `slices` (9): `slices.append-growth`, `slices.backing-array-aliasing`, `slices.copy-overlap`, `slices.full-slice-expressions`, `slices.header-model`, `slices.len-cap-slicing`, `slices.nil-empty`, `slices.preallocation`, `slices.subslice-retention`
- `strings` (3): `strings.building`, `strings.byte-rune-conversions`, `strings.immutable-bytes`
- `unicode` (3): `unicode.invalid-utf8`, `unicode.range-decoding`, `unicode.utf8-encoding`
- `maps` (4): `maps.iteration`, `maps.key-comparability`, `maps.nil-delete-clear`, `maps.operations-comma-ok`
- `structs` (3): `structs.copy-equality`, `structs.field-visibility-tags`, `structs.literals-zero`
- `methods` (3): `methods.method-sets`, `methods.receiver-semantics`, `methods.values-expressions`
- `pointers` (2): `pointers.address-deref-nil`, `pointers.mutation-sharing`
- `embedding` (2): `embedding.field-promotion`, `embedding.method-promotion-ambiguity`
- `interfaces` (5): `interfaces.dynamic-type-value`, `interfaces.implicit-satisfaction`, `interfaces.small-contracts`, `interfaces.type-switches`, `interfaces.typed-nil`
- `errors` (7): `errors.as`, `errors.is`, `errors.panic-recover-boundaries`, `errors.propagation-context`, `errors.sentinels-creation`, `errors.typed`, `errors.wrapping`
- `generics` (4): `generics.approximation-unions`, `generics.constraints-type-sets`, `generics.reusable-data-structures`, `generics.type-parameters-inference`
- `packages` (5): `packages.import-cycles`, `packages.imports-exports-naming`, `packages.initialization-order`, `packages.internal-boundaries`, `packages.test-package-boundaries`
- `io` (6): `io.buffered`, `io.copy-limit-composition`, `io.files-cleanup`, `io.pipe-lifecycle`, `io.reader-writer-contracts`, `io.scanning`
- `encoding` (4): `encoding.binary-marshaling`, `encoding.json-shape-tags`, `encoding.streaming-strictness`, `encoding.text-marshaling`
- `time` (3): `time.duration-arithmetic`, `time.parsing-zones`, `time.timers-tickers`
- `testing` (11): `testing.basics-diagnostics`, `testing.cleanup`, `testing.deterministic-boundaries`, `testing.examples`, `testing.golden`, `testing.helpers`, `testing.parallel`, `testing.race-detector`, `testing.seams-doubles`, `testing.subtests`, `testing.table-driven`
- `fuzzing` (3): `fuzzing.invariants`, `fuzzing.reproduction`, `fuzzing.targets-seeds`
- `benchmarks` (3): `benchmarks.allocations`, `benchmarks.experimental-design`, `benchmarks.loops-timing`
- `goroutines` (6): `goroutines.capture-arguments`, `goroutines.launch-lifecycle`, `goroutines.leak-prevention`, `goroutines.panic-boundaries`, `goroutines.result-error-collection`, `goroutines.state-ownership`
- `channels` (11): `channels.buffering-backpressure`, `channels.cancellation`, `channels.close-semantics`, `channels.directional`, `channels.fan-out-in`, `channels.nil-disabling`, `channels.operations-blocking`, `channels.ownership`, `channels.pipeline`, `channels.range-drain`, `channels.select`
- `sync` (9): `sync.atomics`, `sync.cond`, `sync.data-races-happens-before`, `sync.mutex-critical-sections`, `sync.mutex-ownership`, `sync.once-initialization`, `sync.pool`, `sync.rwmutex-tradeoffs`, `sync.waitgroup-lifecycle`
- `context` (6): `context.cancellation`, `context.deadlines`, `context.derivation`, `context.propagation-api`, `context.timeouts`, `context.values-boundaries`

## Over-complexity risks

- Paired toolchain and library skills retained from v2 may still hide independent commands or API contracts; split only after task evidence shows separate failure modes.
- Very small specialist skills can create artificial sequencing and reporting overhead; the three-distinct-task-form rule is the guardrail.
- Concurrency decomposition is intentionally detailed because ownership and termination bugs are independently observable, but future content review must reject near-duplicate exercises.
- Deprecated stubs increase file size and must be removed only through a later explicit migration after all external references are inventoried.
