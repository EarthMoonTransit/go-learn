# Return a slice window

## Context

Implement Window(input []int, low, high int) ([]int, bool).

## API

```go
func Window(input []int, low, high int) ([]int, bool)
```

## Requirements

- If 0 <= low <= high <= len(input), return input[low:high] and true.
- Otherwise return nil and false without panicking.
- The result length is high-low; its capacity is cap(input)-low.
- Calling the function must not change the caller's slice length, capacity or elements.
- The result shares elements with input; updates through either view must be visible through the other.

## Constraints

- A valid empty window succeeds. Nil input accepts only low=high=0 and returns nil. Keep the signature. A slice value holds a view's position, length and capacity; creating a view does not copy its elements.
