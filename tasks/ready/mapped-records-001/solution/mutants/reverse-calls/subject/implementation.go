package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i := len(values) - 1; i >= 0; i-- {
		value := values[i]
		out[i] = Entry[A, B]{Input: value, Output: f(value)}
	}
	return out
}
