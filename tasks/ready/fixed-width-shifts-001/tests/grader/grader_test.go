package subject_test

import (
	subject "task.local/fixed-width-shifts-001/subject"
	"testing"
)

func TestAllByteValuesAndCounts(t *testing.T) {
	for v := 0; v < 256; v++ {
		for n := uint(0); n < 18; n++ {
			left := v
			ur := v
			sr := v
			if sr >= 128 {
				sr -= 256
			}
			for j := uint(0); j < n; j++ {
				left = left * 2 % 256
				ur /= 2
				if sr < 0 {
					sr = (sr - 1) / 2
				} else {
					sr /= 2
				}
			}
			sl := left
			if sl >= 128 {
				sl -= 256
			}
			a, b, c, d := subject.ShiftByte(uint8(v), n)
			if int(a) != left || int(b) != sl || int(c) != ur || int(d) != sr {
				t.Fatalf("value=%d count=%d: got %d %d %d %d, want %d %d %d %d", v, n, a, b, c, d, left, sl, ur, sr)
			}
		}
	}
}
