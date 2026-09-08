# Build a label with Unicode uppercase conversion

## Context

Complete the specified operations.

## API

```go
func UpperLabel(text string) string
```

## Requirements

- Return the exact prefix INPUT: followed by one space and the uppercased input.
- Use Go Unicode uppercase behavior for every letter.
- Preserve punctuation, digits and whitespace from the input.
- Input is valid UTF-8; empty input returns only the prefix and its space.

## Constraints

- Use only the Go standard library.
