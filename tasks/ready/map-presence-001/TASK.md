# Look up a value or use a fallback

## Context

Implement LookupOr(input map[string]int, key string, fallback int) int.

## API

```go
func LookupOr(input map[string]int, key string, fallback int) int
```

## Requirements

- Return the stored value when key exists, including a stored zero.
- Return fallback only when key does not exist.
- Treat a nil map as having no keys.
- Leave all map entries unchanged.

## Constraints

- Keys may be any strings, including empty strings. Values and fallback range from -1000 through 1000. Keep the signature. A map lookup can return both the value and a boolean reporting key presence. Reading a missing key yields the value type's zero value; that value alone does not establish absence.
