package subject_test

import (
	"task.local/dictionary-add-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	d := subject.Dictionary{}
	e := d.Add("test", "this is just a test")
	if e != nil || d["test"] != "this is just a test" {
		t.Errorf("insert got %v,%v", d, e)
	}
}

func TestContractCase02(t *testing.T) {
	d := subject.Dictionary{"test": "old", "other": "keep"}
	e := d.Add("test", "new")
	if e != subject.ErrExists || d["test"] != "old" || d["other"] != "keep" || len(d) != 2 {
		t.Errorf("existing got %v,%v want unchanged,ErrExists", d, e)
	}
}

func TestContractCase03(t *testing.T) {
	d := subject.Dictionary{"empty": ""}
	if e := d.Add("empty", "new"); e != subject.ErrExists || d["empty"] != "" {
		t.Errorf("present empty got %v,%v", d, e)
	}
}

func TestContractCase04(t *testing.T) {
	d := subject.Dictionary{"other": "keep"}
	e := d.Add("", "")
	v, ok := d[""]
	if e != nil || !ok || v != "" || d["other"] != "keep" || len(d) != 2 {
		t.Errorf("empty insertion got %v,%v", d, e)
	}
}
