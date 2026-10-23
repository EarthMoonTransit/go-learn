package subject_test

import (
	"reflect"
	"strings"
	"task.local/json-record-patch-001/model"
	"testing"
)

func TestRecordDeclaration(t *testing.T) {
	typ := reflect.TypeOf(model.New("", 0, ""))
	if typ.Kind() != reflect.Struct {
		t.Errorf("Record kind=%v; want struct", typ.Kind())
		return
	}
	for _, want := range []struct {
		name, tag string
		kind      reflect.Kind
	}{{"Label", "label", reflect.String}, {"Quantity", "qty", reflect.Int}} {
		field, ok := typ.FieldByName(want.name)
		if !ok {
			t.Errorf("field %s missing; want exported %v with JSON name %q", want.name, want.kind, want.tag)
			continue
		}
		if field.PkgPath != "" || field.Type.Kind() != want.kind {
			t.Errorf("field %s export/type=(%q,%v); want exported %v", want.name, field.PkgPath, field.Type, want.kind)
		}
		if got := strings.Split(field.Tag.Get("json"), ",")[0]; got != want.tag {
			t.Errorf("field %s JSON name=%q; want %q", want.name, got, want.tag)
		}
	}
	field, ok := typ.FieldByName("secret")
	if !ok {
		t.Error("private secret field missing")
	} else if field.PkgPath == "" || field.Type.Kind() != reflect.String {
		t.Errorf("secret export/type=(%q,%v); want private string", field.PkgPath, field.Type)
	}
}
func TestRecordConstructionAndAccessors(t *testing.T) {
	first := model.New(" padded λ ", -7, " first secret ")
	second := model.New("", 0, "second")
	for i, c := range []struct {
		record model.Record
		label  string
		qty    int
		secret string
	}{{first, " padded λ ", -7, " first secret "}, {second, "", 0, "second"}} {
		label, qty := c.record.Values()
		if label != c.label || qty != c.qty || c.record.Secret() != c.secret {
			t.Errorf("record %d values=(%q,%d,%q); want (%q,%d,%q)", i, label, qty, c.record.Secret(), c.label, c.qty, c.secret)
		}
	}
}
