# Forward values with cancellation at both blocking operations

## Context

Implement a forwarding stage whose consumer may stop reading and cancel its context.

## API

```go
func Forward(ctx context.Context, input chan int) <-chan int
```

## Requirements

- Return a non-nil output channel; forward each received value in order, including duplicates and zero.
- Make forwarding progress while input is still open; do not collect until EOF.
- Close output when input ends or ctx is canceled.
- Cancellation must release a stage waiting for input and a stage blocked sending output.
- Preserve caller ownership: do not close input or cancel the parent.
- Do not leave a stage goroutine blocked after termination.

## Constraints

- ctx and input are non-nil. The caller eventually closes input or cancels ctx. An already-winning send may deliver a value during cancellation; no strict cancellation priority is required. Output capacity is unspecified. Preserve the API; use only the standard library.
