# Test complete UTF-8 validity

## Context

Write tests for the specified Unicode contract.

## API

```go
func ValidUTF8(s string) bool
```

## Requirements

- Write tests for whether the entire byte string is valid UTF-8.
- Accept empty input, ASCII, NUL and every correctly encoded scalar, including U+FFFD.
- Reject isolated continuation bytes, invalid leading bytes, overlong encodings and truncated sequences.
- Reject encodings of surrogate values or values above U+10FFFF, including malformed bytes after a valid prefix.

## Constraints

- Use only the Go standard library.
