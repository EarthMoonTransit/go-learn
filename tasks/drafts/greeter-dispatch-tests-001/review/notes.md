# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return "I can speak " + g.LanguageName() + ": " + g.Greet(name).
- The non-nil Greeter may hold any implementing concrete type and state; dispatch both methods through its current dynamic value.
- Forward name exactly, including empty strings, Unicode and spaces. Preserve returned language and greeting text exactly.
- Write local implementing types in your tests and include different types and differently configured values of one type.

Public/reference assertions (actual/expected diagnostics):

- Case01: `got:=subject.SayHello("Flora",greeting{"Italian","Ciao "});want:="I can speak Italian: Ciao Flora!";if got!=want{t.Errorf("Italian got %q want %q",got,want)}`
- Case02: `for _,c:=range []struct{n string;g subject.Greeter;want string}{{"Fabrício",greeting{"Portuguese","Olá "},"I can speak Portuguese: Olá Fabrício!"},{" a ",echoGreeting{},"I can speak Echo:  a  a "},{"",greeting{"", ""},"I can speak : !"},{" x ",greeting{" L ","["},"I can speak  L : [ x !"}}{if got:=subject.SayHello(c.n,c.g);got!=c.want{t.Errorf("dispatch got %q want %q",got,c.want)}}`

Realistic required defect models:

- fixed-language
- fixed-greeting
- trim-name
- empty-name
- missing-separator
- trim-language

Retained source behavior: airport-robot selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
