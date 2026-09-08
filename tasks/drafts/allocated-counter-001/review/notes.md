# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return a non-nil pointer to an integer containing initial.
- initial may be zero, positive or negative.
- The returned integer remains readable and writable after the function returns.
- Each call creates independent storage; updating one result must not change another.

Public/reference assertions (actual/expected diagnostics):

- Case01: `p:=subject.NewCounter(2);if p==nil{t.Fatal("NewCounter(2): got nil want pointer")};if *p!=2{t.Errorf("counter got %d want 2",*p)}`
- Case02: `p:=subject.NewCounter(0);if p==nil{t.Fatal("zero counter: got nil want pointer")};if *p!=0{t.Errorf("counter got %d want 0",*p)}`
- Case03: `p:=subject.NewCounter(-3);if p==nil{t.Fatal("negative counter: got nil want pointer")};if *p!=-3{t.Errorf("counter got %d want -3",*p)}`
- Case04: `a,b:=subject.NewCounter(4),subject.NewCounter(4);if a==nil||b==nil{t.Fatal("got nil counters")};*a=7;if *a!=7||*b!=4{t.Errorf("independent counters got (%d,%d) want (7,4)",*a,*b)}`

Realistic required defect models:

- nil-pointer
- default-zero
- negative-absolute
- shared-storage
- zero-is-nil

Retained source behavior: election-day selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| nonnull initialized storage | 01 | nil-pointer; default-zero |
| zero/negative initial | 02–03 | zero-is-nil; negative-absolute |
| writable independent storage | 04 | shared-storage |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.
