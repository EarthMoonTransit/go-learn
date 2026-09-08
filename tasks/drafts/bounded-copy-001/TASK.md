# Copy a bounded byte stream with error reporting

## Context

Complete the specified operations.

## API

```go
var ErrLimit error
func CopyAtMost(dst io.Writer, src io.Reader, limit int64) (int64, error)
```

## Requirements

- Copy at most limit bytes from src to dst. Return bytes written and the original read/write error; EOF before the limit is successful.
- Return 0,ErrLimit for a negative limit, or 0,nil for zero; neither case may read or write.
- A short write with nil error becomes io.ErrShortWrite. When reading returns bytes and an error together, write the bytes before reporting the error.
- Do not close either stream. Streams are non-nil and follow io contracts; no absolute timing or buffer-size requirement.

## Constraints

- Use only the Go standard library.
