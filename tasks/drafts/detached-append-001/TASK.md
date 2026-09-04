# Repair append isolation for a returned slice

## Context

Repair `AppendDetached` so callers can modify either slice without changing the other.

## API

```go
func AppendDetached(base []int, values ...int) []int
```

## Requirements

- Return the elements of `base` followed by every value.
- Do not change `base` while constructing the result.
- The result must not share writable backing storage with `base`.
- Later changes to the result must not affect `base`.
- Later changes to `base` must not affect the result.
- Handle nil, empty, and spare-capacity inputs.

## Constraints

- Preserve the API.
- Use only the standard library.
