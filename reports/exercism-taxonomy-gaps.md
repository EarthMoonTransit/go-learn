# Exercism taxonomy gaps after pilot seed selection

Explicit uncovered active skills: **94**. These gaps are assigned to an origin strategy instead of forcing an unsuitable Exercism candidate into coverage.

## debugging-tasks (20 skills)

These failure modes need deliberately broken code, mutants, deterministic seams, or test repair; accepting a superficially related algorithm would not provide reviewed evidence.

- `benchmarks`: `benchmarks.allocations`, `benchmarks.experimental-design`, `benchmarks.loops-timing`
- `errors`: `errors.as`, `errors.is`, `errors.wrapping`
- `fuzzing`: `fuzzing.invariants`, `fuzzing.reproduction`, `fuzzing.targets-seeds`
- `slices`: `slices.full-slice-expressions`, `slices.subslice-retention`
- `testing`: `testing.cleanup`, `testing.deterministic-boundaries`, `testing.examples`, `testing.golden`, `testing.helpers`, `testing.parallel`, `testing.race-detector`, `testing.subtests`
- `unicode`: `unicode.invalid-utf8`

## original-tasks (19 skills)

These foundation mechanisms need narrowly specified original tasks because no reviewed Exercism candidate isolates them without story, algorithm, or coverage distortion.

- `functions`: `functions.defer`, `functions.signatures-results`
- `interfaces`: `interfaces.dynamic-type-value`, `interfaces.small-contracts`
- `maps`: `maps.key-comparability`
- `slices`: `slices.copy-overlap`, `slices.header-model`, `slices.nil-empty`, `slices.preallocation`
- `strings`: `strings.byte-rune-conversions`, `strings.immutable-bytes`
- `structs`: `structs.field-visibility-tags`
- `toolchain`: `toolchain.build-run`, `toolchain.docs-environment`, `toolchain.formatting-vetting`, `toolchain.modules-workspaces`
- `types`: `types.constants`, `types.declarations-inference`, `types.iota-enumeration`

## permissive-sources (8 skills)

These mechanisms need a different permissive source with clearer language-specific evidence; Exercism candidates are absent or too indirect for seed selection.

- `arrays`: `arrays.value-semantics`
- `embedding`: `embedding.field-promotion`, `embedding.method-promotion-ambiguity`
- `generics`: `generics.approximation-unions`, `generics.constraints-type-sets`, `generics.reusable-data-structures`
- `interfaces`: `interfaces.typed-nil`
- `methods`: `methods.values-expressions`

## production-derived-tasks (47 skills)

These skills require realistic package, I/O, serialization, time, lifecycle, or concurrency boundaries that the selected foundation seeds do not establish safely.

- `channels`: `channels.buffering-backpressure`, `channels.cancellation`, `channels.close-semantics`, `channels.directional`, `channels.fan-out-in`, `channels.nil-disabling`, `channels.operations-blocking`, `channels.ownership`, `channels.pipeline`, `channels.range-drain`, `channels.select`
- `context`: `context.cancellation`, `context.deadlines`, `context.derivation`, `context.propagation-api`, `context.timeouts`, `context.values-boundaries`
- `encoding`: `encoding.binary-marshaling`, `encoding.json-shape-tags`, `encoding.streaming-strictness`, `encoding.text-marshaling`
- `goroutines`: `goroutines.capture-arguments`, `goroutines.launch-lifecycle`, `goroutines.leak-prevention`, `goroutines.panic-boundaries`, `goroutines.result-error-collection`, `goroutines.state-ownership`
- `io`: `io.buffered`, `io.copy-limit-composition`, `io.files-cleanup`, `io.pipe-lifecycle`
- `packages`: `packages.import-cycles`, `packages.imports-exports-naming`, `packages.initialization-order`, `packages.internal-boundaries`, `packages.test-package-boundaries`
- `sync`: `sync.atomics`, `sync.cond`, `sync.data-races-happens-before`, `sync.mutex-critical-sections`, `sync.mutex-ownership`, `sync.once-initialization`, `sync.pool`, `sync.rwmutex-tradeoffs`, `sync.waitgroup-lifecycle`
- `time`: `time.parsing-zones`, `time.timers-tickers`

## Taxonomy representation findings

- Bitwise expressions are not an active taxonomy v3 skill. `practice/allergies`, `practice/eliuds-eggs`, `practice/grains`, and `practice/secret-handshake` therefore remain deferred or idea-only rather than being mislabeled as arithmetic.
- Regular-expression practice and documentation-comment authoring remain scope-deferred in the Exercism mapping; they are not counted as seed coverage.
- This report records the gap without changing taxonomy v3; any new skill requires a separate ADR and migration plan.
