# Test an overlapping right shift in fixed storage

## Context

Write tests for the specified contract.

## API

```go
func ShiftRight(values []int, index int) int
```

## Requirements

- Return the original final element, shift elements index..len(values)-2 one place right, then write zero at index.
- Modify the supplied storage so aliases observe the shift; preserve elements before index.
- The slice is nonempty and 0 <= index < len(values). Length is unchanged; signed and repeated values are valid.
- The source and destination overlap. Tests must distinguish a correct shift from repeated overwriting while copying forward.

## Constraints

- Use only the Go standard library.
