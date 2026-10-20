package subject_test

import (
	subject "task.local/replace-mask-field-001/subject"
	"testing"
)

func TestReplacementBitByBit(t *testing.T) {
	for b := 0; b < 256; b += 17 {
		for r := 0; r < 256; r += 15 {
			for m := 0; m < 256; m++ {
				var want uint8
				for bit := uint(0); bit < 8; bit++ {
					value := b
					if m&(1<<bit) != 0 {
						value = r
					}
					if value&(1<<bit) != 0 {
						want += 1 << bit
					}
				}
				if got := subject.ReplaceBits(uint8(b), uint8(r), uint8(m)); got != want {
					t.Fatalf("base=%d replacement=%d mask=%d: got %d; want %d", b, r, m, got, want)
				}
			}
		}
	}
}
