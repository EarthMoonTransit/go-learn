# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Complete Inc. A new or zero-value Counter starts at zero; each call adds exactly one, including simultaneous calls on the same Counter.
- Use the Counter-owned mutex to protect its shared value. Do not copy the Counter after first use. Different counters retain independent state.
- The provided Value is called only after all concurrent increments have completed in this lesson; leave its behavior unchanged.
- Public checks include the source three-increment and 1000-worker scenarios. Check and Complete run the race detector; a numerically correct result with a data race is incorrect.

Public/reference assertions (actual/expected diagnostics):

- Case01: `c:=subject.NewCounter();if c==nil||c.Value()!=0{t.Fatal("new counter must start at zero")};c.Inc();c.Inc();c.Inc();if c.Value()!=3{t.Errorf("three increments got%d want3",c.Value())}`
- Case02: `c:=subject.NewCounter();var wg sync.WaitGroup;wg.Add(1000);for range 1000{go func(){defer wg.Done();c.Inc()}()};wg.Wait();if c.Value()!=1000{t.Errorf("concurrent increments got%d want1000",c.Value())}`
- Case03: `var a,b subject.Counter;a.Inc();b.Inc();b.Inc();if a.Value()!=1||b.Value()!=2{t.Errorf("independent zero counters got%d,%d",a.Value(),b.Value())}`

Realistic required defect models:

- no-lock
- unlock-before-update
- reset-increment
- double-increment
- initial-one

Retained source behavior: lgwt:sync/v2/sync.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
