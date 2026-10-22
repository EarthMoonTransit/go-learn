# Golden-file finite procedure

## Fixed isolation boundary

Use the authorized finite manual procedure in docs/plans/finite-v1.md. Create a fresh canonical temporary root for EACH named specimen, with project, home, cache, modcache and tmp. Copy only the listed regular files; reject symlinks and escaping paths. The checkout is data only. Create project/go.mod with module finite.local/subject and go 1.27.0. Use installed Go 1.27 (and its adjacent gofmt), a clean PATH of that bin plus /usr/bin:/bin, GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off, CGO_ENABLED=0, GOFLAGS=-mod=readonly, GOWORK=off, and private HOME/GOCACHE/GOMODCACHE/TMPDIR. Write off to the private Go telemetry mode file before any command. Do not inherit credentials or user configuration.

Every command below runs in project with a 60-second limit. Retain exact argv/cwd, version, status, stdout/stderr, source snapshots and file digests. An expected failure must be the named ordinary diagnostic, never a timeout or process crash. Remove only these disposable roots after saving evidence. These finite recipes are not a new runner/profile and earn no automatic coverage/progress. Reference, errors, method and actual outputs require independent review. Actual release commit binding is deferred; no commit is invented.

## Exact files and specimen edits

For each specimen copy files/subject.go.txt → subject.go, files/report.golden → testdata/report.golden, and reference/subject_test.go.txt → subject_test.go. Learners author their test from files/starter_test.go.txt instead of receiving the reference test. Record SHA256 and actual bytes of the golden before and after every ordinary run. Named source substitutions below are the only permitted implementation edits.

Every fresh copy starts with go version. Reference then runs go test -count=1 ./... successfully; golden hash is unchanged.

Independent wrong implementation copies run go test -count=1 ./... and must fail with the named first-difference diagnostic, while preserving the golden:

- header: replace RECORDS v1 with RECORDS bad; mismatch line 1.
- order: replace loop `for i:=0;i<32;i++` with `for i:=31;i>=0;i--`; mismatch line 2.
- value: replace the argument `i*i` with `i*i+1`; mismatch line 2.
- newline: return strings.TrimSuffix(out.String(), "\n"); mismatch at the final empty line versus EOF.

Restore the original subject.go in each copy and rerun ordinary tests: success, same golden hash. These passing controls distinguish intended assertion errors from a broken test harness.

Silent-auto-update copy: use wrong/auto_update_test.go.txt plus the header defect. go test -count=1 ./... may pass, but its golden hash changes. That observed mutation rejects the method: it approved a regression without an explicit update/review.

## Explicit update specimen

Use a fresh reference copy. Change only RECORDS v1 → RECORDS v2 in subject.go. Run go test -count=1 ./...: mismatch line 1 and unchanged golden. Then run go test -count=1 -run=^TestReportGolden$ -args -update: success, changed golden. Retain before/after bytes and focused diff. Pause before the final run so the author can inspect actual diff against the approved header-only change. Accept only line 1 v1→v2 with all 32 rows/order/newlines unchanged; save that decision. Then go test -count=1 ./... must pass without further file changes. Independent final review inspects this actual update and decision too.

Reviewer checklist: real file read and exact large-output comparison; useful actual/want line diagnostic for four independent errors; no normal-run writes; observed rejection of the silent updater; actual explicit flag/update/diff/review/final run. Process success alone never establishes golden-review acceptance.
