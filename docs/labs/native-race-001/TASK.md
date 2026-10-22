# Interpret and repair a dynamic race

Prerequisites: testing.parallel and sync.data-races-happens-before. First encounter: testing.race-detector.

Exercise(fast bool) starts one writer and one reader, then joins both. Edit only cell.go; keep value int and the accessor signatures.

- Run the unsafe main-path command in PROCEDURE.md. Map the report's read, write and goroutine creation stacks to saved source lines.
- Explain why the common start and later join do not order the conflicting worker accesses.
- Synchronize reads and writes with one per-cell mutex, including the alternate branch. Preserve both observable value checks and the supplied exercised path.
- Rerun that main-path command, then both paths with race checking and ordinary assertions.
- Execute the write-only-lock and main-only-fixed specimens. Explain the surviving conflict and why a clean main-path run leaves the unexecuted alternate branch unknown.
- Submit your repair, outputs and a scoped conclusion about the workloads executed.

Use fresh disposable copies. No sleeps, stress loops, detector suppression or universal race-freedom claim.
