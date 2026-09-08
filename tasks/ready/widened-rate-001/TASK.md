# Combine a narrow rate with a wide input

## Context

Complete the specified operations.

## API

```go
func BandRate(value float64) float32 // supplied
func BandAmount(value float64) float64
```

## Requirements

- Use the supplied BandRate result as the percentage for value.
- Return value multiplied by that float32 rate widened to float64, then divided by 100.
- Preserve the sign and fractional input; zero returns zero.
- Inputs are finite and between -1000000000 and 1000000000; allow absolute error 0.00000001 or relative error 0.000000000001.

## Constraints

- Edit subject/amount.go; subject/rate.go is fixed.
