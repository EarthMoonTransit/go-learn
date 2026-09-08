# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Place a line of width asterisks before the text and an identical line after it.
- Separate the three parts with exactly two added newline bytes; add no final newline.
- Preserve the complete text, including any existing newlines or whitespace.
- width is between 0 and 10000; zero width still adds both separators and empty text is valid.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.AddBorder("Welcome!",10);got!="**********\nWelcome!\n**********"{t.Errorf("result got %q want %q",got,"**********\nWelcome!\n**********")}`
- Case02: `if got:=subject.AddBorder("Hi",2);got!="**\nHi\n**"{t.Errorf("result got %q want %q",got,"**\nHi\n**")}`
- Case03: `if got:=subject.AddBorder("x",0);got!="\nx\n"{t.Errorf("result got %q want %q",got,"\nx\n")}`
- Case04: `if got:=subject.AddBorder("",1);got!="*\n\n*"{t.Errorf("result got %q want %q",got,"*\n\n*")}`
- Case05: `if got:=subject.AddBorder(" a\nb ",2);got!="**\n a\nb \n**"{t.Errorf("result got %q want %q",got,"**\n a\nb \n**")}`

Realistic required defect models:

- extra-newline
- missing-lower
- wrong-width
- trimmed-text
- zero-bypass
- empty-bypass

Retained source behavior: welcome-to-tech-palace selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| both lines/exact width | 01–02 | missing-lower; wrong-width |
| exact inserted newlines | 01–02 | extra-newline |
| zero width/empty text | 03–04 | zero-bypass; empty-bypass |
| middle text preserved | 05 | trimmed-text |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.
