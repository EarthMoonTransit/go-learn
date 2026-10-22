package subject

type reading int16
type step int16

func shift(value reading, delta step) step { return step(value) + delta }
