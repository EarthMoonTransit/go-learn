# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Start each non-nil callback exactly once in its own goroutine. Return a non-nil completion channel promptly, even while callbacks are blocked.
- Close completion only after every callback has returned. Send no values on completion. Empty input completes without invoking anything.
- The provided reference uses WaitGroup Add before launch, Done at callback exit and Wait before close. Your tests must expose skipped, duplicate, serial or premature completion.
- Callbacks may write only their own test-owned slots; read those after testing/synctest.Wait or completion synchronization. Keep independent invocations separate and release every test-owned gate, including on failure.

Public/reference assertions (actual/expected diagnostics):

- Case01: `synctest.Test(t,func(t *testing.T){calls:=0;done:=subject.JoinAll([]func(){func(){calls++}});awaitJoined(t,done);if calls!=1{t.Errorf("single callback calls%d",calls)}})`
- Case02: `synctest.Test(t,func(t *testing.T){release:=make(chan struct{});released:=false;unblock:=func(){if !released{close(release);released=true}};defer unblock();counts:=make([]int,3);work:=make([]func(),3);for i:=range work{work[i]=func(){counts[i]++;<-release}};returned:=make(chan (<-chan struct{}),1);go func(){returned<-subject.JoinAll(work)}();synctest.Wait();var done <-chan struct{};select{case done=<-returned:default:t.Fatal("JoinAll blocked before callbacks finished")};if done==nil{t.Fatal("nil completion")};if !slices.Equal(counts,[]int{1,1,1}){t.Errorf("callbacks not independently started: %v",counts)};select{case <-done:t.Error("completion before workers returned");default:};unblock();awaitJoined(t,done);if !slices.Equal(counts,[]int{1,1,1}){t.Errorf("callback counts%v want[1 1 1]",counts)}})`
- Case03: `synctest.Test(t,func(t *testing.T){a:=subject.JoinAll(nil);b:=subject.JoinAll([]func(){});if a==b{t.Error("independent completions shared")};awaitJoined(t,a);awaitJoined(t,b)})`

Realistic required defect models:

- close-before-wait
- skip-last
- double-work
- serial-work
- nil-empty
- completion-value

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
