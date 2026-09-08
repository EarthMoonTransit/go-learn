# Build two exact border lines around text

## Context

Complete the specified operations.

## API

```go
func AddBorder(text string, width int) string
```

## Requirements

- Place a line of width asterisks before the text and an identical line after it.
- Separate the three parts with exactly two added newline bytes; add no final newline.
- Preserve the complete text, including any existing newlines or whitespace.
- width is between 0 and 10000; zero width still adds both separators and empty text is valid.

## Constraints

- Use only the Go standard library.
