package subject_test

import (
	"task.local/composite-key-table-001/subject"
	"testing"
)

func TestCompositeKeyRange(t *testing.T) {
	table := subject.NewTable()
	for a := -2; a <= 2; a++ {
		for b := -2; b <= 2; b++ {
			for class := -2; class <= 2; class++ {
				key := subject.Key{Parts: [2]int{a, b}, Class: class}
				value := string(rune('A' + (a+2)*25 + (b+2)*5 + class + 2))
				subject.Put(table, &key, value)
			}
		}
	}
	for a := -2; a <= 2; a++ {
		for b := -2; b <= 2; b++ {
			for class := -2; class <= 2; class++ {
				key := subject.Key{Parts: [2]int{a, b}, Class: class}
				want := string(rune('A' + (a+2)*25 + (b+2)*5 + class + 2))
				if got, ok := subject.Lookup(table, key); got != want || !ok {
					t.Fatalf("Lookup(%v) = %q,%t, want %q,true", key, got, ok, want)
				}
			}
		}
	}
	max := int(^uint(0) >> 1)
	for _, key := range []subject.Key{{Parts: [2]int{-max - 1, max}, Class: max}, {Parts: [2]int{max, -max - 1}, Class: -max - 1}} {
		subject.Put(table, &key, "boundary")
		if got, ok := subject.Lookup(table, key); got != "boundary" || !ok {
			t.Errorf("boundary lookup = %q,%t, want boundary,true", got, ok)
		}
	}
}
