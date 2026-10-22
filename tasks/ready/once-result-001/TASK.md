# Cache one completed initialization

## Context
Share one lazy initialization result across callers.

## API
```go
func Memoize(load func() (int, error)) func() (int, error)
```

## Requirements
- Construction must not call load.
- The returned function calls load exactly once, on its first invocation.
- Every invocation returns the complete first value and exact error, including a nonzero value with an error or a zero value without one.
- Concurrent callers wait for that initialization to finish. Cache errors too; never retry.
- Separate Memoize instances have independent initialization and results.
- Use sync.Once and captured result state. Explain how completion publishes the result and why reset/retry needs a different API.

## Constraints
- Load is non-nil, eventually returns and does not panic.
- No reset, retry or resource teardown is requested.
- Method review inspects your actual Once and result publication separately from behavioral/race checks.
