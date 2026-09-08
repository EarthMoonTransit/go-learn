# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Decode exactly one JSON object with optional name string and count integer fields; missing fields take their zero values.
- Reject unknown fields, null, wrong field types, malformed input and any second JSON value or non-whitespace suffix.
- On rejection return zero Record and the supplied non-nil ErrJSON; preserve the sentinel. Whitespace around one object is allowed.
- Do not close src. Standard encoding/json duplicate-key and field-name matching rules apply; no additional schema validation is required.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{text string;want subject.Record}{{"{}",subject.Record{}},{" 
{\"name\":\" λ \",\"count\":-2}	",subject.Record{Name:" λ ",Count:-2}},{"{\"name\":\"a\",\"name\":\"b\"}",subject.Record{Name:"b"}}}{v,e:=subject.DecodeOne(strings.NewReader(c.text));if e!=nil||v!=c.want{t.Errorf("decode%q got%+v,%v want%+v",c.text,v,e,c.want)}}`
- Case02: `sentinel:=subject.ErrJSON;if sentinel==nil{t.Fatal("nil JSON sentinel")};for _,text:=range []string{"","null","[]","{\"extra\":1}","{\"name\":2}","{\"count\":1.5}","{} {}","{} null","{} trailing","{\"name\":\"x\""}{v,e:=subject.DecodeOne(strings.NewReader(text));if v!=(subject.Record{})||e!=sentinel||subject.ErrJSON!=sentinel{t.Errorf("reject%q got%+v,%v wantzero,original",text,v,e)}}`
- Case03: `r:=&jsonReader{Reader:strings.NewReader("{}")};_,e:=subject.DecodeOne(r);if e!=nil||r.closes!=0{t.Errorf("success closes%d err%v",r.closes,e)};r=&jsonReader{Reader:strings.NewReader("{} x")};_,_=subject.DecodeOne(r);if r.closes!=0{t.Errorf("failure closes%d",r.closes)}`

Realistic required defect models:

- allow-unknown
- ignore-trailing
- null-as-zero
- partial-on-reject
- fresh-error
- close-reader

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
