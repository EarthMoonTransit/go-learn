package subject_test

import (
	subject "task.local/promoted-field-selection-001/subject"
	"testing"
)

func TestSelectedFields(t *testing.T) {
	cases := []subject.Record{
		{Header: subject.Header{Code: 13, Name: "primary", Label: "other"}, Left: subject.Left{Value: 23}, Right: subject.Right{Value: -31}, Code: 7},
		{Header: subject.Header{Code: -9, Name: "", Label: "not-empty"}, Left: subject.Left{Value: 0}, Right: subject.Right{Value: 11}, Code: 0},
		{},
	}
	for _, r := range cases {
		outer, embedded, name, left, right := subject.SelectFields(r)
		if outer != r.Code {
			t.Errorf("outer Code=%d; want %d", outer, r.Code)
		}
		if embedded != r.Header.Code {
			t.Errorf("Header.Code=%d; want %d", embedded, r.Header.Code)
		}
		if name != r.Header.Name {
			t.Errorf("promoted Name=%q; want %q", name, r.Header.Name)
		}
		if left != r.Left.Value {
			t.Errorf("Left.Value=%d; want %d", left, r.Left.Value)
		}
		if right != r.Right.Value {
			t.Errorf("Right.Value=%d; want %d", right, r.Right.Value)
		}
	}
}
