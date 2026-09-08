# Test appended values and an indexed shared subslice

## Context

Write tests for the specified contract.

## API

```go
func Sequences() ([]int, []string, []string)
```

## Requirements

- Return ten integers 0..9 in order, five strings "a".."e", and a view containing strings at indices one and two.
- The view shares storage with the five-string slice; view index zero corresponds to original index one.
- Separate calls return independent number and string storage.
- Check values and lengths; do not require an implementation-dependent capacity for the appended integer slice.

## Constraints

- Use only the Go standard library.
