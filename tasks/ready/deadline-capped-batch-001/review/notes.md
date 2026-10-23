# I7 shared budget application review

Independent contract approval precedes authoring: deadline-contract-review.json, section4d7b495931d5b9e00c05bc816362b3c38fc95e9b4a475e025a04ad29f0cdbe2b. Actual source/evidence/method review independently approved by /root/reviewer; batch delivery gates remain separate.

| Requirement | Public assertion | Wrong variants |
| --- | --- | --- |
| ordered exact callbacks/results, including errors and empty | TestBudgetPublicCompositionAndResults | phantom-empty, skip-last, repeat-operation, reverse-order, stop-on-error, drop-error-value, replace-error, return-context-error |
| earliest deadline and zero ceiling | public budget/ceiling/parent matrix | omit-budget, double-budget, zero-is-cap, ignore-ceiling, ceiling-extends-budget, detached-parent |
| shared non-renewable scope | live Done comparison and TestSharedBudgetIsNotRenewedAfterWork | renew-budget, later-parent |
| expired/canceled input still invokes every callback | zero/negative/past/canceled public cases | skip-expired, positive-default |
| owned cancellation after success/error | captured-context state after return | no-release, release-only-error, cancel-before-work |
| cancellation and synchronous cleanup | TestParentCancellationWaitsForCleanup and shared-budget test | detached-parent, return-before-cleanup |

Tests use native synctest and entry-created4/6second timers, never time.Sleep/Now/After or wall thresholds. Callback2entry OR wrapperreturn prevents skipped-callback hangs. Facts are mutex-protected or buffered notifications, not late test errors from background callbacks. Rescue/cleanup release happens before wrapper join and bubble exit. Declared wrong variants have finite escape paths; timeout/deadlock/panic/race is never evidence for these policies.

Actual method must inspect one entry-time allocation, absolute-cap composition, shared final operation scope, synchronous calls and release of every owned helper. Correct alternative avoids unnecessary cap construction. Hidden intermediate cancellation omission can pass behavior but must fail method review; no automatic leak-proof claim. The supplied Result/signature, test timers and context markers earn no extra skills. Old timeout-operation first is unchanged; shared multi-callback budget and independent ceiling are distinct actions.

DEADLINE-CLEANUP-01: focused early-return execution exposed a first callback still waiting its native two-second timer after the wrapper returned. Compact full-suite diagnostics had hidden the later bubble panic; those original receipts are not finite-failure acceptance. The first callback now selects timer OR independent rescue. Assertions still precede release, so normal five-second versus renewed-seven-second observations are unchanged. Fresh canonical/public/control and focused timed probes are required for final evidence. No checker change.

## Independent actual acceptance

Sole independent reviewer /root/reviewer approved actual source, executed evidence, provenance, public feedback and finite method on2026-10-22. Current decision is accepted for official promotion; final batch local gates, final diff review, exactHEADvalidate and merge remain separate.

Actual entry-time budget, absolute cap composition, shared final scope and synchronous ordered callbacks inspected. Every owned helper cancellation is released.22useful public defects and valid alternate pass; hidden intermediate release omission passes behavior but fails finite method review. No automatic leak-proof claim.
