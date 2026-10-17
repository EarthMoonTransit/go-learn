package subject

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		done := make(chan struct{})
		go func() {
			out[i] = Entry[A, B]{Input: value, Output: f(value)}
			close(done)
		}()
		<-done
	}
	return out
}
