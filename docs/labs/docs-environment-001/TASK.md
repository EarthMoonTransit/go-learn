# Read an API and predict build selection

After modules-build-001 and diagnostic-helper-001, inspect the supplied local parse package using go doc. Use Segment's actual documented contract to complete Render in cmd/main.go. Inspect GOOS, GOARCH, GOMOD and GOWORK with go env; explain the selected module, disabled workspace and target.

Before each go list, record which cmd Go files you predict for darwin/arm64 and linux/arm64, explaining filename and build-tag constraints. Compare actual selection, build both targets and run only the native binary. Submit files, predictions, command results and the observed executable formats. Correct the three supplied mistakes: wrong symbol, wrong target and wrong interpretation of Segment's index/presence contract.

Follow PROCEDURE.md in fresh private disposable copies. Only Render is learner code; the parser, tests and platform markers are supplied. The assessed outcome is finding/applying documentation and explaining environment/build selection. This manual first encounter creates no automatic coverage or progress.
