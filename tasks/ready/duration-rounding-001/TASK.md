# Repair bounded duration conversion with truncation

## Context

Repair the supplied operation.

## API

```go
var ErrDuration error
func WholeMinutes(value time.Duration) (int64, error)
```

## Requirements

- Return the number of whole minutes in a duration from zero through 24 hours inclusive, truncating a positive remainder.
- Return 0 and the supplied non-nil ErrDuration for negative durations or values above 24 hours.
- Preserve the error sentinel. Use duration units, not a guessed conversion from nanoseconds.
- No wall-clock input, floating-point rounding or local timezone is involved.

## Constraints

- Use only the Go standard library.
