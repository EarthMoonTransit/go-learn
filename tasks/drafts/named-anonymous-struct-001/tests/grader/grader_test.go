package subject_test

import (
	"task.local/named-anonymous-struct-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	a, b := subject.BuildContacts("", "", 0)
	if a != (subject.Contact{}) || b.Name != "" || b.Email != "" || b.Age != 0 {
		t.Error("zero values changed")
	}
}
