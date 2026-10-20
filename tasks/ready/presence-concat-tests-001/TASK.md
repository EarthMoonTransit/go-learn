# Test presence across concatenated slices

## Context

Write tests for the supplied API:

## API

```go
func ConcatPresent(parts ...[]int) []int
```

## Requirements

- Return all elements of all parts in argument order, preserving duplicates, zeros and negative values.
- Return nil exactly when every part is nil. Calling with no arguments also returns nil.
- If any part is non-nil, return a non-nil slice, even when all parts have length zero.
- Nil parts may appear before, between or after present parts; they contribute no elements.
- Do not change any input element. Result capacity and backing-array sharing are unspecified.

## Constraints

- Write tests in package subject_test using testing. Correct implementations must pass; realistic errors must fail assertions. Do not modify the implementation or require a particular allocation strategy.
