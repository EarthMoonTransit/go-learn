# Test three shapes through a shared area contract

## Context

Write tests for the specified contract.

## API

```go
type Shape interface{Area()float64}
type Rectangle struct{Width,Height float64}
type Circle struct{Radius float64}
type Triangle struct{Base,Height float64}
func (r Rectangle) Area() float64
func (c Circle) Area() float64
func (t Triangle) Area() float64
```

## Requirements

- Rectangle area is Width times Height; Circle area is math.Pi times Radius squared; Triangle area is Base times Height divided by two.
- Each value implements Shape without a declaration of implementation. Exercise heterogeneous Shape values in your tests.
- Dimensions are finite from 0 through 1000; allow absolute error 0.000000001 and reject NaN.
- Zero dimensions and fractional dimensions are valid; results do not depend on previous calls.

## Constraints

- Use only the Go standard library.
