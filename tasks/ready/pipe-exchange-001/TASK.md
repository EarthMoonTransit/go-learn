# Complete a pipe exchange

## Context
Connect two callbacks with a synchronous byte stream and finish both endpoints.

## API
```go
func Exchange(produce func(io.Writer) error, consume func(io.Reader) error) (producerErr, consumerErr error)
```

## Requirements
- Create one io.Pipe; invoke produce once in one worker goroutine and consume once in the calling goroutine.
- Give callbacks its writer and reader; writes wait for reads.
- After produce returns, close the writer with its returned error, allowing the reader to observe EOF or that error.
- After consume returns, close the reader with its returned error, including early nil success, to release unfinished writes.
- Join the producer before Exchange returns.
- Return both exact callback errors, preserving identity even after partial transfer.

## Constraints
- Callbacks are non-nil, do not panic or close endpoints, and eventually return when their I/O completes or fails.
- Actual pipe ownership and close/join placement require method review.
