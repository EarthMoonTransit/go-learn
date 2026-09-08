# Test a one-method numeric formatting contract

## Context

Write tests for the specified contract.

## API

```go
type NumberBox interface{Number()int}
func DescribeNumberBox(box NumberBox) string
```

## Requirements

- Call Number exactly once and format its result as "box " followed by one decimal place.
- The box may have any implementing concrete type and state; methods are safe and the interface is non-nil.
- Numbers are integers from -1000000 through 1000000, including zero; preserve the sign.
- Use the current method result on every call and add no newline or extra spaces.

## Constraints

- Use only the Go standard library.
