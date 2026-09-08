# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the exact prefix INPUT: followed by one space and the uppercased input.
- Use Go Unicode uppercase behavior for every letter.
- Preserve punctuation, digits and whitespace from the input.
- Input is valid UTF-8; empty input returns only the prefix and its space.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got:=subject.UpperLabel("Judy");got!="INPUT: JUDY"{t.Errorf("result got %q want %q",got,"INPUT: JUDY")}`
- Case02: `if got:=subject.UpperLabel("lars");got!="INPUT: LARS"{t.Errorf("result got %q want %q",got,"INPUT: LARS")}`
- Case03: `if got:=subject.UpperLabel("Peter-James");got!="INPUT: PETER-JAMES"{t.Errorf("result got %q want %q",got,"INPUT: PETER-JAMES")}`
- Case04: `if got:=subject.UpperLabel("MJ");got!="INPUT: MJ"{t.Errorf("result got %q want %q",got,"INPUT: MJ")}`
- Case05: `if got:=subject.UpperLabel("éж9");got!="INPUT: \u00c9\u04169"{t.Errorf("result got %q want %q",got,"INPUT: \u00c9\u04169")}`
- Case06: `if got:=subject.UpperLabel(" a ");got!="INPUT:  A "{t.Errorf("result got %q want %q",got,"INPUT:  A ")}`
- Case07: `if got:=subject.UpperLabel("");got!="INPUT: "{t.Errorf("result got %q want %q",got,"INPUT: ")}`

Realistic required defect models:

- no-uppercase
- lowercase
- ascii-only
- trimmed-input
- prefix-missing
- empty-special

Retained source behavior: welcome-to-tech-palace selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
