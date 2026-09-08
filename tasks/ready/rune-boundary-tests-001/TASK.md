# Test rune positions at byte boundaries

## Context

Write tests for the specified Unicode contract.

## API

```go
func RuneBoundary(s string, pos int) (int, bool)
```

## Requirements

- Input is valid UTF-8; pos is a zero-based rune position.
- Return the byte offset at that position and true.
- The position immediately after the final rune is valid and returns len(s).
- Negative positions and positions beyond the rune count return zero and false; position zero in empty input is valid.

## Constraints

- Use only the Go standard library.
