# Test a value-based struct transition

## Context

Write tests for the specified contract.

## API

```go
type Motor struct{Speed, Drain, Battery, Distance int}
func Advance(m Motor) Motor
```

## Requirements

- If Battery is at least Drain, subtract Drain and add Speed to Distance.
- Exactly enough battery permits one advance; insufficient battery leaves every field unchanged.
- Speed and Drain never change. Return the updated value without changing the caller value.
- All fields are nonnegative, Drain is positive, and arithmetic fits int.

## Constraints

- Use only the Go standard library.
