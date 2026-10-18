# Finite native fuzz source contract: hex-codec-v1

This contract describes the accepted Go test language for `fuzz-seed-tests`
with target `FuzzHex`. These are explicit assessment limits. A public Check
reports unsupported source before submitted test code can execute. Check uses
the learner's registered seed values; Complete additionally checks the required
defect models with those same seeds and invariants.

## API and byte values

The task supplies this exact subject API through its generated module import:

```go
func Encode([]byte) []byte
func Decode([]byte) ([]byte, bool)
```

Encode produces two lowercase ASCII hexadecimal digits per input byte. On a
value returned by Encode in the same callback, Decode returns the original
bytes and `true`. Nil and empty slices are equal byte values. Allocation,
capacity, pointer identity and ownership are outside the assessed contract.
Arbitrary malformed encoded input is outside this grammar's input domain.

Public Check runs the submitted tests against the canonical implementation and
every declared correct implementation, with repeated native seed executions.
Assertions that reject a valid representation fail this public check. Passing
public Check establishes correctness of the submitted assertions on those
implementations; full verification also requires detecting every defect model.

## Files, imports and native registration

- Submit one to eight UTF-8 basename `_test.go` files in external package
  `subject_test`, with at most 256 KiB combined source. Files must be selected
  on every platform: Go-ignored names and recognized OS/architecture suffixes
  are unsupported. Build, compiler and line directives are unsupported.
- Import only `testing`, `bytes`, `encoding/hex` and the task's generated
  subject package. Named import aliases work; dot and blank imports do not.
- Package-level declarations are imports and ordinary body-bearing functions.
  Package state, init, methods, generics, and additional native Test, Fuzz,
  Benchmark or Example entrypoints are unsupported.
- Declare exactly `func FuzzHex(*testing.F)`. Its setup creates finite `[]byte`
  seeds or `[][]byte` seed tables, may create immutable F aliases, and calls
  `F.Add`. Setup may range over finite seed tables using fresh `:=` bindings
  or no bindings. It performs no subject, library or submitted-helper calls.
- Register 1..64 seeds, each at most 4096 bytes. Typed nil and empty `[]byte`
  values work. Bare untyped nil is not an Add argument. The learner owns the
  seeds; the checker does not add a hidden corpus to make weak seeds adequate.
- The final top-level setup statement is exactly one `F.Fuzz` registration.
  Its callback is a function literal or a directly named submitted function
  with signature `func(*testing.T, []byte)`. It cannot capture setup locals.
  Native F/T aliases retain their exact pointer types; interface boxing is
  unsupported.

Native execution uses deterministic seed mode. Active randomized fuzzing and
disk corpus discovery are outside this profile.

## Callback statements and helper functions

- Use blocks, if/else with optional local init, returns, and range over
  readonly seed/result byte slices. Range uses fresh `:=` bindings or no
  bindings; `= range` is unsupported.
- Local assignments use one name with `:=` or `=`. A local `var` declaration
  uses one explicit type and one initializer or the zero value. Boolean,
  integer and string locals support ordinary typed scalar expressions,
  comparisons and integer conversions except uintptr. Scalar local
  accumulators may be reassigned; helper parameters may not be reassigned.
  Byte bindings and function aliases are immutable.
- An exact `([]byte, bool)` result is bound by two fresh `:=` identifiers from
  one statically resolved Decode or tuple-returning helper call. Blank names
  are allowed with at least one fresh nonblank binding. Tuple `var`, tuple
  reassignment, grouped declarations, inferred-type `var`, implicit tuple
  argument spreading and naked tuple returns are unsupported.
- Tuple helpers return two explicit values or forward one exact static tuple
  call. A single-result named helper may return its tracked named result.
- Read byte indices and two-index slices. Never write elements or address
  storage. `x[:]` and the exact `x[0:len(x)]` preserve byte lineage. Other
  slices lose lineage even when their runtime byte values happen to match.
  `bytes.Clone` preserves lineage.
