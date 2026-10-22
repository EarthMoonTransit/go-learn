# Parallel subtests: finite procedure

## Fixed files and learner boundary

Copy files/go.mod.txt to go.mod, files/fixture_test.go.txt to fixture_test.go, and exactly one selected test asset to parallel_test.go. These three files are the complete executable module. The learner edits only parallel_test.go and records an explanation separately. Starter files/starter_parallel_test.go.txt contains only an unfinished entry point; it is source-reviewed, not a fifth execution specimen.

The supplied fixture constructor/helper owns a mutex and map. The learner authors the table, t.Run/t.Parallel, per-case allocation placement, atomic lifecycle assertions, completion counting and parent Cleanup. Supplied helper locking and assertions earn no authored synchronization credit. Reference material remains separate from the learner submission.

## Execution boundary

Follow the authorized finite manual procedure in docs/plans/finite-v1.md. Source assets are inert text; execute only the exact commands below from fresh disposable copies after independent review of the actual source bytes and command manifest. Each named specimen gets its own canonical temporary root containing project, home, cache, modcache and tmp. Copy only listed regular files to listed destinations. Reject symlinks, nonregular files and escaping paths. Do not execute in the checkout or use a learner's real home.

Use local /Users/egor/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.27.0.darwin-arm64/bin/go. Each specimen first runs go version with a 10-second outer bound; record the actual version. Require the reviewed Go1.27.0 darwin/arm64 environment. Missing or changed tools and unsupported race infrastructure require review; never download a replacement or interpret unsupported checking as success.

Construct a clean child environment: PATH is that Go bin plus /usr/bin:/bin; GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, GOFLAGS=-mod=readonly, GOWORK=off, and private HOME/GOCACHE/GOMODCACHE/TMPDIR. CGO_ENABLED=0 for ordinary commands and 1 for listed -race commands. Inherit no credentials, CC, CXX, CGO flags or GORACE options. Before the first Go command, write off followed by LF into the private home/Library/Application Support/go/telemetry/mode. GOTELEMETRY alone is not a mode switch.

Every listed test command has an outer 60-second bound and -timeout=20s. All cwd values are the disposable project directory; go denotes the absolute local executable and arguments are passed as literal argv, not shell code. Kill only that command's own process group on timeout and retain the failure. Timeout, deadlock, compile failure or process crash is never the intended evidence. No retries, stress loops, added sleeps, alternate commands or arbitrary-command interface are authorized by this recipe.

Save full source snapshots, per-file hashes, provenance digest, exact argv/cwd/environment, version, exit status, stdout/stderr and timing bounds before deleting the owned temporary roots. Record actual interpretations against the saved source lines. This is finite manual evidence with no runner/profile, automatic mutant coverage or learner progress. Independent actual source/outcome/method acceptance is required; unexecuted or unreviewed work remains unfinished. Original provenance binds a digest now and an actual commit on release.

## Exact specimens and commands

Each of four fresh specimens first runs go version. Test commands are:

| Specimen: selected parallel_test.go asset | Exact commands after version | Required observation |
| --- | --- | --- |
| reference: reference/parallel_test.go.txt | go test -v -parallel=2 -count=1 -timeout=20s -run=^TestParallelTable$ .; then go test -v -parallel=1 -count=1 -timeout=20s -run=^TestParallelTable$ . | Both exit0. Three native PAUSE/CONT pairs; parent body end precedes resumed child markers; parent Cleanup observes completed3. No sibling ordering expectation. |
| parent-read: wrong/parent_read_parallel_test.go.txt | go test -v -parallel=2 -count=1 -timeout=20s -run=^TestParallelTable$ . | Exit1; parent body completed=0; want3 before returning. Parent still returns; every child and Cleanup completes. |
| shared-fixture: wrong/shared_fixture_parallel_test.go.txt | go test -v -parallel=2 -count=1 -timeout=20s -run=^TestParallelTable$ . | Exit1; fixture size2 and size3, each wanting1. Case-to-size assignment is unspecified. The lock covers insertion and assertion; no deliberate data race. |
| no-parallel: wrong/no_parallel_test.go.txt | go test -v -parallel=2 -count=1 -timeout=20s -run=^TestParallelTable$ . | Exit1; child resumed before parent body end; parent body completed=3; want0. No native PAUSE/CONT events. |

Semicolons above separate independently bounded commands; they are not a shell program. There are four version commands and five test commands. The reference's second command checks the absence of a sibling barrier at parallelism1; it is not a performance measurement. Every wrong specimen is a complete alternative source file and is run in its own fresh root.

## Method and feedback review

| Requirement | Actual source/evidence to inspect | Finite wrong observation |
| --- | --- | --- |
| Actual parallel subtests | Learner t.Parallel before any lock; native PAUSE/CONT for alpha/beta/gamma | no-parallel: resumed-before-parent diagnostic and completed3/expected0 |
| Private mutable fixtures | A new fixture per child/case, not just a new pointer variable | shared-fixture: protected map sizes2/3 instead of1 |
| Parent does not read completed child results in its body | Parent atomic count0 before return; end marker before children resume | parent-read: actual0 versus incorrect expected3 |
| Cleanup observes child completion | Deferred completion exactly once and parent Cleanup actual3 | Check actual lifecycle markers; no parent wait is added |
| Correct scheduling interpretation | Predict partial order and compare actual output | No sibling total order, elapsed-time or concurrency-speed claim |

Reviewer checks actual code and saved outcomes separately. The sample contains no mutex held across t.Parallel, parent-body wait or all-child barrier. Children may proceed one at a time at parallelism1. Count markers use atomics, so even wrong scheduling assumptions produce ordinary assertions without a race. These facts are not inferred merely from process success. This is a first encounter; independent application remains open.
