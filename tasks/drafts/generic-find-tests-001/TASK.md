# Test a generic first matching value

## Context

Write tests for the specified contract.

## API

```go
func Find[T any](items []T, predicate func(T) bool) (T, bool)
```

## Requirements

- Return the first matching value and true, or the zero value of T and false when no value matches.
- Call the non-nil predicate once per visited element in input order, stopping at the first match; never call it for empty input.
- Preserve the input slice. The function works with any element type, including non-comparable structs containing slices.
- Use inferred and explicit type arguments in tests; a matched zero value still has found=true.

## Constraints

- Use only the Go standard library.
