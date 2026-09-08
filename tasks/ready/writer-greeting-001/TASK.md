# Write an exact greeting and report write failures

## Context

Complete the specified operations.

## API

```go
func WriteGreeting(w io.Writer, name string) error
```

## Requirements

- Write exactly "Hello, " followed by name, preserving empty text, Unicode and spaces; append no newline.
- Make one Write call on the supplied non-nil writer. Return its exact error if non-nil, even with a partial count.
- If Write reports fewer bytes than requested with nil error, return io.ErrShortWrite; on a full successful write return nil.
- Do not retry, close the writer or write elsewhere. Writer counts are between zero and the supplied byte length.

## Constraints

- Use only the Go standard library.
