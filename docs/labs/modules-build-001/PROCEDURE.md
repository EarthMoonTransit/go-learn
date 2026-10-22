# Finite execution and review procedure

Acceptance status and actual execution evidence are recorded in review/.
This is a fixed specimen under the owner's 2026-10-22 exception to AGENTS rule 9,
not a task-check profile or a reusable command runner.

## Isolation and exact inputs

Create a fresh temporary root with `project/lib`, `project/app`, `project/bin`,
`home`, `cache`, `modcache`, `tmp`. Resolve the temporary root to its canonical
absolute path before deriving cwd or environment paths (on macOS /var and
/private/var must not be mixed). Copy only these regular source files, rejecting
symlinks: `files/number.go.txt` to `project/lib/number.go`,
`files/number_test.go.txt` to `project/lib/number_test.go`,
`files/main.go.txt` to `project/app/main.go`, and `files/main_test.go.txt` to
`project/app/main_test.go`. The source checkout remains data only.

Use a clean child environment with PATH limited to the installed local Go bin,
/usr/bin and /bin; GOENV=off, GOTOOLCHAIN=local, GOPROXY=off, GOSUMDB=off,
CGO_ENABLED=0, GOFLAGS=-mod=readonly. Point HOME, GOCACHE, GOMODCACHE and TMPDIR
inside the temporary root. Before the first command, write `off` to the isolated
Go telemetry mode file (`home/Library/Application Support/go/telemetry/mode`
on macOS; `home/.config/go/telemetry/mode` on Linux). Initially GOWORK=off.
Use an installed Go 1.27 executable (the pinned repository language version),
not a PATH shim selecting an older Go, and a 60-second subprocess timeout for each
command. Do not inherit user configuration or credentials.

## Reference commands

Record argv, working directory, version, stdout/stderr and exit status for every
line. Paths below are relative to `project`; only module/workspace edits listed
here are allowed. All reference commands must exit zero.

| Working directory | Command | Observable result |
| --- | --- | --- |
| `.` | `go version` | Actual toolchain version |
| `lib` | `go mod init finite.local/temporary` | Creates library go.mod |
| `lib` | `go mod edit -module=finite.local/number` | Corrects module path; retain before/after bytes |
| `app` | `go mod init finite.local/cmd` | Creates command go.mod |
| `app` | `go mod edit -require=finite.local/number@v0.0.0` | Records dependency without fetching it |
| `.` | `go work init ./app` | Creates workspace, initially without library |

Immediately after go work init and before go work use, set GOWORK to the absolute path of **this copy's**
`project/go.work` (never the user's workspace). Keep the isolated offline
environment and GOFLAGS=-mod=readonly unchanged.

| Working directory | Command | Observable result |
| --- | --- | --- |
| `.` | `go work use ./lib` | Adds local library; retain before/after bytes |
| `.` | `go test -count=1 ./lib/... ./app/...` | Both actual test suites pass |
| `.` | `go build -o bin/double ./app` | Creates the command binary |
| `.` | `./bin/double` | stdout exactly `14\n` |
| `.` | `go run ./app` | stdout exactly `14\n` |

The app module path is finite.local/cmd, its directory is app, its package name is
main and its local import is finite.local/number. The library's package name is
number. The workspace selects local module directories; it does not rename their
imports. `go build -o` leaves the named binary, while `go run` runs a temporary
build. `go test` compiles and executes assertions in both packages.

## Wrong examples and useful feedback

For each example use a separate fresh temporary root, repeat the setup and apply
only the named change. Keep the real failed command/output and return to a fresh
correct control, rather than accepting a timeout as the expected failure.

- Missing local module: omit `go work use ./lib`, then `go run ./app` must fail
  because finite.local/number cannot resolve offline. Add the missing use entry,
  rerun successfully and retain the repair evidence.
- Wrong declared library path: use `go mod edit -module=finite.local/incorrect`
  in lib, then `go run ./app` must fail. Explain that a directory named lib does
  not override the module path required by the import; correct the module path
  and rerun successfully.
- Non-executable package: in an otherwise correct copy, `go run ./lib` must fail
  with a package-not-main diagnostic. Run ./app successfully in that same copy.

Exact Go diagnostic wording can differ by version. Feedback must identify the
module/package/entry-point mistake and connect it to the actual failed command.
The three expected failures must be ordinary command errors, never timeouts.

## Checklist and evidence boundary

Reviewer verifies actual before/after module and workspace files, distinct
module/package/import/workspace explanations, correct target selection, actual
binary/run output and test assertions, reproduced wrong examples with repairs,
and isolation/version/exit evidence. The learner must perform the edits and
commands; supplied Go application/test code earns no implementation/testing skill
credit. Module/workspace work is the prerequisite encounter before build/run.
Each of the two first slots receives only its own actions. Both independent
application slots remain open. No automatic route, progress or numeric coverage
is changed by this manual review.

Retain evidence with the content digest and dated reviewer result in review/;
delete only the temporary execution roots after collecting it. Actual release
commit binding is deferred to a release manifest; never invent a commit.
