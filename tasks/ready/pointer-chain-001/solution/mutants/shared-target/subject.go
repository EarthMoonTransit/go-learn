package subject

var value = 20

func PointerChain() (*int, **int) { pointer := &value; return pointer, &pointer }
