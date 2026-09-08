package subject

var pointer *int

func PointerChain() (*int, **int) { value := 20; pointer = &value; return pointer, &pointer }
