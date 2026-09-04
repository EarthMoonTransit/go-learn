# Editor API declarations for learner test submissions

Date: 2026-09-05. Separate tooling repair for the first test-writing lesson.

A temporary export of the revised clamp task reproduced the missing declaration:
the learner workspace had TASK.md, go.mod, its marker and subject_test.go, with
no ordinary Go file declaring Clamp. The checker already assembled it correctly
in a disposable workspace. No actual GoLand UI trial is claimed.

The repair generates a read-only declaration aid for supported APIs. Static
Go type checking confirms that the declarations resolve. Source bodies,
comments, values and private declarations are absent. Exact-byte protection,
exclusion from submitted files, rejection of extra implementation files and
compatibility with older no-stub workspaces are covered by tests. Complex and
platform-specific APIs retain the documented checker-only behavior.

Independent review reproduced two edge defects before final acceptance: nested
callback names survived normalization and platform variants produced duplicate
declarations. Editor-only recursive normalization and whole-API fallback repair
them without changing the shared public-API fingerprint. The reviewer reran
both independent regressions and the boundary tests and approved the result.

Final fast tests, race, vet and tagged integration passed. One full 25-package
task-check collection was reused by phase1 validation and report checks. The
editor-only review repairs did not change tasks or verification contracts, so
the collection remained valid while the affected final test graph was rerun.

An actual clamp workflow used temporary learner homes: the new declaration was
present, starter tests passed the correct implementation, the weak suite was
rejected at completion and the full test submission was accepted. An older
workspace without the aid still passed public verification and was not modified.
Authored tests ran only through the disposable learner runner. The real learner
progress file retained its SHA-256. These are automated trials, not human
completion, calibration or mastery evidence.
