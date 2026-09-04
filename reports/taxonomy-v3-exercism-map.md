# Taxonomy v3 Exercism map

- Source commit: `ae1477abfeeb9fbeaf971caee220e91f9e4df790`
- Taxonomy SHA-256: `4cc68407da591576cbce72117bc9a92993d491084ec736c26e5921323b269f16`

All **38** frozen upstream labels are explicit: **21 mapped**, **12 ambiguous**, **2 taxonomy-gap**, and **3 scope-deferred**. Accepted reviewed evidence records: **18**; coverage-eligible records: **0**.

| Upstream label | Status | Skill ids | Accepted evidence |
|---|---|---|---|
| `arithmetic-operators` | mapped | `expressions.arithmetic` | `cars-assemble` |
| `basics` | taxonomy-gap | — | definition-only; no candidate review |
| `booleans` | mapped | `expressions.boolean-logic` | `annalyns-infiltration` |
| `comments` | scope-deferred | — | definition-only; no candidate review |
| `comparison` | mapped | `expressions.comparison` | `vehicle-purchase` |
| `conditionals-if` | mapped | `control-flow.if` | `vehicle-purchase` |
| `conditionals-switch` | mapped | `control-flow.switch` | `blackjack` |
| `constants` | mapped | `types.constants` | `savings-account` |
| `error-wrapping` | mapped | `errors.wrapping` | definition-only; no candidate review |
| `errors` | ambiguous | `errors.as`, `errors.is`, `errors.panic-recover-boundaries`, `errors.propagation-context`, `errors.sentinels-creation`, `errors.typed`, `errors.wrapping` | definition-only; no candidate review |
| `first-class-functions` | mapped | `functions.function-values` | `expenses` |
| `floating-point-numbers` | mapped | `types.floating-point-semantics` | `interest-is-interesting` |
| `for-loops` | mapped | `control-flow.for` | `bird-watcher` |
| `functions` | ambiguous | `functions.closure-capture`, `functions.defer`, `functions.function-values`, `functions.signatures-results`, `functions.variadic` | definition-only; no candidate review |
| `interfaces` | ambiguous | `interfaces.dynamic-type-value`, `interfaces.implicit-satisfaction`, `interfaces.small-contracts`, `interfaces.type-assertions`, `interfaces.type-switches`, `interfaces.typed-nil` | definition-only; no candidate review |
| `maps` | ambiguous | `maps.iteration`, `maps.key-comparability`, `maps.nil-delete-clear`, `maps.operations-comma-ok` | definition-only; no candidate review |
| `methods` | ambiguous | `methods.method-sets`, `methods.receiver-semantics`, `methods.values-expressions` | definition-only; no candidate review |
| `multiple-return-values` | mapped | `functions.signatures-results` | definition-only; no candidate review |
| `numbers` | taxonomy-gap | — | definition-only; no candidate review |
| `packages` | ambiguous | `packages.import-cycles`, `packages.imports-exports-naming`, `packages.initialization-order`, `packages.internal-boundaries`, `packages.test-package-boundaries` | definition-only; no candidate review |
| `pointers` | ambiguous | `pointers.address-deref-nil`, `pointers.mutation-sharing` | definition-only; no candidate review |
| `randomness` | scope-deferred | — | definition-only; no candidate review |
| `range-iteration` | mapped | `control-flow.range` | `chessboard` |
| `regular-expressions` | scope-deferred | — | definition-only; no candidate review |
| `runes` | mapped | `unicode.rune-model` | `logs-logs-logs` |
| `slices` | ambiguous | `slices.append-growth`, `slices.backing-array-aliasing`, `slices.copy-overlap`, `slices.full-slice-expressions`, `slices.header-model`, `slices.len-cap-slicing`, `slices.nil-empty`, `slices.preallocation`, `slices.subslice-retention` | definition-only; no candidate review |
| `string-formatting` | mapped | `strings.formatting` | `party-robot` |
| `stringers` | mapped | `interfaces.stringer-contract` | `meteorology` |
| `strings` | ambiguous | `strings.building`, `strings.byte-rune-conversions`, `strings.immutable-bytes` | definition-only; no candidate review |
| `strings-package` | mapped | `strings.standard-library` | `welcome-to-tech-palace` |
| `structs` | ambiguous | `structs.copy-equality`, `structs.field-visibility-tags`, `structs.literals-zero` | definition-only; no candidate review |
| `time` | ambiguous | `time.duration-arithmetic`, `time.parsing-zones`, `time.timers-tickers` | definition-only; no candidate review |
| `type-assertion` | mapped | `interfaces.type-assertions` | `sorting-room` |
| `type-conversion` | ambiguous | `strings.byte-rune-conversions`, `types.numeric-conversions` | definition-only; no candidate review |
| `type-definitions` | mapped | `types.named-underlying` | `chessboard` |
| `variables` | mapped | `types.declarations-inference` | definition-only; no candidate review |
| `variadic-functions` | mapped | `functions.variadic` | `card-tricks` |
| `zero-values` | mapped | `types.zero-values` | `census` |

`basics` and `numbers` deliberately remain gaps because they do not identify one atomic skill. Ambiguous labels never assign encounters automatically.
