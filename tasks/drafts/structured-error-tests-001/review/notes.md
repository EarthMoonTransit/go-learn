# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Positive counts return a nil error.
- Negative counts return a non-nil *InvalidCountError with the original Count and Reason "negative count"; zero uses Reason "empty count".
- Error formats "<Count> items are invalid: <Reason>" using decimal Count and exact Reason, including for manually constructed errors.
- Check the concrete pointer type, structured fields and displayed text; a generic error with identical text is insufficient.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if e:=subject.ValidateCount(2);e!=nil{t.Errorf("positive got %v wantnil",e)}`
- Case02: `for _,n:=range []int{-20,-1,0}{e:=subject.ValidateCount(n);typed,ok:=e.(*subject.InvalidCountError);if !ok||typed==nil{t.Fatalf("count%d got %T,%v want nonnil *InvalidCountError",n,e,e)};reason:="negative count";if n==0{reason="empty count"};want:=fmt.Sprintf("%d items are invalid: %s",n,reason);if typed.Count!=n||typed.Reason!=reason||e.Error()!=want{t.Errorf("count%d got %+v,%q want count%d,%q,%q",n,typed,e.Error(),n,reason,want)}}`
- Case03: `e:=&subject.InvalidCountError{Count:7,Reason:"custom reason"};if got:=e.Error();got!="7 items are invalid: custom reason"{t.Errorf("custom Error got %q",got)}`

Realistic required defect models:

- zero-valid
- negative-valid
- erased-count
- generic-error
- same-reason
- fixed-reason-text
- typed-nil-success

Retained source behavior: the-farm selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
