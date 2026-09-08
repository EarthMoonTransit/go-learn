# Sum the tail of each supplied slice

## Context

Complete the specified operations.

## API

```go
func SumTails(groups ...[]int) []int
```

## Requirements

- Return one sum per input slice, in the order supplied.
- Exclude only the first element of each slice; include every remaining integer, including negatives.
- Nil, empty and one-element groups produce zero; no groups produce zero results.
- Preserve all input elements. Every tail sum fits int; nil or empty zero-length output is acceptable.

## Constraints

- Use only the Go standard library.
