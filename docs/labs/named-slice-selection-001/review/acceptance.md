# Actual author observations — independent acceptance pending

The exact approved recipe ran once from 2026-10-23T05:53:55.090459+00:00 to 2026-10-23T05:54:27.732667+00:00. All 14 fresh copies produced all 28 expected observations. Every command stayed within its 60-second bound. All source snapshots remained identical, all disposable roots were removed, and no timeout or panic was used as evidence.

Execution receipt SHA-256: `97d1d6beae64a6cb57a601c302fa7612f5ded0318d3271ed7a52ce9003536d1d`. Full argv, isolated environment, actual stdout/stderr/status and before/after source snapshots are retained in execution.json. Source digest `05066c3c89e0aece9f611dfbed2adb68162ffd3d8aade62406d67970cfef1dd4` and manifest/script remain unchanged.

| Specimen | Actual second-command exit | Observation |
| --- | --- | --- |
| reference | 0 | All four semantic/static-method client tests pass. |
| embedded | 0 | All four semantic/static-method client tests pass. |
| exact-container | 1 | ./positive_test.go:20:23: Numbers does not satisfy []Number (possibly missing ~ for []Number in []Number); ./positive_test.go:26:25: Words does not satisfy []Word (possibly missing ~ for []Word in []Word) |
| exact-elements | 1 | ./positive_test.go:20:23: Number does not satisfy sliceops.Key (possibly missing ~ for int in sliceops.Key); ./positive_test.go:26:25: Word does not satisfy sliceops.Key (possibly missing ~ for string in sliceops.Key) |
| erased-container | 1 | ./positive_test.go:21:9: got.Kind undefined (type []Number has no field or method Kind); ./positive_test.go:22:79: got.Kind undefined (type []Number has no field or method Kind) |
| any-element | 1 | ./api.go:11:6: invalid operation: value < limit (type parameter E cannot use operator <) |
| alias-input | 1 |     semantics_test.go:18: input changed during selection: got=[-2 -2 0 1 2 2] want=[3 -2 -2 0 1 2] |
| inclusive-cutoff | 1 |     semantics_test.go:15: integer selection=[3 -2 -2 0 1 2]; want strict ordered values [-2 -2 0 1 2];     semantics_test.go:32: string selection=[b a b  a]; want strict ordered values [a  a] |
| reject-float | 1 | ./probe_test.go:6:20: float64 does not satisfy sliceops.Key (float64 missing in ~int \| ~string) |
| reject-bool | 1 | ./probe_test.go:6:20: bool does not satisfy sliceops.Key (bool missing in ~int \| ~string) |
| reject-array | 1 | ./probe_test.go:6:20: in call to sliceops.Below, S (type [2]int) does not satisfy ~[]E |
| reject-mismatch | 1 | ./probe_test.go:6:21: Numbers does not satisfy ~[]string (Numbers missing in ~[]string) |
| reject-constraint-value | 1 | ./probe_test.go:6:12: cannot use type sliceops.Key outside a type constraint: interface contains type constraints |
| overbroad-float | 0 | Compiles successfully with no tests selected; the float64 client is incorrectly admitted. |

The compiler rejected the intended element/container constraints and the non-basic Key value. The missing-approximation variants are distinct: exact elements exclude named Number/Word, whereas the exact slice excludes named Numbers/Words. The erased-result variant is rejected at real method calls on the inferred result, so assignment permissiveness cannot hide loss of the named container. The unconstrained element variant fails because its generic body cannot justify `<`.

Both semantic defects compiled and failed ordinary public assertions. alias-input changes the caller slice; inclusive-cutoff includes values equal to the limit. The valid embedded-constraint alternative passes the same tests with a different output-building loop. The overbroad-float control succeeds natively but violates the declared Key type set: a successful command is deliberately insufficient for acceptance.

The learner must still explain the related S/E relation, preserved result S, exact underlying int|string set, permitted operations and why Key is constraint-only while the basic method interface is a valid value. Finite examples do not establish the universal type set; source/method review is required. Supplied tests and probe bodies are assessment evidence, not learner-authored work. No automatic coverage, mastery or unrelated application is credited.

Original provenance has project identity, creation record and external sources []; release commit binding remains the parent integration responsibility. The preexecution review is copied verbatim to preexecution-review.json. Actual independent acceptance remains separate.

## Independent actual acceptance

Sole reviewer accepted actual source/method/native evidence2026-10-23:20immutable assets,14freshcopies28commands and28snapshots verified. Nine compilerrejections andtwoordinarysemantic failures match required contracts; overbroadcompile-success rejected by method. Exactlytwo manual applications; finalbatchdelivery separate.
