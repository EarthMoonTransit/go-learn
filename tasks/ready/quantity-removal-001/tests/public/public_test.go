package subject_test

import (
	"maps"
	"task.local/quantity-removal-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	items := map[string]int{"x": 12, "keep": 7}
	units := map[string]int{"half": 6}
	if !subject.RemoveQuantity(items, units, "x", "half") || items["x"] != 6 || items["keep"] != 7 {
		t.Errorf("subtract got%v", items)
	}
	if !subject.RemoveQuantity(items, units, "x", "half") {
		t.Error("exact removal rejected")
	}
	if _, ok := items["x"]; ok {
		t.Errorf("zero key remains %v", items)
	}
	if !maps.Equal(units, map[string]int{"half": 6}) || !maps.Equal(items, map[string]int{"keep": 7}) {
		t.Errorf("maps changed %v %v", items, units)
	}
}

func TestPublicCase02(t *testing.T) {
	for _, c := range []struct{ item, unit string }{{"missing", "u"}, {"x", "bad"}, {"x", "large"}} {
		items := map[string]int{"x": 3, "keep": 9}
		units := map[string]int{"u": 1, "large": 4}
		a, b := maps.Clone(items), maps.Clone(units)
		if subject.RemoveQuantity(items, units, c.item, c.unit) || !maps.Equal(items, a) || !maps.Equal(units, b) {
			t.Errorf("rejected %v changed maps %v %v", c, items, units)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	items := map[string]int{"": 2, "X": 3}
	units := map[string]int{"": 1}
	if !subject.RemoveQuantity(items, units, "", "") || items[""] != 1 {
		t.Errorf("empty keys got%v", items)
	}
	if subject.RemoveQuantity(items, units, "x", "") {
		t.Error("case-normalized key accepted")
	}
}
