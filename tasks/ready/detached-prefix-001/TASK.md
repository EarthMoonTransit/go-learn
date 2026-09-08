# Repair a prefix that retains input storage

## Context

Repair the supplied operation.

## API

```go
func DetachedPrefix(in []int, n int) []int
```

## Requirements

- Return the first n values with length and capacity exactly n; preserve their order.
- The result owns independent storage: mutations through input or result must not affect the other.
- Bounds satisfy 0 <= n <= len(in). Nil input and n==0 are valid, with zero capacity.
- Preserve input during the call; no garbage-collector timing or absolute memory measurement is required.

## Constraints

- Use only the Go standard library.
