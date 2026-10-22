# Finite procedure and exact inputs

Execution and actual acceptance are pending. This is one original manual specimen under docs/plans/finite-v1.md, outside tasks/ready. No new runner/profile or automatic progress is involved. Execute only after exact preexecution review.

Every specimen uses a fresh canonical disposable root with project/home/cache/modcache/tmp. Copy only the initial regular files listed below, removing their final .txt. Reject symlinks, escapes and unlisted files. The source checkout remains data only. Record exact source snapshots before and after edits. Remove only these temporary roots after retaining actual evidence.

Use the installed local Go1.27 executable, record resolved path and go version, and use a clean child environment: PATH=local-Go-bin:/usr/bin:/bin; private HOME/GOCACHE/GOMODCACHE/TMPDIR; GOENV=off,GOTOOLCHAIN=local,GOPROXY=off,GOSUMDB=off,GOWORK=off,GOFLAGS=-mod=readonly,CGO_ENABLED=0,GOTELEMETRY=off. No inherited credentials or GOOS/GOARCH overrides. Write off to the private telemetry mode file before commands. No go.work, external modules or network. Every command, including compilation, has a 60-second maximum, project cwd, and retained argv/status/stdout/stderr/duration. Expected failures require the named useful diagnostic, not a timeout or unrelated error.

Initial file mapping (every listed source has destination relative to project):

- `files/go.mod.txt` → `go.mod`
- `files/providers/fixed/fixed.go.txt` → `providers/fixed/fixed.go`
- `files/providers/mutable/mutable.go.txt` → `providers/mutable/mutable.go`
- `files/external/source.go.txt` → `external/source.go`
- `files/consumer/render.go.txt` → `consumer/render.go`
- `files/consumer/render_test.go.txt` → `consumer/render_test.go`
- `files/caller/doc.go.txt` → `caller/doc.go`
- `files/caller/accept_test.go.txt` → `caller/accept_test.go`

Correct replacement files:

- `reference/consumer/render.go.txt` → `consumer/render.go`

Each named row starts with `go version` in a fresh initial copy. All commands are project-relative. Retain predictions before the respective command and actual source replacements.

| Specimen | Exact command sequence and expected observation | Exact correction / evidence |
| --- | --- | --- |
| `interface-reference` | Before edit, `go test -count=1 ./consumer ./caller` fails ordinary actual/want assertions from the TODO starter. After the correct edit, `go test -count=1 ./...` succeeds. Then `go list -f '{{.ImportPath}} {{.Imports}}' ./consumer ./providers/fixed ./providers/mutable ./external` succeeds and each listed production import list is empty. | Retain learner-authored interface/body and explain why Text is necessary, Width and SetText are not, and providers need no consumer import. |
| `interface-overbroad` | Start with correct source except declare `type TextSource interface { Text() string; Width() int }`. Run only `go test -count=1 ./caller`: compiler failure must identify `*external.Source` and missing Width. Targeting caller avoids attributing this observation to mutable-provider tests. | Remove only `Width() int` from the consumer interface, then run the identical caller command successfully and `go test -count=1 ./...` successfully. Record before/after method sets and diagnostic interpretation. |
| `interface-concrete-coupling` | Replace consumer source with `package consumer`, import `finite.local/consumer-contract/providers/fixed`, and `func Render(source fixed.Value) string { return "[" + source.Text() + "]" }`. `go test -count=1 ./caller` must fail because `*external.Source` cannot be used as fixed.Value. | Replace that concrete parameter/import with the correct local named interface above. Repeat the caller command and all-package tests successfully. |
| `interface-any-method-control` | Replace consumer source with `package consumer` and `func Render(source any) string { return "[" + source.(interface{ Text() string }).Text() + "]" }`. `go test -count=1 ./...` succeeds. | Reviewer **rejects method**, because the public consumer parameter accepts arbitrary values and defers contract checking to runtime; named consumer contract was not authored. Apply correct source; the same tests still succeed. This distinguishes method review from compiler/test pass; it is not an automatic mutant. |

Only consumer/render.go is learner-editable. The production consumer and all three provider packages have no imports in the reference. The consumer tests import the two concrete providers; caller tests import consumer and external. The learner authors the interface; providers, tests and the stub do not supply it.

Wrong replacement map: wrong/overbroad/consumer/render.go.txt, wrong/concrete/consumer/render.go.txt and wrong/any/consumer/render.go.txt each replace consumer/render.go in their respective copies. Correct them with reference/consumer/render.go.txt. The overbroad repair removes only Width from the interface. The any control must pass behavior but fail method review; do not convert it into an automatic defect claim.

Review actual named declaration/parameter, minimal required operation, one-call behavior, exact bytes and producer independence. Another interface name or equivalent string construction is valid. Supplied provider/testing mechanisms earn no learner credit. Record both compiler interpretations and method rejection of the passing any control. This is the first small-contract repair only; independent application remains open.

All actual records belong in review/execution.json; acceptance.md remains pending until the sole reviewer inspects source, real output, useful failure interpretations and cleanup. Preserve original-manual provenance, digest and actual release binding. A described or passing-only procedure is unfinished.
