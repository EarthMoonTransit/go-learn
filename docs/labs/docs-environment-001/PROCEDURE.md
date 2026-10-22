# Exact docs/environment procedure

Use the owner-authorized finite procedure in docs/plans/finite-v1.md. Each of the FOUR specimens below gets a fresh canonical disposable root with project/home/cache/modcache/tmp; copy only regular listed files, reject symlinks and escapes. Source checkout is data only. Remove only these roots after retaining actual evidence. No new runner/profile or automatic coverage.

For each copy files/parse/segment.go.txt→parse/segment.go and all four files/cmd/*.go.txt→cmd/*.go. Write go.mod exactly `module finite.local/docs` plus `go 1.27.0`. Record files before/after edits. Only edit Render: the reference replaces `_ = parse.Segment; return "TODO"` with `part, ok := parse.Segment(text, n); return fmt.Sprintf("%s:%t", part, ok)`. Whitespace formatting is allowed. No changes to parser/tests/platform markers.

Use installed local Go1.27, clean PATH of its bin plus /usr/bin:/bin; private HOME/GOCACHE/GOMODCACHE/TMPDIR, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, CGO_ENABLED=0, GOFLAGS=-mod=readonly, GOWORK=off. Write off to private telemetry mode before commands. Reference explicitly sets GOOS=darwin,GOARCH=arm64 after verifying host is darwin/arm64. Foreign commands explicitly set GOOS=linux,GOARCH=arm64; never execute their output. No inherited credentials/configuration. Every listed command has60s timeout, project cwd, actual argv/environment overrides/status/stdout/stderr/version captured. Expected failures must have the named diagnostic, never timeout/crash.

Reference commands, starting with the unchanged caller stub and go version:

1. `go doc ./parse Segment` (zero): read zero-based indexing, present empty segments, absent negative/out-of-range, exact content. Save an interpretation of these actual results, THEN apply the reference Render edit above and record the resulting source.
2. `go env GOOS GOARCH GOMOD GOWORK` (zero): darwin,arm64,canonical disposable project/go.mod,off.
3. BEFORE execution save prediction: cmd GoFiles main.go+platform_other.go; platform_linux.go excluded by filename, !linux tag admits other. Then `go list -f '{{.GoFiles}}' ./cmd` (zero, predicted pair; tests are not GoFiles).
4. `go test -count=1 ./...` (zero).
5. `go build -o bin/segment ./cmd` (zero); run absolute disposable bin/segment (zero, `non-linux|α:true|:true|:false` plus newline).
6. Set only GOOS=linux; `go env GOOS GOARCH GOMOD GOWORK` (zero,linux/arm64,same module/off).
7. BEFORE execution save prediction: main.go+platform_linux.go; !linux excludes other. `go list -f '{{.GoFiles}}' ./cmd` (zero, predicted pair).
8. `go build -o bin/segment-linux ./cmd` (zero). Inspect first20 bytes using host file read, not execution: ELF64 little-endian machine AArch64(183). Native header is Mach-O64 little-endian. Retain header bytes and interpretation.

Wrong-symbol copy: reference Render; go version; `go doc ./parse Segments` must fail with no-symbol diagnostic. Correct only symbol to Segment; same command passes and documented facts are interpreted. No test/build needed for this documentation lookup defect.

Wrong-target copy: reference Render; go version with GOOS=linux; `go env GOOS GOARCH GOMOD GOWORK` exposes linux/arm64 when native darwin/arm64 is intended. Record the incorrect native-file expectation before `go list -f '{{.GoFiles}}' ./cmd`; actual main.go+platform_linux.go disproves it. Correct GOOS=darwin; env verifies repair. Record corrected prediction before the same go list; build `go build -o bin/segment ./cmd`, then run only this repaired native executable with expected reference output. No foreign executable is run.

Wrong-caller copy: go version; Render uses `parse.Segment(text,n+1)` and the same formatting. `go doc ./parse Segment` succeeds; `go test -count=1 ./...` must fail with named Render actual/want assertions, proving readable docs alone do not imply correct use. Repair only n+1→n; repeat tests, build native bin/segment, run it; all pass.

Finite review checklist: actual symbol discovery and facts applied in caller; actual environment interpreted; timestamped predictions precede go list; filename/tag selections explained for both targets; native execution and foreign header-only evidence; all three realistic mistakes produce useful evidence and actual repairs; private offline cleanup and provenance verified. Supplied parser/testing mechanisms earn no implementation/testing credit.
