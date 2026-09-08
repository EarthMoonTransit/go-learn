# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- NewStack returns an empty non-nil stack. The zero value of Stack also works; receivers are non-nil.
- Push stores one value; Pop removes and returns the most recently pushed value with true.
- An empty Pop returns the zero value of T and false without panicking. IsEmpty reflects the current count.
- Support arbitrary types including slices, repeated empty pops and reuse after draining. Separate stack instances are independent.

Public/reference assertions (actual/expected diagnostics):

- Case01: `s:=subject.NewStack[int]();if s==nil{t.Fatal("nil new stack")};if !s.IsEmpty(){t.Error("new nonempty")};s.Push(123);s.Push(456);if s.IsEmpty(){t.Error("populated empty")};for _,want:=range []int{456,123}{got,ok:=s.Pop();if got!=want||!ok{t.Errorf("LIFO got%d,%v want%d,true",got,ok,want)}};for i:=0;i<2;i++{v,ok:=s.Pop();if v!=0||ok||!s.IsEmpty(){t.Errorf("empty got%d,%v",v,ok)}};s.Push(7);v,ok:=s.Pop();if v!=7||!ok{t.Error("reuse failed")}`
- Case02: `var a,b subject.Stack[string];a.Push("a");b.Push("b");a.Push("");for _,want:=range []string{"","a"}{v,ok:=a.Pop();if v!=want||!ok{t.Errorf("first instance got%q,%v",v,ok)}};v,ok:=b.Pop();if v!="b"||!ok{t.Errorf("second instance got%q,%v",v,ok)}`
- Case03: `s:=subject.NewStack[[]int]();if s==nil{t.Fatal("nil stack")};s.Push([]int{1,2});v,ok:=s.Pop();if !ok||!slices.Equal(v,[]int{1,2}){t.Errorf("slice element got%v,%v",v,ok)};v,ok=s.Pop();if v!=nil||ok{t.Errorf("slice zero got%v,%v",v,ok)}`

Realistic required defect models:

- fifo
- peek-only
- false-success
- empty-true
- always-empty
- replace-push
- nil-new

Retained source behavior: lgwt:generics/stack.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
