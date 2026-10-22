# Finite-v1 review

Contract and actual content/method/execution independently accepted by /root/reviewer2026-10-22. Final batch gates remain separate.

| Requirement | Public assertion | Defect/method |
| --- | --- | --- |
| initial quiet interval starts at call, no early or late return | TestQuietIntervalStartsAtCall at virtual9s and11s for quiet10s | half-interval: returned early; double-interval: required boundary |
| each value resets interval, including zero and duplicate | TestEveryArrivalRestartsQuietInterval | no-reset: returned before reset deadline; half-interval: early; double-interval: required boundary |
| all values preserve arrival order, zero and duplicates | TestClosedInputPreservesAllValues plus arrival sequence | drop-zero/deduplicate/reverse-order: actual and expected values |
| quiet returns false and closure returns true | empty-open, reset, closed-empty and closed-sequence tests | true-on-idle/false-on-close: actual and expected exhaustion flag |
| closed input returns promptly, including after accepted values | TestClosedInputPreservesAllValues; TestCloseAfterAcceptedValueReturnsPromptly | wait-on-close: closure must return without waiting; fixture advances virtual clock for cleanup |
| single owned timer, Stop on all exits, Reset, input ownership and no producer | actual source inspection and finite learner checklist | Stop is method evidence, not automatic timer-leak proof; no co-ready priority promised or tested; ticker application stays open |

Every behavioral mutant must fail the public-only copied fixture with a useful assertion. Timeout, panic, compile error or cleanup failure alone does not establish detection. Method checks inspect actual code; signatures alone earn no mechanism credit.

Independent actual method and distinct-action acceptance: Actual one owned Timer, initial interval, Reset after accepted values, Stop on every return and caller input ownership accepted. Virtual-time behavior checks avoid co-ready priority; no automatic timer-leak proof. Ticker application remains open.
