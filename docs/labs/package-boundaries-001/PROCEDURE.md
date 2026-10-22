# Finite procedure and exact inputs

Execution and actual acceptance are pending. This is one original manual specimen under docs/plans/finite-v1.md, outside tasks/ready. No new runner/profile or automatic progress is involved. Execute only after exact preexecution review.

Every specimen uses a fresh canonical disposable root with project/home/cache/modcache/tmp. Copy only the initial regular files listed below, removing their final .txt. Reject symlinks, escapes and unlisted files. The source checkout remains data only. Record exact source snapshots before and after edits. Remove only these temporary roots after retaining actual evidence.

Use the installed local Go1.27 executable, record resolved path and go version, and use a clean child environment: PATH=local-Go-bin:/usr/bin:/bin; private HOME/GOCACHE/GOMODCACHE/TMPDIR; GOENV=off,GOTOOLCHAIN=local,GOPROXY=off,GOSUMDB=off,GOWORK=off,GOFLAGS=-mod=readonly,CGO_ENABLED=0,GOTELEMETRY=off. No inherited credentials or GOOS/GOARCH overrides. Write off to the private telemetry mode file before commands. No go.work, external modules or network. Every command, including compilation, has a 60-second maximum, project cwd, and retained argv/status/stdout/stderr/duration. Expected failures require the named useful diagnostic, not a timeout or unrelated error.

Initial file mapping (every listed source has destination relative to project):

- `files/go.mod.txt` → `go.mod`
- `files/cycle/producer/source.go.txt` → `cycle/producer/source.go`
- `files/cycle/consumer/render.go.txt` → `cycle/consumer/render.go`
- `files/cycle/consumer/render_test.go.txt` → `cycle/consumer/render_test.go`
- `files/cmd/cycle/main.go.txt` → `cmd/cycle/main.go`
- `files/domain/internal/format/format.go.txt` → `domain/internal/format/format.go`
- `files/domain/internal/format/format_test.go.txt` → `domain/internal/format/format_test.go`
- `files/domain/facade.go.txt` → `domain/facade.go`
- `files/domain/facade_test.go.txt` → `domain/facade_test.go`
- `files/domain/probe/probe.go.txt` → `domain/probe/probe.go`
- `files/domain/probe/probe_test.go.txt` → `domain/probe/probe_test.go`
- `files/cmd/outside/main.go.txt` → `cmd/outside/main.go`

Correct replacement files:

- `reference/cycle/consumer/render.go.txt` → `cycle/consumer/render.go`
- `reference/cycle/producer/source.go.txt` → `cycle/producer/source.go`
- `reference/domain/facade.go.txt` → `domain/facade.go`
- `reference/cmd/outside/main.go.txt` → `cmd/outside/main.go`

Each named row starts with `go version` in a fresh initial copy. All commands are project-relative. Retain predictions before the respective command and actual source replacements.

| Specimen | Exact command sequence and expected observation | Permitted correction / proof |
| --- | --- | --- |
| `cycle-original` | Save predicted cycle path. `go test -count=1 ./cycle/consumer` fails with import-cycle diagnostic naming consumer and producer. | Apply only the two cycle edits above. `go test -count=1 ./cycle/consumer ./cycle/producer` succeeds; `go run ./cmd/cycle` succeeds with stdout `[alpha]\n`. Record removal of both production back-edges and preservation of renderer bytes. Unrelated internal defect is not executed in this copy. |
| `cycle-alias-only` | Keep original producer and replace only consumer with the alias-only specimen below. `go test -count=1 ./cycle/consumer` still fails with the same import cycle. | Apply the two correct cycle edits, then run identical targeted tests and cycle command successfully. Explain that a type alias does not remove an import edge. |
| `internal-outside` | Save parent-tree prediction. `go test -count=1 ./domain/probe ./domain/internal/format` succeeds, proving a legal inside caller and formatter behavior. Independently, `go run ./cmd/outside` fails with internal-package-not-allowed diagnostic naming domain/internal/format. | Apply only facade and outside-caller edits. `go test -count=1 ./domain ./domain/probe ./domain/internal/format` succeeds; the **same** `go run ./cmd/outside` succeeds with stdout `<beta>\n`. Explain why cmd/outside is outside domain despite the shared module. Cycle subtree is not executed in this copy. |
| `boundaries-reference` | Apply both pairs of correct edits to a fresh copy. `go test -count=1 ./...` succeeds. `go list -f '{{.ImportPath}} {{.Imports}}' ./cycle/consumer ./cycle/producer ./domain ./domain/probe ./cmd/cycle ./cmd/outside` succeeds and matches the repaired graph. `go run ./cmd/cycle` prints `[alpha]\n`; `go run ./cmd/outside` prints `<beta>\n`. | This is the combined correct control. It cannot substitute for either independently observed initial failure. Review actual graph, source ownership, unchanged tests/outputs and delegated facade. |

Only cycle/consumer/render.go, cycle/producer/source.go, domain/facade.go and cmd/outside/main.go are learner-editable. No module changes, added/moved files, copied formatter, build tags, init workaround or test/output edits. The finite chosen repair is a public facade, not caller relocation.

Initial graph: cmd/cycle imports consumer and producer; consumer imports producer; producer imports consumer through Record. Cmd/outside imports domain/internal/format illegally; domain/probe imports it legally from inside parent domain. Domain initially has an empty facade stub. The two failing commands target independent subgraphs, so neither failure masks the other or depends on the stub facade test.

The alias-only specimen replaces cycle/consumer/render.go with wrong/alias-only/cycle/consumer/render.go.txt while retaining the original producer. The cycle fix replaces both cycle files with reference counterparts. The internal fix replaces facade and cmd/outside with reference counterparts. The combined reference applies all four replacements.

Before the failing commands record the predicted cycle and internal parent boundary. After repair the consumer owns Source; neither cycle implementation imports the other. Domain and domain/probe may import domain/internal/format; cmd/outside instead uses domain.Format. Review actual source and delegation rather than copied output implementation. Retain separate cycle and internal compiler failures, alias-only failure, legal inside caller, repaired commands, combined tests and actual import listing. These are two first outcomes, not a further small-interface application or automatic coverage.

All actual records belong in review/execution.json; acceptance.md remains pending until the sole reviewer inspects source, real output, useful failure interpretations and cleanup. Preserve original-manual provenance, digest and actual release binding. A described or passing-only procedure is unfinished.
