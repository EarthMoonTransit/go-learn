# Widen before addition

## Context

Repair a sum that loses information before returning an int.

## API

```go
func WideSum(a, b int8) int
```

## Requirements

- Return the mathematical sum of the two signed inputs.
- Preserve negative input values.
- Accept the complete int8 range, from -128 through 127, for both inputs.
- Return results outside int8 range without wrapping or clamping.

## Constraints

- Edit subject/subject.go. Keep the signature. An int holds every possible result here. Converting an already-overflowed sum cannot recover its original value.
