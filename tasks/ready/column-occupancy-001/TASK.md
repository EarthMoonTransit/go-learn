# Count true values in a selected map column

## Context

Complete the specified operations.

## API

```go
func CountColumn(grid map[string][]bool, key string) int
```

## Requirements

- Return the number of true elements in the column stored under key.
- False elements contribute nothing; count both first and final elements.
- Missing keys, nil maps and empty columns return zero.
- Keys are exact strings and columns may have any length; preserve the map and its column elements.

## Constraints

- Use only the Go standard library.
