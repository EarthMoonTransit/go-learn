package subject

import (
	"encoding/json"
	"task.local/json-record-patch-001/model"
)

type patchFields struct {
	Secret   *string `json:"secret"`
	Label    *string `json:"label"`
	Quantity *int    `json:"qty"`
}

func Apply(base model.Record, data []byte) (model.Record, error) {
	var patch patchFields
	if err := json.Unmarshal(data, &patch); err != nil {
		return base, err
	}
	label, qty := base.Label, base.Quantity
	if patch.Label != nil {
		label = *patch.Label
	}
	if patch.Quantity != nil {
		qty = *patch.Quantity
	}
	secret := base.Secret()
	if patch.Secret != nil {
		secret = *patch.Secret
	}
	return model.New(label, qty, secret), nil
}
func Encode(value model.Record) ([]byte, error) { return json.Marshal(value) }
