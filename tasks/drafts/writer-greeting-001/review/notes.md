# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Write exactly "Hello, " followed by name, preserving empty text, Unicode and spaces; append no newline.
- Make one Write call on the supplied non-nil writer. Return its exact error if non-nil, even with a partial count.
- If Write reports fewer bytes than requested with nil error, return io.ErrShortWrite; on a full successful write return nil.
- Do not retry, close the writer or write elsewhere. Writer counts are between zero and the supplied byte length.

Public/reference assertions (actual/expected diagnostics):

- Case01: `var b bytes.Buffer;e:=subject.WriteGreeting(&b,"Chris");if e!=nil||b.String()!="Hello, Chris"{t.Errorf("greeting got%q,%v",b.String(),e)}`
- Case02: `for _,name:=range []string{""," λ "}{w:=&writerProbe{n:-1};e:=subject.WriteGreeting(w,name);if e!=nil||w.text!="Hello, "+name||w.calls!=1{t.Errorf("name%q got%q,%v calls%d",name,w.text,e,w.calls)}}`
- Case03: `sentinel:=errors.New("write failed");w:=&writerProbe{n:2,err:sentinel};if e:=subject.WriteGreeting(w,"x");e!=sentinel||w.calls!=1{t.Errorf("failure got%v calls%d",e,w.calls)}`
- Case04: `w:=&writerProbe{n:2};if e:=subject.WriteGreeting(w,"x");e!=io.ErrShortWrite||w.calls!=1{t.Errorf("short got%v calls%d",e,w.calls)}`

Realistic required defect models:

- ignore-error
- ignore-short
- fresh-error
- extra-newline
- trim-name
- retry-write

Retained source behavior: lgwt:di/v2/di.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
