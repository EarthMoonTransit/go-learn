package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		out[len(values)-1-i] = Entry[A, B]{Input: value, Output: f(value)}
	}
	return out
}
