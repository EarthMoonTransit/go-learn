# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Perform one nonblocking select: a received input value returns that value and value; a closed empty input returns zero and closed; ready cancellation returns zero and canceled; otherwise return zero and pending. Result labels are literal strings.
- Nil channels disable their cases. Do not close or send to input channels. Consume at most one selected input value.
- If input and cancellation are both ready, either ready result is valid; no priority is promised. The default result is only valid when no case is ready.
- Buffered zeros and negative numbers remain values. Cancellation is signaled by a closed or readable channel.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if n,s:=subject.TryReceive(nil,nil);n!=0||s!="pending"{t.Errorf("disabled got%d,%s",n,s)};open:=make(chan int);if n,s:=subject.TryReceive(open,nil);n!=0||s!="pending"{t.Errorf("unready got%d,%s",n,s)};close(open);if n,s:=subject.TryReceive(open,nil);n!=0||s!="closed"{t.Errorf("closed got%d,%s",n,s)}`
- Case02: `c:=make(chan int,2);c<-0;c<- -3;close(c);for _,want:=range []int{0,-3}{if n,s:=subject.TryReceive(c,nil);n!=want||s!="value"{t.Errorf("value got%d,%s want%d,value",n,s,want)}};stop:=make(chan struct{});close(stop);if n,s:=subject.TryReceive(nil,stop);n!=0||s!="canceled"{t.Errorf("cancel got%d,%s",n,s)}`
- Case03: `c:=make(chan int,1);c<-7;stop:=make(chan struct{});close(stop);n,s:=subject.TryReceive(c,stop);if !(s=="value"&&n==7||s=="canceled"&&n==0){t.Errorf("both ready got%d,%s",n,s)};if s=="canceled"&&len(c)!=1{t.Error("unselected input consumed")};if s=="value"&&len(c)!=0{t.Error("selected input not consumed")}`

Realistic required defect models:

- zero-is-closed
- ignore-cancel
- wrong-default
- absolute-value
- closed-as-value

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
