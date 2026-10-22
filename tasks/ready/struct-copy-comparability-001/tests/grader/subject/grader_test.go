package subject_test

import (
	subject "task.local/struct-copy-comparability-001/subject"
	"testing"
)

func checkStructTuple(t *testing.T, input [7]int) {
	t.Helper()
	source, copied, same := subject.StructStages(input[0], input[1], input[2], input[3], input[4], input[5], input[6])
	wantSource := subject.RecordView{Count: input[0], Inline: [2]int{input[1], input[2]}, Pointed: input[6]}
	wantCopy := subject.RecordView{Count: input[4], Inline: [2]int{input[5], input[2]}, Pointed: input[6]}
	if source != wantSource || copied != wantCopy || !same {
		t.Fatalf("input=%v: source=%+v copied=%+v sameTarget=%t; want %+v/%+v/true", input, source, copied, same, wantSource, wantCopy)
	}
}

func TestGraderFiniteStructInputs(t *testing.T) {
	values := [3]int{-2, 0, 3}
	for code := 0; code < 2187; code++ {
		var input [7]int
		remainder := code
		for i := range input {
			input[i] = values[remainder%3]
			remainder /= 3
		}
		checkStructTuple(t, input)
	}
	first := [7]int{-9, 9, -9, 9, 9, -9, -9}
	checkStructTuple(t, first)
	checkStructTuple(t, [7]int{9, -9, 9, -9, -9, 9, 9})
	checkStructTuple(t, first)
}

func TestGraderRecordEqualityMatrix(t *testing.T) {
	a, b, c := -9, -9, 9
	targets := []*int{nil, &a, &b, &c}
	var records []subject.Record
	for _, count := range []int{-9, 0, 9} {
		for _, first := range []int{-9, 0, 9} {
			for _, last := range []int{-9, 0, 9} {
				for _, target := range targets {
					records = append(records, subject.Record{Count: count, Inline: [2]int{first, last}, Target: target})
				}
			}
		}
	}
	for i, left := range records {
		for j, right := range records {
			want := left.Count == right.Count && left.Inline[0] == right.Inline[0] && left.Inline[1] == right.Inline[1] && left.Target == right.Target
			if got := subject.EqualRecord(left, right); got != want {
				t.Fatalf("record pair %d/%d equality=%t; want %t from every field and pointer identity", i, j, got, want)
			}
		}
	}
}
