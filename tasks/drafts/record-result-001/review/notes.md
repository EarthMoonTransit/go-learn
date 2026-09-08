# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return a non-nil pointer to a Record containing Label and Count exactly.
- Return a nil error for every valid input; negative counts and empty labels are valid.
- Each call returns independent storage, even for identical inputs.
- Changing a returned record must not affect earlier or later results.

Public/reference assertions (actual/expected diagnostics):

- Case01: `r,e:=subject.NewRecord("data",7);if e!=nil||r==nil||*r!=(subject.Record{Label:"data",Count:7}){t.Fatalf("record got%v,%v",r,e)}`
- Case02: `a,e:=subject.NewRecord("",-1);b,f:=subject.NewRecord("",-1);if e!=nil||f!=nil||a==nil||b==nil||*a!=(subject.Record{Count:-1})||*b!=(subject.Record{Count:-1}){t.Fatalf("empty negative got%v,%v %v,%v",a,e,b,f)};a.Count=99;a.Label="x";if *b!=(subject.Record{Count:-1}){t.Errorf("results share storage %v",b)};c,e:=subject.NewRecord("",-1);if e!=nil||c==nil||*c!=(subject.Record{Count:-1})||a.Count!=99{t.Errorf("later changed state a%v c%v err%v",a,c,e)}`

Realistic required defect models:

- nil-pointer
- count-erased
- negative-rejected
- success-error
- shared-record

Retained source behavior: ardan:topics/go/language/functions/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
