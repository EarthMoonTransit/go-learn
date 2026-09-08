# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return nil for a nil interface, int for dynamic type int, string for dynamic type string, and other for every other type. These results are literal strings.
- NamedInt is a distinct named type and returns other. An alias of int remains int.
- A typed nil pointer or slice is not a nil interface and returns other.
- The stored value does not change classification; include zero and empty values.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.TypeLabel(3);got!="int"{t.Errorf("int got%q",got)}`
- Case02: `type alias=int;for _,c:=range []struct{v any;want string}{{nil,"nil"},{0,"int"},{alias(9),"int"},{"","string"},{"int","string"},{subject.NamedInt(3),"other"},{int64(3),"other"},{3.0,"other"},{true,"other"},{(*int)(nil),"other"},{[]int(nil),"other"},{struct{}{},"other"}}{if got:=subject.TypeLabel(c.v);got!=c.want{t.Errorf("dynamic %T:%v got%q want%q",c.v,c.v,got,c.want)}}`

Realistic required defect models:

- named-is-int
- wide-is-int
- typed-nil-is-nil
- empty-string-other
- default-string
- zero-int-other

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
