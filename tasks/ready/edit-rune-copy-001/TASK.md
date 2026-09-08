# Change a rune through a copied representation

## Context

Complete the specified Unicode operation.

## API

```go
func EditRune(s string, index int, replacement rune) (string, bool)
```

## Requirements

- Input is valid UTF-8 and replacement is a valid scalar.
- Treat index as a zero-based rune index and replace exactly that rune.
- Return the resulting UTF-8 string and true, preserving all other code points.
- For negative or out-of-range index, return the original string and false.

## Constraints

- Use only the Go standard library.
