package subject_test

import (
	"maps"
	"task.local/dictionary-delete-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	d := subject.Dictionary{"test": "value", "keep": "x"}
	if e := d.Delete("test"); e != nil || !maps.Equal(d, subject.Dictionary{"keep": "x"}) {
		t.Errorf("delete got%v,%v", d, e)
	}
}

func TestPublicCase02(t *testing.T) {
	d := subject.Dictionary{"": "", "A": "x"}
	if e := d.Delete(""); e != nil || len(d) != 1 || d["A"] != "x" {
		t.Errorf("empty present got%v,%v", d, e)
	}
}

func TestPublicCase03(t *testing.T) {
	for _, d := range []subject.Dictionary{nil, {}, {"A": "x"}} {
		before := maps.Clone(d)
		e := d.Delete("a")
		if e != subject.ErrAbsent || !maps.Equal(d, before) {
			t.Errorf("absent got%v,%v want unchanged,ErrAbsent", d, e)
		}
	}
}
