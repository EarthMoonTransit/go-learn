package subject_test

import (
	"math"
	"slices"
	"task.local/category-total-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	original := subject.ErrCategory
	if original == nil {
		t.Fatal("nil original sentinel")
	}
	defer func() {
		if subject.ErrCategory != original {
			t.Error("sentinel identity changed")
		}
	}()
	sentinel := subject.ErrCategory
	if sentinel == nil {
		t.Fatal("nil sentinel")
	}
	defer func() {
		if subject.ErrCategory != sentinel {
			t.Error("sentinel changed")
		}
	}()
	in := []subject.Record{{Day: 1, Amount: 2.5, Category: "a"}, {Day: 2, Amount: -1, Category: "a"}, {Day: 2, Amount: 9, Category: "b"}, {Day: 4, Amount: 3, Category: "a"}}
	before := slices.Clone(in)
	n, e := subject.CategoryTotal(in, subject.Period{From: 1, To: 2}, "a")
	if e != nil || math.IsNaN(n) || math.Abs(n-1.5) > 1e-9 {
		t.Errorf("total got %g,%v want1.5,nil", n, e)
	}
	if !slices.Equal(in, before) {
		t.Errorf("input changed %v", in)
	}
}

func TestPublicCase02(t *testing.T) {
	original := subject.ErrCategory
	if original == nil {
		t.Fatal("nil original sentinel")
	}
	defer func() {
		if subject.ErrCategory != original {
			t.Error("sentinel identity changed")
		}
	}()
	for _, p := range []subject.Period{{From: 9, To: 10}, {From: 3, To: 1}} {
		n, e := subject.CategoryTotal([]subject.Record{{Day: 2, Amount: 8, Category: "a"}}, p, "a")
		if (math.IsNaN(n) || math.Abs(n) > 1e-9) || e != nil {
			t.Errorf("known outside got %g,%v want0,nil", n, e)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	original := subject.ErrCategory
	if original == nil {
		t.Fatal("nil original sentinel")
	}
	defer func() {
		if subject.ErrCategory != original {
			t.Error("sentinel identity changed")
		}
	}()
	for _, in := range [][]subject.Record{nil, {{Category: "A", Amount: 5}}} {
		sentinel := subject.ErrCategory
		n, e := subject.CategoryTotal(in, subject.Period{}, "a")
		if sentinel == nil || e != sentinel || (math.IsNaN(n) || math.Abs(n) > 1e-9) || subject.ErrCategory != sentinel {
			t.Errorf("unknown got %g,%v want original sentinel", n, e)
		}
	}
}

func TestPublicCase04(t *testing.T) {
	original := subject.ErrCategory
	if original == nil {
		t.Fatal("nil original sentinel")
	}
	defer func() {
		if subject.ErrCategory != original {
			t.Error("sentinel identity changed")
		}
	}()
	n, e := subject.CategoryTotal([]subject.Record{{Day: 0, Amount: 0, Category: ""}}, subject.Period{}, "")
	if (math.IsNaN(n) || math.Abs(n) > 1e-9) || e != nil {
		t.Errorf("empty category got %g,%v want0,nil", n, e)
	}
}
