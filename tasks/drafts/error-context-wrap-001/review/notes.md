# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return nil when err is nil, regardless of operation text.
- Otherwise return an error with text operation + ": " + err.Error(). Preserve spaces and empty operation text.
- The returned error must unwrap directly to the original err, preserving its identity.
- Repeated wrapping adds one context layer each time; do not flatten earlier layers or modify the original error.

Public/reference assertions (actual/expected diagnostics):

- Case01: `cause:=errors.New("failed");wrapped:=subject.Wrap("read",cause);if wrapped==nil||wrapped.Error()!="read: failed"||errors.Unwrap(wrapped)!=cause{t.Errorf("wrapped got%v unwrap%v",wrapped,errors.Unwrap(wrapped))}`
- Case02: `cause:=errors.New("bad");first:=subject.Wrap("",cause);second:=subject.Wrap(" outer ",first);if first==nil||second==nil||first.Error()!=": bad"||second.Error()!=" outer : : bad"||errors.Unwrap(first)!=cause||errors.Unwrap(second)!=first||cause.Error()!="bad"{t.Errorf("layers got%v,%v",first,second)}`
- Case03: `for _,op:=range []string{"","read"}{if got:=subject.Wrap(op,nil);got!=nil{t.Errorf("nil cause got%v",got)}}`

Realistic required defect models:

- text-only
- no-context
- nil-wrapped
- trim-operation
- flatten-chain

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
