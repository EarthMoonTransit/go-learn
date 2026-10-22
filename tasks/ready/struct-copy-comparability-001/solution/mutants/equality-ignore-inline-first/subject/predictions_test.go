package subject_test

import (
	subject "task.local/struct-copy-comparability-001/subject"
	"testing"
)

func TestPredictShallowCopy(t *testing.T) {
	source, copied, same := subject.StructStages(2, -5, 7, -4, 9, 3, 8)
	wantSource := subject.RecordView{Count: 2, Inline: [2]int{-5, 7}, Pointed: 8}
	wantCopy := subject.RecordView{Count: 9, Inline: [2]int{3, 7}, Pointed: 8}
	if source != wantSource || copied != wantCopy || !same {
		t.Errorf("shallow copy source=%+v copy=%+v sameTarget=%t; want %+v/%+v/true", source, copied, same, wantSource, wantCopy)
	}
}

func TestPredictEqualityAndComparability(t *testing.T) {
	a, b := 6, 6
	left := subject.Record{Count: 1, Inline: [2]int{2, 3}, Target: &a}
	right := subject.Record{Count: 1, Inline: [2]int{2, 3}, Target: &b}
	if subject.EqualRecord(left, right) {
		t.Error("different pointer identities must compare unequal even when both integers are 6")
	}
	if !subject.EqualRecord(left, left) {
		t.Error("the identical record must compare equal")
	}
	// Count is an int. Inline is an array whose two int elements are comparable.
	// Target is a comparable pointer: equality uses identity, not the pointed value.
	// Every field of Record is therefore comparable. SliceRecord.Base is Record,
	// but Extra is a slice, so SliceRecord is not comparable, even with nil Extra.
	record, sliceRecord := subject.ComparableShapes()
	if !record || sliceRecord {
		t.Errorf("comparability=%t,%t; want Record=true, SliceRecord=false", record, sliceRecord)
	}
}
