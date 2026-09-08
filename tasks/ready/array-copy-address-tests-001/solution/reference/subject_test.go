package subject_test

import (
	"task.local/array-copy-address-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	source := &[5]string{"a", "b", "c", "d", "e"}
	out, ptrs := subject.CopyArray(source)
	if out == nil || *out != *source {
		t.Fatalf("copy got%v want%v", out, source)
	}
	for i, p := range ptrs {
		if p != &out[i] {
			t.Errorf("address%d not copied element", i)
		}
	}
}

func TestContractCase02(t *testing.T) {
	source := &[5]string{"", "λ", "x", "x", "last"}
	a, ptrs := subject.CopyArray(source)
	b, _ := subject.CopyArray(source)
	if a == nil || b == nil {
		t.Fatal("nil copy")
	}
	if *a != *source || *b != *source {
		t.Fatal("values mismatch")
	}
	a[0] = "changed"
	source[1] = "new"
	if source[0] != "" || a[1] != "λ" || b[0] != "" || b[1] != "λ" {
		t.Errorf("arrays share state source%v a%v b%v", source, a, b)
	}
	*ptrs[4] = "through pointer"
	if a[4] != "through pointer" || source[4] != "last" || b[4] != "last" {
		t.Error("element pointer targets wrong storage")
	}
}
