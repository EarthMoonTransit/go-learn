# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return one Entry for every input key with its exact associated value. Output order is unspecified.
- Do not omit zero, negative or empty-key entries; do not duplicate entries.
- Preserve the input map. Changing the returned entries must not change input values.
- Nil or empty maps return zero entries; keys may contain Unicode and spaces.

Public/reference assertions (actual/expected diagnostics):

- Case01: `got:=subject.Entries(map[string]int{"a":1});if len(got)!=1||got[0]!=(subject.Entry{Key:"a",Value:1}){t.Errorf("single got%v",got)}`
- Case02: `in:=map[string]int{"a":1,"":0,"λ":-2," x ":7,"last":9};before:=maps.Clone(in);got:=subject.Entries(in);if len(got)!=len(in){t.Fatalf("entries length got%d want%d",len(got),len(in))};seen:=map[string]bool{};for _,e:=range got{v,ok:=in[e.Key];if !ok||seen[e.Key]||v!=e.Value{t.Errorf("missing/duplicate/wrong entry%v",e)};seen[e.Key]=true};if !maps.Equal(in,before){t.Errorf("input changed %v",in)};got[0].Value=99;if !maps.Equal(in,before){t.Error("result shares mutable state")}`
- Case03: `for _,in:=range []map[string]int{nil,{}}{if got:=subject.Entries(in);len(got)!=0{t.Errorf("empty got%v",got)}}`

Realistic required defect models:

- first-only
- positive-only
- empty-key-ignored
- erase-values
- destructive
- duplicate-entry
- empty-placeholder

Retained source behavior: ardan:topics/go/language/maps/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
