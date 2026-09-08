# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Construct a Contact and an anonymous struct with the declared three fields, both storing the inputs exactly.
- Empty names/emails, spaces, Unicode and age zero are valid; age is 0..150.
- Returned values are independent value copies. Changing either result does not affect the other or later calls.
- Do not swap same-typed fields or discard the anonymous result.

Public/reference assertions (actual/expected diagnostics):

- Case01: `a,b:=subject.BuildContacts("Ada","a@example.test",37);if a.Name!="Ada"||a.Email!="a@example.test"||a.Age!=37||b.Name!=a.Name||b.Email!=a.Email||b.Age!=37{t.Errorf("fields got%+v %+v",a,b)}`
- Case02: `a,b:=subject.BuildContacts(" λ ","",0);if a!=(subject.Contact{Name:" λ "})||b.Name!=" λ "||b.Email!=""||b.Age!=0{t.Errorf("zero/exact got%+v %+v",a,b)};a.Name="changed";b.Age=9;c,d:=subject.BuildContacts("", " x ",150);if b.Name!=" λ "||a.Age!=0||c.Email!=" x "||c.Age!=150||d.Email!=c.Email||d.Age!=150{t.Errorf("independence got%+v %+v %+v %+v",a,b,c,d)}`

Realistic required defect models:

- swap-fields
- anonymous-zero
- age-lost
- trim-name
- default-age

Retained source behavior: ardan:topics/go/language/struct_types/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
