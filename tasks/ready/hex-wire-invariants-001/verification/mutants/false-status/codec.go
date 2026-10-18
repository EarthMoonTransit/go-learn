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
	return encodeValue(input)
}

func Decode(wire []byte) ([]byte, bool) {
	value, _ := decodeValue(wire)
	return value, false
}
