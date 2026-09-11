package subject_test

import (
	"reflect"
	"task.local/bounded-lru-cache-001/subject"
	"testing"
)

type pair struct{ key, value int }

func snapshot(c subject.Cache) []pair {
	out := []pair{}
	c.Range(func(k, v int) bool { out = append(out, pair{k, v}); return true })
	return out
}
func wantEntries(t *testing.T, c subject.Cache, want ...pair) {
	t.Helper()
	if want == nil {
		want = []pair{}
	}
	if got := snapshot(c); !reflect.DeepEqual(got, want) {
		t.Fatalf("Range oldest-to-newest = %v; want %v", got, want)
	}
}
func wantGet(t *testing.T, c subject.Cache, k, v int, ok bool) {
	t.Helper()
	if got, found := c.Get(k); got != v || found != ok {
		t.Fatalf("Get(%d) = (%d,%t); want (%d,%t)", k, got, found, v, ok)
	}
}
func newCache(t *testing.T, capacity int) subject.Cache {
	t.Helper()
	c := subject.New(capacity)
	if c == nil {
		t.Fatal("New returned nil; want usable Cache")
	}
	return c
}
func TestPublicEmptyZeroAndIndependent(t *testing.T) {
	a, b := newCache(t, 2), newCache(t, 2)
	wantEntries(t, a)
	wantGet(t, a, 9, 0, false)
	a.Set(-4, 0)
	b.Set(6, 21)
	wantEntries(t, a, pair{-4, 0})
	wantEntries(t, b, pair{6, 21})
	wantGet(t, a, -4, 0, true)
	wantGet(t, a, 6, 0, false)
	z := newCache(t, 0)
	z.Set(3, 8)
	z.Set(3, 9)
	wantGet(t, z, 3, 0, false)
	wantEntries(t, z)
	z.Clear()
	z.Set(4, 2)
	wantEntries(t, z)
}
func TestPublicGetRefreshAndEviction(t *testing.T) {
	c := newCache(t, 2)
	c.Set(8, 80)
	c.Set(2, 20)
	wantGet(t, c, 8, 80, true)
	wantEntries(t, c, pair{2, 20}, pair{8, 80})
	c.Set(5, 50)
	wantGet(t, c, 2, 0, false)
	wantEntries(t, c, pair{8, 80}, pair{5, 50})
}
func TestPublicUpdateRefreshWithoutEviction(t *testing.T) {
	c := newCache(t, 2)
	c.Set(7, 70)
	c.Set(-3, 30)
	c.Set(7, 71)
	wantEntries(t, c, pair{-3, 30}, pair{7, 71})
	c.Set(7, 72)
	wantEntries(t, c, pair{-3, 30}, pair{7, 72})
	c.Set(4, 40)
	wantEntries(t, c, pair{7, 72}, pair{4, 40})
	wantGet(t, c, -3, 0, false)
}
func TestPublicMissDoesNotChangeOrder(t *testing.T) {
	c := newCache(t, 2)
	c.Set(5, 0)
	c.Set(1, 10)
	wantGet(t, c, 99, 0, false)
	wantEntries(t, c, pair{5, 0}, pair{1, 10})
	c.Set(3, 30)
	wantEntries(t, c, pair{1, 10}, pair{3, 30})
}
func TestPublicRangeStopAndNoMutation(t *testing.T) {
	c := newCache(t, 3)
	c.Set(9, 90)
	c.Set(-2, 20)
	c.Set(4, 40)
	var seen []pair
	c.Range(func(k, v int) bool { seen = append(seen, pair{k, v}); return len(seen) < 2 })
	if want := []pair{{9, 90}, {-2, 20}}; !reflect.DeepEqual(seen, want) {
		t.Fatalf("Range stopping on second callback = %v; want %v", seen, want)
	}
	calls := 0
	c.Range(func(k, v int) bool { calls++; return false })
	if calls != 1 {
		t.Fatalf("Range stopping immediately called %d times; want 1", calls)
	}
	wantEntries(t, c, pair{9, 90}, pair{-2, 20}, pair{4, 40})
	c.Set(6, 60)
	wantEntries(t, c, pair{-2, 20}, pair{4, 40}, pair{6, 60})
}
func TestPublicClearAndReuse(t *testing.T) {
	c := newCache(t, 2)
	c.Set(1, 11)
	c.Set(2, 22)
	c.Clear()
	c.Clear()
	wantEntries(t, c)
	wantGet(t, c, 1, 0, false)
	wantGet(t, c, 2, 0, false)
	c.Set(3, 33)
	c.Set(4, 44)
	c.Set(5, 55)
	wantEntries(t, c, pair{4, 44}, pair{5, 55})
}

func TestPublicZeroKeyAndNegativeValue(t *testing.T) {
	c := newCache(t, 2)
	c.Set(0, -17)
	c.Set(-8, -23)
	wantEntries(t, c, pair{0, -17}, pair{-8, -23})
	wantGet(t, c, 0, -17, true)
	wantEntries(t, c, pair{-8, -23}, pair{0, -17})
	c.Set(-8, -29)
	wantGet(t, c, -8, -29, true)
}
