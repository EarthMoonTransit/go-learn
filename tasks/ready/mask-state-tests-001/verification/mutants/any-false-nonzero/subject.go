package subject

type State struct {
	Any, All           bool
	Set, Clear, Toggle uint8
}

func Observe(bits, mask uint8) State {
	return State{Any: false, All: bits&mask == mask, Set: bits | mask, Clear: bits &^ mask, Toggle: bits ^ mask}
}
