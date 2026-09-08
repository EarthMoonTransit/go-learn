# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- ZeroValues returns a zero-valued string, int and bool through interfaces, in that order; none is a nil interface.
- LiteralValues returns string "go", int 42, bool true and float32 converted from literal 3.14.
- Preserve concrete dynamic types: for example, float64(42) is not an int even though numerically equal.
- Calls are deterministic and independent; no input or output is read from the environment.

Public/reference assertions (actual/expected diagnostics):

- Case01: `a,b,c:=subject.ZeroValues();s,okS:=a.(string);n,okN:=b.(int);v,okV:=c.(bool);if !okS||!okN||!okV||s!=""||n!=0||v{t.Errorf("zero got%T:%v %T:%v %T:%v",a,a,b,b,c,c)}`
- Case02: `a,b,c,d:=subject.LiteralValues();s,okS:=a.(string);n,okN:=b.(int);v,okV:=c.(bool);if !okS||!okN||!okV||s!="go"||n!=42||!v||d!=float32(3.14){t.Errorf("literal got%T:%v %T:%v %T:%v float%g",a,a,b,b,c,c,d)}`

Realistic required defect models:

- nil-zero
- floating-int
- false-literal
- zero-not-zero
- truncated-pi

Retained source behavior: ardan:topics/go/language/variables/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
