# Finite workspace and artifact procedure

This is an independent application after modules-build-001. The first created
modules/workspace and built a command. Here two local modules declare the same
module path; only core belongs in the intended workspace. Shadow is a deliberately
wrong implementation selected by the app's initial relative replace. Supplied
application/library/tests earn no implementation credit.

Copy manifest.initial_files as regular text into one fresh private disposable
project, remove `.txt`, and create only its bin directory. Initial go.work selects
app alone. Only go.work, app/go.mod and BuildLabel's constant value are editable;
all other source/module files stay unchanged. Never copy both core and shadow
into the workspace simultaneously. No external module fetching, real user config
or source-checkout execution. Local Go1.27, private HOME/config/caches/tmp and
telemetry-off files, offline/readonly module flags and <=60second command limits
are mandatory. GOWORK is explicitly either this copy's go.work or off per command.

Predict and run the initial workspace app tests: selected dependency gives-6,
not6. Add core with go work use ./core, then test app/core and inspect go list -m.
Switch only GOWORK to off and run from app: the wrong local replace is exposed
again. Inspect the selected module directory. Repair app/go.mod with a relative
replace to ../core; verify standalone tests/list and workspace tests. Directory,
module path, package and workspace are different facts; retain actual files and
selected paths, replacing private root names only when normalizing evidence.

Build the intended app into bin/inspect; execute it and retain hash/output
`6 before`. Change only the supplied BuildLabel to after. The retained executable
still prints before; go run uses current source and prints after. Rebuild the
same artifact, then it prints after and has a different hash. A passing go test
or go run does not refresh the earlier binary. Preserve binaries as hashes/size
in snapshots, not decoded text or committed executable blobs. Every executable
command is the fixed reviewed bin/inspect built inside that copy.

The fixed workspace-only repair control passes with core in go.work but still
fails standalone. A separate wrong replacement to ../absent is hidden by workspace
selection yet produces a native missing-local-module error standalone; repair and
retest. Stale-artifact control deliberately executes an old but valid binary:
exit0 with before is expected evidence and fails a claimed refreshed-artifact
method, never a fabricated command error. Rebuild demonstrates correction.
A valid alternate uses ../core/. and builds from app with GOWORK=off; all reference
mode and artifact observations remain required.

Exact manifest contains5fresh copies and60commands including Go versions and
reviewed local binary executions. Retain cwd/argv/environment, before/after
module/workspace/label bytes, actual outputs/status/duration and binary hashes.
Unexpected outcomes stop without silent retry. Every root is removed. Source and
exact command recipe require independent preexecution approval; actual results
need separate review. Two manual application credits only; no new runner/profile,
automatic coverage/progress, extra package/testing or mastery claim.
