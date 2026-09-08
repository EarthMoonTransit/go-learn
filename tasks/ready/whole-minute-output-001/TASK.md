# Repair conversion of a fractional rate

## Context

Repair the supplied operation.

## API

```go
func SuccessfulPerMinute(rate int, percent float64) int
```

## Requirements

- Return the whole successful units per minute from an hourly rate and percentage.
- Apply the percentage and divide by 60 before discarding the fractional result toward zero.
- Preserve fractional percentages during the calculation.
- rate is between 0 and 1000000; percent is finite and between 0 and 100 inclusive.

## Constraints

- Use only the Go standard library.
