# Own a finite producer's output

## Context

Implement the specified operation.

## API

```go
func Publish(values []int, ready <-chan struct{}) <-chan int
```

## Requirements

- Return a fresh non-nil output channel without waiting for ready.
- One producer waits for one receive from ready, either a value or closure, before sending output.
- Send every input value exactly once in order, including zero and negative values, then close the producer-owned output.
- Preserve input values. Concurrent calls own independent outputs and workers.

## Constraints

- Ready is non-nil and caller-owned. The caller eventually releases it, drains output and preserves values until output closes. Inputs contain at most 32 values. Cancellation and early consumer exit are outside this contract.

Use only the standard library. Review must identify the sole sender/closer and why output closes after its final send. The receive-only gate type is supplied protection; creating and closing your output is the assessed ownership action.
