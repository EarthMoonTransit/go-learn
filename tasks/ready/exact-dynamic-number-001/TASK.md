# Extract a number only from the exact dynamic type

## Context

Complete the specified operations.

## API

```go
type FancyNumber struct{Text string}
func(n FancyNumber)Value()string{return n.Text}
type FancyNumberBox interface{Value()string}
func Extract(box FancyNumberBox) int
```

## Requirements

- Only a dynamic FancyNumber value is accepted; convert its Text with strconv.Atoi and return the integer.
- Return zero for invalid or overflowing integers, nil interfaces, pointers to FancyNumber and all other implementing types.
- Do not trim whitespace or parse fractional text. Leading signs and leading zeroes follow strconv.Atoi.
- Do not call Value on unsupported types; dynamic type identity determines eligibility.

## Constraints

- Use only the Go standard library.
