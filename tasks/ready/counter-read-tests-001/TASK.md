# Test optional pointer reads

## Context

Write tests for the specified contract.

## API

```go
func CounterValue(counter *int) int
```

## Requirements

- A nil pointer returns zero without panicking.
- A non-nil pointer returns the exact integer it points to, including zero and negative values.
- Read the current value on every call, including after the caller changes it.
- Leave the pointed-to integer unchanged.

## Constraints

- Use only the Go standard library.
