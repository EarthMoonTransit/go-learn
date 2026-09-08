# Repair a numeric type set for named types

## Context

Repair the supplied operation.

## API

```go
// numeric admits int, int64, float64 and types with those underlying types.
func Double[T numeric](value T) T
```

## Requirements

- Double returns value plus itself, preserving its type.
- The private numeric constraint must admit int, int64 and float64 and named types with those underlying types.
- Do not admit strings or other numeric kinds. Use a union of approximated underlying types.
- Inputs are finite and doubling does not overflow. Both negative and zero values are valid.

## Constraints

- Use only the Go standard library.
