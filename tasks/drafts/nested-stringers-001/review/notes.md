# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Temperature units format as "°C" or "°F"; speed units as "km/h" or "mph". Only declared units are supplied.
- Temperature and Speed format their integer followed by one space and the formatted unit.
- Reading formats "<Location>: <Temperature>, Wind <Direction> at <Speed>, <Humidity>% Humidity".
- String and fmt.Sprint produce identical results, preserving text, zero and negative integers.
- Implement all five value methods without recursive receiver formatting.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{s fmt.Stringer;want string}{{subject.Celsius,"°C"},{subject.Fahrenheit,"°F"},{subject.KmPerHour,"km/h"},{subject.MilesPerHour,"mph"},{subject.Temperature{Degree:21,Unit:subject.Celsius},"21 °C"},{subject.Temperature{Degree:-17,Unit:subject.Fahrenheit},"-17 °F"},{subject.Speed{Magnitude:18,Unit:subject.KmPerHour},"18 km/h"},{subject.Speed{Magnitude:0,Unit:subject.MilesPerHour},"0 mph"}}{if got:=c.s.String();got!=c.want{t.Errorf("String got %q want %q",got,c.want)};if got:=fmt.Sprint(c.s);got!=c.want{t.Errorf("Sprint got %q want %q",got,c.want)}}`
- Case02: `r:=subject.Reading{Location:"Athens",Temperature:subject.Temperature{Degree:21,Unit:subject.Celsius},Direction:"N",Speed:subject.Speed{Magnitude:16,Unit:subject.KmPerHour},Humidity:63};want:="Athens: 21 °C, Wind N at 16 km/h, 63% Humidity";if got:=r.String();got!=want{t.Errorf("reading got %q want %q",got,want)};if got:=fmt.Sprint(r);got!=want{t.Errorf("formatted got %q want %q",got,want)}`
- Case03: `r:=subject.Reading{Location:" x ",Temperature:subject.Temperature{Degree:0,Unit:subject.Fahrenheit},Direction:"",Speed:subject.Speed{Magnitude:-2,Unit:subject.MilesPerHour},Humidity:0};want:=" x : 0 °F, Wind  at -2 mph, 0% Humidity";if got:=fmt.Sprint(r);got!=want{t.Errorf("boundary got %q want %q",got,want)}`

Realistic required defect models:

- celsius-always
- speed-unit-always
- no-unit-space
- absolute-temperature
- no-humidity-percent
- direction-erased
- trim-location

Retained source behavior: meteorology selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
