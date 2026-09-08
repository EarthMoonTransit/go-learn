package subject

type Stack[T any] struct{ values []T }

func NewStack[T any]() *Stack[T]  { return new(Stack[T]) }
func (s *Stack[T]) Push(value T)  { s.values = append(s.values, value) }
func (s *Stack[T]) IsEmpty() bool { return len(s.values) == 0 }
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, true
	}
	i := len(s.values) - 1
	v := s.values[i]
	s.values = s.values[:i]
	return v, true
}
