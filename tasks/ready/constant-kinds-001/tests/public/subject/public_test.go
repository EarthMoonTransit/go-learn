package subject_test

import (
	"task.local/constant-kinds-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	name, port, ratio := subject.Values()
	if name != "localhost" || port != 8080 || ratio != 1.5 {
		t.Errorf("values got%q,%d,%g wantlocalhost,8080,1.5", name, port, ratio)
	}
}
