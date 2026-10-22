package subject_test

import (
	subject "task.local/iota-layout-tests-001/subject"
	"testing"
)

func TestLayoutExpandsEveryPosition(t *testing.T) {
	codes, masks, restart := subject.Layout()
	if codes != [3]int{10, 12, 13} {
		t.Fatalf("codes=%v; want [10 12 13] including blank declaration position", codes)
	}
	if masks != [3]uint8{1, 4, 8} {
		t.Fatalf("masks=%v; want [1 4 8] from repeated shift expressions", masks)
	}
	if restart != [2]int{0, 1} {
		t.Fatalf("second-block values=%v; want [0 1] after resetting iota", restart)
	}
}
