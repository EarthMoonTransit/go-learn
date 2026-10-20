package subject

func ShiftByte(value uint8, count uint) (uint8, int8, uint8, int8) {
	if count > 7 {
		count = 7
	}
	return value << count, int8(value) << count, value >> count, int8(value) >> count
}
