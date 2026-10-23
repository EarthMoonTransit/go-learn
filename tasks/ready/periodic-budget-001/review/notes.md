# I10 timer independent content acceptance

Original approved154-word contract; task-init and ADR0008 sources=[] provenance. Sole independent actual source/evidence/method review is approved; see review/independent-review.json.

| Basic requirement | Public check | Defect / useful assertion |
| --- | --- | --- |
| First visit follows period | FirstPeriodAndRemaining | immediate-first, half-period, double-period: visit count or exact duration |
| Actual elapsed/remaining at callback entry | DelayedVisitUsesCurrentElapsed | tick-count-elapsed, tick-timestamp-elapsed, remaining-equals-budget: actual/wanted pair |
| Keep native ticker cadence | DelayedVisitUsesCurrentElapsed | restart-period-after-callback: delayed next callback missing before a newly reset period |
| Independent finite total budget | IndependentTotalBudget, period longer than budget | budget-only-on-tick, reset-budget-after-visit, false-budget-result: missing boundary return/bool |
| Synchronous callback and false-stop | ImmediateCallbackStop, delayed/held tests | ignore-visit-stop, true-stop-result, background-visit |
| No callback after budget | HeldCallbackOutlivesBudget | allow-expired-visit: expired visit count; background worker returns before held callback release |

Fourteen required wrong variants all keep finite budget or callback-stop escape. No zero/negative-ticker panic fixture. Tests observe virtual time only with NewTimer/Stop, channel gates and synctest.Wait. Assertions inspect held state while callbacks durably wait; every cleanup path first releases its callback gate, then sets the rescue flag/advances finite budget time. No mutex held during Wait; no Fatal bypasses release. Timing observations use offsets away from tick/budget ties. When budget/ticker are co-ready after a held callback, only no-expired-callback policy is asserted, not select priority or event replay. The current elapsed sample is independent of old tick timestamp; no every-missed-tick guarantee is claimed.

## Actual method and alternatives

Learner owns one real Ticker and one independent unreset budget Timer, derives elapsed from a call-entry origin, subtracts it from budget, invokes callback synchronously and Stops both clocks on all exits. time.Now/Since occurs only in implementation under native synctest; test AST guard stays unchanged. Correct Now().Sub and helper cleanup controls pass. Omitted Stops intentionally pass behavior and must fail method review; no automatic timer-leak proof or GC-based ownership credit. No worker/cancellation/context/component credit. Firsts idle-stream and duration-rounding remain unchanged preparation.

## Independent actual acceptance

Reviewer /root/reviewer accepted this final source and finite evidence on2026-10-23; exact immutable decision is retained in review/independent-review.json. Behavior-passing method controls are explicitly rejected by the required checklist. Official lifecycle transitions are complete; integration and all mandatory batch local/CI gates remain separate.
