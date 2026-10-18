# Unmapped Exercism concepts

- Source commit: `ae1477abfeeb9fbeaf971caee220e91f9e4df790`
- Taxonomy SHA-256: `8fc463c93634564e7972117ea7f7c1f90d4e5cb82659f2f631698ef68ba4ab64`

Frozen concept inventory: **38**; mapped: **21**; ambiguous: **12**; taxonomy gaps: **2**; scope-deferred: **3**. Fuzzy matching, stemming, and topic normalization are disabled.

The authoritative map is `curriculum/mappings/exercism-go.yaml`. Candidate evidence is separately reviewed and never contributes ready-task coverage.

## Taxonomy gaps

| Upstream concept | Rationale | Candidate evidence |
|---|---|---|
| `basics` | The label is an umbrella and cannot identify one atomic current skill. | `annalyns-infiltration`, `armstrong-numbers`, `cars-assemble`, `darts`, `difference-of-squares`, `hamming` (+6 more) |
| `numbers` | The umbrella label spans literals, arithmetic, conversions, integer division, and representation. | `bird-watcher`, `blackjack`, `booking-up-for-beauty`, `card-tricks`, `cars-assemble`, `census` (+11 more) |

## Scope-deferred concepts

| Upstream concept | Rationale | Candidate evidence |
|---|---|---|
| `comments` | Comment authoring needs a separate decision on the fundamentals track boundary. | `weather-forecast` |
| `randomness` | Random number generation is a standard-library topic requiring a separate track-boundary decision. | `animal-magic`, `baffling-birthdays`, `dnd-character`, `robot-name` |
| `regular-expressions` | Regular expressions are a library topic and are not automatically part of language fundamentals. | `parsing-log-files`, `pig-latin`, `wordy` |

## Ambiguous or partial correspondence

These labels are broader than one atomic skill and must not assign coverage automatically.

| Upstream concept | Candidate skills | Rationale |
|---|---|---|
| `errors` | `errors.as`, `errors.is`, `errors.panic-recover-boundaries`, `errors.propagation-context`, `errors.sentinels-creation`, `errors.typed`, `errors.wrapping` | The source label does not distinguish creation, propagation, wrapping, typed errors, or panic boundaries. |
| `functions` | `functions.closure-capture`, `functions.defer`, `functions.function-values`, `functions.signatures-results`, `functions.variadic` | The umbrella label cannot select one of the distinct function mechanics. |
| `interfaces` | `interfaces.dynamic-type-value`, `interfaces.implicit-satisfaction`, `interfaces.small-contracts`, `interfaces.type-assertions`, `interfaces.type-switches`, `interfaces.typed-nil` | The label spans satisfaction, dynamic values, type inspection, nil behavior, and contract design. |
| `maps` | `maps.iteration`, `maps.key-comparability`, `maps.nil-delete-clear`, `maps.operations-comma-ok` | The label spans lookup, mutation, nil behavior, iteration, and key comparability. |
| `methods` | `methods.method-sets`, `methods.receiver-semantics`, `methods.values-expressions` | The label does not distinguish receiver and method-set semantics from method values. |
| `packages` | `packages.import-cycles`, `packages.imports-exports-naming`, `packages.initialization-order`, `packages.internal-boundaries`, `packages.test-package-boundaries` | The label does not identify imports, initialization, layout, visibility, or test boundaries separately. |
| `pointers` | `pointers.address-deref-nil`, `pointers.mutation-sharing` | The label spans addressability, nil, sharing, mutation, and receiver selection. |
| `slices` | `slices.append-growth`, `slices.backing-array-aliasing`, `slices.copy-overlap`, `slices.full-slice-expressions`, `slices.header-model`, `slices.len-cap-slicing`, `slices.nil-empty`, `slices.preallocation`, `slices.subslice-retention` | The label spans representation, slicing, append, aliasing, copy, and allocation. |
| `strings` | `strings.building`, `strings.byte-rune-conversions`, `strings.immutable-bytes` | The label does not distinguish byte semantics, construction, conversion, or allocation. |
| `structs` | `structs.copy-equality`, `structs.field-visibility-tags`, `structs.literals-zero` | The label spans literals, zero values, copying, equality, visibility, and tags. |
| `time` | `time.duration-arithmetic`, `time.parsing-zones`, `time.timers-tickers` | The label does not distinguish durations, parsing and zones, or timers and tickers. |
| `type-conversion` | `strings.byte-rune-conversions`, `types.numeric-conversions` | The label does not identify numeric conversion versus string and slice conversion. |

## Mapped labels

| Upstream concept | Taxonomy skill | Accepted reviewed candidate evidence |
|---|---|---|
| `arithmetic-operators` | `expressions.arithmetic` | `cars-assemble` |
| `booleans` | `expressions.boolean-logic` | `annalyns-infiltration` |
| `comparison` | `expressions.comparison` | `vehicle-purchase` |
| `conditionals-if` | `control-flow.if` | `vehicle-purchase` |
| `conditionals-switch` | `control-flow.switch` | `blackjack` |
| `constants` | `types.constants` | `savings-account` |
| `error-wrapping` | `errors.wrapping` | definition-only; no candidate review |
| `first-class-functions` | `functions.function-values` | `expenses` |
| `floating-point-numbers` | `types.floating-point-semantics` | `interest-is-interesting` |
| `for-loops` | `control-flow.for` | `bird-watcher` |
| `multiple-return-values` | `functions.signatures-results` | definition-only; no candidate review |
| `range-iteration` | `control-flow.range` | `chessboard` |
| `runes` | `unicode.rune-model` | `logs-logs-logs` |
| `string-formatting` | `strings.formatting` | `party-robot` |
| `stringers` | `interfaces.stringer-contract` | `meteorology` |
| `strings-package` | `strings.standard-library` | `welcome-to-tech-palace` |
| `type-assertion` | `interfaces.type-assertions` | `sorting-room` |
| `type-definitions` | `types.named-underlying` | `chessboard` |
| `variables` | `types.declarations-inference` | definition-only; no candidate review |
| `variadic-functions` | `functions.variadic` | `card-tricks` |
| `zero-values` | `types.zero-values` | `census` |

Topics are retained in candidate metadata as low-confidence inventory labels and never converted directly to skills. Prerequisites are dependency evidence, not skill encounters.
