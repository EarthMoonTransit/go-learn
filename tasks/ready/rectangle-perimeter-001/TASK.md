# Repair perimeter from independent struct fields

## Context

Repair the supplied operation.

## API

```go
type Rectangle struct{Width,Height float64}
func Perimeter(rect Rectangle) float64
```

## Requirements

- Return twice the sum of Width and Height.
- Each dimension is finite between zero and 1000; zero and fractional dimensions are valid.
- Preserve the input value; Width and Height contribute independently.
- Allow absolute error 0.000000001 and reject NaN.

## Constraints

- Use only the Go standard library.
