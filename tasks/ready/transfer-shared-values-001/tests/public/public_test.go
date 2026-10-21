package subject_test

import (
	subject "task.local/transfer-shared-values-001/subject"
	"testing"
)

func TestSharedTransfer(t *testing.T) {
	cases := []struct{ dst, src, limit, amount, wantDst, wantSrc int }{
		{4, 9, 3, 3, 7, 6}, {4, 9, 20, 9, 13, 0}, {7, 7, 2, 2, 9, 5}, {8, 0, 4, 0, 8, 0},
		{4, 9, 0, 0, 4, 9}, {4, 9, -5, 0, 4, 9}, {0, 1, 1, 1, 1, 0},
	}
	for _, c := range cases {
		dst, src := c.dst, c.src
		da, sa := &dst, &src
		got := subject.Transfer(&dst, &src, c.limit)
		if got != c.amount || *da != c.wantDst || *sa != c.wantSrc {
			t.Fatalf("dst=%d src=%d limit=%d: amount/dst/src=%d/%d/%d; want %d/%d/%d", c.dst, c.src, c.limit, got, *da, *sa, c.amount, c.wantDst, c.wantSrc)
		}
	}
	dst, src := 1, 8
	a := subject.Transfer(&dst, &src, 3)
	b := subject.Transfer(&dst, &src, 9)
	if a != 3 || b != 5 || dst != 9 || src != 0 {
		t.Fatalf("repeated transfers=%d,%d dst/src=%d/%d; want 3,5 9/0", a, b, dst, src)
	}
}
func TestTransferNilAndSharedIdentity(t *testing.T) {
	x := 13
	if got := subject.Transfer(&x, &x, 5); got != 0 || x != 13 {
		t.Fatalf("same pointer: got amount/value=%d/%d; want 0/13", got, x)
	}
	for _, limit := range []int{-1, 0, 4} {
		x = 17
		if got := subject.Transfer(nil, &x, limit); got != 0 || x != 17 {
			t.Fatalf("nil destination limit=%d: amount/source=%d/%d; want 0/17", limit, got, x)
		}
		x = 19
		if got := subject.Transfer(&x, nil, limit); got != 0 || x != 19 {
			t.Fatalf("nil source limit=%d: amount/destination=%d/%d; want 0/19", limit, got, x)
		}
	}
	if got := subject.Transfer(nil, nil, 5); got != 0 {
		t.Fatalf("both nil: amount=%d; want 0", got)
	}
}
