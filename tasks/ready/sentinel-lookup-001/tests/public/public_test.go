package subject_test

import (
	"task.local/sentinel-lookup-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	sentinel := subject.ErrMissing
	if sentinel == nil {
		t.Fatal("ErrMissing: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrMissing != sentinel {
			t.Error("ErrMissing changed during call; want original supplied sentinel")
		}
	}()

	d := subject.Dictionary{"test": "this is just a test"}
	v, e := d.Search("test")
	if v != "this is just a test" || e != nil {
		t.Errorf("present got (%q,%v)", v, e)
	}
}

func TestPublicCase02(t *testing.T) {
	sentinel := subject.ErrMissing
	if sentinel == nil {
		t.Fatal("ErrMissing: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrMissing != sentinel {
			t.Error("ErrMissing changed during call; want original supplied sentinel")
		}
	}()

	v, e := subject.Dictionary{"key": ""}.Search("key")
	if v != "" || e != nil {
		t.Errorf("present empty got (%q,%v) want empty,nil", v, e)
	}
}

func TestPublicCase03(t *testing.T) {
	sentinel := subject.ErrMissing
	if sentinel == nil {
		t.Fatal("ErrMissing: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrMissing != sentinel {
			t.Error("ErrMissing changed during call; want original supplied sentinel")
		}
	}()

	v, e := subject.Dictionary{"test": "value"}.Search("unknown")
	if v != "" || e != sentinel {
		t.Errorf("absent got (%q,%v) want empty,ErrMissing", v, e)
	}
}

func TestPublicCase04(t *testing.T) {
	sentinel := subject.ErrMissing
	if sentinel == nil {
		t.Fatal("ErrMissing: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrMissing != sentinel {
			t.Error("ErrMissing changed during call; want original supplied sentinel")
		}
	}()

	var d subject.Dictionary
	v, e := d.Search("")
	if v != "" || e != sentinel {
		t.Errorf("nil map got (%q,%v) want empty,ErrMissing", v, e)
	}
}

func TestPublicCase05(t *testing.T) {
	sentinel := subject.ErrMissing
	if sentinel == nil {
		t.Fatal("ErrMissing: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrMissing != sentinel {
			t.Error("ErrMissing changed during call; want original supplied sentinel")
		}
	}()

	d := subject.Dictionary{"": "v", "A": "x"}
	v, e := d.Search("")
	if v != "v" || e != nil {
		t.Errorf("empty key got (%q,%v)", v, e)
	}
	_, e = d.Search("a")
	if e != sentinel {
		t.Errorf("case lookup got %v want ErrMissing", e)
	}
	if len(d) != 2 || d[""] != "v" || d["A"] != "x" {
		t.Errorf("dictionary changed: %v", d)
	}
}
