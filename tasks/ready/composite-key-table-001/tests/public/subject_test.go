package subject_test

import (
	"task.local/composite-key-table-001/subject"
	"testing"
)

func TestCompositeEqualityAndReplacement(t *testing.T) {
	table := subject.NewTable()
	base := subject.Key{Parts: [2]int{2, 5}, Class: 7}
	if got, ok := subject.Lookup(table, base); got != "" || ok {
		t.Errorf("empty lookup = %q,%t, want empty,false", got, ok)
	}
	subject.Put(table, &base, "first")
	equal := subject.Key{Parts: [2]int{2, 5}, Class: 7}
	if got, ok := subject.Lookup(table, equal); got != "first" || !ok {
		t.Errorf("equal key lookup = %q,%t, want first,true", got, ok)
	}
	subject.Put(table, &equal, "replacement")
	if got, ok := subject.Lookup(table, base); got != "replacement" || !ok {
		t.Errorf("replacement lookup = %q,%t, want replacement,true", got, ok)
	}
	for _, key := range []subject.Key{{Parts: [2]int{3, 5}, Class: 7}, {Parts: [2]int{2, 6}, Class: 7}, {Parts: [2]int{2, 5}, Class: 8}} {
		subject.Put(table, &key, "distinct")
		if got, ok := subject.Lookup(table, base); got != "replacement" || !ok {
			t.Errorf("changed component collided: base = %q,%t, want replacement,true", got, ok)
		}
		if got, ok := subject.Lookup(table, key); got != "distinct" || !ok {
			t.Errorf("distinct key lookup = %q,%t, want distinct,true", got, ok)
		}
	}
}

func TestInsertionCopiesCallerKey(t *testing.T) {
	table := subject.NewTable()
	key := subject.Key{Parts: [2]int{-2, 4}, Class: 6}
	saved := key
	subject.Put(table, &key, "saved")
	key.Parts[0], key.Parts[1], key.Class = 8, 9, 10
	if got, ok := subject.Lookup(table, saved); got != "saved" || !ok {
		t.Errorf("original copied key = %q,%t, want saved,true after caller mutation", got, ok)
	}
	if got, ok := subject.Lookup(table, key); got != "" || ok {
		t.Errorf("mutated caller key = %q,%t, want empty,false", got, ok)
	}
}

func TestPresenceAndIndependentTables(t *testing.T) {
	a, b := subject.NewTable(), subject.NewTable()
	key := subject.Key{}
	subject.Put(a, &key, "")
	for i := 0; i < 2; i++ {
		if got, ok := subject.Lookup(a, key); got != "" || !ok {
			t.Errorf("present empty value = %q,%t, want empty,true", got, ok)
		}
	}
	if got, ok := subject.Lookup(b, key); got != "" || ok {
		t.Errorf("independent table = %q,%t, want empty,false", got, ok)
	}
}
