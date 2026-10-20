# Gather an array from an index map

## Context

Implement:

## API

```go
func Gather(source [8]int, indexes [8]int) ([8]int, bool)
```

## Requirements

- An index map is valid when every entry is between 0 and 7 inclusive.
- For valid maps, result element i equals source[indexes[i]], and the boolean is true.
- Every lookup uses the original source value. Repeated indexes are valid; the map need not be a permutation.
- If any index is invalid, return the unchanged source and false. Do not return a partially transformed array or panic.
- Preserve both input values. Zero and negative source elements are ordinary data.

## Constraints

- Use only the standard library. Any implementation satisfying this value contract is accepted.
