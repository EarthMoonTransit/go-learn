# Finite native reproduction procedure

Prerequisites: accepted native modules/build, ordinary test assertions, existing
fuzz target/seed concepts. This is the first fuzz reproduction encounter, not a
new automatic profile. The target and original minimal corpus are supplied.

All source assets below are inert text. A fresh canonical disposable project
receives the five `files/` assets with only final `.txt` removed. Reject symlinks,
escapes/nonregular/unlisted files. Permitted learner edits are split.go and a new
regression_test.go. The tests, oracle, go.mod and corpus remain unchanged. Source
assets under reference/ and wrong/ are complete alternative files, not an engine.

Use local Go1.27, GOTOOLCHAIN=local, GOENV=off, GOPROXY=off, GOSUMDB=off,
GOWORK=off, GOFLAGS=-mod=readonly, CGO_ENABLED=0 and fresh private HOME/config,
GOCACHE/GOMODCACHE/TMPDIR. Write off to the private telemetry mode file before
any Go command. Child environment is constructed, not inherited. Each command
has a60-second outer bound; go test has -timeout=20s. No source-checkout execution.
Retain exact argv/cwd/environment/version/source snapshots/exit/stdout/stderr;
remove each disposable root after its final evidence snapshot.

Exact common argv prefix is the absolute installed Go executable. Every copy
first runs `go version`. All other commands execute at the disposable project:

| Copy and state | Literal argv after go | Required observation |
| --- | --- | --- |
| reference, supplied broken | test -run ^FuzzSplitPair/double-colon$ -count=1 -timeout=20s | Assertion failure for :: with separator count2, actual accepted pair, expected empty/empty/false |
| same, install reference/regression_test.go.txt before repair | test -run ^TestSplitPairRegression$ -count=1 -timeout=20s | Useful retained-regression assertion failure |
| same, install reference/split.go.txt | test -run ^TestSplitPairRegression$ -count=1 -timeout=20s | Success on identical regression |
| same | test ./... -count=1 -timeout=20s | Baseline, retained Test and native seeds pass |
| same | test -run ^$ -fuzz ^FuzzSplitPair$ -fuzztime=1x -parallel=1 -timeout=20s | Bounded native invocation passes; retain actual native output |
| weakened, broken plus wrong/weakened_regression_test.go.txt | test -run ^TestSplitPairRegression$ -count=1 -timeout=20s | Pass is insufficient: source only logs, no assertion |
| same unchanged broken implementation | test -run ^FuzzSplitPair/double-colon$ -count=1 -timeout=20s | Unchanged oracle still rejects actual defect |
| same, replace implementation and regression with references | test ./... -count=1 -timeout=20s | Repaired implementation and meaningful retained assertion pass |
| omitted, correct implementation with no regression_test.go | test ./... -count=1 -timeout=20s | Behavioral pass; method rejected because ordinary regression absent |
| same, add reference/regression_test.go.txt | test -run ^TestSplitPairRegression$ -count=1 -timeout=20s | Retained regression present and passes |

There are three fresh copies, three version commands and ten test commands.
No campaign discovers or minimizes a failure: two colons are a supplied original
minimal input for this specified bug, because fewer than two cannot violate the
multiple-colon condition. A clean active invocation is evidence for its actual
workload only. A weak Test beside a correct implementation is method rejection,
not promised automatic detection. Compiler errors, timeouts and unrelated failures
do not establish the expected ordinary assertion. Exact source/commands need
independent approval before execution; actual retained results need review after.

Checklist: learner's ordinary Test fails before fix and passes after; independent
oracle/corpus preserved; root cause and bytes/minimality explained; all declared
other inputs preserved; no deleted expectation or log-only replacement accepted;
actual native output and scope stated. Manual evidence adds no automatic progress,
coverage, fuzz-target authorship or claim of learner mastery.
