# Benchmark finite procedure

## Fixed isolation boundary

Use the authorized finite manual procedure in docs/plans/finite-v1.md. Create a fresh canonical temporary root for EACH named specimen, with project, home, cache, modcache and tmp. Copy only the listed regular files; reject symlinks and escaping paths. The checkout is data only. Create project/go.mod with module finite.local/subject and go 1.27.0. Use installed Go 1.27 (and its adjacent gofmt), a clean PATH of that bin plus /usr/bin:/bin, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, CGO_ENABLED=0, GOFLAGS=-mod=readonly, GOWORK=off, and private HOME/GOCACHE/GOMODCACHE/TMPDIR. Write off to the private Go telemetry mode file before any command. Do not inherit credentials or user configuration.

Every command below runs in project with a 60-second limit. Retain exact argv/cwd, version, status, stdout/stderr, source snapshots and file digests. An expected failure must be the named ordinary diagnostic, never a timeout or process crash. Remove only these disposable roots after saving evidence. These finite recipes are not a new runner/profile and earn no automatic coverage/progress. Reference, errors, method and actual outputs require independent review. Actual release commit binding is deferred; no commit is invented.

## Exact files and commands

Every specimen copies files/subject.go.txt → subject.go, files/subject_test.go.txt → subject_test.go, and exactly one benchmark source → benchmark_test.go. Reference uses reference/benchmark_test.go.txt. Learners instead fill files/starter_benchmark_test.go.txt; the provided entry point itself earns no credit.

Reference: go version; go test -count=1 ./...; go test -run=^$ -bench=BenchmarkJoin -benchmem -count=3 -benchtime=20ms ./.... All commands succeed. The source checks equivalence before resetting each timed sub-benchmark. Retain all 24 measurements: four fixtures × two implementations × three runs. Report observed min/max ns/op, B/op and allocs/op without asserting a universal winner. Inputs are built before the timed loops; the package sink retains the final output.

Independent wrong specimens:

- setup-inside: install wrong/setup_inside_test.go.txt. Execute the same three commands. They may all pass; review rejects the input-copy allocation inside b.N because the stated operation excludes fixture setup. Compare the actual output and code, without a fixed allocation threshold.
- missing-separator: use the reference benchmark and replace only Builder's WriteString("/") with WriteString(""). Run go version and go test -count=1 ./...; require the useful output-equivalence assertion to fail. Do not benchmark semantically unequal output. Restore slash and rerun correctness successfully.
- single-tiny: install wrong/single_tiny_test.go.txt. Run go version; go test -count=1 ./...; go test -run=^$ -bench=BenchmarkJoin -benchmem -count=1 -benchtime=20ms ./.... Actual measurements may be valid for this one input but cannot support the required representative, repeated comparison. Reviewer rejects experimental-design credit, not execution.

## Three distinct review checklists

1. Timing: learner-authored b.N loop, setup before ResetTimer, equivalent intended work, observable result; identify exact lines. Supplied functions are not a timing skill action.
2. Allocations: report actual B/op and allocs/op across fixtures, distinguish bytes from allocation events, connect repeated concatenation and builder capacity growth to temporary/result storage and lifetime. No claim that zero allocations or fewer allocations alone guarantees speed.
3. Experimental design: learner-selected sizes and byte patterns, equivalent results before comparison, repeated outputs and variability, scoped conclusions. Empty is only a control. Record machine/toolchain as context, not acceptance thresholds.

Wrong method samples are actually executed and remain incomplete despite process success. The independent reviewer verifies the actual reference and analysis separately from correctness. These are first encounters only; independent application slots remain open.
