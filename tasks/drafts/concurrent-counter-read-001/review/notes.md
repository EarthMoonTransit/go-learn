# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Complete Value so it reads the current counter safely while other goroutines call Inc. Inc is supplied and already locks the Counter mutex.
- Use the same synchronization for reads and writes. Value must not mutate the count; it returns a value observed during that call.
- A new/zero counter reads zero. Repeated reads after all increments return the exact count; separate counters remain independent.
- The original source reads only after Wait. This extension explicitly permits overlapping reads and writes; race-free execution is required in public and full checks.

Public/reference assertions (actual/expected diagnostics):

- Case01: `var c subject.Counter;if c.Value()!=0{t.Error("zero read")};c.Inc();c.Inc();c.Inc();if c.Value()!=3||c.Value()!=3{t.Error("stable read after three increments")}`
- Case02: `c:=subject.NewCounter();var wg sync.WaitGroup;wg.Add(2);go func(){defer wg.Done();for range 1000{c.Inc()}}();go func(){defer wg.Done();for range 1000{if v:=c.Value();v<0||v>1000{t.Errorf("concurrent read %d outside0..1000",v)}}}();wg.Wait();if c.Value()!=1000||c.Value()!=1000{t.Error("final count must be stable1000")}`
- Case03: `a,b:=subject.NewCounter(),subject.NewCounter();a.Inc();if a.Value()!=1||b.Value()!=0{t.Error("counter state is shared")}`

Realistic required defect models:

- unlocked-read
- unlock-before-read
- zero-read
- read-increments
- read-clears

Retained source behavior: lgwt:sync/v2/sync.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
