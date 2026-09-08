package subject

func ValidUTF8(s string) bool {
	for i := 0; i < len(s); {
		b := s[i]
		if b < 0x80 {
			i++
			continue
		}
		n, min, r := 0, 0, rune(0)
		switch {
		case b >= 0xc0 && b <= 0xdf:
			n = 2
			min = 0x80
			r = rune(b & 0x1f)
		case b >= 0xe0 && b <= 0xef:
			n = 3
			min = 0x800
			r = rune(b & 0x0f)
		case b >= 0xf0 && b <= 0xf7:
			n = 4
			min = 0x10000
			r = rune(b & 7)
		default:
			return false
		}
		if len(s)-i < n {
			return false
		}
		for j := 1; j < n; j++ {
			v := s[i+j]
			if v&0xc0 != 0x80 {
				return false
			}
			r = r<<6 | rune(v&0x3f)
		}
		_ = min
		if r >= 0xd800 && r <= 0xdfff {
			return false
		}
		if r > 0x10ffff {
			return false
		}
		i += n
	}
	return true
}
