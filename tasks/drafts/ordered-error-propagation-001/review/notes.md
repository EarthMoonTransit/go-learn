# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Call Amount with count, then Factor, each exactly once on success; return Amount times Factor divided by count.
- On the first error stop immediately and return zero with that exact error; do not call a later dependency.
- The calculator is non-nil, count is 1..1000, successful values are finite with magnitude <=1000. Allow absolute error 0.000000001 for numeric results.
- Preserve error identity even when the dependency also returns a nonzero numeric value.

Public/reference assertions (actual/expected diagnostics):

- Case01: `c:=&calculator{amount:60.5,factor:1.3};n,e:=subject.Divide(c,5);if e!=nil||math.IsNaN(n)||math.Abs(n-15.73)>1e-9||!slices.Equal(c.calls,[]string{"amount","factor"})||c.count!=5{t.Errorf("divide got %g,%v calls%v count%d",n,e,c.calls,c.count)}`
- Case02: `sentinel:=errors.New("amount failed");c:=&calculator{amount:99,amountErr:sentinel,factor:2};n,e:=subject.Divide(c,4);if e!=sentinel||math.IsNaN(n)||math.Abs(n)>1e-9||!slices.Equal(c.calls,[]string{"amount"}){t.Errorf("first error got %g,%v calls%v",n,e,c.calls)}`
- Case03: `sentinel:=errors.New("factor failed");c:=&calculator{amount:12,factor:99,factorErr:sentinel};n,e:=subject.Divide(c,3);if e!=sentinel||math.IsNaN(n)||math.Abs(n)>1e-9||!slices.Equal(c.calls,[]string{"amount","factor"}){t.Errorf("second error got %g,%v calls%v",n,e,c.calls)}`

Realistic required defect models:

- ignore-amount-error
- ignore-factor-error
- fresh-error
- drop-division
- wrong-count
- double-amount

Retained source behavior: the-farm selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
