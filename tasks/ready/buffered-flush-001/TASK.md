# Write buffered text and report flush failures

## Context

Complete the specified operations.

## API

```go
func WriteBuffered(writer *bufio.Writer, text string) error
```

## Requirements

- Append all bytes of text to the supplied non-nil buffered writer and flush before returning success, including any bytes already pending.
- Return the original write or flush error; a short underlying write with nil error becomes io.ErrShortWrite.
- Preserve text bytes exactly, including Unicode and empty text. Empty text still flushes pending bytes; an empty buffer with empty text performs no underlying write.
- Preserve pending byte order and leave ownership with the caller. The underlying writer follows io contracts; callers must not rely on a particular number or size of nonempty write calls.

## Constraints

- Use only the Go standard library.
