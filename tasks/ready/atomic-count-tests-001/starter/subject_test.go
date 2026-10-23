package subject_test

import (
	subject "task.local/atomic-count-tests-001/subject"
	"testing"
)

func TestOnePositiveUpdate(t *testing.T) {
	counts := map[string]int{"a": 2}
	err := subject.ApplyCounts(counts, []subject.Delta{{Key: "a", Amount: 1}})
	if err != nil || counts["a"] != 3 {
		t.Errorf("single update count=%d error=%v; want3,nil", counts["a"], err)
	}
}
