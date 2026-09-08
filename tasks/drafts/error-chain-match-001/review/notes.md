# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the result of Go error identity matching: direct identity, wrapped causes and any cause of a joined error qualify.
- Respect a source error implementing Is(error) bool. Equal message text alone does not establish identity.
- Two nil errors match; exactly one nil does not. Preserve errors and do not unwrap only a single level.
- Input error chains are finite and follow the standard error conventions.

Public/reference assertions (actual/expected diagnostics):

- Case01: `target:=errors.New("x");other:=errors.New("x");wrapped:=fmt.Errorf("outer: %w",fmt.Errorf("inner: %w",target));for _,c:=range []struct{err,target error;want bool}{{nil,nil,true},{target,nil,false},{nil,target,false},{target,target,true},{other,target,false},{wrapped,target,true},{errors.Join(other,wrapped),target,true},{errors.Join(other),target,false}}{if got:=subject.Matches(c.err,c.target);got!=c.want{t.Errorf("match %v/%v got%v want%v",c.err,c.target,got,c.want)}}`
- Case02: `target:=errors.New("marker");err:=&semanticError{target:target};if !subject.Matches(err,target)||subject.Matches(err,errors.New("marker")){t.Error("custom Is semantics lost")}`

Realistic required defect models:

- identity-only
- text-equality
- single-cause
- nil-never
- reverse-match

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
