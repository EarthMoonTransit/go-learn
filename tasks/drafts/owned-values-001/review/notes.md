# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return a non-nil receive-only channel containing a snapshot of values in order, followed by closure. All values and closure are available when this function returns.
- The function owns the sending side and closes it. Use sufficient buffering so no receiver or background producer is needed to complete creation.
- Preserve input storage and keep separate calls independent. Changes to the input after return must not affect queued values.
- Empty input returns an already closed channel. Do not leave a goroutine waiting for a receiver; an unused result must not retain a worker.

Public/reference assertions (actual/expected diagnostics):

- Case01: `input:=[]int{1,0,-2};c:=subject.OwnedValues(input);if c==nil{t.Fatal("nil stream")};if !slices.Equal(input,[]int{1,0,-2}){t.Error("input changed")};input[0]=9;for _,want:=range []int{1,0,-2}{select{case got,ok:=<-c:if !ok||got!=want{t.Errorf("queued got%d,%v want%d,true",got,ok,want)};default:t.Fatal("stream value must be ready at return")}};select{case _,ok:=<-c:if ok{t.Error("stream not closed after values")};default:t.Error("closure must be ready at return")}`
- Case02: `for _,v:=range [][]int{nil,{}}{c:=subject.OwnedValues(v);if c==nil{t.Fatal("nil empty stream")};select{case _,ok:=<-c:if ok{t.Error("empty stream sent value")};default:t.Error("empty stream not closed")}};a:=subject.OwnedValues([]int{3});b:=subject.OwnedValues([]int{4});if a==b{t.Error("streams share channel")};if <-a!=3||<-b!=4{t.Error("streams share values")}`

Realistic required defect models:

- no-close
- skip-zero
- reverse-values
- mutate-input
- nil-empty

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
