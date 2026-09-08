# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the number of whole minutes in a duration from zero through 24 hours inclusive, truncating a positive remainder.
- Return 0 and the supplied non-nil ErrDuration for negative durations or values above 24 hours.
- Preserve the error sentinel. Use duration units, not a guessed conversion from nanoseconds.
- No wall-clock input, floating-point rounding or local timezone is involved.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{v time.Duration;want int64}{{0,0},{time.Minute-time.Nanosecond,0},{time.Minute,1},{89*time.Second,1},{119*time.Second,1},{2*time.Minute,2},{24*time.Hour,1440}}{n,e:=subject.WholeMinutes(c.v);if n!=c.want||e!=nil{t.Errorf("duration %v got%d,%v want%d,nil",c.v,n,e,c.want)}}`
- Case02: `sentinel:=subject.ErrDuration;if sentinel==nil{t.Fatal("nil duration sentinel")};for _,v:=range []time.Duration{-time.Nanosecond,24*time.Hour+time.Nanosecond,time.Duration(1<<63-1)}{n,e:=subject.WholeMinutes(v);if n!=0||e!=sentinel||subject.ErrDuration!=sentinel{t.Errorf("invalid duration%v got%d,%v want0,original",v,n,e)}}`

Realistic required defect models:

- round-nearest
- seconds-as-minutes
- exclude-day
- negative-accepted
- fresh-error
- swallow-invalid

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
