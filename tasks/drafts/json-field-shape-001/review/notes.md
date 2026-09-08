# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Record has exported Name and Count and private secret. NewRecord stores all three supplied values.
- EncodeRecord returns a JSON object with name always present, count omitted only when zero, and no secret field.
- Use JSON tags on Record fields. Preserve negative counts and exact Unicode/text content; object key order and escaping spelling are unspecified.
- Return encoding errors normally. Do not mutate the Record or read environment state.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{name string;count int}{{"",0},{" λ <x> ",3},{"negative",-2}}{item:=subject.NewRecord(c.name,c.count,"private");data,e:=subject.EncodeRecord(item);if e!=nil{t.Fatal(e)};var got map[string]any;if e=json.Unmarshal(data,&got);e!=nil{t.Fatal(e)};want:=map[string]any{"name":c.name};if c.count!=0{want["count"]=float64(c.count)};if !reflect.DeepEqual(got,want)||item.Name!=c.name||item.Count!=c.count{t.Errorf("JSON got%v want%v item%+v",got,want,item)}}`

Realistic required defect models:

- default-names
- count-always
- omit-empty-name
- hide-count
- wrong-count
- trim-name

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
