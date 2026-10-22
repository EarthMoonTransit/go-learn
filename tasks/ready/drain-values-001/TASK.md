# Drain until the producer closes

## Context

Drain a finite caller-owned stream.

## API

```go
func Collect(in <-chan int) []int
```

## Requirements

- Receive every value until the producer closes input; temporary emptiness is not completion.
- Preserve order, zeros, negative values and duplicates.
- Return a non-nil empty slice when the closed input contained no values.
- Each call returns only its own input's values.
- Use range over the channel. Explain who closes input and why the loop terminates.

## Constraints

- Input is non-nil and caller-owned. Its producer closes it after at most 32 int values. Collect must keep receiving until that closure; there is no early-exit or cancellation request. The producer may pause.

The receive-only signature is supplied protection. Method review checks your actual range loop and the defined producer close path separately from behavioral tests. A correct manual receive loop may pass those tests but does not satisfy the prescribed method.
