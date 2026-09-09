# Register test cleanup after successful acquisition

## Context

Complete the specified operations.

## API

```go
type Scope interface{Helper();Cleanup(func());Errorf(format string,args ...any)}
func AcquireForTest(t Scope, acquire func() (io.Closer, error)) (io.Closer, error)
```

## Requirements

- Register this helper once, then call acquire exactly once. On error return nil and the same error, registering no cleanup.
- On success acquire provides a non-nil resource. Return that exact resource and register one Cleanup callback; do not close it yet.
- Cleanup closes the resource once. Report a close error once with Errorf("cleanup: %v", err); successful cleanup reports nothing.
- Use Cleanup so the testing scope owns lifetime, including failure paths. Multiple acquisitions retain independent resources; cleanup executes in LIFO registration order.

## Constraints

- Use only the Go standard library.
