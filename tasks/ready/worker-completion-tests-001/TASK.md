# Test asynchronous start and explicit completion

## Context

Write tests for the specified contract.

## API

```go
func Start(work func()) <-chan struct{}
```

## Requirements

- Start launches non-nil work in a new goroutine and returns a non-nil completion channel without waiting for work to finish.
- Call work exactly once. Close the completion channel only after work returns; send no values on it.
- Calls own independent workers and completion channels. Callbacks return normally after any supplied gate is released.
- Use controlled channels to hold and release work, then await completion. Use the supplied synctest.Test/Wait scaffold to advance runnable goroutines deterministically; no sleeps or wall-clock deadlines.

## Constraints

- Use only the Go standard library.
