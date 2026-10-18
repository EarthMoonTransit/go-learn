package subject

import "encoding/hex"

func Encode(input []byte) []byte {
	return hex.AppendEncode(nil, input)
}

func Decode(wire []byte) ([]byte, bool) {
	value, err := hex.DecodeString(string(wire))
	return value, err == nil
}
