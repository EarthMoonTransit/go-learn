# Test typed error extraction through a cause tree

## Context

Write tests for the specified contract.

## API

```go
type Detail struct{Code int;Message string}
func(e *Detail)Error()string{return e.Message}
func FindDetail(err error) (*Detail, bool)
```

## Requirements

- Return the first *Detail found by errors.As and true, preserving pointer identity.
- Search wrapped and joined errors in standard traversal order; respect custom As methods.
- Return nil,false when no match exists, including a nil error. Matching a typed-nil *Detail returns nil,true.
- Do not synthesize a new Detail from its fields or message. Error trees are finite.

## Constraints

- Use only the Go standard library.
