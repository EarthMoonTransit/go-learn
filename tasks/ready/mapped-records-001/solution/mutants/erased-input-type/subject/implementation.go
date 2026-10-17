package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[any, B] {
	out := make([]Entry[any, B], len(values))
	for i, value := range values {
		out[i] = Entry[any, B]{Input: value, Output: f(value)}
	}
	return out
}
