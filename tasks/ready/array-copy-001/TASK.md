# Replace an array element

## Context

Implement Replaced(input [3]int, index int, value int) [3]int.

## API

```go
func Replaced(input [3]int, index int, value int) [3]int
```

## Requirements

- For indices 0, 1 and 2, return an array with that element replaced.
- Preserve the other two elements.
- For any other index, return the unchanged array without panicking.
- Calling the function must leave the caller's array unchanged.

## Constraints

- Elements and value range from -1000 through 1000. Index may be any int. Arrays have fixed length and are passed as values. The input parameter is a local array value, not a slice. Keep the signature. You may use len(input) to check the upper boundary.
