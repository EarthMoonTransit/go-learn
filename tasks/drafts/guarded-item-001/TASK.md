# Repair slice index guards

## Context

Repair the supplied operation.

## API

```go
func GetItem(values []int, index int) int
```

## Requirements

- For a valid zero-based index, return the exact stored integer.
- The first and final indexes are valid.
- For a negative index or an index at or beyond len(values), return -1.
- Nil and empty slices have no valid index; leave the input unchanged.

## Constraints

- Use only the Go standard library.
