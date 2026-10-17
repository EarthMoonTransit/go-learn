package subject

func mapEntries[A any, B interface{ ~string | ~[]byte }](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		out[i] = Entry[A, B]{Input: value, Output: f(value)}
	}
	return out
}
