package subject

func RemoveQuantity(items, units map[string]int, item, unit string) bool {
	n, ok := items[item]
	amount, _ := units[unit]
	if !ok || amount > n {
		return false
	}
	if amount == n {
		delete(items, item)
	} else {
		items[item] = n - amount
	}
	return true
}
