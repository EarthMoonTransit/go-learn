# Repair a transfer between shared integer values

## Context

Repair the supplied operation.

## API

```go
func Transfer(dst, src *int, limit int) int
```

## Requirements

- If either pointer is nil, return zero and change no value.
- If dst and src point to the same integer, return zero without changing it. Equal values at different addresses are separate integers.
- A zero or negative limit returns zero without changes.
- Otherwise transfer min(limit, *src): subtract that amount from source, add it to destination, and return it.
- Caller aliases observe both changes. Repeated calls use the current values; an empty source transfers zero.

## Constraints

- Initial pointee values are nonnegative and all sums fit int. Use only the standard library; equivalent implementations are accepted.
