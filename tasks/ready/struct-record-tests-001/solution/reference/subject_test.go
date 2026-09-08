package subject_test

import (
	"reflect"
	"task.local/struct-record-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	want := subject.Entry{Name: "record", Age: 29, Attributes: map[string]string{"street": "Main St."}}
	got := subject.NewEntry(want.Name, want.Age, want.Attributes)
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("record got %+v want %+v", got, want)
	}
}

func TestContractCase02(t *testing.T) {
	got := subject.NewEntry("", 0, nil)
	if got == nil || got.Name != "" || got.Age != 0 || got.Attributes != nil {
		t.Errorf("zero record got %+v", got)
	}
}

func TestContractCase03(t *testing.T) {
	got := subject.NewEntry("n", -2, map[string]string{})
	if got == nil || got.Age != -2 || got.Attributes == nil || len(got.Attributes) != 0 {
		t.Errorf("empty map/negative age got %+v", got)
	}
}

func TestContractCase04(t *testing.T) {
	a := map[string]string{"k": "v"}
	subject.NewEntry("n", 3, a)
	if len(a) != 1 || a["k"] != "v" {
		t.Errorf("attributes changed: %v", a)
	}
}

func TestContractCase05(t *testing.T) {
	a := subject.NewEntry("a", 1, nil)
	b := subject.NewEntry("b", 2, nil)
	if a == nil || b == nil {
		t.Fatal("got nil record")
	}
	a.Name = "changed"
	if b.Name != "b" {
		t.Errorf("independent record got %q want b", b.Name)
	}
}
