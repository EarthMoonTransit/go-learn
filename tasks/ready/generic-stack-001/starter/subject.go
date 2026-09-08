package subject

type Stack[T any] struct{ values []T }

func NewStack[T any]() *Stack[T]   { return nil }
func (s *Stack[T]) Push(value T)   {}
func (s *Stack[T]) IsEmpty() bool  { return false }
func (s *Stack[T]) Pop() (T, bool) { var zero T; return zero, false }
