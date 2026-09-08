package subject

import (
	"reflect"
	"testing"
)

func TestFieldContract(t *testing.T) {
	r := NewRecord("x", 2, " exact secret ")
	if r.secret != " exact secret " {
		t.Errorf("private secret got%q want exact value", r.secret)
	}
	typ := reflect.TypeOf(r)
	for _, c := range []struct{ name, tag string }{{"Name", "name"}, {"Count", "count,omitempty"}} {
		f, ok := typ.FieldByName(c.name)
		if !ok || f.Tag.Get("json") != c.tag {
			t.Errorf("field %s JSON tag got%q want%q", c.name, f.Tag.Get("json"), c.tag)
		}
	}
	f, ok := typ.FieldByName("secret")
	if !ok || f.PkgPath == "" {
		t.Error("secret must remain unexported")
	}
}
