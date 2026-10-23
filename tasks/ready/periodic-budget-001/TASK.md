# Run periodic work within a budget

## Context
Run periodic callbacks.

## API
```go
func RunPeriodic(period, budget time.Duration, visit func(elapsed, remaining time.Duration) bool) bool
```

## Requirements
- Own one time.Ticker and one budget time.Timer, starting at entry; the first callback follows a tick.
- Invoke visit synchronously for received ticks. Measure elapsed time at callback entry, and pass elapsed and budget-elapsed; never derive them from tick count or the tick timestamp.
- Never start a callback when elapsed reaches budget. Return true when the budget expires, including when period exceeds budget.
- False from visit ends with false; otherwise continue. Join each callback before returning.
- Stop both owned clock objects on every exit. Do not launch workers or reset either clock.
- Missed ticks need not all be delivered. Method review checks clock ownership and arithmetic.

## Constraints
- Period and budget are positive and at most one minute. visit is nonnil, finite and nonpanicking.
