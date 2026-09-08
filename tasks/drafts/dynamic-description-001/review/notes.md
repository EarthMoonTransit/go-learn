# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Dynamic int and finite float64 values format as "number " plus one decimal place using fmt formatting.
- NumberBox formats "box " plus its Number result with one decimal place.
- FancyNumberBox formats "fancy " plus one decimal place: parse Text only for exact FancyNumber values; otherwise use zero. Invalid or overflowing integers use zero.
- NumberBox takes precedence when both interfaces match. Nil and unsupported types return "unsupported".
- Interface method receivers are non-nil. Do not call Value on unsupported FancyNumberBox implementations.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,c:=range []struct{v any;want string}{{42,"number 42.0"},{7.11,"number 7.1"},{-3.2,"number -3.2"},{numberHolder{16},"box 16.0"},{subject.FancyNumber{Text:"16"},"fancy 16.0"},{subject.FancyNumber{Text:"bad"},"fancy 0.0"},{otherFancy{},"fancy 0.0"},{bothBox{},"box 9.0"},{nil,"unsupported"},{"42","unsupported"},{int64(42),"unsupported"}}{if got:=subject.Describe(c.v);got!=c.want{t.Errorf("Describe %T got %q want %q",c.v,got,c.want)}}`

Realistic required defect models:

- int-unsupported
- float-integer
- no-box
- all-fancy-values
- nil-number
- default-number

Retained source behavior: sorting-room selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
