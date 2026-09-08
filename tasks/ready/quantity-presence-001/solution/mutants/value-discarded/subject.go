package subject

func Quantity(values map[string]int, key string) (int, bool) { _, ok := values[key]; return 0, ok }
