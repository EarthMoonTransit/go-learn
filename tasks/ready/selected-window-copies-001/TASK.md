# Select independently owned windows

## Context

Collect selected slice intervals.

## API

```go
type Window struct{ Start, End int }
func SelectedCopies(input []int, windows []Window) [][]int
```

## Requirements

- Skip empty windows. Copy others in request order, preserving element order and input.
- Outer length/capacity equal the selected count; inner length/capacity equal End-Start. Empty output nilness is unspecified.
- Each result owns independent storage. Result mutations cannot affect input or peers; input mutations cannot affect results, including duplicate or overlapping windows.
- Count selections before collection. Use length preallocation/indexed writes or zero-length capacity preallocation/append, without zero-value prefix elements.
- Allocate from each selected interval only. Capped headers or whole-input clones do not satisfy storage ownership. Explain the lifetime/allocation plan.

## Constraints

- Standard library only. 0<=Start<=End<=len(input), input length≤256, at most8 windows. No GC timing or allocation threshold.
