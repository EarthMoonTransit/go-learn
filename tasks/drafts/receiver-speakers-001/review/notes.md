# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- English.Speak returns "Hello World"; (*Chinese).Speak returns "你好世界".
- SayHello calls Speak once on any supplied non-nil Speaker and returns its exact result.
- BuiltinGreetings returns English then Chinese output by using values or pointers that satisfy Speaker.
- Keep the declared receiver forms: English values and pointers qualify; Chinese requires a pointer. Calls are independent.

Public/reference assertions (actual/expected diagnostics):

- Case01: `var a subject.Speaker=subject.English{};var b subject.Speaker=&subject.English{};var c subject.Speaker=&subject.Chinese{};for _,x:=range []struct{s subject.Speaker;want string}{{a,"Hello World"},{b,"Hello World"},{c,"你好世界"}}{if got:=x.s.Speak();got!=x.want{t.Errorf("speaker got%q want%q",got,x.want)}};if got:=subject.BuiltinGreetings();got!=([2]string{"Hello World","你好世界"}){t.Errorf("builtins got%v",got)}`
- Case02: `s:=&customSpeaker{text:" exact "};if got:=subject.SayHello(s);got!=" exact "||s.calls!=1{t.Errorf("dispatch got%q calls%d",got,s.calls)};s.text="";if got:=subject.SayHello(s);got!=""||s.calls!=2{t.Errorf("current dispatch got%q calls%d",got,s.calls)}`

Realistic required defect models:

- same-language
- known-only
- double-call
- trim-result
- reverse-builtins

Retained source behavior: ardan:topics/go/language/interfaces/exercises/template1/template1.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.

The source header says both receivers are values, while its detailed body requires a Chinese pointer receiver. This adaptation follows that detailed body and preserves the bilingual outputs and interface dispatch.
