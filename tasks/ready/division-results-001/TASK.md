# Repair multiple results

## Context

Repair Divide(total, divisor int) (quotient, remainder int, ok bool).

## API

```go
func Divide(total, divisor int) (quotient, remainder int, ok bool)
```

## Requirements

- If divisor is zero, return (0, 0, false) without panicking.
- Otherwise return the integer quotient, the remainder, and true, in that order.
- Integer division truncates toward zero.
- A nonzero remainder has the same sign as total.

## Constraints

- Inputs are between -1000 and 1000 inclusive. Go uses / for the quotient and % for the remainder. The existing conditional branch handles the zero case; preserve its observable behavior. Keep the signature. An exact division still succeeds, even when its remainder is zero.
