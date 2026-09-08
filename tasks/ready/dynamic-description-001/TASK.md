# Repair description dispatch by dynamic type

## Context

Repair the supplied operation.

## API

```go
type NumberBox interface{Number()int}
type FancyNumberBox interface{Value()string}
type FancyNumber struct{Text string}
func(n FancyNumber)Value()string{return n.Text}
func Describe(v any) string
```

## Requirements

- Dynamic int and finite float64 values format as "number " plus one decimal place using fmt formatting.
- NumberBox formats "box " plus its Number result with one decimal place.
- FancyNumberBox formats "fancy " plus one decimal place: parse Text only for exact FancyNumber values; otherwise use zero. Invalid or overflowing integers use zero.
- NumberBox takes precedence when both interfaces match. Nil and unsupported types return "unsupported".
- Interface method receivers are non-nil. Do not call Value on unsupported FancyNumberBox implementations.

## Constraints

- Use only the Go standard library.
