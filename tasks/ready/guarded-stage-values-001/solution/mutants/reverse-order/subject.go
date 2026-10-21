package subject

func RunStages(initial int, steps []func(int) (int, bool)) (int, int) {
	value, count := initial, 0
	for i := len(steps) - 1; i >= 0; i-- {
		step := steps[i]
		if step == nil {
			continue
		}
		next, ok := step(value)
		if !ok {
			break
		}
		value = next
		count++
	}
	return value, count
}
