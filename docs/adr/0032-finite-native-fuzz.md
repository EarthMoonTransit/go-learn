# ADR 0032: Finite native fuzz-seed verification

- Status: accepted design for a separate tooling batch; no live activation
- Date: 2026-10-18
- Extends: ADR 0007 and ADR 0019

## Context and activation boundary

`fuzz-seed-tests` is reserved by Task Contract v2. Ordinary test failure does
not establish that a learner registered a native target, exercised its seeds,
called the intended implementation, or checked an invariant. A mutually wrong
encoder and decoder can preserve a round trip. A panic, race, timeout, broken
observer or truncated result can also produce a failing process.

Accept the independently reviewed finite design below. This documentation
batch changes no executable profile, schema, task, taxonomy, route or learner
state. The reservation remains enforced until a separate implementation PR
passes all acceptance gates. That PR changes schema v2 and the collection
contract as explicitly described here. New content follows in a separate PR.

The first grammar is `hex-codec-v1`, with exactly one target named `FuzzHex`.
It assesses deterministic native seed execution and value properties of a
lossless lowercase hexadecimal codec. It does not implement arbitrary Go fuzz
programs, randomized search, parser errors or crash reproduction. P21 and the
benchmark profile require their own source, assessment and tooling decisions.
No new skill id or new meaning for an existing skill is introduced.

## Declarative task contract

Add `verification.fuzz` to `schemas/task-v2.schema.json`, required exactly when
`profile` is `fuzz-seed-tests` and forbidden for every other profile. Its closed
shape has three required fields:

| Field | Finite meaning |
|---|---|
| `grammar` | Constant `hex-codec-v1`. |
| `target` | Constant `FuzzHex`; no inference from package patterns. |
| `correct_variants` | Two to four unique identifiers with the same identifier syntax as `required_mutants`. Each names an additional checker-owned correct implementation. |

Retain `activity: fuzz-test-writing`, `submission.kind: tests`, empty
`editable_paths` and `workspace_skeleton`. Require each of `packages`,
`public_checks`, `grader_checks` and `reference_checks` to equal `["./subject"]`.
They continue to select packages. Require `fuzz_seconds: 0`; no active fuzzing
or disk corpus is part of this profile. Retain `command_timeout_seconds: 1..60`
and `determinism_runs: 2..3`. `benchmark_runs` retains its existing schema range
and is unused here. `required_mutants` must be nonempty. `race_required` and
`vet` retain their declared meanings with the fuzz-specific outcome rule below.

Use the existing mutation-submission layout for learner starter/reference test
files, canonical `verification/reference`, and `verification/mutants/<id>`.
Add exact `verification/correct/<id>` directories for the declared correct
variants. Missing, extra, duplicate, unsafe or symlinked variant paths fail
validation. All implementations expose the same canonical exported API.
The two additional P20 correct variants must independently cover loop encoding
and nil versus non-nil empty decoder results; the canonical implementation is a
third value-equivalent implementation. None is a learner-editable artifact.

The exact callable subject API is non-generic, nonvariadic:

```go
func Encode([]byte) []byte
func Decode([]byte) ([]byte, bool)
```

On the exercised domain, Encode returns two lowercase ASCII hex digits for each
input byte. Decode consumes a value returned by Encode in the same callback and
returns the original bytes with `true`. Nil and empty byte slices are equal
values for this contract. Allocation, capacity, pointer identity and ownership
are not assessed. Teacher implementations and mutants must be synchronous,
must not mutate their arguments and must satisfy this finite API boundary;
this is an authoring/review obligation, not a purity theorem inferred by AST.
Their public Encode/Decode functions must not call either public API internally,
directly or through private helpers. The observer wraps learner-origin call
boundaries; it does not intercept calls inside trusted implementations. Private
implementation helpers and their library calls remain possible under the stated
authoring guarantee. Composed public-API implementations require a separately
reviewed extension and correct-variant proof before admission to this profile.
Malformed arbitrary wire input is outside this first assessment domain.

The generated module path comes from the existing task module policy; it is
not `stageb.local` from the private proof. Resolve the API by exact Go object,
package path and signature, including through import aliases and checked local
function aliases. Public native event matching uses that same generated package
identity. A similarly named function or a library oracle never counts as SUT
participation.

## Task-visible finite source language

The implementation must publish the complete frame below in the learner
documentation and link it from each task using this grammar. Unsupported source
gets a stable diagnostic with file/position and the relevant rule. These are
explicit assessment limits, not general restrictions on idiomatic Go.

