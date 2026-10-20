package subject

func ShiftByte(value uint8, count uint) (uint8, int8, uint8, int8) {
	if count == 0 {
		return 0, 0, 0, 0
	}
	return value << count, int8(value) << count, value >> count, int8(value) >> count
}
