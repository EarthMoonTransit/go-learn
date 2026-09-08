# Return independently allocated integer counters

## Context

Complete the specified operations.

## API

```go
func NewCounter(initial int) *int
```

## Requirements

- Return a non-nil pointer to an integer containing initial.
- initial may be zero, positive or negative.
- The returned integer remains readable and writable after the function returns.
- Each call creates independent storage; updating one result must not change another.

## Constraints

- Use only the Go standard library.
