package subject_test

import (
	"bytes"
	"task.local/hex-seed-classes-001/subject"
	"testing"
)

func FuzzHex(f *testing.F) {
	// Extend the finite corpus; the round-trip check is supplied.
	f.Add([]byte{0x44})
	f.Fuzz(func(t *testing.T, input []byte) {

		wire := subject.Encode(input)
		decoded, ok := subject.Decode(wire)
		if !ok {
			t.Errorf("Decode(Encode(%x)) returned false", input)
		}
		if !bytes.Equal(decoded, input) {
			t.Errorf("decoded bytes = %x, want %x", decoded, input)
		}

	})
}
