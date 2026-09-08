# Classify finite values infinities and NaN

## Context

Complete the specified operations.

## API

```go
func FloatKind(value float64) string
```

## Requirements

- Return "nan" for any NaN value, "positive-infinity" for positive infinity and "negative-infinity" for negative infinity.
- Return "finite" for every other value, including signed zero, subnormal values and the largest finite float64.
- Classify the original float64 without narrowing it to float32 or comparing to an arbitrary magnitude threshold.
- NaN is unequal to itself; use an appropriate classification operation.

## Constraints

- Use only the Go standard library.
