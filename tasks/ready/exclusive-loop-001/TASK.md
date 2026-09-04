# Repair an exclusive loop

## Context

Repair CountEven(limit int) int. It counts the even integers in the half-open interval [0, limit).

## API

```go
func CountEven(limit int) int
```

## Requirements

- Include zero when limit is positive.
- Exclude limit itself.
- Return zero when limit is zero.
- Count every even integer in the interval exactly once.

## Constraints

- Inputs range from 0 through 10000. For example, limit 5 includes 0, 2 and 4, so the result is 3. The starter already contains the parity check and accumulator. The for statement has an initialization, a condition checked before each iteration, and an update. Keep the signature; do not add input/output code.
