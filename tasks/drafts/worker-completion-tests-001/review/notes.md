# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Start launches non-nil work in a new goroutine and returns a non-nil completion channel without waiting for work to finish.
- Call work exactly once. Close the completion channel only after work returns; send no values on it.
- Calls own independent workers and completion channels. Callbacks return normally after any supplied gate is released.
- Use controlled channels to hold and release work, then await completion. The provided five-second guard detects deadlocks; it is not a performance target.

Public/reference assertions (actual/expected diagnostics):

- Case01: `var calls atomic.Int32;done:=subject.Start(func(){calls.Add(1)});awaitClosed(t,done);if calls.Load()!=1{t.Errorf("callback calls%d want1",calls.Load())}`
- Case02: `release:=make(chan struct{});var once sync.Once;unblock:=func(){once.Do(func(){close(release)})};defer unblock();started:=make(chan struct{});var startOnce sync.Once;var calls atomic.Int32;returned:=make(chan (<-chan struct{}),1);go func(){returned<-subject.Start(func(){calls.Add(1);startOnce.Do(func(){close(started)});<-release})}();var done <-chan struct{};select{case done=<-returned:case <-time.After(5*time.Second):t.Fatal("Start blocked waiting for callback")};if done==nil{t.Fatal("nil completion")};select{case <-started:case <-time.After(5*time.Second):t.Fatal("callback did not start")};select{case _,ok:=<-done:t.Errorf("completion before release, open=%v",ok);default:};unblock();awaitClosed(t,done);if calls.Load()!=1{t.Errorf("controlled callback calls%d want1",calls.Load())}`
- Case03: `var a,b atomic.Int32;first:=subject.Start(func(){a.Add(1)});second:=subject.Start(func(){b.Add(1)});if first==second{t.Error("calls share completion channel")};awaitClosed(t,first);awaitClosed(t,second);if a.Load()!=1||b.Load()!=1{t.Errorf("independent calls%d,%d",a.Load(),b.Load())}`

Realistic required defect models:

- inline-work
- never-close
- close-before-work
- skip-work
- double-work
- nil-channel

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.

This first concurrency lesson introduces launch/lifetime and blocking channel operations together (two concepts). Start exposes a completion channel to make asynchronous start and joining observable. The separate leak-prevention encounter moves to the subsequent directional-channel ownership lesson, before context cancellation. The testing/synctest harness uses Test/Wait to observe runnable work and blocked operations with no real clock dependency. Existing mutation-tests and race verification support it without a checker extension.
