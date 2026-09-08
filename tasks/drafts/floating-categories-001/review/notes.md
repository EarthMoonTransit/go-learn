# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return "nan" for any NaN value, "positive-infinity" for positive infinity and "negative-infinity" for negative infinity.
- Return "finite" for every other value, including signed zero, subnormal values and the largest finite float64.
- Classify the original float64 without narrowing it to float32 or comparing to an arbitrary magnitude threshold.
- NaN is unequal to itself; use an appropriate classification operation.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{v float64;want string}{{math.NaN(),"nan"},{math.Float64frombits(0xfff8000000000001),"nan"},{math.Inf(1),"positive-infinity"},{math.Inf(-1),"negative-infinity"}}{if got:=subject.FloatKind(c.v);got!=c.want{t.Errorf("kind got%q want%q",got,c.want)}}`
- Case02: `for _,v:=range []float64{0,math.Copysign(0,-1),-1,1,math.SmallestNonzeroFloat64,-math.SmallestNonzeroFloat64,math.MaxFloat64,-math.MaxFloat64}{if got:=subject.FloatKind(v);got!="finite"{t.Errorf("finite %g got%q",v,got)}}`

Realistic required defect models:

- nan-equality
- infinity-sign-lost
- negative-missed
- narrow-input
- magnitude-threshold

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
