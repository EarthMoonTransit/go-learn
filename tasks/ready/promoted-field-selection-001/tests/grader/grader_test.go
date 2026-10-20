package subject_test

import (
	subject "task.local/promoted-field-selection-001/subject"
	"testing"
)

func TestIndependentFieldValues(t *testing.T) {
	for i := -17; i < 19; i++ {
		r := subject.Record{Header: subject.Header{Code: i * 3, Name: string(rune('A' + i + 17)), Label: "label"}, Left: subject.Left{Value: i + 20}, Right: subject.Right{Value: i - 40}, Code: i}
		a, b, n, l, q := subject.SelectFields(r)
		if a != i || b != i*3 || n != r.Header.Name || l != i+20 || q != i-40 {
			t.Fatalf("record %+v: got %d %d %q %d %d", r, a, b, n, l, q)
		}
	}
}
