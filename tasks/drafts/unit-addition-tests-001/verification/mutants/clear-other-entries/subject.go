package subject

func AddUnit(values, units map[string]int, key, unit string) bool {
	amount, ok := units[unit]
	if !ok {
		return false
	}
	old := values[key]
	clear(values)
	values[key] = old + amount
	return true
}
