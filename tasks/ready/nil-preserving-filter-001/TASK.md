# Preserve slice presence while filtering

## Context

Filter values while preserving whether the input slice is present.

## API

```go
func KeepNonnegative(values []int) []int
```

## Requirements

- Return only values greater than or equal to zero, in input order, preserving duplicates.
- For a nil input, return nil.
- For every non-nil input, return a non-nil slice, even when it is empty or no values qualify.
- Do not change any input element.

## Constraints

- Capacity and backing-array sharing of the result are unspecified. Do not test or depend on them. Use only the standard library; no goroutines are needed.
