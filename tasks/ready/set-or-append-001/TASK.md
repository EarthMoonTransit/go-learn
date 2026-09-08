# Replace an element or append a new value

## Context

Complete the specified operations.

## API

```go
func SetItem(values []int, index, value int) []int
```

## Requirements

- For a valid index, replace exactly that element and return the resulting slice.
- Valid-index replacement must modify the input storage so existing views observe the new element.
- For either an invalid negative index or an index at or beyond length, append value to the end.
- Preserve the order and values of other elements; nil and empty input accept an appended value.

## Constraints

- Use only the Go standard library.
