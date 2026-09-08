package subject_test

import (
	"math"
	"task.local/category-total-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	original := subject.ErrCategory
	if original == nil {
		t.Fatal("nil original sentinel")
	}
	defer func() {
		if subject.ErrCategory != original {
			t.Error("sentinel identity changed")
		}
	}()
	n, e := subject.CategoryTotal([]subject.Record{{Day: -2, Amount: 3, Category: "a"}, {Day: -2, Amount: 4, Category: "b"}}, subject.Period{From: -2, To: -2}, "a")
	if e != nil || (math.IsNaN(n) || math.Abs(n-3) > 1e-9) {
		t.Errorf("single day got %g,%v want3,nil", n, e)
	}
}
