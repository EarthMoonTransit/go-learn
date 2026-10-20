package subject

type State struct {
	Any, All           bool
	Set, Clear, Toggle uint8
}

func Observe(bits, mask uint8) State {
	return State{Any: bits&mask != 0, All: bits&mask == mask, Set: mask, Clear: bits &^ mask, Toggle: bits ^ mask}
}
