package subject

func Observe() (any, any) {
	original := block{2, 4, 6}
	modified := original
	modified[0] = 99
	return original, modified
}
