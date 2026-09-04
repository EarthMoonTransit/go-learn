package subject_test

import (
	"task.local/clamp-tests-001/subject"
	"testing"
)

func TestClampBelow(t *testing.T) {
	if got := subject.Clamp(-4, 0, 10); got != 0 {
		t.Fatalf("Clamp(-4,0,10)=%d, want 0", got)
	}
}
func TestClampLowerEquality(t *testing.T) {
	if got := subject.Clamp(2, 2, 8); got != 2 {
		t.Fatalf("Clamp(2,2,8)=%d, want 2", got)
	}
}
func TestClampInside(t *testing.T) {
	if got := subject.Clamp(5, 0, 10); got != 5 {
		t.Fatalf("Clamp(5,0,10)=%d, want 5", got)
	}
}
func TestClampUpperEquality(t *testing.T) {
	if got := subject.Clamp(8, 2, 8); got != 8 {
		t.Fatalf("Clamp(8,2,8)=%d, want 8", got)
	}
}
func TestClampAbove(t *testing.T) {
	if got := subject.Clamp(14, 0, 10); got != 10 {
		t.Fatalf("Clamp(14,0,10)=%d, want 10", got)
	}
}
func TestClampSingleton(t *testing.T) {
	if got := subject.Clamp(20, 3, 3); got != 3 {
		t.Fatalf("Clamp(20,3,3)=%d, want 3", got)
	}
}
