package subject

func SelectValue(initial, next int, useNext bool) int {
	return map[bool]int{false: initial, true: next}[useNext]
}
