package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, any] {
	out := make([]Entry[A, any], len(values))
	for i, value := range values {
		out[i] = Entry[A, any]{Input: value, Output: f(value)}
	}
	return out
}
