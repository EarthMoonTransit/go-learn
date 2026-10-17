# Save value-bound and pointer-bound method values

## Context

Complete Capture so two saved calls retain different receiver semantics after the original counter changes.

## API

```go
type Counter struct { Value int }
func (c Counter) ReadValue() int
func (c *Counter) ReadPointer() int
func Capture(c *Counter) (func() int, func() int)
```

## Requirements

- Return two non-nil functions in value-bound, pointer-bound order.
- The first keeps the receiver value observed when Capture is called.
- The second follows the current Value of that same pointed-to counter.
- Both supplied methods return Value without modifying the receiver.
- Keep separate captures independent, including repeated captures of one counter.

## Constraints

- c is non-nil; all int field values are valid. Edit subject/subject.go. Preserve the methods. For Value 6, capture both calls, then set Value to 9: the saved calls return 6 and 9 respectively. No I/O or concurrency.
