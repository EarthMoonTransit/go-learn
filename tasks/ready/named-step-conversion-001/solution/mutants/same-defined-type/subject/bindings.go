package subject

type reading int16
type step = reading

func shift(value reading, delta step) reading { return value + reading(delta) }
