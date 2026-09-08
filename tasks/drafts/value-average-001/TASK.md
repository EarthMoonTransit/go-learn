# Implement a value method for a fractional average

## Context

Complete the specified operations.

## API

```go
type Sample struct { Label string; Attempts, Hits int }
func (s Sample) Average() float64
```

## Requirements

- Average is Hits divided by Attempts using floating-point division; zero attempts returns zero.
- Inputs satisfy 0 <= Hits <= Attempts <= 1000000. Label is unrelated to the result.
- Use the declared value receiver: both Sample values and *Sample pointers must expose Average.
- Do not change any field of the receiver. Absolute error up to 0.000000001 is allowed.

## Constraints

- Use only the Go standard library.
