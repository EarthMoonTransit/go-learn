# Drain buffered values before observing channel closure

## Context

Repair the supplied operation.

## API

```go
func CollectClosed(input <-chan int) []int
```

## Requirements

- Input is non-nil, buffered and already closed. Return all queued values in send order, then stop at closed-and-empty.
- Zero and negative values are ordinary data. Closed empty input returns an empty result; nil or allocated empty slices are both valid.
- Do not close the caller-owned input or invent an extra zero after its buffer is drained. Repeated calls on the drained channel return empty.
- The public fixtures fully prepare channels before calling the function; no scheduling or timing assumptions are needed.

## Constraints

- Use only the Go standard library.
