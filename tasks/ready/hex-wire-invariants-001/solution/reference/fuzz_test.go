package subject_test

import (
	"bytes"
	"encoding/hex"
	"task.local/hex-wire-invariants-001/subject"
	"testing"
)

func FuzzHex(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{0x00})
	f.Add([]byte{0x80})
	f.Add([]byte{0x0d})
	f.Add([]byte{0x12, 0x34, 0x56})
	f.Add([]byte{0x77, 0x77, 0x77, 0x77, 0x77, 0x77, 0x77, 0x77})
	f.Fuzz(func(t *testing.T, input []byte) {

		wire := subject.Encode(input)
		expected := hex.AppendEncode(nil, input)
		if !bytes.Equal(wire, expected) {
			t.Fatalf("wire = %q, want canonical lowercase %q", wire, expected)
		}
		decoded, ok := subject.Decode(wire)
		if !ok {
			t.Errorf("Decode(Encode(%x)) returned false", input)
		}
		if !bytes.Equal(decoded, input) {
			t.Errorf("decoded bytes = %x, want %x", decoded, input)
		}

	})
}
