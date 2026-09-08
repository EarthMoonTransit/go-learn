# Encode one Unicode scalar

## Context

Complete the specified Unicode operation.

## API

```go
func EncodeScalar(r rune) (string, int, bool)
```

## Requirements

- Accept scalar values from U+0000 through U+10FFFF, excluding U+D800 through U+DFFF.
- For a valid scalar, return its UTF-8 string, byte count and true.
- For any other rune value, return an empty string, zero and false.
- NUL and the replacement character U+FFFD are valid scalars.

## Constraints

- Use only the Go standard library.
