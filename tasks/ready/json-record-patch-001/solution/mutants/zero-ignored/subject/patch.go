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
	label, qty := base.Label, base.Quantity
	if patch.Label != nil && *patch.Label != "" {
		label = *patch.Label
	}
	if patch.Quantity != nil && *patch.Quantity != 0 {
		qty = *patch.Quantity
	}
	return model.New(label, qty, base.Secret()), nil
}
func Encode(value model.Record) ([]byte, error) { return json.Marshal(value) }
