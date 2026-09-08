# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Parse text with Go layout "1/2/2006 15:04:05", retaining time.Parse acceptance rules, including optional fractional seconds.
- Successful values use UTC and preserve date, hour, minute, second and fractional second. Do not read the current time or local timezone.
- On any parse error return the zero time and supplied non-nil ErrDate. Preserve that sentinel.
- Month/day order follows the layout; invalid calendar dates and trailing text are errors.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{text string;want time.Time}{{"7/13/2020 20:32:00",time.Date(2020,7,13,20,32,0,0,time.UTC)},{"11/28/1984 2:02:02",time.Date(1984,11,28,2,2,2,0,time.UTC)},{"2/29/2112 11:59:59",time.Date(2112,2,29,11,59,59,0,time.UTC)},{"3/4/2024 05:06:07.125",time.Date(2024,3,4,5,6,7,125000000,time.UTC)}}{v,e:=subject.ParseTimestamp(c.text);if e!=nil||!v.Equal(c.want)||v.Location()!=time.UTC{t.Errorf("parse%q got%v,%v want%v UTC",c.text,v,e,c.want)}}`
- Case02: `sentinel:=subject.ErrDate;if sentinel==nil{t.Fatal("nil date sentinel")};for _,text:=range []string{"","2/29/2023 00:00:00","13/1/2024 00:00:00","1/1/2024 24:00:00","1/1/2024 00:00:00 extra"}{v,e:=subject.ParseTimestamp(text);if !v.IsZero()||e!=sentinel||subject.ErrDate!=sentinel{t.Errorf("invalid%q got%v,%v wantzero,original",text,v,e)}}`

Realistic required defect models:

- day-month
- discard-time
- discard-fraction
- swallow-error
- fresh-error
- fixed-zone

Retained source behavior: booking-up-for-beauty selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
