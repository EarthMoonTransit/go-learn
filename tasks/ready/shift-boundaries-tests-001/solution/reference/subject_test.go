package subject_test

import (
	subject "task.local/shift-boundaries-tests-001/subject"
	"testing"
)

func TestWindowBoundaries(t *testing.T) {
	max := ^uint(0)
	cases := []struct {
		word          uint16
		offset, width uint
		want          uint16
	}{
		{0xB6D3, 0, 0, 0}, {0xB6D3, 0, 4, 3}, {0xB6D3, 4, 8, 0x6D},
		{0xB6D3, 8, 4, 6}, {0xB6D3, 12, 4, 11}, {0xB6D3, 15, 16, 1},
		{0xB6D3, 1, 16, 0x5B69}, {0xB6D3, 0, 15, 0x36D3},
		{0xB6D3, 0, 16, 0xB6D3}, {0xB6D3, 0, 17, 0xB6D3}, {0xB6D3, 0, max, 0xB6D3},
		{0xB6D3, 16, 16, 0}, {0xB6D3, 32, 8, 0}, {0xB6D3, max, max, 0},
		{0, 0, 16, 0}, {0xFFFF, 8, 16, 255}, {1, 0, 1, 1},
	}
	for _, c := range cases {
		got := subject.Window(c.word, c.offset, c.width)
		if got != c.want {
			t.Errorf("Window(%#04x,%d,%d)=%#04x; want %#04x", c.word, c.offset, c.width, got, c.want)
		}
	}
}
