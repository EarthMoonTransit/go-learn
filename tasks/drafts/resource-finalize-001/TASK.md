# Refactor resource cleanup and error propagation

## Context

Refactor `Store` so one write attempt is followed by reliable cleanup without losing either failure.

## API

```go
type Resource interface {
    Write([]byte) (int, error)
    Close() error
}

func Store(resource Resource, data []byte) error
```

## Requirements

- Call `Write` exactly once before cleanup.
- Call `Close` exactly once after every write attempt.
- Return write errors with operation context that preserves `errors.Is`.
- Return `io.ErrShortWrite` when a successful write consumes fewer bytes than supplied.
- Return a close error when writing succeeds.
- Preserve both write and close errors when both occur.

## Constraints

- Preserve the declared API.
- Do not recover panics or retry operations.
- Use only the standard library.
