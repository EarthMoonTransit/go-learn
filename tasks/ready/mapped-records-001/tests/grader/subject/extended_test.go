package subject

import (
	"reflect"
	"testing"
)

type p18Bytes []byte

func TestRepeatedAndSingletonInputs(t *testing.T) {
	for _, input := range [][]p18Reading{{}, {0}, {-7, -7, 4, 4, 0, -7}} {
		calls := 0
		got := mapEntries(input, func(v p18Reading) []byte { calls++; return []byte{byte(v + 30), byte(calls)} })
		want := make([]Entry[p18Reading, []byte], len(input))
		for i, value := range input {
			want[i] = Entry[p18Reading, []byte]{value, []byte{byte(value + 30), byte(i + 1)}}
		}
		// Either nil or allocated empty output is permitted.
		rows, ok := any(got).([]Entry[p18Reading, []byte])
		if !ok {
			t.Fatalf("result type = %T, want []Entry[p18Reading, []byte]", got)
		}
		if len(input) == 0 {
			if len(rows) != 0 {
				t.Errorf("zero inputs produced %d entries, want 0", len(rows))
			}
		} else {
			p18Assert(t, "repeated values", got, want)
		}
		if calls != len(input) {
			t.Errorf("callback calls = %d, want %d", calls, len(input))
		}
	}
}

func TestArrayInputAndNamedBytesOutput(t *testing.T) {
	input := [][2]int{{2, 8}, {0, -3}, {7, 7}}
	before := append([][2]int(nil), input...)
	got := mapEntries(input, func(v [2]int) p18Bytes { return p18Bytes{byte(v[0] + 20), byte(v[1] + 20)} })
	want := []Entry[[2]int, p18Bytes]{{[2]int{2, 8}, p18Bytes{22, 28}}, {[2]int{0, -3}, p18Bytes{20, 17}}, {[2]int{7, 7}, p18Bytes{27, 27}}}
	p18Assert(t, "array to named bytes", got, want)
	if !reflect.DeepEqual(input, before) {
		t.Errorf("array input = %v, want unchanged %v", input, before)
	}
}

func TestIndependentInvocationsAndNilMappedValues(t *testing.T) {
	first := mapEntries([]p18Reading{3, 0}, func(p18Reading) []byte { return nil })
	second := mapEntries([]p18Reading{-2}, func(p18Reading) []byte { return []byte{} })
	p18Assert(t, "nil mapped values", first, []Entry[p18Reading, []byte]{{3, nil}, {0, nil}})
	p18Assert(t, "empty mapped values", second, []Entry[p18Reading, []byte]{{-2, []byte{}}})
}
