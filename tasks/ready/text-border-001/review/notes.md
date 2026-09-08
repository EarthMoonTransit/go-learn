# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

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


## Independent review

Public checks distinguish exact border width, both newline separators, absence of an added final newline, zero width, empty input and preservation of existing whitespace/newlines.

Reviewed the actual task contract, starter, public/reference assertions, all 6 required defects and route dependencies. The declared operations are substantive permitted adaptations. Exact source artifact hashes and verbatim MIT notices match the pinned local source checkouts, including Exercism authors/contributors where supplied and the literal LGWT copyright text. Private reference/mutant/review material stays outside learner payloads.

The author-provided targeted report passes, with the strong reference detecting all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full batch gates, allocation reconciliation and T07 acceptance remain separate. Automated fixtures are not learner calibration.
