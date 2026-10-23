package subject

import (
	"bytes"
	"encoding/json"
	"task.local/json-record-patch-001/model"
)

type patchFields struct {
	Label    *string `json:"label"`
	Quantity *int    `json:"qty"`
}

func Apply(base model.Record, data []byte) (model.Record, error) {
	var patch patchFields
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&patch); err != nil {
		return base, err
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
