package subject

func AddUnit(values, units map[string]int, key, unit string) bool {
	amount, ok := units[unit]
	if !ok {
		values[key] = 0
		return false
	}
	values[key] += amount
	return true
}
