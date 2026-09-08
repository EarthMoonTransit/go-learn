# Test replacement of every matching rune

## Context

Write tests for the specified Unicode contract.

## API

```go
func Replace(log string, oldRune, newRune rune) string
```

## Requirements

- Write tests for Replace; inputs are valid UTF-8 and both rune arguments are valid scalars.
- Replace every occurrence of oldRune with newRune, preserving all other code points and their order.
- Replacement may increase or decrease UTF-8 byte length.
- Empty input, absent oldRune and equal rune arguments leave the text unchanged.

## Constraints

- Use only the Go standard library.
