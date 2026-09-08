# Repair scaling without sharing input storage

## Context

Repair the supplied operation.

## API

```go
func ScaledCopy(values []float64, units int) []float64
```

## Requirements

- Input quantities represent two units; multiply each by units divided by two.
- Preserve order and length, including zero and fractional quantities.
- Return independent output storage: changing it later must not change input values.
- Input values are finite between 0 and 10000, units is 0 through 100; preserve input even for units 0 or 2.
- Nil or empty input returns zero values; allow absolute floating error 0.000000001.

## Constraints

- Use only the Go standard library.
