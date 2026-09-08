# Repair alternating slice updates

## Context

Repair the supplied operation.

## API

```go
func CorrectAlternating(counts []int) []int
```

## Requirements

- Add one to elements at indexes 0, 2, 4 and every subsequent even index.
- Leave odd-indexed values unchanged and preserve element order.
- Apply the correction to the input slice itself and return the corrected values.
- Empty or nil input has zero output elements; increments fit int.

## Constraints

- Use only the Go standard library.