### Files, declarations and registration

- Accept one to eight UTF-8 basename `_test.go` files, at most 256 KiB combined,
  in external package `subject_test`. Apply the existing submission filesystem
  policy before parsing. Build directives, compiler directives and alternate
  platform filename selection cannot remove admitted source from execution;
  reject them rather than analyze one program and execute another.
- Imports are only `testing`, `bytes`, `encoding/hex` and the generated subject
  package. Named import aliases work; dot and blank imports do not. Only imports
  and ordinary body-bearing non-generic functions appear at package scope.
  Package state, init, methods and additional Test/Fuzz/Benchmark/Example native
  entrypoints are unsupported.
- The target has exact signature `func FuzzHex(*testing.F)`. Its setup permits
  fresh finite `[]byte` values, finite `[][]byte` seed tables, immutable native
  F aliases, `F.Add` and range over finite seed tables. It performs no SUT,
  library or helper computation. It adds 1..64 learner-owned `[]byte` seeds,
  each at most 4096 bytes, then makes exactly one final top-level `F.Fuzz` call.
  Bare untyped nil is not a native Add argument; typed nil and empty seeds work.
- The callback is a function literal or a directly named submitted function,
  with exact signature `func(*testing.T, []byte)`. It cannot capture setup
  locals. Aliases retain exact native pointer types; interface boxing is not
  supported. The learner owns all grading seeds. The checker never inserts a
  seed or corpus value to turn a surviving submission into an adequate one.

### Callback, helpers and data lineage

- Permit blocks, single-name `:=`/`=` assignments and single explicitly typed
  `var` declarations (with one initializer or the zero value), if/else with
  optional local init, returns and range over readonly seed/result byte slices.
  Both setup and callback/helper ranges use fresh `:=` bindings or no bindings;
  `= range` is unsupported even for an existing mutable scalar. Local boolean,
  integer and string values support ordinary typed scalar expressions, comparisons and
  integer conversions except uintptr. Scalar local accumulators may be
  reassigned; byte bindings and function aliases are immutable. Compound
  assignment, ordinary for loops, switch, labels, defer, go, channels, maps,
  interfaces, pointers, closures other than the native callback, and recursion
  are unsupported. Helper parameters cannot be reassigned as accumulators.
- An exact `([]byte, bool)` tuple can be bound only by two fresh `:=` identifiers
  from one statically resolved Decode or exact tuple-returning helper call.
  Blank identifiers are allowed, with at least one fresh nonblank binding.
  Grouped declarations, inferred-type `var`, tuple `var`, tuple reassignment and
  implicit tuple argument spreading are unsupported. Tuple helpers return two
  explicit values or forward one exact static tuple call; naked tuple returns
  are unsupported. Single-result named helpers may use the tracked named result.
- Read byte indices and two-index slices; never write elements or address
  storage. `x[:]` and the exact `x[0:len(x)]` preserve lineage. Other slicing
  loses lineage even when an execution happens to produce equal bytes.
  `bytes.Clone` preserves the input's lineage. Integer/scalar calculations and
  byte literals can implement an independent per-index wire oracle.
- SUT Encode accepts only the current callback Seed or its full readonly clone.
  Its return has Encoded lineage. SUT Decode accepts only that callback's
  Encoded return or full clone; its byte return has Decoded lineage. Neither
  Decoded nor unrelated bytes can feed another SUT call. Actual-context helper
  checks enforce these rules at every call site; hypothetical helper validation
  never grants call-site or runtime participation credit.
- Allow `bytes.Equal`, `bytes.Clone`, builtin `len`, and the exact library
  function `encoding/hex.AppendEncode` as an independent oracle. Its destination
  must syntactically be `nil` or `[]byte(nil)` (parentheses allowed), and its
  input must be the current Seed/clone. An alias, empty slice or cloned nil is
  not a permitted destination. Its output is unrelated bytes: it cannot stand
  in for SUT Encode or feed SUT Decode. Other encoding/hex calls are unsupported.
  Import aliases and immutable function aliases retain these constraints.
- Permit only native T `Helper`, `Error`, `Errorf`, `Fatal`, `Fatalf`, through
  the current callback T or a checked helper parameter. Diagnostic arguments are
  scalars or readonly bytes. No log, skip, cleanup, parallel, subtest, name or
  environment methods. Calls are synchronous and statically resolved. Variadic
  argument expansion is unsupported. No arbitrary builtin calls, make/new,
  append/copy, dynamic dispatch, external execution, I/O or observer imports.
