# Create local modules, then build the intended command

Manual first encounters: `toolchain.modules-workspaces`, then `toolchain.build-run`.
This unit has no automatic completion or coverage credit. Complete it before the
automatic route; basic package/function syntax is supplied in the source files.

Work only in a fresh disposable copy following the authorized procedure in
[finite-v1](../../plans/finite-v1.md#authorized-finite-manual-procedure-owner-decision-2026-10-22).
The four supplied `.txt` files become `lib/number.go`, `lib/number_test.go`,
`app/main.go` and `app/main_test.go`. Keep their contents unchanged. Write the
module/workspace files yourself; do not change the repository or your real Go
configuration. Only local Go, the standard library and these two modules are used.

1. Create a module for the library, then deliberately correct its module path to
   `finite.local/number`. Create the command module as `finite.local/cmd` and
   declare its dependency on `finite.local/number v0.0.0`.
2. Create a workspace containing both module directories. Explain the difference
   between their module paths, package names, import paths and the workspace.
3. Build the executable command into a disposable `bin/double` file. Execute it
   and separately use `go run`; both must print exactly `14` followed by newline.
4. Test both intended packages. Explain why successful compilation alone does
   not establish that the assertions passed, and where the built executable is.
5. In a second disposable copy omit the library from the workspace. Reproduce
   the failed command, identify the missing local dependency and correct it.
   Also explain the diagnostic produced by trying to run the library package.

Submit the actual go.mod/go.work files, exact commands, Go version, exit statuses,
outputs and short explanations. Each command has a 60-second limit. Reference
commands, isolation details and the finite assessment checklist are in
[PROCEDURE.md](PROCEDURE.md). No timing threshold, network dependency or future
runner is involved. A written procedure without actual execution is unfinished.
