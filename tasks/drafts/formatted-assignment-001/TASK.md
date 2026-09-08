# Format a multiline record with width and precision

## Context

Complete the specified operations.

## API

```go
func Welcome(name string) string
func FormatAssignment(name string, slot int, direction string, distance float64, peer string) string
```

## Requirements

- Welcome returns "Welcome, <name>!".
- FormatAssignment returns three lines: Welcome(name), "Slot <slot>: <direction>, <distance> m.", and "Peer: <peer>.".
- Format slot as decimal with minimum width three and leading zeroes; format distance with one decimal place using fmt rounding.
- Preserve all supplied strings exactly, including Unicode, whitespace and empty strings. Separate lines with one newline and add no final newline.
- Slot is 0..9999; distance is finite from 0 through 10000.

## Constraints

- Use only the Go standard library.
