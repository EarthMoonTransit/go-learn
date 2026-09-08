# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Add amount to the receiver Units, visible through every alias of the same pointer.
- Repeated calls accumulate from the current balance; separate Balance values remain independent.
- Receiver is non-nil, initial Units and amounts are nonnegative and sums fit int. Zero deposits leave Units unchanged.

Public/reference assertions (actual/expected diagnostics):

- Case01: `b:=subject.Balance{};b.Deposit(10);if b.Units!=10{t.Errorf("first got%d want10",b.Units)}`
- Case02: `b:=subject.Balance{Units:20};alias:=&b;other:=subject.Balance{Units:20};for _,amount:=range []int{3,0,7}{before:=b.Units;b.Deposit(amount);if b.Units!=before+amount||alias.Units!=before+amount||other.Units!=20{t.Errorf("deposit%d got%d alias%d other%d",amount,b.Units,alias.Units,other.Units)}};other.Deposit(1);if other.Units!=21||b.Units!=30{t.Errorf("independent got%d,%d",b.Units,other.Units)}`

Realistic required defect models:

- replace-balance
- no-update
- subtract
- zero-reset
- shared-total

Retained source behavior: lgwt:pointers/v4/wallet.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
