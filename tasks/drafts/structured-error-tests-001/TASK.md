# Test structured validation errors and their pointer form

## Context

Write tests for the specified contract.

## API

```go
type InvalidCountError struct{Count int;Reason string}
func(e *InvalidCountError)Error()string{return fmt.Sprintf("%d items are invalid: %s",e.Count,e.Reason)}
func ValidateCount(n int) error
```

## Requirements

- Positive counts return a nil error.
- Negative counts return a non-nil *InvalidCountError with the original Count and Reason "negative count"; zero uses Reason "empty count".
- Error formats "<Count> items are invalid: <Reason>" using decimal Count and exact Reason, including for manually constructed errors.
- Check the concrete pointer type, structured fields and displayed text; a generic error with identical text is insufficient.

## Constraints

- Use only the Go standard library.
