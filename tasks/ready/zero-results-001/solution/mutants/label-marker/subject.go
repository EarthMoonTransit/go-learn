package subject

func Defaults() (int, bool, string, []int, map[string]int) {
	var count int
	var enabled bool
	var label string = "unset"
	var values []int
	var lookup map[string]int
	return count, enabled, label, values, lookup
}
