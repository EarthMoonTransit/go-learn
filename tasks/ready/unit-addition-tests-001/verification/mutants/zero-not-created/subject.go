package subject

func AddUnit(values, units map[string]int, key, unit string) bool {
	amount, ok := units[unit]
	if !ok {
		return false
	}
	if amount != 0 {
		values[key] += amount
	}
	return true
}
