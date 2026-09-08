# Test independent exact-match closures

## Context

Write tests for the specified contract.

## API

```go
type Record struct{Day int;Amount float64;Category string}
func InCategory(category string) func(Record) bool
```

## Requirements

- Return a non-nil predicate that accepts exactly matching Category strings.
- Matching is case-sensitive and whitespace-sensitive; empty categories are valid.
- Day and Amount do not affect the result.
- Predicates from separate calls retain independent categories and can be called repeatedly in any order.

## Constraints

- Use only the Go standard library.
