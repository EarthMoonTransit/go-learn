# Repair each malformed decoding unit

## Context

Complete the specified Unicode operation.

## API

```go
func RepairUTF8(s string) string
```

## Requirements

- Read arbitrary bytes as UTF-8 units from left to right.
- Preserve every valid scalar, including encoded U+FFFD and NUL.
- Replace each malformed one-byte decoding unit with the UTF-8 encoding of U+FFFD.
- Do not combine adjacent malformed units; empty input returns an empty string.

## Constraints

- Use only the Go standard library.
