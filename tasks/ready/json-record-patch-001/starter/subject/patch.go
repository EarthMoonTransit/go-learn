package subject

import "task.local/json-record-patch-001/model"

func Apply(base model.Record, data []byte) (model.Record, error) { return base, nil }
func Encode(value model.Record) ([]byte, error)                  { return nil, nil }
