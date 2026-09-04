package subject_test

import (
	"reflect"
	"task.local/append-value-001/subject"
	"testing"
)

func TestPushExamples(t *testing.T) {
	input := []int{2, 7}
	out := subject.Push(input, 9)
	if !reflect.DeepEqual(out, []int{2, 7, 9}) {
		t.Fatalf("Push=%v, want [2 7 9]", out)
	}
	if len(input) != 2 {
		t.Fatal("caller length changed")
	}
}
