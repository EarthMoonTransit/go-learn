# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return a non-nil closure whose initial state is start.
- Each call adds delta to that closure state and returns the updated value.
- Closures returned by separate NewCounter calls have independent state, even when starts are equal.
- Start and accumulated results fit int; negative and zero deltas are valid. Calls are sequential.

Public/reference assertions (actual/expected diagnostics):

- Case01: `f:=subject.NewCounter(10);if f==nil{t.Fatal("nil closure")};if got:=f(2);got!=12{t.Errorf("first got%d want12",got)}`
- Case02: `a:=subject.NewCounter(5);b:=subject.NewCounter(5);c:=subject.NewCounter(-2);if a==nil||b==nil||c==nil{t.Fatal("nil closure")};for _,x:=range []struct{f func(int)int;d,w int}{{a,3,8},{b,0,5},{a,-4,4},{c,1,-1},{b,2,7},{a,0,4}}{if got:=x.f(x.d);got!=x.w{t.Errorf("counter delta%d got%d want%d",x.d,got,x.w)}}`

Realistic required defect models:

- no-persistence
- shared-state
- start-erased
- negative-ignored
- previous-value

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
