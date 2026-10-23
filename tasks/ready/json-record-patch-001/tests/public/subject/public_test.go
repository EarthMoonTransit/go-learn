package subject_test

import (
	"encoding/json"
	"fmt"
	"task.local/json-record-patch-001/model"
	"task.local/json-record-patch-001/subject"
	"testing"
)

func checkRecord(t *testing.T, r model.Record, label string, qty int, secret string) {
	t.Helper()
	gotLabel, gotQty := r.Values()
	if gotLabel != label || gotQty != qty || r.Secret() != secret {
		t.Errorf("record result=(%q,%d,%q); want (%q,%d,%q)", gotLabel, gotQty, r.Secret(), label, qty, secret)
	}
}
func checkEncoded(t *testing.T, r model.Record, label string, qty int, secret string) {
	t.Helper()
	data, err := subject.Encode(r)
	if err != nil {
		t.Errorf("Encode error=%v; want nil", err)
		return
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(data, &object); err != nil || object == nil {
		t.Errorf("encoded object=%q error=%v; want object", data, err)
		return
	}
	wantKeys := 1
	if qty != 0 {
		wantKeys++
	}
	if len(object) != wantKeys {
		t.Errorf("JSON keys=%v; want label and qty only when nonzero", object)
	}
	raw, ok := object["label"]
	var gotLabel string
	if !ok || json.Unmarshal(raw, &gotLabel) != nil || gotLabel != label {
		t.Errorf("JSON label=(%q,present=%t); want %q always present", raw, ok, label)
	}
	raw, ok = object["qty"]
	if ok != (qty != 0) {
		t.Errorf("JSON qty presence=%t; want %t for qty=%d", ok, qty != 0, qty)
	} else if ok {
		var gotQty int
		if err := json.Unmarshal(raw, &gotQty); err != nil || gotQty != qty {
			t.Errorf("JSON qty=%q error=%v; want %d", raw, err, qty)
		}
	}
	checkRecord(t, r, label, qty, secret)
}
func TestEncodeDeclaredShape(t *testing.T) {
	cases := []struct {
		label  string
		qty    int
		secret string
	}{{"", 0, "hidden"}, {"  λ <x>  ", -4, " private "}, {"next", 9, "another"}}
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			r := model.New(c.label, c.qty, c.secret)
			checkEncoded(t, r, c.label, c.qty, c.secret)
		})
	}
}
func TestOptionalPatchValues(t *testing.T) {
	const initial = " original λ "
	cases := []struct {
		name, data, label string
		qty               int
	}{
		{"absent", "{}", initial, -3}, {"top-null", "null", initial, -3}, {"null-fields", `{"label":null,"qty":null}`, initial, -3},
		{"empty-label", `{"label":""}`, "", -3}, {"zero-qty", `{"qty":0}`, initial, 0},
		{"both", `{"label":"updated","qty":4}`, "updated", 4}, {"null-with-value", `{"label":null,"qty":2}`, initial, 2},
		{"negative", `{"qty":-8}`, initial, -8}, {"unicode", `{"label":"\u03bb"}`, "λ", -3},
		{"unknown-private", `{"other":9,"secret":"replace","Secret":"again","label":"safe"}`, "safe", -3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			base := model.New(initial, -3, " keep ")
			got, err := subject.Apply(base, []byte(c.data))
			if err != nil {
				t.Errorf("Apply error=%v; want nil for %s", err, c.name)
			}
			checkRecord(t, got, c.label, c.qty, " keep ")
			checkRecord(t, base, initial, -3, " keep ")
		})
	}
}
func TestPatchFailureIsAtomic(t *testing.T) {
	inputs := []string{`{"label":"partial","qty":"bad"}`, `{"qty":0,"label":false}`, `{"label":`, "[]", "1", `"text"`, "true", ""}
	for i, data := range inputs {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			base := model.New("before", 6, "unchanged secret")
			got, err := subject.Apply(base, []byte(data))
			if err == nil {
				t.Errorf("Apply error=nil for invalid input %q; want decoding error", data)
			}
			checkRecord(t, got, "before", 6, "unchanged secret")
			checkRecord(t, base, "before", 6, "unchanged secret")
		})
	}
}
