# P17/P19: bounded concurrency source decisions

Baseline: `f458a1aad95626d74c901749fa1664862db0e036`. Exact registry IDs,
commits, artifact hashes, planning hashes and decision boundaries are in
`catalog/intake/2026-10-18/p17-p19-source-review.json`. Independent source review approved both decisions; task content and
exact-head CI/merge gates remain separate. This source-only batch changes no registry scope, task, taxonomy,
theory, learner state or release/export status. It adds no coverage encounter.

## P17: use the suitable finite fan-out donor

Select Ardan's `topics/go/concurrency/channels/exercises/template2/template2.go`
at `4ae5a46bfc665909a7cb29e11c18775330c63401`, already inspected in the S3
bounded template set. The existing registry adapt mode permits consideration;
this P17 decision selects the exact donor. S3 supplied inspection evidence and
explicitly selected no content donor. Its operation
starts a finite group, receives one result per goroutine, and collects all values
in a slice. A suitable donor exists; calling P17 original would be unjustified.

The adaptation replaces random generation and printing with supplied operations.
It extends results with input index, value and exact error, retains nonzero values
returned with errors, and waits for every started operation after any failure.
This retains the collect-all operation and adds a useful error/accounting policy.
No fail-fast, cancellation or panic recovery is implied. Do not copy upstream
answer implementations or test fixtures. Preserve the Apache-2.0 header/full
license and applicable notices, credit contributors and mark modifications.

Selected-path history was checked through the GitHub API because the verified
local cache is shallow. Commit `50e6bde1` introduces fan-out by changing an
existing template in 2016; it is not the first creation of the file. The two rename commits
share blob `90847ada` with each other; it differs from the earlier fan-out blob. Jacob Walker's `b71a1313` adds slice collection in 2018;
`4dcac3b3` refines comments and yields the blob matching the registry pin.
William Kennedy and Jacob Walker are in pinned CONTRIBUTORS. Only the root
LICENSE applies among selected ancestor notice files. Current repository
metadata says nonfork. These are bounded findings, not exhaustive origin proof.

SHAD tparallel concerns testing parent/child scheduling, rather than indexed
results/errors from supplied operations. It remains reference-only outside the
existing LRU allowlist. Ardan's filter/open-ended templates and bounded ping-pong are less
direct donors than template2. No restricted code or prose is repurposed.

P17 assesses `goroutines.result-error-collection` as an implementation transfer.
JoinAll only accounts for void callback completion. Add its direct taxonomy
prerequisite `goroutines.state-ownership` alongside launch, WaitGroup, error and
slice prerequisites; the proposal's list alone omitted that requirement.
Place it after JoinAll, both mutex tasks and P15, before P19. Stable addresses
are appended after verifying the current maximum; prior addresses/roles stay.

Public feedback must show zero jobs, mixed results/errors, nonzero value plus
exact error, exactly-once calls, controlled reverse completion and input-index
preservation. Hold jobs to establish concurrency and continued waiting after an
early error. Release all gates on failures. Require race, deterministic gates
and compiled wrong variants; a timeout or compiler error proves no defect kill.

## P19: a specifically justified original

None of the six inspected S3 Ardan templates captures a preexisting `=` iteration
variable behind a gate released after the loop. The append-race exercise does
not capture its loop index. Reusing only their loop/goroutine shell would be
nominal adaptation. This is a bounded comparison of queued material, not a claim
that no matching exercise exists anywhere.

The pinned Go specification distinguishes per-iteration variables introduced
by range `:=` from assignment to existing variables through `=`. Its header
describes Go1.27; the development checkout itself must not be called a Go1.27
release. The memory model's channel-close ordering explains why a post-loop gate
can make bad capture deterministic and race-free. These documents supply facts
only; no examples, prose or task implementation are copied.

Existing category/period closures retain constructor state; method-value capture
binds receivers; StartJob launches one supplied callback; range-offset observes
element copies. None assesses the proposed reused-variable capture contract.
An original debugging task is justified, with `goroutines.capture-arguments`
as primary. Preserve the supplied post-loop gate and joining mechanics; avoid
teaching obsolete pre1.22 assumptions about `:=`. Accept equivalent correct
argument binding and per-iteration declarations. No output order is promised.

Public checks need exact multisets including repeated, zero and negative values,
nil/empty inputs and independent batches, with safe release/join on failures.
Race remains required but cannot replace deterministic wrong-value assertions.
The editable mechanics determine any extra prerequisites beyond closure capture,
launch, channel operations, range and WaitGroup. Place it after P17.

Future original provenance follows ADR0008: project identity, actual creation
record, content digest and empty external task sources. Semantic references stay
in this decision evidence. Bind the digest to an actual commit at release;
never fabricate a commit or legacy registry entry. Final task content review,
public feedback proof, official promotion and all repository gates remain required.
