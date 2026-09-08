# Test insertion without replacing existing entries

## Context

Write tests for the specified contract.

## API

```go
type Dictionary map[string]string
var ErrExists=errors.New("key exists")
func (d Dictionary) Add(key, value string) error
```

## Requirements

- For a missing key store the supplied value and return nil.
- For a present key return the supplied ErrExists sentinel and preserve its previous value.
- A present empty value is still an existing entry; empty keys and new empty values are valid.
- Preserve unrelated entries. The dictionary is non-nil; a fresh error with matching text is not the supplied sentinel.

## Constraints

- Use only the Go standard library.
