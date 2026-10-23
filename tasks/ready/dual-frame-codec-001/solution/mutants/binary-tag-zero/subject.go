package subject

import "errors"

type Frame struct {
	Tag     byte
	Payload []byte
}

var ErrFrame = errors.New("invalid frame")

func writeHexPair(dst []byte, value byte) {
	const digits = "0123456789abcdef"
	dst[0] = digits[value>>4]
	dst[1] = digits[value&15]
}
func readHexPair(src []byte) (byte, bool) {
	if len(src) != 2 {
		return 0, false
	}
	var result byte
	for _, b := range src {
		var digit byte
		switch {
		case b >= '0' && b <= '9':
			digit = b - '0'
		case b >= 'a' && b <= 'f':
			digit = b - 'a' + 10
		default:
			return 0, false
		}
		result = result*16 + digit
	}
	return result, true
}
func (f Frame) MarshalBinary() ([]byte, error) {
	n := len(f.Payload)
	if n > 32 {
		return nil, ErrFrame
	}
	data := make([]byte, 3+n)
	data[0] = 1
	data[1] = 0
	data[2] = byte(n)
	copy(data[3:], f.Payload)
	return data, nil
}
func (f *Frame) UnmarshalBinary(data []byte) error {
	if len(data) < 3 || data[0] != 1 || data[2] > 32 || len(data) != 3+int(data[2]) {
		return ErrFrame
	}
	tag := data[1]
	payload := append([]byte(nil), data[3:]...)
	*f = Frame{Tag: tag, Payload: payload}
	return nil
}
func (f Frame) MarshalText() ([]byte, error) {
	n := len(f.Payload)
	if n > 32 {
		return nil, ErrFrame
	}
	data := make([]byte, 8+2*n)
	data[0] = '1'
	data[1] = '/'
	data[4] = '/'
	data[7] = '/'
	writeHexPair(data[2:4], f.Tag)
	writeHexPair(data[5:7], byte(n))
	for i, b := range f.Payload {
		writeHexPair(data[8+2*i:10+2*i], b)
	}
	return data, nil
}
func (f *Frame) UnmarshalText(data []byte) error {
	if len(data) < 8 || data[0] != '1' || data[1] != '/' || data[4] != '/' || data[7] != '/' {
		return ErrFrame
	}
	tag, ok := readHexPair(data[2:4])
	if !ok {
		return ErrFrame
	}
	n, ok := readHexPair(data[5:7])
	if !ok || n > 32 || len(data) != 8+2*int(n) {
		return ErrFrame
	}
	payload := make([]byte, int(n))
	for i := range payload {
		value, valid := readHexPair(data[8+2*i : 10+2*i])
		if !valid {
			return ErrFrame
		}
		payload[i] = value
	}
	*f = Frame{Tag: tag, Payload: payload}
	return nil
}
