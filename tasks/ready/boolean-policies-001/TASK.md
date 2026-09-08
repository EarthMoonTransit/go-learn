# Implement four boolean policies

## Context

Complete the specified operations.

## API

```go
func FastAllowed(blocked bool) bool
func AnyActive(a, b, c bool) bool
func SignalAllowed(blocked, ready bool) bool
func ReleaseAllowed(a, b, ready, override bool) bool
```

## Requirements

- FastAllowed returns the negation of blocked.
- AnyActive returns true when at least one of its three inputs is true.
- SignalAllowed requires ready to be true and blocked to be false.
- ReleaseAllowed requires b to be false and either override to be true or both a false and ready true.

## Constraints

- Use only the Go standard library.
