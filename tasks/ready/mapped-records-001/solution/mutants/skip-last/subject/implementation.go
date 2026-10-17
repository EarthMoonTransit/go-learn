package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	if len(values) > 0 {
		values = values[:len(values)-1]
	}
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		out[i] = Entry[A, B]{Input: value, Output: f(value)}
	}
	return out
}
