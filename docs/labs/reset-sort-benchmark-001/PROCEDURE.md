# Finite mutable-input benchmark procedure

This is one manual application of benchmarks.loops-timing, benchmarks.allocations
and benchmarks.experimental-design. It introduces no new concepts and grants no
automatic coverage or mastery. Prerequisites are the accepted join-benchmark-001
firsts, testing.basics-diagnostics, toolchain.formatting-vetting,
slices.backing-array-aliasing, slices.copy-overlap, functions.function-values and
control-flow.for. The sorting implementations are supplied and earn no learner
credit. Unlike the static-input join first, every operation here consumes and
changes its working state.

## Learner files and semantic evidence

Copy files/go.mod.txt to go.mod and files/sorts.go.txt to sorts.go. Copy the two
files/starter_* templates to benchmark_test.go and equivalence_test.go. Author only
those two test files. Empty signatures are not completion. Own the case generation,
iteration/reset/timer/observation code, equivalence assertions and interpretation.
Assertions must independently compare both operations against a sorted copy,
including nil/empty, signed values, duplicates and every selected benchmark case;
check exact values and length and preserve immutable templates. No time or
allocation threshold establishes correctness.

The reference chooses sizes 32 and 256 (eightfold), sorted and reverse unique
values, and an unordered eleven-value duplicate pattern. The small case exposes
constant costs; the larger one and contrasting orders expose sensitivity to work
shape. Neither represents all production distributions. The valid alternative
chooses 64/256, a different duplicate pattern and a freshly cloned working slice
on each iteration with timing stopped. Its setup allocation lifetime differs from
the reference's reused working buffer; both measure the stated sorting operation.

## Complete finite file bindings

Each disposable project contains exactly four regular files. Every copy installs
files/go.mod.txt as go.mod, files/sorts.go.txt as sorts.go,
reference/benchmark_test.go.txt as benchmark_test.go and
reference/equivalence_test.go.txt as equivalence_test.go, except these substitutions:

| Copy | Substitution |
| --- | --- |
| reference | none |
| clone-outside-timer | alternatives/clone-outside-timer/benchmark_test.go.txt |
| no-restore | wrong/no-restore/benchmark_test.go.txt |
| timed-clone | wrong/timed-clone/benchmark_test.go.txt |
| single-order | wrong/single-order/benchmark_test.go.txt |
| unequal-sort | wrong/unequal-sort/sorts.go.txt replaces sorts.go |

The full immutable asset list and hashes are in review/source-manifest.json;
provenance.json binds the original content digest. Starter templates are inspected,
not substituted for accepted reference evidence. No other code, dependency or
command is part of this finite procedure.

## Exact commands and observations

Use the locally installed Go 1.27.0 binary, invoking argv directly in each fresh
project directory. Every command has a 60-second maximum:

1. `go version`
2. `go test -run=^TestSortEquivalence$ -count=1 ./...`
3. `go test -run=^$ -bench=^BenchmarkSort$ -benchmem -count=3 -benchtime=50x ./...`

Reference and clone-outside-timer run all three successfully: six input families
by two implementations by three repetitions, 36 native measurements each.
No-restore also runs all three; correctness passes, then a real ordinary benchmark
assertion names the case, iteration, actual pre-state and wanted original. It
copies input only once, so later unsorted-case iterations receive sorted state.
A crash/timeout does not count. Source review also catches omitted restoration if
a learner removes that assertion.

Timed-clone runs all three successfully but is method-wrong: input clone/pre-state
checking and result validation are inside its measured interval. Review points to
those lines; extra allocations are observations, not a required numeric threshold.
Single-order runs all three with command 3 changed to `-count=1`: two measurements
of one sorted size. Commands can pass, but missing input families, second size and
repetitions make the comparison incomplete.

Unequal-sort runs only commands 1 and 2. Merge changes a value after sorting; the
ordinary exact-result assertion must name Merge, input, actual and expected values.
Do not benchmark unequal operations. Total: six fresh copies, exactly 17 commands.
No adaptive repetitions or correction in the same copy are permitted.

The native fixed 50x bounds the overhead of repeated timer transitions; source
still uses framework b.N. These short runs support local observations and ranges,
not statistical significance or a universal speed ranking.

## Isolation, review and interpretation

Apply docs/plans/finite-v1.md's authorized finite manual procedure. Before execution,
the sole reviewer approves these exact source hashes, the fixed one-off script and
argv. The checkout remains data only. Reject symlinks, path escapes and extra files;
create a fresh project/home/config/cache/modcache/tmp for each specimen. Use a clean
child environment, local Go, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off,
GOWORK=off, CGO_ENABLED=0, GOFLAGS=-mod=readonly, and private HOME/GOCACHE/GOMODCACHE/
TMPDIR. Write telemetry mode off before any command. Inherit no user configuration
or credentials. Keep exact argv/cwd/version/status/stdout/stderr, source snapshots,
hashes and cleanup. Kill only the owned command process group on timeout, retain
unexpected results, stop, and seek a new reviewed continuation; never rerun silently.

Independent acceptance inspects actual learner lines and output:

- Timing: each implementation/iteration starts with equivalent original state;
  restoration, working-buffer allocation and validation are outside timing; only
  the sort is timed; results remain observable. Reference stops timing before
  ResetTimer and each restoration, then starts immediately before the operation.
- Allocations: distinguish bytes from allocation events, reference working-buffer
  reuse from alternative clone lifetime, and Merge scratch allocated inside the
  measured operation. Untimed setup still influences GC/cache context; zero reported
  operation allocations does not mean the entire experiment allocates nothing.
- Design: independently authored equivalence, two justified sizes and three orders,
  36 actual repeated observations, min/max and restrained per-case interpretation.
  A sparse or contaminated experiment is rejected even if native commands pass.

Actual output analysis and evidence are retained in review/. Only the sole reviewer
can accept this manual evidence; expected behavior or process success alone is not
acceptance. Release binding uses a future actual commit, never an invented one.
