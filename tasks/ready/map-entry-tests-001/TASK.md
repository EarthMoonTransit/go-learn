# Test unordered enumeration of complete map entries

## Context

Write tests for the specified contract.

## API

```go
type Entry struct{Key string;Value int}
func Entries(values map[string]int) []Entry
```

## Requirements

- Return one Entry for every input key with its exact associated value. Output order is unspecified.
- Do not omit zero, negative or empty-key entries; do not duplicate entries.
- Preserve the input map. Changing the returned entries must not change input values.
- Nil or empty maps return zero entries; keys may contain Unicode and spaces.

## Constraints

- Use only the Go standard library.
