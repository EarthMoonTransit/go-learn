# Repair Unicode byte offsets

## Context

Repair the supplied Unicode operation.

## API

```go
func RuneOffsets(s string) []int
```

## Requirements

- Input is valid UTF-8.
- Return the byte index at which each decoded rune starts, in input order.
- Include ASCII, combining code points and multibyte code points separately.
- Empty input returns a slice with zero elements; either nil or empty is accepted.

## Constraints

- Use only the Go standard library.
