# Return a shared view with no spare append capacity

## Context

Complete the specified operations.

## API

```go
func CappedView(in []int, start, end int) []int
```

## Requirements

- Return the elements in in[start:end] with length and capacity both end-start.
- Existing view elements share input storage, so changing one changes the corresponding input element.
- Appending an element to the returned view must allocate independent storage and preserve input elements outside the view.
- Bounds satisfy 0 <= start <= end <= len(in). Empty ranges and nil input are valid; do not mutate input during the call.

## Constraints

- Use only the Go standard library.
