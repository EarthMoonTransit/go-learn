# Exact finite selected-tools procedure

Only the six fixed specimens and command entries in manifest.json may run, each
inside a fresh canonical disposable copy of initial_files. Use the approved finite
manual procedure, no source-checkout execution. Pinned local Go1.27 and its gofmt
are the only executables; private home/config/cache/modcache/tmp, telemetry-off,
GOENV=off/GOWORK=off/GOTOOLCHAIN=local/GOPROXY=off/GOSUMDB=off/CGO_ENABLED=0.
GOFLAGS is exactly -mod=readonly or -mod=readonly -tags=preview as each manifest
entry says. Inspect its actual value. No network, arbitrary command or global
configuration. Each command≤60seconds; retain argv/cwd/environment/status/output,
duration, predictions and source snapshots; remove only owned disposable roots.

Only format/preview.go may change. Supplied default, tests, main and module remain
identical. The initial preview has a final return after unreachable ordinary code,
so it is well typed and go vet, not compilation, must diagnose unreachable code.
Tests isolating behavior explicitly disable implicit vet; explicit vet is separate.

Reference and correct-alternate: after go version inspect default env, predict
its GoFiles, list, test and vet; all pass. Inspect preview env, predict its files,
list, observe nonempty gofmt diff (native Go1.27 exit1 is expected), apply gofmt and require empty second diff with exit0.
Preview vet must find unreachable code. Read actual QueryEscape and PathEscape
documentation and record their relevant distinction. Install wrong/api-only source
(removes only unreachable code, retains wrong path API); preview vet passes but
ordinary query assertions fail. Install reference or alternate alias source;
preview and default test/vet all pass. The alternate uses a local alias/result.

Wrong-mode: validate only default first, then record why that does not establish
preview. Explicit preview env/list/vet exposes the selected defect. Install correct
reference and require preview test/vet success. Format-only: native format diff,
apply and empty diff leave preview vet failing; correct reference then passes.
Unreachable-only and wrong-api: install wrong/api-only source first. Native docs
lookups both succeed, but actual API interpretation and ordinary query assertions
are still necessary. Vet passes while tests fail; correct reference then passes.
Their distinct submitted explanations are rejected by actual method review, not
counted as distinct new implementation defects.

Keep original meaningful diagnostics; no compile failure, timeout or crash earns
ordinary feedback. Actual source review verifies only permitted edits, formatting
idempotence, target interpretation before listing, documented semantics applied,
and analyzer versus behavior distinction. No automatic coverage or learner state.
Exact source/command preexecution review and independent actual acceptance follow.

A nonempty gofmt -d result is a successful formatting observation with native exit1, not an ordinary test/vet defect diagnosis. The empty second diff requires exit0.
