# Repair ordered calls and preserve dependency errors

## Context

Repair the supplied operation.

## API

```go
type Calculator interface{Amount(int)(float64,error);Factor()(float64,error)}
func Divide(c Calculator, count int) (float64, error)
```

## Requirements

- Call Amount with count, then Factor, each exactly once on success; return Amount times Factor divided by count.
- On the first error stop immediately and return zero with that exact error; do not call a later dependency.
- The calculator is non-nil, count is 1..1000, successful values are finite with magnitude <=1000. Allow absolute error 0.000000001 for numeric results.
- Preserve error identity even when the dependency also returns a nonzero numeric value.

## Constraints

- Use only the Go standard library.
