package subject_test

import (
	subject "task.local/fixed-width-shifts-001/subject"
	"testing"
)

func TestShiftByte(t *testing.T) {
	cases := []struct {
		v  uint8
		n  uint
		ul uint8
		sl int8
		ur uint8
		sr int8
	}{{0, 0, 0, 0, 0, 0}, {255, 0, 255, -1, 255, -1}, {129, 1, 2, 2, 64, -64}, {128, 7, 0, 0, 1, -1}, {1, 7, 128, -128, 0, 0}, {255, 7, 128, -128, 1, -1}, {1, 8, 0, 0, 0, 0}, {129, 8, 0, 0, 0, -1}, {255, 9, 0, 0, 0, -1}, {127, 9, 0, 0, 0, 0}, {37, 2, 148, -108, 9, 9}, {165, 3, 40, 40, 20, -12}, {128, ^uint(0), 0, 0, 0, -1}}
	for _, c := range cases {
		ul, sl, ur, sr := subject.ShiftByte(c.v, c.n)
		if ul != c.ul || sl != c.sl || ur != c.ur || sr != c.sr {
			t.Errorf("ShiftByte(%d,%d)=(%d,%d,%d,%d); want (%d,%d,%d,%d)", c.v, c.n, ul, sl, ur, sr, c.ul, c.sl, c.ur, c.sr)
		}
	}
}
