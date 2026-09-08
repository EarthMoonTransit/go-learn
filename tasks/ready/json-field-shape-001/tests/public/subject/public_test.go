package subject_test

import (
	"encoding/json"
	"reflect"
	"task.local/json-field-shape-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		name  string
		count int
	}{{"", 0}, {" λ <x> ", 3}, {"negative", -2}} {
		item := subject.NewRecord(c.name, c.count, "private")
		data, e := subject.EncodeRecord(item)
		if e != nil {
			t.Fatal(e)
		}
		var got map[string]any
		if e = json.Unmarshal(data, &got); e != nil {
			t.Fatal(e)
		}
		want := map[string]any{"name": c.name}
		if c.count != 0 {
			want["count"] = float64(c.count)
		}
		if !reflect.DeepEqual(got, want) || item.Name != c.name || item.Count != c.count {
			t.Errorf("JSON got%v want%v item%+v", got, want, item)
		}
	}
}
