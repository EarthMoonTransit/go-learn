package subject

func RunStages(initial int, steps []func(int) (int, bool)) (int, int) {
	value, count := initial, 0
	for _, step := range steps {
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
	for i := range steps {
		steps[i] = nil
	}
	return value, count
}
