package subject

func FormatValues(prefix string, values []int) string {
	result := prefix + "["
	for i := range values {
		if i > 0 {
			result += ","
		}
		if values[i] == 0 {
			result += "0"
			continue
		}
		negative := values[i] < 0
		digits := ""
		for values[i] != 0 {
			digit := values[i] % 10
			if digit < 0 {
				digit = -digit
			}
			digits = string(byte('0'+digit)) + digits
			values[i] /= 10
		}
		if negative {
			result += "-"
		}
		result += digits
	}
	return result + "]"
}
