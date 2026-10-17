# Own a child timeout while preserving the operation result

## Context

Wrap a cooperative operation with an owned timeout budget.

## API

```go
func RunWithTimeout(parent context.Context, budget time.Duration,
    operation func(context.Context) (int, error)) (int, error)
```

## Requirements

- Derive a child with budget; an earlier parent deadline wins, and parent cancellation and values propagate.
- Invoke operation exactly once with that child, including an already-canceled parent.
- Return its integer and error unchanged, including a nonzero result with an error.
- Zero or negative budget makes the child expired before invocation.
- Release the child on return without canceling the parent.
- Wait for operation to return, including cleanup after Done; do not return while it remains in background.

## Constraints

- parent and operation are non-nil. Operation cooperates with Done and returns normally after any supplied cleanup gate is released. Cancellation does not abort arbitrary code. Preserve the API; use only the standard library.
