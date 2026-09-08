# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Create an integer initialized to 20 and a pointer variable referring to it.
- Return the integer pointer and the address of that pointer variable; neither result is nil and the second dereferences to the first.
- Mutating the integer through either path is shared; replacing the pointer variable must not change the original integer.
- Calls allocate independent integer storage and pointer variables. Compare identities without numeric addresses.

Public/reference assertions (actual/expected diagnostics):

- Case01: `p,pp:=subject.PointerChain();if p==nil||pp==nil||*pp!=p||*p!=20||**pp!=20{t.Fatalf("invalid chain p%v pp%v",p,pp)};**pp=7;if *p!=7{t.Error("dereference paths disagree")};other:=33;*pp=&other;if *p!=7||**pp!=33{t.Errorf("pointer replacement changed target got%d,%d",*p,**pp)}`
- Case02: `a,ap:=subject.PointerChain();b,bp:=subject.PointerChain();if a==nil||b==nil||ap==nil||bp==nil{t.Fatal("nil chain")};*a=9;if *b!=20||*ap!=a||*bp!=b||ap==bp{t.Error("calls share pointer or value storage")}`

Realistic required defect models:

- wrong-initial
- nil-pointer-variable
- parallel-values
- shared-target
- shared-pointer-variable

Retained source behavior: ardan:topics/go/language/pointers/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
