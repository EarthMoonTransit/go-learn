# Return a stable error for an absent dictionary key

## Context

Complete the specified operations.

## API

```go
type Dictionary map[string]string
var ErrMissing=errors.New("missing key")
func (d Dictionary) Search(key string) (string, error)
```

## Requirements

- Return the exact stored value and nil for a present key, including a present empty value.
- For an absent key return empty text and the supplied ErrMissing sentinel.
- Nil and empty dictionaries have no keys; keys are exact strings and may be empty.
- Preserve all dictionary entries and ErrMissing. A newly created error with the same text does not equal the supplied sentinel.

## Constraints

- Use only the Go standard library.