- Allow at most 16 helpers (excluding the target and a directly named callback),
  each with at most eight parameters, including at most four byte parameters.
  Parameters are supported scalars, `[]byte` or `*testing.T`; results are zero,
  one supported scalar/byte value, or exactly `([]byte, bool)`. Function aliases
  are local values, not helper parameters/results. Reject cyclic helper graphs.
  Check all submitted syntax, including unused helpers and unreachable code.
  Return joins preserve only proven common lineage. For structural validation
  of each helper, try the finite Seed/Encoded/Decoded/Other role combinations
  for its byte parameters (at most 4^4); at least one context must validate.
  Actual calls are then rechecked with their real origins. This role search
  shares a non-resettable 100,000-step analysis budget with all source analysis;
  it must fail explicitly when exhausted. Expanded literals/byte strings are
  bounded at 4096 bytes before allocation, including keyed-literal gaps.

Do not require one textual reference solution shape. The private proof accepts
direct/named callbacks, imports and function aliases, multiple files, helper
returns/tuples, readonly clones, per-index wire checks and the library oracle.
Production admission must preserve these accepted forms and reject invalid
counterparts with public diagnostics.

## Native eligibility and invariant detection

Run only seed mode, with the exact native target selected and no fuzz cache or
corpus. Use the isolation service, fresh temporary modules and minimal bounded
environment. Authored test code never executes in a source checkout. A trusted
stdlib-only admission helper parses and type-checks the original files before
any submitted test binary can run. Rewrite only the statically resolved native
registration/assertion and SUT call boundaries to add a checker-owned observer.
The library oracle remains distinct. Compile and run the admitted rewritten
program, not an unrelated generated test that simulates the learner's assertions.

Use one canonical admission source shared by CLI/service entrypoints and their
disposable helper. Do not assume the calling executable is `task-check`:
learner APIs and repository test binaries invoke the same service. Bind helper
requests/responses to original source hashes, generated package identity,
grammar/target, runtime and protocol versions. Verify closed, complete response
schemas before trusting rewritten source or metadata. Operation-scoped reuse
may avoid rebuilding trusted helpers/exports only for validated unchanged
inputs; no unversioned global success cache may confer admission.

For every learner seed, observe one native callback, actual API entry/return
values, native assertions and callback end. Encode must complete at least once.
An assertion-free callback must also complete Decode; an invariant assertion
may stop earlier after a completed Encode, for example when its wire is wrong.
A Fatal assertion is a legitimate terminal assertion, including before Decode;
a bare early return or uncalled helper cannot manufacture participation.

Each callback allows at most eight total SUT calls and 64 assertions. Bound
every API value at 8192 bytes, the complete observer trace at 16 MiB/8192
records, raw native output at 8 MiB and converted native JSON at 32 MiB. The
64-seed maximum within quotas has a conservative trace upper bound of
13,620,355 bytes/5314 records with this exact fixed target. Check quotas across
the whole run: a later breach invalidates an earlier real assertion.
These are admission, command and capture limits under ADR 0004/0005's existing
isolation model, not a hard process-heap limit or an OS sandbox. Scalar string
operations and diagnostic formatting are not proved allocation-bounded by the
private mechanism experiment.

Validate the entire observer protocol, native `test2json` stream and independent
raw process result. Reject unknown/duplicate/noncanonical fields, missing or
extra events, incorrect package/target/seed order, mismatched byte lineage,
untrusted receivers, pending calls, assertions after Fatal and incomplete
parent/package finality. Bind the protocol to the pinned toolchain event form;
unsupported runtime output gets an explicit invalid result, not a guessed kill.

Only a complete eligible native assertion failure is Detection. A complete
assertion-free successful run is eligible but detects nothing. Compile errors,
admission failures, panic, unexpected Goexit, race reports, unexplained process exit,
timeout, observer I/O errors, capture truncation, over-quota output and protocol
disagreement are invalid and never Detection, including after a real assertion.
Raw exit 1 alone is insufficient. For this profile this strict rule overrides
ADR 0019's generic compiled-race-failure detection rule; other profiles keep
their existing race semantics. Every invocation uses `-race` when required.

## Public, full and task-check behavior

