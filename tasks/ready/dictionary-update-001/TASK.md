# Repair typed errors and existing-key updates

## Context

Repair the supplied operation.

## API

```go
type Dictionary map[string]string
type DictionaryError string
func(e DictionaryError)Error()string{return string(e)}
const ErrAbsent=DictionaryError("cannot update absent key")
func (d Dictionary) Update(key, value string) error
```

## Requirements

- For an existing key replace its value exactly and return nil; a stored empty value counts as existing.
- For an absent key return ErrAbsent, preserving the entire dictionary; nil maps contain no keys.
- Empty keys and new empty values are valid. Preserve unrelated entries and do not normalize key spelling.
- The failure has dynamic type DictionaryError and Error text "cannot update absent key"; a newly allocated generic error is insufficient.

## Constraints

- Use only the Go standard library.
