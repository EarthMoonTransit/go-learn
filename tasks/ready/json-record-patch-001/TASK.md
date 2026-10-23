# Apply optional JSON fields across a package boundary
## Context
Implement model.Record and subject.
## API
```go
// model
func New(label string, qty int, secret string) Record
func (Record) Values()(string,int)
func (Record) Secret()string
// subject
func Apply(base model.Record,data []byte)(model.Record,error)
func Encode(value model.Record)([]byte,error)
```
## Requirements
- Author Record's Label string and Quantity int exports, private secret, and JSON tags: label always present, qty omitted exactly when zero.
- New stores values; accessors return them. Use exported fields and the private-state accessor across packages.
- Apply accepts objects or null. Absent/null fields preserve base; explicit empty label or zero qty replaces it. Ignore unknown fields.
- Preserve secret and base. Invalid JSON or types return unchanged base and a non-nil error.
- Encode honors tags without leaking secret. Use encoding/json with learner-declared structs; no maps or manual JSON.
## Constraints
- Valid UTF-8, ≤256 input bytes; standard library. JSON spelling/order unspecified; method review required.
