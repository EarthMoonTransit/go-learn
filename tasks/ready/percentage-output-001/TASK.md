# Calculate a fractional successful output

## Context

Complete the specified operations.

## API

```go
func SuccessfulPerHour(rate int, percent float64) float64
```

## Requirements

- rate is an integer from zero through 1000000; percent is finite and between zero and 100.
- Return rate multiplied by percent divided by 100 as a float64.
- Preserve fractional percentages and fractional output; do not truncate to an integer.
- Zero rate or zero percent gives zero; 100 percent gives the original rate. Absolute error up to 1e-9 is accepted.

## Constraints

- Use only the Go standard library.
