package subject

func ShiftByte(value uint8, count uint) (uint8, int8, uint8, int8) {
	count %= 8
	return value << count, int8(value) << count, value >> count, int8(value) >> count
}
