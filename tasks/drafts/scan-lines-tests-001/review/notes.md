# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Collect lines using bufio.ScanLines semantics: remove newline and an optional preceding carriage return; include a final nonempty line without a newline.
- Preserve spaces, Unicode and interior empty lines. Empty input returns no lines; nil and empty result slices are equivalent.
- Return collected lines together with the scanner error, preserving a supplied reader error identity. Do not close src.
- Accept lines up to 2048 bytes. Reject a 10000-byte unbroken token with a non-nil error; exact buffer capacity is unspecified.

Public/reference assertions (actual/expected diagnostics):

- Case01: `v,e:=subject.ScanLines(strings.NewReader("a
b"));if e!=nil||!slices.Equal(v,[]string{"a","b"}){t.Errorf("final line got%q,%v",v,e)}`
- Case02: `for _,c:=range []struct{text string;want []string}{{"",nil},{"
",[]string{""}},{" a 

λ
",[]string{" a ","","λ"}},{strings.Repeat("x",2048),[]string{strings.Repeat("x",2048)}}}{v,e:=subject.ScanLines(strings.NewReader(c.text));if e!=nil||!slices.Equal(v,c.want){t.Errorf("lines got%q,%v want%q",v,e,c.want)}}`
- Case03: `cause:=errors.New("read failed");r:=&lineReader{data:[]byte("ok
"),err:cause};v,e:=subject.ScanLines(r);if !slices.Equal(v,[]string{"ok"})||e!=cause||r.closes!=0{t.Errorf("reader error got%q,%v closes%d",v,e,r.closes)};r=&lineReader{data:[]byte("ok"),err:io.EOF};v,e=subject.ScanLines(r);if e!=nil||!slices.Equal(v,[]string{"ok"})||r.closes!=0{t.Errorf("reader success got%q,%v closes%d",v,e,r.closes)}`
- Case04: `v,e:=subject.ScanLines(strings.NewReader(strings.Repeat("x",10000)));if e==nil||len(v)!=0{t.Errorf("oversize token got%q,%v wantempty,error",v,e)}`

Realistic required defect models:

- ignore-error
- trim-lines
- skip-empty
- tiny-buffer
- unbounded-token
- close-reader
- drop-last

Retained source behavior: None selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
