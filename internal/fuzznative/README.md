# Finite native seed kernel

This internal package implements the bounded `hex-codec-v1` mechanism accepted
by [ADR 0032](../../docs/adr/0032-finite-native-fuzz.md). It is not connected to
tasking or learner verification yet. `fuzz-seed-tests` remains reserved by the
current checker; the v2 schema and `task-check/v3` collection contract are
unchanged. This package grants no task, promotion or coverage credit.

`Engine` receives an existing disposable `isolation.Environment`. Each native
invocation gets a new module. Original learner test sources are parsed and
type-checked by a trusted frontend built from embedded canonical sources in
that environment. The parent validates the complete response before compiling
the rewritten test program. Neither helper building nor authored execution
uses a source checkout. The engine does not own the environment's lifetime.

Requests bind the generated module, fixed grammar and target, actual queried
Go 1.27 runtime and target, helper protocol version, canonical source digest,
original source bytes and hash-pinned compiler exports. Helper and export
reuse is confined to one engine and its environment. The helper checks export
contents again on admission. This follows the existing filesystem/process
hygiene model; it is not an OS sandbox or a hard heap bound.

The teacher API is exactly `Encode([]byte) []byte` and
`Decode([]byte) ([]byte, bool)`. Teacher implementations are trusted repository
inputs: they must be synchronous, preserve arguments, and never call either
public API internally. The observer wraps learner-origin call boundaries.
The package does not infer those authoring guarantees from teacher ASTs.

`Result.Valid` requires complete agreement between the observer trace, native
test events and raw process status. Only an eligible assertion failure sets
`Detected`. Compile/admission errors, panic, unexpected Goexit, race, timeout,
I/O failures, quota breaches and truncated protocols are invalid, including
after an assertion. The canonical observation excludes diagnostic text and
normalizes nil/empty byte values; it retains API values, Decode status, native
event structure and callback outcomes for later repetition checks.

Ordinary tests run no authored callbacks. They cover the 61 independently
reviewed finite source examples, additional selection/directive/write guards,
closed helper response bindings, protocol/raw-status disagreement, semantic
canonicalization and bounded capture/filesystem reads. The source examples and
tagged synthetic codecs were authored in this repository's private mechanism
proof; no external task text, solution, vectors or tests were copied.

Tagged `taskintegration` tests exercise the standalone helper and actual
`Engine.Verify` only in disposable modules. They cover all 28 accepted source
forms with real compiler exports, three correct implementations, eleven codec
defects, repeated observations, weak round-trip/status checks, removal of
critical seed classes, all assertion wrappers, minimum/maximum task-id lengths,
late raw failures and quotas, the 64-seed maximum and race mode. Offline parser
fixtures are explicitly separate from evidence of an actual native run.

Activation is a separate tooling change. It must add the complete declarative
descriptor/layout, task-check and learner public/full orchestration, all correct
variants and repetitions, strict mutant eligibility, public-feedback/service/CLI
regressions, learner-state preservation and v4 freshness migration required by
ADR 0032. Existing profile behavior and learner data remain under their current
contracts until those gates pass. No alternate Go SDK execution is claimed by
runtime-mismatch controls.
