package subject_test

import (
	"reflect"
	"task.local/named-block-001/subject"
	"testing"
)

func TestNamedArray(t *testing.T) {
	original, _ := subject.Observe()
	typ := reflect.TypeOf(original)
	if typ.Name() != "block" || typ.Kind() != reflect.Array || typ.Len() != 3 {
		t.Fatalf("type=%v, want a defined block with underlying [3]int", typ)
	}
}
