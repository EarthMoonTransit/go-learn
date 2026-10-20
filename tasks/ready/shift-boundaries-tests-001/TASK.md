# Test extraction at fixed-width shift boundaries

## Context

Write tests for:

## API

```go
func Window(word uint16, offset, width uint) uint16
```

## Requirements

- Shift word right by offset positions, filling vacated positions with zeros.
- Keep only the lowest width bits of the shifted value.
- Width zero returns zero; widths of 16 or more retain the entire shifted value.
- Offsets of 16 or more return zero. Counts are not reduced modulo 16.
- Every uint count is valid, including its maximum value; the function must not panic.

## Constraints

- Write tests in package subject_test using testing. Cover ordinary extraction and boundaries. Correct implementations must pass and realistic defects must fail assertions. Do not modify the implementation. Negative counts and compile-time constant representability are outside this runtime contract.
