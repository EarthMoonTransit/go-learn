# Finite tagged-text reproduction mechanism

Normative companion to [ADR 0033](../adr/0033-finite-text-reproduction.md).
Status: independently reviewed design; **no implementation, parser execution or observed proof**.
Every control below is a future obligation. Any executable prototype needs its
own independently reviewed exact source/diff and disposable launcher first.

## Source identities and assembly

Use the registry/commit/path/hash in ADR 0033 and the admitted source record.
The adaptation mapping retains the count, 0/1/2/3 tags, two shared text fields,
two-byte number and the helper's unchecked payload/tail return expression.
Local additions are the readonly guards, error policy, renamed Record API,
input-value preservation rule, helper bool result and missing-length guard.
The numeric digit policy is local; do not attribute it to upstream Atoi.
The readonly scaffold returns an ordinary malformed error for a text tag without
a following length byte **before** entering takeText. There is no text_enter
on that path; a present zero length is a different, ordinary helper call with
declared0. The helper's own empty-input guard is still required and independently
checked by direct teacher behavior tests, which confer no replay witness.

Four separately hashed source manifests are mandatory:

| Assembly | Exact source / purpose |
|---|---|
| Raw donor | Exact admitted file bytes, unchanged package/functions, separate module. Locally authored driver only. Establish selected fault relation, not global equivalence. |
| Adapted original | Readonly scaffold and Cause type plus frozen helper. It retains the unchecked payload and tail expressions in one return node, after a missing-length guard, with the new bool result. |
| Wrapped original | Deterministically instrumented adapted original plus byte-identical submitted regression. Actual Decode entry and pre-helper checkpoints; catch only a proven selected original fault, then synthetic no-records/error continuation. |
| Fixed / correct variant | Same scaffold, Cause and submitted regression; learner or specified correct helper. No expected-fault continuation. Teacher behavior tests run in separate assemblies. |

The future source manifest lists sorted relative regular-file names, bytes'
SHA-256, modes, module/API object identity, grammar/rewriter/protocol/kernel
versions and compiler flags. Original excludes the learner repair path; fixed
excludes original and correct variants. The regression hash must be identical
in all replay assemblies before rewriting. The helper/scaffold/public API and
Cause objects are type-checked together; duplicate definitions cannot shadow
them. Rewritten-source manifests and test binary digests bind actual execution.
All generated imports, names and files are checked for collisions; deterministic
fresh names must not depend on learner-controlled identifiers alone.

Maintain a source-site manifest with the frozen helper function object and the
complete return-node byte/line span, including both payload `[:N]` and tail
`[N:]` bounds sites. Disable inlining for the subject compilation with explicit
pinned compiler flags and verify the flags/binary binding. A native recover
snapshot must contain `runtime.Error` and the frozen helper frame in that span,
in an active genuine Decode/pre-helper lineage. Missing/ambiguous frames fail
closed. Runtime frame names and source path mapping must be proved on each
supported pinned runtime/GOOS/GOARCH; unknown formats are unsupported, never a
guessed witness. No one-PC, one operator, evaluation-order or panic-text rule.

