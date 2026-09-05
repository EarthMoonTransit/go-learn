# Validate and bound a value

## Context

Implement ordered decisions for a possibly invalid interval.

## API

```go
func Bounded(value, low, high int) (int, bool)
```

## Requirements

- When low > high, return (0, false) for every value.
- Otherwise return true as the second result.
- For value below low, return low; for value above high, return high.
- Keep values inside the inclusive interval unchanged.
- An interval with low == high is valid and always returns that endpoint.

## Constraints

- Edit subject/subject.go. Every int input is valid. Keep the signature. Compare endpoints directly; subtraction can overflow at the limits of int.
