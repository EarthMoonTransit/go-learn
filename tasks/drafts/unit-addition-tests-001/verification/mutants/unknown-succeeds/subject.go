package subject

func AddUnit(values, units map[string]int, key, unit string) bool {
	amount, ok := units[unit]
	if !ok {
		return true
	}
	values[key] += amount
	return true
}
