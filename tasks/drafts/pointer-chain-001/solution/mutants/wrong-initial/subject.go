package subject

func PointerChain() (*int, **int) { value := 0; pointer := &value; return pointer, &pointer }
