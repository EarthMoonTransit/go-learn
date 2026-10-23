package subject_test

import (
	"encoding/json"
	"task.local/json-record-patch-001/model"
	"task.local/json-record-patch-001/subject"
	"testing"
)

func TestFinitePatchCombinations(t *testing.T) {
	for _, label := range []string{"", `"quoted"`, "空白"} {
		for _, qty := range []int{-13, 0, 17} {
			base := model.New("old", 44, "preserved")
			data, err := json.Marshal(struct {
				Label    string `json:"label"`
				Quantity int    `json:"qty"`
			}{label, qty})
			if err != nil {
				t.Fatal(err)
			}
			got, err := subject.Apply(base, data)
			if err != nil {
				t.Errorf("finite patch error=%v; want nil", err)
			}
			checkRecord(t, got, label, qty, "preserved")
			checkEncoded(t, got, label, qty, "preserved")
			checkRecord(t, base, "old", 44, "preserved")
		}
	}
}
