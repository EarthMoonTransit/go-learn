package subject_test

import (
	subject "task.local/struct-copy-comparability-001/subject"
	"testing"
)

func publicStructStages(t *testing.T, input [7]int) {
	t.Helper()
	source, copied, same := subject.StructStages(input[0], input[1], input[2], input[3], input[4], input[5], input[6])
	if source.Count != input[0] || source.Inline != [2]int{input[1], input[2]} {
		t.Errorf("source Count/Inline=%d/%v; want %d/%v after modifying only the copy", source.Count, source.Inline, input[0], [2]int{input[1], input[2]})
	}
	if copied.Count != input[4] {
		t.Errorf("copied Count=%d; want %d", copied.Count, input[4])
	}
	if copied.Inline != [2]int{input[5], input[2]} {
		t.Errorf("copied Inline=%v; want %v with only index 0 changed", copied.Inline, [2]int{input[5], input[2]})
	}
	if source.Pointed != input[6] || copied.Pointed != input[6] {
		t.Errorf("shared pointee observations=%d/%d; want %d/%d after writing through the copied pointer", source.Pointed, copied.Pointed, input[6], input[6])
	}
	if !same {
		t.Error("sameTarget=false; want true: struct assignment preserves the pointer identity")
	}
}

func TestPublicScalarInlineAndPointeeStages(t *testing.T) {
	publicStructStages(t, [7]int{2, -5, 7, -4, 9, 3, 8})
	publicStructStages(t, [7]int{-9, 6, -2, 9, 0, 0, 0})
}

func TestPublicEqualityChecksEveryField(t *testing.T) {
	target, equalValue, otherValue := 4, 4, -3
	base := subject.Record{Count: 2, Inline: [2]int{-5, 7}, Target: &target}
	changedCount, changedFirst, changedLast := base, base, base
	changedCount.Count = -2
	changedFirst.Inline[0] = 5
	changedLast.Inline[1] = -7
	sameValue, other, withNil := base, base, base
	sameValue.Target, other.Target, withNil.Target = &equalValue, &otherValue, nil
	cases := []struct {
		label       string
		left, right subject.Record
		want        bool
	}{
		{"identical record", base, base, true},
		{"different Count", base, changedCount, false},
		{"different Inline[0]", base, changedFirst, false},
		{"different Inline[1]", base, changedLast, false},
		{"equal pointed values do not make different Target pointers equal", base, sameValue, false},
		{"different target and pointed value", base, other, false},
		{"one nil Target", base, withNil, false},
		{"other nil Target", withNil, base, false},
		{"matching nil Targets", withNil, withNil, true},
	}
	for _, tc := range cases {
		if got := subject.EqualRecord(tc.left, tc.right); got != tc.want {
			t.Errorf("%s: EqualRecord=%t; want %t", tc.label, got, tc.want)
		}
	}
}

func TestPublicAllFieldComparability(t *testing.T) {
	record, sliceRecord := subject.ComparableShapes()
	if !record {
		t.Error("Record comparable=false; want true: int, [2]int and *int are comparable")
	}
	if sliceRecord {
		t.Error("SliceRecord comparable=true; want false: Extra is a slice, including when nil or empty")
	}
}