Go slice high bounds can use capacity; the omitted high bound of a tail is
length. Preserve the caller's actual representation when invoking the subject.
Record visible bytes/len/cap before evaluation; never read outside len or
normalize cap. See the primary [slice specification](https://go.dev/ref/spec#Slice_expressions)
and [frame API](https://pkg.go.dev/runtime#CallersFrames). Immutable factual
references are registry `golang-go`, commit
`2c31595783c74564c1299a5401dc6e606f379cc7`, paths `doc/go_spec.html` and
`src/runtime/symtab.go`; these are facts/API references, not copied task code.
Source/stack deductions alone do not prove the native classification.

## Closed learner language

Publish this entire language in the canonical readonly
`docs/REPRODUCTION_SOURCE.md` at activation, copy it into each task, include its
hash in freshness and expose it in learner statements/export. A shortened
prompt links it. Unsupported forms receive file:position plus a stable rule id.
The rules describe assessment limits, not general Go style requirements.

### Files, types and declarations

Exactly `subject/repair.go` in package subject and
`subject/regression_test.go` in package subject_test, combined at most 128 KiB
UTF-8, each at most 64 KiB and 16,384 AST nodes combined. Reject platform/build,
compiler/linker/line directives, alternate files, cgo and hidden source selection.
Comments are otherwise inert. Analyze all declarations and unreachable syntax.

Repair has exactly takeText plus at most eight private helpers. Regression has
ReproductionCause, 1..16 exact native `func TestX(*testing.T)` functions and at
most sixteen private helpers. Names must satisfy Go's native Test discovery
rule. No other Test/Fuzz/Benchmark/Example entrypoints, init, TestMain, globals,
package constants, type declarations, methods, generics or variadic functions.
No result names, naked returns or function literals. Report getter is not a
test. Native entrypoints may not call one another or be aliased. No submitted
expression may refer to ReproductionCause as a value or call target, directly
or through an alias, even in dead code. Only generated trusted TestMain calls
it, once before m.Run. A private pure construction helper may serve both the
getter and learner tests, with ordinary per-invocation budget charging. Include
accepted shared-helper and rejected direct/aliased-getter controls.

Imports in repair: only bytes. Imports in regression: only testing, bytes and
the generated task subject path. Named aliases are accepted, blank/dot imports
are rejected. Resolve actual package objects from trusted compiler exports;
never use spelling as proof. No arbitrary errors, fmt, reflect, runtime,
filesystem, process, network, unsafe, observer or external imports.

Supported value types are bool, byte, int, string, []byte, [][]byte,
subject.Record, []subject.Record, subject.Stage and subject.Cause. Records,
Cause and their fields retain readonly checker-owned type identity. Error is
only an opaque value from Decode, usable for nil comparison, forwarding and
local binding; no type assertions, method calls, formatting or fabricated
error. *testing.T is only the actual current native argument, an immutable
local alias or a checked helper parameter. No other pointers, arrays, maps,
channels, interfaces, function-valued parameters/results or type aliases.

Private helpers have at most eight parameters and zero to three results, all
from those types; *testing.T cannot be returned. Repair cannot reference subject
Decode/Record/Cause, tests, T or report helpers. The report call graph cannot
call Decode or T. Test helper graphs may call Decode/T with proven objects.
All graphs are acyclic with maximum depth sixteen. Aliases of ordinary helper,
Decode and permitted bytes functions are local single-assignment bindings,
not mutable storage; their ultimate object/signature determines every rule.

### Expressions and statements

The allowlist is exhaustive; Go type checking also applies:

- Parentheses; supported identifiers; bool, integer, rune and string literals;
  nil for slices/error; untyped constants must fit int64 before conversions.
  Integer constants are limited to 64 lexical bytes; decoded strings to 8192
  bytes. No floats, complex, unsafe pointers or constant exponent explosions.
- Unary `+`, `-`, `!`; binary `+`, `-`, `*`, `/`, `%`, `==`, `!=`, `<`, `<=`,
  `>`, `>=`, `&&`, `||` on supported Go operands. String concatenation is
  bounded below. No shifts or other bitwise operators in v1.
- Read indices of slices/strings; two- and three-index slices of byte slices,
  and two-index string slices. All omitted bounds retain normal Go semantics.
  Slice constructors `[]byte{...}`, `[][]byte{...}`, `[]subject.Record{...}`;
  positional slice entries only, at most 64 table rows / 255 records. Nested
  byte entries may use Go's elided literal type. Keyed Record/Cause literals
  accept each actual field at most once, including omitted fields/defaults.
  Field reads on the exact Record/Cause types are accepted.
- Conversions only between supported numeric types, Stage/string, []byte/string,
  and nil to supported slices. Calls only to resolved submitted helpers,
  Decode, the allowed T methods, bytes.Equal, bytes.Clone, and builtins len,
  cap, make, append, copy. len/cap apply only to supported Go operand kinds.
  make constructs only []byte with length and optional capacity. append acts
  only on []byte with byte arguments or one final []byte expansion; expansion
  is otherwise prohibited. copy takes two byte slices. No panic, recover, new,
  delete, clear, print, min/max or dynamic dispatch.
- Blocks; empty statements; local `var x T` or `var x T = expression` and
  `const x = scalar-constant`; assignments and short declarations of one or
  more locals with explicit arity (or one exact tuple call). Blanks allowed;
  no implicit tuple spreading into call arguments. Mutable parameters/locals
  may be rebound except T/function aliases. Element assignment to []byte is
  allowed and checked by actual input-value behavior tests; record/Cause
  field writes and assignment into tables/record slices are unsupported.
- If/else, including simple local init; expression switch with optional init,
  case expressions and default, no type switch/fallthrough; explicit returns
  with exact result arity, or forwarding a resolved exact tuple call.
- Range only over []byte, [][]byte and []Record, with fresh `:=` or no bindings.
  Break/continue without labels are accepted in their Go-valid context. No
  ordinary for loops, iterator ranges, goto, labels, defer, go or select.
  No compound assignment or increment/decrement; scalar `x = x + 1` works.
- Standalone calls only for no-result helpers, Decode (an executed direct
  no-panic regression is valid), copy, and permitted T methods. A call to
  takeText always remains on the typed private helper path.

Allowed T methods: Helper, Error, Errorf, Fatal, Fatalf, through the current
native receiver only. No Run, Parallel, Skip, Cleanup, Log, Name, environment or
deadline APIs. Formatting accepts a constant format of at most 256 bytes with
only `%%`, `%v`, `%s`, `%q`, `%d`, `%x`, `%t`, no flags, widths, precision or
argument indexing; at most eight scalar/string/byte arguments. Non-f methods
have the same argument cap/types. Errors count as failed assertions; Fatal
also records Goexit as an expected assertion termination. A bare Goexit cannot
be admitted and is tested only as a privileged runner-boundary control.

Admission records each native test, static Decode callsite, helper object and
T assertion site using original relative path and byte offsets. All compiler
selections, signatures, alias assignments, receiver origins and call graph
edges are checked before rewriting. Do not execute initializers to discover
objects. The frontend is trusted stdlib-only code in a disposable helper;
compile exports come only from pinned stdlib and checker source. The same
frontend/rewriter/validator must serve task-check and learner public/full.

At least these alternatives must pass the later proof: guard-before-split,
combined len/extent guard, switch guard and copied-payload helpers; table,
guard/error and direct no-panic regressions; import/function/T aliases and
acyclic tuple helpers; literal/string/append/clone/subslice representations,
both text tags and ordinary spare capacity. A source form passing admission
does not prove behavior, report truth or actual witness execution.

## Resource arithmetic

Admission has a single non-resettable 200,000-step budget across original
syntax/object/call-graph analysis. Charge every visited node and graph edge;
no recursive context search or uncharged helper expansion. Depth <=64 in AST
and protocol data. Exceeding source/analysis quotas fails before learner-native
launch. Runtime has one 250,000-operation budget per child, shared by report,
all tests and all helper calls; charge function entry, statement, loop iteration
and each allocating expression before evaluation. Never reset per test/helper.

Instrumentation evaluates each operand once in original Go order and charges
before allocation. It must preserve short circuiting, argument evaluation,
assignments, ranges and original positions. The later diff must prove these
transformations; unsupported lowering is an admission error, not a silent
change of program. Before make/append/clone/conversion/concatenation, check sizes
with overflow-safe arithmetic. Each byte backing request/string is <=8192;
Decode visible input <=4096 and cap <=32768. Use the real Go make/append/clone
operations after the guard, preserving evaluation, values and alias behavior;
do not replace append with a different allocation policy. For append/clone
charge a conservative 32768-byte backing allowance before evaluation, even
when no allocation occurs, and check the actual capacity <=32768 afterward.
The pinned runtime's byte-allocation growth bound for <=8192 requested bytes
must be established in the source/runtime proof before execution approval.
If the destination already has sufficient capacity, charge that same allowance;
never append beyond requested length8192. Make's requested cap is honored and
charged directly. String conversions/concatenation charge their resulting byte
length. Spare capacity is admitted. Total charged payload allowance is <=4 MiB
per child, even if reused/freed; tables <=64 rows, records
<=255. Do not claim a hard process heap bound: runtime/GC/compiler/metadata
overhead is governed by command deadlines/capture limits and ADR 0004/0005.

The allowance is a conservative logical charge, not measured heap usage.
Complete charge table (use checked unsigned arithmetic before allocation):

| Operation | Charged bytes |
|---|---|
| []byte literal with n entries, make([]byte,n,c) | 32+n for literal; 32+c for make (c=n when omitted), including zero length. |
| [][]byte literal with n rows | 32+32*n for its backing/header; each nested literal charges separately. Existing slice-valued entries add no duplicate payload charge. |
| []Record literal with n entries | 32+64*n for its backing/header; each Record literal charges 64 again conservatively; any new string expressions charge separately. |
| Record or Cause literal | 64 or 128 respectively, even if compiled on stack. Existing field values are copied; their already allocated backing is not charged again. |
| byte append or bytes.Clone | 32+32768 per invocation, including zero-size/no-allocation cases. |
| []byte(string), string([]byte), string concatenation | 32+result byte length per invocation, even if optimized to avoid allocation. |
| readonly Decode records allocation | 32+64*C before the actual make, including C=0. |
| original helper's string conversion | 32+N at the trusted pre-helper checkpoint before invoking the unmodified unchecked return. Fixed/correct helper conversion is charged at its actual expression instead. |

String literals are charged 32+decoded length each time their expression is
evaluated (including constants at each use); integer-to-string conversion is
unsupported. Slicing, assignment, range binding, field reads and copy allocate
no new backing and add no allocation charge. They still consume operation
budget. Zero declarations of scalars/structs/nil slices allocate no backing;
stack/local-value overhead is outside this allowance. The permitted target
runtime has pointer/int size <=8, string header <=16, slice header <=24 and
Record size <=40; admission verifies those target sizes from compiler metadata
before relying on the conservative 32/64/128 charges. Unsupported widths fail
closed. Observer snapshots/sidecars/compiler memory are governed by their own
finite event/capture quotas, not falsely described as learner backing bytes.
Nested constructor charges accumulate, without resetting at a table row,
helper, test or record. The later proof needs largest-admitted and just-over
repeated byte/table/record/Cause constructors, nested/shared rows and readonly
C=255 allocation, with native value/alias behavior preserved.

Exactly 1..16 tests and one report, <=64 Decode calls over the whole child,
<=256 assertions and <=3840 pre-helper events. The latter limit is cumulative
across calls; valid large prefixes may consume it. Each input has <=4096
bytes, each text value <=255, and up to 255 records with two strings and an int.
Theoretical sum of returned text payload is bounded by the consumed input
bytes (overwrites only reduce it); enforce sum <=4096 independently.

Use JSONL trace <=16 MiB, <=8192 events, each <=64 KiB; raw combined native
output <=8 MiB; test2json <=32 MiB; frontend request <=2 MiB, response <=4 MiB.
Byte strings use canonical padded base64, not JSON text escapes. Record outputs
encode string byte values as base64. Bound names/sites/ids to 128 ASCII bytes,
diagnostics to 8192 bytes and stack frames to 64 frames with 512-byte paths.
Private frame sidecar has a separate 4 MiB ceiling. Trace text-return events
carry remainder hash/length rather than duplicating every suffix. Per-kind
serialized ceilings, including the event envelope and newline, are 1280 bytes
for text_enter, 1024 for text_return, 2048 for expected_fault, 8192 for
call_begin, 32768 for call_end, 1024 for assertion/test_begin/test_end, and
16384 for begin/report/end. Expected_fault replaces a text_return and occurs
at most64 times. Therefore a conservative trace ceiling is
`3840*(1280+1024) + 64*1024 + 64*(8192+32768) + 256*1024 + 32*1024 + 3*16384`
= 11,878,400 bytes and at most8099 events, below the enclosing ceilings.
Unknown or over-limit values invalidate the whole child. Prototype must compute
an exact serialized worst-case bound for the final closed wire shape and
exercise the largest admissible trace; the caps above are ceilings, not an
already-proved trace-size theorem. If limits conflict, revise the design before
claiming maximal-form acceptance. A quota failure is never a semantic rejection.

Build, vet, frontend, native and converter are separately counted subprocesses,
each with the declared 1..60-second deadline, at most one concurrent native
child. Original/fixed/each correct variant run in separate clean disposable
workspaces for each of 2..3 repetitions. No fuzz campaign/cache/corpus. Compiler
or worker reuse is operation-scoped to validated identical identities only;
native success and learner observations are never reused. Preserve input hashes
before/after; remove temporary homes/modules/binaries on all exits.

## Privileged observation and finality

Threat boundary is the finite admitted source plus the existing disposable
isolation model, not an OS sandbox for arbitrary Go. Learners cannot import or
name observer APIs, forge events, access trace paths, run processes or share
global state. Trusted instrumentation gets a private observer package and a
run-scoped trace destination unavailable to the learner language. The native
observer is not trusted simply because stdout resembles JSON.

Frontend protocol version `text-admission/v1` is a strict closed request/response:

| Message | Required fields |
|---|---|
| Request | version, grammar, module, runtime identity, kernel SHA-256, source manifests, source bytes, trusted exports manifest, limits |
| Response | same identity, request SHA-256, original source hashes, accepted bool, diagnostics, typed site/test/object map, rewritten sources and hashes |

Rejected responses have empty rewrite/map; accepted responses have a complete
map and exact source selection. Reject duplicate, unknown, case-aliased fields,
trailing JSON, unsafe paths and hashes disagreeing with returned bytes. The
parent recomputes response/hash/identity checks, not just accepted=true.
Runtime identity includes full Go version, executable/export fingerprints,
GOOS/GOARCH, race mode and compiler flags; changed kernel/worker/public contract
or originals require fresh admission. No current-executable assumption.

Observer protocol `text-observer/v1` is strict JSONL. Every event has exactly
version, monotonically contiguous seq starting at 1, run_id, kind and payload.
run_id is a fresh parent-generated 128-bit nonce bound to the assembly/request.
Every payload below is closed; missing/noncanonical/extra/duplicate keys fail.
Numbers are bounded canonical nonnegative integers except Value, exit and the
three Cause numeric fields (Go int values, so a negative answer can receive a
semantic Cause diagnostic rather than masquerading as a broken protocol).
Hashes are 64 lowercase hex; all byte values are canonical base64. No timestamps
or arbitrary labels in semantic matching.

| kind | payload fields and rule |
|---|---|
| begin | request_hash, assembly_hash, regression_hash, repair_hash, runtime_hash, phase (`original`, `fixed`, `correct:<id>`), repetition, tests (ordered exact names) |
| report | case, stage, payload_offset, declared, available; emitted from the actual ReproductionCause return before tests |
| test_begin | test; exactly once for every declared test, in the pinned native order |
| call_begin | test, call (1..64 contiguous), site, occurrence, input, length, capacity; actual Decode argument captured before subject evaluation |
| text_enter | call, ordinal, record, tag_offset, tag, payload_offset, declared, available, helper_input_hash; readonly scaffold immediately before actual takeText call |
| text_return | call, ordinal, value, rest_hash, rest_length, ok; only after actual helper return; actual values are preserved even when a fixed helper returns the wrong value/suffix/status |
| expected_fault | call, ordinal, source_hash, site_span, frames_hash; only real original runtime fault with retained private raw frame evidence |
| call_end | call, outcome (`return` or `synthetic-error`), records, error_present, input_after; matches exactly one open call |
| assertion | test, site, ordinal, fatal; emitted only for executed Error/Errorf/Fatal/Fatalf, all are failures |
| test_end | test, outcome (`pass`, `assertion-fail`, `fatal`); defer records completion but does not manufacture success |
| end | test_count, call_count, text_count, assertion_count, expected_fault_count, m_run_exit, budget_used, allocation_charged, prior_events_hash |

The trace is append-only through the trusted writer; I/O error poisons the run.
end is emitted by generated trusted TestMain only after m.Run returns, then
the trace is flushed/closed and the process exits with the native code. A
failed write/flush/close must force a nonzero distinct runner failure. Native
exit zero plus a syntactically valid prefix without end is insufficient.
Store the private raw-frame sidecar under the same bounded writer and bind it
by frames_hash. It contains panic runtime type and complete normalized frames,
never learner-readable secrets. The validator rechecks source/frame mapping.

One synchronous test/call is active at a time; helper nesting stays within it.
At text_enter independently parse the captured input from offset zero through
the declared prefix, without invoking task/helper code. Recompute current
record, tag position, extent, remaining visible bytes and helper-input hash.
Match the actual pre-helper checkpoint to those independent facts. Every return
matches an actual open call/checkpoint; matching its values to the helper
contract is a separate semantic assertion, not a wire-integrity prerequisite.
A selected fault has
present length and declared>available, real runtime.Error and exact frozen
helper-site frame while this checkpoint is open. For minimal credit additionally
require C=1, first tag1/2, length3, declared1, available0, payload offset3.

Only original may substitute synthetic error after expected_fault. Emit no
text_return for that interrupted helper. Any other panic aborts/poisons the
run; never recover it as a success or produce a fallback. Fixed faults may be
captured diagnostically by a trusted outer boundary but remain invalid. Check
caller visible input values before/after even when original recovers. More
than one outstanding checkpoint, wrong ordinal, a return without an open helper,
unattributed fault or checkpoint after terminal failure is invalid.

For fixed replay, an authentic wrong value, remainder hash/length, ok or changed
visible input produces a public `behavior` assertion with actual/expected
facts. Do not label it forged protocol merely because it violates the task.
Mark that call semantically diverged, preserve the actual values and continue
the learner child unchanged. Subsequent checkpoints in this diverged call are
checked for actual trusted object/call/ordinal binding, structural consistency
and quotas, but cannot grant witness credit or be required to match the
normative parser cursor: an honest wrong remainder may change its control
flow. Each new Decode call starts its own snapshot and nondiverged independent
model. Subsequent unexpected panic, budget breach or incomplete finality still
invalidates the entire submission; an earlier behavior assertion does not
rescue it. No wrong fixed output is replaced by teacher output.

Original and trusted correct-helper mismatches instead indicate a broken
trusted assembly and invalidate it; do not blame a learner for faulty teacher
code. Original faults are allowed only with the independently valid prefix and
real selected helper-frame evidence above. A fixed call is paired for success
only if it never diverged, returned the required malformed result, preserved
input and belongs to a fully valid complete run. The later proof must retain
an authentic complete wrong-helper replay with a live minimal call, its useful
behavior diagnostic, and a late-invalid counterpart that overrides it. These
engine boundary controls are separate from product execution ordering.

Run native binary with exact declared test selection and no cache, capture raw
process exit separately, then convert the complete captured output with pinned
`go tool test2json`. Validate its full action/package/test stream against the
observer and independently discovered tests: exactly one run and matching
terminal outcome per test, exactly one final package result, no skip/pause/
continue/subtest/unexplained exit, no pending calls, and actual child termination.
A passing child requires exit0, package pass, every test pass, zero assertions,
successful observer end/close and no truncation/race/deadline/runner error.
Completed assertion failures require corresponding native failed tests/package
and exit1; unexplained exit1 is invalid. A genuine early fault never overrides
a late assertion, panic, race, timeout, missing end or converter discrepancy.

Repeats compare normalized complete semantic events and return values within
each assembly. Exclude nonce, capacity, temp paths, frames/PC, panic/diagnostic
text and timing; retain visible inputs, test/site/occurrence, causal facts,
returns, errors-as-presence and finality. Pair each minimal original call with
the same test/site/occurrence and bytes in fixed and correct assemblies; require
the expected local malformed result there, without fallback. Cause must match
one of these live minimal calls. Additional same-cause faults at later fields/
records are legal but cannot satisfy Cause/minimality. Repetition success cannot
rescue an invalid run or replace missing submitted regression.

## Collection evidence

Activation adds a closed optional `reproduction` object to QualityReport,
required exactly for reproduction-tests and forbidden elsewhere. Required keys:
grammar, protocol, kernel_digest, contract_digest, source_manifest_digest,
submissions, passed. Each submission has id, patch_digest, admission,
expected_rejection and checks. id is starter/reference/a declared required wrong
patch; each occurs once. expected_rejection is null for starter/reference and
the declared category for wrong patches. Learner service uses the same structure
with id learner and null expected_rejection, without author-only patch rows.
No report_digest is invented until the actual report has executed.

### Exact plan and execution order

Task-check schedules public and full checks for starter and reference, and full
checks for every required wrong patch (1..16). Learner Check/Complete schedules
just public/full respectively. Full behavior assembles public+grader tests;
public behavior contains public tests only. A check has scope, repetitions,
rows, behavior, assertions, witness, cause, finality, classification, passed.
The five obligation fields are `pass`, `semantic-fail`, `invalid`, or `not-run`.
Admission is pass/invalid and is performed once per submitted source identity,
before any behavior or replay native launch. Native builds still type-check each
exact assembly; admission reuse never skips compilation/source binding.

Every check predeclares a rectangular row inventory in this order:

1. Behavior rows for r=1..R, assembly `behavior-public` or `behavior-full`.
2. Replay rows for r=1..R, in assembly order original, fixed, then sorted
   correct:<id> (K=2..4). Every row uses fresh disposable native state.

Thus there are R*(3+K) rows/check, R=2..3, at most21. A task-check has at most
(4+16)*21=420 planned rows, plus one admission per submitted patch. A row may
include source validation, build, optional vet, one native child and conversion;
these are separate counted processes, not five native runs. One native child
runs at a time. The original and teacher variants never contain the learner
repair. Behavior rows never supply replay witnesses.

On admission invalidity all planned rows are not-run. On an infrastructure or
invalid-native result, stop scheduling children for this submission, retain
that actual result and explicitly block its remaining rows. A complete behavior
assertion failure still allows the remaining behavior repetitions to finish;
if any fail, all replay rows are not-run with reason behavior-dependency.
Behavior requires all repetitions pass before replay starts. Once replay starts,
complete the entire planned matrix despite semantic/assertion failures; stop
only for invalidity/infrastructure. This exposes late invalidity and permits
paired witness/Cause evaluation without silently dropping failing children.
Changing outcomes or normalized semantic traces across repetitions is invalid,
not a semantic wrong-submission rejection.

### Closed honest row forms

Every row has exactly key `{scope,assembly,repetition}` and a discriminator
state, plus the fields of one form below. Hashes are lowercase SHA-256 of real
retained bytes; null means absent, never an all-zero placeholder. Raw sidecars
are saved with source/evidence manifests; no real learner/private paths.

| state | Additional required fields |
|---|---|
| `not-run` | reason (`admission-dependency`, `behavior-dependency`, `invalid-predecessor`, `infrastructure-predecessor`), blocked_by (exact admission or earlier row key). No command, binary, trace, exit or counters. |
| `prelaunch` | stage (`assembly`, `build`, `vet`), classification (`invalid-source`, `infrastructure`), source_manifest_digest, request_digest (null before request exists), rewritten_digest (null before rewrite exists), runtime_digest, commands, diagnostic. No native exit or trace fields. |
| `native` | request_digest, rewritten_digest, binary_digest, runtime_digest, commands, native_result, observation. |

Admission itself is a closed record with state (`pass`, `invalid-source`,
`infrastructure`), request_digest (nullable), response_digest (nullable),
rewritten_digest (nullable), commands and diagnostic. Pass requires all three
non-null verified hashes; invalid/source frontend response may have no rewrite.
A failed worker launch has no response. A compile failure is invalid-source;
launch/transport/missing-tool/unsupported-race failures are infrastructure.
A missing artifact is not described as a native or semantic failure.

Each command record contains kind (`frontend`, `build`, `vet`, `native`,
`convert`), started bool, runtime_digest, argv_digest, output_digest (nullable),
exit (nullable integer), timed_out bool, truncated bool, diagnostic. A command
not started has null output/exit; a started command has the captured output's
hash even if empty, and an exit only if a process status was actually obtained.
A timeout may have an actual exit status but cannot imply successful completion.
Only scheduled/attempted commands occur; later absent commands are not fabricated.

native_result is exactly classification (`pass`, `assertion-fail`,
`invalid-native`, `infrastructure`), diagnostic and raw_exit (nullable only when
no exit status was available). It uses process/native finality, not witness or
Cause truth. Behavior teacher assertions and learner assertions are distinguished
by the row assembly. observation is null if no complete valid observation exists;
otherwise it has trace_digest, native_digest, report_digest (null for behavior),
semantic_digest, calls, minimal_candidates, extra_faults, assertions, operations,
allocated_bytes and behavior_assertions. Behavior rows use their standard
complete native teacher-test evidence, with zero replay counters and null
trace_digest/report_digest; native_digest is required. Replay rows require
trace_digest/report_digest and their complete observer evidence. For a corrupt
or incomplete trace, raw captured evidence remains in command/output and
sidecar manifests; absent validated observation prevents credit. Do not discard
an actual late-invalid native_result because an earlier trace prefix looked good.

### Deterministic aggregation

Aggregate only after all expected rows are present and identity/freshness is
valid. `finality: pass` means every executed native result is structurally
complete and all skipped rows have valid dependencies; it does **not** mean
assertion success. A completed assertion-fail has valid finality but cannot
pass the check. Invalid/infrastructure dominates every earlier semantic finding,
sets finality invalid, classification invalid-source/invalid-native/infrastructure
as appropriate, and prevents wrong-submission detection.

Otherwise evaluate in this priority order, retaining all available diagnostics:

1. behavior: teacher assertion failures or independent fixed-replay public
   value/remainder/input-preservation assertions. A wrong trusted original or
   correct helper is invalid infrastructure instead. Behavior-dependency skips
   make assertions/witness/cause not-run. A complete divergent fixed replay is
   behavior semantic-fail; no candidate from it is credited.
2. assertions: every learner replay has zero executed Error/Fatal assertions.
   If false, classification is test-assertion (provided behavior passed).
3. witness: after behavior and assertions pass, each repetition has an original
   minimal event paired with the actual same test/site/occurrence/bytes in fixed
   and every correct variant, all with successful finality. One complete row
   cannot prove this cross-assembly predicate. Missing live/minimal pairing is
   witness semantic-fail; cause is then not-run.
4. cause: after paired witness success, the actual report from each replay
   equals one paired minimal witness's bytes/stage/offset/extent/available and
   agrees across assemblies/repetitions. Only this gate produces Cause success
   and report_digest-backed diagnosis. Wrong facts are cause semantic-fail.

If replay assertion failure precludes paired-witness acceptance, witness/cause
are not-run; raw candidates remain diagnostic only. In every other dependency
failure the unavailable later obligations are not-run with the corresponding
higher-priority reason. Any nonexecuted obligation is never pass. Overall
classification is the first failure above or pass; successful reference/learner
checks require admission, behavior, assertions, witness, cause and finality
all pass with no skipped row. A retained early genuine fault cannot override
late assertion failure, invalidity or missing later phases.

Starter authoring gate requires admitted source, all public/full behavior rows
structurally complete with semantic assertion failure on every repetition and
no invalidity; replay and unanswered Cause are correctly not-run. Reference
requires both public and full checks pass. Each required wrong patch must
produce its declared highest-priority category on every full repetition with
all earlier prerequisites pass and no invalid result. Later dependent gates
may be not-run, never fabricated successes. This is detected=1 for that patch;
all-pass is survived=1; any invalid/unstable result is neither detected nor
survived and fails the quality gate. All required wrong patches must be detected
for the task report's mutant gate. A single matching failure among mixed
repetitions is insufficient. Review/promotion gates remain independent.

Add execution counters `reproduction_runs` (started check plans),
`reproduction_go_runs` (attempted Go-command launches),
`reproduction_workspaces` (actually created disposable workspaces),
`reproduction_helper_runs` (attempted frontend launches),
`reproduction_native_runs` (actually started learner/teacher native binaries),
`reproduction_expected_faults` (events from complete validated original
observations only). A failed process start increments attempted launch counters
but not native_runs; not-run rows increment none. An invalid overall original
observation contributes zero expected_faults, regardless of its prefix. A
complete assertion-fail can contain actual faults diagnostically but yields no
passing task/witness credit. Preserve existing counters' meanings; never infer
counters from planned row counts.

Bind production internal/repronative worker/observer/grammar/validator source,
schema, public contract and actual runtime into verifier freshness. Task input
closure includes all original/variant/scaffold/submitted source bytes. Both
quality-report and embedded collection schemas, AggregateTask/report consumers,
validator and coverage must carry or validate the new object without dropping
its failures. v5 is an explicit incompatible evidence version; v4 remains
immutable history. No ready counts or progress identities are changed by this
evidence migration; no counter alone confers an assessed encounter.

## Required staged proof and acceptance matrix

Before execution, review the concrete prototype and launcher with manifests,
all command arguments/environment, maximum children, row inventory, deadlines,
quotas, cleanup, evidence locations and source-checkout preservation. Nothing
here authorizes an unbounded experiment or direct source-checkout execution.

| Stage | Required controls and recorded evidence |
|---|---|
| Source / compile | Raw donor hash/license; exact adaptation and helper-site mapping; original/fixed file exclusion; same regression bytes; every intended runtime control compiles; compile failures labelled separately. |
| Admission | All listed alternate correct source forms, aliases and tuple helpers; reject directives, init/TestMain, recover, external/observer import, recursion, fake subjects/T receivers, over-budget source before native launch. Analyze dead code too. |
| Raw relation | Both minimal tags, full and spare-capacity forms, raw and adapted original each really fault; wrapped original actual lineage/frames match source-family and continues synthetically. No claim of global raw/local equivalence. |
| Positive native | Two distinct correct repairs plus reference, direct no-panic/table/error tests, valid calls and missing/header/tag/numeric errors, input preservation, extra faults at second field/later record; each twice with declared race/vet. |
| Semantic negative | Omit only minimal executed call; declare-but-never-call; only later/nonminimal/other-site cases; each Cause field/case wrong independently; blanket errors/lost values/wrong remainder/exact boundary/input mutation; completed late assertion after a true witness. |
| Invalid native | Unchanged buggy helper, unrelated late fault, wrong-site original fault, unexplained Goexit/exit, actual race/timeout, observer write/close failure and capture truncation. Privileged fault-injection fixtures are explicitly outside learner grammar and cannot become admitted-language evidence. |
| Protocol | Unknown/duplicate/missing/out-of-order events, mismatched identity/phase/worker/runtime/source/rewrite, pending call/checkpoint, forged terminal success, malformed frame sidecar, incomplete package/child/converter result; no synthetic JSON is described as a native run. |
| Bounds | Constructor/operation/trace maxima and just-over controls; exact serialized upper bound of final wire shape; both ordinary and race modes where supported. Quota mismatch returns to design. |
| Integration later | Same kernel in task-check/public/full; two-path readonly overlay guards; task-promote; exported/reloaded bytes still execute; public diagnostics and incorrect counterparts; old catalog/route/history/current/Next/workspace preservation in temporary homes. |

Preserve actual results, source/runtime digests and independent review receipts;
planned counts are not observed counts. Public and full feedback must distinguish
repair behavior, live minimal witness, Cause mismatch, unsupported syntax and
invalid runtime. Prototype approval is not production acceptance. Production
and content retain ADR 0033's separate PRs, ordered tests, one bank, shared
consumers and exact-head validate gates. Actual learner VCS operation, remaining
foundation modes/types/depth and full component/semantic acceptance remain open.
