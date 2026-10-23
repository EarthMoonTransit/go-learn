# Diagnose the selected source before repair

Prerequisites: modules-build, format-vet and docs-environment.

The supplied query renderer has default and preview implementations. Edit only preview.go.

- Inspect go env GOFLAGS and go list; predict and record selected files before each listing with preview disabled and enabled.
- Find the actual url documentation needed for query escaping. Explain the distinction from path escaping and repair the preview call without altering input.
- Record the initial formatting diff, apply gofmt, then show an empty second diff. Separate whitespace changes from functional edits.
- Capture and interpret the preview vet finding; remove its cause. Demonstrate why default-mode success alone cannot validate preview.
- Run supplied assertions and vet in both modes. Retain wrong-mode, formatting-only and wrong-API evidence, then show repairs.

Use the exact reviewed disposable procedure. Submit source snapshots, predictions and actual tool output; supplied tests earn no test-writing credit.

