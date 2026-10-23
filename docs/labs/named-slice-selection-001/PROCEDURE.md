# Finite related-parameter API procedure

This manual unit assesses independent applications of constraints-type-sets and
approximation-unions after named-numeric-constraint-001. It introduces no new
concept or automatic coverage. That scalar first supplied Double; this learner
owns Key, the S/E relationship, preserved named container type and the whole
selection body. Supplied clients/tests are assessment, not learner authorship.

## Learner work and meaning

Install files/starter_api.go.txt as api.go. Author only that file. Its unconstrained
identity TODO is incomplete: it supplies no working type set or container relation.
Below's ordered type parameters are S then E, input is S, limit is E and result is
S. Key denotes exactly types whose underlying type is int or string. S's slice
relation must use the same E. Use comparisons and slice construction justified by
those constraints. Do not erase to a fixed concrete slice or element type.

Reference uses S ~[]E and E Key, with Key's ~int | ~string union. The embedded
alternative decomposes the same union and names a parameterized slice constraint;
it fills an indexed result and returns its selected prefix. Both preserve S and
its methods. Empty-nonnil capacity/nilness and allocation counts are unspecified.
Nil input stays nil; nonempty selected storage is independent of input. Values
are strictly below the limit, preserving order and duplicates.

The positive client directly calls Kind on the inferred named result and places
it in a basic method-only interface. Assignment to a named slice alone would not
prove the static type because Go assignment rules may allow an unnamed slice.
The negative Key value client separately demonstrates a non-basic constraint
cannot be an ordinary value interface. Finite type probes do not enumerate all Go
types; independent source/type-set reasoning establishes the complete boundary.

## Fixed project files

Each specimen copies files/go.mod.txt to go.mod, files/semantics_test.go.txt to
semantics_test.go, files/positive_test.go.txt to positive_test.go and exactly one
API asset to api.go. Only named probe specimens add one probes/*_test.go.txt as
probe_test.go. All assets are inert text in this checkout. Reject symlinks, extra
files and path escapes. The source manifest pins the complete lab asset set,
provenance and the fixed execution recipe before execution.

## Exact finite commands

Use pinned local Go 1.27.0 with direct argv in each fresh project directory.
Every command has a 60-second maximum:

- V: `go version`
- T: `go test -v -count=1 ./...`
- C: `go test -run=^$ -count=1 ./...`

C compiles all files and runs no tests. Probes put their operations in an uncalled
function, so even the intentionally overbroad accepted client executes no probe
body. Invalid programs never execute. Fourteen fresh copies run two commands each:

| Specimen | API asset | Probe | Commands | Expected observation |
| --- | --- | --- | --- | --- |
| reference | reference/api.go.txt | none | V,T | values, storage, nil and actual named-method clients pass |
| embedded | alternatives/embedded/api.go.txt | none | V,T | equivalent constraints and different body pass |
| exact-container | wrong/exact-container/api.go.txt | none | V,T | defined slice fails S admission; compiler identifies slice relation |
| exact-elements | wrong/exact-elements/api.go.txt | none | V,T | defined Number/Word fails Key admission |
| erased-container | wrong/erased-container/api.go.txt | none | V,T | actual result.Kind is undefined on unnamed []E |
| any-element | wrong/any-element/api.go.txt | none | V,T | generic E does not support ordered comparison |
| alias-input | wrong/alias-input/api.go.txt | none | V,T | ordinary input-preservation/storage assertion fails |
| inclusive-cutoff | wrong/inclusive-cutoff/api.go.txt | none | V,T | ordinary strict-selection assertion fails |
| reject-float | reference/api.go.txt | float | V,C | float64 violates element constraint |
| reject-bool | reference/api.go.txt | bool | V,C | bool violates element constraint |
| reject-array | reference/api.go.txt | array | V,C | array violates slice relation |
| reject-mismatch | reference/api.go.txt | mismatch | V,C | named int slice cannot satisfy slice-of-string relation |
| reject-constraint-value | reference/api.go.txt | constraint-value | V,C | Key cannot be used outside a constraint |
| overbroad-float | wrong/overbroad-float/api.go.txt | float | V,C | compile succeeds; type-contract wrong because float64 was admitted |

Native failures must name the intended type/operation. Compile failures are
expected manual type-boundary evidence only for those rows; they are not automatic
mutant detection. The two semantic defects must compile and fail useful ordinary
assertions. The overbroad variant must compile successfully: process success alone
does not prove the intended type set. A panic, timeout or unrelated failure is not
credit. No unreviewed probe extension or repeat-in-place is allowed.

## Isolation and acceptance

Reuse docs/plans/finite-v1.md's authorized finite procedure. Sole reviewer must
approve exact source/script/manifest hashes, 28 argv and expected families before
any authored execution. Each copy gets private project/home/config/cache/modcache/
tmp, a clean explicit environment and a telemetry-off mode file written before the
first command. Use local Go, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off,
GOWORK=off, CGO_ENABLED=0 and readonly module flags. Do not inherit credentials,
user configuration or network dependencies. Checkout files remain data only.

Retain version, argv/cwd/environment, timed start/end, stdout/stderr, exits and
before/after source snapshots and hashes. Remove each owned root after saving
results, including unexpected failures. On timeout kill only that command's own
process group and stop; retain failure without credit. A correction needs a newly
reviewed finite continuation. There is no general compiler runner or new profile.

Actual independent review checks Key's full union, related S/E, result S and method
witness, permitted generic operations, semantic behavior, both native rejection
and overbroad-success controls, and explanation of basic versus non-basic interfaces.
Original content keeps project identity, creation record, empty external sources
and digest; actual release commit binding is deferred. No automatic progress,
learner mastery or unrelated skill/component credit is inferred from these runs.
