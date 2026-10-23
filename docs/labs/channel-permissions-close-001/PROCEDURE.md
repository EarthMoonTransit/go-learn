# Finite channel permissions and closure procedure

This original manual unit remains outside tasks/ready. The bounded contract was
approved before authoring; exact source/argv preexecution approval and actual
acceptance remain separate. No automatic completion or coverage is claimed.

## Inputs and edits

manifest.json is the complete exact file map and command list. Initial regular
files map to go.mod, channel/report.go, channel/observe.go,
channel/observe_test.go, cmd/demo/main.go and PREDICTIONS.md. Remove only the final
.txt from the reviewed Go/module assets. Reject symlinks, path escapes and
unlisted assets. Only channel/observe.go and PREDICTIONS.md are editable.

The supplied Report/Reading types, synchronous recovery wrapper, tests and demo
formatter are observation infrastructure. No helper signature or channel action
is supplied in the starter. Author the producer and consumer parameter permissions
and actual operations. The observation tests include empty input and a queued zero;
the last two receives must both report zero, false, true. Input has at most four
values. The sufficient buffer and absence of goroutines are actual method checks.

Record predictions before every behavioral/compiler command. Completed reference
predictions are kept separately in reference/PREDICTIONS.md; they are assessment
material, not learner scaffolding. Source snapshots preserve initial, predicted,
wrong/probe, and repaired states with hashes.

## Exact commands

Use the absolute installed Go1.27 path in manifest.json. Resolve and record it;
if absent, stop without downloading or falling back. Every row begins with
`go version`, in its own fresh disposable project root. Each command is limited
to 60 seconds; retain argv/cwd, actual stdout/stderr, exit, duration and timeout.

| Specimen | Commands after go version | Replacements and expected outcome |
| --- | --- | --- |
| reference | go test -count=1 ./...; go test -count=1 ./...; go run ./cmd/demo | First starter ordinary failure, then reference replacement, tests/demo success. |
| producer-receive | go build ./...; go build ./...; go test -count=1 ./... | Install producer probe, native send-only receive rejection, restore reference, build/test success. |
| receiver-send | Same three commands | Install receiver send probe, native receive-only send rejection, restore, build/test success. |
| receiver-close | Same three commands | Install receiver close probe, native receive-only close rejection, restore, build/test success. |
| missing-close | go test -count=1 ./...; go test -count=1 ./... | Wrong replacement then reference. Ordinary Ready and panic-flag failures, then success. |
| discard-buffer | Same two commands | Owner discards receives before observation; first values mismatch, then reference success. |
| zero-as-closure | Same two commands | Queued-zero OK becomes false; ordinary indexed mismatch, then reference success. |
| suppress-panic-attempts | Same two commands | Owner attempts omitted and flags false; ordinary flag failures, then reference success. |
| unrestricted-helpers | go test -count=1 ./...; go run ./cmd/demo; go build ./...; go build ./...; go build ./... | Ordinary wide-helper control passes, then the three wide-helper probe replacements each compile. Probe bodies are NEVER run. Method rejects permissive signatures. |
| hardcoded-panic-flags | go test -count=1 ./...; go run ./cmd/demo | Behavior passes; actual method rejects flags replacing real operations. |
| correct-alternate | go test -count=1 ./...; go run ./cmd/demo | Named direction types, other helper names, larger sufficient buffer and alternate result assembly pass. |

Exactly 11 copies and 40 commands including versions. For each restricted probe,
the reference differs by the single inserted illegal statement; restoration removes
only that statement. Widening parameters is not an acceptable repair. Probe paths,
diagnostic fragments and exact success stdout are in manifest.json. Wrong runtime
specimens must fail with named ordinary assertions, not compiler errors or process
failures. Native compiler rows require the relevant direction/operation diagnostic,
not unrelated syntax errors. Expected nonzero exits never justify a timeout, crash
or race. Preserve any unexpected result and stop; do not extend/rerun the manifest
without review. Success exits on the two method-wrong controls are expected.

## Isolation and cleanup

The one-off fixed recipe accepts only --approved-execution and verifies the
independent approval's exact content digest, manifest hash and its own script hash
before creating a copy. It executes no arbitrary user commands or projects.
Use private project, home, config, cache, modcache and temp directories, clean
child environment, local Go and offline stdlib only. Set GOENV=off,
GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, GOWORK=off,
GOFLAGS=-mod=readonly, CGO_ENABLED=0, GOTELEMETRY=off. Write the private telemetry
mode file before any Go command. No inherited credentials, user's workspace,
network, toolchain download or GOOS/GOARCH override is used.

Reference and missing-close both finish without goroutines: receives are
nonblocking; owner send uses select/send/default; each owner operation has its own
recovery call. A closed-channel send case is ready and panics even with default.
An open empty channel in the missing-close wrong case instead defaults (capacity
zero) or accepts the send (positive capacity). Its first owner close does not panic.
The wrong case therefore terminates with useful observations, not a stranded worker.

Retain actual source texts/hashes and prediction events before commands; retain
actual output before removing only each owned temporary root. Raw evidence goes
to the batch .tmp receipt, normalized evidence to review/execution.json, with
independent review afterward. The original content digest excludes review evidence
and provenance.json and binds sorted relative-path LF file-SHA256 LF records.
No release commit is invented.

## Finite method acceptance

Inspect actual helper parameter declarations/calls, sufficient owned buffering,
all sends before the single initial close, actual comma-ok receives, Ready values,
real nonblocking owner send and second close inside separate recovery calls, and
predictions compared with actual native outputs. Directions limit operations but
do not prove sole ownership or change buffering. Fixed types/recovery/formatter
earn no authorship credit. Runtime results alone cannot detect hardcoded flags or
unnecessarily wide helper types; keep both passing controls method-rejected.
Other helper names, named directional types, equivalent local assembly and a larger
sufficient buffer are valid. No AST, profile, runner or automatic coverage changes.

Compared with the unchanged firsts, this unit adds learner-selected helper APIs,
native invalid-operation diagnoses/repairs and explicit closed-state observations
and safe post-close attempts. It claims only channels.directional and
channels.close-semantics applications. Root owns manual index/backlog/route and
delivery integration; source/evidence acceptance belongs to the sole reviewer.
