package subject_test

import (
	"task.local/quantity-presence-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if a, b := subject.Quantity(map[string]int{"item": 12}, "item"); a != 12 || b != true {
		t.Errorf("Quantity: got (%v,%v), want (%v,%v)", a, b, 12, true)
	}
}

func TestPublicCase02(t *testing.T) {
	if a, b := subject.Quantity(map[string]int{"item": 0}, "item"); a != 0 || b != true {
		t.Errorf("Quantity: got (%v,%v), want (%v,%v)", a, b, 0, true)
	}
}

func TestPublicCase03(t *testing.T) {
	if a, b := subject.Quantity(map[string]int{"item": 12}, "absent"); a != 0 || b != false {
		t.Errorf("Quantity: got (%v,%v), want (%v,%v)", a, b, 0, false)
	}
}

func TestPublicCase04(t *testing.T) {
	if a, b := subject.Quantity(nil, "item"); a != 0 || b != false {
		t.Errorf("Quantity: got (%v,%v), want (%v,%v)", a, b, 0, false)
	}
}

func TestPublicCase05(t *testing.T) {
	if a, b := subject.Quantity(map[string]int{"": -2}, ""); a != -2 || b != true {
		t.Errorf("Quantity: got (%v,%v), want (%v,%v)", a, b, -2, true)
	}
}

func TestPublicCase06(t *testing.T) {
	m := map[string]int{"a": 3, "b": 4}
	subject.Quantity(m, "a")
	if len(m) != 2 || m["a"] != 3 || m["b"] != 4 {
		t.Errorf("input changed: %v", m)
	}
}

func TestPublicCase07(t *testing.T) {
	if a, b := subject.Quantity(map[string]int{"A": 3}, "a"); a != 0 || b != false {
		t.Errorf("Quantity: got (%v,%v), want (%v,%v)", a, b, 0, false)
	}
}
