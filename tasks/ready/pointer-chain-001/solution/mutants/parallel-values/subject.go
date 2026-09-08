package subject

func PointerChain() (*int, **int) {
	value := 20
	pointer := &value
	other := 20
	second := &other
	return pointer, &second
}
