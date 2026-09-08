# Release after an operation even when it panics

## Context

Complete the specified operations.

## API

```go
func WithRelease(work func() int, release func()) int
```

## Requirements

- Call work once and return its integer result unchanged.
- Call release exactly once after work finishes and before WithRelease returns.
- If work panics, still call release and let the original panic continue unchanged. Do not recover it.
- Both callbacks are non-nil; release does not panic. Deferred calls run during normal return and panic unwinding.

## Constraints

- Use only the Go standard library.
