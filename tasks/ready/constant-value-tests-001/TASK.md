# Test fractional constants and their exported uses

## Context

Write tests for the specified contract.

## API

```go
const Rate = 0.05
const Days int = 365
func Rate32() float32
func DayCount() int
func AnnualFraction(days int) float64
```

## Requirements

- The supplied Rate constant is 0.05 and Days is the integer constant 365.
- Rate32 returns Rate converted to float32, and DayCount returns Days.
- AnnualFraction returns days divided by Days using floating-point division; days is 0..10000.
- Check exported constants as well as functions. Rate32 must equal float32(0.05); allow absolute error 0.000000001 for AnnualFraction and reject NaN.

## Constraints

- Use only the Go standard library.
