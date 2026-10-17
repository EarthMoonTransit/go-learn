package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		out[i] = Entry[A, B]{Output: f(value)}
	}
	return out
}
