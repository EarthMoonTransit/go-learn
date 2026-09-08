# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the first *Detail found by errors.As and true, preserving pointer identity.
- Search wrapped and joined errors in standard traversal order; respect custom As methods.
- Return nil,false when no match exists, including a nil error. Matching a typed-nil *Detail returns nil,true.
- Do not synthesize a new Detail from its fields or message. Error trees are finite.

Public/reference assertions (actual/expected diagnostics):

- Case01: `d:=&subject.Detail{Code:7,Message:"bad"};got,ok:=subject.FindDetail(d);if got!=d||!ok{t.Errorf("direct got%v,%v",got,ok)}`
- Case02: `a:=&subject.Detail{Code:1,Message:"a"};b:=&subject.Detail{Code:2,Message:"b"};wrapped:=fmt.Errorf("outer: %w",a);for _,e:=range []error{wrapped,errors.Join(errors.New("other"),wrapped,b)}{got,ok:=subject.FindDetail(e);if got!=a||!ok{t.Errorf("first cause got%v,%v",got,ok)}}`
- Case03: `for _,e:=range []error{nil,errors.New("bad")}{d,ok:=subject.FindDetail(e);if d!=nil||ok{t.Errorf("miss got%v,%v",d,ok)}};var typed *subject.Detail;d,ok:=subject.FindDetail(typed);if d!=nil||!ok{t.Errorf("typed nil got%v,%v",d,ok)}`
- Case04: `want:=&subject.Detail{Code:9,Message:"custom"};got,ok:=subject.FindDetail(&projectedError{detail:want});if got!=want||!ok{t.Errorf("custom As got%v,%v",got,ok)}`

Realistic required defect models:

- direct-only
- copy-detail
- false-found
- nil-is-miss
- fabricate-miss

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
