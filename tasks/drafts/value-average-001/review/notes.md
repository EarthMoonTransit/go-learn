# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Average is Hits divided by Attempts using floating-point division; zero attempts returns zero.
- Inputs satisfy 0 <= Hits <= Attempts <= 1000000. Label is unrelated to the result.
- Use the declared value receiver: both Sample values and *Sample pointers must expose Average.
- Do not change any field of the receiver. Absolute error up to 0.000000001 is allowed.

Public/reference assertions (actual/expected diagnostics):

- Case01: `s:=subject.Sample{Label:"a",Attempts:4,Hits:1};got:=s.Average();if math.IsNaN(got)||math.Abs(got-0.25)>1e-9{t.Errorf("average got %g want .25",got)};if s!=(subject.Sample{Label:"a",Attempts:4,Hits:1}){t.Errorf("receiver changed: %v",s)}`
- Case02: `if got:=(subject.Sample{}).Average();got!=0{t.Errorf("zero got %g want0",got)}`
- Case03: `s:=subject.Sample{Attempts:3,Hits:2};for _,got:=range []float64{s.Average(),(&s).Average(),(subject.Sample{Attempts:3,Hits:2}).Average()}{if math.IsNaN(got)||math.Abs(got-2.0/3)>1e-9{t.Errorf("value/pointer got %g want2/3",got)}}`

Realistic required defect models:

- integer-division
- reversed-ratio
- zero-one
- percentage
- label-filter

Retained source behavior: ardan:topics/go/language/methods/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
