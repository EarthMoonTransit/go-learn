# Take a value from a map without zero-value ambiguity

## Context

Implement a lookup that removes exactly one existing entry after reading it.

## API

```go
func Take(values map[string]int, key string) (value int, ok bool)
```

## Requirements

- Return the stored value and `true` when `key` exists.
- Treat an existing zero value as present.
- Delete the entry only when it exists.
- Return `0, false` for a missing key.
- Return `0, false` for a nil map without panicking.
- Preserve every unrelated entry.

## Constraints

- Mutate the supplied map only as described.
- Use only the standard library.
