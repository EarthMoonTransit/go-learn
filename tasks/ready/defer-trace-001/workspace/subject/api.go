package subject

const Checkpoint = 1000001

func Trace(seed int) []int {
	var values []int
	record := func(value int) { values = append(values, value) }
	checkpoint := func() { values = append(values, Checkpoint) }
	runTrace(seed, record, checkpoint)
	return values
}
