# Review a stable report with a golden file

Prerequisites: testing.basics-diagnostics and io.files-cleanup (resource-finalize-001, then filesystem-record-tests-001). This manual first encounter assesses testing.golden.

Report returns exactly `RECORDS v1` plus newline, followed by 32 rows indexed 00 through 31: `NN | item-NN | square`, each with a final newline. The source and canonical baseline file are supplied. Write your own test reading testdata/report.golden and comparing the complete output, with a focused first-difference diagnostic. Ordinary tests must never rewrite expected data.

Demonstrate rejection of header, order, value and trailing-newline errors. Then enact the specifically approved header-only format change to `RECORDS v2`. Require an ordinary mismatch, an explicit local update flag, retained before/after diff, review of the intentional delta and a final ordinary passing test. No implicit CI update is allowed.

Follow PROCEDURE.md in fresh disposable copies. Submit your actual test, normal-run file hashes, explicit-update diff and decision. Supplied report generation and baseline earn no implementation credit. A test that silently rewrites its expectation is incomplete even if go test succeeds.
