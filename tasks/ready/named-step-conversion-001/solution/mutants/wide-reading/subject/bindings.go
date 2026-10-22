package subject

type reading int32
type step int16

func shift(value reading, delta step) reading { return value + reading(delta) }
