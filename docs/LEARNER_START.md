# Start in GoLand

The repository contains the course engine. Each exercise gets its own private
Go module. Read the exercise's TASK.md and edit only the files it names.

Before the first exercise, understand these two toolchain basics:

- A module is a group of packages identified by the module line in go.mod.
  A package groups Go files in a directory; their package declaration names it.
  The course supplies go.mod and the verification harness.
- Go compiles a package before running its tests. Compilation errors point to
  invalid code; a failing assertion means the code compiled but its behavior
  did not match a requirement. Here, learn check assembles and tests a disposable
  copy. learn complete runs the full verification.

This is guided preparation, not an automatically completed task or proof of
toolchain mastery. Later toolchain exercises will assess these skills directly.

## Your first workspace

Run learn next, then the printed learn start command. Open the printed Workspace
folder as a GoLand project. For the first exercise the editable file is
subject/bindings.go; subject/api.go is supplied plumbing and stays unchanged.

Use the GoLand terminal:

- learn roadmap: see the course blocks and your place.
- learn next: see the next lesson, its preparation and the start command.
- learn check: try your current code.
- learn hint: request one progressive hint.
- learn complete: submit for full verification.
- learn pause: stop the session clock before a break.
- learn resume TASK-ID: continue an existing workspace.
- learn review: read the post-completion review prompts.
- learn feedback "your feedback": save a local note.

Keep TASK.md and go.mod unchanged. New lessons can have different editable files;
the task always names them. Public and full checks use temporary copies, so
hidden material never needs to be opened in GoLand.

An earlier completion remains completed even when it appears later in the route.
Read-only commands do not start a lesson, consume a hint or change your timer.
