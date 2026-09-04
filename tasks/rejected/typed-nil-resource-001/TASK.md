# Detect typed nil values stored in an interface

## Context

Repair `IsNilResource` so it recognizes both a nil interface and nil values held behind the interface.

## API

```go
type Resource interface {
    Close() error
}

func IsNilResource(value Resource) bool
```

## Requirements

- Return `true` for a nil interface.
- Return `true` for typed nil pointers stored in `Resource`.
- Return `true` for other nil-capable dynamic values that implement `Resource`.
- Return `false` for non-nil pointers.
- Return `false` for non-nil value implementations.
- Do not panic for any valid `Resource` implementation.

## Constraints

- Preserve the declared API.
- Use only the standard library.
