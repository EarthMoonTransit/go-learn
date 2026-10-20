package subject_test

import (
	subject "task.local/mask-state-tests-001/subject"
	"testing"
)

func TestZeroMaskAny(t *testing.T) {
	if subject.Observe(0, 0).Any {
		t.Fatal("zero mask: Any=true; want false")
	}
}
