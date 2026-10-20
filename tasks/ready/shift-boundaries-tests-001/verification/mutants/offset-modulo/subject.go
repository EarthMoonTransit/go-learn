package subject

func Window(word uint16, offset, width uint) uint16 {
	offset %= 16
	shifted := word >> offset
	if width >= 16 {
		return shifted
	}
	return shifted & ((uint16(1) << width) - 1)
}
