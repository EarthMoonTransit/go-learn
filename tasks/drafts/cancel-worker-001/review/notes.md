# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Derive a cancelable child from the non-nil parent. Return that child, its CancelFunc and a non-nil worker completion channel.
- The worker waits for child cancellation, then closes completion without sending values. Explicit child cancellation and parent cancellation must both finish it.
- Canceling the child must not cancel its parent. Repeated calls to cancel are safe; an already canceled parent produces a canceled child and a finishing worker.
- Preserve context values and parent cancellation cause behavior through standard context.WithCancel derivation. Tests control cancellation and use synctest; no real sleep is required.

Public/reference assertions (actual/expected diagnostics):

- Case01: `synctest.Test(t,func(t *testing.T){parent,parentCancel:=context.WithCancel(context.Background());defer parentCancel();child,cancel,done:=subject.StartCancelable(parent);if child==nil||cancel==nil||done==nil{t.Fatal("nil child/cancel/completion")};defer cancel();synctest.Wait();if child.Err()!=nil{t.Error("child canceled early")};select{case <-done:t.Error("worker completed before cancel");default:};cancel();cancel();awaitCanceled(t,done);if child.Err()!=context.Canceled||parent.Err()!=nil{t.Errorf("child cancellation child%v parent%v",child.Err(),parent.Err())}})`
- Case02: `synctest.Test(t,func(t *testing.T){type key struct{};parent,parentCancel:=context.WithCancelCause(context.WithValue(context.Background(),key{},"value"));cause:=errors.New("parent stopped");defer parentCancel(cause);child,cancel,done:=subject.StartCancelable(parent);if child==nil||cancel==nil||done==nil{t.Fatal("nil result")};defer cancel();if child.Value(key{})!="value"{t.Error("parent value lost")};parentCancel(cause);awaitCanceled(t,done);if child.Err()!=context.Canceled||context.Cause(child)!=cause{t.Errorf("parent cancel child%v cause%v",child.Err(),context.Cause(child))}})`
- Case03: `synctest.Test(t,func(t *testing.T){parent,stop:=context.WithCancel(context.Background());stop();child,cancel,done:=subject.StartCancelable(parent);if child==nil||cancel==nil||done==nil{t.Fatal("nil result")};defer cancel();awaitCanceled(t,done);if child.Err()!=context.Canceled{t.Error("already canceled parent ignored")}})`

Realistic required defect models:

- detached-child
- complete-early
- return-parent
- cancel-immediately
- completion-value

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
