package subject

func RunStages(initial int, steps []func(int) (int, bool)) (int, int) {
	value, count := initial, 0
	for _, step := range steps {
		if step == nil {
			continue
		}
		next, ok := step(value)
		if !ok {
			count++
			break
		}
		value = next
		count++
	}
	return value, count
}