Public learner Check applies source admission and the complete native boundary
to the canonical implementation and every declared correct variant, repeated
`determinism_runs` times. Every run must be eligible and assertion-free.
Apply vet when declared. Incorrect assumptions about nil/empty representation
must fail publicly on a valid implementation. Use the same admission and native
eligibility kernel for public and full checks; avoid a weaker public parser.

Within each implementation, repetitions must agree on canonical native
action/package/test events and the complete observer sequence, including seed
values, API entry/return values and status, assertions and callback outcomes.
Normalize nil/empty payloads by byte-value equality and ignore execution
timings, temporary paths and diagnostic text, as distinct from structural
native events. Do not require traces of different correct implementations to
be byte-identical. A repeated invalid or changing observation fails determinism.

Full verification additionally runs every required mutant with only the
learner's admitted seeds and assertions, with the same repetition/race/limits.
Every mutant must have eligible Detection on every repetition; a survivor or
invalid run fails full verification with different diagnostics. Diagnostic
teacher probes may help classify a defect but never count as learner detection.
Hidden mutation adequacy remains under the test-submission exception in
`QUALITY_GATES.md`; public Check need not disclose mutants or supply their kills.
No performance-speed threshold or randomized search duration establishes
Detection; bounded command deadlines remain eligibility limits.

Task-check requires the reference learner tests to pass every correct variant
and detect every required mutant under these rules. The starter must be
admitted, eligible and correct on all correct variants. Every starter run on
every mutant must also be eligible, with at least one required mutant producing
eligible non-Detection on every repetition. An invalid-on-mutant starter cannot
establish intentional insufficiency even if another mutant survives. Teacher
variants/mutants preserve the canonical exported API.
An unused helper or textual SUT call cannot replace executed participation.

Map every public requirement to a public regression and useful diagnostic:
wrong/missing target, seeds and registration; unsupported source/aliases and
API signatures; missing or dead participation; invalid correct-implementation
assertions; resource and protocol failures. Include passing counterparts and
wrong variants, rather than merely proving a mutant fails the full grader.
Keep weakening controls for roundtrip-only, status-only and value-only checks,
paired wrong codecs and removal of each critical seed class. Correct-form
acceptance and nil/empty equivalence are mandatory protections against grading
only one author's preferred harness.

## Migration and implementation acceptance

The executable collection contract advances from `task-check/v3` to
`task-check/v4`; update both collection/report schemas and every consumer.
Old collections fail closed, without relabeling or in-place upgrades. Include
the new grammar/observer/admission protocol versions, trusted source and runtime
identity in the checker contract/freshness boundary, alongside the changed v2
schema and existing policy inputs. After implementation inputs stabilize,
produce one fresh full `task-check -all -output` and use that same file for
validator, task-report and scoped coverage checks. This documentation-only
decision can reuse the current unchanged-input v3 collection.

Preserve task schema v1 and all existing executable-profile behavior. Existing
tasks, ids, content digests, skills, route addresses, editable solutions,
workspaces, completion evidence and progress bytes remain unchanged. The new
optional descriptor is forbidden outside its profile; it grants no old task
new fuzz credit. Do not inspect or migrate real learner data. Use temporary
learner homes to prove old public/full behavior, current/Next/readiness/resume,
stored history and read-only commands remain intact.

The private mechanism proof is summarized in
`docs/plans/fuzz-hex-mechanism-proof.md`. Its 61 compile/admission, 67 malformed
protocol and 305 native controls passed independent review. They support this
bounded design; they do not certify production integration, promotion, content,
P20 delivery, P21 reproduction or numeric/semantic foundation coverage.

Before activating the profile, add meaningful source/protocol/outcome/service,
CLI and learner public/full regressions in disposable environments. Include
generated module identities at task-id boundaries, alternate correct variants,
helper-response tampering, all accepted forms, late failures after assertions,
missing native callbacks, paired-mutant survivors and temporary-home history
preservation. Recalculate and test trace bounds if the wire representation
changes. The new metadata is rejected until its complete executor is present;
there is no partial fallback to generic test failure.

Run ordered fast tests, race tests, vet, then one taskintegration run; collect
the bank once and validate/report from that file. Use phase1 only for this
infrastructure batch, retaining explicit coverage deficits. Require independent
final diff review and exact-head successful GitHub `validate` before merge.
Content/source/profile work stays in separate reviewable batches. Foundation
release still requires actual full `-coverage` and independent semantic review.
