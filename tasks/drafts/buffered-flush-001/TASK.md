# Write buffered text and report flush failures

## Context

Complete the specified operations.

## API

```go
func WriteBuffered(dst io.Writer, text string) error
```

## Requirements

- Write all bytes of text to dst using a buffered writer and flush before returning success.
- Return the original write or flush error; a short underlying write with nil error becomes io.ErrShortWrite.
- Preserve text bytes exactly, including Unicode and empty text. Empty text performs no underlying write.
- Do not close dst. The writer follows io contracts; callers must not rely on a particular number or size of nonempty write calls.

## Constraints

- Use only the Go standard library.
