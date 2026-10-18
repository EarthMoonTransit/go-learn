package subject

func Encode(input []byte) []byte {
	const digits = "0123456789abcdef"
	var wire []byte
	for _, value := range input {
		wire = append(wire, digits[value/16], digits[value%16])
	}
	return wire
}

func Decode(wire []byte) ([]byte, bool) {
	if len(wire)%2 != 0 {
		return nil, false
	}
	var value []byte
	for i := 0; i < len(wire); i += 2 {
		hi, okHi := digit(wire[i])
		lo, okLo := digit(wire[i+1])
		if !okHi || !okLo {
			return nil, false
		}
		value = append(value, hi*16+lo)
	}
	return value, true
}

func digit(c byte) (byte, bool) {
	switch {
	case c >= '0' && c <= '9':
		return c - '0', true
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10, true
	default:
		return 0, false
	}
}
