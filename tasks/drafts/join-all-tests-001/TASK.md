# Test joining independently owned worker callbacks

## Context

Write tests for the specified contract.

## API

```go
func JoinAll(work []func()) <-chan struct{}
```

## Requirements

- Start each non-nil callback exactly once in its own goroutine. Return a non-nil completion channel promptly, even while callbacks are blocked.
- Close completion only after every callback has returned. Send no values on completion. Empty input completes without invoking anything.
- The provided reference uses WaitGroup Add before launch, Done at callback exit and Wait before close. Your tests must expose skipped, duplicate, serial or premature completion.
- Callbacks may write only their own test-owned slots; read those after testing/synctest.Wait or completion synchronization. Keep independent invocations separate and release every test-owned gate, including on failure.

## Constraints

- Use only the Go standard library.
