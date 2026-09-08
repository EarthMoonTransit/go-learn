# Return zero and literal values with their concrete types

## Context

Complete the specified operations.

## API

```go
func ZeroValues() (any, any, any)
func LiteralValues() (any, any, any, float32)
```

## Requirements

- ZeroValues returns a zero-valued string, int and bool through interfaces, in that order; none is a nil interface.
- LiteralValues returns string "go", int 42, bool true and float32 converted from literal 3.14.
- Preserve concrete dynamic types: for example, float64(42) is not an int even though numerically equal.
- Calls are deterministic and independent; no input or output is read from the environment.

## Constraints

- Use only the Go standard library.
