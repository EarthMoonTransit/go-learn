package subject_test

import (
	subject "task.local/replace-mask-field-001/subject"
	"testing"
)

func TestSelectedReplacement(t *testing.T) {
	cases := []struct{ base, replacement, mask, want uint8 }{
		{255, 0, 15, 240}, {128, 1, 1, 129}, {0, 255, 15, 15}, {165, 60, 0, 165},
		{165, 60, 255, 60}, {165, 60, 90, 189}, {195, 3, 3, 195},
	}
	for _, c := range cases {
		got := subject.ReplaceBits(c.base, c.replacement, c.mask)
		if got != c.want {
			t.Errorf("ReplaceBits(%08b,%08b,%08b)=%08b; want %08b", c.base, c.replacement, c.mask, got, c.want)
		}
	}
}
