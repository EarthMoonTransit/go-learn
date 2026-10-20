# Repair fixed-width signed and unsigned shifts

## Context

Repair the function to return unsigned left, signed left, unsigned right and signed right shifts, in that order.

## API

```go
func ShiftByte(value uint8, count uint) (uint8, int8, uint8, int8)
```

## Requirements

- Left-shift value as uint8; discard bits outside that type's width.
- Reinterpret value as int8, then left-shift it as int8.
- Right-shift value as uint8, filling high bits with zero.
- Reinterpret value as int8, then right-shift it, extending the sign.
- Match the results for the exact count, including zero, eight and larger values. Reducing the count modulo eight changes those results.

## Constraints

- Use only the standard library. Count is unsigned: negative counts and constant-expression representability are outside this task. Preserve the signature and do not panic.
