# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- For amount <= Units subtract amount and return nil. Withdrawing exactly the balance is allowed.
- If amount exceeds Units return the original non-nil ErrInsufficient sentinel and leave Units unchanged.
- The receiver is non-nil; initial Units and amount are nonnegative integers <=1000000. Zero withdrawals are valid.
- Update the shared receiver visible through aliases, preserve the sentinel, and use the current balance on repeated calls.

Public/reference assertions (actual/expected diagnostics):

- Case01: `b:=subject.Balance{Units:20};alias:=&b;for _,amount:=range []int{10,10,0}{before:=b.Units;e:=b.Withdraw(amount);if e!=nil||b.Units!=before-amount||alias.Units!=before-amount{t.Errorf("withdraw%d got%d,%v want%d,nil",amount,b.Units,e,before-amount)}}`
- Case02: `b:=subject.Balance{Units:20};sentinel:=subject.ErrInsufficient;if sentinel==nil{t.Fatal("nil sentinel")};e:=b.Withdraw(100);if e!=sentinel||b.Units!=20||subject.ErrInsufficient!=sentinel{t.Errorf("failed got%d,%v want20,original sentinel",b.Units,e)};if e=b.Withdraw(1);e!=nil||b.Units!=19{t.Errorf("after failure got%d,%v want19,nil",b.Units,e)}`
- Case03: `b:=subject.Balance{};sentinel:=subject.ErrInsufficient;if e:=b.Withdraw(1);e!=sentinel||b.Units!=0{t.Errorf("zero balance got%d,%v",b.Units,e)}`

Realistic required defect models:

- strict-boundary
- subtract-on-failure
- fresh-error
- no-subtract
- add-instead
- failure-nil

Retained source behavior: lgwt:pointers/v4/wallet.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
