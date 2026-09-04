# Build and query a structured error chain

## Context

Implement a typed contextual error and helpers that inspect its unwrap chain.

## API

```go
type FieldError struct {
    Field string
    Err   error
}

func (e *FieldError) Error() string
func (e *FieldError) Unwrap() error
func WrapField(field string, err error) error
func HasCause(err, target error) bool
func FindField(err error) (*FieldError, bool)
```

## Requirements

- Include the field and cause text in `FieldError.Error`.
- Preserve the cause through `Unwrap`.
- Return nil from `WrapField` when `err` is nil.
- Match targets through the chain with `errors.Is`.
- Extract a wrapped `*FieldError` with `errors.As`.
- Do not classify errors by message equality.

## Constraints

- Preserve error identity and structured field data.
- Use only the standard library.
