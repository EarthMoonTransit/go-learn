# Test exact dynamic types and the default branch

## Context

Write tests for the specified contract.

## API

```go
type NamedInt int
func TypeLabel(value any) string
```

## Requirements

- Return nil for a nil interface, int for dynamic type int, string for dynamic type string, and other for every other type. These results are literal strings.
- NamedInt is a distinct named type and returns other. An alias of int remains int.
- A typed nil pointer or slice is not a nil interface and returns other.
- The stored value does not change classification; include zero and empty values.

## Constraints

- Use only the Go standard library.
