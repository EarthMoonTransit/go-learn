package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		boxed := interface{}(value)
		out[i] = Entry[A, B]{Input: boxed.(A), Output: f(boxed.(A))}
	}
	return out
}
