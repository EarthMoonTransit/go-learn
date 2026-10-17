# Refactor resource cleanup and error propagation

## Context

Refactor Store so one write attempt is followed by cleanup without losing failures.

## API

```go
type Resource interface {
    Write([]byte) (int, error)
    Close() error
}
func Store(resource Resource, data []byte) error
```

## Requirements

- Call Write(data) exactly once, including empty data.
- Call Close exactly once afterward, including write failure or panic.
- Write errors have context text containing write, in any case, while preserving errors.Is.
- A short write with nil write error returns io.ErrShortWrite.
- Return the close error when writing succeeds.
- Preserve both the write failure or short-write error and the close error when both occur.

## Constraints

- Resource is non-nil. Write returns a count from 0 through len(data), possibly a short count with nil error; it may panic. Close returns normally. Do not retry operations or recover panics. Preserve the API. Edit subject/subject.go. Use only the standard library.
