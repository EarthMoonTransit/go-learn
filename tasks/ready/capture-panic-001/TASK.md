# Contain a callback panic at an explicit boundary

## Context

Complete the specified operations.

## API

```go
func Capture(work func() error) (err error, panicked bool)
```

## Requirements

- Call work exactly once in the current goroutine. For a normal return, preserve its exact error and return panicked=false.
- For a panic carrying a string or int, recover at this boundary and return a non-nil error with text "panic: <value>" and panicked=true.
- Callback panic values are limited to strings and ints; empty strings and zero are valid panic values.
- Do not rerun work or leave background work. Separate calls are independent.

## Constraints

- Use only the Go standard library.
