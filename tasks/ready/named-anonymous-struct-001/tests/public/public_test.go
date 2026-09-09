package subject_test

import (
	"task.local/named-anonymous-struct-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	a, b := subject.BuildContacts("Ada", "a@example.test", 37)
	if a.Name != "Ada" || a.Email != "a@example.test" || a.Age != 37 || b.Name != a.Name || b.Email != a.Email || b.Age != 37 {
		t.Errorf("fields got%+v %+v", a, b)
	}
}

func TestPublicCase02(t *testing.T) {
	a, b := subject.BuildContacts(" λ ", "", 0)
	if a != (subject.Contact{Name: " λ "}) || b.Name != " λ " || b.Email != "" || b.Age != 0 {
		t.Errorf("zero/exact got%+v %+v", a, b)
	}
	a.Name = "changed"
	b.Age = 9
	c, d := subject.BuildContacts("", " x ", 150)
	if b.Name != " λ " || a.Age != 0 || c.Email != " x " || c.Age != 150 || d.Email != c.Email || d.Age != 150 {
		t.Errorf("independence got%+v %+v %+v %+v", a, b, c, d)
	}
}
