package subject_test

import (
	"task.local/pointer-chain-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	p, pp := subject.PointerChain()
	if p == nil || pp == nil || *p != 20 || *pp != p {
		t.Error("fresh call invalid")
	}
}
