# Work in a private Go workspace

Your goal is to edit one exercise in GoLand and understand what each check tells
you. Start with the existing numbered route; the task ID is its stable command
name. Reading this page does not start a work session.

Open an exercise from Practice below. Start creates a private workspace and
prints its path. Open that directory in GoLand with the Open in GoLand button,
or copy the path if the launcher is unavailable. Read TASK.md before editing.
Change only the files the task marks editable. The repository's canonical task
package is the input to verification; it is not your working directory.

A Go file belongs to a package. A module's go.mod defines its module path and
language version. Keep the workspace's supplied go.mod. GoLand should use the
project's Go 1.27 toolchain. Imports name packages, not arbitrary source files.
Here is a separate example of a local constant and a runnable example test:

```go
package example

import "fmt"

func Example() {
    const separator = "::"
    fmt.Println("left" + separator + "right")
    // Output: left::right
}
```

This example illustrates a file's structure; do not paste it into the exercise.
Its package and output are unrelated to the task's requested declarations.

After editing, use Check public tests in the UI or `learn check`. A public
failure should identify a visible requirement or a compilation problem. Correct
that issue and run Check again. Complete runs the full verifier; only successful
full verification records completion. Passing public tests is useful feedback,
but is not a promise that every requirement has been tested on every input.

For test-writing exercises, supplied editor declarations help GoLand resolve
names but do not implement the tested behavior. Use the learner Check for the
real test-submission workflow. When pausing work, use Pause explicitly. You can
return to Theory at any time; opening an explanation does not consume a hint,
change your current task or record an attempt.

Common mistake: opening the canonical repository package instead of the printed
workspace. Confirm the path before editing; keep existing work rather than
starting over if a workspace needs attention.
