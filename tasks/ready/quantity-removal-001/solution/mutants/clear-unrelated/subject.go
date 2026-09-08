package subject

func RemoveQuantity(items, units map[string]int, item, unit string) bool {
	n, ok := items[item]
	amount, valid := units[unit]
	if !ok || !valid || amount > n {
		return false
	}
	if amount == n {
		delete(items, item)
	} else {
		items[item] = n - amount
	}
	clear(items)
	return true
}
