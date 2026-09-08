package subject

func Quantity(values map[string]int, key string) (int, bool) { n, _ := values[key]; return n, true }
