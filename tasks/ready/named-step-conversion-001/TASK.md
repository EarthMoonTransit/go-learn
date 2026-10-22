# Convert between distinct named scalar types

## Context

Repair subject/bindings.go. The supplied Observe wrapper constructs reading(raw) and step(delta), invokes shift, and exposes their dynamic types and values.

## API

```go
func Observe(raw, delta int16) (any, any, any, int16)
```

## Requirements

- Define reading and step as distinct types with underlying type int16. Neither is an alias; they are different types.
- Implement shift(value reading, delta step) reading to return their mathematical sum.
- Observe's first two slots retain the input values with types reading and step, respectively.
- Its third slot contains the sum with type reading; the fourth contains the same number as int16.
- Preserve signed values, including negative changes and zero.

## Constraints

- Inputs are between -1000 and 1000. Edit only subject/bindings.go; preserve the supplied wrapper. Equivalent definitions and conversions are accepted.
