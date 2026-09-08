# Propagate the caller context through a composed operation

## Context

Complete the specified operations.

## API

```go
func CallWithContext(ctx context.Context, operation func(context.Context) (int, error)) (int, error)
```

## Requirements

- Call the non-nil operation exactly once with the exact supplied non-nil context. Return its integer and error unchanged, including a nonzero result with an error.
- Pass through cancellation, deadlines and values by preserving context identity. Do not replace, derive or cancel the context.
- Invoke operation even if the context is already canceled; that operation owns its cancellation response.
- Use synchronous composition: return only after the operation returns, and do not launch a background callback.

## Constraints

- Use only the Go standard library.
