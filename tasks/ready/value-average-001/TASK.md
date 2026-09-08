# Implement a value method for a fractional average

## Context

Complete the specified operations.

## API

```go
type Sample struct { Label string; Attempts, Hits int }
func (s Sample) Average() float64
func MethodFunctions() (func(Sample) float64, func(*Sample) float64)
```

## Requirements

- Average is Hits divided by Attempts using floating-point division; zero attempts returns zero.
- Inputs satisfy 0 <= Hits <= Attempts <= 1000000. Label is unrelated to the result.
- MethodFunctions returns non-nil functions corresponding to Sample.Average and (*Sample).Average method expressions. Each computes Average from its current argument; calls are independent.
- Do not change any field of the receiver. Absolute error up to 0.000000001 is allowed.

## Constraints

- Use only the Go standard library.
