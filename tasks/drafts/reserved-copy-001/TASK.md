# Copy values with an exact capacity contract

## Context

Complete the specified operations.

## API

```go
func ReservedCopy(in []int, extra int) []int
```

## Requirements

- Return an independent copy with length len(in) and capacity len(in)+extra.
- Preserve input values and order; later output changes and appends must not affect input.
- Extra is 0..1000, input length <=1000. Nil and empty input are valid, including positive extra capacity.
- The exact capacity is observable API behavior, not a timing or allocation benchmark.

## Constraints

- Use only the Go standard library.
