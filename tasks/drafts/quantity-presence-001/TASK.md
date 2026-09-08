# Return quantity and map-key presence

## Context

Complete the specified operations.

## API

```go
func Quantity(values map[string]int, key string) (int, bool)
```

## Requirements

- For a present key return its stored integer and true.
- A present zero is different from a missing key.
- For a missing key, including reads from a nil map, return zero and false.
- Use exact string keys, including the empty string, and leave all entries unchanged.

## Constraints

- Use only the Go standard library.
