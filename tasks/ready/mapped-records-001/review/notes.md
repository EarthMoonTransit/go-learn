# P18 author notes

Source decision: PR83, merge f458a1aad95626d74c901749fa1664862db0e036;
`catalog/intake/2026-10-17/p18-source-review.json`. Only Birsi Task9.2 is an
API/operation donor. The source has no completed body; no upstream cases,
solution implementations, adjacent First/Max/Keys sections or other restricted
sources were copied. MIT notice, exact commit/path/artifact hashes are retained.

This is substantive mapping-operation adaptation with an explicit typed-record
extension, not another Reduce accumulator repair. The private function declaration
is absent from the starter; its erased non-generic stub keeps public/full packages
compiling. Entry is immutable supplied support explained in TASK, not an assessed
generic data structure. Existing structs/slices/functions/named-type prerequisites
occur earlier than the planned insertion after Reduce. No new theory or learner
state access. Time estimate is provisional, not learner-calibrated.

The exported Entry API is identical in starter/reference/mutants. A private helper
is intentionally the assessed API. Same-package workspace tests are supported by
the current contract. go/types checks real declaration/type identities, unrestricted
constraints and a synthesized explicit/inferred call matrix. Explicit generic calls
never enter actual test Go files. Constraint/type assertions are declaration evidence;
normal compiler failures/timeouts are never accepted as behavioral detection.

## Public requirement feedback

| Requirement | Public assertion | Compiling wrong variants | Expected useful diagnostic |
| --- | --- | --- | --- |
| Two unrestricted independent parameters; exact typed API and calls | TestGenericDeclarationAndCallMatrix | no-type-parameters, comparable-input, restricted-output, erased-input-type, erased-output-type | actual parameter count/constraint/type and wanted two/any/typed result |
| Preserve each input value and mapped output at its index | TestMappingValuesAndCallOrder, TestIndependentStructAndStringTypes | zero-input, zero-output, reverse-result, skip-last | actual and expected typed entries |
| Exactly once, forward callback order | TestMappingValuesAndCallOrder | double-call, reverse-calls | actual callback inputs or stateful mapped values versus expected order/count |
| Empty input has zero entries and zero calls | TestEmptyInputsDoNotCallMapper | empty-call | actual callback calls, want 0; nil and allocated empty both tested |
| Input elements are not overwritten | TestInputSliceIsNotOverwritten | overwrite-input | actual input, want unchanged snapshot |
| No goroutine launches | TestNoGoroutineLaunch | goroutine-call | actual go-statement location, want synchronous mapping |
| No interface-value erasure or reflection | TestNoValueErasureOrReflection | boxed-values, reflection | actual interface value type/location or reflect import, want direct typed mapping |

All required variants must compile with both canonical suites. Runtime A cases
are comparable, so the intentionally too-narrow input constraint remains compiling
and is rejected by the explicit unrestricted-constraint assertion. The synthesized
matrix also admits a named non-comparable struct input. Runtime B covers string,
[]byte and named []byte; the text-only union similarly compiles but fails the
unrestricted assertion. A blanket comparable output or one-parameter declaration
that fails native compilation is not a required mutant and earns no behavioral kill.

The goroutine-call variant waits after each launch, preserving values and callback
order without a race. Its public failure therefore proves the explicit no-go
requirement, rather than relying on a timeout or incidental behavior.

Grader adds singleton/repeated/zero values, array inputs, named byte outputs,
independent calls and nil-versus-empty mapped values. Empty outer result remains
representation-tolerant. Mapping callbacks are supplied and do not mutate inputs;
no deep-copy, allocation-count or output ownership promise is added.

## Executed feedback evidence

An isolated catalog copy used the exact authored files and current contracts.
Only that fixture received explicitly provisional static-review metadata to permit
existing learner verification before canonical final approval. Official
`task-promote` performed its ready transition; no canonical gate was bypassed.
Its technical report passes: starter public/full compile then fail, repeated
reference/full and vet pass, all 16 required mutants compile and are detected.

A single disposable CheckSession then ran public-only verification for the
starter and every required mutant: 17 assertion failures, no compiler errors or
timeout-based detections. Reference passes public and full verification. Actual
public diagnostics include `type parameters = 0, want 2`,
`constraint = comparable, want unrestricted any`, exact typed entries and callback
counts versus expected values, and `goroutine launch at implementation.go:7:3;
want synchronous mapping without go statements`. The corresponding public tests
and wrong variants are listed above. Local evidence is retained under
`.tmp/p18-content/{review-check.yaml,public-proof.json,public-proof.log}`;
canonical ready/full-bank acceptance remains a separate final gate.

Routing appends stable address 07.22 and executes after generic-reduce-001,
before named-numeric-constraint-001. All 166 existing route tuples (block, address,
role, skills) and mapping rows except proposed execution order remain identical;
the original CSV is byte-identical. No real learner home was accessed. No fixture
completion is human learning evidence.
