# Formatting/vet finite procedure

## Fixed isolation boundary

Use the authorized finite manual procedure in docs/plans/finite-v1.md. Create a fresh canonical temporary root for EACH named specimen, with project, home, cache, modcache and tmp. Copy only the listed regular files; reject symlinks and escaping paths. The checkout is data only. Create project/go.mod with module finite.local/subject and go 1.27.0. Use installed Go 1.27 (and its adjacent gofmt), a clean PATH of that bin plus /usr/bin:/bin, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, CGO_ENABLED=0, GOFLAGS=-mod=readonly, GOWORK=off, and private HOME/GOCACHE/GOMODCACHE/TMPDIR. Write off to the private Go telemetry mode file before any command. Do not inherit credentials or user configuration.

Every command below runs in project with a 60-second limit. Retain exact argv/cwd, version, status, stdout/stderr, source snapshots and file digests. An expected failure must be the named ordinary diagnostic, never a timeout or process crash. Remove only these disposable roots after saving evidence. These finite recipes are not a new runner/profile and earn no automatic coverage/progress. Reference, errors, method and actual outputs require independent review. Actual release commit binding is deferred; no commit is invented.

## Exact file and command manifest

For every specimen copy files/subject.go.txt → subject.go and files/subject_test.go.txt → subject_test.go. The only semantic edits replace the Label format string exactly as listed. Record subject.go before and after each edit/format operation.

Reference sequence (separate fresh copy):

1. go version — zero.
2. gofmt -d subject.go — on pinned Go 1.27, exit 1 with nonempty formatting diff and empty stderr; this status reports a difference, not a parse failure.
3. gofmt -w subject.go — zero.
4. gofmt -d subject.go — zero, empty diff; bytes unchanged on repeat.
5. go vet ./... — nonzero, Sprintf %d/string diagnostic.
6. Edit only `item:%d` → `item:%s`; retain the semantic diff.
7. gofmt -w subject.go; gofmt -d subject.go — both zero, final diff empty.
8. go vet ./...; go test -count=1 ./... — both zero.

Wrong-format specimen: fresh original files; go version; gofmt -w subject.go; go vet ./... must fail with the same targeted format diagnostic. Correct the verb to %s; go vet ./... and go test -count=1 ./... must pass. A format pass alone did not repair the type mismatch.

Clean-vet/wrong-result specimen: fresh original files, replace `item:%d` with `label:%s`. Run go version, gofmt -w subject.go, go vet ./... (zero), then go test -count=1 ./... (nonzero, actual label: versus expected item:). Correct prefix to item:; vet and tests pass. This is ordinary test failure, not a compiler/vet failure.

Reviewer checklist: actual formatting diff and idempotence; correct interpretation of the printf finding; narrow semantic repair; clean-vet counterexample with useful actual/want diagnostic; private offline execution receipts and source snapshots. No supplied code mechanism receives independent implementation credit.
