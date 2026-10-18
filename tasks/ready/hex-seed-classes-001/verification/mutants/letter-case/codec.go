package subject

import "encoding/hex"

func encodeValue(input []byte) []byte {
	return hex.AppendEncode(nil, input)
}

func decodeValue(wire []byte) ([]byte, bool) {
	for _, c := range wire {
		if c >= 'A' && c <= 'F' {
			return nil, false
		}
	}
	value, err := hex.DecodeString(string(wire))
	return value, err == nil
}

func Encode(input []byte) []byte {
	wire := encodeValue(input)
	for i, c := range wire {
		if c >= 'a' && c <= 'f' {
			wire[i] = c - ('a' - 'A')
		}
	}
	return wire
}

func Decode(wire []byte) ([]byte, bool) {
	return decodeValue(wire)
}
