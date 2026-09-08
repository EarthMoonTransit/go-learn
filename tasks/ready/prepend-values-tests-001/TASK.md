# Test a variadic prefix operation

## Context

Write tests for the specified contract.

## API

```go
func PrependItems(values []int, prefix ...int) []int
```

## Requirements

- Write tests for placing every prefix argument before the original slice elements.
- Preserve the order of both the prefix arguments and the original values.
- With no prefix arguments, return values equal to the original input.
- Nil or empty input accepts a prefix; zero total elements may be returned as nil or an empty slice.

## Constraints

- Use only the Go standard library.
