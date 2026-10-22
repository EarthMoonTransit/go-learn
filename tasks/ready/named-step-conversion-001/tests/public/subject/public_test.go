package subject_test

import (
	"reflect"
	subject "task.local/named-step-conversion-001/subject"
	"testing"
)

func TestDistinctTypesAndConversion(t *testing.T) {
	for _, c := range [][2]int16{{30, 7}, {-40, 13}, {5, -9}, {0, 0}, {1000, 1000}, {-1000, -1000}} {
		first, delta, result, raw := subject.Observe(c[0], c[1])
		values := []any{first, delta, result}
		names := []string{"reading", "step", "reading"}
		want := []int64{int64(c[0]), int64(c[1]), int64(c[0]) + int64(c[1])}
		for i, v := range values {
			typ := reflect.TypeOf(v)
			if typ == nil || typ.Name() != names[i] || typ.PkgPath() == "" || typ.Kind() != reflect.Int16 {
				t.Fatalf("inputs=%v slot=%d type=%T; want distinct package type %s with int16 underlying", c, i, v, names[i])
			}
			if got := reflect.ValueOf(v).Int(); got != want[i] {
				t.Fatalf("inputs=%v slot=%d value=%d; want %d", c, i, got, want[i])
			}
		}
		if reflect.TypeOf(first) == reflect.TypeOf(delta) || reflect.TypeOf(first) != reflect.TypeOf(result) {
			t.Fatalf("type identities: %T,%T,%T; want reading,distinct step,reading", first, delta, result)
		}
		if raw != c[0]+c[1] {
			t.Fatalf("inputs=%v raw result=%d; want %d", c, raw, c[0]+c[1])
		}
	}
}
