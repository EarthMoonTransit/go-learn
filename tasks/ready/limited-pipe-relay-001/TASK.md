# Relay a bounded pipe prefix

## Context
Copy a producer's prefix.

## API
```go
func Relay(dst io.Writer, produce func(io.Writer) error, limit int64) (n int64, producerErr, copyErr error)
```

## Requirements
- Create io.Pipe; run produce once in one worker. Close the writer with its exact callback error.
- In the caller, use io.Copy through a bounded pipe reader into dst; read at most limit without buffering the whole input.
- Return written byte count and exact copy error. Early EOF succeeds; nil-error short writes yield io.ErrShortWrite. Reaching limit succeeds without another read.
- On copy completion, close the reader with copyErr, including nil, releasing unfinished writes.
- Join the producer before returning its exact callback error separately from copyErr.
- Never close dst. Method review checks actual pipe, adapter, closure and join.

## Constraints
- 0≤limit≤16. Producer writes at most 32 bytes, never closes its endpoint or panics, and eventually returns when I/O completes/fails. Callbacks/streams are nonnil; streams obey io contracts.
