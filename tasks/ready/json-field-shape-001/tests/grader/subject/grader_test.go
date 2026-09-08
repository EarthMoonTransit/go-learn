package subject_test

import (
	"encoding/json"
	"task.local/json-field-shape-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	data, e := subject.EncodeRecord(subject.NewRecord("x", 1, "hidden"))
	var got map[string]any
	if e != nil || json.Unmarshal(data, &got) != nil || got["name"] != "x" || len(got) != 2 {
		t.Errorf("object got%s,%v", data, e)
	}
}
