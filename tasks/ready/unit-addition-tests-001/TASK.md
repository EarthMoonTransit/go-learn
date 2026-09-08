# Test map updates using named units

## Context

Write tests for the specified contract.

## API

```go
func AddUnit(values, units map[string]int, key, unit string) bool
```

## Requirements

- values is non-nil; units may be nil. The two maps do not share storage. Quantities are nonnegative and additions fit int.
- An absent unit returns false and leaves values unchanged.
- A present unit returns true and adds its quantity to the named entry, creating an absent entry.
- A present zero unit still succeeds and creates an absent key with zero.
- Preserve other entries and the entire units map; repeated calls accumulate.

## Constraints

- Use only the Go standard library.
