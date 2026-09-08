package subject_test

import (
	"task.local/dictionary-add-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	sentinel := subject.ErrExists
	if sentinel == nil {
		t.Fatal("ErrExists: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrExists != sentinel {
			t.Error("ErrExists changed during call; want original supplied sentinel")
		}
	}()

	d := subject.Dictionary{}
	e := d.Add("test", "this is just a test")
	if e != nil || d["test"] != "this is just a test" {
		t.Errorf("insert got %v,%v", d, e)
	}
}

func TestContractCase02(t *testing.T) {
	sentinel := subject.ErrExists
	if sentinel == nil {
		t.Fatal("ErrExists: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrExists != sentinel {
			t.Error("ErrExists changed during call; want original supplied sentinel")
		}
	}()

	d := subject.Dictionary{"test": "old", "other": "keep"}
	e := d.Add("test", "new")
	if e != sentinel || d["test"] != "old" || d["other"] != "keep" || len(d) != 2 {
		t.Errorf("existing got %v,%v want unchanged,ErrExists", d, e)
	}
}

func TestContractCase03(t *testing.T) {
	sentinel := subject.ErrExists
	if sentinel == nil {
		t.Fatal("ErrExists: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrExists != sentinel {
			t.Error("ErrExists changed during call; want original supplied sentinel")
		}
	}()

	d := subject.Dictionary{"empty": ""}
	if e := d.Add("empty", "new"); e != sentinel || d["empty"] != "" {
		t.Errorf("present empty got %v,%v", d, e)
	}
}

func TestContractCase04(t *testing.T) {
	sentinel := subject.ErrExists
	if sentinel == nil {
		t.Fatal("ErrExists: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrExists != sentinel {
			t.Error("ErrExists changed during call; want original supplied sentinel")
		}
	}()

	d := subject.Dictionary{"other": "keep"}
	e := d.Add("", "")
	v, ok := d[""]
	if e != nil || !ok || v != "" || d["other"] != "keep" || len(d) != 2 {
		t.Errorf("empty insertion got %v,%v", d, e)
	}
}
