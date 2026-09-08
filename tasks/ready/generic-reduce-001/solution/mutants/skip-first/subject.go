package subject

func Reduce[A, B any](items []A, f func(B, A) B, initial B) B {
	result := initial
	for _, v := range items[1:] {
		result = f(result, v)
	}
	return result
}
