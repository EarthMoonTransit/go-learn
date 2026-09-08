# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Call Helper exactly once before any assertion reporting; a *testing.T satisfies the supplied Reporter interface.
- When got equals want report nothing. Otherwise call Errorf once with format "got %d, want %d" and the two integer arguments in that order.
- Do not panic or terminate the caller. Zero and negative values are valid.
- The public recording reporter makes helper registration, ordering and diagnostics observable.

Public/reference assertions (actual/expected diagnostics):

- Case01: `r:=&recordReporter{};subject.AssertInt(r,3,3);if !slices.Equal(r.events,[]string{"helper"})||len(r.messages)!=0{t.Errorf("equal events%v messages%v",r.events,r.messages)}`
- Case02: `r:=&recordReporter{};subject.AssertInt(r,-2,0);if !slices.Equal(r.events,[]string{"helper","error"})||!slices.Equal(r.messages,[]string{"got -2, want 0"})||r.format!="got %d, want %d"||!reflect.DeepEqual(r.args,[]any{-2,0}){t.Errorf("helper diagnostic events%v messages%v format%q args%v",r.events,r.messages,r.format,r.args)};subject.AssertInt(r,7,7);if len(r.events)!=3||len(r.messages)!=1{t.Error("later equality reported")};subject.AssertInt(t,1,1)`

Realistic required defect models:

- missing-helper
- helper-after-error
- report-equality
- swapped-arguments
- wrong-format

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
