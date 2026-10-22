# Finite constants contexts procedure

Prerequisites: constant-kinds, numeric conversions, fixed-width-shifts and modules-build. The only editable file is project/limits/constants.go. Positive consumers and all listed context probes are supplied assessment material; they earn no test/consumer authorship credit. Scale/Mask/Marker declarations themselves start absent.

Each row uses a fresh canonical disposable root. Source assets are inert data in this checkout: never run them here. Copy only the regular files in manifest.json, rejecting symlinks and escaped paths. Keep project/client paths inside that root. Snapshot initial and installed sources; install the listed reference edits, then the row's named wrong overlay if any. No other edits/files or retries.

Use local Go1.27.0 with private HOME/config/cache/modcache/tmp, PATH=Go-bin:/usr/bin:/bin, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, GOWORK=off, GOFLAGS=-mod=readonly, CGO_ENABLED=0 and GOTELEMETRY=off. Write off to the private telemetry mode file before commands. Inherit no credentials, build flags or user workspace. Every command, including the built executable, has a60-second outer limit; tests also use20s. Retain argv/cwd/environment, Go version, source snapshots, start/end/status/stdout/stderr. Timeout or unrelated syntax/import error earns no credit. Delete only the owned temporary root after saving actual results, including failure paths.

Exact source/command approval must precede execution. review/execution.json retains real results and review/independent-review.json records actual acceptance; a command list alone is unfinished. This manual unit creates no automatic coverage, progress or mastery. Sources are original ADR0008 material; provenance binds creation and content hashes, with actual release commit binding deferred. Method review accepts equivalent implementation/private names while enforcing the stated learner-owned actions.

Each row starts with `go version` in project. Tests mean `go test -count=1 -timeout=20s` with the shown package. Exactly9copies/20commands including versions.

| Row | Commands after version | Expected observation/prediction |
|---|---|---|
| reference | go build ./...; test ./...; go run ./cmd/values | All pass; stdout exactly255 255 2 followed by newline. |
| default | go build ./probe | Inferred default int cannot represent Scale=2^80. |
| signed | go build ./probe | Mask255 is not representable as int8. |
| overflow | go build ./probe | Mask+1=256 is not representable as uint8. |
| typed-target | go build ./probe | Typed uint8 Marker cannot be implicitly assigned to uint16. |
| wrong-mask-type | test ./accept | Incorrectly typed uint8 Mask fails unchanged positive uint16 context. |
| wrong-marker-kind | go build ./probe | Incorrectly untyped Marker permits forbidden uint16 assignment: expected command success, method/contract rejection. |
| wrong-scale-type | test ./accept | uint64 Scale declaration overflows. |
| wrong-variable | test ./accept | Variable Marker fails unchanged constant-expression use. |

Check actual source declarations, exact compile-time arithmetic and each value/type/defaulting interpretation. Marker conversion succeeds only in the explicit permitted positive context; it does not repair the deliberately invalid implicit probe. Record the wrong-marker-kind success honestly rather than relabeling it a test failure. Invalid declaration diagnostics must concern the named context, not unrelated syntax/module defects. This is the missing independent types.constants application only; no iota claim. All source snapshots and actual compiler/behavior outputs must be independently accepted.
