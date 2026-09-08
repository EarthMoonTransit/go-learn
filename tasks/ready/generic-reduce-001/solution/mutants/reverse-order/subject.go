package subject

func Reduce[A, B any](items []A, f func(B, A) B, initial B) B {
	result := initial
	for i := len(items) - 1; i >= 0; i-- {
		result = f(result, items[i])
	}
	return result
}
