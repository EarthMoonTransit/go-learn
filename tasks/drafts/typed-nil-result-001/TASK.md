# Handle an interface containing a nil pointer

## Context

Complete the specified operations.

## API

```go
type Result struct{Code int}
func(r *Result)CodeValue()int{return r.Code}
type Coder interface{CodeValue()int}
func ReadCode(value Coder) (int, bool)
```

## Requirements

- Return (0,false) for a nil Coder interface or an interface holding a nil *Result pointer; do not call its method.
- For a non-nil *Result return its Code and true, including zero and negative codes.
- For any other implementing type call CodeValue once and return its result with true. Such values have safe method receivers.
- Preserve the pointed Result. An interface holding a typed nil pointer is not itself nil.

## Constraints

- Use only the Go standard library.
