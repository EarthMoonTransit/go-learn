package subject_test

import (
	"task.local/kind-condition-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.RequiresPermit("car"); got != true {
		t.Errorf("RequiresPermit: got %v, want %v", got, true)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.RequiresPermit("truck"); got != true {
		t.Errorf("RequiresPermit: got %v, want %v", got, true)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.RequiresPermit("bike"); got != false {
		t.Errorf("RequiresPermit: got %v, want %v", got, false)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.RequiresPermit(""); got != false {
		t.Errorf("RequiresPermit: got %v, want %v", got, false)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.RequiresPermit("CAR"); got != false {
		t.Errorf("RequiresPermit: got %v, want %v", got, false)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.RequiresPermit(" car"); got != false {
		t.Errorf("RequiresPermit: got %v, want %v", got, false)
	}
}

func TestPublicCase07(t *testing.T) {
	if got := subject.RequiresPermit("truck "); got != false {
		t.Errorf("RequiresPermit: got %v, want %v", got, false)
	}
}