- Ordinary for loops, switch, labels, compound assignment, defer, go,
  channels, maps, interfaces, pointers, dynamic dispatch and recursion are
  unsupported. The only permitted closure is the native callback literal.
- Use at most 16 helpers, excluding the target and a directly named callback.
  Each has at most eight parameters, including at most four byte parameters.
  Parameters are supported scalars, `[]byte` or `*testing.T`. Results are zero,
  one supported scalar/byte value, or exactly `([]byte, bool)`. Function aliases
  are local values, rather than helper parameters or results.
- All syntax is checked, including unused helpers and unreachable code. Helper
  graphs must be acyclic. Every actual call is checked with its real byte
  origins. Hypothetical helper validation gives no executed participation
  credit. For structural checking, each byte parameter receives a finite
  Seed, Encoded, Decoded or Other role. At least one role assignment must
  validate; there are at most 4^4 assignments per helper. Actual call contexts
  are checked separately. Structural analysis shares a 100,000-step budget across the whole
  source; expanded literals and byte strings are bounded at 4096 bytes before
  allocation, including gaps in keyed literals.

## Subject calls and independent checks

The callback input has Seed lineage. Encode accepts only that current Seed or
its full readonly clone. Encode's return has Encoded lineage. Decode accepts
only that callback's Encoded return or full clone. Decode's byte return has
Decoded lineage. Decoded and unrelated bytes cannot feed another subject API
call. Helper calls and return joins preserve only proven common lineage.

Import aliases and immutable local function aliases can refer to the exact
subject Encode and Decode functions. Similar names and independent library
calls do not count as subject participation.

The submitted source must contain statically resolved Encode and Decode calls.
This source-presence rule is separate from the executed participation required
for every callback.

Allowed library operations are `bytes.Equal`, `bytes.Clone`, builtin `len`, and
the exact `encoding/hex.AppendEncode` function as an independent wire oracle.
Its destination must syntactically be literal `nil` or `[]byte(nil)`;
parentheses are allowed. Its input must be the current Seed or full clone.
An alias, empty slice or cloned nil is not an allowed destination. Oracle output
has unrelated lineage and cannot replace subject Encode or feed subject Decode.
Other encoding/hex calls are unsupported. The same rules apply through import
aliases and immutable local function aliases.

An independent per-index lowercase wire check may use readonly byte indexing,
finite literals and integer/scalar calculations. A round trip by itself need
not distinguish mutually wrong encoder and decoder implementations.

## Assertions and eligible execution

Only current native T `Helper`, `Error`, `Errorf`, `Fatal` and `Fatalf` are
permitted, directly or through checked helper parameters. Diagnostic arguments
are scalars or readonly bytes. Variadic argument expansion is unsupported.
Log, skip, cleanup, parallel, subtest, name and environment methods are
unsupported. Arbitrary builtins, make/new, append/copy, external execution,
I/O and checker-observer imports are unsupported.

Every seed must enter one native callback and complete Encode at least once.
An assertion-free callback must also complete Decode. An invariant assertion
may stop after a completed Encode when its wire is wrong; Fatal and an Error
followed by return can be legitimate terminal checks. A bare early return or
an unused helper cannot replace executed participation.

Each callback permits at most eight subject calls and 64 assertions. Each API
byte value is bounded at 8192 bytes. Source, command, native output and complete
observation limits apply to the entire run; a later quota or capture failure
invalidates an earlier assertion. Complete observation is bounded at 16 MiB
and 8192 records, raw native output at 8 MiB, and converted native JSON at
32 MiB. Repetitions must agree on observed native
event structure, seed/API byte values, Decode status and callback outcomes.
Nil/empty byte values are normalized; diagnostic text and timing are excluded.

Only a complete eligible native assertion failure counts as detecting a defect.
Admission or compile errors, panic, unexpected Goexit, race, unexplained exit,
timeout, I/O failure, truncation and quota breaches are invalid, including after
a real assertion. Public failures include the relevant source rule or native
diagnostic. Full verification distinguishes ineligible execution from a defect
model that survives the learner's seeds and assertions.
