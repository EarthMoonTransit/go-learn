# ADR 0019: Race-aware verification across public checks and defect models

Status: accepted for implementation; release gated by independent review and validate.

## Problem

Task Contract v2 ran the race detector only as an auxiliary reference/full learner check. A no-lock Counter can preserve every ordinary assertion while racing. Its realistic defect model survived task-check, and learner Check could pass code later rejected by Complete.

## Decision

For v2 tasks with `race_required: true`, run every public/full, starter/reference and mutant test invocation with `-race`, including test-writing submissions and workspace patches. Keep explicit auxiliary race/platform reporting and vet. A compiled race-test failure counts as an observed defect; assembly, compilation, unsupported platforms and timeouts never count as detection. For `race_required: false` and legacy tasks behavior is unchanged. No task schema or skill identifiers change, and no new profile is introduced.

The collection execution contract advances from `task-check/v2` to `task-check/v3`; both collection and aggregate-report schemas enforce it. A collection produced under ordinary-only mutation semantics must not certify race-aware content.

## Migration and recovery

Regenerate one full collection with the new checker, then reuse it in validator and report-check. Old v1/v2 collections are rejected rather than upgraded or relabeled. Rebuild/reinstall local tooling as usual. Task IDs, authored task digests, route numbers and progress v3 storage remain unchanged; existing solutions, completions, attempts, hints and time are retained. A subsequent Check/Complete applies the stronger verification; historical completion is not erased. No automatic filesystem migration occurs.

## Evidence

Fast tests cover request propagation in fixed, mutation and workspace profiles and rejection of old collections. A real disposable-workspace regression proves a behavior-correct but unsynchronized implementation passes ordinary tests, fails race verification, and is detected as a mutant; public learner verification exposes the race diagnostic. Tooling acceptance includes fast/race/vet/integration plus one shared full collection. Concurrency content follows as a separate batch.
