# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Choose the string that comes first under Go string comparison.
- Return selected: followed by one space and the chosen string.
- Swapping the arguments does not change the result; equal strings return that same string after the prefix.
- Empty strings are valid; preserve exact case and whitespace and compare the complete strings.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.ChooseLabel("Beta","Delta");got!="selected: Beta"{t.Errorf("result got %q want %q",got,"selected: Beta")}`
- Case02: `if got:=subject.ChooseLabel("Delta","Beta");got!="selected: Beta"{t.Errorf("result got %q want %q",got,"selected: Beta")}`
- Case03: `if got:=subject.ChooseLabel("abz","aba");got!="selected: aba"{t.Errorf("result got %q want %q",got,"selected: aba")}`
- Case04: `if got:=subject.ChooseLabel("z","A");got!="selected: A"{t.Errorf("result got %q want %q",got,"selected: A")}`
- Case05: `if got:=subject.ChooseLabel("x","x");got!="selected: x"{t.Errorf("result got %q want %q",got,"selected: x")}`
- Case06: `if got:=subject.ChooseLabel("x","");got!="selected: "{t.Errorf("result got %q want %q",got,"selected: ")}`
- Case07: `if got:=subject.ChooseLabel(" a","a");got!="selected:  a"{t.Errorf("result got %q want %q",got,"selected:  a")}`

Realistic required defect models:

- first-only
- largest
- first-byte-only
- case-insensitive
- empty-skipped
- equal-empty
- trimmed-output
- prefix-missing

Retained source behavior: vehicle-purchase selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| symmetric choice | 01–02 | first-only; largest |
| complete strings | 03 | first-byte-only |
| case-sensitive comparison | 04,08 | case-insensitive |
| equal inputs | 05 | equal-empty |
| empty input | 06 | empty-skipped |
| whitespace/prefix | 07 | trimmed-output; prefix-missing |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.
