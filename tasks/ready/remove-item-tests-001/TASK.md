# Test slice removal and invalid indexes

## Context

Write tests for the specified contract.

## API

```go
func RemoveItem(values []int, index int) []int
```

## Requirements

- For a valid index, remove exactly that element and preserve the order of the remaining values.
- Both the first and final element can be removed.
- For a negative index or an index at or beyond len(values), return unchanged values.
- Nil and empty inputs remain empty; input storage may be reused or changed for a valid removal.

## Constraints

- Use only the Go standard library.
