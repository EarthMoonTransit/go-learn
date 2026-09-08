# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Call work once and return its integer result unchanged.
- Call release exactly once after work finishes and before WithRelease returns.
- If work panics, still call release and let the original panic continue unchanged. Do not recover it.
- Both callbacks are non-nil; release does not panic. Deferred calls run during normal return and panic unwinding.

Public/reference assertions (actual/expected diagnostics):

- Case01: `events:=[]string{};n:=subject.WithRelease(func()int{events=append(events,"work");return 7},func(){events=append(events,"release")});if n!=7||!slices.Equal(events,[]string{"work","release"}){t.Errorf("normal got%d events%v",n,events)}`
- Case02: `token:=&struct{n int}{3};events:=[]string{};func(){defer func(){if got:=recover();got!=token{t.Errorf("panic got%v want original",got)}}();subject.WithRelease(func()int{events=append(events,"work");panic(token)},func(){events=append(events,"release")})}();if !slices.Equal(events,[]string{"work","release"}){t.Errorf("panic events%v want work,release",events)}`

Realistic required defect models:

- normal-only
- release-first
- double-release
- result-erased
- double-work

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
