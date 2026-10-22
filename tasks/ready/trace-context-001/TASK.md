# Carry request trace metadata

## Context
Attach request-scoped metadata while preserving the parent's control signals.

## API
```go
func WithTrace(parent context.Context, id string) context.Context
func Trace(ctx context.Context) (string, bool)
```

## Requirements
- WithTrace returns a child containing the exact trace id; an empty id is present.
- Trace returns ("", false) when no string trace value exists, including a wrong dynamic type.
- Nested trace values shadow ancestors without changing them.
- Preserve parent cancellation, deadline and unrelated values.
- Author a private comparable key type and safe typed accessor. Avoid collisions with unrelated string or foreign keys.
- Explain why this metadata belongs to the request and is not an optional business argument.

## Constraints
- Parent and ctx are non-nil. IDs are arbitrary strings.
- Actual key design and safe assertion require method review; the starter supplies neither mechanism.
