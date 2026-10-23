package subject

import (
	"encoding/json"
	"task.local/json-record-patch-001/model"
)

type patchFields struct {
	Label    *string `json:"label"`
	Quantity *int    `json:"qty"`
}

func Apply(base model.Record, data []byte) (model.Record, error) {
	var patch patchFields
	if err := json.Unmarshal(data, &patch); err != nil {
		return base, err
	}
	var raw map[string]json.RawMessage
	_ = json.Unmarshal(data, &raw)
	if string(raw["label"]) == "null" {
		v := ""
		patch.Label = &v
	}
	if string(raw["qty"]) == "null" {
		v := 0
		patch.Quantity = &v
	}
	label, qty := base.Label, base.Quantity
	if patch.Label != nil {
		label = *patch.Label
	}
	if patch.Quantity != nil {
		qty = *patch.Quantity
	}
	return model.New(label, qty, base.Secret()), nil
}
func Encode(value model.Record) ([]byte, error) { return json.Marshal(value) }
