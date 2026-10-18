# Finite native seed kernel

This internal package implements the bounded `hex-codec-v1` mechanism accepted
by [ADR 0032](../../docs/adr/0032-finite-native-fuzz.md). Tasking uses this engine
for the explicit `fuzz-seed-tests` descriptor and learner public/full checks.
The `task-check/v4` collection binds production verifier sources and the actual
Go runtime. This mechanism and its synthetic fixtures grant no catalog or
coverage credit. New course content requires a separate reviewed content batch.

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

Tasking runs the canonical and every declared correct implementation on public
Check, repeating each with declared vet/race checks. Full verification first
requires that public correctness, then eligible detection of every required
mutant on every repetition. Starter insufficiency requires one same eligible
survivor across all repetitions. No hidden seeds repair a weak submission.
The complete learner source language is published in
[FUZZ_SOURCE.md](../../docs/FUZZ_SOURCE.md), copied readonly into workspaces and
included inline in learner statements. Existing task content identities and
learner state formats retain their meanings. Old collections must be regenerated;
relabeling v3 output does not create v4 evidence. Runtime mismatch controls are
offline metadata controls; no alternate Go SDK execution is claimed.
