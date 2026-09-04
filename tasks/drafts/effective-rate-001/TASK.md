# Compute a fractional effective rate

## Context

Compute the completed-item rate from an hourly capacity and a fractional success percentage.

## API

```go
func EffectiveRate(itemsPerHour int, successPercent float64) int
```

## Requirements

- Treat `successPercent` as a percentage in the inclusive range 0 through 100.
- Multiply the capacity by the full fractional percentage.
- Convert only the final result to `int`; conversion truncates toward zero.
- Return zero when either input is zero.
- Preserve fractional percentages such as `99.5`.

## Constraints

- Inputs satisfy `itemsPerHour >= 0`.
- Use only the standard library.
