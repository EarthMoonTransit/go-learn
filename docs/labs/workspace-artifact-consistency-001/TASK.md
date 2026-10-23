# Repair module selection and refresh an executable
## Context
The app's local replacement points to a wrong library copy.
## Requirements
- Predict and reproduce the wrong result. Add the intended core module to the workspace and inspect the selected module.
- Disable workspace selection, reproduce the standalone discrepancy, then repair the app's relative replacement. Preserve supplied source behavior and verify both modes.
- Build the intended command into a named disposable executable and retain its output and hash.
- Change only the permitted source label. Compare the old executable with go run, then rebuild and verify the updated executable.
- Explain workspace versus module selection and retained versus temporary build artifacts using actual files, diagnostics and outputs.
## Constraints
- Edit only go.work, app/go.mod and the label constant. Use fixed local offline copies, bounded commands and independent method review; no automatic coverage or progress.
