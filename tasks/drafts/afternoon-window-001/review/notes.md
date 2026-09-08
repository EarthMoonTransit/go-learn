# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Parse with layout "Monday, January 2, 2006 15:04:05" using time.Parse rules, including its weekday handling and optional fractional seconds.
- Return true only for parsed times at or after 12:00 and before 18:00; return false for other valid times.
- On invalid input return false and the supplied non-nil ErrDate. Preserve its identity.
- The result depends only on text; do not use the current date or machine timezone.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{text string;want bool}{{"Thursday, May 13, 2010 20:32:00",false},{"Friday, March 8, 1974 12:02:02",true},{"Friday, September 9, 2112 11:59:59",false},{"Monday, January 1, 2024 12:00:00",true},{"Monday, January 1, 2024 17:59:59.999",true},{"Monday, January 1, 2024 18:00:00",false}}{got,e:=subject.IsAfternoon(c.text);if got!=c.want||e!=nil{t.Errorf("window%q got%v,%v want%v,nil",c.text,got,e,c.want)}}`
- Case02: `sentinel:=subject.ErrDate;if sentinel==nil{t.Fatal("nil date sentinel")};for _,text:=range []string{"","Monday, February 30, 2024 13:00:00","Monday, January 1, 2024 25:00:00","1/1/2024 13:00:00"}{v,e:=subject.IsAfternoon(text);if v||e!=sentinel||subject.ErrDate!=sentinel{t.Errorf("invalid %q got%v,%v",text,v,e)}}`

Realistic required defect models:

- exclude-noon
- include-eighteen
- or-window
- wrong-layout
- swallow-error
- fresh-error

Retained source behavior: booking-up-for-beauty selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
