package subject_test

import (
	"task.local/receiver-reset-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	e := subject.Entry{Name: "record", Age: 29, Attributes: map[string]string{"street": "Main St."}}
	alias := &e
	alias.Clear()
	if e.Name != "" || e.Age != 0 || e.Attributes != nil {
		t.Errorf("reset got %+v want zero record", e)
	}
}

func TestPublicCase02(t *testing.T) {
	m := map[string]string{"k": "v"}
	e := subject.Entry{Attributes: m}
	e.Clear()
	if len(m) != 1 || m["k"] != "v" {
		t.Errorf("old map changed: %v", m)
	}
}

func TestPublicCase03(t *testing.T) {
	e := subject.Entry{}
	e.Clear()
	e.Clear()
	if e.Name != "" || e.Age != 0 || e.Attributes != nil {
		t.Errorf("repeat reset got %+v", e)
	}
}
