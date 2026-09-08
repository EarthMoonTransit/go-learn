package subject_test

import (
	"task.local/category-total-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	n, e := subject.CategoryTotal([]subject.Record{{Day: -2, Amount: 3, Category: "a"}, {Day: -2, Amount: 4, Category: "b"}}, subject.Period{From: -2, To: -2}, "a")
	if e != nil || n != 3 {
		t.Errorf("single day got %g,%v want3,nil", n, e)
	}
}
