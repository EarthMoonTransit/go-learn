package subject

import (
	"reflect"
	"testing"
)

type p18Reading int
type p18Readings []p18Reading
type p18Tag struct{ Code int }

func p18Assert(t *testing.T, label string, got, want any) {
	t.Helper()
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Fatalf("%s result type = %T, want %T; retain input and output types", label, got, want)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s entries = %#v, want %#v", label, got, want)
	}
}

func TestMappingValuesAndCallOrder(t *testing.T) {
	values := p18Readings{9, -4, 0}
	before := append(p18Readings(nil), values...)
	calls := []p18Reading{}
	got := mapEntries(values, func(value p18Reading) []byte {
		calls = append(calls, value)
		return []byte{byte(value + 10), byte(len(calls))}
	})
	want := []Entry[p18Reading, []byte]{{9, []byte{19, 1}}, {-4, []byte{6, 2}}, {0, []byte{10, 3}}}
	p18Assert(t, "named reading to bytes", got, want)
	if !reflect.DeepEqual(calls, []p18Reading{9, -4, 0}) {
		t.Errorf("callback inputs = %v, want [9 -4 0] exactly once in input order", calls)
	}
	if !reflect.DeepEqual(values, before) {
		t.Errorf("input after mapping = %v, want unchanged %v", values, before)
	}
}

func TestIndependentStructAndStringTypes(t *testing.T) {
	values := []p18Tag{{Code: 12}, {Code: -3}, {Code: 0}}
	got := mapEntries(values, func(value p18Tag) string {
		if value.Code < 0 {
			return "low"
		}
		return "set"
	})
	p18Assert(t, "struct to text", got, []Entry[p18Tag, string]{{p18Tag{12}, "set"}, {p18Tag{-3}, "low"}, {p18Tag{0}, "set"}})
}

func TestEmptyInputsDoNotCallMapper(t *testing.T) {
	for _, c := range []struct {
		name  string
		input []p18Reading
	}{{"nil", nil}, {"empty", []p18Reading{}}} {
		t.Run(c.name, func(t *testing.T) {
			calls := 0
			got := mapEntries(c.input, func(p18Reading) []byte { calls++; return []byte{99} })
			rows, ok := any(got).([]Entry[p18Reading, []byte])
			if !ok {
				t.Fatalf("empty result type = %T, want []Entry[p18Reading, []byte]", got)
			}
			if len(rows) != 0 {
				t.Errorf("empty input entries = %v, want zero entries", rows)
			}
			if calls != 0 {
				t.Errorf("empty input callback calls = %d, want 0", calls)
			}
		})
	}
}

func TestInputSliceIsNotOverwritten(t *testing.T) {
	values := []p18Reading{31, -8, 5}
	before := append([]p18Reading(nil), values...)
	_ = mapEntries(values, func(v p18Reading) []byte { return []byte{byte(v + 10)} })
	if !reflect.DeepEqual(values, before) {
		t.Errorf("input after mapping = %v, want unchanged %v", values, before)
	}
}
