# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Input is non-nil, buffered and already closed. Return all queued values in send order, then stop at closed-and-empty.
- Zero and negative values are ordinary data. Closed empty input returns an empty result; nil or allocated empty slices are both valid.
- Do not close the caller-owned input or invent an extra zero after its buffer is drained. Repeated calls on the drained channel return empty.
- The public fixtures fully prepare channels before calling the function; no scheduling or timing assumptions are needed.

Public/reference assertions (actual/expected diagnostics):

- Case01: `c:=make(chan int,4);for _,v:=range []int{3,0,-2,7}{c<-v};close(c);if got:=subject.CollectClosed(c);!slices.Equal(got,[]int{3,0,-2,7}){t.Errorf("drain got%v want[3 0 -2 7]",got)};if got:=subject.CollectClosed(c);len(got)!=0{t.Errorf("already drained got%v",got)}`
- Case02: `c:=make(chan int);close(c);if got:=subject.CollectClosed(c);len(got)!=0{t.Errorf("closed empty got%v",got)};one:=make(chan int,1);one<-0;close(one);if got:=subject.CollectClosed(one);!slices.Equal(got,[]int{0}){t.Errorf("queued zero got%v",got)}`

Realistic required defect models:

- zero-terminates
- skip-negative
- reverse-order
- extra-closed-zero
- drop-first

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
