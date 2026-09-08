package subject_test

import (
	"reflect"
	"task.local/struct-record-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	want := subject.Entry{Name: "record", Age: 29, Attributes: map[string]string{"street": "Main St."}}
	got := subject.NewEntry(want.Name, want.Age, want.Attributes)
	if !reflect.DeepEqual(got, &want) {
		t.Errorf("record got %+v want %+v", got, want)
	}
}

// Add cases for the remaining requirements in TASK.md.
