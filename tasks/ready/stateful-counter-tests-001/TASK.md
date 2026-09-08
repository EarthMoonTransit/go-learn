# Test independent closures with persistent state

## Context

Write tests for the specified contract.

## API

```go
func NewCounter(start int) func(int) int
```

## Requirements

- Return a non-nil closure whose initial state is start.
- Each call adds delta to that closure state and returns the updated value.
- Closures returned by separate NewCounter calls have independent state, even when starts are equal.
- Start and accumulated results fit int; negative and zero deltas are valid. Calls are sequential.

## Constraints

- Use only the Go standard library.
