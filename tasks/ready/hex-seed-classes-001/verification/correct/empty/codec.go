package subject

import "encoding/hex"

func Encode(input []byte) []byte {
	wire := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(wire, input)
	return wire
}

func Decode(wire []byte) ([]byte, bool) {
	value := make([]byte, hex.DecodedLen(len(wire)))
	n, err := hex.Decode(value, wire)
	return value[:n], err == nil
}
