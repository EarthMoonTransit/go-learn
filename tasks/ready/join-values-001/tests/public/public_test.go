package subject_test

import (
	"task.local/join-values-001/subject"
	"testing"
)

func TestBuildValues(t *testing.T) {
	values := []int{3, -2, 0}
	got := subject.FormatValues("v=", values)
	if got != "v=[3,-2,0]" {
		t.Fatalf("list=%q", got)
	}
	if values[0] != 3 || values[1] != -2 || values[2] != 0 {
		t.Fatal("input values changed")
	}
}
func TestBuildEmptyValues(t *testing.T) {
	if got := subject.FormatValues("n=", nil); got != "n=[]" {
		t.Fatalf("empty list=%q", got)
	}
}
