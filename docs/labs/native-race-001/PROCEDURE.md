# Native race interpretation: finite procedure

## Fixed files and learner boundary

Copy files/go.mod.txt to go.mod, files/exercise.go.txt to exercise.go, files/exercise_test.go.txt to exercise_test.go, and exactly one selected cell asset to cell.go. These four files are the complete executable module. Learners edit only cell.go: preserve the value int field and load(fast bool) int / store(value int, fast bool) signatures; add a per-cell mutex and synchronize both accessors. Record actual report interpretation separately. The reference removes the redundant bypass; both supplied tests still call the same harness with their original distinct flag values.

The harness registers two workers before launching, releases them with one common start-channel close, performs one load and one store, joins both and reads the final field after the join. The sample channel is buffered1. The observed value may be0 or42; final must be42. No choice between permitted reader/writer schedules is an assertion. The learner authors the access synchronization and report explanation. Supplied worker launch, WaitGroup, tests and access signatures earn no learner authorship credit.

## Execution boundary

Follow the authorized finite manual procedure in docs/plans/finite-v1.md. Source assets are inert text; execute only the exact commands below from fresh disposable copies after independent review of the actual source bytes and command manifest. Each named specimen gets its own canonical temporary root containing project, home, cache, modcache and tmp. Copy only listed regular files to listed destinations. Reject symlinks, nonregular files and escaping paths. Do not execute in the checkout or use a learner's real home.

Use local /Users/egor/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.27.0.darwin-arm64/bin/go. Each specimen first runs go version with a 10-second outer bound; record the actual version. Require the reviewed Go1.27.0 darwin/arm64 environment. Missing or changed tools and unsupported race infrastructure require review; never download a replacement or interpret unsupported checking as success.

Construct a clean child environment: PATH is that Go bin plus /usr/bin:/bin; GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, GOFLAGS=-mod=readonly, GOWORK=off, and private HOME/GOCACHE/GOMODCACHE/TMPDIR. CGO_ENABLED=0 for ordinary commands and 1 for listed -race commands. Inherit no credentials, CC, CXX, CGO flags or GORACE options. Before the first Go command, write off followed by LF into the private home/Library/Application Support/go/telemetry/mode. GOTELEMETRY alone is not a mode switch.

Every listed test command has an outer 60-second bound and -timeout=20s. All cwd values are the disposable project directory; go denotes the absolute local executable and arguments are passed as literal argv, not shell code. Kill only that command's own process group on timeout and retain the failure. Timeout, deadlock, compile failure or process crash is never the intended evidence. No retries, stress loops, added sleeps, alternate commands or arbitrary-command interface are authorized by this recipe.

Save full source snapshots, per-file hashes, provenance digest, exact argv/cwd/environment, version, exit status, stdout/stderr and timing bounds before deleting the owned temporary roots. Record actual interpretations against the saved source lines. This is finite manual evidence with no runner/profile, automatic mutant coverage or learner progress. Independent actual source/outcome/method acceptance is required; unexecuted or unreviewed work remains unfinished. Original provenance binds a digest now and an actual commit on release.

## Exact specimens and commands

Each of four fresh specimens first runs go version. The fixed literal argv forms are:

- R-main: go test -race -v -count=1 -timeout=20s -run=^TestMainPath$ .
- R-alt: go test -race -v -count=1 -timeout=20s -run=^TestAlternatePath$ .
- R-all: go test -race -v -count=1 -timeout=20s '-run=^Test(Main|Alternate)Path$' .
- N-all: go test -v -count=1 -timeout=20s '-run=^Test(Main|Alternate)Path$' .

The quotes show shell-safe grouping only; literal argv contains no quote characters. Parentheses and the vertical bar belong to the regex argument and must never form a shell pipeline.

| Specimen: selected cell.go asset | Commands after version | Required actual observation |
| --- | --- | --- |
| unsafe-main: files/starter_cell.go.txt | R-main | Exit1 and WARNING: DATA RACE; read/write sites in cell.go and both goroutine origins in exercise.go. Value assertions are not the failure cause. |
| write-only-lock: wrong/write_only_lock_cell.go.txt | R-main | Exit1 and actual race between the bare load and protected store. Locking only the writer supplies no common synchronization. |
| main-only-fixed: wrong/main_only_fixed_cell.go.txt | R-main, then R-alt | Main exits0 clean; alternate exits1 with a race from the bypassed accesses. Retain and interpret both outputs together. |
| reference: reference/cell.go.txt | R-main, then R-all, then N-all | All exit0. Same main-path command repairs the original exercised conflict; both paths then pass race and ordinary value checks. |

There are four version commands and seven test commands. Each expected race requires both actual conflicting access stacks and goroutine creation stacks in retained output. Accept read-first or write-first report ordering; never invent stack ordering, line numbers or memory addresses. A clean wrong-copy race run is not expected detection: retain it as unverified and stop for review, rather than adding retries or stress. Both wrong paths terminate after finite operations. If a learner repair differs from the reviewed reference source, inspect its bytes and execute the reference command sequence in a separate fresh copy before acceptance.

## Method and feedback review

| Requirement | Actual source/evidence to inspect | Finite boundary |
| --- | --- | --- |
| Interpret native report | Actual read/write source lines and both worker creation sites, with matching source hashes | The detector output is dynamic evidence, not a value assertion or an automatic mutant score |
| Explain missing happens-before edge | Common start releases both; sample send and Wait join the caller after accesses | Neither construct orders the conflicting accesses relative to each other |
| Synchronize reads as well as writes | Both accessors share one per-cell mutex in the accepted source | write-only-lock still reports a conflict |
| Preserve exercised path and result contract | Same R-main command before/after; supplied harness/tests unchanged | Observed0/42 and final42 do not prescribe a winner |
| Understand an unexercised branch | Actual clean R-main and racing R-alt from main-only-fixed source | Clean subset says nothing dynamically about the omitted branch; source inspection already exposes it |
| Scope clean evidence honestly | Correct R-main/R-all/N-all outcomes and actual source review | No all-workloads or universal race-freedom claim; no performance claim for the fast flag |

The after-Wait final field read is safe because both workers have completed. Reference eliminates the unnecessary fast bypass but does not delete a supplied test or change the harness. An alternative ownership/atomic solution requires separate actual-source/specimen review because this fixed learner file retains value int for the harness. This first encounter does not create independent application or automatic coverage credit.
