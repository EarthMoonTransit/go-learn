# Reduce two integers to their greatest common divisor

## Context

Implement the greatest common divisor for two bounded non-negative integers.

## API

```go
func GCD(a, b int) int
```

## Requirements

- Return the non-negative greatest common divisor.
- Define GCD(0, 0) as zero; with exactly one zero input, return the other.
- Use a condition-only for loop, with % or %= in its body for remainder reduction; omit initialization and post clauses.
- Do not use recursion or perform I/O.
- Return the same result when the inputs are exchanged.

## Constraints

- Inputs range from zero through 1,000,000 inclusive. Edit only subject/implementation.go. GCD(18, 30) returns 6; GCD(7, 4) returns 1. Public checks report an incorrect loop form or direct recursion. Keep the signature; use no external dependencies.
