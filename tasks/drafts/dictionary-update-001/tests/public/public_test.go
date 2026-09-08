package subject_test

import (
	"task.local/dictionary-update-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	d := subject.Dictionary{"test": "old", "other": "keep"}
	e := d.Update("test", "new")
	if e != nil || d["test"] != "new" || d["other"] != "keep" || len(d) != 2 {
		t.Errorf("update got %v,%v", d, e)
	}
}

func TestPublicCase02(t *testing.T) {
	for _, d := range []subject.Dictionary{nil, {}, {"A": "keep"}} {
		e := d.Update("a", "new")
		typed, ok := e.(subject.DictionaryError)
		if !ok || typed != subject.ErrAbsent || e.Error() != "cannot update absent key" {
			t.Errorf("absent got %T,%v want DictionaryError ErrAbsent", e, e)
		}
		if len(d) > 0 && (len(d) != 1 || d["A"] != "keep") {
			t.Errorf("absent changed map %v", d)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	d := subject.Dictionary{"": "", "a": "v"}
	if e := d.Update("", "new"); e != nil || d[""] != "new" {
		t.Errorf("existing empty got %v,%v", d, e)
	}
	if e := d.Update("a", ""); e != nil || d["a"] != "" || len(d) != 2 {
		t.Errorf("new empty got %v,%v", d, e)
	}
}
