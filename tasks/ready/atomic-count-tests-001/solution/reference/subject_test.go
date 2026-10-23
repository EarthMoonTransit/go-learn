package subject_test

import (
	"maps"
	"slices"
	subject "task.local/atomic-count-tests-001/subject"
	"testing"
)

var originalCountError = subject.ErrCount

type countCase struct {
	name    string
	seed    map[string]int
	deltas  []subject.Delta
	want    map[string]int
	wantErr error
}

func checkCountOutcome(t *testing.T, counts, want map[string]int, err, wantErr error, deltas, before []subject.Delta) {
	t.Helper()
	if err != wantErr {
		t.Errorf("update error=%v; want exact %v", err, wantErr)
	}
	if !maps.Equal(counts, want) {
		t.Errorf("updated counts=%v; want %v", counts, want)
	}
	if !slices.Equal(deltas, before) {
		t.Errorf("delta input=%v; want unchanged %v", deltas, before)
	}
	if subject.ErrCount != originalCountError || originalCountError == nil {
		t.Errorf("count sentinel=%v; want original nonnil %v", subject.ErrCount, originalCountError)
	}
}
func TestAtomicCountPartitions(t *testing.T) {
	cases := []countCase{
		{"empty-keeps-existing", map[string]int{"keep": 4}, nil, map[string]int{"keep": 4}, nil},
		{"existing-and-unrelated", map[string]int{"n": 6, "keep": 9}, []subject.Delta{{"n", -2}, {"n", 3}}, map[string]int{"n": 7, "keep": 9}, nil},
		{"missing-positive", map[string]int{"keep": 8}, []subject.Delta{{"new", 3}}, map[string]int{"keep": 8, "new": 3}, nil},
		{"repeated-key", map[string]int{"n": 2}, []subject.Delta{{"n", 3}, {"n", 2}}, map[string]int{"n": 7}, nil},
		{"existing-reaches-zero", map[string]int{"n": 2}, []subject.Delta{{"n", -2}}, map[string]int{"n": 0}, nil},
		{"missing-becomes-present-zero", map[string]int{"keep": 2}, []subject.Delta{{"new", 0}}, map[string]int{"keep": 2, "new": 0}, nil},
		{"first-negative", map[string]int{"n": 1, "keep": 7}, []subject.Delta{{"n", -2}}, map[string]int{"n": 1, "keep": 7}, originalCountError},
		{"late-rejection-rolls-back", map[string]int{"a": 5, "b": 1, "keep": 8}, []subject.Delta{{"a", 1}, {"b", -2}}, map[string]int{"a": 5, "b": 1, "keep": 8}, originalCountError},
		{"negative-then-recovery", map[string]int{"n": 1}, []subject.Delta{{"n", -2}, {"n", 2}}, map[string]int{"n": 1}, originalCountError},
		{"positive-before-subtraction", map[string]int{"n": 0}, []subject.Delta{{"n", 3}, {"n", -3}}, map[string]int{"n": 0}, nil},
		{"amount-extremes", map[string]int{"n": 100}, []subject.Delta{{"n", 10}, {"n", -10}}, map[string]int{"n": 100}, nil},
		{"missing-negative", map[string]int{"keep": 8}, []subject.Delta{{"new", -1}}, map[string]int{"keep": 8}, originalCountError},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			counts := make(map[string]int, len(c.seed))
			for k, v := range c.seed {
				counts[k] = v
			}
			deltas := slices.Clone(c.deltas)
			before := slices.Clone(deltas)
			err := subject.ApplyCounts(counts, deltas)
			checkCountOutcome(t, counts, c.want, err, c.wantErr, deltas, before)
		})
	}
}
