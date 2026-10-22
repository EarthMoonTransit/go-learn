package subject_test

import (
	"reflect"
	subject "task.local/named-step-conversion-001/subject"
	"testing"
)

func TestConversionSmallGrid(t *testing.T) {
	for raw := -17; raw <= 17; raw++ {
		for delta := -9; delta <= 9; delta++ {
			a, b, c, n := subject.Observe(int16(raw), int16(delta))
			for _, v := range []any{a, b, c} {
				if typ := reflect.TypeOf(v); typ == nil || typ.Kind() != reflect.Int16 {
					t.Fatalf("observed type=%T; want defined int16 type", v)
				}
			}
			if reflect.ValueOf(a).Int() != int64(raw) || reflect.ValueOf(b).Int() != int64(delta) || reflect.ValueOf(c).Int() != int64(raw+delta) || n != int16(raw+delta) {
				t.Fatalf("raw=%d delta=%d: got %v,%v,%v,%d; want %d,%d,%d,%d", raw, delta, a, b, c, n, raw, delta, raw+delta, raw+delta)
			}
		}
	}
}
