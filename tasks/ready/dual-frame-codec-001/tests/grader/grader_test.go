package subject_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGraderIndependentVectorsAndAllTruncations(t *testing.T) {
	for _, tag := range []byte{1, 15, 16, 128, 255} {
		for _, n := range []int{1, 2, 7, 31, 32} {
			payload := make([]byte, n)
			for i := range payload {
				payload[i] = byte(i*37) + tag
			}
			binaryWire := append([]byte{1, tag, byte(n)}, payload...)
			textWire := []byte("1/" + hex.EncodeToString([]byte{tag}) + "/" + hex.EncodeToString([]byte{byte(n)}) + "/" + hex.EncodeToString(payload))
			name := fmt.Sprintf("tag-%02x-size-%d", tag, n)
			frameVector(t, name, tag, payload, binaryWire, textWire)
		}
	}
	for i, wire := range [][]byte{{1, 0x82, 2, 0xab, 0xff}, []byte("1/82/02/abff")} {
		for end := 0; end < len(wire); end++ {
			rejectFrame(t, frameFormats[i], fmt.Sprintf("%s-prefix-%d", frameFormats[i].name, end), bytes.Clone(wire[:end]))
		}
	}
}
