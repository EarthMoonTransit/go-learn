# Finite-v1 review

Contract approved by /root/reviewer before authoring. Actual content/method and canonical/public-only7/7 accepted by /root/reviewer2026-10-22; applicable race passes.

| Requirement | Public assertion | Defect/method |
| --- | --- | --- |
| lazy construction | TestLazyAndCompletedPublication before calls | eager; release initializer before reporting further results |
| exactly once, wait for completion, complete value/error identity | TestLazyAndCompletedPublication held/released/late phases | repeated/retry-error/drop-value/drop-error/return-early; phase/result/count diagnostics |
| independent zero/nonzero successes | TestIndependentInstancesAndZero | shared-instances |
| actual sync.Once and captured publication | reference source and finite explanation | nonpanic lifecycle only; errors cached not retried |

Public-only specimens must detect every behavioral defect with useful assertions. Panic, timeout or compile failure does not establish detection. Method review examines actual source; behavior pass alone earns no method credit.

The initial fixture waited for synctest quiescence while Once mutex contenders existed; mutex waits are not durable blocking, so that attempt is rejected (no timeout earns detection). Corrected fixture observes the held first call alone, starts three other callers, releases the initializer before Wait, then checks every complete result and total count. It does not claim the start acknowledgements prove every contender reached Once before release. Actual Once completion/publication is independently method-reviewed; race and wrong-result/count assertions remain mandatory.

Independent application evidence: This is a bounded first; reference uses Once completion to publish captured value/error and caches failures. No reset/teardown/panic semantics credited. Corrected public-only7/7/race and canonical7/7/race pass.
