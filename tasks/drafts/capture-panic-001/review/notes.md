# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Call work exactly once in the current goroutine. For a normal return, preserve its exact error and return panicked=false.
- For a panic carrying a string or int, recover at this boundary and return a non-nil error with text "panic: <value>" and panicked=true.
- Callback panic values are limited to strings and ints; empty strings and zero are valid panic values.
- Do not rerun work or leave background work. Separate calls are independent.

Public/reference assertions (actual/expected diagnostics):

- Case01: `calls:=0;sentinel:=errors.New("work failed");for _,cause:=range []error{nil,sentinel}{err,p:=subject.Capture(func()error{calls++;return cause});if err!=cause||p{t.Errorf("normal got%v,%v want%v,false",err,p,cause)}};if calls!=2{t.Errorf("normal calls%d want2",calls)}`
- Case02: `for _,v:=range []any{"broken","",7,0}{calls:=0;err,p:=subject.Capture(func()error{calls++;panic(v)});want:=fmt.Sprintf("panic: %v",v);if err==nil||err.Error()!=want||!p||calls!=1{t.Errorf("panic got%v,%v calls%d want%q,true,1",err,p,calls,want)}}`

Realistic required defect models:

- no-recover
- text-without-prefix
- false-panic
- normal-error-lost
- work-twice

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
