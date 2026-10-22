# Finite configuration project procedure

Prerequisites: modules-build, consumer-text-contract, package-boundaries, package-inventory, external-window-tests and writer-double, plus earlier slice/map/function/string operations. The learner owns five files listed in manifest.reference, including the new project/go.mod; supplied adapter, client and public tests are unchanged. The starter contains no consumer interface, module, double, assertions, formatter or composition implementation. Author external tests named TestLearnerInjectedSnapshot and TestLearnerMissingAndEmpty; their recording dependency and assertions must cover the TASK requirements.

Each row uses a fresh canonical disposable root. Source assets are inert data in this checkout: never run them here. Copy only the regular files in manifest.json, rejecting symlinks and escaped paths. Keep project/client paths inside that root. Snapshot initial and installed sources; install the listed reference edits, then the row's named wrong overlay if any. No other edits/files or retries.

Use local Go1.27.0 with private HOME/config/cache/modcache/tmp, PATH=Go-bin:/usr/bin:/bin, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, GOWORK=off, GOFLAGS=-mod=readonly, CGO_ENABLED=0 and GOTELEMETRY=off. Write off to the private telemetry mode file before commands. Inherit no credentials, build flags or user workspace. Every command, including the built executable, has a60-second outer limit; tests also use20s. Retain argv/cwd/environment, Go version, source snapshots, start/end/status/stdout/stderr. Timeout or unrelated syntax/import error earns no credit. Delete only the owned temporary root after saving actual results, including failure paths.

Exact source/command approval must precede execution. review/execution.json retains real results and review/independent-review.json records actual acceptance; a command list alone is unfinished. This manual unit creates no automatic coverage, progress or mastery. Sources are original ADR0008 material; provenance binds creation and content hashes, with actual release commit binding deferred. Method review accepts equivalent implementation/private names while enforcing the stated learner-owned actions.

Set only the explicit additional fixture environment A=ambient and EMPTY=; ABSENT is unset. Each row begins with `go version` from project. All test commands use `go test -count=1 -timeout=20s` with the arguments below. Commands run in project unless prefixed client.

| Row | Commands after version | Expected observations |
|---|---|---|
| reference | test ./...; go build -o bin/snapshot ./cmd/snapshot; ./bin/snapshot; go list -f '{{.ImportPath}} {{.Imports}}' ./snapshot ./snapshot/internal/row ./processenv ./cmd/snapshot; client: test ./... | All pass, exact stdout below, expected production graph, independent provider accepted. |
| concrete | client: test ./... | Compiler rejects client.Third as processenv.Source. |
| overbroad | client: test ./... | Compiler identifies missing Environ on client.Third. |
| bypass | test -run '^TestPublic' ./snapshot; test -run '^TestLearner' ./snapshot | Both fail useful injected-value assertions; ambient differs from requested injected value. |
| empty | test -run '^TestPublic' ./snapshot; test -run '^TestLearner' ./snapshot | Both fail present-empty versus missing assertions. |
| twice | test -run '^TestLearner' ./snapshot | Learner-authored recording assertion rejects duplicated lookup history. |
| any | test ./...; client: test ./... | Behavior passes; method review rejects the missing named compile-time consumer contract. |
| internal | client: test ./forbidden | Compiler rejects internal import outside its parent tree. |
| module | go build -o bin/snapshot ./cmd/snapshot | Incorrect module path cannot resolve intended configkit imports offline. |

Exactly9 copies/25 commands including version. Command stdout is exactly:

```text
A="ambient"
EMPTY=""
ABSENT=<missing>
A="ambient"
```

Each line includes newline. The import listing must show snapshot→snapshot/internal/row, row→fmt, processenv→os and cmd→snapshot/processenv/fmt; snapshot has no processenv/os import. Check actual source delegation, module/import decisions, minimal consumer-owned Lookup method, independent client, composition root and learner-authored deterministic double/assertions. Learner-only probes establish injection, presence and interaction diagnostics; supplied public fixtures do not establish authored tests. The any control is honestly a passing behavior/method rejection. The component and two applications (small-contracts and seams-doubles) require actual evidence. No additional module, package-boundary or toolchain applications are inferred from the component.
