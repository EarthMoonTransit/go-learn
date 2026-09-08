# Delete a present key while preserving other entries

## Context

Complete the specified operations.

## API

```go
type Dictionary map[string]string
type DictionaryError string
func(e DictionaryError)Error()string{return string(e)}
const ErrAbsent=DictionaryError("cannot delete absent key")
func (d Dictionary) Delete(key string) error
```

## Requirements

- Delete a present key and return nil, including a key with an empty stored value.
- For an absent key return ErrAbsent and preserve the entire dictionary; nil maps contain no keys.
- Keys are exact strings and may be empty. Preserve all unrelated entries.
- The error must be the supplied DictionaryError constant, not a fresh generic error.

## Constraints

- Use only the Go standard library.
