package subject_test

import (
	"fmt"
	"reflect"
	"sort"
	"task.local/bounded-lru-cache-001/subject"
	"testing"
)

// Logical access counters form an independent deterministic oracle.
func TestGraderOperationTraces(t *testing.T) {
	for _, capacity := range []int{0, 1, 2, 4} {
		t.Run(fmt.Sprint(capacity), func(t *testing.T) {
			c := subject.New[int, int](capacity)
			values, access := map[int]int{}, map[int]int{}
			tick := 0
			keys := []int{0, -7, 12, 0, 42, -7, 3, 12, 0, -7, 9, 42}
			for round := 0; round < 3; round++ {
				for step, key := range keys {
					tick++
					if step%4 == 1 {
						got, ok := c.Get(key)
						want, exists := values[key]
						if got != want || ok != exists {
							t.Fatalf("capacity=%d round=%d step=%d Get(%d)=(%d,%t); want (%d,%t)", capacity, round, step, key, got, ok, want, exists)
						}
						if exists {
							access[key] = tick
						}
					} else {
						value := round*20 - step
						c.Set(key, value)
						if capacity > 0 {
							values[key] = value
							access[key] = tick
							if len(values) > capacity {
								oldest := key
								for k, when := range access {
									if when < access[oldest] {
										oldest = k
									}
								}
								delete(values, oldest)
								delete(access, oldest)
							}
						}
					}
					order := []int{}
					for k := range values {
						order = append(order, k)
					}
					sort.Slice(order, func(i, j int) bool { return access[order[i]] < access[order[j]] })
					expected := []pair{}
					for _, k := range order {
						expected = append(expected, pair{k, values[k]})
					}
					wantEntries(t, c, expected...)
				}
				if round == 1 {
					c.Clear()
					values = map[int]int{}
					access = map[int]int{}
					wantEntries(t, c)
				}
			}
		})
	}
}

type traceKey struct {
	Bucket int
	Name   string
}

// Logical access counters form an independent deterministic oracle.
func TestGraderGenericOperationTraces(t *testing.T) {
	for _, capacity := range []int{0, 1, 2, 4} {
		t.Run(fmt.Sprint(capacity), func(t *testing.T) {
			c := subject.New[traceKey, []string](capacity)
			values, access := map[traceKey][]string{}, map[traceKey]int{}
			tick := 0
			keys := []int{0, -7, 12, 0, 42, -7, 3, 12, 0, -7, 9, 42}
			for round := 0; round < 3; round++ {
				for step, raw := range keys {
					key := traceKey{raw % 2, fmt.Sprint(raw)}
					tick++
					if step%4 == 1 {
						got, ok := c.Get(key)
						want, exists := values[key]
						if !reflect.DeepEqual(got, want) || ok != exists {
							t.Fatalf("capacity=%d round=%d step=%d Get(%v)=(%v,%t); want (%v,%t)", capacity, round, step, key, got, ok, want, exists)
						}
						if exists {
							access[key] = tick
						}
					} else {
						value := []string{fmt.Sprint(round), fmt.Sprint(step)}
						c.Set(key, value)
						if capacity > 0 {
							values[key] = value
							access[key] = tick
							if len(values) > capacity {
								oldest := key
								for k, when := range access {
									if when < access[oldest] {
										oldest = k
									}
								}
								delete(values, oldest)
								delete(access, oldest)
							}
						}
					}
					order := []traceKey{}
					for k := range values {
						order = append(order, k)
					}
					sort.Slice(order, func(i, j int) bool { return access[order[i]] < access[order[j]] })
					expected := []item[traceKey, []string]{}
					for _, k := range order {
						expected = append(expected, item[traceKey, []string]{k, values[k]})
					}
					genericEntries(t, c, expected...)
				}
				if round == 1 {
					c.Clear()
					values = map[traceKey][]string{}
					access = map[traceKey]int{}
					genericEntries(t, c)
				}
			}
		})
	}
}
