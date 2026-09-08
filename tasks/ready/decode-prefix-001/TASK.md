# Distinguish malformed bytes from a replacement rune

## Context

Complete the specified Unicode operation.

## API

```go
func DecodePrefix(s string) (rune, int, bool)
```

## Requirements

- Decode only the first UTF-8 unit; later bytes do not affect validity of that unit.
- For a valid leading scalar return the rune, its consumed byte count and true.
- For malformed leading bytes return U+FFFD, one and false.
- For empty input return zero, zero and false; encoded U+FFFD is valid and consumes three bytes.

## Constraints

- Use only the Go standard library.
