package subject_test

import (
	"reflect"
	"task.local/bounded-lru-cache-001/subject"
	"testing"
)

type item[K comparable, V any] struct {
	key   K
	value V
}

func genericNew[K comparable, V any](t *testing.T, capacity int) *subject.Cache[K, V] {
	t.Helper()
	c := subject.New[K, V](capacity)
	if c == nil {
		t.Fatal("New returned nil; want usable cache")
	}
	return c
}
func genericEntries[K comparable, V any](t *testing.T, c *subject.Cache[K, V], want ...item[K, V]) {
	t.Helper()
	got := []item[K, V]{}
	c.Range(func(k K, v V) bool { got = append(got, item[K, V]{k, v}); return true })
	if len(got) != len(want) {
		t.Fatalf("Range length = %d; want %d (got %#v, want %#v)", len(got), len(want), got, want)
	}
	for i := range got {
		if got[i].key != want[i].key || !reflect.DeepEqual(got[i].value, want[i].value) {
			t.Fatalf("Range entry %d = (%#v,%#v); want (%#v,%#v)", i, got[i].key, got[i].value, want[i].key, want[i].value)
		}
	}
}
func genericGet[K comparable, V any](t *testing.T, c *subject.Cache[K, V], key K, want V, present bool) {
	t.Helper()
	got, ok := c.Get(key)
	if ok != present || !reflect.DeepEqual(got, want) {
		t.Fatalf("Get(%#v) = (%#v,%t); want (%#v,%t)", key, got, ok, want, present)
	}
}
func TestPublicZeroCache(t *testing.T) {
	var c subject.Cache[string, []int]
	genericEntries(t, &c)
	genericGet(t, &c, "", nil, false)
	c.Clear()
	c.Set("a", []int{7, 2})
	c.Set("", nil)
	genericGet(t, &c, "a", nil, false)
	genericGet(t, &c, "", nil, false)
	genericEntries(t, &c)
	c.Clear()
	c.Set("b", []int{9})
	genericEntries(t, &c)
}
func TestPublicGenericSlicesAndRecency(t *testing.T) {
	c := genericNew[string, []int](t, 2)
	c.Set("", []int{2, 5})
	c.Set("beta", []int{8})
	genericGet(t, c, "", []int{2, 5}, true)
	genericEntries(t, c, item[string, []int]{"beta", []int{8}}, item[string, []int]{"", []int{2, 5}})
	c.Set("beta", []int{-3, 4})
	genericGet(t, c, "absent", nil, false)
	c.Set("next", []int{6, 7})
	genericGet(t, c, "", nil, false)
	genericEntries(t, c, item[string, []int]{"beta", []int{-3, 4}}, item[string, []int]{"next", []int{6, 7}})
	var seen []item[string, []int]
	c.Range(func(k string, v []int) bool { seen = append(seen, item[string, []int]{k, v}); return false })
	if want := []item[string, []int]{{"beta", []int{-3, 4}}}; !reflect.DeepEqual(seen, want) {
		t.Fatalf("Range stop = %#v; want %#v", seen, want)
	}
	genericEntries(t, c, item[string, []int]{"beta", []int{-3, 4}}, item[string, []int]{"next", []int{6, 7}})
	c.Clear()
	c.Clear()
	genericGet(t, c, "next", nil, false)
	c.Set("one", []int{1})
	c.Set("two", []int{2})
	c.Set("three", []int{3})
	genericEntries(t, c, item[string, []int]{"two", []int{2}}, item[string, []int]{"three", []int{3}})
}
func TestPublicTypedMissesAndStoredNil(t *testing.T) {
	slices := genericNew[int, []int](t, 2)
	slices.Set(1, []int{3})
	genericGet(t, slices, 2, nil, false)
	slices.Set(2, nil)
	genericGet(t, slices, 2, nil, true)
	pointers := genericNew[string, *int](t, 2)
	n := 8
	pointers.Set("number", &n)
	genericGet(t, pointers, "missing", nil, false)
	if got, ok := pointers.Get("number"); !ok || got != &n {
		t.Fatalf("Get(number) = (%p,%t); want original pointer %p,true", got, ok, &n)
	}
	pointers.Set("nil", nil)
	genericGet(t, pointers, "nil", nil, true)
	type record struct {
		Label  string
		Values []int
	}
	records := genericNew[int, record](t, 2)
	records.Set(1, record{"sample", []int{2, 9}})
	genericGet(t, records, 9, record{}, false)
	records.Set(2, record{})
	genericGet(t, records, 2, record{}, true)
	genericGet(t, records, 1, record{"sample", []int{2, 9}}, true)
}
func TestPublicNamedCompositeKeysAndStructValues(t *testing.T) {
	type name string
	type key struct {
		Group name
		Index int
	}
	type value struct {
		Label  name
		Values []int
	}
	c := genericNew[key, value](t, 2)
	a, b := key{"same", 1}, key{"same", 2}
	c.Set(a, value{"a", []int{11}})
	c.Set(b, value{"b", []int{22}})
	genericGet(t, c, a, value{"a", []int{11}}, true)
	c.Set(b, value{"updated", []int{4, 5}})
	c.Set(key{}, value{"zero", []int{6}})
	genericGet(t, c, a, value{}, false)
	genericEntries(t, c, item[key, value]{b, value{"updated", []int{4, 5}}}, item[key, value]{key{}, value{"zero", []int{6}}})
	named := genericNew[name, []int](t, 2)
	named.Set(name("x"), []int{1})
	named.Set(name(""), []int{2})
	genericGet(t, named, name("x"), []int{1}, true)
	genericGet(t, named, name(""), []int{2}, true)
}
func TestPublicPointerKeyIdentity(t *testing.T) {
	a, b := 7, 7
	c := genericNew[*int, string](t, 3)
	c.Set(&a, "first")
	c.Set(&b, "second")
	c.Set(nil, "nil key")
	genericGet(t, c, &a, "first", true)
	genericGet(t, c, &b, "second", true)
	genericGet(t, c, nil, "nil key", true)
	genericEntries(t, c, item[*int, string]{&a, "first"}, item[*int, string]{&b, "second"}, item[*int, string]{nil, "nil key"})
}
func TestPublicAssignmentAndIndependentGenericInstances(t *testing.T) {
	a, b := genericNew[string, []int](t, 2), genericNew[string, []int](t, 1)
	input := []int{5, 9}
	a.Set("x", input)
	b.Set("y", []int{8})
	input[0] = 17
	genericGet(t, a, "x", []int{17, 9}, true)
	got, ok := a.Get("x")
	if !ok || len(got) != 2 {
		t.Fatalf("Get(x) = (%v,%t); want two retained values", got, ok)
	}
	got[1] = 21
	genericGet(t, a, "x", []int{17, 21}, true)
	genericEntries(t, b, item[string, []int]{"y", []int{8}})
	a.Clear()
	genericEntries(t, b, item[string, []int]{"y", []int{8}})
}
