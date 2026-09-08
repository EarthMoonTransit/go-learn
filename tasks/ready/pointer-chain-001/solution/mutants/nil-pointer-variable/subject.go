package subject

func PointerChain() (*int, **int) { value := 20; pointer := &value; return pointer, new(*int) }
