# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Register this helper once, then call acquire exactly once. On error return nil and the same error, registering no cleanup.
- On success acquire provides a non-nil resource. Return that exact resource and register one Cleanup callback; do not close it yet.
- Cleanup closes the resource once. Report a close error once with Errorf("cleanup: %v", err); successful cleanup reports nothing.
- Use Cleanup so the testing scope owns lifetime, including failure paths. Multiple acquisitions retain independent resources; cleanup executes in LIFO registration order.

Public/reference assertions (actual/expected diagnostics):

- Case01: `s:=&cleanupScope{};r:=&closeProbe{};calls:=0;got,e:=subject.AcquireForTest(s,func()(io.Closer,error){calls++;return r,nil});if got!=r||e!=nil||calls!=1||s.helpers!=1||len(s.cleanups)!=1||r.calls!=0{t.Fatalf("setup got%v,%v calls%d helpers%d cleanup%d closes%d",got,e,calls,s.helpers,len(s.cleanups),r.calls)};s.finish();if r.calls!=1||len(s.errors)!=0{t.Errorf("cleanup closes%d errors%v",r.calls,s.errors)}`
- Case02: `cause:=errors.New("acquire failed");s:=&cleanupScope{};unused:=&closeProbe{};calls:=0;got,e:=subject.AcquireForTest(s,func()(io.Closer,error){calls++;return unused,cause});if got!=nil||e!=cause||calls!=1||s.helpers!=1||len(s.cleanups)!=0||unused.calls!=0{t.Errorf("failed acquisition got%v,%v calls%d cleanups%d closes%d",got,e,calls,len(s.cleanups),unused.calls)}`
- Case03: `var order []string;cause:=errors.New("close failed");s:=&cleanupScope{};a:=&closeProbe{name:"a",order:&order};b:=&closeProbe{name:"b",order:&order,err:cause};for _,r:=range []*closeProbe{a,b}{if _,e:=subject.AcquireForTest(s,func()(io.Closer,error){return r,nil});e!=nil{t.Fatal(e)}};s.finish();if !slices.Equal(order,[]string{"b","a"})||a.calls!=1||b.calls!=1||!slices.Equal(s.errors,[]string{"cleanup: close failed"}){t.Errorf("order%v closes%d,%d errors%v",order,a.calls,b.calls,s.errors)}`

Realistic required defect models:

- close-immediately
- no-registration
- ignore-close-error
- discard-acquire-error
- double-close
- missing-helper

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
