# Keep the appended result

## Context

Repair Push(input []int, value int) []int.

## API

```go
func Push(input []int, value int) []int
```

## Requirements

- Return all original elements followed by value.
- The returned length must be len(input)+1.
- Nil and allocated-empty input must both work.
- Existing elements of input must remain unchanged.
- The caller's original slice length must remain unchanged.

## Constraints

- Keep the signature. Inputs contain at most 100 elements; each element and value lies between -1000 and 1000. Spare capacity may or may not exist. No particular capacity growth factor is required. Storage may be reused: this task does not require the result to be independent of input.
