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
	for i := 0; i+1 < len(wire); i += 2 {
		wire[i], wire[i+1] = wire[i+1], wire[i]
	}
	return wire
}

func Decode(wire []byte) ([]byte, bool) {
	normal := append([]byte(nil), wire...)
	for i := 0; i+1 < len(normal); i += 2 {
		normal[i], normal[i+1] = normal[i+1], normal[i]
	}
	return decodeValue(normal)
}
