# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return (0,false) for a nil Coder interface or an interface holding a nil *Result pointer; do not call its method.
- For a non-nil *Result return its Code and true, including zero and negative codes.
- For any other implementing type call CodeValue once and return its result with true. Such values have safe method receivers.
- Preserve the pointed Result. An interface holding a typed nil pointer is not itself nil.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,v:=range []subject.Coder{nil,(*subject.Result)(nil)}{n,ok:=subject.ReadCode(v);if n!=0||ok{t.Errorf("nil %T got%d,%v",v,n,ok)}}`
- Case02: `for _,code:=range []int{0,-3,9}{r:=&subject.Result{Code:code};n,ok:=subject.ReadCode(r);if n!=code||!ok||r.Code!=code{t.Errorf("result got%d,%v stored%d want%d,true",n,ok,r.Code,code)}}`
- Case03: `c:=&otherCoder{code:4};n,ok:=subject.ReadCode(c);if n!=4||!ok||c.calls!=1{t.Errorf("other got%d,%v calls%d",n,ok,c.calls)}`

Realistic required defect models:

- interface-nil-only
- all-result-invalid
- zero-absent
- other-rejected
- double-call

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
