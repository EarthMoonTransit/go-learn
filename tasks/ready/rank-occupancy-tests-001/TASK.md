# Test a one-based rank across map columns

## Context

Write tests for the specified contract.

## API

```go
func CountRank(grid map[string][]bool, rank int) int
```

## Requirements

- Each present column contains exactly eight boolean elements; the map may be nil or empty.
- For ranks 1 through 8, count true elements at that one-based rank across all columns.
- Ranks outside 1 through 8 return zero; true elements at other ranks contribute nothing.
- Preserve every column and map entry; iteration order must not affect the count.

## Constraints

- Use only the Go standard library.
