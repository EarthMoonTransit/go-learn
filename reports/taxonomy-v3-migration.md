# Taxonomy v3 migration

Migration: taxonomy v2 → v3. Domains: **28**. Assignable skills: **142**. Deprecated compatibility stubs: **35**. v2 skills inventoried: **99**.

Operations: **64 preserved**, **34 split**, **1 merged**, **0 renamed**, **0 deprecated-without-replacement**; additive v3 skills: **8**. Legacy task resolutions: **0**.

Coverage contract changes from **396 primary / 594 total** to **409 primary / 693 total** encounters.

## Split, merge, rename, and deprecate operations

| v2 skill | Operation | v3 assignable targets | Rationale |
|---|---|---|---|
| `benchmarks.loops-allocations` | split | `benchmarks.allocations`, `benchmarks.loops-timing` | Benchmark timing boundaries and allocation reporting need different controls. |
| `channels.pipeline-fan` | split | `channels.cancellation`, `channels.fan-out-in`, `channels.pipeline` | Stage composition, fan topology, and end-to-end cancellation have different invariants. |
| `channels.select-nil` | split | `channels.nil-disabling`, `channels.select` | Select readiness/fair choice and nil-channel case disabling are independent rules. |
| `channels.send-receive-close` | split | `channels.close-semantics`, `channels.operations-blocking`, `channels.ownership` | Send/receive blocking, post-close behavior, and the right to close are independently testable. |
| `context.deadlines-timeouts` | split | `context.deadlines`, `context.timeouts` | Absolute inherited deadlines and relative timeout budgeting require different reasoning. |
| `context.derive-cancel` | split | `context.cancellation`, `context.derivation` | Parent-child derivation and cancel-function ownership have different leak and propagation failures. |
| `encoding.text-binary-custom` | split | `encoding.binary-marshaling`, `encoding.text-marshaling` | Text and binary representations have different malformed-input and framing contracts. |
| `errors.creation-propagation` | split | `errors.propagation-context`, `errors.sentinels-creation` | Creating classifiable errors and propagating contextual failures are separate API decisions. |
| `errors.typed-as` | split | `errors.as`, `errors.typed` | Designing a typed error and extracting one from a chain require different target-shape knowledge. |
| `errors.wrapping-is` | split | `errors.is`, `errors.wrapping` | Building an unwrap chain and inspecting it with errors.Is are independently testable. |
| `functions.first-class-closures` | split | `functions.closure-capture`, `functions.function-values` | Calling and passing function values is independent from capture lifetime and mutation. |
| `fuzzing.invariants-reproduction` | split | `fuzzing.invariants`, `fuzzing.reproduction` | Stating a property and preserving a minimized failure are separate review steps. |
| `goroutines.leaks-panics` | split | `goroutines.leak-prevention`, `goroutines.panic-boundaries` | Blocked-worker leaks and goroutine-local panic containment require different diagnostics. |
| `goroutines.ownership-results` | split | `goroutines.result-error-collection`, `goroutines.state-ownership` | Mutable-state ownership and complete result/error collection fail independently. |
| `interfaces.assertions-switches` | split | `interfaces.type-assertions`, `interfaces.type-switches` | Single assertions and multi-branch type switches have different control-flow and exhaustiveness tests. |
| `interfaces.dynamic-values-nil` | split | `interfaces.dynamic-type-value`, `interfaces.typed-nil` | Dynamic representation can be understood without correctly handling a nil pointer stored in an interface. |
| `io.buffering-scanning` | split | `io.buffered`, `io.scanning` | Flush ownership and token-size/scanning semantics fail independently. |
| `io.copy-limit-pipe` | split | `io.copy-limit-composition`, `io.pipe-lifecycle` | Stream copying/limiting and synchronous pipe ownership have different blocking behavior. |
| `maps.iteration-comparability` | split | `maps.iteration`, `maps.key-comparability` | Unspecified iteration order and key-type comparability are separate language rules. |
| `methods.receivers-method-sets` | split | `methods.method-sets`, `methods.receiver-semantics` | Receiver mutation/copying and method-set membership are independently testable. |
| `packages.init-order-cycles` | split | `packages.import-cycles`, `packages.initialization-order` | Initialization ordering is a runtime prediction problem; import cycles are a compile-time architecture problem. |
| `packages.internal-layout-tests` | split | `packages.internal-boundaries`, `packages.test-package-boundaries` | Internal import visibility and same/external test package access are independent boundaries. |
| `pointers.receiver-choice` | merged | `methods.receiver-semantics` | Receiver choice belongs with receiver semantics rather than general pointer operations. |
| `slices.append-aliasing` | split | `slices.append-growth`, `slices.backing-array-aliasing` | Append allocation behavior and cross-slice aliasing require different counterexamples. |
| `slices.capacity-bounds` | split | `slices.full-slice-expressions`, `slices.len-cap-slicing` | Ordinary len/cap calculations do not require full-slice capacity control. |
| `slices.growth-allocation` | split | `slices.append-growth`, `slices.preallocation` | Growth behavior and deliberate preallocation are related but independently measurable. |
| `slices.model` | split | `slices.header-model`, `slices.nil-empty` | Slice header sharing and nil-versus-empty API state are independently observable. |
| `strings.construction-conversion` | split | `strings.building`, `strings.byte-rune-conversions` | Efficient construction and representation conversion have different allocation and correctness contracts. |
| `sync.mutex-invariants` | split | `sync.mutex-critical-sections`, `sync.mutex-ownership` | Lock ownership/lifecycle and the invariant protected inside critical sections are independently reviewable. |
| `testing.examples-golden` | split | `testing.examples`, `testing.golden` | Executable examples and reviewed golden snapshots have different maintenance contracts. |
| `testing.helpers-cleanup` | split | `testing.cleanup`, `testing.helpers` | Failure attribution and resource cleanup attach to different test lifecycles. |
| `testing.parallel-race` | split | `testing.parallel`, `testing.race-detector` | Parallel test scheduling and interpreting dynamic race reports are different skills. |
| `testing.table-subtests` | split | `testing.basics-diagnostics`, `testing.subtests`, `testing.table-driven` | Basic testing diagnostics, data-driven cases, and subtest lifecycle are three distinct capabilities. |
| `types.constants-iota` | split | `types.constants`, `types.iota-enumeration` | Typed and untyped constants do not require iota, while enumeration patterns add independent declaration semantics. |
| `unicode.utf8-runes` | split | `unicode.rune-model`, `unicode.utf8-encoding` | Code-point semantics and UTF-8 byte encoding are distinct mechanisms. |

## Additive v3 skills

| Skill | Rationale |
|---|---|
| `channels.directional` | Directional channel types encode ownership and API permissions independently from runtime blocking. |
| `interfaces.stringer-contract` | Implementing fmt.Stringer exposes receiver, formatting, and interface-contract mistakes not covered by general formatting. |
| `slices.subslice-retention` | Small subslices can retain large backing arrays even when functional behavior is correct. |
| `strings.formatting` | Format verbs and argument contracts form an independently testable standard-library boundary. |
| `strings.standard-library` | Searching, splitting, trimming, and replacing with package strings is a distinct reusable API skill. |
| `sync.data-races-happens-before` | Race freedom and happens-before reasoning are prerequisites for selecting synchronization primitives safely. |
| `testing.deterministic-boundaries` | Deterministic control of time, randomness, scheduling, environment, and I/O is a prerequisite for trustworthy tests. |
| `types.floating-point-semantics` | Floating-point representation, comparison, NaN, infinity, and rounding are independently testable and were a v2 gap. |

Every v2 id appears exactly once in the machine-readable migration. Changed ids remain as non-assignable deprecated stubs; they are not global aliases.
