# Implement a reusable generic LIFO stack

## Context

Complete the specified operations.

## API

```go
type Stack[T any] struct{values []T}
func NewStack[T any]() *Stack[T]
func (s *Stack[T]) Push(value T)
func (s *Stack[T]) IsEmpty() bool
func (s *Stack[T]) Pop() (T, bool)
```

## Requirements

- NewStack returns an empty non-nil stack. The zero value of Stack also works; receivers are non-nil.
- Push stores one value; Pop removes and returns the most recently pushed value with true.
- An empty Pop returns the zero value of T and false without panicking. IsEmpty reflects the current count.
- Support arbitrary types including slices, repeated empty pops and reuse after draining. Separate stack instances are independent.

## Constraints

- Use only the Go standard library.
