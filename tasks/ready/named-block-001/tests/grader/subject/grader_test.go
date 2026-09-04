package subject_test

import (
	"reflect"
	"task.local/named-block-001/subject"
	"testing"
)

func TestIdentityAndValueCopy(t *testing.T) {
	original, modified := subject.Observe()
	for _, v := range []any{original, modified} {
		typ := reflect.TypeOf(v)
		if typ.Name() != "block" || typ.PkgPath() != "task.local/named-block-001/subject" || typ.Kind() != reflect.Array {
			t.Fatal("incorrect named array identity")
		}
		if typ.Len() != 3 || typ.Elem() != reflect.TypeOf(int(0)) {
			t.Fatal("incorrect underlying array type")
		}
	}
	for i, want := range []int64{2, 4, 6} {
		if reflect.ValueOf(original).Index(i).Int() != want {
			t.Fatal("original value changed")
		}
	}
	for i, want := range []int64{99, 4, 6} {
		if reflect.ValueOf(modified).Index(i).Int() != want {
			t.Fatal("incorrect copied value")
		}
	}
}
