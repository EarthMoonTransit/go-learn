package subject

type reading = int16
type step int16

func shift(value reading, delta step) reading { return value + reading(delta) }
