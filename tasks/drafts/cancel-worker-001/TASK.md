# Repair derived cancellation and worker shutdown

## Context

Repair the supplied operation.

## API

```go
func StartCancelable(parent context.Context) (context.Context, context.CancelFunc, <-chan struct{})
```

## Requirements

- Derive a cancelable child from the non-nil parent. Return that child, its CancelFunc and a non-nil worker completion channel.
- The worker waits for child cancellation, then closes completion without sending values. Explicit child cancellation and parent cancellation must both finish it.
- Canceling the child must not cancel its parent. Repeated calls to cancel are safe; an already canceled parent produces a canceled child and a finishing worker.
- Preserve context values and parent cancellation cause behavior through standard context.WithCancel derivation. Tests control cancellation and use synctest; no real sleep is required.

## Constraints

- Use only the Go standard library.
