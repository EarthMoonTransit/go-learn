# Review notes

## Decision

Returned to draft. The workspace boundary itself prevents editing the grader,
but the checks do not prove incremental output, Unicode whitespace handling, or
nil-error short writes. The read-error mutant loses bytes returned with the
error; it does not distinguish a correct read-all implementation. The API block
also shows `pipeline.Run` as invalid Go syntax.

## Required repair

Use a coordinated reader/writer to prove output before a later read, a short
writer returning `(n < len(p), nil)`, and Unicode whitespace cases with distinct
mutants. Clarify editable versus fixed package APIs and record whether the
upstream reader/writer API shape influenced authoring.
