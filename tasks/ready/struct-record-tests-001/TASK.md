# Test construction of a pointer-backed record

## Context

Write tests for the specified contract.

## API

```go
type Entry struct{Name string;Age int;Attributes map[string]string}
func NewEntry(name string, age int, attributes map[string]string) *Entry
```

## Requirements

- Return a non-nil pointer with the supplied Name, Age and Attributes contents.
- Empty name and zero age are preserved; all integer ages are accepted.
- A nil Attributes input remains nil; a non-nil empty map remains non-nil.
- Do not change the supplied map contents. Each call returns a distinct record; map storage may be shared or copied.

## Constraints

- Use only the Go standard library.
