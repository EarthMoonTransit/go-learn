package subject_test

import (
	"slices"
	"task.local/named-cases-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	var calls []string
	t.Run("group", func(parent *testing.T) {
		cases := []subject.Case{{Name: "first", Check: func(child *testing.T) {
			if child == parent || child.Name() != parent.Name()+"/first" {
				t.Errorf("child first name%q", child.Name())
			}
			calls = append(calls, "first")
		}}, {Name: "second", Check: func(child *testing.T) {
			if child == parent || child.Name() != parent.Name()+"/second" {
				t.Errorf("child second name%q", child.Name())
			}
			calls = append(calls, "second")
		}}}
		subject.RunNamed(parent, cases)
		if !slices.Equal(calls, []string{"first", "second"}) || cases[0].Name != "first" || cases[1].Name != "second" {
			t.Errorf("serial cases calls%v names%s,%s", calls, cases[0].Name, cases[1].Name)
		}
	})
}

func TestPublicCase02(t *testing.T) {
	subject.RunNamed(t, nil)
	subject.RunNamed(t, []subject.Case{})
	calls := 0
	t.Run("single", func(parent *testing.T) {
		subject.RunNamed(parent, []subject.Case{{Name: "case-1", Check: func(child *testing.T) {
			calls++
			if child.Name() != parent.Name()+"/case-1" {
				t.Errorf("single name%q", child.Name())
			}
		}}})
	})
	if calls != 1 {
		t.Errorf("single calls%d want1", calls)
	}
}
