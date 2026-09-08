package subject_test

import (
	"task.local/receiver-speakers-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.SayHello(&subject.Chinese{}); got != "你好世界" {
		t.Errorf("pointer dispatch got%q", got)
	}
}
