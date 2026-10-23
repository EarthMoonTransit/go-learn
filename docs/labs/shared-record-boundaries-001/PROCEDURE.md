# Finite manual procedure

This application follows package-boundaries-001, external-window-tests-001 and
injected-config-project-001. Only import-cycle repair and test-package choice
receive application credit. The earlier cycle task used an interface; here the
learner relocates neutral data and composition. Both test declarations and bodies
are initially absent and must be authored independently.

Copy only the five production/module assets listed in manifest.initial_files
into a new private offline disposable project, removing their `.txt` suffix.
Keep the two inert test templates outside the compiled tree until full authored
package declarations and test bodies are ready. Predict both cycle edges. Execute
`go build -o /dev/null ./cmd/preview` before installing tests: Go scans test package headers even during build, so
empty test files would mask the native cycle diagnosis. Never execute authored code from this repository.
Use local Go1.27, private caches/home/config and telemetry-off files, no network,
no auto toolchain download, and at most60seconds for each command.

The six editable paths are record/entry.go, store/store.go, codec/codec.go,
cmd/preview/main.go, store/normalize_test.go and store/consumer_test.go.
Keep Load/Encode/normalize algorithms unchanged. Move only declarations,
imports and composition. Write private normalization cases and an outside-client
Load/Entry test; label the authored tests TestLearner*. Explain why private access
requires package store, whereas the client boundary requires package store_test.
The supplied checks are installed separately after repair and receive no learner
credit. Run learner-only tests first, then all tests. Inspect actual imports with
`go list -f '{{.ImportPath}} -> {{join .Imports " "}}' ./record ./store ./codec`.
Run `go run ./cmd/preview`: output must be exactly `alpha:1` plus newline.

In a fresh copy, the external private-access probe must fail at native compile
only (`go test -c -o /dev/null ./store`); restore the reviewed private test and
repeat compilation and tests. Other finite specimens and every exact argv are
fixed in manifest.json. Wrong normalization and Load values must fail learner-only
assertions without any supplied public assertion running. Record predictions
before commands and preserve actual sources, diagnostics, status and cleanup.

The reference and valid alternative are internal review material. Both-internal
and exported-private controls can pass tests; source review rejects their package
boundary or exported private helper. Compiler diagnostics never substitute for
ordinary semantic feedback. The recipe performs9fresh copies and36commands;
independent preexecution approval of exact source/script hashes is mandatory.
No generic runner, checker profile, automatic coverage or learner progress is added.
