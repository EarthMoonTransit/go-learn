# Author properties and retain a reduced run-compaction failure
## Context
CompactRuns must keep one byte per adjacent equal run; the supplied implementation is faulty.
## Requirements
- Author FuzzCompact with representative finite seeds and a deterministic callback accepting arbitrary byte values.
- Check independent run-boundary semantics, input preservation and idempotence. Explain why idempotence or global uniqueness alone is insufficient or false.
- Reproduce a failing authored seed; execute the finite reduction sequence and retain the smallest meaningful failure in corpus data and an ordinary Test.
- Show both retained regressions fail before repair, then repair only CompactRuns and rerun the same assertions, seeds and bounded native invocation.
- Reject weakened properties and log-only regressions even when commands pass; retain actual source, diagnostic and reduction evidence.
## Constraints
- Work only in reviewed offline disposable copies. Nil and empty outputs are equivalent. No random-discovery or automated-minimization claim; source/method review gives manual evidence only.
