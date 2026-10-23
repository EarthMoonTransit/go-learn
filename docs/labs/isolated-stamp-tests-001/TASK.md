# Isolate parallel stamp fixtures

WriteStamp writes stamp.txt as label, newline, decimal Unix seconds, newline.

- Introduce your own narrow clock seam while preserving WriteStamp(dir,label string) error and its real-clock default.
- Author TestIsolatedStamps with three distinct labels and fixed clocks. Acquire independent private directories and register cleanup immediately with each owning case.
- Run cases and their readback descendants with t.Parallel. Each descendant reads after its case body returns; its directory must survive until that descendant finishes.
- Assert complete file contents, successful writes, and one clock call per write. Ancestor Cleanup must find distinct paths already removed; retain useful failure diagnostics.
- Repeat with parallelism one and three. Explain the permitted parent/descendant/cleanup order without asserting sibling order.
- Follow PROCEDURE.md in disposable copies. Submit your seam, tests, native output and ownership explanation. No sleeps, ambient test clocks, process-global mutation, sibling barriers or parent-body waits.
